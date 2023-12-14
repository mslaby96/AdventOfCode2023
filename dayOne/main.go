package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var spelledOutToNumeric = map[string]string{
    "one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
}

func findAllDigits(input []string) {
	p2 := 0
  for _, line := range input {
		p2Digits := []int{}
    for i, c := range line {

      if '0' <= c && c <= '9' {
				digit, _ := strconv.Atoi(string(c))
				p2Digits = append(p2Digits, digit)
			}

      for d, val := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
				if strings.HasPrefix(line[i:], val) {
					p2Digits = append(p2Digits, d+1)
				}
			}

    }

		if len(p2Digits) > 0 {
      v, err := strconv.Atoi(strconv.Itoa(p2Digits[0]) + strconv.Itoa(p2Digits[len(p2Digits)-1]))
      if err != nil {
        fmt.Println("Error")
      }
			p2 += v
		}
  }
	fmt.Println(p2)
}


func findNumbers(input []string) ([]int, error) {
  var matches []int
  for _, line := range input {
    re := regexp.MustCompile(`\d`)
    numberArray := re.FindAllString(line, -1)
    sum := ""
    if len(numberArray) > 0{
      sum = numberArray[0] + numberArray[len(numberArray)-1]
    }else {
      continue
    }
    intSum, err := strconv.Atoi(sum)
    if err != nil {
      fmt.Println("error inside findNumbers")
      return nil ,err;
    } else {
      matches = append(matches, intSum)
    }
  }
  return matches, nil
}

func addNumbers(input []int) int {
  answer := 0
  for _, number := range input {
    answer += number
  }
  return answer
}


func main() {
  var userInput int
  var env string
  fmt.Println("Define which environment you want to use 1. InputData, 2. Test :")
  _, err := fmt.Scan(&userInput)
  if err != nil {
    fmt.Println("There was problem with user input: ", err)
  }
  switch userInput {
  case 1:
    env = "InputData"
  case 2:
    env = "test"
  }
  fmt.Println(env)
  inputString := ""
  if env == "test"{
  inputString = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
  } else {
  filePath := "./inputData.txt"
  content, err := os.ReadFile(filePath)
  if err != nil {
    fmt.Println("Error reading a file:", err)
    return
  }
    inputString = string(content)
  }
  splitInputString := strings.Split(inputString, "\n")
  numbers, err := findNumbers(splitInputString)
  if err != nil {
    fmt.Println("There was an error: ", err)
  } else {
    answer := addNumbers(numbers)
    fmt.Println("Answer is ", answer)
  }

  findAllDigits(splitInputString)
}
