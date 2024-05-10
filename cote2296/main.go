package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var buildingCount int
var quarter [2]int
var dp [][2]int
var buildings []building

type building struct {
	x      int
	y      int
	profit int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	buildingCount, _ = strconv.Atoi(scanner.Text())

	buildings = make([]building, buildingCount)
	dp = make([][2]int, buildingCount)

	for i := 0; i < buildingCount; i++ {
		scanner.Scan()
		tempBuild := building{}
		tempBuild.x, _ = strconv.Atoi(scanner.Text())

		scanner.Scan()
		tempBuild.y, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		tempBuild.profit, _ = strconv.Atoi(scanner.Text())
		buildings[i] = tempBuild

	}

	//1,3분면 최대값
	sort.Slice(buildings, func(i, j int) bool {
		if buildings[i].x == buildings[j].x {
			return buildings[i].y < buildings[j].y
		}
		return buildings[i].x < buildings[j].x
	})

	result := 0

	for i := 0; i < buildingCount; i++ {
		dp[i][0] = buildings[i].profit
		for j := 0; j < i; j++ {
			if buildings[j].y < buildings[i].y {
				dp[i][0] = max(dp[i][0], dp[j][0]+buildings[i].profit)
			}

		}
		result = max(result, dp[i][0])
	}

	//2,4분면 최대값
	sort.Slice(buildings, func(i, j int) bool {
		if buildings[i].x == buildings[j].x {
			return buildings[i].y < buildings[j].y
		}
		return buildings[i].x > buildings[j].x
	})

	for i := 0; i < buildingCount; i++ {
		dp[i][1] = buildings[i].profit
		for j := 0; j < i; j++ {
			if buildings[j].y < buildings[i].y {
				dp[i][1] = max(dp[i][1], dp[j][1]+buildings[i].profit)
			}
		}
		result = max(result, dp[i][1])
	}

	fmt.Println(result)

}
