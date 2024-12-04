package day4

import (
	"aoc2024/cmd"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Coord struct {
	x int
	y int
}

var directions = []Coord{
	{1, 1},
	{1, 0},
	{1, -1},
	{0, 1},
	{0, -1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

var diagonal1 = []Coord{
	{1, 1},
	{-1, -1},
}
var diagonal2 = []Coord{
	{1, -1},
	{-1, 1},
}

var day1_1Cmd = &cobra.Command{
	Use:   "day4",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day1_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day4_data.txt")
	ans1, ans2 := Solve(string(dat))

	fmt.Printf("ANSWER 1: %d\n", ans1)
	fmt.Printf("ANSWER 2: %d\n", ans2)
}

func Solve(data string) (int, int) {
	ans1 := 0
	ans2 := 0

	var grid [][]rune

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		var runes []rune
		for _, r := range line {
			runes = append(runes, r)
		}
		grid = append(grid, runes)
	}

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'X' {
				ans1 += StartXMASSearch(&grid, x, y)
			} else if grid[x][y] == 'A' {
				ans2 += StartX_MASSearch(&grid, x, y)
			}
		}
	}

	return ans1, ans2
}

func StartXMASSearch(grid *[][]rune, x int, y int) int {
	results := 0
	searchSlice := []rune{'M', 'A', 'S'}
	for _, d := range directions {
		results += Search(grid, x, y, d.x, d.y, searchSlice)
	}

	return results
}
func Search(grid *[][]rune, x int, y int, dx int, dy int, searchSlice []rune) int {
	if len(searchSlice) == 0 {
		return 1
	}

	pop, searchSlice := searchSlice[0], searchSlice[1:]

	x += dx
	y += dy

	if x < 0 || y < 0 || x == len(*grid) || y == len((*grid)[0]) {
		return 0
	}

	if (*grid)[x][y] == pop {
		return Search(grid, x, y, dx, dy, searchSlice)
	}

	return 0
}

func StartX_MASSearch(grid *[][]rune, x int, y int) int {
	diag1_M := false
	diag1_S := false
	diag2_M := false
	diag2_S := false
	for _, d := range diagonal1 {
		dx := x + d.x
		dy := y + d.y

		if dx < 0 || dy < 0 || dx == len(*grid) || dy == len((*grid)[0]) {
			continue
		}

		r := (*grid)[dx][dy]
		if r == 'M' {
			diag1_M = true
		} else if r == 'S' {
			diag1_S = true
		}
	}

	for _, d := range diagonal2 {
		dx := x + d.x
		dy := y + d.y

		if dx < 0 || dy < 0 || dx == len(*grid) || dy == len((*grid)[0]) {
			continue
		}

		r := (*grid)[dx][dy]
		if r == 'M' {
			diag2_M = true
		} else if r == 'S' {
			diag2_S = true
		}
	}

	if diag1_M && diag1_S && diag2_M && diag2_S {
		return 1
	}

	return 0
}
