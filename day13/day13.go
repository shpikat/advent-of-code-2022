package day13

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
)

func part1(input string) (int, error) {
	pairs := bytes.Split(bytes.TrimSpace([]byte(input)), []byte("\n\n"))

	sum := 0
	for index, pair := range pairs {
		leftRaw, rightRaw, found := bytes.Cut(pair, []byte("\n"))
		if !found {
			return 0, errors.New("expected a pair: " + string(pair))
		}

		var left, right []interface{}
		err := json.Unmarshal(leftRaw, &left)
		if err != nil {
			return 0, err
		}
		err = json.Unmarshal(rightRaw, &right)
		if err != nil {
			return 0, err
		}

		compared := Compare(left, right)
		if compared < 0 {
			sum += index + 1
		}
	}
	return sum, nil
}

func part2(input string) (int, error) {
	lines := bytes.Split(bytes.TrimSpace([]byte(input)), []byte("\n"))

	var packets [][]interface{}
	for _, line := range lines {
		if len(line) != 0 {
			var packet []interface{}
			err := json.Unmarshal(line, &packet)
			if err != nil {
				return 0, err
			}
			packets = append(packets, packet)

		}
	}

	var dividers [][]interface{}
	for _, s := range [][]byte{[]byte("[[2]]"), []byte("[[6]]")} {
		var divider []interface{}
		err := json.Unmarshal(s, &divider)
		if err != nil {
			return 0, err
		}
		dividers = append(dividers, divider)
	}

	packets = append(packets, dividers...)

	sort.Slice(packets, func(i, j int) bool {
		return Compare(packets[i], packets[j]) < 0
	})

	key := 1
	for _, divider := range dividers {
		n, found := sort.Find(len(packets), func(i int) int {
			return Compare(divider, packets[i])
		})
		if !found {
			return 0, fmt.Errorf("divider %v not found", divider)
		}
		key *= n + 1
	}

	return key, nil
}

func Compare(left, right []interface{}) int {
	for i := 0; i < len(left) && i < len(right); i++ {
		// JSON representation of Number is float64
		leftNumber, isLeftNumber := left[i].(float64)
		rightNumber, isRightNumber := right[i].(float64)
		if isLeftNumber && isRightNumber {
			leftInteger := int(leftNumber)
			rightInteger := int(rightNumber)
			if leftInteger != rightInteger {
				return leftInteger - rightInteger
			}
		} else {
			var leftList, rightList []interface{}
			if isLeftNumber {
				leftList = []interface{}{leftNumber}
			} else {
				leftList = left[i].([]interface{})
			}
			if isRightNumber {
				rightList = []interface{}{rightNumber}
			} else {
				rightList = right[i].([]interface{})
			}
			cmp := Compare(leftList, rightList)
			if cmp != 0 {
				return cmp
			}
		}
	}
	return len(left) - len(right)
}
