package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go_gin_mcis/internal/dto"
	"go_gin_mcis/internal/model"
	"go_gin_mcis/pkg/db"
	"go_gin_mcis/pkg/logger"
	"go_gin_mcis/pkg/middleware"
	"go_gin_mcis/pkg/result"
	"io"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	RegisterPrivateRoutes(func(r *gin.RouterGroup) {
		r.POST("/endcase_scheduler", middleware.AutoBind(EndcaseScheduler, reflect.TypeOf(dto.EndcaseSchedulerDTO{})))
	})
}

// APIResponse 接口响应外层结构
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

func EndcaseScheduler(c *gin.Context, data any) {
	schedulerDTO := data.(*dto.EndcaseSchedulerDTO)
	logger.Infof("Received EndcaseScheduler request: %+v", schedulerDTO)
	// 发送post请求向接口获取数据 http://www.wuyouyunyu.com:8090/mcis-web/api/woman/getEndCase.do
	body := map[string]any{
		"login_no": schedulerDTO.LoginNo,
		"password": schedulerDTO.Password,
		"data": map[string]string{
			"begin_date": schedulerDTO.Data.BeginDate,
			"end_date":   schedulerDTO.Data.EndDate,
		},
	}
	var apiResp APIResponse
	logger.Infof("Request body for external API: %+v", body)
	jsonData, _ := json.Marshal(body)
	resp, err := http.Post(
		"http://www.wuyouyunyu.com:8090/mcis-web/api/woman/getEndCase.do",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		logger.Errorf("Error requesting external API: %s", err.Error())
		result.Fail(c, http.StatusInternalServerError, "请求远程服务失败")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Errorf("close body err: %s", err.Error())
		}
	}(resp.Body)
	// 读取响应内容
	res, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(res, &apiResp)
	if err != nil {
		logger.Errorf("Error parsing external API response: %s", err.Error())
		result.Fail(c, http.StatusInternalServerError, "解析远程服务响应失败")
		return
	}
	if !apiResp.Success {
		result.Fail(c, http.StatusBadRequest, "远程服务返回失败状态")
		return
	}
	// 循环遍历数据进行状态结案
	for _, woman := range apiResp.Data {
		if woman.IsEnd == 1 {
			err := SaveEndcase(db.DB, woman)
			if err != nil {
				logger.Errorf("Failed to save endcase for ID %d: %s", woman.ID, err)
			}
		}
	}

}
func SaveEndcase(db *gorm.DB, woman PregnantWoman) error {
	var wmID int

	// 1️⃣ 先查 jj_upload_record
	err := db.Model(&model.JJUploadRecord{}).
		Select("rowId").
		Where("type = ? AND outId = ?", "wcbj_basic", woman.ID).
		Scan(&wmID).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 2️⃣ 如果没查到，再查 gen_business 表
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

	// 3️⃣ 如果找到了，就更新
	if wmID != 0 {
		err = db.Model(&model.GenBusiness{}).
			Where("id = ?", wmID).
			Updates(map[string]any{
				"end_way":        woman.EndTypeName,
				"end_way_name":   woman.EndTypeName,
				"end_way_time":   woman.EndCaseDate,
				"end_way_remark": "平台自动同步",
			}).Error

		if err != nil {
			return fmt.Errorf("更新失败: %w", err)
		}

		logger.Infof("✅ 成功更新结案记录，wm_id=%d", wmID)
	} else {
		logger.Errorf("⚠️ 未找到匹配的孕妇记录: %+v", woman)
	}

	return nil
}
