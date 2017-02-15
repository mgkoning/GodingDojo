package main

import "strconv"

func main() {

}

func parseAccount(line1, line2, line3 string) (string, string) {
	result, ok := "", true

	for i := 0; i < 9; i++ {
		from, to := 3*i, 3*(i+1)
		digitLine1 := line1[from:to]
		digitLine2 := line2[from:to]
		digitLine3 := line3[from:to]
		digit, digitOk := parseDigit(digitLine1, digitLine2, digitLine3)
		result += digit
		ok = ok && digitOk
	}
	status := ""
	switch {
	case !ok:
		status = "ILL"
	case !validChecksum(result):
		status = "ERR"
	}
	return result, status
}

var digitMap = map[string]string{
	" _ | ||_|": "0",
	"     |  |": "1",
	" _  _||_ ": "2",
	" _  _| _|": "3",
	"   |_|  |": "4",
	" _ |_  _|": "5",
	" _ |_ |_|": "6",
	" _   |  |": "7",
	" _ |_||_|": "8",
	" _ |_| _|": "9",
}

func parseDigit(line1, line2, line3 string) (string, bool) {
	digit, ok := digitMap[line1+line2+line3]
	if !ok {
		return "?", false
	}
	return digit, true
}

func validChecksum(number string) bool {
	result, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return false
	}
	weight, sum := int64(1), int64(0)
	for result != 0 {
		leastSignificantDigit := result % 10
		sum += leastSignificantDigit * weight
		weight++
		result /= 10
	}
	return sum%11 == 0
}
