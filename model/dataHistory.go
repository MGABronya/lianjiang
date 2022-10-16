// @Title  dataHistory
// @Description  定义数据操作历史记录
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package model

// DataHistory			定义数据操作历史记录
type DataHistory struct {
	UserId      uint   `json:"user_id" gorm:"type:uint;not null"`             // 用户Id
	CreatedAt   Time   `json:"create_at" gorm:"type:timestamp;not null"`      // 操作时间
	StartTime   string `json:"start_time" gorm:"type:varchar(50);not null"`   // 起始时间
	EndTime     string `json:"end_time" gorm:"type:varchar(50);not null"`     // 终止时间
	StationName string `json:"station_name" gorm:"type:varchar(50);not null"` // 站名
	System      string `json:"system" gorm:"type:varchar(50);not null"`       // 制度
	Option      string `json:"option" gorm:"type:varchar(20);not null;"`      // 操作方法
	Time        string `json:"time" gorm:"type:varchar(20);not null;"`        // 参考时间
}
