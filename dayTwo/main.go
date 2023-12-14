package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func getInputData() ([]string, error) {
  filePath := "./inputData.txt"
  content, err := os.ReadFile(filePath)
  if err != nil {
    return nil,err
  }
  inputString := string(content)

  // testData

//   inputString = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

  splitInputString := strings.Split(inputString, "\n")
  return splitInputString, nil
}

func getMinNumberOfCubes(set string, minCubes map[string]int) {
  cubesFromSet := strings.Split(set, ",")
  for _, cubeFromSet := range cubesFromSet{
    cubeInfoArray := strings.Fields(cubeFromSet)
    quantity, err := strconv.Atoi(cubeInfoArray[0])
    if err != nil {
      return
    }
    color := cubeInfoArray[1]
    if minCubes[color] < quantity{
    minCubes[color] = quantity
    }
  }
}


func checkIfOkay(color string, quantity int, minCubes map[string]int) int {
  inputCubes := map[string]int {
    "red": 12,
    "green": 13,
    "blue": 14,
  }
  
  if quantity > inputCubes[color] {
    return -1
  }

  return 0


}

func checkCubesInSet(set string, minCubes map[string]int) bool {
  cubesFromSet := strings.Split(set, ",")
  for _, cubeFromSet := range cubesFromSet{
    cubeInfoArray := strings.Fields(cubeFromSet)
    quantity, err := strconv.Atoi(cubeInfoArray[0])
    if err != nil {
      return false
    }
    color := cubeInfoArray[1]
    answer := checkIfOkay(color, quantity, minCubes)
    if answer != 0 {
      return false
    }
  }
  return true

}

func readSet(set string, setIsOk bool, minCubes map[string]int) bool{
answer := checkCubesInSet(set, minCubes)
    if !answer {
    setIsOk = false
    }
  return setIsOk
}

func readGame(inputString string, sumOfIds int) (int, int, error) {
  minCubes := map[string]int {
    "red": 0,
    "green": 0,
    "blue": 0,
  }

  newString := strings.Split(inputString, ":")
  game := newString[0]
  gameFields := strings.Fields(game)
  if len(gameFields) > 1{
  gameId := gameFields[1]
  gameIdInt, err := strconv.Atoi(gameId)
  if err != nil{
    return -1, 0, err
  }
  sets := strings.Split(newString[1], ";")
  isSetOk := true
  for _, set := range sets {
    getMinNumberOfCubes(set,minCubes)
    isSetOk = readSet(set, isSetOk, minCubes)
  }
    minCubePower := 1
    for _, value := range minCubes {
      minCubePower *= value
    }
  if isSetOk {
      return gameIdInt, minCubePower, nil
  }else{
      return 0, minCubePower, nil
    }
  }
  return -1, 0, nil
}


func main() {
  sumOfIds := 0
  sumOfPowers := 0
  inputStrings, err := getInputData()
  if err != nil {
    fmt.Println("There was an error with getInputData: ", err)
  }
  for _, inputString := range inputStrings {

    sumOfIdsFromInputString, minCubePower, err := readGame(inputString, sumOfIds)
    if err != nil {
      return
    }
    if sumOfIdsFromInputString == -1 {
      continue
    }
    sumOfIds += sumOfIdsFromInputString
    sumOfPowers += minCubePower
  }
  fmt.Println("Sum of ids: ",sumOfIds)
  fmt.Println("Sum of powers: ", sumOfPowers)
}
