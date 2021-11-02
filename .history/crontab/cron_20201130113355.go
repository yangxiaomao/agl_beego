/*
 * @Author: yxm
 * @Date: 2020-11-17 01:56:01
 * @LastEditTime: 2020-11-30 11:33:54
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /beeapi/crontab/cron.go
 */
package crontab

import (
	models "beeapi/models/mymysql"
	"beeapi/util"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

//	 second minute hour day month week   command
//顺序：秒      分    时   日   月    周      命令
func Init() {

	//测试使用

	tk := toolbox.NewTask("task1", "0 0 17 * * 5", task)
	// err := tk.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	toolbox.AddTask("task1", tk)

}

func task() error {
	var obj models.Md5
	obj.Username = u.GetString("username")
	obj.Password = util.Md5(u.GetString("password"))
	// 事务处理过程
	if userid, err := models.AddUser(obj, o); err != nil {
		beego.Error("失败！" + err.Error())
		returnData = ReturnJSONData(10001, "失败！"+err.Error(), userRes)
	} else {

	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	return nil
}