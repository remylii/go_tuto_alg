package main

import (
  "testing"
)

func TestOkaneSuccess(t *testing.T) {
  expected_maps := map[string]int{"10000": 1, "5000": 1, "1000": 4, "500": 1, "100": 4, "50": 1, "10": 4, "5": 1, "1": 4}
  res := okane(19999)

  for k, expected := range expected_maps {
    value, ok := res[k];
    if ok == false {
      t.Fatalf("failed. key %s doesn't exists.", k)
    }

    if value != expected {
      t.Fatalf("failed. return value (%d) is not equals %d", value, expected)
    }
  }
}

