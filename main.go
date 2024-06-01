package main

import (
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		os.Stdout.WriteString("provide an input file")
		os.Exit(1)
	}

	file, err := os.ReadFile(os.Args[1])

	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}

	interpreter := new_interpreter(parse(file))

	for i, arg := range os.Args[2:] {
		as_int, err := strconv.Atoi(arg)
		if err != nil {
			os.Stdout.WriteString("could not accept argument at ")
			os.Stdout.WriteString(strconv.Itoa(i))
			os.Stdout.WriteString(" failed to parse as byte")
			os.Stdout.WriteString("\n")
		} else {
			interpreter.tape[i] = byte(as_int)
		}
	}

	interpreter.run(false)
}
