package go17monip

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"testing"
	"time"
)

func Example() {
	ip := net.ParseIP("217.171.224.66")
	fmt.Println(IpFind(ip))
}
func Benchmark(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
		IpFind(net.ParseIP(ip))
	}
}

func Test1(t *testing.T) {
	var m = map[string]uint32{
		"217.171.224.66":  3651919938,
		"80.101.64.8":     1348812808,
		"188.204.155.170": 3167525802,
		"74.86.158.107":   1247190635,
		"83.83.140.49":    1397984305,
		"195.73.215.250":  3276396538,
		"37.74.94.85":     625630805,
		"62.195.45.68":    1052978500,
		"82.169.113.78":   1386836302,
		"145.53.147.68":   2436207428,
	}
	for k, v := range m {
		ip := net.ParseIP(k)
		if ip == nil {
			t.Error("parseip fail:", k, v)
			continue
		}
		country := IpFind(ip)
		t.Log("country:", country)

		if r := Ip2long(ip); r != v {
			t.Error("ip2long fail:", r, v)
		}
		if r := Long2ip(v).String(); r != k {
			t.Error("long2ip fail:", r, v)
		}
	}
}

func Test2(t *testing.T) {
	t.SkipNow()
	for i := math.MaxUint8; i >= 0; i-- {
		v, ok := mCCMAP[byte(i)]
		if !ok {
			t.Errorf("%d not exists: %#v", i, v)
		}
	}
}
