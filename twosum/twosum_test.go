package twosum_test

import (
	"reflect"
	"testing"

	"github.com/brodiep21/leetcode/twosum"
)

func TestTwoSum(t *testing.T) {
	t.Run("Testing 2 length array", func(t *testing.T) {
		force := []int{2, 4}
		target := 6
		got := twosum.TwoSum(force, target)

		want := []int{0, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("Testing a 4 length array with a negative number", func(t *testing.T) {
		force := []int{2, -20, 6, 15}
		target := 14
		got := twosum.TwoSum(force, target)

		want := []int{1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("Testing a 4 length array with all negative numbers", func(t *testing.T) {
		force := []int{-3, -6, -20, -13}
		target := 19
		got := twosum.TwoSum(force, target)

		want := []int{1, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
}
