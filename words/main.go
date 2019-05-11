package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const SIZE int = 2

type stateMatrix struct {
	Trained map[[SIZE]string]strAndProb
}

type strAndProb struct {
	Str  string
	Prob float64
}

type trainingMatrix map[[SIZE]string]map[string]float64

var model stateMatrix

func main() {
	// rand.Seed(time.Now().UTC().UnixNano())
	input := loadText("../input/commie.txt")
	// for i := 0; i < len(input); i++ {
	// 	fmt.Println(input[i])
	// }
	train(&model, input)
	fmt.Println(generate(&model, [SIZE]string{"implied", "your"}))
}

func generate(model *stateMatrix, input [SIZE]string) []string {
	ret := []string{}
	key := input
	var probability float64 = 1
	for true {
		if len(ret) > 6 {
			break
		}
		if _, exists := model.Trained[key]; exists {
			probability *= model.Trained[key].Prob
			fmt.Println(model.Trained[key])

			if probability > 0.9 {
				ret = append(ret, model.Trained[key].Str)
				last := model.Trained[key].Str
				for i := 0; i < len(key)-1; i++ {
					key[i] = key[i+1]
				}
				key[len(key)-1] = last
			} else {
				break
			}
		} else {
			break
		}
	}
	return ret
}

func train(model *stateMatrix, input []string) {
	matrix := make(trainingMatrix)
	for i := 0; i < len(input)-SIZE; i++ {
		key := [SIZE]string{}
		for x := 0; x < SIZE; x++ {
			key[x] = input[i+x]
		}
		val := input[i+SIZE]
		if matrix[key] == nil {
			matrix[key] = make(map[string]float64)
		}
		if matrix[key][val] == 0 {
			matrix[key][val] = 1
		} else {
			matrix[key][val]++
		}
	}
	model.Trained = make(map[[SIZE]string]strAndProb)
	for key := range matrix {
		var sum float64
		var biggestSkey string
		var biggestVal float64
		for skey := range matrix[key] {
			sum += matrix[key][skey]
			if matrix[key][skey] > biggestVal {
				biggestVal = matrix[key][skey]
				biggestSkey = skey
			}
		}
		model.Trained[key] = strAndProb{biggestSkey, biggestVal / sum}
		// fmt.Println(key, sum, biggestVal, biggestSkey)
	}
	// fmt.Println(model.Trained)
}

func loadText(path string) []string {
	unProcessed, _ := ioutil.ReadFile(path)
	ret := strings.Fields(string(unProcessed))
	for i := 0; i < len(ret); i++ {
		ret[i] = strings.ToLower(ret[i])
	}
	return ret
}
