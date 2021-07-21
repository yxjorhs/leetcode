package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello World!\n")
}

// 障碍点值为1
const obstacleVal = 1


type findWayStore struct {
	setup
}

func obstacleSet(mp *[][]int, x int, y int) {
	
}

// {mp}二维数组地图
func findWay(mp *[][]int) {
	// 动态规划，从点(0,0)出发，尝试向右、下移动，记录到每个格子所需的最小步数
	store := [][]
}

func move() {}
