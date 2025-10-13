// internal/service/upload_record_service.go
package service

import (
	"go_gin_mcis/internal/dto"
	"go_gin_mcis/internal/model"
	"go_gin_mcis/pkg/db"
)

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
		tx = tx.Where("created_at >= ?", *query.UploadTime)
	}
	if query.EndTime != nil {
		tx = tx.Where("created_at <= ?", *query.EndTime)
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
