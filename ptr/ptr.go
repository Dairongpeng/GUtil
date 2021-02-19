package main

import "fmt"

// 详细解释Go语言指针，参考文章：http://c.biancheng.net/view/21.html, http://c.biancheng.net/golang/

// 要明白指针需要明白指针地址、指针类型和指针取值的区别。
// 每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
// Go语言中使用在变量名前面添加&操作符（前缀）来获取变量的内存地址（取地址操作），格式如下：

// ptr := &v    // v 的类型为 T

// 其中 v 代表被取地址的变量，变量 v 的地址使用变量 ptr 进行接收，ptr 的类型为*T，称做 T 的指针类型，*代表指针。
// 提示：变量、指针和地址三者的关系是，每个变量都拥有地址，指针的值就是地址。

func Ptr() {
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p", &cat, &str) // 运行结果 0xc042052088 0xc0420461b0
}

func Ptr2() {
	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr) // ptr type: *string
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr) // address: 0xc0420401b0
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value) // value type: string
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value) // value: Malibu Point 10880, 90265
}

// 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

// 变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
// 1、对变量进行取地址操作使用&操作符，可以获得这个变量的指针变量。
// 2、指针变量的值是指针地址。
// 3、对指针变量进行取值操作使用*操作符，可以获得指针变量指向的原变量的值。

// 交换函数。传入两个指针地址，交换这两个指针解引用后各自对应的值
func swap(a, b *int) {
	// 取a指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}

// 创建指针的另外一种方式new()函数
// str := new(string)
// *str = "Go语言教程"
// fmt.Println(*str)

func main() {
	Ptr2()
}
