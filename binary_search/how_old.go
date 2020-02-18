package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var default_min, default_max = 20, 35

func main() {
  n := getValue()
  if default_min > n || n > default_max {
    fmt.Printf("入力は %d 〜 %d でお願いします\n", default_min, default_max)
    return
  }

  res := judgement(n)
  fmt.Println(res)
}

func getValue() int {
  sc.Scan()
  n, e := strconv.Atoi(sc.Text())

  if e != nil {
    panic(e)
  }

  return n
}

func judgement(n int) string {
  var p int

  min, max := default_min, default_max
  diff := (default_max - default_min) / 2

  for diff > 1 {
    diff = (max - min) / 2
    p = min + diff
    if min <= n && n <= p {
      fmt.Printf("%d <= n <= %d \n", min, p)
      max = p
    } else {
      fmt.Printf("%d < n <= %d \n", p, max)
      min = p + 1
    }
  }

  if min < n {
    return strconv.Itoa(max) + " = " + strconv.Itoa(n)
  }

  return strconv.Itoa(min) + " = " + strconv.Itoa(n)
}

