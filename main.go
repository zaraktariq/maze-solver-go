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

func isValidPt(p Point, maze Maze) bool {
	// Check if the point is within the bounds of the maze
	return p.x >= 0 && p.x < len(maze) && p.y >= 0 && p.y < len(maze[0]) &&
		// Check if the cell at the point's coordinates is not a wall
		maze[p.x][p.y] != 1
}

func solveMaze(maze Maze, start, end Point) []Point {
	// Initialize an empty stack and a map to keep track of visited points
	stack := []Point{start}
	visited := make(map[Point]bool)

	// Iterate until the stack is empty
	for len(stack) > 0 {
		// Pop a point from the stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// If the current point is the end point, return the path
		if current == end {
			return []Point{start, end}
		}

		// Mark the current point as visited
		visited[current] = true

		// Explore valid neighbors
		for _, dir := range directions {
			next := Point{current.x + dir.x, current.y + dir.y}
			// Check if the next point is within bounds, not a wall, and not visited
			if isValidPt(next, maze) && !visited[next] {
				stack = append(stack, next)
			}
		}
	}

	// If the end point is not reached, return an empty path
	return nil
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
