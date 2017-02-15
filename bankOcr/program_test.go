package main

import (
	"strings"
	"testing"
)

var input0 = []string{
	" _  _  _  _  _  _  _  _  _ ",
	"| || || || || || || || || |",
	"|_||_||_||_||_||_||_||_||_|",
}

var input1 = []string{
	"                           ",
	"  |  |  |  |  |  |  |  |  |",
	"  |  |  |  |  |  |  |  |  |",
}
var input2 = splitLines(
	` _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
|_ |_ |_ |_ |_ |_ |_ |_ |_ `)

var allDigits = splitLines(
	`    _  _     _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|`)

var illDigits = splitLines(
	`    _  _     _  _  _  _  _ 
  | _| _||_| _ |_   ||_||x|
  ||_  _|  | _||_|  ||_| _|`)

func splitLines(input string) []string {
	return strings.Split(input, "\n")
}

func TestParseAccount0(t *testing.T) {
	expectEqual("000000000", "", input0, t)
}

func TestParseAccount1(t *testing.T) {
	expectEqual("111111111", "ERR", input1, t)
}

func TestParseAccount2(t *testing.T) {
	expectEqual("222222222", "ERR", input2, t)
}

func TestParseAccountAllDigits(t *testing.T) {
	expectEqual("123456789", "", allDigits, t)
}

func TestParseAccountIllDigits(t *testing.T) {
	expectEqual("1234?678?", "ILL", illDigits, t)
}

func TestChecksum(t *testing.T) {
	number := "457508000"
	if !validChecksum(number) {
		t.Errorf("Expected valid checksum for %v", number)
	}
	number = "664371495"
	if validChecksum(number) {
		t.Errorf("Expected invalid checksum for %v", number)
	}
}

func expectEqual(expected, expectedStatus string, input []string, t *testing.T) {
	var actual, status = parseAccount(input[0], input[1], input[2])
	if actual != expected {
		t.Errorf("Result should be %v but was %v", expected, actual)
	}
	if status != expectedStatus {
		t.Errorf("Result status should be %v but was %v", expectedStatus, status)
	}
}
