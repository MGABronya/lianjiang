// @Title  RowAll
// @Description  定义点集描述信息
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package model

import (
	"time"
)

// RowAll			定义点集描述信息
type RowAll struct {
	StartTime                       time.Time `json:"start_time" gorm:"type:timestamp;primary_key;autoIncrement:false"`  // 记录的开始日期
	EndTime                         time.Time `json:"end_time" gorm:"type:timestamp;primary_key;autoIncrement:false"`    // 记录的终止日期
	StationName                     string    `json:"station_name" gorm:"type:char(50);primary_key;autoIncrement:false"` // 站点名称
	Cod                             string    `json:"Cod" gorm:"type:char(50);"`                                         // 化学需氧量
	FiveDaysNiochemicalOxygenDemand string    `json:"five_days_niochemicaloxygendemand" gorm:"type:char(50);"`           // 五日生化需氧量
	Se                              string    `json:"Se" gorm:"type:char(50);"`                                          // 硒
	As                              string    `json:"As" gorm:"type:char(50);"`                                          // 砷
	Hg                              string    `json:"Hg" gorm:"type:char(50);"`                                          // 汞
	Fluoride                        string    `json:"fluoride" gorm:"type:char(50);"`                                    // 氟化物
	Petroleum                       string    `json:"petroleum" gorm:"type:char(50);"`                                   // 石油类
	FecalColiform                   string    `json:"fecal_coliform" gorm:"type:char(50);"`                              // 粪大肠菌群
	Temperature                     string    `json:"temperature" gorm:"type:char(50);"`                                 // 水温
	PH                              string    `json:"pH" gorm:"type:char(50);"`                                          // pH值
	DO                              string    `json:"DO" gorm:"type:char(50);"`                                          // 溶解氧
	EC                              string    `json:"EC" gorm:"type:char(50);"`                                          // 电导率
	Turbidity                       string    `json:"turbidity" gorm:"type:char(50);"`                                   // 浊度
	CODMII                          string    `json:"CODMII" gorm:"type:char(50);"`                                      // 高锰酸盐指数
	NH_N                            string    `json:"NH_N" gorm:"type:char(50);"`                                        // 氨氮
	TP                              string    `json:"TP" gorm:"type:char(50);"`                                          // 总磷
	TN                              string    `json:"TN" gorm:"type:char(50);"`                                          // 总氮
	CODcr                           string    `json:"CODcr" gorm:"type:char(50);"`                                       // CODcr
	CN                              string    `json:"CN" gorm:"type:char(50);"`                                          // 氰化物
	VolatilePenol                   string    `json:"volatile_penol" gorm:"type:char(50);"`                              // 挥发酚
	Cr                              string    `json:"Cr" gorm:"type:char(50);"`                                          // 六价铬
	Cu                              string    `json:"Cu" gorm:"type:char(50);"`                                          // 铜
	Zn                              string    `json:"Zn" gorm:"type:char(50);"`                                          // 锌
	Pb                              string    `json:"Pb" gorm:"type:char(50);"`                                          // 铅
	Cd                              string    `json:"Cd" gorm:"type:char(50);"`                                          // 镉
	LAS                             string    `json:"LAS" gorm:"type:char(50);"`                                         // 阴离子表面活性剂
	SOx                             string    `json:"SOx" gorm:"type:char(50);"`                                         // 硫化物
	CumulativeDischarge             string    `json:"cumulative_discharge" gorm:"type:char(50);"`                        // 累计流量
	WaterDischarge                  string    `json:"water_discharge" gorm:"type:char(50);"`                             // 水流量
	WaterLevel                      string    `json:"water_level" gorm:"type:char(50);"`                                 // 水位
	PeriodCumulativeFlow            string    `json:"period_cumulative_flow" gorm:"type:char(50);"`                      // 时段累积流量
	SectionalMeanVelocity           string    `json:"sectional_mean_velocity" gorm:"type:char(50);"`                     // 断面平均流速
	SectionalArea                   string    `json:"sectional_area" gorm:"type:char(50);"`                              // 断面面积
	TotalCumulativeFlow             string    `json:"total_cumulativeflow" gorm:"type:char(50);"`                        // 总累积流量
	CurrentInstantaneousVelocity    string    `json:"current_instantaneous_velocity" gorm:"type:char(50);"`              // 当前瞬时流速
	InstantaneousDelivery           string    `json:"instantaneous_delivery" gorm:"type:char(50);"`                      // 瞬时流量
	CreatedAt                       Time      `json:"created_at" gorm:"type:timestamp"`                                  // 记录的存入日期
}
