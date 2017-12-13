package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
)

var cycleLimit = flag.Uint64("l", 0, "Cycle limit for calculating palindromic number (default 0 means no limit)")
var bigNumber = flag.String("n", "0", "Sets test number for calculating palindromic number (accepst only base 10 numbers)")
var truncateOperationLogs = flag.Bool("nolog", false, "Truncates reversal add operation logs")

func main() {
	flag.Parse()

	testNumber, validNumber := new(big.Int).SetString(*bigNumber, 10)

	if validNumber {
		reversalAdditionPalindromeCalculator(testNumber, *cycleLimit)
	} else {
		log.Fatalf("The number you entered (%v) is not valid!", *bigNumber)
	}
}

func reverseString(data string) string {
	dataAsRunes := []rune(data)

	for fI, rI := 0, len(dataAsRunes)-1; fI < rI; fI, rI = fI+1, rI-1 {
		dataAsRunes[fI], dataAsRunes[rI] = dataAsRunes[rI], dataAsRunes[fI]
	}

	return string(dataAsRunes)
}

func isPalindrome(testString string) bool {
	if reverseString(testString) == testString {
		return true
	}

	return false
}

func reversalAdditionPalindromeCalculator(pN *big.Int, cycleLimit uint64) {
	if isPalindrome(pN.String()) {
		fmt.Printf("%v is already palindromic number!\n", pN)
		return
	}

	var firstNum *big.Int
	var secondNum *big.Int
	res := new(big.Int).Set(pN)
	counter := big.NewInt(0)

	for {
		firstNum = new(big.Int).Set(res)
		secondNum, _ = new(big.Int).SetString(reverseString(firstNum.String()), 10)

		res = new(big.Int).Add(firstNum, secondNum)

		if !*truncateOperationLogs {
			fmt.Printf("%v + %v = %v\n", firstNum, secondNum, res)
		}

		counter = new(big.Int).Add(counter, big.NewInt(1))

		if isPalindrome(res.String()) {
			break
		}

		if cycleLimit > 0 && counter.Uint64() == cycleLimit {
			fmt.Printf("Cycle limit (%v) exceeded for number %v\n", cycleLimit, pN)
			return
		}
	}

	if !*truncateOperationLogs {
		fmt.Println("--- Done ---")
	}

	fmt.Printf("Found in %v steps. %v become %v as palindromic number\n", counter, pN, res.String())
}
