package day20

import (
	"strconv"
	"strings"
)

type Node struct {
	value int
	prev  *Node
	next  *Node
}

func (n *Node) moveBefore(other *Node) {
	n.prev.next, n.next.prev = n.next, n.prev
	n.prev, n.next = other.prev, other
	other.prev.next, other.prev = n, n
}

type File struct {
	data        []Node
	indexOfZero int
}

func (f File) mix() {
	// When wrapping we deduct one to account for the number actually being moved
	wrap := len(f.data) - 1

	for i := range f.data {
		current := &f.data[i]
		if current.value != 0 {
			moves := current.value % wrap
			// We could have chosen the shorter path back or forth, but it's less code this way
			if moves < 0 {
				moves += wrap
			}
			target := current.next
			for i := 0; i < moves; i++ {
				target = target.next
			}

			current.moveBefore(target)
		}
	}
}

func (f File) getCoordinates() int {
	sum := 0
	current := &f.data[f.indexOfZero]
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			current = current.next
		}
		sum += current.value
	}
	return sum
}

func (f File) applyDecryptionKey(decryptionKey int) {
	for i := range f.data {
		f.data[i].value *= decryptionKey
	}
}

func NewFile(lines []string) (File, error) {
	var file File
	file.data = make([]Node, len(lines))

	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return file, err
		}

		file.data[i].value = n
		if n == 0 {
			file.indexOfZero = i
		}
	}

	size := len(file.data)
	last := size - 1
	for i := range file.data[:last] {
		file.data[i].next = &file.data[i+1]
		file.data[i+1].prev = &file.data[i]
	}
	file.data[last].next = &file.data[0]
	file.data[0].prev = &file.data[last]
	return file, nil
}

func part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	file, err := NewFile(lines)
	if err != nil {
		return 0, err
	}

	file.mix()
	sum := file.getCoordinates()

	return sum, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	file, err := NewFile(lines)
	if err != nil {
		return 0, err
	}

	file.applyDecryptionKey(811589153)
	for i := 0; i < 10; i++ {
		file.mix()
	}
	sum := file.getCoordinates()

	return sum, nil
}
