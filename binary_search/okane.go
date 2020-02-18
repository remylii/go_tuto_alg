package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
  n := getInputValue()

  res := okane(n)

  keys := []string{"10000", "5000", "1000", "500", "100", "50", "10", "5", "1"}
  for _, v := range keys {
    fmt.Printf("%s x %d\n", v, res[v])
  }
}

func getInputValue() int {
  sc.Scan()

  n, e := strconv.Atoi(sc.Text())
  if e != nil {
    panic(e)
  }

  return n
}

func okane(price int) map[string]int {
  yens := []int{10000, 5000, 1000, 500, 100, 50, 10, 5, 1}
  result := make(map[string]int, len(yens))

  for _, yen := range yens {
    cnt := price / yen
    result[strconv.Itoa(yen)] = cnt
    price = price  - (yen * cnt)
  }

  return result
}
