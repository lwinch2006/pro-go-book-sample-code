package tests

import (
	"chapter31/services"
	"fmt"
	"sort"
	"testing"
)

func Test_UnitTesting1_1_Total_ShouldPass(t *testing.T) {
	testValues := []int{10, 20, 30}
	_, total := services.UnitTesting1_1(testValues)
	expected := 60

	if total != expected {
		t.Fatalf("Expected (%v), Got (%v)", expected, total)
	}
}

func Test_UnitTesting1_1_Sort_ShouldPass(t *testing.T) {
	testValues := []int{10, 20, 30}
	sortedNumbers, _ := services.UnitTesting1_1(testValues)

	if !sort.IntsAreSorted(sortedNumbers) {
		t.Fatalf("Numbers not sorted: %v", sortedNumbers)
	}
}

func Test_UnitTesting1_1_Sort_SubTests_ShouldPass(t *testing.T) {
	testValues := [][]int{
		{10, 20, 30},
		{11, 21, 31},
		{12, 22, 32},
	}

	for index, testValue := range testValues {
		t.Run(fmt.Sprintf("Running subtest %v", index), func(subT *testing.T) {
			sortedNumbers, _ := services.UnitTesting1_1(testValue)

			if !sort.IntsAreSorted(sortedNumbers) {
				subT.Fatalf("Numbers not sorted: %v", sortedNumbers)
			}
		})
	}
}

func Test_UnitTesting1_1_Total_SubTests_ShouldPass(t *testing.T) {
	testValues := [][]int{
		{10, 20, 30},
		{11, 21, 31},
		{12, 22, 32},
	}
	expectedTotals := []int{23, 34, 46}

	for index, testValue := range testValues {
		t.Run(fmt.Sprintf("Running subtest %v", index), func(subT *testing.T) {
			if t.Failed() {
				subT.SkipNow()
			}

			_, total := services.UnitTesting1_1(testValue)
			if total != expectedTotals[index] {
				subT.Fatalf("Expected (%v), Got (%v)", expectedTotals[index], total)
			}
		})
	}

}
