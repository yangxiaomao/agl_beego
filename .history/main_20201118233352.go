/*
 * @Author: your name
 * @Date: 2020-10-31 02:12:08
 * @LastEditTime: 2020-11-18 05:44:44
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/main.go
 */
package main

import (
	_ "beeapi/routers"
	"beeapi/util"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

//初始化函数
func init() {

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
