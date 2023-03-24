package main

import (
	"fmt"
	"log"

	iptables "github.com/coreos/go-iptables/iptables"
)

// TODO: implement This
func Flushfirewall() {
	ipt, err := iptables.New()
	if err != nil {
		log.Panic(err)
	}

	result, err := ipt.ListChains("filter")
	if err != nil {
		log.Panic(err)
	}

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
