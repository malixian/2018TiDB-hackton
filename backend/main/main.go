package main

import (
	"log"
	"net/http"
	"TiDB_hackthon/backend/server"
)

func main() {
	log.Print("server start")
	server.NewRedisClinet()
	http.HandleFunc("/savePlan", server.SavePlan) //设置访问的路由
	http.HandleFunc("/getPlan", server.GetPlan)
	err := http.ListenAndServe(":8001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
