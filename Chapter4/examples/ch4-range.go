package main

import (
	"log"
	"strconv"
)

const LENGTH = 10

var list []string

func main() {
	for i := 0; i < LENGTH; i++ {
		v := "Entry #" + strconv.FormatInt(int64(i), 10)
		list = append(list, v)
	}

	for j := range list {
		log.Println(list[j])
	}
}
