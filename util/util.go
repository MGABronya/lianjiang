// @Title  util
// @Description  收集各种需要使用的工具函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:47
package util

import (
	"context"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"lianjiang/common"
	"lianjiang/model"
	"math"
	"math/rand"
	"net/smtp"
	"path"
	"regexp"
	"strconv"
	"time"

	"os"

	"github.com/tealeg/xlsx"

	"github.com/extrame/xls"

	"github.com/jordan-wright/email"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 点集字段映射表
var PointMap = map[string]string{
	"监测断面":     "StationName",
	"监测指标":     "Time",
	"监测时间":     "Time",
	"时间":       "Time",
	"水温":       "Temperature",
	"pH":       "PH",
	"化学需氧量":    "Cod",
	"五日生化需氧量":  "FiveDaysNiochemicalOxygenDemand",
	"硒":        "Se",
	"砷":        "As",
	"汞":        "Hg",
	"氟化物":      "Fluoride",
	"石油类":      "Petroleum",
	"粪大肠菌群":    "FecalColiform",
	"溶解氧":      "DO",
	"电导率":      "EC",
	"浊度":       "Turbidity",
	"高锰酸盐指数":   "CODMII",
	"氨氮":       "NH_N",
	"总磷":       "TP",
	"总氮":       "TN",
	"CODcr":    "CODcr",
	"氰化物":      "CN",
	"挥发酚":      "VolatilePenol",
	"六价铬":      "Cr",
	"铜":        "Cu",
	"锌":        "Zn",
	"铅":        "Pb",
	"镉":        "Cd",
	"阴离子表面活性剂": "LAS",
	"硫化物":      "SOx",
	"累计流量":     "CumulativeDischarge",
	"水流量":      "WaterDischarge",
	"总累积流量":    "TotalCumulativeFlow",
	"水位":       "WaterLevel",
	"时段累积流量":   "PeriodCumulativeFlow",
	"断面平均流速":   "SectionalMeanVelocity",
	"当前瞬时流速":   "CurrentInstantaneousVelocity",
	"瞬时流量":     "InstantaneousDelivery",
	"断面面积":     "SectionalArea",
}

// 行唯一字段映射表
var RowOneMap = map[string]string{
	"水质类别":  "water_quality_classification",
	"主要污染物": "key_pollutant",
}

// 行多字段映射表
var RowAllMap = map[string]string{
	"分项类别": "item_category",
}

// 制度映射表
var SysMap = map[string]string{
	"小时制": "hour",
	"月度制": "month",
}

// 文件内容的标记点映射表
var OptMap = map[string]string{
	"hour":  "时间",
	"month": "监测断面",
}

// 站名注册表
var StationMap = map[string]string{
	"海门湾桥闸":  "haimen_bay_bridge_gate",
	"汕头练江水站": "lian_jiang_water_station",
	"青洋山桥":   "lian_jiang_water_station",
	"新溪西村":   "xinxi_village",
	"万兴桥":    "wanxing_bridge",
	"流仙学校":   "liuxian_school",
	"仙马闸":    "xianma_brake",
	"华侨学校":   "huaqiao_school",
	"港洲桥":    "gangzhou_bridge",
	"云陇":     "yunlong",
	"北港水":    "beixiangshui",
	"官田水":    "guantianshui",
	"北港河闸":   "beixiang_penstock",
	"峡山大溪":   "xiashan_stream",
	"井仔湾闸":   "jingzai_wan_sluice",
	"东北支流":   " northeast_branch",
	"西埔桥闸":   "xipu_bridge_sluice",
	"五福桥":    "wufu_bridge",
	"成田大寮":   "narita_daliao",
	"新坛港":    "xitan_port",
	"瑶池港":    "yaochi_port",
	"护城河闸":   "moat_locks",
	"和平桥":    "peace_bridge",
}

// 数据注册表
var DataMap = map[string]float64{}

// @title    Read
// @description   读取文件内容
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     file_path string		文件位置
// @return    res [][]string, err error		res为读出的内容，err为可能出现的错误
func Read(file_path string) (res [][]string, err error) {

	extName := path.Ext(file_path)

	if extName == ".csv" {
		return ReadCsv(file_path)
	} else if extName == ".xls" {
		return ReadXls(file_path)
	} else if extName == ".xlsx" {
		return ReadXlsx(file_path)
	}
	return nil, nil
}

// @title    ReadCsv
// @description   读取Csv文件内容
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     file_path string		文件位置
// @return    res [][]string, err error		res为读出的内容，err为可能出现的错误
func ReadCsv(file_path string) (res [][]string, err error) {
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// TODO 初始化csv-reader
	reader := csv.NewReader(file)
	// TODO 设置返回记录中每行数据期望的字段数，-1 表示返回所有字段
	reader.FieldsPerRecord = -1
	// TODO 允许懒引号（忘记遇到哪个问题才加的这行）
	reader.LazyQuotes = true
	// TODO 返回csv中的所有内容
	records, read_err := reader.ReadAll()
	if read_err != nil {
		return nil, read_err
	}
	return records, nil
}

// @title    ReadXls
// @description   读取Xls文件内容
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     file_path string		文件位置
// @return    res [][]string, err error		res为读出的内容，err为可能出现的错误
func ReadXls(file_path string) (res [][]string, err error) {
	if xlFile, err := xls.Open(file_path, "utf-8"); err == nil {
		fmt.Println(xlFile.Author)
		// TODO 第一个sheet
		sheet := xlFile.GetSheet(0)
		if sheet.MaxRow != 0 {
			temp := make([][]string, sheet.MaxRow)
			for i := 0; i < int(sheet.MaxRow); i++ {
				row := sheet.Row(i)
				data := make([]string, 0)
				if row.LastCol() > 0 {
					for j := 0; j < row.LastCol(); j++ {
						col := row.Col(j)
						data = append(data, col)
					}
					temp[i] = data
				}
			}
			res = append(res, temp...)
		}
	} else {
		return nil, err
	}
	return res, nil
}

// @title    ReadXlsx
// @description   读取Xlsx文件内容
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     file_path string		文件位置
// @return    res [][]string, err error		res为读出的内容，err为可能出现的错误
func ReadXlsx(file_path string) (res [][]string, err error) {
	if xlFile, err := xlsx.OpenFile(file_path); err == nil {
		for index, sheet := range xlFile.Sheets {
			// TODO 第一个sheet
			if index == 0 {
				temp := make([][]string, len(sheet.Rows))
				for k, row := range sheet.Rows {
					var data []string
					for _, cell := range row.Cells {
						data = append(data, cell.Value)
					}
					temp[k] = data
				}
				res = append(res, temp...)
			}
		}
	} else {
		return nil, err
	}
	return res, nil
}

// @title    GetFiles
// @description   获取一个目录下的所有文件
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     folder string	指定目录
// @return    []string    所有文件的文件名
func GetFiles(folder string) []string {
	files, _ := ioutil.ReadDir(folder)
	res := make([]string, 10)
	for _, file := range files {
		if file.IsDir() {
			GetFiles(folder + file.Name())
			continue
		} else {
			res = append(res, file.Name())
		}
	}
	return res
}

// @title    PathExists
// @description   判断文件夹是否存在
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     path string	指定目录
// @return    bool, error    查看文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// @title    Mkdir
// @description   建立文件夹
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     path string	指定路径
// @return   error    查看是否出错
func Mkdir(dir string) error {
	exist, err := PathExists(dir)
	if err != nil {
		return err
	}

	if !exist {
		// TODO 创建文件夹
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// @title    StringToFloat
// @description   从字符串中提取各式各样的浮点数
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     s string		一串字符串
// @return    float64, bool		表示解析出来的浮点数，ok表示解析是否成功
func StringToFloat(s string) (float64, bool) {
	// TODO 优先查看数据注册表
	data, ok := DataMap[s]
	if ok {
		return data, ok
	}
	k := len(s)
	// TODO 尝试取出前缀数字，以此来滤过符号单位
	for k >= 0 {
		_, err := strconv.ParseFloat(s[0:k], 64)
		if err != nil {
			k--
		} else {
			break
		}
	}
	// TODO 成功取出数字
	if k > 0 {
		data, err := strconv.ParseFloat(s[0:k], 64)
		if err != nil {
			return 0, false
		}
		// TODO 查看是否有科学计数法
		if k+4 <= len(s) && s[k:(k+4)] == "×10" {
			// TODO 尝试读出后缀数字
			data1, ok := StringToFloat(s[(k + 4):])
			if !ok {
				data1 = 0
			} else if data1 == 0 {
				data1 = 1
			}
			data *= math.Pow(10, data1)
		}
		return data, true
	}
	return 0, false
}

// @title    RandomString
// @description   生成一段随机的字符串
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     n int		字符串的长度
// @return    string    一串随机的字符串
func RandomString(n int) string {
	var letters = []byte("qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	// TODO 不断用随机字母填充字符串
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// @title    VerifyEmailFormat
// @description   用于验证邮箱格式是否正确的工具函数
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     email string		一串字符串，表示邮箱
// @return    bool    返回是否合法
func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// @title    VerifyMobileFormat
// @description   用于验证手机号格式是否正确的工具函数
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     mobileNum string		一串字符串，表示手机号
// @return    bool    返回是否合法
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// @title    VerifyQQFormat
// @description   用于验证QQ号格式是否正确的工具函数
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     QQNum string		一串字符串，表示QQ
// @return    bool    返回是否合法
func VerifyQQFormat(QQNum string) bool {
	regular := "[1-9][0-9]{4,10}"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(QQNum)
}

// @title    VerifyQQFormat
// @description  用于验证Icon是否为默认图片的工具函数
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     Icon string		一串字符串，表示图像名称
// @return    bool    返回是否合法
func VerifyIconFormat(Icon string) bool {
	regular := "MGA[1-9].jpg"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(Icon)
}

// @title    isEmailExist
// @description   查看email是否在数据库中存在
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func IsEmailExist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email = ?", email).First(&user)
	return user.ID != 0
}

// @title    isNameExist
// @description   查看name是否在数据库中存在
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func IsNameExist(db *gorm.DB, name string) bool {
	var user model.User
	db.Where("name = ?", name).First(&user)
	return user.ID != 0
}

var ctx context.Context = context.Background()

// @title    SendEmailValidate
// @description   发送验证邮件
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    em []string       接收一个邮箱字符串
// @return   string, error     返回验证码和error值
func SendEmailValidate(em []string) (string, error) {
	mod := `
	尊敬的%s，您好！

	您于 %s 提交的邮箱验证，本次验证码为%s，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。
	此邮箱为系统邮箱，请勿回复。
`
	e := email.NewEmail()
	e.From = "mgAronya <2829214609@qq.com>"
	e.To = em
	// TODO 生成6位随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	t := time.Now().Format("2006-01-02 15:04:05")
	// TODO 设置文件发送的内容
	content := fmt.Sprintf(mod, em[0], t, vCode)
	e.Text = []byte(content)
	// TODO 设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "2829214609@qq.com", "rmdtxokuuqyrdgii", "smtp.qq.com"))
	return vCode, err
}

// @title    SendEmailPass
// @description   发送密码邮件
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    em []string       接收一个邮箱字符串
// @return   string, error     返回验证码和error值
func SendEmailPass(em []string) string {
	mod := `
	尊敬的%s，您好！

	您于 %s 提交的邮箱验证，已经将密码重置为%s，为了保证账号安全。切勿向他人泄露，并尽快更改密码，感谢您的理解与使用。
	此邮箱为系统邮箱，请勿回复。
`
	e := email.NewEmail()
	e.From = "mgAronya <2829214609@qq.com>"
	e.To = em
	// TODO 生成8位随机密码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	password := fmt.Sprintf("%08v", rnd.Int31n(100000000))
	t := time.Now().Format("2006-01-02 15:04:05")

	db := common.GetDB()

	// TODO 创建密码哈希
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "密码加密失败"
	}

	// TODO 更新密码
	err = db.Model(&model.User{}).Where("email = ?", em[0]).Updates(model.User{
		Password: string(hasedPassword),
	}).Error

	if err != nil {
		return "密码更新失败"
	}

	// TODO 设置文件发送的内容
	content := fmt.Sprintf(mod, em[0], t, password)
	e.Text = []byte(content)
	// TODO 设置服务器相关的配置
	err = e.Send("smtp.qq.com:25", smtp.PlainAuth("", "2829214609@qq.com", "rmdtxokuuqyrdgii", "smtp.qq.com"))

	if err != nil {
		return "邮件发送失败"
	}

	return "密码已重置"
}

// @title    IsEmailPass
// @description   验证邮箱是否通过
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    em []string       接收一个邮箱字符串
// @return   string, error     返回验证码和error值
func IsEmailPass(email string, vertify string) bool {
	client := common.GetRedisClient(0)
	V, err := client.Get(ctx, email).Result()
	if err != nil {
		return false
	}
	return V == vertify
}

// @title    SetRedisEmail
// @description   设置验证码，并令其存活五分钟
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    email string, v string       接收一个邮箱和一个验证码
// @return   void
func SetRedisEmail(email string, v string) {
	client := common.GetRedisClient(0)

	client.Set(ctx, email, v, 300*time.Second)
}
