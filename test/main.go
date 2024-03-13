package main

import (
	"fmt"
	"github.com/clearcodecn/v2ray-core/v2raystart"
	"time"
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
