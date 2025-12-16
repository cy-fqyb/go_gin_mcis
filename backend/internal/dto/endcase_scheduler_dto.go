package dto

type EndcaseSchedulerDTO struct {
	LoginNo  string                  `json:"login_no" binding:"required" label:"登录号"`
	Password string                  `json:"password" binding:"required" label:"密码"`
	Data     EndcaseSchedulerDtoData `json:"data" binding:"required"`
}

type EndcaseSchedulerDtoData struct {
	BeginDate string `json:"begin_date" binding:"required" label:"开始日期"`
	EndDate   string `json:"end_date" binding:"required" label:"结束日期"`
}
