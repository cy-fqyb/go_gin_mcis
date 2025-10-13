package model

type VwUploadRecord struct {
	RecordId    int    `json:"record_id" gorm:"column:record_id"`
	UploadType  string `json:"upload_type" gorm:"column:upload_type"`
	IdCardNo    string `json:"id_card_no" gorm:"column:id_card_no"`
	Name        string `json:"name" gorm:"column:name"`
	SuccessFlag string `json:"success_flag" gorm:"column:success_flag"`
	ErrorMsg    string `json:"error_msg" gorm:"column:error_msg"`
	UploadMsg   string `json:"upload_msg" gorm:"column:upload_msg"`
	UploadTime  string `json:"upload_time" gorm:"column:upload_time"`
	ReturnMsg   string `json:"return_msg" gorm:"column:return_msg"`
}
