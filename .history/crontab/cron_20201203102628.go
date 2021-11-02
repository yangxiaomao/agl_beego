/*
 * @Author: yxm
 * @Date: 2020-11-17 01:56:01
 * @LastEditTime: 2020-12-03 10:16:28
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

	// 批量创建MD5int原文 + 密文

	tk := toolbox.NewTask("batch_create_md5_int", "0 0 17 * * 5", batchCreateMd5Int)
	// err := tk.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	toolbox.AddTask("batch_create_md5_int", tk)

	// 批量创建MD5string原文 + 密文

	tk1 := toolbox.NewTask("batch_create_md5_string", "0 0 17 * * 5", batchCreateMd5String)

	toolbox.AddTask("batch_create_md5_string", tk1)

}

func batchCreateMd5Int() error {
	/* for 循环 */
	for a := 15181970; a < 100000000; a++ {
		var obj models.Md5Int
		o := orm.NewOrm()
		obj.OriginalString = strconv.Itoa(a)
		fmt.Println(strconv.Itoa(a))
		obj.DenseString = util.Md5(strconv.Itoa(a))
		// 事务处理过程
		if md5id, err := models.AdMd5Int(&obj, o); err != nil {
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

func batchCreateMd5String() error {
	/* for 循环 */
	for a := 1; a < 100000000; a++ {
		var obj models.Md5Int
		o := orm.NewOrm()
		obj.OriginalString = strconv.Itoa(a)
		fmt.Println(strconv.Itoa(a))
		obj.DenseString = util.Md5(strconv.Itoa(a))
		// 事务处理过程
		if md5id, err := models.AdMd5Int(&obj, o); err != nil {
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

    /*
    * 用户唯一邀请码生成
    * 2020-12-03
    * yxm
    */

func createCode(userid int64) error {
        sourceString := [
            0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
            'a', 'b', 'c', 'd', 'e', 'f',
            'g', 'h', 'i', 'j', 'k', 'l',
            'm', 'n', 'o', 'p', 'q', 'r',
            's', 't', 'u', 'v', 'w', 'x',
            'y', 'z'
        ]

        num := userid;
        $code = '';
        while ($num) {
            $mod = $num % 35;
            $num = (int) ($num / 35);
            $code = "{$sourceString[$mod]}{$code}";
        }
        //判断code的长度
        if (empty($code[6])) {
            $code = str_pad($code, 6, '0', STR_PAD_RIGHT);
        }
        return nil
    }
