/*
 * @Author: your name
 * @Date: 2020-10-31 04:24:59
 * @LastEditTime: 2020-12-12 16:29:49
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/routers/router.go
 */
package routers

import (
	adminv1 "beeapi/controllers/admin/v1"
	clientv1 "beeapi/controllers/client/v1"

	"github.com/astaxie/beego"
)

func init() {

	// Register routers.
	beego.Router("/", &clientv1.LoginController{}, "*:GetHome")
	// Indicate AppController.Join method to handle POST requests.
	beego.Router("/join", &clientv1.AppController{}, "post:Join")

	// Long polling.
	beego.Router("/lp", &clientv1.LongPollingController{}, "get:Join")
	beego.Router("/lp/post", &clientv1.LongPollingController{})
	beego.Router("/lp/fetch", &clientv1.LongPollingController{}, "get:Fetch")

	// WebSocket.
	beego.Router("/ws", &clientv1.WebSocketController{})
	beego.Router("/ws/join", &clientv1.WebSocketController{}, "get:Join")

	// 曲线图
	beego.Router("/diagram", &clientv1.WebSocketController{}, "get:Diagram")
	// 曲线图WebSocket获取数据
	beego.Router("/ws/join", &clientv1.WebSocketController{}, "get:Join")

	//Api接口RSA加密
	beego.Router("/api1/client/rsa_encryption", &clientv1.RsaController{}, "*:RsaEncrypting")
	//Api接口RSA解密
	beego.Router("/api1/client/rsa_decryption", &clientv1.RsaController{}, "*:RsaDecrypting")

	//Mongodb测试路由
	beego.Router("/api1/client/mongo_test", &clientv1.RsaController{}, "*:MongoTest")

	//后台接口RSA加密
	beego.Router("/api1/admin/rsa_encryption", &adminv1.AdminRsaController{}, "*:RsaEncrypting")
	//后台接口RSA解密
	beego.Router("/api1/admin/rsa_decryption", &adminv1.AdminRsaController{}, "*:RsaDecrypting")

	//api1版本路由组
	ns := beego.NewNamespace("/api1",
		//Client端路由组
		beego.NSNamespace("/client",
			//用户路由组
			beego.NSNamespace("/user",
				//用户个人中心接口
				beego.NSRouter("/user_personal_center", &clientv1.UserController{}, "*:UserPersonalCenter"),
				//用户注册
				beego.NSRouter("/user_register", &clientv1.LoginController{}, "*:UserRegister"),
				//用户登录
				beego.NSRouter("/user_login", &clientv1.LoginController{}, "*:UserLogin"),
				//用户签到
				beego.NSRouter("/user_sign_in", &clientv1.UserController{}, "*:UserSignIn"),
			),
		),
		//Admin端路由组
		beego.NSNamespace("/admin",
			//管理员用户路由组
			beego.NSNamespace("/user",
				//用户注册
				beego.NSRouter("/register", &adminv1.AdminLoginController{}, "*:AdminUserRegister"),
				//用户登录
				beego.NSRouter("/login", &adminv1.AdminLoginController{}, "*:AdminUserLogin"),
			),
			//管理员商品路由组
			beego.NSNamespace("/goods",
				//商品分类列表
				beego.NSRouter("/goods_category_list", &adminv1.AdminGoodsGategoryController{}, "*:GoodsCategoryList"),
				//商品分类添加
				beego.NSRouter("/goods_category_add", &adminv1.AdminGoodsGategoryController{}, "*:AddGoodsCategory"),
				//商品分类更新
				beego.NSRouter("/goods_category_update", &adminv1.AdminGoodsGategoryController{}, "*:UpdateGoodsCategory"),
				//商品列表
				beego.NSRouter("/goods_list", &adminv1.AdminGoodsController{}, "*:GoodsList"),
				//商品添加
				beego.NSRouter("/goods_add", &adminv1.AdminGoodsController{}, "*:AddGoods"),
				//商品更新
				beego.NSRouter("/goods_update", &adminv1.AdminGoodsController{}, "*:UpdateGoods"),
			),

			//MD5路由组
			beego.NSNamespace("/md5",
				//根据密文查原文
				beego.NSRouter("/search_original_int", &adminv1.AdminMd5Controller{}, "*:AdminSearchOriginalInt"),
			),
		),
	)
	beego.AddNamespace(ns)
}
