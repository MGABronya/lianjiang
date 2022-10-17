// @Title  main
// @Description  程序的入口，读取配置，调用初始化函数以及运行路由
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:48
package main

import (
	"fmt"
	"lianjiang/common"
	"lianjiang/model"
	"lianjiang/util"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// @title    main
// @description   程序入口，完成一些初始化工作后将开始监听
// @auth      MGAronya（张健）             2022-9-16 10:49
// @param     void			没有入参
// @return    void			没有回参
func main() {
	InitConfig()
	common.InitDB()
	client0 := common.InitRedis(0)
	defer client0.Close()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	// TODO 定时备份映射表
	go func() {
		for {
			log.Println("Backup mapping in progress...")
			// TODO 执行备份功能
			BackUp()
			log.Println("The backup mapping is complete...")
			now := time.Now()
			// TODO 计算下一个4:00
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 4, 00, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

// @title    InitConfig
// @description   读取配置文件并完成初始化
// @auth      MGAronya（张健）             2022-9-16 10:49
// @param     void			没有入参
// @return    void			没有回参
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	// TODO 如果发生错误，终止程序
	if err != nil {
		panic(err)
	}
}

// @title    BackUp
// @description   备份映射文件
// @auth      MGAronya（张健）             2022-9-16 10:49
// @param     void			没有入参
// @return    void			没有回参
func BackUp() {
	db := common.GetDB()
	for id, v := range util.MapMap {
		for _, key := range v.Keys() {
			value, _ := v.Get(key)
			db.Create(&model.MapBackup{
				Id:    id,
				Key:   key,
				Value: fmt.Sprint(value),
			})
		}
	}
}
