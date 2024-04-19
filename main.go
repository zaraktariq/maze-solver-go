package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Maze [][]int

// Directions representing the four cardinal directions
var directions = []Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (maze Maze) PrintMaze() {
	// iterate over rows
	for _, row := range maze {
		// iterate over cells in current row
		for _, cell := range row {
			// print
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

func isValidPoint(p Point, maze Maze) bool {
	// Check if the point is within the maze
	if p.x < 0 || p.x >= len(maze[0]) ||
		p.y < 0 || p.y >= len(maze) {
		return false
	}
	// Check if the current point is not a wall
	if maze[p.y][p.x] == 1 {
		return false
	}
	// Otherwise, point is valid
	return true
}

// solveMaze solves the maze using depth-first search (DFS) with an explicit stack
func solveMaze(maze Maze, start, end Point) []Point {
	// Initialize an empty stack and a map to keep track of visited points
	stack := make([]Point, 0)
	visited := make(map[Point]bool)

	// Push the start point onto the stack
	stack = append(stack, start)

	// Iterate until the stack is empty
	for len(stack) > 0 {
		// Pop a point from the stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// If the current point is the end point, return the path
		if current == end {
			// Reconstruct and return the path
			return reconstructPath(start, current, visited, maze)
		}

		// Mark the current point as visited
		visited[current] = true

		// Push valid neighbors onto the stack
		for _, dir := range directions {
			next := Point{current.x + dir.x, current.y + dir.y}
			if isValidPoint(next, maze) && !visited[next] {
				stack = append(stack, next)
			}
		}
	}

	// If the end point is not reached, return an empty path
	return nil
}

// reconstructPath reconstructs the path from start to end point
func reconstructPath(start, end Point, visited map[Point]bool, maze Maze) []Point {
	path := make([]Point, 0)
	current := end
	for current != start {
		path = append([]Point{current}, path...)
		for _, dir := range directions {
			previous := Point{current.x - dir.x, current.y - dir.y}
			if isValidPoint(previous, maze) && visited[previous] {
				current = previous
				break
			}
		}
	}
	return append([]Point{start}, path...)
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

	// Define start and end points
	start := Point{0, 0}
	end := Point{4, 4}

	// Solve the maze
	path := solveMaze(maze, start, end)

	// Print the maze
	maze.PrintMaze()

	// Print the path
	if path != nil {
		fmt.Println("Path found:")
		for _, p := range path {
			fmt.Printf("(%d, %d) -> ", p.x, p.y)
		}
		fmt.Println("End")
	} else {
		fmt.Println("No path found")
	}
}
