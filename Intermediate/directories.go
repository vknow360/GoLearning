package main

import (
	"os"
	"path/filepath"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	handle(os.Mkdir("subdir", 0755))
	handle(os.Remove("subdir"))
	os.WriteFile("subdir/file", []byte(""), 0755)

	os.MkdirAll("subdir/parent/child", 0755)

	result, err := os.ReadDir("./Intermediate")
	handle(err)

	for _, entry := range result {
		println(entry.Name(), entry.IsDir(), entry.Type())
	}

	handle(filepath.WalkDir("./Intermediate", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			println(err)
			return err
		}
		println(path)
		return nil
	}))

	handle(os.RemoveAll("./subdir"))
}
