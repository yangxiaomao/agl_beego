/*
 * @Author: yxm
 * @Date: 2020-11-17 01:56:01
 * @LastEditTime: 2020-11-18 03:05:42
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /beeapi/crontab/cron.go
 */
package crontab

import (
	"fmt"
	"time"
)

//	 second minute hour day month week   command
//顺序：秒      分    时   日   月    周      命令
func Init() {

	//测试使用

	// tk := toolbox.NewTask("task1", "0/5 * * * * *", task)
	// err := tk.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// toolbox.AddTask("task1", tk)

}

func task() error {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	return nil
}
