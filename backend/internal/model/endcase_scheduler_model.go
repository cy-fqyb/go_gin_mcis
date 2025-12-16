package model

type JJUploadRecord struct {
	RowID int    `gorm:"column:rowId"`
	Type  string `gorm:"column:type"`
	OutID int    `gorm:"column:outId"`
}

func (JJUploadRecord) TableName() string {
	return "jj_upload_record"
}

type GenBusiness struct {
	ID            int    `gorm:"column:id"`
	WmName        string `gorm:"column:wm_name"`
	WmIdentityNo  string `gorm:"column:wm_identityno"`
	MenstrualLast string `gorm:"column:menstrual_last"`
	EndWay        string `gorm:"column:end_way"`
	EndWayName    string `gorm:"column:end_way_name"`
	EndWayTime    string `gorm:"column:end_way_time"`
	EndWayRemark  string `gorm:"column:end_way_remark"`
}

func (GenBusiness) TableName() string {
	return "gen_business_1570758835271_tab"
}
