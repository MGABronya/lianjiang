// @Title  HourController
// @Description  该文件用于提供操作小时点集的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	"lianjiang/common"
	"lianjiang/model"
	"lianjiang/util"
	"log"
	"path"
	"reflect"
	"strconv"
	"time"

	"lianjiang/response"

	"github.com/gin-gonic/gin"
)

// @title    HourUpload
// @description   用户上传小时点集文件
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func HourUpload(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")

	user := tuser.(model.User)

	if user.Level < 2 {
		response.Fail(ctx, nil, "权限不足")
		return
	}

	file, err := ctx.FormFile("file")

	//TODO 数据验证
	if err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// TODO 验证文件格式
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".xls":  true,
		".xlsx": true,
		".csv":  true,
	}

	// TODO 格式验证
	if _, ok := allowExtMap[extName]; !ok {
		response.Fail(ctx, nil, "文件后缀有误")
		return
	}

	// TODO 将文件存入本地
	ctx.SaveUploadedFile(file, "./hour/"+file.Filename)

	// TODO 解析文件
	res, err := util.Read("./hour/" + file.Filename)

	// TODO 解析有误
	if err != nil || res == nil {
		response.Fail(ctx, nil, "文件解析有误")
		return
	}

	// TODO 将文件中的数据存入数据库
	db := common.GetDB()

	// TODO 第一列字段为时间time
	index := make([]string, 1)

	index[0] = "time"

	// TODO n和m用于标记表的范围
	n := 0
	m := 0

	// TODO 取出站点名
	name := res[3][0][18:]

	// TODO 取出映射
	for i := 1; ; i++ {
		// TODO 遍历到了边缘则记录并退出
		if i == len(res[5]) || res[5][i] == "" {
			n = i
			break
		}
		// 将映射存入index
		tmp, ok := util.HourMap[res[5][i]]
		if ok {
			index = append(index, tmp)
		}
	}

	// TODO 找到数据的开头
	for i := 0; ; i++ {
		if res[i][0] == "时间" {
			m = i + 1
			break
		}
		// TODO 没有数据头，退出
		if i > 10 {
			response.Fail(ctx, nil, "文件格式有误")
			return
		}
	}

	// TODO 一行一行的遍历数据，将数据存入数据库
	for i := m; i < len(res); i++ {
		var point model.PointHour
		// TODO 空数据则退出
		if res[i][0] == "" {
			break
		}

		// TODO 损坏、错误数据直接滤过
		t, err := strconv.ParseFloat(res[i][0], 64)
		if err != nil {
			continue
		}

		// TODO 如果出现了数据读出损坏，尝试修复数据
		if t < 40000 {
			// TODO 如果是递增或者递减，则测算出损坏数据
			if i > m+3 {
				var t1, t2, t3 float64

				// TODO 取出前三位数据
				t1, err = strconv.ParseFloat(res[i-1][0], 64)
				if err != nil {
					continue
				}

				t2, err = strconv.ParseFloat(res[i-2][0], 64)
				if err != nil {
					continue
				}

				t3, err = strconv.ParseFloat(res[i-3][0], 64)
				if err != nil {
					continue
				}

				// TODO 不满足递增或者递减，滤过这条数据
				if (t3-t2)-(t2-t3) > 0.00001 {
					continue
				}

				// TODO 满足则计算处预测值
				t = t1 + t1 - t2
			} else {
				var t1, t2, t3 float64
				// TODO 取出后三位数据
				t1, err = strconv.ParseFloat(res[i+1][0], 64)
				if err != nil {
					continue
				}

				t2, err = strconv.ParseFloat(res[i+2][0], 64)
				if err != nil {
					continue
				}

				t3, err = strconv.ParseFloat(res[i+3][0], 64)
				if err != nil {
					continue
				}

				// TODO 不满足递增或者递减，滤过这条数据
				if (t3-t2)-(t2-t3) > 0.00001 {
					continue
				}

				// TODO 满足则计算处预测值
				t = t1 - (t2 - t1)
			}
		}

		// TODO 计算正确时间
		point.Time = time.Unix(int64((t-25569)*24*60*60)-8*60*60, 0)

		// 遍历每一列，尝试取出数据
		for j := 1; j < n; j++ {
			k := len(res[i][j])
			tmp := 0.0
			var err error
			// TODO 尝试取出前缀数字，以此来滤过符号单位
			for k >= 0 {
				_, err = strconv.ParseFloat(res[i][j][0:k], 64)
				if err != nil {
					k--
				} else {
					break
				}
			}
			// TODO 成功取出数字，利用反射机制写入结构体
			if k > 0 {
				tmp, err = strconv.ParseFloat(res[i][j][0:k], 64)
				if err == nil {
					reflect.ValueOf(&point).Elem().FieldByName(index[j]).SetFloat(tmp)
				}
			}
		}
		// TODO 标记上站点后存入数据库
		point.StationName = name
		db.Create(&point)
	}

	response.Success(ctx, nil, "更新成功")
}
