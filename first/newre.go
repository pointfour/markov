package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type stateMatrix struct {
	Matrix [27][27]float64
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	rawInput, _ := ioutil.ReadFile("../names.txt")
	stringRawInput := string(rawInput)
	sanatizedInput := sanatize(&stringRawInput)
	input := []byte(*sanatizedInput)
	matrix := stateMatrix{Matrix: [27][27]float64{}}
	train(input, &matrix)
	generated := matrix.generate(0, true)
	fmt.Println(*generated)
}

func train(str []byte, matrix *stateMatrix) (bool, error) {
	for i := 0; i < len(str)-1; i++ {
		first := numValue(str[i])
		second := numValue(str[i+1])
		// fmt.Println(string(str[i]) + "," + string(str[i+1]) + "," + strconv.FormatInt(int64(first), 10) + "," + strconv.FormatInt(int64(second), 10))
		matrix.Matrix[first][second]++ //= matrix.Matrix[first][second] + 1
	}
	for i := 0; i < len(matrix.Matrix); i++ {
		var sum float64
		for s := 0; s < len(matrix.Matrix[i]); s++ {
			sum += matrix.Matrix[i][s]
		}
		for s := 0; s < len(matrix.Matrix[i]); s++ {
			matrix.Matrix[i][s] /= sum
		}
	}
	return true, nil
}

func (matrix stateMatrix) generate(starter byte, oneWord bool) *string {
	generate := []byte{deValuerize(starter)}
	seed := starter
	for i := 0; i < 1000; i++ {
		getTo := rand.Float64()
		var at float64
		for j := byte(0); j < byte(len(matrix.Matrix[seed])); j++ {
			at += matrix.Matrix[seed][j]
			if at > getTo {
				// fmt.Printf("(%f", at)
				// fmt.Printf(",%f)\n", getTo)
				seed = j
				if seed == 26 && oneWord {
					ret := string(generate)
					return &ret
				}
				// fmt.Println(seed)
				generate = append(generate, deValuerize(seed))
				break
			}
		}
	}
	ret := string(generate)
	return &ret
}

func sanatize(str *string) *string {
	reg, err := regexp.Compile("[^a-z ]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(strings.ToLower(*str), "")
	return &processedString

}

func numValue(letter byte) byte {

	if letter == 32 {
		return 26
	} else if letter >= 97 && letter <= 122 {
		return letter - 97
	}
	return 100
}

func deValuerize(letter byte) byte {
	if letter == 26 {
		return 32
	}
	return letter + 97
}
