package main

import (
	"github.com/NattapatN/path-finder/handler"
)

func main() {
	handler.Setup()
	handler.FindPath("A", "G")
}
