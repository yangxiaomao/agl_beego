/*
 * @Author: your name
 * @Date: 2020-10-29 19:05:34
 * @LastEditTime: 2020-12-16 11:55:29
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /src/beeapi/controllers/v1/user.go
 */
package controllers

// Operations about Game
type GameController struct {
	baseController
}

/**
 * @description: 	3D旋转拼图小游戏
 * @param {*GameController} yxm --- 2020-12-16
 * @return {*}	json
 */
func (g *GameController) RotatingJigsaw() {
	this.Data["IsDiagram"] = true
	this.Data["IsLongPolling"] = true

	this.TplName = "rotating_jigsaw.html"
}
