// @Title  mapBackup
// @Description  定义映射备份
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package model

// MapBackup			定义映射操作历史记录
type MapBackup struct {
	CreatedAt Time   `json:"create_at" gorm:"type:timestamp;not null"` // 操作时间
	Table     string `json:"table" gorm:"type:varchar(50);not null"`   // 映射表
	Key       string `json:"key" gorm:"type:varchar(50);not null"`     // 主键
	Value     string `json:"value" gorm:"type:varchar(50);not null"`   // 对应值
}
