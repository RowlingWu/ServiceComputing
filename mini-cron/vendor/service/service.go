package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"manager"
	"net/http"
	"net/url"
)

// glManager 是全局的任务管理器，负责定时任务的创建、删除
var glManager = manager.MyManager

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	sbody, _ := ioutil.ReadAll(r.Body) // 读取http的json参数
	body, _ := url.QueryUnescape(string(sbody))
	defer r.Body.Close()

	var task manager.Task
	err := json.Unmarshal([]byte(body), &task) // 从json中解析Task的内容

	str := ""                           // str 要写入 ResponseWriter 的json内容
	statusCode := http.StatusBadRequest // statusCode 返回状态码

	if err != nil { // json解析参数出错
		fmt.Println("json error: ", err)
		str = err.Error()
	} else { // json解析无误，判断请求是POST还是DELETE
		if r.Method == "POST" {
			ok, err := glManager.Create(&task) // 创建定时任务

			switch {
			case ok: // 任务创建成功，返回200
				str = "{\"ok\":true, \"id\":\"" + task.ID + "\"}"
				statusCode = 200
			case !ok: // 已存在该任务，任务创建失败，返回409
				str = "{\"ok\":false, \"error\":\"The task " + task.ID + " already exists.\"}"
				statusCode = 409
			case err != nil: // 创建任务时发生其他错误
				fmt.Println("manager.create() error:", err)
				str = err.Error()
			}
		} else { // DELETE方法。但其实curl -X DELETE 是无法传参数的，故此url我用了PUT参数代替DELETE
			if err := glManager.Destroy(task.ID); err != nil { // 没有该任务，删除该任务时出错，返回404
				fmt.Println(err)
				str = "{\"ok\":false, \"error\":\"The task " + task.ID + " is not found.\"}"
				statusCode = 404
			} else { // 删除该任务无误，返回200
				str = "{\"ok\":true, \"id\":\"" + task.ID + "\"}"
				statusCode = 200
			}
		}
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(str))
}
