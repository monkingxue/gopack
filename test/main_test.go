package test

import (
	"testing"
	"github.com/gopack/src"
)

func TestName(t *testing.T) {
	var b src.Bundle
	println(b.Build("/Users/qyh/go/src/github.com/gopack/test/main.js","/Users/qyh/go/src/github.com/gopack/dest/"))
}