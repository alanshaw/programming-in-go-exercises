package main

import "fmt"

func main () {
  nums := []int{1, 4, 5, 9}
  total := sum(nums)
  fmt.Println(total)
}

func sum (nums []int) int {
  total := 0
  for i := 0; i < len(nums); i++ {
    total += nums[i]
  }
  return total
}