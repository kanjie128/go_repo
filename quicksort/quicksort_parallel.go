package main

import "fmt"

func quicksort(nums []int, ch chan int, level int, threads int) {
	level = level * 2
	if len(nums) == 1 {
		ch <- nums[0]
		close(ch)
		return
	}
	if len(nums) == 0 {
		close(ch)
		return
	}

	less := make([]int, 0)
	greater := make([]int, 0)
	left := nums[0]
	nums = nums[1:]

	for _, num_data := range nums {
		switch {
		case num_data <= left:
			less = append(less, num_data)
		case num_data > left:
			greater = append(greater, num_data)
		}
	}

	left_ch := make(chan int, len(less))
	right_ch := make(chan int, len(greater))

	if level <= threads {
		go quicksort(less, left_ch, level, threads)
		go quicksort(greater, right_ch, level, threads)
	} else {
		quicksort(less, left_ch, level, threads)
		quicksort(greater, right_ch, level, threads)
	}

	for i := range left_ch {
		ch <- i
	}
	ch <- left
	for i := range right_ch {
		ch <- i
	}
	close(ch)
	return
}

func main() {
	x := []int{12, 34, 55, 29, 43, 32, 11, 1, 2}
	ch := make(chan int)
	go quicksort(x, ch, 0, 0)
	for v := range ch {
		fmt.Println(v)
	}
}
