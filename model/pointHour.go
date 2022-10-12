// @Title  pointHour
// @Description  定义小时制点集
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// PointHour			定义小时制点集
type PointHour struct {
	ID                           uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`                     // 点的id
	Time                         time.Time `json:"time" gorm:"type:timestamp;index:idx_name,unique"`        // 记录的创建日期
	StationName                  string    `json:"station_name" gorm:"type:char(36);index:idx_name,unique"` // 站点名称
	Temperature                  float64   `json:"temperature" gorm:"type:double;"`                         // 水温
	PH                           float64   `json:"pH" gorm:"type:double;"`                                  // pH值
	DO                           float64   `json:"DO" gorm:"type:double;"`                                  // 溶解氧
	EC                           float64   `json:"EC" gorm:"type:double;"`                                  // 电导率
	Turbidity                    float64   `json:"turbidity" gorm:"type:double;"`                           // 浊度
	CODMII                       float64   `json:"CODMII" gorm:"type:double;"`                              // 高锰酸盐指数
	NH_N                         float64   `json:"NH_N" gorm:"type:double;"`                                // 氨氮
	TP                           float64   `json:"TP" gorm:"type:double;"`                                  // 总磷
	TN                           float64   `json:"TN" gorm:"type:double;"`                                  // 总氮
	CODcr                        float64   `json:"CODcr" gorm:"type:double;"`                               // CODcr
	CN                           float64   `json:"CN" gorm:"type:double;"`                                  // 氰化物
	VolatilePenol                float64   `json:"volatile_penol" gorm:"type:double;"`                      // 挥发酚
	Cr                           float64   `json:"Cr" gorm:"type:double;"`                                  // 六价铬
	Cu                           float64   `json:"Cu" gorm:"type:double;"`                                  // 铜
	Zn                           float64   `json:"Zn" gorm:"type:double;"`                                  // 锌
	Pb                           float64   `json:"Pb" gorm:"type:double;"`                                  // 铅
	Cd                           float64   `json:"Cd" gorm:"type:double;"`                                  // 镉
	LAS                          float64   `json:"LAS" gorm:"type:double;"`                                 // 阴离子表面活性剂
	SOx                          float64   `json:"SOx" gorm:"type:double;"`                                 // 硫化物
	CumulativeDischarge          float64   `json:"cumulative_discharge" gorm:"type:double;"`                // 累计流量
	WaterDischarge               float64   `json:"water_discharge" gorm:"type:double;"`                     // 水流量
	WaterLevel                   float64   `json:"water_level" gorm:"type:double;"`                         // 水位
	PeriodCumulativeFlow         float64   `json:"period_cumulative_flow" gorm:"type:double;"`              // 时段累积流量
	SectionalMeanVelocity        float64   `json:"sectional_mean_velocity" gorm:"type:double;"`             // 断面平均流速
	SectionalArea                float64   `json:"sectional_area" gorm:"type:double;"`                      // 断面面积
	TotalCumulativeFlow          float64   `json:"total_cumulativeflow" gorm:"type:double;"`                // 总累积流量
	CurrentInstantaneousVelocity float64   `json:"current_instantaneous_velocity" gorm:"type:double;"`      // 当前瞬时流速
	InstantaneousDelivery        float64   `json:"instantaneous_delivery" gorm:"type:double;"`              // 瞬时流量
	CreatedAt                    Time      `json:"created_at" gorm:"type:timestamp"`                        // 记录的创建日期
}

// @title    BeforeCreate
// @description   计算出一个uuid
// @auth      MGAronya（张健）             2022-9-16 10:19
// @param     scope *gorm.Scope
// @return    error
func (pointHour *PointHour) BeforeCreate(scope *gorm.DB) error {
	pointHour.ID = uuid.NewV4()
	return nil
}
