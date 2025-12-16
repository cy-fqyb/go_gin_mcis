package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go_gin_mcis/config"
	"go_gin_mcis/internal/model"
	"go_gin_mcis/pkg/logger"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

// ------------------------ 数据结构 ------------------------

type APIResponse struct {
	Success bool            `json:"success"`
	Code    int             `json:"code"`
	Data    []PregnantWoman `json:"data"`
}

type PregnantWoman struct {
	ID              int    `json:"ID"`
	WMName          string `json:"WM_NAME"`
	WMIDNo          string `json:"WM_IDNO"`
	WMLastMenstrual string `json:"WM_MENSTRUAL_LAST"`
	IsEnd           int    `json:"IS_END"`
	EndCaseDate     string `json:"END_CASE_DATE"`
	EndTypeCode     any    `json:"END_TYPE_CODE"`
	EndTypeName     string `json:"END_TYPE_NAME"`
}

// ------------------------ 定时任务结构 ------------------------

type EndCaseScheduler struct {
	cron       *cron.Cron
	db         *gorm.DB
	jobID      cron.EntryID
	running    bool
	jobTarget  string
	apiURL     string
	mu         sync.Mutex
	client     *http.Client
	runningJob bool // 关键：防止任务重入
}

func NewEndCaseScheduler(db *gorm.DB, apiURL string) *EndCaseScheduler {
	return &EndCaseScheduler{
		cron:   cron.New(cron.WithSeconds()),
		db:     db,
		apiURL: apiURL,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

// Start 启动任务
func (s *EndCaseScheduler) Start(expr string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.running {
		if s.jobID != 0 {
			s.cron.Remove(s.jobID)
			logger.Infof("♻️ 移除旧任务 (jobID: %d)", s.jobID)
		}
	}

	id, err := s.cron.AddFunc(expr, s.runJob)
	if err != nil {
		return err
	}
	s.jobID = id

	if !s.running {
		s.cron.Start()
		s.running = true
	}

	logger.Infof("🕓 调度器启动 (cron: %s, jobID: %d)", expr, s.jobID)
	return nil
}

// ------------------------ 主任务逻辑 ------------------------

func (s *EndCaseScheduler) runJob() {
	start := time.Now()

	// =============================
	// 🔒 防止任务重入：非阻塞判断
	// =============================
	s.mu.Lock()
	if s.runningJob {
		logger.Infof("⚠️ 上一次任务尚未完成，本次执行跳过")
		s.mu.Unlock()
		return
	}
	s.runningJob = true
	s.mu.Unlock()

	logger.Info("🚀 Running scheduled job: syncing end-case data...")

	// 执行完释放锁
	defer func() {
		s.mu.Lock()
		s.runningJob = false
		s.mu.Unlock()
		logger.Infof("✅ Job completed in %v", time.Since(start))
	}()

	// ------------------- 原逻辑不动 -------------------

	beginDate := config.GetConf().EndCaseServer.BeginDate
	endDate := config.GetConf().EndCaseServer.EndDate
	if beginDate == "" {
		beginDate = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	}
	if endDate == "" {
		endDate = time.Now().Format("2006-01-02")
	}

	body := map[string]any{
		"login_no": config.GetConf().EndCaseServer.LoginNo,
		"password": config.GetConf().EndCaseServer.Password,
		"data": map[string]string{
			"begin_date": beginDate,
			"end_date":   endDate,
		},
	}

	jsonData, _ := json.Marshal(body)
	logger.Infof("请求外部接口参数: %+v", body)

	resp, err := s.client.Post(s.apiURL, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		logger.Errorf("接口请求失败: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Errorf("接口返回状态码异常: %d", resp.StatusCode)
		return
	}

	resBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("响应读取失败: %s", err.Error())
		return
	}

	var apiResp APIResponse
	if err := json.Unmarshal(resBytes, &apiResp); err != nil {
		snippet := string(resBytes)
		if len(snippet) > 400 {
			snippet = snippet[:400] + "..."
		}
		logger.Errorf("接口响应解析失败: %s\n响应片段: %s", err.Error(), snippet)
		return
	}

	if !apiResp.Success {
		logger.Errorf("接口返回失败状态: %+v", apiResp)
		return
	}

	var recordNotFoundCount int64
	for _, woman := range apiResp.Data {
		if woman.IsEnd == 1 {
			err := SaveEndcase(s.db, woman, &recordNotFoundCount)
			if err != nil {
				logger.Errorf("保存结案记录失败，ID=%d: %s", woman.ID, err.Error())
			}
		}
	}
}

// ------------------------ 数据保存逻辑 ------------------------

func SaveEndcase(db *gorm.DB, woman PregnantWoman, recordNotFoundCount *int64) error {
	var wmID int

	err := db.Model(&model.JJUploadRecord{}).
		Select("rowId").
		Where("type = ? AND outId = ?", "wcbj_basic", woman.ID).
		Scan(&wmID).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if wmID == 0 {
		err = db.Model(&model.GenBusiness{}).
			Select("id").
			Where("wm_name = ? AND LOWER(wm_identityno) = LOWER(?) AND menstrual_last = ? AND end_way = ''",
				woman.WMName, woman.WMIDNo, woman.WMLastMenstrual).
			Scan(&wmID).Error

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if wmID != 0 {
		err = db.Model(&model.GenBusiness{}).
			Where("id = ?", wmID).
			Updates(map[string]any{
				"end_way":        woman.EndTypeName,
				"end_way_name":   woman.EndTypeName,
				"end_way_time":   woman.EndCaseDate,
				"end_way_remark": "平台自动同步",
				"update_time":    time.Now().Format("2006-01-02 15:04:05"),
			}).Error

		if err != nil {
			return fmt.Errorf("更新失败: %w", err)
		}

		logger.Infof("✅ 成功更新结案记录，wm_id=%d，wm_identityno=%s", wmID, woman.WMIDNo)
	} else {
		*recordNotFoundCount = *recordNotFoundCount + 1
	}

	return nil
}

// Stop 停止任务
func (s *EndCaseScheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return
	}

	if s.jobID != 0 {
		s.cron.Remove(s.jobID)
		logger.Infof("🛑 已移除任务 (jobID: %d)", s.jobID)
		s.jobID = 0
	}

	if len(s.cron.Entries()) == 0 {
		s.cron.Stop()
		s.running = false
		logger.Info("🛑 EndCaseScheduler stopped")
	}
}

// Reload 重启任务
func (s *EndCaseScheduler) Reload(newExpr string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	logger.Infof("🔁 Reloading EndCaseScheduler with new cron: %s", newExpr)

	if s.jobID != 0 {
		s.cron.Remove(s.jobID)
		logger.Infof("♻️ 移除旧任务 (jobID: %d)", s.jobID)
	}

	id, err := s.cron.AddFunc(newExpr, s.runJob)
	if err != nil {
		return err
	}
	s.jobID = id

	if !s.running {
		s.cron.Start()
		s.running = true
	}

	logger.Infof("✅ EndCaseScheduler reloaded (cron: %s, jobID: %d)", newExpr, s.jobID)
	return nil
}
