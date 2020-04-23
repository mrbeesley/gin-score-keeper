package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mrbeesley/gin_score_keeper/controllers"
)

func main() {
	gm := controllers.NewGameManager()
	input := ""
	for {
		readOnly, msg := controllers.ManageGame(gm, input)
		if msg == "break" {
			break
		}
		fmt.Println(msg)
		if !readOnly {
			i, err := handleInput()
			if err != nil {
				break
			} else {
				input = i
			}
		}
	}
}

func handleInput() (string, error) {
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		println("There was an Error reading input: ", err.Error())

	} else {
		line = strings.TrimSuffix(line, "\n")
	}
	return line, err
}
