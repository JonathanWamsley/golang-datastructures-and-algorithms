package fibonacci

// source: https://www.youtube.com/watch?v=oBt53YbR9Kk&ab_channel=freeCodeCamp.org
// adapted from javascript to go code
//
// Write a function fib(n) that takes in a number as an argument.
// The function should return the n-th number of the Fibonacci sequence.

// The 0th number of the sequence is 0.
// The 1st number of the sequence is 1.

// fib(7) = fib(6) + fib(5)
// 		= fib(5) + fib(4) + fib(4) + fib(3)
//= fib(4) + fib(3) + fib(3) + fib(2) + fib(3) + fib(2) + fib(3) + fib(2)
//= fib(3) + fib(2) + fib(2) + fib(1) + fib(2) + fib(1) + 1 + fib(2) + fib(1) + 1 + fib(2) + fib(1) + 1
// = fib(2) + fib(1) + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1
// = 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 = 13
// O(2^n) since each call creates 2 more calls for each branch
// fib(50) = 1,125,899,906,842,624 ... very large amount of calls to make
// We can search for patterns within to improve performance
// can use a dictionary to keep track of fib seq. If it has not been seen, sore it

// Fibonacci - base case recursion
func Fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// FibonacciMemoization - uses a map to store previously seen elements
func FibonacciMemoization(n int) int {
	fibMap := map[int]int{1: 1, 2: 1}
	return fibonacciMemoization(n, fibMap)
}

// memoization stores previously seen fibonacci values
func fibonacciMemoization(n int, fibMap map[int]int) int {
	if value, ok := fibMap[n]; ok {
		return value
	}
	fibMap[n] = fibonacciMemoization(n-1, fibMap) + fibonacciMemoization(n-2, fibMap)
	return fibMap[n]
}

// tabulation
// building an array and using iteration instead of recurssion

// FibonacciTabulationA - uses past values to build new
func FibonacciTabulationA(n int) int {
	table := make([]int, n+1)
	table[0] = 0
	table[1] = 1
	for i := 2; i < len(table); i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[len(table)-1]
}

// FibonacciTabulationB - build from the bottom up
func FibonacciTabulationB(n int) int {
	table := make([]int, n+2)
	table[1] = 1
	for i := 0; i < n; i++ {
		table[i+1] += table[i]
		table[i+2] += table[i]
	}
	return table[n]
}
