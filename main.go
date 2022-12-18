package main

import (
	"fmt"

	"github.com/NattapatN/maze-generator/generator"
	"github.com/NattapatN/path-finder/pathfinder"
)

var maze [][][4]int

func main() {
	generator.Initial(40, 30)
	maze = generator.GetMaze()
	generator.PrintMaze()

	pathfinder.InitMaze(maze)
	pathfinder.FindPath([2]int{0, 0})
	fmt.Println("Resolve Path :")
	pathfinder.PrintMazeWithMaze()
}
