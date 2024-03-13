package main

import (
	"fmt"
	"time"
	"v2ray.com/core/v2raystart"
)

func main() {
	ch := make(chan struct{})
	srv, err := v2raystart.Start("default_client_config.json", ch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(srv)

	time.Sleep(10 * time.Hour)
}
