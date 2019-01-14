package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	all = flag.String("a", "all", "contains dot files")
)

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()
	fmt.Println(walk(getPath()))
	return 0
}

func walk(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, walk(filepath.Join(path, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(path, file.Name()))
	}
	return paths
}

func getPath() string {
	var path string
	if args := flag.Args(); len(args) > 0 {
		path = args[0]
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
	} else {
		path = "./"
	}
	return path
}
