package main

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	MathRegexp("input.txt", "output.txt")
}

var re *regexp.Regexp = regexp.MustCompile("(\\d+)\\s*([-+/*])\\s*(\\d+)")

func MathRegexp(input, output string) {
	fs := os.DirFS(".")
	sliceIn := readInFile(input, fs)
	regexpCompile(sliceIn, re, output)

}

func regexpCompile(sliceIn []string, re *regexp.Regexp, outputFile string) {

	fileOutput := outFile(outputFile)
	defer fileOutput.Close()
	writer := bufio.NewWriter(fileOutput)
	for _, s := range sliceIn {
		result := exprHelp(matches(s, re))
		if result == "" {
			continue
		}
		fileOutput.WriteString(strings.Replace(s, "?", result, 1))
	}
	writer.Flush()
}

func outFile(output string) *os.File {
	_ = os.Remove(output)
	fOut, err := os.OpenFile(output, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	return fOut
}

func readInFile(input string, fsys fs.FS) []string {

	f, err := fsys.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sl := []string{}
	for scanner.Scan() {
		sl = append(sl, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sl
}

func matches(input string, re *regexp.Regexp) (int, int, string) {
	matches := re.FindAllStringSubmatch(input, -1)
	if matches == nil {
		return 0, 0, ""
	}
	num1, _ := strconv.Atoi(matches[0][1])
	num2, _ := strconv.Atoi(matches[0][3])
	operation := matches[0][2]

	return num1, num2, operation
}

func exprHelp(num1, num2 int, operation string) string {
	var result int
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		return ""
	}

	return strconv.Itoa(result) + "\n"
}
