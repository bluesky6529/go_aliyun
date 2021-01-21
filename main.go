package main

import (
	"log"

	"github.com/bluesky6529/go_aliyun/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
