package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/zengming00/go17monip"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	startTime := time.Now()
	for i := 0; i < 1000000; i++ {
		ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
		go17monip.IpFind(net.ParseIP(ip))
	}
	fmt.Println("end ,find ", time.Now().Sub(startTime))
}
