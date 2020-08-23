package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
	"!@#$%^&*()-_=+,.?/:;{}[]`~"

func genPass(num int) string {
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
	rand.Seed(time.Now().UnixNano())
	var num, length int
	flag.IntVar(&num, "n", 1, "specify amount of passwords to generate")
	flag.IntVar(&length, "l", 15, "specify the length of generated passwords")
	flag.Parse()

	i := 0
	for i != num {
		fmt.Println(genPass(length))
		i++
	}
}
