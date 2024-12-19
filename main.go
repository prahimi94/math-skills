package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	inputFileName := os.Args[1]

	/* this part is for reading input file */
	data, err := os.ReadFile(inputFileName)
	if checkError(err) {
		return
	}
	inputStringData := string(data) // data is in the format of bytes so it should be converted to string
	inputStringArray := strings.Split(inputStringData, "\n")
	var inputNumbersArray []float64
	for _, v := range inputStringArray {
		if v == "" {
			continue
		}
		intValue, err := strconv.ParseFloat((strings.TrimSpace(string(v))), 64)
		if checkError(err) {
			return
		}

		inputNumbersArray = append(inputNumbersArray, intValue)
	}

	calculatedAverage := round(calcAverage(inputNumbersArray))
	calculatedMedian := round(calcMedian(inputNumbersArray))
	calculatedVariance := round(calcVariance(inputNumbersArray))
	calculatedStandardDeviation := round(calcStandardDeviation(inputNumbersArray))
	fmt.Printf("Average: %v\nMedian: %v\nVariance: %v\nStandardDeviation: %v\n", calculatedAverage, calculatedMedian, calculatedVariance, calculatedStandardDeviation)
}

func checkError(e error) bool {
	if e != nil {
		fmt.Printf("problem occured: %v \n", e.Error())
		return true

	}
	return false
}

func calcAverage(input []float64) float64 {
	/* The average, or mean, is the sum of a set of values divided by the number of values. */
	var output float64
	var inputSum float64
	inputCnt := float64(len(input)) /* inputCnt is number of values */
	for _, v := range input {
		inputSum += v
	}
	output = float64(inputSum / inputCnt)
	return output
}

func calcMedian(input []float64) float64 {
	var output float64
	/* for calculating median we should Sort the dataset at first */
	slices.Sort(input)
	inputCnt := len(input) /* inputCnt is number of values */
	if len(input)%2 == 0 {
		/* If inputCnt is even, the median is the average of the values at positions inputCnt/2 and (inputCnt/2)+1 but our index starts from 0 so the median is the average of the values at positions (inputCnt/2)-1 and inputCnt/2 */
		output = calcAverage([]float64{input[(inputCnt/2)-1], input[inputCnt/2]})
	} else {
		/* If inputCnt is odd, the median is the value at position (n+1)/2.  but our index starts from 0 so the median is  the value at position n/2 */
		output = float64(input[inputCnt/2])
	}
	return output
}

func calcVariance(input []float64) float64 {
	var output float64
	/*  Variance measures how much the values in a dataset differ from the average (mean). It quantifies the spread of the data points. */

	inputCnt := len(input)
	inputAvg := calcAverage(input)
	for _, v := range input {
		output += math.Pow(float64(v-float64(inputAvg)), 2)
	}
	output /= float64(inputCnt)
	return output
}

func calcStandardDeviation(input []float64) float64 {
	/* Standard deviation is the square root of the variance. It provides a measure of the average distance of each data point from the mean. */

	output := math.Sqrt(calcVariance(input))
	return output
}

func round(number float64) int64 {
	if number > 9223372036854775807 || number < -9223372036854775808 {
		fmt.Println("can not convert the float number to integer")
		os.Exit(1)
	}
	return int64(math.Round(number))
}
