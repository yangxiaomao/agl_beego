/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-12-16 13:14:19
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/v1/user.go
 */
package controllers

import "github.com/astaxie/beego"

// Operations about Game
type SortController struct {
	baseController
}

/**
 * @description: 	3D旋转拼图小游戏
 * @param {*GameController} yxm --- 2020-12-16
 * @return {*}	json
 */
func (g *SortController) RotatingJigsaw() {
	beego.Info(123)
	g.TplName = "rotatingjigsaw.html"
}
