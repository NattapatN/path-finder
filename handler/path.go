package handler

import "fmt"

func findAllPosiblePath(start, end string) [][]string {
	pathList := [][]string{}
	//find connected path
	connectedPath := findConnectedPath(start)
	fmt.Println("Connected path :", connectedPath)

	return pathList
}

func findConnectedPath(point string) []string {
	connectecPath := []string{}
	for _, v := range paths {
		if v.a == point {
			connectecPath = append(connectecPath, v.b)
		} else if v.b == point {
			connectecPath = append(connectecPath, v.a)
		}
	}

	return connectecPath
}
