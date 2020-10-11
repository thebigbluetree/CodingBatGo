package list1

// Given an array of ints, return true if 6 appears as either the first or last
// element in the array. The array will be length 1 or more.
func FirstLast6(i []int) bool {
	if (i[0] == 6) || (i[len(i)-1] == 6) {
		return true
	} else {
		return false
	}
}

// Given 2 arrays of ints, a and b, return true if they have the same first
// element or they have the same last element. Both arrays will be length 1 or more.
func CommonEnd(a, b []int) bool {

	if (a[0] == b[0]) || (a[len(a)-1] == b[len(b)-1]) {
		return true
	} else {
		return false
	}
}

// Given an array of ints length 3, return a new array with the elements in
// reverse order, so [3]int{1, 2, 3} becomes [3]int{3, 2, 1}.
func Reverse3(a [3]int) [3]int {
	var b [3]int
	b[0] = a[2]
	b[1] = a[1]
	b[2] = a[0]
	return b

	//for unlimited number of elements using slice.
	// var b = make([]int, 0, len(a))
	// for i := len(a) - 1; i >= 0; i-- {
	// 	b = append(b, a[i])
	// }

	// return b
}

// Given 2 int arrays, a and b, each length 3, return a new array length 2
// containing their middle elements.
func MiddleWay(a, b [3]int) [2]int {
	var c [2]int
	c[0] = a[1]
	c[1] = b[1]
	return c
}

// Given an array of ints, return true if the array is length 1 or more, and
// the first element and the last element are equal.
func SameFirstLast(i []int) bool {
	if (len(i) >= 1) && (i[0] == i[len(i)-1]) {
		return true
	} else {
		return false
	}
}

// Given an array of ints length 3, return the sum of all the elements
func Sum3(i [3]int) int {
	sum := 0
	for _, v := range i {
		sum = sum + v
	}
	return sum
}

// Given an array of ints length 3, figure out which is larger, the first or
// last element in the array, and set all the other elements to be that value.
// Return the changed array.
func MaxEnd3(a [3]int) [3]int {

	if a[0] > a[2] {
		b := [3]int{a[0], a[0], a[0]}
		return b
	} else {
		b := [3]int{a[2], a[2], a[2]}
		return b
	}
}

// Given an array of ints, return a new array length 2 containing the first and
// last elements from the original array. The original array will be length 1 or more.
func MakeEnds(i []int) [2]int {
	var j [2]int
	j[0] = i[0]
	j[1] = i[len(i)-1]
	return j
}

// Return an int array length 3 containing the first 3 digits of pi, {3, 1, 4}.
func MakePi() []int {
	return []int{3, 1, 4}
}

// Given an array of ints length 3, return an array with the elements
// "rotated left" so [3]int{1, 2, 3} yields [3]int{2, 3, 1}.
func RotateLeft3(a [3]int) [3]int {
	var b [3]int
	b[0] = a[1]
	b[1] = a[2]
	b[2] = a[0]
	return b

}

// Given an array of ints, return the sum of the first 2 elements in the array.
// If the array length is less than 2, just sum up the elements that exist,
// returning 0 if the array is length 0.
func Sum2(a []int) int {
	if len(a) == 0 {
		return 0
	} else if len(a) == 1 {
		return a[0]
	} else {
		return a[0] + a[1]
	}
}

// Given an int array length 2, return true if it contains a 2 or a 3.
func Has23(i [2]int) bool {
	if (i[0] == 2 || i[0] == 3) || (i[1] == 2 || i[1] == 3) {
		return true
	} else {
		return false
	}

}
