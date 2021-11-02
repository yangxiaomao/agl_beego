/*
 * @Author: your name
 * @Date: 2020-10-31 04:24:59
 * @LastEditTime: 2020-11-18 05:39:47
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
	//Api接口RSA加密
	beego.Router("/api1/client/rsa_encryption", &clientv1.RsaController{}, "*:RsaEncrypting")
	//Api接口RSA解密
	beego.Router("/api1/client/rsa_decryption", &clientv1.RsaController{}, "*:RsaDecrypting")

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
		),
	)
	beego.AddNamespace(ns)
}