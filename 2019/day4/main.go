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

// Day 4

// Part 1

// You arrive at the Venus fuel depot only to discover it's protected by a
// password. The Elves had written the password on a sticky note, but someone
// threw it out.

// However, they do remember a few key facts about the password:
// - It is a six-digit number.
// - The value is within the range given in your puzzle input.
// - Two adjacent digits are the same (like 22 in 122345).
// - Going from left to right, the digits never decrease; they only ever
//   increase or stay the same (like 111123 or 135679).

//Other than the range rule, the following are true:
// - 111111 meets these criteria (double 11, never decreases).
// - 223450 does not meet these criteria (decreasing pair of digits 50).
// - 123789 does not meet these criteria (no double).

// How many different passwords within the range given in your puzzle input meet
// these criteria?

// Part 2

// An Elf just remembered one more important detail: the two adjacent matching
// digits are not part of a larger group of matching digits.

// Given this additional criterion, but still ignoring the range rule, the
// following are now true:

// - 112233 meets these criteria because the digits never decrease and all
//   repeated digits are exactly two digits long.
// - 123444 no longer meets the criteria (the repeated 44 is part of a larger
//   group of 444).
// - 111122 meets the criteria (even though 1 is repeated more than twice, it
//   still contains a double 22).

// How many different passwords within the range given in your puzzle input meet
// all of the criteria?

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
	inputRange := strings.Split(input, "-")
	start, err := strconv.ParseUint(inputRange[0], 10, 64)
	if err != nil {
		log.Fatalf("invalid start range %s: %v", inputRange[0], err)
	}
	end, err := strconv.ParseUint(inputRange[1], 10, 64)
	if err != nil {
		log.Fatalf("invalid end range %s: %v", inputRange[1], err)
	}
	if start > end {
		log.Fatalf("invalid range: %d>%d", start, end)
	}

	fmt.Println(countValidPINs(start, end, -1))
	fmt.Println(countValidPINs(start, end, 2))

	if stopProfiling != nil {
		stopProfiling()
	}
}

func countValidPINs(start, end uint64, maxrep int) uint64 {
	var count uint64
	for i := start; i <= end; i++ {
		if isValidPIN(i, maxrep) {
			count++
		}
	}

	return count
}

func getNextValidPIN(n uint64, maxrep int) uint64 {
	// TODO
	return n
}

func isValidPIN1(n uint64, _ int) bool {
	doubleDigits := false
	i := n % 10
	n = n / 10
	j := n % 10
	for n > 0 {
		switch {
		case i == j:
			doubleDigits = true
		case i < j:
			return false
		}

		i = j
		n = n / 10
		j = n % 10
	}

	return doubleDigits
}

func isValidPIN(n uint64, maxrep int) bool {
	i := n % 10
	n = n / 10
	j := n % 10
	reps := 1
	repOK := false
	for n > 0 {
		switch {
		case i == j:
			reps++
		case i < j:
			return false
		default:
			if reps > 1 {
				repOK = repOK ||
					(maxrep == -1) ||
					(reps == maxrep)
				reps = 1
			}
		}

		i = j
		n = n / 10
		j = n % 10
	}
	if reps > 1 {
		repOK = repOK ||
			(maxrep == -1) ||
			(reps == maxrep)
	}

	return repOK
}
