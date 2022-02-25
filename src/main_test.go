package main


import (
    "testing"
)

func TestHello(t *testing.T) {

 want := "MongoDB"

 got := "MongoDB"

 if want != got {
  t.Fatalf("want %s, got %s\n", want, got)
 }
}