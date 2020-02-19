package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	for i := 0; i < n; i++ {

		var lines []int
		t := strconv.Itoa(nextInt())
		for j := 0; j < 4; j++ {
			v, e := strconv.Atoi(t[j : j+1])
			if e != nil {
				panic(e)
			}
			lines = append(lines, v)
		}

		current := lines[0]
		same := 0
		sp := false
		for k, val := range lines {
			if k == 0 {
				continue
			}

			if current == val {
				same++
			} else {
				if k == 2 {
					sp = true
				}
			}
			current = val
		}

		OutputCount(same, sp)
	}
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())

	if e != nil {
		panic(e)
	}

	return i
}

func OutputCount(s int, sp bool) {
	if s == 0 {
		fmt.Println("No Pair")
		return
	}

	if s == 1 {
		fmt.Println("One Pair")
		return
	}

	if s == 3 {
		fmt.Println("Four Card")
		return
	}

	if s == 2 {
		if sp == true {
			fmt.Println("Two Pair")
		} else {
			fmt.Println("Three Card")
		}
		return
	}
}
