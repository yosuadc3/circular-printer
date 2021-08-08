package main

import (
  "fmt"
)

var alphabetCircle = []string{}

func main(){
  for i := 'A'; i <= 'Z'; i++ {
    alphabetCircle = append(alphabetCircle, string(i))
  }
  stringToCount := "BZA"
  fmt.Println(getTime(stringToCount))
}

func indexOf(element string, data []string) (int) {
   for index, value := range data {
       if element == value {
         return index
       }
   }
   return -1    //not found.
}

func abs(x int) int {
	if x < 0 {
    return -x
	}
	return x
}

func getTime(words string) int {
  var ( 
    pointer = 0
    steps = 0
    clockwise = 0
    counterClockwise = 0
  )
  var state = ""
  for _, word := range words {
    if (string(word) == "A"){
      counterClockwise = abs(pointer - 1)
      clockwise = abs(25 - (pointer - 1))
    } else {
      clockwise = abs(indexOf(string(word), alphabetCircle) - pointer)
      counterClockwise = abs(26 - ((indexOf(string(word), alphabetCircle) - pointer)))
    }
    if clockwise < counterClockwise {
      steps += clockwise
      state = "clockwise"
    } else {
      steps += counterClockwise
      state = "counter clockwise"
    }
    fmt.Println(alphabetCircle[pointer],"to",string(word),"takes", steps , state)
    pointer = indexOf(string(word), alphabetCircle)
  }
  return steps
}
