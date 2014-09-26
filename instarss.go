package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

const REDISHOST = "redis:6379"

var c, err = redis.Dial("tcp", REDISHOST)

func sendData(input string) string {
	if err != nil {
		log.Fatal(err)
	}

	c.Do("SET", "test-string", input)
	r, err := redis.String(c.Do("GET", "test-string"))

	if err != nil {
		log.Fatal(err)
	}

	return r
}

func main() {
	r := sendData("test")
	fmt.Println(r)
	defer c.Close()
}
