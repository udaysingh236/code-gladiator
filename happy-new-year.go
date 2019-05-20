package main

import (
	"bufio"
	"bytes"
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

func arrayToString(input []int) string {
	var buffer bytes.Buffer
	for i := 0; i < len(input); i++ {
		buffer.WriteString(strconv.Itoa(input[i]))

	}

	return buffer.String()
}

func doResults(numTickets []int, tempChan chan<- string) {
	var incl = numTickets[0]
	var excl int
	var exclNew int
	maxSum := numTickets[0]
	var exclArray []int
	var inclArray []int
	if incl > 0 {
		inclArray = append(inclArray, incl)
		// inclArray[0] = incl
	}
	for index := 1; index < len(numTickets); index++ {
		if maxSum < numTickets[index] {
			maxSum = numTickets[index]
		}
		// current max sum excluding index
		if incl > excl {
			exclNew = incl
		} else {
			exclNew = excl
		}

		tempExcl := exclArray
		//max sum sequence excluding current element or arr[i]
		if incl > excl {
			exclArray = inclArray
		} else {
			exclArray = exclArray
		}
		if numTickets[index] > 0 {
			// inclArray = append(tempExcl, numTickets[index])
			inclArray = append([]int{numTickets[index]}, tempExcl...)
		}

		// current max sum including i
		incl = excl + numTickets[index]
		excl = exclNew

	}
	if incl > excl {
		// fmt.Println("Inclusive array is and max sum is: ", inclArray, incl)
		tempChan <- arrayToString(inclArray)
	} else {
		// fmt.Println("Exclusive array is and max sum is: ", exclArray, excl)
		tempChan <- arrayToString(exclArray)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numRead, err := fmt.Scan(&testNum)
	// fmt.Println("testNum: ", testNum)
	if err != nil {
		panic("Problem with first read")
	}
	if numRead == 0 {
		panic("Problem with first read")
	}
	var finalResult = []chan string{}
	for index := 0; index < testNum; index++ {
		var (
			numTickets []int
			numPlayers []int
		)
		// numRead, err = fmt.Scan(&numPlayers)
		// // fmt.Println("numPlayers: ", numPlayers)
		// if err != nil {
		// 	fmt.Println("numPlayers: ", numPlayers, err)
		// 	panic("Problem with seconds read")
		// }
		// if numRead == 0 {
		// 	panic("Problem with first read")
		// }

		// fmt.Scanf("%d", &numRead)
		// fmt.Println("numRead: ", numRead)

		tmpChan := make(chan string)
		finalResult = append(finalResult, tmpChan)
		for i := 0; i < 2 && scanner.Scan(); i++ {
			// fmt.Println("In the input loop")
			switch i {
			case 0:
				numPlayers = numbers(scanner.Text())
				// fmt.Println("numPlayers: ", numPlayers)
			case 1:
				numTickets = numbers(scanner.Text())
				// fmt.Println("numTickets", numTickets)
			}
		}
		if len(numTickets) == numPlayers[0] {
			go doResults(numTickets, tmpChan)
		} else {
			fmt.Println("You have provided less number of tickets, you promised", numPlayers[0])
		}
	}

	for _, value := range finalResult {
		fmt.Println(<-value)
	}
	// fmt.Println(numVillians)
	// fmt.Println(numHeros)
}
