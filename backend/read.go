package main

import (
	"strings"
)

func toIntUnary(str string) int {
	switch str {
	case "satu":
		return 1
	case "dua":
		return 2
	case "tiga":
		return 3
	case "empat":
		return 4
	case "lima":
		return 5
	case "enam":
		return 6
	case "tujuh":
		return 7
	case "delapan":
		return 8
	case "sembilan":
		return 9
	case "sepuluh":
		return 10
	case "sebelas":
		return 11
	case "seratus":
		return 100
	case "seribu":
		return 1000
	default:
		return -1
	}
}

func toIntUnit(str string) int {
	switch str {
	case "puluh":
		return 10
	case "ratus":
		return 100
	case "ribu":
		return 1000
	case "juta":
		return 1000000
	case "milyar":
		return 1000000000
	default:
		return -1
	}
}

func readUtils(str string) int {
	var s = strings.Split(str, " ")
	var i = 0
	var sum = 0
	var bound = len(s)

	for i < bound {
		for s[i] != "ribu" && s[i] != "juta" && s[i] != "milyar" {
			var temp = 0
			temp = temp + toIntUnary(s[i])
			i = i + 1
			if i < bound {
				if toIntUnit(s[i]) != -1 {
					temp = temp * toIntUnit(s[i])
					i = i + 1
				} else if s[i] == "belas" {
					temp = temp + 10
					i = i + 1
				}
			}
			sum = sum + temp
			if i >= bound {
				break
			}
		}
		if i < bound {
			sum = sum * toIntUnit(s[i])
			i = i + 1
		}
	}
	return sum
}
