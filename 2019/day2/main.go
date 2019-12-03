// Copyright (c) 2019, Theodor-Iulian Ciobanu
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:

// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

// Day 2

// Part 1

// On the way to your gravity assist around the Moon, your ship computer beeps
// angrily about a "1202 program alarm". On the radio, an Elf is already
// explaining how to handle the situation: "Don't worry, that's perfectly
// norma--" The ship computer bursts into flames.

// You notify the Elves that the computer's magic smoke seems to have escaped.
// "That computer ran Intcode programs like the gravity assist program it was
// working on; surely there are enough spare parts up there to build a new
// Intcode computer!"

// An Intcode program is a list of integers separated by commas
// (like 1,0,0,3,99). To run one, start by looking at the first integer (called
// position 0). Here, you will find an opcode - either 1, 2, or 99. The opcode
// indicates what to do; for example, 99 means that the program is finished and
// should immediately halt. Encountering an unknown opcode means something went
// wrong.

// Opcode 1 adds together numbers read from two positions and stores the result
// in a third position. The three integers immediately after the opcode tell you
// these three positions - the first two indicate the positions from which you
// should read the input values, and the third indicates the position at which
// the output should be stored.

// Part 2

// "Good, the new computer seems to be working correctly! Keep it nearby during
// this mission - you'll probably use it again. Real Intcode computers support
// many more features than your new one, but we'll let you know what they are as
// you need them."

// "However, your current priority should be to complete your gravity assist
// around the Moon. For this mission to succeed, we should settle on some
// terminology for the parts you've already built."

// Intcode programs are given as a list of integers; these values are used as
// the initial state for the computer's memory. When you run an Intcode program,
// make sure to start by initializing memory to the program's values. A position
// in memory is called an address (for example, the first value in memory is at
// "address 0").

// Opcodes (like 1, 2, or 99) mark the beginning of an instruction. The values
// used immediately after an opcode, if any, are called the instruction's
// parameters. For example, in the instruction 1,2,3,4, 1 is the opcode; 2, 3,
// and 4 are the parameters. The instruction 99 contains only an opcode and has
// no parameters.

// The address of the current instruction is called the instruction pointer; it
// starts at 0. After an instruction finishes, the instruction pointer increases
// by the number of values in the instruction; until you add more instructions
// to the computer, this is always 4 (1 opcode + 3 parameters) for the add and
// multiply instructions. (The halt instruction would increase the instruction
// pointer by 1, but it halts the program instead.)

// "With terminology out of the way, we're ready to proceed. To complete the
// gravity assist, you need to determine what pair of inputs produces the output
// 19690720."

// The inputs should still be provided to the program by replacing the values at
// addresses 1 and 2, just like before. In this program, the value placed in
// address 1 is called the noun, and the value placed in address 2 is called the
// verb. Each of the two input values will be between 0 and 99, inclusive.

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var startProfiling, stopProfiling func()

func main() {
	if startProfiling != nil {
		startProfiling()
	}

	var input string
	fmt.Scanf("%s", &input)
	progStr := strings.Split(input, ",")
	origProg := make([]int, len(progStr))
	for i, val := range progStr {
		var err error
		origProg[i], err = strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
	}

	// bootstrap gravity assist
	prog := append(origProg[:0:0], origProg...)
	prog[1], prog[2] = 12, 2
	runProg(prog)
	fmt.Println(prog[0])

LOOP:
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			prog = append(origProg[:0:0], origProg...)
			prog[1], prog[2] = i, j
			runProg(prog)
			if prog[0] == 19690720 {
				fmt.Println(prog[1]*100 + prog[2])
				break LOOP
			}
		}
	}

	if stopProfiling != nil {
		stopProfiling()
	}
}

func runProg(prog []int) {
	ip := 0
	for {
		switch prog[ip] {
		case 1:
			src1 := prog[ip+1]
			src2 := prog[ip+2]
			dest := prog[ip+3]
			prog[dest] = prog[src1] + prog[src2]
		case 2:
			src1 := prog[ip+1]
			src2 := prog[ip+2]
			dest := prog[ip+3]
			prog[dest] = prog[src1] * prog[src2]
		case 99:
			return
		default:
			log.Fatalf("unkown opcode on pos %d: %d", ip, prog[ip])
		}
		ip += 4
	}
}
