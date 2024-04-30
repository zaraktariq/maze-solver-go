## Maze Solver
This Go program is designed to solve mazes loaded from a text file. It provides functionality to read maze data from the file, display available mazes, allow the user to choose a maze to solve, and then find a path from the start to the end of the maze.

### Usage
 - Load Mazes: The program reads maze data from a text file named maze.txt.
 - Select Maze: After loading the maze data, it displays the available mazes and prompts the user to choose a maze to solve.
 - Solve Maze: Once the user selects a maze, the program employs a depth-first search algorithm to find a path from the start to the end of the maze.
 - Display Path: If a path is found, it displays the coordinates of each point along the path. If no path is found, it informs the user accordingly.

### File Format
 - The maze data in the text file should be formatted as follows:
 - Each maze is separated by a # delimiter.
 - Each line within a maze represents a row of the maze grid.
 - Each value in a row represents a cell in the maze, where 0 denotes an empty cell and 1 denotes a wall.
