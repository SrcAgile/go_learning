package Array

import (
	"testing"
)

// func TestArray(t *testing.T) {
// 	fmt.Println("array")
// }
//
// func TestArrayTraval(t *testing.T) {
// 	arr3 := [...]int{1, 2, 3, 4, 5}
// 	for i := 0; i < len(arr3); i++ {
// 		t.Log(arr3[i])
// 	}
// 	for idx, val := range arr3 {
// 		t.Log(idx, val)
// 	}
// }
//
// func TestTruc(t *testing.T) {
// 	//	arr4 := []int{1, 2, 3, 4, 5}
// 	// arr4 = append(arr4, 5)
// 	arr5 := [5]int{1, 2, 3, 4, 5}
// 	arr1 := arr5[1:len(arr5)]
// 	for idx, val := range arr1 {
// 		t.Log(idx, val)
// 	}
// 	t.Log()
// }

func TestSliceAnaylise(t *testing.T) {
	arr1 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := arr1[:2]
	arr3 := arr1[:2]
	arr2[1] = 10
	t.Log(arr3[1])
	// arr4 := make([]int, 3, 5)
	// t.Log(len(arr4), cap(arr4))
	// //t.Log(arr2 == arr3)
}

// func TestSliceGrowing(t *testing.T) {
// 	arr1 := []int{}
// 	for i := 0; i < 10; i++ {
// 		arr1 = append(arr1, i)
// 		t.Log(len(arr1), cap(arr1))
// 	}
// }
