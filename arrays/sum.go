package arrays

// 將欲加總的數放進同一個 slices，並傳入 SumAll 函式，就可以得到總和
// 不過總和會是以一個 slice 回傳
func SumAll(numberArrays ...[]int) []int {

	arrayCapability := len(numberArrays)
	sumArray := make([]int, arrayCapability)

	for i, numbers := range numberArrays {
		sumArray[i] = Sum(numbers)
	}

	return sumArray
}

func Sum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}
