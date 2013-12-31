package main

import (
    "testing"
)

func simpleTest(t *testing.T) {
    if 1 == 2 {
        t.Error("Integer Value Error")
    }
}
