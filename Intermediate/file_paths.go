package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	relativePath := "./data/file.txt"
	absolutePath := "C:/users/xyz"

	joinedPath := filepath.Join("downloads", "file.zip")
	fmt.Println(joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	println(normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	println(file)
	println(dir)

	println(filepath.IsAbs(absolutePath))
	println(filepath.Ext(relativePath))

	absPath, _ := filepath.Abs(relativePath)
	println(absPath)
}
