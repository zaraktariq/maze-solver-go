package main

import (
	"fmt"
	_ "fmt"
)

type Point struct {
	x int
	y int
}

type Maze [][]int

func (maze Maze) PrintMaze() {
	for _, row := range maze {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

func main() {
	// Initialize a sample maze
	maze := Maze{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	// Print the maze
	maze.PrintMaze()
}
