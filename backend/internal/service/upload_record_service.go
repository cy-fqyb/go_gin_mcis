// internal/service/upload_record_service.go
package service

import (
	"go_gin_mcis/internal/dto"
	"go_gin_mcis/internal/model"
	"go_gin_mcis/pkg/db"
)

// GetUploadRecord 普通查询
func GetUploadRecord(query dto.UploadRecordQuery) ([]model.VwUploadRecord, error) {
	var records []model.VwUploadRecord
	tx := db.DB.Model(&model.VwUploadRecord{})

	if query.RecordId != nil {
		tx = tx.Where("record_id = ?", *query.RecordId)
	}
	if query.Name != nil {
		tx = tx.Where("name LIKE ?", "%"+*query.Name+"%")
	}
	if query.Status != nil {
		tx = tx.Where("status = ?", *query.Status)
	}
	if query.UploadTime != nil {
		tx = tx.Where("upload_time >= ?", *query.UploadTime)
	}
	if query.EndTime != nil {
		tx = tx.Where("end_time <= ?", *query.EndTime)
	}
	if query.Limit > 0 {
		tx = tx.Limit(query.Limit)
	}
	if query.Offset > 0 {
		tx = tx.Offset(query.Offset)
	}

	if err := tx.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

// GetUploadRecordByType 按类型动态查询
func GetUploadRecordByType(recordType string, query dto.UploadRecordQuery) ([]model.VwUploadRecord, error) {
	type joinConfig struct {
		uploadType string
		joins      []string
	}

	configs := map[string]joinConfig{
		"wcbj_basic": {
			uploadType: "基本信息",
			joins: []string{
				"JOIN gen_business_1570758835271_tab b ON r.rowId = b.id",
			},
		},
		"wcbj_shoucha": {
			uploadType: "首查记录",
			joins: []string{
				"JOIN gen_business_1611718425241_tab a ON a.id = r.rowId",
				"JOIN gen_business_1570758835271_tab b ON a.wm_id = b.id",
			},
		},
		"wcbj_fucha": {
			uploadType: "复查记录",
			joins: []string{
				"JOIN gen_business_1570764820027_tab f ON r.rowId = f.id",
				"JOIN gen_business_1570758835271_tab b ON f.wm_id = b.id",
			},
		},
		"wcbj_wm_risk": {
			uploadType: "高危记录",
			joins: []string{
				"JOIN wm_factor_tab w ON r.rowId = w.id",
				"JOIN gen_business_1570758835271_tab b ON w.wm_id = b.id",
			},
		},
		"wcbj_fenmian": {
			uploadType: "分娩记录",
			joins: []string{
				"JOIN gen_business_1571363457629_tab m ON r.rowId = m.id",
				"JOIN gen_business_1570758835271_tab b ON m.wm_id = b.id",
			},
		},
	}

	cfg, ok := configs[recordType]
	if !ok {
		return nil, nil
	}

	querySQL := `
		r.id AS record_id,
		? AS upload_type,
		b.wm_identityno AS id_card_no,
		b.wm_name AS name,
		CASE WHEN r.success IS NOT NULL AND r.error IS NULL THEN '1' ELSE '0' END AS success_flag,
		r.error AS error_msg,
		r.returnData AS return_msg,
		r.uploadData AS upload_msg,
		r.uploadTime AS upload_time
	`

	dbConn := db.DB.Table("jj_upload_record AS r").Select(querySQL, cfg.uploadType)
	for _, j := range cfg.joins {
		dbConn = dbConn.Joins(j)
	}

	dbConn = dbConn.Where("r.type = ?", recordType)

	// 动态条件
	if query.RecordId != nil {
		dbConn = dbConn.Where("r.id = ?", *query.RecordId)
	}
	if query.Name != nil && *query.Name != "" {
		dbConn = dbConn.Where("b.wm_name LIKE ?", "%"+*query.Name+"%")
	}
	if query.IdCardNo != nil && *query.IdCardNo != "" {
		dbConn = dbConn.Where("b.wm_identityno = ?", *query.IdCardNo)
	}
	if query.Status != nil {
		switch *query.Status {
		case 1:
			dbConn = dbConn.Where("r.success IS NOT NULL AND r.error IS NULL")
		case 0:
			dbConn = dbConn.Where("r.error IS NOT NULL")
		}
	}
	if query.UploadTime != nil && *query.UploadTime != "" {
		dbConn = dbConn.Where("r.uploadTime >= ?", *query.UploadTime)
	}
	if query.EndTime != nil && *query.EndTime != "" {
		dbConn = dbConn.Where("r.uploadTime <= ?", *query.EndTime)
	}

	// 分页
	if query.Limit == 0 {
		query.Limit = 20
	}
	dbConn = dbConn.Limit(query.Limit).Offset(query.Offset)

	var records []model.VwUploadRecord
	if err := dbConn.Scan(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
