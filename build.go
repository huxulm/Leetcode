package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

func main() {
	root := "."
	entries := getEntries(root)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	start, step := 1, 49
	for _, entry := range entries {
		num, err := strconv.Atoi(entry.Name()[:strings.Index(entry.Name(), ".")])
		if err != nil {
			continue
		}

		for num > start+step {
			start += step + 1
		}

		numDir := fmt.Sprintf("%04d-%04d", start, start+step)
		if _, err := os.Stat(path.Join(root, numDir)); os.IsNotExist(err) {
			os.Mkdir(path.Join(root, numDir), 0777)
		}

		moveDir(entry.Name(), path.Join(root, numDir, entry.Name()))
	}
}

func moveDir(from, to string) {
	os.Rename(from, to)
}

func getEntries(root string) (res []fs.DirEntry) {
	cnt := 0
	entries, _ := os.ReadDir(".")
	res = make([]fs.DirEntry, 0, len(entries))
	for i := range entries {
		if entries[i].IsDir() && strings.Index(entries[i].Name(), ".") > 0 {
			res = append(res, entries[i])
			cnt++
		}
	}
	return res[:cnt]
}
