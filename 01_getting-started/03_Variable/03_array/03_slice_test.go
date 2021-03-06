/*

空数组， 这个就和Python的 list 很相似了

https://www.digitalocean.com/community/tutorials/understanding-arrays-and-slices-in-go

slice

在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。
在Go里面这种数据结构叫slice

slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。

对于slice有几个有用的内置函数：

len 获取slice的长度
cap 获取slice的最大容量
append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数

*/
package _3_array

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// 和声明array一样，只是少了长度
	var fslice [10]int
	fmt.Println(fslice)

	slice := []byte{'a', 'b', 'c', 'd'}
	var z []byte
	z = slice[0:1]

	/*
		len 获取slice的长度
		cap 获取slice的最大容量
		append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
		copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
	*/
	fmt.Println(z)
	fmt.Println(cap(fslice)) //
	fmt.Println(len(slice))

	var array [10]int
	var newSlice []int
	newSlice = array[2:4:7]
	fmt.Println(newSlice)
}

func TestSliceCURD(t *testing.T) {
	// 切片的操作
	strSlice := []string{"a", "z", "1", "2", "3"}

	// 添加
	strSlice = append(strSlice, "666")

	t.Log(strSlice)

	// 修改 和数组一样
	strSlice[0] = "aaaa"
	t.Log(strSlice)

	// 删除使用append实现 重新赋值
	// 删除第一个元素 [1:]
	//strSlice = strSlice[1:]
	//
	//t.Log(strSlice)
	//
	//// 删除第二个元素  a = append(a[:i], a[i+1:]...) // 删除中间第i个元素(只删除一个)
	//strSlice = append(strSlice[:1], strSlice[2:]...)
	//t.Log(strSlice)
	//
	//strSlice = remove(strSlice, 2)
	//
	//t.Log(strSlice)
	//
	//strSlice = newRemove01(strSlice, 1)
	//t.Log(strSlice)

	strSlice = newRemove02(strSlice, 1)
	t.Log(strSlice)
}

// 利用上面的方式 封装程一个函数 效率很低 因为是copy重新赋值操作
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func newRemove01(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

//参考 https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func newRemove02(s []string, i int) []string {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
