package main

import (
	"fmt"
	"time"
)

// +kubebuilder:rbac:groups=streaming.sneaksanddata.com,resources=teststreamdefinitions,verbs=get;list;watch;create;update;patch;delete
func main() {
	for {
		fmt.Println("running stream")
		time.Sleep(30 * time.Second)
	}
}
