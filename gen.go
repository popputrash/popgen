package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
	"!@#$%^&*()-_=+,.?/:;{}[]`~"

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

	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "quantity", "length")
		os.Exit(2)
	}

	num, err := strconv.Atoi(os.Args[1])
	length, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	i := 0
	for i != num {
		fmt.Println(genPass(length))
		i++
	}
}
