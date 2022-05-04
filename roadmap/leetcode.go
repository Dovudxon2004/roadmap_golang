package roadmap

import (
	"sort"
	"strconv"
	"strings"
)

func ContainsDUplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func ContainsNearbyDuplicate(nums []int, k int) bool {
	for i, v := range nums {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if nums[j] == v {
				return true
			}
		}
	}
	return false
}
func ContainsNearbyDuplicateWithMap(nums []int, k int) bool {
	record := make(map[int]bool)
	for i, v := range nums {
		if record[v] {
			return true
		}
		record[v] = true
		if len(record) > k+1 {
			delete(record, nums[i-k])
		}

	}
	return false
}

func IsPalindrome(x int) bool {
	s := strconv.Itoa(x)
	runes := make([]string, len(s))
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		runes[i], runes[j] = string(s[j]), string(s[i])
	}
	s1 := strings.Join(runes, "")
	if s1 == s {
		return true
	}
	return false
}

func SumOddLengthSubarrays(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += (((i+1)*(len(arr)-i) + 1) / 2) * arr[i]
	}
	return sum
}
func CountEven(num int) int {
	sum := 0
	sum = num/1000 + (num%1000)/100 + ((num%1000)%100)/10 + ((num%1000)%100)%10
	return (num - sum%2) / 2
}
