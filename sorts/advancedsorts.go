package sorts

import "math/rand"

// 1、归并排序
func MergeSort(arr *[]int) {

	if arr == nil || len(*arr) < 2 {
		return
	}

	Process(arr, 0, len(*arr)-1)

}

// 递归
func Process(arr *[]int, L int, R int) {
	// 如果L==R的时候，递归出口，当前分支有序
	if L == R {
		return
	}

	mid := L + ((R - L) / 2)

	Process(arr, L, mid)

	Process(arr, mid+1, R)

	// 两个分支的递归走完，当前两个区间的数组即有序，调用当前的合并函数
	Merge(arr, L, mid, R)
}

// 合并两个有序数组
func Merge(arr *[]int, L int, M int, R int) {
	// 辅助数组。R - L + 1是动态的，所以用声明的方法不行，必须用Make来初始化切片
	// var help [R - L + 1]int
	help := make([]int, R-L+1)

	i := 0
	p1 := L
	p2 := M + 1

	// p1和p2都没越界
	for p1 <= M && p2 <= R {
		if (*arr)[p1] <= (*arr)[p2] {
			help[i] = (*arr)[p1]
			i++
			p1++
		} else {
			help[i] = (*arr)[p2]
			i++
			p2++
		}
	}

	// 经过上述循环，此时要不然左数组越界，要不然右数组越界
	for p1 <= M {
		help[i] = (*arr)[p1]
		i++
		p1++
	}

	for p2 <= R {
		help[i] = (*arr)[p2]
		i++
		p2++
	}

	// 把有序的数组，拷贝回原数组
	for j := 0; j < len(help); j++ {
		(*arr)[L+j] = help[j]
	}

}

// 2、随机快排序
func QuickSort(arr *[]int) {

	if arr == nil || len(*arr) < 2 {
		return
	}

	ProcessByQuick(arr, 0, len(*arr)-1)

}

func ProcessByQuick(arr *[]int, L int, R int) {

	if L >= R {
		return
	}

	// 随机选择位置，与arr[R]上的数做交换
	Swap(arr, L+rand.Intn(R-L+1), R)
	// 整体进行荷兰国旗划分，返回中间相等区域的左右边界，每次荷兰国旗问题，都可以搞定一批位置
	equalArea := NetherlandsFlag(arr, L, R)

	ProcessByQuick(arr, L, equalArea[0]-1)
	ProcessByQuick(arr, equalArea[1]+1, R)

}

//  小于arr[R]放左侧  等于arr[R]放中间  大于arr[R]放右边
//  返回中间区域的左右边界
func NetherlandsFlag(arr *[]int, L int, R int) []int {

	// 不存在荷兰国旗问题
	if L > R {
		return []int{-1, -1}
	}

	// 已经都是等于区域，由于用R做划分返回R位置
	if L == R {
		return []int{L, R}
	}

	less := L - 1 // 小于区域的右边界
	more := R     // 大于区域的左边界
	index := L
	for index < more {
		// 当前值等于右边界，不做处理，index++
		if (*arr)[index] == (*arr)[R] {
			index++
		} else if (*arr)[index] < (*arr)[R] {
			less++
			Swap(arr, index, less)
			index++
		} else {
			more--
			Swap(arr, index, more)
		}
	}
	// 比较完之后，把R位置的数，调整到等于区域的右边，至此大于区域才是真正意义上的大于区域
	Swap(arr, more, R)
	less++
	return []int{less, more}
}
