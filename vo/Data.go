// @Title  Data
// @Description  用于定义行字段传回点
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package vo

import "time"

// Data			定义传回的点
type Data struct {
	StartTime time.Time `json:"start_time"` // 初始时间
	EndTime   time.Time `json:"end_time"`   // 初始时间
	Field     string    `json:"field"`      // 字段值
}
