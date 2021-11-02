/*
 * @Author: your name
 * @Date: 2020-10-31 02:12:08
 * @LastEditTime: 2020-11-28 00:17:57
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/main.go
 */
package main

import (
	_ "beeapi/routers"
	"beeapi/util"

	"github.com/Echosong/beego_blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

//初始化函数
func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true

	//初始化日志
	util.InitLogs()

}

func main() {
	orm.Debug = true

	// crontab.Init()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
