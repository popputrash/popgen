package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

	args := os.Args

	i, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(genPass(i))

}
