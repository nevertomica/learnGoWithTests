package arrays

// 將欲加總的數放進同一個 slices，並傳入 SumAll 函式，就可以得到總和
// 不過總和會是以一個 slice 回傳
// func 透過初始一個 empty slice，後續透過 append 加添元素
func SumAll(numberArrays ...[]int) []int {

	arrayCapability := len(numberArrays)
	sumArray := make([]int, 0, arrayCapability)

	for _, numbers := range numberArrays {
		// 只要不超過 sumArray 底層陣列長度，都會指向同一個底層陣列
		sumArray = append(sumArray, Sum(numbers))
	}

	return sumArray
}

// 加總 slice 中的所有元素
func Sum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll 的改良，加總每一個除去第一個元素的 slice 總和
// 如果傳入 []int{3} 加總是 []int{0}，另外如果是空的 slice 也是回傳 []int{0}
func SumAllTails(slices ...[]int) []int {
	subSlice := make([][]int, 0, len(slices))
	for _, arr := range slices {
		if len(arr) > 1 {
			subSlice = append(subSlice, arr[1:])
		} else {
			subSlice = append(subSlice, []int{})
		}
	}
	return SumAll(subSlice...)

}
