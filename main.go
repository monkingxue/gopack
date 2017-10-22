package main

import (
	cp "github.com/monkingxue/gopack/src/composite"
	"fmt"
)

func main(){
	m, l := cp.CreateModule("a", `
	var a = require('b');
	var d = 1;
	var c = require('d');
	console.log(a)`)
	fmt.Println(m.Code)
	fmt.Println(l)
}
