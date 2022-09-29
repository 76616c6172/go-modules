package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// like input() in python
func Input(Msg string) string {
	fmt.Println(Msg)
	input_buffer := bufio.NewReader(os.Stdin)
	usr_input, _ := input_buffer.ReadString('\n')
	return strings.TrimSpace(usr_input)
}
