/*
 * @Author: yxm
 * @Date: 2020-11-17 01:56:01
 * @LastEditTime: 2020-12-01 20:25:03
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /beeapi/crontab/cron.go
 */
package crontab

import (
	models "beeapi/models/mymysql"
	"beeapi/util"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	/* for 循环 */
	for a := 8223990; a < 100000000; a++ {
		var obj models.Md5
		o := orm.NewOrm()
		obj.OriginalString = strconv.Itoa(a)
		fmt.Println(strconv.Itoa(a))
		obj.DenseString = util.Md5(strconv.Itoa(a))
		// 事务处理过程
		if md5id, err := models.AdMd5(&obj, o); err != nil {
			beego.Error("失败！" + err.Error())
			fmt.Println("失败")
		} else {
			fmt.Println(md5id)
			fmt.Println("成功")
		}
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	return nil
}
