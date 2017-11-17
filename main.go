package gocodicefiscale

import (
	"strings"
	"regexp"
)

func ValidCF(input string) bool {
	PATTERN := `(?:[B-DF-HJ-NP-TV-Z](?:[AEIOU]{2}|[AEIOU]X)|[AEIOU]{2}[AEIOUX]|[B-DF-HJ-NP-TV-Z]{2}[A-Z]){2}[\dLMNP-V]{2}(?:[A-EHLMPR-T](?:[04LQ][1-9MNP-V]|[1256LMRS][\dLMNP-V])|[DHPS][37PT][0L]|[ACELMRT][37PT][01LM])(?:[A-MZ][1-9MNP-V][\dLMNP-V]{2}|[A-M][0L](?:[\dLMNP-V][1-9MNP-V]|[1-9MNP-V][0L]))[A-Z]`

	input = strings.TrimSpace(input)
	if len(input) != 16 {
		return false
	}

	match, err := regexp.MatchString(PATTERN, input)
	if err != nil || !match {
		return false
	}

	// checksum
	cfInput := input[:15]
	controlCode := input[15]

	checksum := computeChecksum(cfInput)
	if checksum != controlCode {
		return false
	}

	return true
}

func computeChecksum(cf string) uint8 {
	digits := "0123456789"
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	evenControlcode := make(map[int32]int)
	for idx, char := range digits {
		evenControlcode[char] = idx
	}
	for idx, char := range chars {
		evenControlcode[char] = idx
	}

	values := []int{1, 0, 5, 7, 9, 13, 15, 17, 19, 21, 2, 4, 18, 20, 11, 3, 6, 8, 12, 14, 16, 10, 22, 25, 24, 23}

	odd_controlcode := make(map[int32]int)
	for idx, char := range digits {
		odd_controlcode[char] = values[idx]
	}
	for idx, char := range chars {
		odd_controlcode[char] = values[idx]
	}

	code := 0
	for idx, char := range cf {
		if idx%2 == 0 {
			code += odd_controlcode[char]
		} else {
			code += evenControlcode[char]
		}
	}

	return chars[code%26]
}
