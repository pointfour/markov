package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type stateMatrix struct {
	state map[[2]byte]map[byte]float64
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	input := loadText("../input.txt")
	model := stateMatrix{state: make(map[[2]byte]map[byte]float64)}
	train(model.state, input)
	fmt.Println(model.generate())
}

func (m stateMatrix) generate() string {
	seed := [2]byte{84, 104} //{69, 109}
	ret := []byte{seed[0], seed[1]}

	for i := 0; i < 1000; i++ {
		getTo := rand.Float64()
		var at float64
		for key, value := range m.state[seed] {
			at += value
			if at > getTo {
				ret = append(ret, key)
				seed = [2]byte{seed[1], key}
				break
			}
		}
	}
	return string(ret)
}

func train(state map[[2]byte]map[byte]float64, input []byte) {
	norm := make(map[[2]byte]float64)
	for i := 0; i < len(input)-2; i++ {
		key := [2]byte{input[i], input[i+1]}
		val := input[i+2]
		if state[key] == nil {
			state[key] = make(map[byte]float64)
		}
		if state[key][val] == 0 {
			state[key][val] = 1
			norm[key] = 1
		} else {
			state[key][val]++
			norm[key]++
		}
		// fmt.Println(state[key][val])
	}
	for key := range state {
		for skey := range state[key] {
			state[key][skey] /= norm[key]
		}
	}

}

func loadText(path string) []byte {
	ret, _ := ioutil.ReadFile(path)
	return ret
}
