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

// Day 1

// Part 1

// At the first Go / No Go poll, every Elf is Go until the Fuel Counter-Upper.
// They haven't determined the amount of fuel required yet.

// Fuel required to launch a given module is based on its mass. Specifically, to
// find the fuel required for a module, take its mass, divide by three, round
// down, and subtract 2.

// Part 2

// During the second Go / No Go poll, the Elf in charge of the Rocket Equation
// Double-Checker stops the launch sequence. Apparently, you forgot to include
// additional fuel for the fuel you just added.

// Fuel itself requires fuel just like a module - take its mass, divide by
// three, round down, and subtract 2. However, that fuel also requires fuel, and
// that fuel requires fuel, and so on. Any mass that would require negative fuel
// should instead be treated as if it requires zero fuel; the remaining mass, if
// any, is instead handled by wishing really hard, which has no mass and is
// outside the scope of this calculation.

package main

import (
	"fmt"
	"io"
	"log"
	"math"
)

var startProfiling, stopProfiling func()

func main() {
	if startProfiling != nil {
		startProfiling()
	}

	var total1, total2 int64
	for {
		var input int64
		_, err := fmt.Scanf("%d\n", &input)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		total1 += int64(math.Floor(float64(input)/3)) - 2
		for {
			input = int64(math.Floor(float64(input)/3)) - 2
			if input <= 0 {
				break
			}

			total2 += input
		}
	}
	fmt.Printf("%d\n%d\n", total1, total2)

	if stopProfiling != nil {
		stopProfiling()
	}
}
