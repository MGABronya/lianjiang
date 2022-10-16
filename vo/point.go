// @Title  Point
// @Description  用于定义传回点
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package vo

import "time"

// Point			定义传回的点
type Point struct {
	Time  time.Time `json:"time"`  // 时间
	Field float64   `json:"field"` // 字段值
}
