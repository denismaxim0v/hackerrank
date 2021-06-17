package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Input:
// 9
// 10 5 20 20 4 5 2 25 1

// Output
// 2 4

func breakingTheRecords(scores []int32) []int32 {
	var hi, lo, countH, countL int32
	hi = scores[0]
	lo = scores[0]
	for i := 0; i < len(scores); i++ {
		if scores[i] > hi {
			hi = scores[i]
			countH++
		} else if scores[i] < lo {
			lo = scores[i]
			countL++
		}
	}

	res := []int32{countH, countL}

	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	strs := strings.Split(line[0:len(line)-1], " ")
	nums := make([]int32, len(strs))
	var num int
	for i, str := range strs {
		if num, err = strconv.Atoi(str); err != nil {
			log.Fatal(err)
		}
		nums[i] = int32(num)
	}

	fmt.Println(breakingTheRecords(nums))
}
