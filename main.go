package main

import (
	"github.com/monkingxue/gopack/src"
	"fmt"
)

func main(){
	m, l := src.CreateModule("a", `
	var a = require('b');
	var d = 1;
	var c = require('d');
	console.log(a)`)
	fmt.Println(m.Code)
	fmt.Println(l)
}
