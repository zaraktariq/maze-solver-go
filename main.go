package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func loadMazeFromFile(filename string) ([][][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var mazes [][][]int
	var maze [][]int
	scanner := bufio.NewScanner(file)

	// Check if the line is a delimiter
	for scanner.Scan() {
		// Read the next line from the file
		line := scanner.Text()

		// Check if the line is a delimiter
		if line == "#" {
			// Add the current maze to the list of mazes
			mazes = append(mazes, maze)
			// Reset the maze slice for the next maze
			maze = nil
		} else {
			// If the line is not a delimiter, it contains maze data
			// Process the line as maze data
			// Split the line into individual values
			parts := strings.Fields(line)
			// Create a slice to store the values of the current row in the maze
			row := make([]int, len(parts))
			// Iterate over the parts of the line and convert them to integers
			for i, part := range parts {
				value, err := strconv.Atoi(part)
				if err != nil {
					// If an error occurs while converting, return the error
					return nil, err
				}
				// Store the converted value in the current row
				row[i] = value
			}
			// Append the current row to the maze
			maze = append(maze, row)
		}
	}
	return mazes, nil
}

func main() {
	mazes, err := loadMazeFromFile("maze.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Available Mazes:")
	for i := range mazes {
		fmt.Printf("Maze %d\n", i+1)
	}

	var mazeChoice int
	fmt.Println("Enter the maze number you want to solve:")
	fmt.Scanf("%d", &mazeChoice)

	// Validate user input
	if mazeChoice < 1 || mazeChoice > len(mazes) {
		fmt.Println("Invalid maze number. Please choose a number between 1 and", len(mazes))
		return
	}

	mazeIndex := mazeChoice - 1
	maze := mazes[mazeIndex]

	// Define start and end points (assuming they are same for all mazes)
	start := Point{0, 0}
	end := Point{len(maze) - 1, len(maze[0]) - 1}

	// Solve the maze
	path := solveMaze(maze, start, end)

	// Print the path
	fmt.Printf("Maze %d:\n", mazeChoice)
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
