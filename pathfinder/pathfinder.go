package pathfinder

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var Maze [][][4]int
var wall map[string]int
var blockX, blockY int
var path [][2]int
var visitedMaze [][]bool
var round = 0

func InitMaze(maze [][][4]int) {
	//define wall
	wall = make(map[string]int)
	wall["up"] = 0
	wall["left"] = 1
	wall["right"] = 2
	wall["down"] = 3
	//assign maze
	Maze = maze
	//assign block x, y
	blockY = len(Maze)
	if blockY > 0 {
		blockX = len(Maze[0])
	}
	initVisitedMaze()
}

func FindPath(startLocation [2]int) {
	currentLocation := startLocation
	path = append(path, currentLocation)
	for currentLocation[0] < blockY-1 || currentLocation[1] < blockX-1 {
		visitedMaze[currentLocation[0]][currentLocation[1]] = true
		availableMove := findAvailableMove(currentLocation)
		if len(availableMove) > 0 {
			path = append(path, availableMove[len(availableMove)-1])
		} else {
			path = path[:len(path)-1]
		}
		if len(path) > 0 {
			currentLocation = path[len(path)-1]
		}
		round++
	}
}

func findAvailableMove(current [2]int) [][2]int {
	availableMove := [][2]int{}
	//check left
	if current[1] > 0 && !visitedMaze[current[0]][current[1]-1] && Maze[current[0]][current[1]][wall["left"]] == 0 {
		availableMove = append(availableMove, [2]int{current[0], current[1] - 1})
	}
	//check right
	if current[1] < blockX-1 && !visitedMaze[current[0]][current[1]+1] && Maze[current[0]][current[1]][wall["right"]] == 0 {
		availableMove = append(availableMove, [2]int{current[0], current[1] + 1})
	}
	//check up
	if current[0] > 0 && !visitedMaze[current[0]-1][current[1]] && Maze[current[0]][current[1]][wall["up"]] == 0 {
		availableMove = append(availableMove, [2]int{current[0] - 1, current[1]})
	}
	//check down
	if current[0] < blockY-1 && !visitedMaze[current[0]+1][current[1]] && Maze[current[0]][current[1]][wall["down"]] == 0 {
		availableMove = append(availableMove, [2]int{current[0] + 1, current[1]})
	}
	return availableMove
}

func initVisitedMaze() {
	for i := 0; i < blockY; i++ {
		visitedMaze = append(visitedMaze, []bool{})
		for j := 0; j < blockX; j++ {
			visitedMaze[i] = append(visitedMaze[i], false)
		}
	}
}

func PrintMazeWithMaze() {
	yellowWithUnderLine := color.New(color.FgHiYellow).Add(color.Underline)
	yellow := color.New(color.FgHiYellow)
	time.Sleep(time.Second / 10)
	// fmt.Printf("\x1bc")
	for i := 0; i < blockX; i++ {
		if i == 0 {
			fmt.Print("   ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
	for i, value := range Maze {
		for j, vj := range value {
			if vj[1] == 1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			isPath := checkIsPath([2]int{i, j})
			if vj[3] == 1 {
				if isPath {
					yellowWithUnderLine.Print("*")
				} else {
					fmt.Print("_")
				}
			} else {
				if isPath {
					yellow.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println("|")
	}
	fmt.Println("Moving :", round, "rounds.")
}

func checkIsPath(location [2]int) bool {
	for _, v := range path {
		if location[0] == v[0] && location[1] == v[1] {
			return true
		}
	}
	return false
}
