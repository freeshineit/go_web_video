package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func Test(t *testing.T) {
	t.Run("xxx", func(t *testing.T) { fmt.Printf("xxx") })
}
