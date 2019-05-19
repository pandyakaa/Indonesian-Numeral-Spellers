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
	if s[0] == "negatif" {
		i = 1
	}
	var sum = 0
	var bound = len(s)

	for i < bound {
		var temp2 = 0
		for s[i] != "ribu" && s[i] != "juta" && s[i] != "milyar" {
			var temp = 0
			temp = temp + toIntUnary(s[i])
			i = i + 1
			if i < bound {
				if toIntUnit(s[i]) != -1 && s[i] != "ribu" && s[i] != "juta" && s[i] != "milyar" {
					temp = temp * toIntUnit(s[i])
					i = i + 1
				} else if s[i] == "belas" {
					temp = temp + 10
					i = i + 1
				}
			}
			temp2 = temp2 + temp
			if i >= bound {
				break
			}
		}
		if i < bound {
			temp2 = temp2 * toIntUnit(s[i])
		}
		sum = sum + temp2
		i = i + 1
	}
	if s[0] == "negatif" {
		return sum * -1
	}
	return sum
}

func isValidText(str string) bool {
	var s = strings.Split(str, " ")

	if len(s) == 1 {
		if toIntUnary(s[0]) == -1 {
			return false
		}
	} else {

		var i = 0
		for i < len(s) {
			if s[i] == "satu" {
				if i+1 < len(s) {
					if s[i+1] == "belas" || s[i+1] == "puluh" || s[i+1] == "ratus" || s[i+1] == "ribu" {
						return false
					}
				}
			}
			i = i + 1
		}

		i = 0
		for i < len(s) {
			if s[i] == "puluh" {
				if i-1 >= 0 {
					if toIntUnary(s[i-1]) > 9 || toIntUnary(s[i-1]) == -1 {
						return false
					}
				}

				if i+1 < len(s) {
					if s[i+1] == "belas" || s[i+1] == "puluh" || s[i+1] == "ratus" || toIntUnary(s[i+1]) > 9 {
						return false
					}
				}
			}
			i = i + 1
		}

		i = 0
		for i < len(s) {
			if s[i] == "belas" {

				if i-1 >= 0 {
					if toIntUnary(s[i-1]) > 9 || toIntUnary(s[i-1]) == -1 {
						return false
					}
				}

				if i+1 < len(s) {
					if s[i+1] == "belas" || s[i+1] == "puluh" || s[i+1] == "ratus" || toIntUnary(s[i+1]) != -1 {
						return false
					}
				}
			}
			i = i + 1
		}

		i = 0
		for i < len(s) {
			if s[i] == "ratus" {

				if i-1 >= 0 {
					if toIntUnary(s[i-1]) > 9 || toIntUnary(s[i-1]) == -1 {
						return false
					}
				}

				if i+1 < len(s) {
					if s[i+1] == "belas" || s[i+1] == "puluh" || s[i+1] == "ratus" {
						return false
					}
				}
			}
			i = i + 1
		}

		i = 0
		for i < len(s) {
			if i+1 < len(s) {
				if toIntUnary(s[i]) != -1 && toIntUnary(s[i+1]) != -1 && toIntUnary(s[i]) < 100 {
					return false
				}
			}
			i = i + 1
		}

	}
	return true
}
