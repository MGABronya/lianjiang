// @Title  personal_update
// @Description  user的个人基本信息
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:47
package vo

import "gorm.io/gorm"

// PersonalChange			个人基本信息，在更新个人信息时使用
type PersonalChange struct {
	gorm.Model        // gorm模板
	Name       string `json:"Newname" gorm:"type:varchar(20);"`        // 名称
	Email      string `json:"Newemail" gorm:"type:varchar(50);unique"` // 邮箱
	Telephone  string `json:"Newtelephone" gorm:"type:varchar(20)"`    // 电话
	Blog       string `json:"Newblog" gorm:"type:varchar(25)"`         // 博客
	QQ         string `json:"Newqq" gorm:"type:varchar(20)"`           // QQ
	Sex        bool   `json:"Newsex" gorm:"type:boolean"`              // 性别
	Address    string `json:"Newaddress" gorm:"type:varchar(20)"`      // 地址
	Hobby      string `json:"Newhobby" gorm:"type:varchar(50)"`        // 爱好
	Verify     string `json:"Verify"`                                  // 验证码
}
