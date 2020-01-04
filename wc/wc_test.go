package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestGetCounts(t *testing.T) {
	testString := "Here is test data let's see if we get what is expected."
	expectedResults := []uint64{1, 12, 56, 56, 55}
	scanner := bufio.NewScanner(strings.NewReader(testString))
	r := getCounts(scanner)

	for index, element := range r {
		if element != expectedResults[index] {
			t.Errorf("Total %d was incorrect. got: %d, want: %d", index, element, expectedResults[index])
		} else {
			fmt.Println(element, expectedResults[index])
		}
	}
}
