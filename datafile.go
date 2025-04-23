//пакет чтение из файла 
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats читает значение float64bp из каждый строчки файла.
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = appand(numbers, number)
	}
	err = file.Close()
	if err !=nil{
		return numbers, err
		}
	
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
