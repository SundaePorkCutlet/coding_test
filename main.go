package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var treeCount int
var wantHeight int
var treeHeight []int
var maxTreeHeight = 0
var result = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	treeCount, _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	wantHeight, _ = strconv.Atoi(scanner.Text())

	treeHeight = make([]int, treeCount)

	for i := 0; i < treeCount; i++ {
		scanner.Scan()
		treeHeight[i], _ = strconv.Atoi(scanner.Text())
		if treeHeight[i] > maxTreeHeight {
			maxTreeHeight = treeHeight[i]
		}
	}

	left := 0
	right := maxTreeHeight
	for left <= right {
		mid := (left + right) / 2
		if check(mid) {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println(result)

}

func check(cutterHeight int) bool {
	max := 0
	for i := 0; i < len(treeHeight); i++ {
		if treeHeight[i] > cutterHeight {
			max += treeHeight[i] - cutterHeight
		}
	}
	return max >= wantHeight
}
