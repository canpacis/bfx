package main

import (
	"syscall"
)

type interpreter struct {
	tape    [30000]byte
	pointer int
	ast     ast
	fd      int
}

func (i *interpreter) run(add_new_line bool) {
	i.run_context(i.ast.body)

	if add_new_line {
		syscall.Write(i.fd, []byte{10})
	}
}

func (i *interpreter) run_context(body []op) {
	for _, op := range body {
		switch op := op.(type) {
		case mutate_op:
			if op.mutation == increment {
				i.tape[i.pointer]++
			} else {
				i.tape[i.pointer]--
			}
		case move_op:
			if op.dir == move_right {
				i.pointer++
			} else {
				i.pointer--
			}

			if i.pointer < 0 {
				i.pointer = 0
			}
			if i.pointer > 30000 {
				i.pointer = 30000
			}
		case loop_op:
			for i.tape[i.pointer] != 0 {
				i.run_context(op.body)
			}
		case io_op:
			if op.io == write {
				syscall.Write(i.fd, []byte{i.tape[i.pointer]})
			} else {
				b := [1]byte{}
				syscall.Read(i.fd, b[:])
				i.tape[i.pointer] = b[0]
			}
		case target_op:
			i.fd = int(i.tape[i.pointer])
		}
	}
}

func new_interpreter(program ast) interpreter {
	return interpreter{
		tape:    [30000]byte{},
		pointer: 0,
		ast:     program,
		fd:      1,
	}
}
