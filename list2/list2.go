package list2

// Return the number of even ints in the given array. Note: the % "mod" operator
// computes the remainder, e.g. 5 % 2 is 1.
func CountEvens(a []int) int {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			count++
		}
	}
	return count
}

// Return the sum of the numbers in the array, returning 0 for an empty array.
// Except the number 13 is very unlucky, so it does not count and numbers that
// come immediately after a 13 also do not count.
func Sum13(a []int) int {
	sum := 0
	var unluckyind int
	if len(a) == 0 {
		return 0
	} else if len(a) == 1 && a[0] == 13 {
		return 0
	} else if len(a) == 2 && (a[0] == 13) {
		return 0
	} else if len(a) == 2 && a[1] == 13 {
		return a[0]
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] == 13 {
				unluckyind = i
			}
		}

		for j := 0; j < unluckyind; j++ {
			sum += a[j]
		}

		return sum
	}
}

// Given an array length 1 or more of ints, return the difference between the
// largest and smallest values in the array.
func BigDiff(a []int) int {
	var small int
	var large int
	if len(a) == 1 {
		return a[0]
	} else {
		small = a[0]
		large = a[0]
		for i := 0; i < len(a); i++ {
			if small < a[i] {
				small = a[i]
				continue
			}

		}
		for i := 0; i < len(a); i++ {
			if large > a[i] {
				large = a[i]
				continue
			}

		}
	}
	return large - small
}

// Return the sum of the numbers in the array, except ignore sections of numbers
//starting with a 6 and extending to the next 7 (every 6 will be followed by at
// least one 7). Return 0 for no numbers.
func Sum67(a []int) int {
	return 0
}

// Return the "centered" average of an array of ints, which we'll say is the mean
// average of the values, except ignoring the largest and smallest values in the array.
// If there are multiple copies of the smallest value, ignore just one copy, and
// likewise for the largest value. Use int division to produce the final average.
// You may assume that the array is length 3 or more.
func CenteredAverage(a []int) int {
	return 0
}

// Given an array of ints, return true if the array contains a 2 next to a 2 somewhere.
func Has22(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] == 2 && a[i+1] == 2 {
			return true
		}
	}
	return false
}
