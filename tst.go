package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var testNum int

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func doResults(numVillians, numHeros []int, testNum int, finalResult chan<- []string) {
	var rmIndex int
	var rmValue int
	for index1 := 0; index1 < len(numHeros); index1++ {
		for rmIndex := 0; rmIndex < len(numHeros); rmIndex++ {
			if numVillians[rmIndex] < numHeros[index1] {
				if rmValue > numVillians[rmIndex] {
					rmValue = numVillians[rmIndex]
				}
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numRead, err := fmt.Scanf("%d", &testNum)
	if err != nil {
		panic("Problem with first read")
	}
	if numRead == 0 {
		panic("Problem with first read")
	}
	var finalResult = make(chan []string)
	for index := 0; index < testNum; index++ {
		var (
			numVillians, numHeros []int
			numPlayers            int
		)
		numRead, err = fmt.Scanf("%d", &numPlayers)
		if err != nil {
			panic("Problem with seconds read")
		}
		if numRead == 0 {
			panic("Problem with first read")
		}
		for i := 1; i <= 2 && scanner.Scan(); i++ {
			switch i {
			case 1:
				numVillians = numbers(scanner.Text())
			case 2:
				numHeros = numbers(scanner.Text())
			}
		}
		go doResults(numVillians, numHeros, testNum, finalResult)
	}

	// fmt.Println(numVillians)
	// fmt.Println(numHeros)
}
