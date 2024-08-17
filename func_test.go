package main_test

import (
	"fmt"
	"slices"
	"testing"

	gfy "github.com/frontside/gfy"
)

func TestMap(t *testing.T) {

	testMapFunc := func(x int) string {
		return fmt.Sprintf("%da", x)
	}

	actual := []string{}

	for idx, v := range gfy.Map[int, string]([]int{1, 2, 3, 4}, testMapFunc) {
		actual = append(actual, fmt.Sprintf("%d", idx))
		actual = append(actual, v)
	}

	expected := []string{"0", "1a", "1", "2a", "2", "3a", "3", "4a"}

	if !slices.Equal(actual, expected) {
		t.Errorf("TestMap expected %v got %v", expected, actual)
	}

}
