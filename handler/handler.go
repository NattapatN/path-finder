package handler

import "fmt"

type Path struct {
	a      string
	b      string
	weight int
}

var paths []Path

func Setup() {
	paths = append(paths, Path{a: "A", b: "B", weight: 4})
	paths = append(paths, Path{a: "A", b: "D", weight: 3})
	paths = append(paths, Path{a: "B", b: "E", weight: 3})
	paths = append(paths, Path{a: "D", b: "E", weight: 2})
	paths = append(paths, Path{a: "B", b: "F", weight: 6})
	paths = append(paths, Path{a: "E", b: "G", weight: 1})
	paths = append(paths, Path{a: "B", b: "G", weight: 5})
	paths = append(paths, Path{a: "F", b: "G", weight: 1})
}

func FindPath(start, stop string) {
	allpath := findAllPosiblePath(start, stop)
	fmt.Println("All posible path :", allpath)
}
