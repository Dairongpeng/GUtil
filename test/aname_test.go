package test

import (
	"fmt"
	"testing"
)

type A struct {
	name string
}

func (a A) Name() string {
	return a.name
}

func TestHello(t *testing.T) {
	a := A{name: "eggo"}

	// 可以直接调用方法，是go的语法糖。
	fmt.Println(a.Name())
	// 等价于
	fmt.Println(A.Name(a))
}
