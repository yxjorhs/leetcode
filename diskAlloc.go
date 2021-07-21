package main

import (
	"fmt"
)

func main() {
	diskSize := 0
	v := 0
	arr := []int{}

	fmt.Scan(&diskSize)
	for {
		sc, _ := fmt.Scan(&v)
		if sc == 0 {
			break
		}
		arr = append(arr, v)
	}
	fmt.Printf("%d\n", diskAlloc(diskSize, arr))
}

func diskAlloc(diskSize int, files []int) int {
	stack := []int{}
	stackSize := 0
	minLeave := diskSize

	for i := 0; i < len(files); i++ {
		if files[i] >= diskSize {
			stack = []int{}
			stackSize = 0

			if files[i] > diskSize {
				continue
			}
		}
		for stackSize+files[i] > diskSize {
			stackSize -= stack[0]
			stack = stack[1:]
		}
		stack = append(stack, files[i])
		stackSize += files[i]
		if diskSize-stackSize < minLeave {
			minLeave = diskSize - stackSize
		}
	}

	return minLeave
}
