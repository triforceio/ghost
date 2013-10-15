package main

import (
  "testing"
)

func TestSHA256Fmt(t *testing.T) {
  testStr := "ruby+node+go"
  expected := "f4aeb2"

  if name := FmtName(testStr); name != expected{
    t.Errorf("Expected %s got %s", name, expected)
  }
}
