package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

var directions = []Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} // Up, Left, Down, Right

func solveMaze(maze [][]int, start, end Point) []Point {
	var path []Point
	visited := make(map[Point]bool)
	stack := []Point{start}

	// Iterate until the stack is empty
	for len(stack) > 0 {
		// Pop a point from the stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// If the current point is the end point, return the path
		if current == end {
			path = append(path, current)
			return path
		}

		if isValidPoint(current, maze) && !visited[current] {
			path = append(path, current)
			visited[current] = true

			// Explore valid neighbors
			for _, dir := range directions {
				next := Point{current.x + dir.x, current.y + dir.y}
				// Check if the next point is within bounds, not a wall, and not visited
				if isValidPoint(next, maze) && !visited[next] {
					stack = append(stack, next)
				}
			}
		} else {
			// If the current point is invalid or visited, skip it
			continue
		}
	}

	// If the end point is not reached, return an empty path
	return nil
}

func isValidPoint(p Point, maze [][]int) bool {
	return p.x >= 0 && p.x < len(maze) && p.y >= 0 && p.y < len(maze[0]) && maze[p.x][p.y] != 1
}

func main() {
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	// Define start and end points
	start := Point{0, 0}
	end := Point{4, 3}

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
