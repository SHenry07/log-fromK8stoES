package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"logpush/utils"
	//"github.com/gin-gonic/gin"
)

func main() {

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://192.168.100.69:2379"},
		//DialTimeout: 2 * time.Second,
	})
	if err != nil {
		log.Panicf("Cant connet etcd server err: %v", err)
	}
	ip := utils.Init()
	key := fmt.Sprintf("/logagent/%s/collect_log1", ip)
	//val, err := json.Marshal("ddd")
	val := `[{"path":"./nginx.log","topic":"app_log"},{"path":"./catalina.out","topic":"nginx_log"},{"path":"./mysql.log","topic":"sql_log"}]`
	_, err = client.Put(context.Background(), key, val)
	if err != nil {
		log.Printf("etcd: put key to etcd failed: %v", err)
	}

}
