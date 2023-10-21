package day21

import (
	"errors"
	"strconv"
	"strings"
)

type Job struct {
	left       string
	right      string
	operation  string
	value      int
	calculated bool
}

type Jobs map[string]*Job

type Reversed struct {
	reverse Operation
	known   string
	next    string
}

type Operation func(a, b int) int

func (j Jobs) solve(name string) int {
	job := j[name]
	if !job.calculated {
		job.value = Operations[job.operation](j.solve(job.left), j.solve(job.right))
		job.calculated = true
	}
	return job.value
}

func (j Jobs) reverseFor(name string) map[string]Reversed {
	path := make(map[string]Reversed, len(j))
	j.find(path, Root, name)
	return path
}

func (j Jobs) find(path map[string]Reversed, branch string, name string) bool {
	if branch == "" {
		return false
	}
	if branch == name {
		return true
	}
	job := j[branch]
	if j.find(path, job.left, name) {
		path[branch] = Reversed{
			ReverseOperationsForFirstOperand[job.operation],
			job.right,
			job.left,
		}
		return true
	}
	if j.find(path, job.right, name) {
		path[branch] = Reversed{
			ReverseOperationsForSecondOperand[job.operation],
			job.left,
			job.right,
		}
		return true
	}

	return false
}

func (j Jobs) parse(lines []string) error {
	for _, line := range lines {
		name, description, found := strings.Cut(line, ": ")
		if !found {
			return errors.New("unexpected format for line " + line)
		}
		splits := strings.Split(description, " ")

		var job Job
		if len(splits) == 1 {
			n, err := strconv.Atoi(description)
			if err != nil {
				return err
			}
			job.value = n
			job.calculated = true
		} else {
			job.left = splits[0]
			job.right = splits[2]
			job.operation = splits[1]
		}
		j[name] = &job
	}
	return nil
}

const (
	Root  = "root"
	Human = "humn"
)

var (
	Operations = map[string]Operation{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	ReverseOperationsForFirstOperand = map[string]Operation{
		"+": func(result, b int) int { return result - b },
		"-": func(result, b int) int { return result + b },
		"*": func(result, b int) int { return result / b },
		"/": func(result, b int) int { return result * b },
	}
	ReverseOperationsForSecondOperand = map[string]Operation{
		"+": func(result, a int) int { return result - a },
		"-": func(result, a int) int { return a - result },
		"*": func(result, a int) int { return result / a },
		"/": func(result, a int) int { return a / result },
	}
)

func part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	jobs := make(Jobs, len(lines))
	err := jobs.parse(lines)
	if err != nil {
		return 0, err
	}

	answer := jobs.solve(Root)

	return answer, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	jobs := make(Jobs, len(lines))
	err := jobs.parse(lines)
	if err != nil {
		return 0, err
	}

	// If we expect both branches to match, we can say it's a subtraction with the result of 0
	jobs[Root].operation = "-"
	expected := 0

	reversed := jobs.reverseFor(Human)

	current := Root
	for current != Human {
		r := reversed[current]
		expected = r.reverse(expected, jobs.solve(r.known))
		current = r.next
	}

	return expected, nil
}
