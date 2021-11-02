/*
 * @Author: your name
 * @Date: 2020-12-10 17:19:01
 * @LastEditTime: 2020-12-15 10:22:00
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /beeapi/controllers/client/v1/websocket.go
 */
// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/samples/WebIM/models"
	"github.com/gorilla/websocket"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	baseController
}

// Get method handles GET requests for WebSocketController.
func (this *WebSocketController) Get() {
	// Safe check.
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	this.TplName = "websocket.html"
	this.Data["IsWebSocket"] = true
	this.Data["UserName"] = uname
}

// Join method handles WebSocket requests for WebSocketController.
func (this *WebSocketController) Join() {
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "不是websocket握手", 400)
		return
	} else if err != nil {
		beego.Error("无法设置WebSocket连接:", err)
		return
	}

	// 加入聊天室。
	Join(uname, ws)
	defer Leave(uname)

	// 消息接收循环。
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		beego.Info(string(p))
		publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}

	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		// Immediately send event to WebSocket users.
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				// User disconnected.
				unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}

// Get method handles GET requests for WebSocketController.
func (this *WebSocketController) Diagram() {

	this.Data["IsDiagram"] = true
	this.Data["IsLongPolling"] = true

	this.TplName = "diagram.html"
}

// Join method handles WebSocket requests for WebSocketController.
func (this *WebSocketController) DiagramData() {
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "不是websocket握手", 400)
		return
	} else if err != nil {
		beego.Error("无法设置WebSocket连接:", err)
		return
	}

	// 加入聊天室。
	Join(uname, ws)
	defer Leave(uname)

	// 消息接收循环。
	for {
		ticker := time.NewTicker(time.Second) // 每隔1s进行一次打印
		var num int64
		num = 1

		returnData := make(map[string]interface{})
		categoriesData := []string{}
		seriesData := []int{}
		for {
			<-ticker.C

			categoriesData = append(categoriesData, "09:30")
			if(len(categoriesData) >= 10){
				categoriesData = categoriesData[0, 10]
			}
			beego.Info(len(categoriesData))

			seriesData = append(seriesData, rand.Intn(20))
			returnData["categories"] = categoriesData
			returnData["series"] = seriesData
			returnData["num"] = num
			mjson, _ := json.Marshal(returnData)

			beego.Info(string(mjson))

			publish <- newEvent(models.EVENT_MESSAGE, uname, string(mjson))
			num++
		}

	}
}
