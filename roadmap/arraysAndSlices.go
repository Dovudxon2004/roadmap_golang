package roadmap

import "fmt"

func ArraysDontChange() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr1 := arr
	arr1[0] = 0
	fmt.Println(arr, arr1)
}
func SlicesChange() {
	sl := []int{1, 2, 3, 4, 5}
	sl1 := sl
	fmt.Printf("%d, %d", len(sl1), cap(sl1))
	sl1 = append(sl1, 6)
	fmt.Printf("%d, %d", len(sl1), cap(sl1))
	sl1[0] = 0
	fmt.Println(sl1, sl)
}
