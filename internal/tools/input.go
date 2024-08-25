package tools

import (
	"bufio"
	"fmt"
	"strings"
)

func ReadInput(prompt string, r *bufio.Reader) string {
	fmt.Printf("%v: ", prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		return ""
	}
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSpace(input)
	return input
}
