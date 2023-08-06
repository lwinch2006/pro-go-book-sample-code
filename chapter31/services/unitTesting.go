package services

import (
	"chapter31/utils"
	"log"
	"os"
	"sort"
)

func UnitTesting1_1(numbers []int) (sortedNumbers []int, total int) {
	sortedNumbers = make([]int, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Ints(sortedNumbers)

	for _, val := range sortedNumbers {
		total += val
	}

	return
}

func UnitTesting1() {
	utils.Printfln("UnitTesting1()")

	numbers := []int{1, 4, 2, 3}
	sortedNumbers, total := UnitTesting1_1(numbers)

	utils.Printfln("Sorted numbers: %v", sortedNumbers)
	utils.Printfln("Total: %v", total)
}

func unitTesting2_2() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func UnitTesting2() {
	utils.Printfln("UnitTesting2()")

	numbers := []int{1, 4, 2, 3}
	sortedNumbers, total := UnitTesting1_1(numbers)

	unitTesting2_2()

	log.Printf("Sorted numbers: %v", sortedNumbers)
	log.Printf("Total: %v", total)
}

func unitTesting3_3() (newLogger *log.Logger, file *os.File) {
	var err error

	if file, err = os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666); err == nil {
		newLogger = log.New(file, "", log.Lshortfile|log.Ltime)
	}

	return
}

func UnitTesting3() {
	utils.Printfln("UnitTesting3()")

	numbers := []int{1, 4, 2, 3}
	sortedNumbers, total := UnitTesting1_1(numbers)

	newLogger, file := unitTesting3_3()
	defer file.Close()

	newLogger.Printf("Sorted numbers: %v", sortedNumbers)
	newLogger.Printf("Total: %v", total)
}
