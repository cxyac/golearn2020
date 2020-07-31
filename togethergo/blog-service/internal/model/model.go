package model
//公共model 一些数据创建和更新删除信息
type Model struct {
	ID uint32 `gorm:"primary_key" json:"id"`
	CreateBy string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeleteOn uint32 `json:"delete_on"`
	IsDel uint8 `json:"is_del"`
}

