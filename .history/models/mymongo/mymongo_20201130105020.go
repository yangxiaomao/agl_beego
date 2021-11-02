/*
 * @Author: your name
 * @Date: 2020-11-27 19:34:05
 * @LastEditTime: 2020-11-30 10:50:20
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /beeapi/models/mymongo/mymongo.go
 */
package mymongo

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

// Conn return mongodb session.
func Conn() *mgo.Session {
	return session.Copy()
}

/*
func Close() {
	session.Close()
}
*/

func init() {
	// url := beego.AppConfig.String("mongodb::url")

	// sess, err := mgo.Dial(url)
	// if err != nil {
	// 	panic(err)
	// }

	// session = sess
	// session.SetMode(mgo.Monotonic, true)
}
