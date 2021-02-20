package sorts

// 1、选择排序。golang中数组是值传递的，这里地址传递数组，效率高，但是传地址的话，形参修改，实参也会更改。
func SelectionSort(arr *[]int) {

	if arr == nil || len(*arr) < 2 {
		return
	}

	for i := 0; i < len(*arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(*arr); j++ {
			// golang中没有三目运算符，设计师觉得三目运算难以理解，if-else虽然繁琐点，但更清晰。
			// 想要取数组指针的值，需要对数组首地址解引用，再按下标读取
			if (*arr)[j] < (*arr)[minIndex] {
				minIndex = j
			}
			Swap(arr, i, minIndex)
		}
	}

}

// 2、冒泡排序
func BubbleSort(arr *[]int) {

	if arr == nil || len(*arr) < 2 {
		return
	}

	for e := len(*arr) - 1; e > 0; e-- {
		// 每次内循环搞定后面e相应位置的一个数
		for i := 0; i < e; i++ {
			// i < e 最后一个位置的下一个位置就是e，不越界
			if (*arr)[i] > (*arr)[i+1] {
				Swap(arr, i, i+1)
			}
		}
	}

}

// 3、插入排序
func InsertionSort(arr *[]int) {

	if arr == nil || len(*arr) < 2 {
		return
	}

	for i := 1; i < len(*arr); i++ {
		for j := i - 1; j >= 0 && (*arr)[j] > (*arr)[j+1]; j-- {
			Swap(arr, j, j+1)
		}
	}
}

func Swap(arr *[]int, i int, j int) {
	tem := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = tem
}

//func main() {
//	// 自动推导数组长度。但是由于形参是没指定数组大小的切片，所以这里要传入不指定大小的切片
//	// arr := [...]int{1, 2, 4, 7, 9, 3}
//	// 切片
//	arr := []int{1, 2, 4, 7, 9, 3}
//	SelectionSort(&arr)
//	for i := 0; i < len(arr); i++ {
//
//		fmt.Printf("%d\n", arr[i])
//
//	}
//}
