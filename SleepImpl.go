package main

import (
	"fmt"
	"time"
)

func userSleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	fmt.Println("entered")
	userSleep(9)
	fmt.Println("exit")
}
