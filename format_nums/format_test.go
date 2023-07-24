// Save this file as main_test.go and run "go test -bench ."

package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/dustin/go-humanize"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func formatIntByDividing(number int) string {
	var output string
	negative := false
	if number < 0 {
		negative = true
		number = -number
	}
	for {
		numberPart := number % 1000
		outputGroup := strconv.Itoa(numberPart)
		number /= 1000
		if number == 0 {
			output = outputGroup + output
			break
		}
		if numberPart >= 100 {
			output = "," + outputGroup + output
		} else if numberPart >= 10 {
			output = ",0" + outputGroup + output
		} else {
			output = ",00" + outputGroup + output
		}
	}
	if negative {
		output = "-" + output
	}
	return output
}

func formatIntByInserting(number int) string {
	output := strconv.Itoa(number)
	startOffset := 3
	if number < 0 {
		startOffset++
	}
	for outputIndex := len(output); outputIndex > startOffset; {
		outputIndex -= 3
		output = output[:outputIndex] + "," + output[outputIndex:]
	}
	return output
}

func formatIntByCollecting(number int) string {
	input := strconv.Itoa(number)
	inputLength := len(input)
	numberOfDigits := inputLength
	if number < 0 {
		numberOfDigits--
	}
	numberOfCommas := (numberOfDigits - 1) / 3
	if numberOfCommas == 0 {
		return input
	}
	outputLength := inputLength + numberOfCommas
	output := make([]byte, outputLength)
	if number < 0 {
		input, output[0] = input[1:], '-'
		inputLength--
	}
	for inputIndex, outputIndex, indexInGroup := inputLength-1, outputLength-1, 0; ; {
		output[outputIndex] = input[inputIndex]
		if inputIndex == 0 {
			return string(output)
		}
		if indexInGroup++; indexInGroup == 3 {
			outputIndex--
			indexInGroup = 0
			output[outputIndex] = ','
		}
		inputIndex--
		outputIndex--
	}
}

func formatIntByCopying(number int) string {
	input := strconv.Itoa(number)
	inputLength, startOffset := len(input), 0
	if number < 0 {
		startOffset = 1
	}
	numberOfCommas := (inputLength - startOffset - 1) / 3
	if numberOfCommas == 0 {
		return input
	}
	outputLength := inputLength + numberOfCommas
	buffer := make([]byte, outputLength)
	startOffset += 3
	inputIndex, outputIndex := inputLength, outputLength
	for inputIndex > startOffset {
		inputIndex -= 3
		outputIndex -= 3
		copy(buffer[outputIndex:outputIndex+3], input[inputIndex:])
		outputIndex--
		buffer[outputIndex] = ','
	}
	if outputIndex > 0 {
		copy(buffer[:outputIndex], input)
	}
	return string(buffer)
}

func formatIntWithBuffer(number int) string {
	input := strconv.Itoa(number)
	startOffset := 0
	var buffer bytes.Buffer
	if number < 0 {
		startOffset = 1
		buffer.WriteByte('-')
	}
	inputLength := len(input)
	numberOfCommas := (inputLength - startOffset - 1) / 3
	if numberOfCommas == 0 {
		return input
	}
	commaIndex := 3 - ((inputLength - startOffset) % 3)
	if commaIndex == 3 {
		commaIndex = 0
	}
	for inputIndex := startOffset; inputIndex < inputLength; inputIndex++ {
		if commaIndex == 3 {
			buffer.WriteRune(',')
			commaIndex = 0
		}
		commaIndex++
		buffer.WriteByte(input[inputIndex])
	}
	return buffer.String()
}

var integerGrouping = regexp.MustCompile("(\\d+)(\\d{3})")

func formatIntWithRexExp(number int) string {
	input := strconv.Itoa(number)
	for {
		previousInput := input
		input = integerGrouping.ReplaceAllString(input, "$1,$2")
		if previousInput == input {
			return input
		}
	}
}

func formatIntWithHumanize(number int) string {
	return humanize.Comma(int64(number))
}

var integerPrinter = message.NewPrinter(language.English)

func formatIntWithTextMessage(number int) string {
	return integerPrinter.Sprint(number)
}

func formatIntUsingRecursion(number int) string {
	if number < 0 {
		return "-" + formatIntUsingRecursion(-number)
	}
	if number < 1000 {
		return fmt.Sprintf("%d", number)
	}
	return formatIntUsingRecursion(number/1000) + "," + fmt.Sprintf("%03d", number%1000)
}

func formatIntForBenchmarks(number int) string {
	str := strconv.Itoa(number)
	l_str := len(str)
	digits := l_str
	if number < 0 {
		digits--
	}
	commas := (digits+2)/3 - 1
	l_buf := l_str + commas
	var sbuf [32]byte // pre allocate buffer at stack rather than make([]byte,n)
	buf := sbuf[0:l_buf]
	// copy str from the end
	for s_i, b_i, c3 := l_str-1, l_buf-1, 0; ; {
		buf[b_i] = str[s_i]
		if s_i == 0 {
			return string(buf)
		}
		s_i--
		b_i--
		// insert comma every 3 chars
		c3++
		if c3 == 3 && (s_i > 0 || number > 0) {
			buf[b_i] = ','
			b_i--
			c3 = 0
		}
	}
}

func delimitNumeral(i int, delim rune) string {
	src := strconv.Itoa(i)
	strLen := len(src)
	negative := i < 0
	outStr := ""
	digitCount := 0
	for i := strLen - 1; i >= 0; i-- {

		outStr = src[i:i+1] + outStr
		if digitCount == 2 && ((i > 0 && !negative) || (i > 1 && negative)) {
			outStr = string(delim) + outStr
			digitCount = 0
		} else {
			if src[i:i+1] != "-" {
				digitCount++
			}
		}
	}

	return outStr
}

func formatRune(number int) string {
	return delimitNumeral(number, ',')
}

var testCases = []struct {
	number int
	result string
}{
	{1, "1"},
	{10, "10"},
	{100, "100"},
	{1000, "1,000"},
	{10000, "10,000"},
	{100000, "100,000"},
	{1000000, "1,000,000"},
	{10000000, "10,000,000"},
	{100000000, "100,000,000"},
	{1000000000, "1,000,000,000"},
	{10000000000, "10,000,000,000"},
	{100000000000, "100,000,000,000"},
	{1000000000000, "1,000,000,000,000"},
	{-1, "-1"},
	{-10, "-10"},
	{-100, "-100"},
	{-1000, "-1,000"},
	{-10000, "-10,000"},
	{-100000, "-100,000"},
	{-1000000, "-1,000,000"},
	{-10000000, "-10,000,000"},
	{-100000000, "-100,000,000"},
	{-1000000000, "-1,000,000,000"},
	{-10000000000, "-10,000,000,000"},
	{-100000000000, "-100,000,000,000"},
	{-1000000000000, "-1,000,000,000,000"},
}

func testFormatInt(t *testing.T, formatInt func(int) string, number int, expected string) {
	actual := formatInt(number)
	if actual != expected {
		t.Errorf("Expected %s, but got %s for %d.", expected, actual, number)
	}
}

func makeTest(formatInt func(int) string) func(t *testing.T) {
	return func(t *testing.T) {
		for _, testCase := range testCases {
			actual := formatInt(testCase.number)
			if actual != testCase.result {
				t.Errorf("Expected %s, but got %s for %d.", testCase.result, actual, testCase.number)
			}
		}
	}
}

var tests = []struct {
	name     string
	function func(t *testing.T)
}{
	{"by dividing", makeTest(formatIntByDividing)},
	{"by inserting", makeTest(formatIntByInserting)},
	{"by collecting", makeTest(formatIntByCollecting)},
	{"by copying", makeTest(formatIntByCopying)},
	{"with buffer", makeTest(formatIntWithBuffer)},
	{"with regexp", makeTest(formatIntWithRexExp)},
	{"with humanize", makeTest(formatIntWithHumanize)},
	{"with text/message", makeTest(formatIntWithTextMessage)},
	{"using recursion", makeTest(formatIntUsingRecursion)},
	{"for benchmarks", makeTest(formatIntForBenchmarks)},
	{"format rune", makeTest(formatRune)},
}

func Test(t *testing.T) {
	for _, test := range tests {
		function := test.function
		t.Run(test.name, func(t *testing.T) {
			function(t)
		})
	}
}

func makeBenchmark(formatInt func(int) string) func() {
	return func() {
		formatInt(1)
		formatInt(10)
		formatInt(100)
		formatInt(1000)
		formatInt(10000)
		formatInt(100000)
		formatInt(1000000)
		formatInt(10000000)
		formatInt(100000000)
		formatInt(1000000000)
		formatInt(10000000000)
		formatInt(100000000000)
		formatInt(1000000000000)
		formatInt(-1)
		formatInt(-10)
		formatInt(-100)
		formatInt(-1000)
		formatInt(-10000)
		formatInt(-100000)
		formatInt(-1000000)
		formatInt(-10000000)
		formatInt(-100000000)
		formatInt(-1000000000)
		formatInt(-10000000000)
		formatInt(-100000000000)
		formatInt(-1000000000000)
	}
}

var benchmarks = []struct {
	name     string
	function func()
}{
	{"by dividing", makeBenchmark(formatIntByDividing)},
	{"by inserting", makeBenchmark(formatIntByInserting)},
	{"by collecting", makeBenchmark(formatIntByCollecting)},
	{"by copying", makeBenchmark(formatIntByCopying)},
	{"with buffer", makeBenchmark(formatIntWithBuffer)},
	{"with regexp", makeBenchmark(formatIntWithRexExp)},
	{"with humanize", makeBenchmark(formatIntWithHumanize)},
	{"with text/message", makeBenchmark(formatIntWithTextMessage)},
	{"using recursion", makeBenchmark(formatIntUsingRecursion)},
	{"for benchmarks", makeBenchmark(formatIntForBenchmarks)},
	{"format rune", makeBenchmark(formatRune)},
}

func Benchmark(b *testing.B) {
	for _, benchmark := range benchmarks {
		function := benchmark.function
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				function()
			}
		})
	}
}
