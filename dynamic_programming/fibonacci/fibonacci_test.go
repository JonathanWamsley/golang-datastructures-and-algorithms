package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert.EqualValues(t, 1, Fibonacci(1))
	assert.EqualValues(t, 1, Fibonacci(2))
	assert.EqualValues(t, 2, Fibonacci(3))
	assert.EqualValues(t, 3, Fibonacci(4))
	assert.EqualValues(t, 5, Fibonacci(5))
	assert.EqualValues(t, 8, Fibonacci(6))
	assert.EqualValues(t, 13, Fibonacci(7))
	assert.EqualValues(t, 21, Fibonacci(8))
	assert.EqualValues(t, 21, Fibonacci(8))
}

func TestFibonacciMemoization(t *testing.T) {
	assert.EqualValues(t, 1, FibonacciMemoization(1))
	assert.EqualValues(t, 1, FibonacciMemoization(2))
	assert.EqualValues(t, 2, FibonacciMemoization(3))
	assert.EqualValues(t, 3, FibonacciMemoization(4))
	assert.EqualValues(t, 5, FibonacciMemoization(5))
	assert.EqualValues(t, 8, FibonacciMemoization(6))
	assert.EqualValues(t, 13, FibonacciMemoization(7))
	assert.EqualValues(t, 21, FibonacciMemoization(8))
	assert.EqualValues(t, 21, FibonacciMemoization(8))
	assert.EqualValues(t, 12586269025, FibonacciMemoization(50))
}

func TestFibonacciTabulationA(t *testing.T) {
	assert.EqualValues(t, 1, FibonacciTabulationA(1))
	assert.EqualValues(t, 1, FibonacciTabulationA(2))
	assert.EqualValues(t, 2, FibonacciTabulationA(3))
	assert.EqualValues(t, 3, FibonacciTabulationA(4))
	assert.EqualValues(t, 5, FibonacciTabulationA(5))
	assert.EqualValues(t, 8, FibonacciTabulationA(6))
	assert.EqualValues(t, 13, FibonacciTabulationA(7))
	assert.EqualValues(t, 21, FibonacciTabulationA(8))
	assert.EqualValues(t, 12586269025, FibonacciTabulationA(50))
}

func TestFibonacciTabulationB(t *testing.T) {
	assert.EqualValues(t, 1, FibonacciTabulationB(1))
	assert.EqualValues(t, 1, FibonacciTabulationB(2))
	assert.EqualValues(t, 2, FibonacciTabulationB(3))
	assert.EqualValues(t, 3, FibonacciTabulationB(4))
	assert.EqualValues(t, 5, FibonacciTabulationB(5))
	assert.EqualValues(t, 8, FibonacciTabulationB(6))
	assert.EqualValues(t, 13, FibonacciTabulationB(7))
	assert.EqualValues(t, 21, FibonacciTabulationB(8))
	assert.EqualValues(t, 12586269025, FibonacciTabulationB(50))
}
