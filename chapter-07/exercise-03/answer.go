package main

import "fmt"

func main () {
  fmt.Println(greatest(3, 6, 2, 87, 138, 4))
}

func greatest (nums...int) int {
  if len(nums) == 0 {
    panic("List of numbers expected")
  }

  max := nums[0]

  for i := 1; i < len(nums); i++ {
    if nums[i] > max {
      max = nums[i]
    }
  }

  return max
}