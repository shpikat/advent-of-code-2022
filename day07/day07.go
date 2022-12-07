package day07

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	RootDir       = ""
	PathSeparator = "/"
)

var (
	pattern = regexp.MustCompile(`^(dir|\d+) (.+)$`)
)

type File struct {
	name string
	size int
}

type Directory struct {
	name        string
	size        int
	directories []*Directory
	files       []File
}

func (d *Directory) ComputeSize() int {
	sum := 0
	for _, file := range d.files {
		sum += file.size
	}
	for _, dir := range d.directories {
		sum += dir.ComputeSize()
	}
	d.size = sum
	return d.size
}

func part1(input string) (int, error) {
	dirs, err := parseTerminalOutput(input)
	if err != nil {
		return 0, err
	}

	dirs[RootDir].ComputeSize()

	sum := 0
	for _, dir := range dirs {
		if dir.size <= 100000 {
			sum += dir.size
		}
	}

	return sum, nil
}

func part2(input string) (int, error) {
	dirs, err := parseTerminalOutput(input)
	if err != nil {
		return 0, err
	}

	size := dirs[RootDir].ComputeSize()
	required := 30000000 - (70000000 - size)

	min := math.MaxInt
	for _, dir := range dirs {
		if dir.size >= required && dir.size < min {
			min = dir.size
		}
	}

	return min, nil
}

func parseTerminalOutput(output string) (map[string]*Directory, error) {
	dirs := map[string]*Directory{
		RootDir: {
			name: RootDir,
		},
	}
	var path []string
	var fullPath string
	var current *Directory
	for _, line := range strings.Split(strings.TrimSpace(output), "\n") {
		if line[0] == '$' {
			current = nil
			command, argument, _ := strings.Cut(line[2:], " ")
			switch command {
			case "cd":
				switch argument {
				case "/":
					path = []string{RootDir}
				case "..":
					path = path[:len(path)-1]
				default:
					path = append(path, argument)
				}
			case "ls":
				fullPath = strings.Join(path, PathSeparator)
				var exists bool
				current, exists = dirs[fullPath]
				if !exists {
					name := RootDir
					if len(path) != 0 {
						name = path[len(path)-1]
					}
					current = &Directory{
						name: name,
					}
					dirs[fullPath] = current
				}
			}
		} else {
			submatch := pattern.FindStringSubmatch(line)
			name := submatch[2]
			if submatch[1] == "dir" {
				p := fullPath + PathSeparator + name
				dir, exists := dirs[p]
				if !exists {
					dir = &Directory{
						name: name,
					}
					dirs[p] = dir
				}
				current.directories = append(current.directories, dir)
			} else {
				size, err := strconv.Atoi(submatch[1])
				if err != nil {
					return nil, err
				}
				file := File{
					name: name,
					size: size,
				}
				current.files = append(current.files, file)
			}
		}
	}

	return dirs, nil
}
