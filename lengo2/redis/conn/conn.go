package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.56.110:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()
}
