package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

type Maze [][]int

func (maze Maze) PrintMaze() {
	// iterate iver rows
	for _, row := range maze {
		// iterate over cells in current row
		for _, cell := range row {
			// print
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

func isValidPt(p Point, maze Maze) int {
	// Check if the point is within the maze
	if p.x < 0 || p.x > len(maze) ||
		p.y < 0 || p.y > len(maze[0]) {
		return -1
	}
	// Check if the current point is not a wall
	if maze[p.y][p.x] == 1 {
		return -1
	}
	// Otherwise, return the value of the cell at the given point
	return maze[p.y][p.x]
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
