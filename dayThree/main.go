package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getInputData() (string, error){
  filePath := "./inputData.txt"
  content, err := os.ReadFile(filePath)
  if err != nil {
    return "",err
  }
  inputString := string(content)
  return inputString, nil
}

func multiplyNumberPartsByCouple (numberStringArray []string) {
  acc := 0
  for i:=0; i<len(numberStringArray); i+=2 {
    numberInt1, err := strconv.Atoi(numberStringArray[i])
    if err != nil {
      return
    }
    numberInt2, err := strconv.Atoi(numberStringArray[i+1])
    if err != nil {
      return
    }
    acc += numberInt1 * numberInt2
  }
  fmt.Println("p2 answer:", acc)
}


func findWholeNumberParts(partNumbersWithGearIndexesMatrix[][]int, stringMatrix [][]string) []string {
  numberCharacters := "1234567890"
  var numberStringArray []string
  for _,partNumberIndex := range partNumbersWithGearIndexesMatrix {
    stringArray := stringMatrix[partNumberIndex[0]]
    if stringMatrix != nil{
      if partNumberIndex[1] != 0 {
        numberString := ""
        for j:=0; j<len(stringArray); j++ {
          if !strings.Contains(numberCharacters,stringArray[j]) && len(numberString) > 0 {
            break
          }else if strings.Contains(numberCharacters,stringArray[j]) && len(numberString) == 0 {
            for k:=j; k<=partNumberIndex[1]; k++ {
              if strings.Contains(numberCharacters, stringArray[k]) {
                numberString = stringArray[j]
              }else if !strings.Contains(numberCharacters, stringArray[k]) {
                numberString = ""
                break
              }
            }
          } else if strings.Contains(numberCharacters,stringArray[j]) && len(numberString) > 0 {
            numberString += stringArray[j]
          }
        }
        numberStringArray = append(numberStringArray, numberString)
      }
    }
  }
  return numberStringArray
}

func findGears(stringMatrix [][]string) [][]int{
  numberCharacters := "1234567890"
  gearCharacter := "*"
  var gearIndexMatrix [][]int
  var partNumbersWithGearIndexesMatrix [][]int
  for i:=0; i<len(stringMatrix); i++ {
    for j:=0; j<len(stringMatrix[i]); j++{
      if stringMatrix[i][j] == gearCharacter {
        adjacentNumbers := 0
        var twoPartNumbersWithGearMatrix [][]int
        for k:=i-1; k<=i+1; k++{
          for m:=j-1; m<=j+1; m++{
            if strings.Contains(numberCharacters, stringMatrix[k][m]) {
              adjacentNumbers += 1
              twoPartNumbersWithGearMatrix = append(twoPartNumbersWithGearMatrix, []int{k,m})
              if adjacentNumbers == 2 {
                index := []int{i,j}
                gearIndexMatrix = append(gearIndexMatrix, index)
                partNumbersWithGearIndexesMatrix = append(partNumbersWithGearIndexesMatrix, twoPartNumbersWithGearMatrix...)
              }
              if k + 1 == i || k - 1 == i {
                if strings.Contains(numberCharacters, stringMatrix[k][m+1]){
                  break
                }
              }
            }
          }
        }
      }
    }
  }
  return partNumbersWithGearIndexesMatrix
}

func checkIfSymbol(stringMatrix [][]string, partNumbersArray []string, indexesToCheck[][]int){
  characters := "!@#$%^&*(){}[]<>+=-_,;:'\"/\\|`~?"
  answer := 0
  for i,indexToCheck := range indexesToCheck {
    partNumber := partNumbersArray[i]
    columnLen := len(partNumber)
    
    rowStart := 0
    rowEnd := 1
    if indexToCheck[0] != 0{
      rowStart = indexToCheck[0] - 1
      rowEnd = indexToCheck[0] + 1
    } 
    if indexToCheck[0] >= (len(stringMatrix) - 2) {
      rowStart = indexToCheck[0] - 1
      rowEnd = len(stringMatrix) - 2
    }

    columnStart := 0
    if indexToCheck[1] - columnLen >= 0 {
      columnStart = indexToCheck[1] - columnLen 
    }
    columnEnd := indexToCheck[1] + 1
    if columnEnd >= len(stringMatrix[rowStart]) {
      columnEnd = len(stringMatrix[rowStart]) - 1
    }

    for k := rowStart; k <= rowEnd; k++{
      for j := columnStart; j <= columnEnd; j++{
        if strings.Contains(characters, stringMatrix[k][j]) {
            partNumberInt, err := strconv.Atoi(partNumbersArray[i])
            if err != nil {
              fmt.Println("error")
              return
            }
            answer += partNumberInt
      } 
    }
  }
}
  fmt.Println("p1 answer:", answer)
}


func makeAMatrixOutOfInput(input string) [][]string {
  rows := strings.Split(input, "\n")
  for i, row := range rows {
    rows[i] = strings.TrimSpace(row)
  }
  var stringMatrix [][]string
  for _, row := range rows {
    rowAsSlice := strings.Split(row, "")
    stringMatrix = append(stringMatrix, []string(rowAsSlice))
  }

  return stringMatrix
}

func makeArrayFromString(input string, inputMatrix [][]string) {
  var partNumbersArray []string
  var indexesToCheck [][]int
  inputArray := strings.Split(input, "\n")
  for index, input := range inputArray {
    if len(input) > 0 {
    charsArray := strings.Split(input, "")
    numberString := ""
    var endIndex []int
      for index2, char := range charsArray {
        
        if !unicode.IsDigit(rune(charsArray[index2][0])){
          continue
        }
        numberString += char
        if index2 + 1 <= len(charsArray) - 1{
          if !unicode.IsDigit(rune(charsArray[index2+1][0])){
            endIndex = []int{index, index2}
            indexesToCheck = append(indexesToCheck, endIndex)
            partNumbersArray = append(partNumbersArray, numberString)
            numberString = ""
          }
        } else if index2 == len(charsArray) - 1{
            endIndex = []int{index, index2}
            indexesToCheck = append(indexesToCheck, endIndex)
            partNumbersArray = append(partNumbersArray, numberString)
            numberString = ""
        }
      }
    }
  }
  checkIfSymbol(inputMatrix, partNumbersArray, indexesToCheck)


}

func main() {
  inputData, err := getInputData()
  if err != nil {
    return
  }
//   inputData = `467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..595.755.
// .....*....
// ...$.*....
// .664.598..`


  inputMatrix := makeAMatrixOutOfInput(inputData)
  makeArrayFromString(inputData, inputMatrix)
  partNumbersWithGearIndexesMatrix := findGears(inputMatrix)
  numberStringArray := findWholeNumberParts(partNumbersWithGearIndexesMatrix, inputMatrix)
  multiplyNumberPartsByCouple(numberStringArray)
}
