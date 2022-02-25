package main


import (
    "testing"
)

func TestHello(t *testing.T) {

 want := "Connected to MongoDB"

 got := "Connected to MongoDB"

 if want != got {
  t.Fatalf("want %s, got %s\n", want, got)
 }
}