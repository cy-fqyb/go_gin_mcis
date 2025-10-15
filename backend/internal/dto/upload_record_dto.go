// internal/dto/upload_record_query.go
package dto

// UploadRecordQuery 用于接收前端查询条件
type UploadRecordQuery struct {
	RecordId   *int    `json:"record_id"`
	UploadType *string `json:"upload_type"`
	Name       *string `json:"Name"`
	IdCardNo   *string `json:"idCardNo"`
	Status     *int    `json:"status"`
	UploadTime *string `json:"uploadTime"`
	EndTime    *string `json:"endTime"`
	Limit      int     `json:"limit"`
	Offset     int     `json:"offset"`
}
