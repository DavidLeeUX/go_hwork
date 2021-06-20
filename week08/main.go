/*
 * @Descripttion:
 * @version: xv1.0
 * @Author: changwei5
 * @Date: 2021-06-20 21:07:22
 * @LastEditors: changwei5
 * @LastEditTime: 2021-06-20 21:35:10
 */
package main

import (
	"encoding/json"
	"fmt"
	_ "net/http/pprof"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "10.210.100.214:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
}
func main() {
	start := time.Now()
	WriteRedis()
	fmt.Println("耗时: ", time.Since(start))
}

func WriteRedis() {
	max := 10000
	data := [1024]byte{}
	p := client.Pipeline()
	for i := 0; i < max; i++ {
		b, _ := json.Marshal(data)
		err := p.Set(fmt.Sprintf("key%v", i), b, time.Hour).Err()
		//fmt.Printf("status: %v\n", d)
		if err != nil {
			fmt.Printf("eroor:%v\n", err)
		}
	}
	x, e := p.Exec()
	if e != nil {
		fmt.Printf("error: %v \n", e)
	}
	fmt.Printf("result: %v \n", x)
}
