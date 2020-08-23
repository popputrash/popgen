package main

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func genPass(num int) string {
	rand.Seed(time.Now().UnixNano())
	a := []rune(charset)
	var n = num
	var ret = ""
	for n > 0 {
		var randI = rand.Intn(len(charset))
		ret += string(a[randI])
		n--
	}
	return ret
}

func main() {

}
