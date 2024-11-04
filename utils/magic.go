package utils

import "strings"

// MagicSum returns the sum of n + n
func MagicSum(n int) int {
	return n + n
}

// MagicPow returns n raised to the power of n
func MagicPow(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= n
	}
	return result
}

// MagicOdd returns true if n is an odd number, false otherwise
func MagicOdd(n int) bool {
	return n%2 != 0
}

// MagicGrade returns a grade string based on n
func MagicGrade(n int) string {
	switch n {
	case 0:
		return "Zero"
	case 1:
		return "Bad"
	case 2:
		return "Ok"
	case 3:
		return "Nice"
	case 4:
		return "Awesome"
	case 5:
		return "Exceptional"
	default:
		return "Unknown"
	}
}

// MagicName returns a slice of strings containing [yourname] repeated n times
func MagicName(n int) []string {
	name := "Anisa" // Replace with your actual name if needed
	return strings.Split(strings.Repeat(name+",", n), ",")[:n]
}

// MagicTria returns the sum of numbers from 1 to n
func MagicTria(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// MagicChange modifies n by doubling its value
func MagicChange(n *int) {
	*n *= 2
}