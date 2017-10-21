package main

import "github.com/monkingxue/gopack/src"

func main(){
	src.CreateModule("a", "var a = require('b'); var c = require('b')")
}
