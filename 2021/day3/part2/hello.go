package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type RatingNode struct {
	zeroNode              *RatingNode
	oneNode               *RatingNode
	bitIndex              int
	ratings               []string
	ratingsProcessedCount int
}

type RatingTree struct {
	root *RatingNode
}

func createNewNode(bitIndex int) *RatingNode {
	return &RatingNode{bitIndex: bitIndex, ratings: make([]string, 0)}
}

func (ratingTree *RatingTree) insert(rating string) {
	if ratingTree.root == nil {
		ratingTree.root = createNewNode(0)
	}
	ratingTree.root.insert(rating)
}

func (ratingNode *RatingNode) insert(rating string) {
	ratingNode.ratingsProcessedCount++
	ratingLength := len(rating)
	if ratingNode.bitIndex == ratingLength {
		ratingNode.ratings = append(ratingNode.ratings, rating)
		return
	}
	ratingBitValue := rating[ratingNode.bitIndex : ratingNode.bitIndex+1]
	if ratingBitValue == "0" {
		if ratingNode.zeroNode == nil {
			ratingNode.zeroNode = createNewNode(ratingNode.bitIndex + 1)
		}
		ratingNode.zeroNode.insert(rating)
		return
	} else if ratingBitValue == "1" {
		if ratingNode.oneNode == nil {
			ratingNode.oneNode = createNewNode(ratingNode.bitIndex + 1)
		}
		ratingNode.oneNode.insert(rating)
		return
	}
}

func createRatingTree() *RatingTree {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	ratingTree := &RatingTree{}

	for scanner.Scan() {
		currentTextValue := scanner.Text()
		ratingTree.insert(currentTextValue)
	}

	//ratingTree.insert("10")
	//ratingTree.insert("11")
	//ratingTree.insert("11")
	//ratingTree.insert("00")

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return ratingTree
}

func (ratingTree *RatingTree) calculate(mostCommon bool) []string {
	return processStuff(*ratingTree.root, mostCommon)
}

func processStuff(ratingNode RatingNode, mostCommon bool) []string {
	// exit condition
	if len(ratingNode.ratings) > 0 {
		return ratingNode.ratings
	}

	if ratingNode.oneNode != nil && ratingNode.zeroNode != nil && ratingNode.oneNode.ratingsProcessedCount == ratingNode.zeroNode.ratingsProcessedCount {
		if mostCommon {
			return processStuff(*ratingNode.oneNode, mostCommon)
		} else {
			return processStuff(*ratingNode.zeroNode, mostCommon)
		}
	}

	// oneNodeOverZeroNode := ratingNode.oneNode != nil && ratingNode.zeroNode == nil
	if ratingNode.oneNode != nil && ratingNode.zeroNode == nil {
		return processStuff(*ratingNode.oneNode, mostCommon)
	} else if ratingNode.oneNode == nil && ratingNode.zeroNode != nil {
		return processStuff(*ratingNode.zeroNode, mostCommon)
	}

	if mostCommon {
		if ratingNode.oneNode.ratingsProcessedCount > ratingNode.zeroNode.ratingsProcessedCount {
			return processStuff(*ratingNode.oneNode, mostCommon)
		} else {
			return processStuff(*ratingNode.zeroNode, mostCommon)
		}
	} else {
		if ratingNode.oneNode.ratingsProcessedCount < ratingNode.zeroNode.ratingsProcessedCount {
			return processStuff(*ratingNode.oneNode, mostCommon)
		} else {
			return processStuff(*ratingNode.zeroNode, mostCommon)
		}
	}
}

func binaryStringNumberToNumber(binaryStringNumber string) int64 {
	output, err := strconv.ParseInt(binaryStringNumber, 2, 64)
	if err != nil {
		panic(err)
	}
	return output
}

func main() {
	ratingTree := createRatingTree()
	co2 := ratingTree.calculate(false)
	oxygen := ratingTree.calculate(true)

	co2Number := binaryStringNumberToNumber(co2[0])
	oxygenNumber := binaryStringNumberToNumber(oxygen[0])
	//gammaRateString := repeatsCountToBinaryNumber(ratingTree)
	//gammaRateNumber := binaryStringNumberToNumber(gammaRateString)
	//epsilon := gammaRateNumber ^ (0b111111111111)
	//result := gammaRateNumber * epsilon
	fmt.Println(co2Number)
	fmt.Println(oxygenNumber)

	result := co2Number * oxygenNumber
	fmt.Println(result)
}
