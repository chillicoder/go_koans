package about_arrays

import "./koans"
import "fmt"

func TestCreatingArray(t *koans.T) {
	var array [10]int
	t.AssertTrue(koans.Int__ == len(array)) //Length of array is part of its type.
}

func TestArraysAreValues(t *koans.T) {
	array1 := [...]int{1, 2, 3}
	var array2 [3]int
	array2 = array1
	t.AssertTrue(koans.String__ == fmt.Sprintf("%v", array2))
}

func TestAccessingArrayElements(t *koans.T) {
	a := [...]string{"peanut", "butter", "and", "jelly"}

	t.AssertTrue(koans.String__ == a[0])
	t.AssertTrue(koans.String__ == a[3])
}
