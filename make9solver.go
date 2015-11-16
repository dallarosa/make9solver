package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	numbers   NumberSet
	operators OperatorSet
	mainSize  int
)

type NumberSet []int

func (ns *NumberSet) Set(value string) (err error) {
	for _, v := range strings.Split(value, ",") {
		intval, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*ns = append(*ns, intval)
	}
	return nil
}

func (ns *NumberSet) String() string {
	return ""
}

type OperatorSet []string

func (ns *OperatorSet) Set(value string) (err error) {
	*ns = strings.Split(value, ",")
	return nil
}

func (ns *OperatorSet) String() string {
	return strings.Join(*ns, ",")
}

type NumOp struct {
	Numbers   []int
	Operators []string
}

func (nop *NumOp) Calculate() (result int) {

	result = nop.Numbers[0]
	for i := 0; i < len(nop.Operators); i++ {
		result = operator(result, nop.Numbers[i+1], nop.Operators[i])
	}
	return result
}

func operator(a int, b int, op string) int {

	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "x":
		return a * b
	}
	return 0
}

func init() {
	flag.Var(&numbers, "numbers", "set of numbers")
	flag.Var(&operators, "operators", "set of operators")

	flag.Parse()
}

func main() {
	fmt.Println(operators)
	fmt.Println(numbers)
	mainSize = int(math.Min(float64(len(numbers)-1), float64(len(operators)+1)))
	nums := GetNumbers(numbers, mainSize)
	ops := opPermutations(operators, mainSize-1)
	res := make([]NumOp, 0)
	for _, nv := range nums {
		for _, ov := range ops {
			if len(ov) == len(nv)-1 {
				res = append(res, NumOp{Numbers: nv, Operators: ov})
			}
		}
	}
	for _, c := range res {
		sum := c.Calculate()
		if sum == 9 {
			fmt.Println(c)
		}
	}
}

func GetNumbers(nSet []int, size int) (result []NumberSet) {
	result = make([]NumberSet, 0)
	tr := permutations(nSet, size)
	for _, v := range tr {
		if len(v) > 1 && len(v) <= size {
			result = append(result, v)
		}
	}
	return
}

// Combination: A Combination of "n" elements (C(n,S)) is all the possible combinations of size "n" of a given set of elements "S"
// Example: Given the set A={1,2,3,4,5,6}, all the C(1,A) will be all the possible combinations of 1 element, using the elements in A
// Result: C(1,A) = {{1},{2},{3},{4},{5},{6}}
// C(2,A) = {{1,2},{1,3},{1,4},{1,5},{1,6},{2,3},{2,4},{2,5},{2,6},{3,4},{3,5},{3,6},{4,5},{5,6}}
//
// In this case, as we have operators, the combinations will be an combination of numbers and operators:
//
// N = 1: <n> <op> <n>
// N = 2: <n> <op> <n> <op> <n>
// ...
//
// For N = 1, you'll need 1 operators and 2 numbers.
// For N = 2, 2 operators and 3 numbers.
// ...
// N = n, n operators and n + 1 numbers
//
// The maximum N will be limited by the number of elements in each set: N <= len(operators) && N <= len(numbers) - 1

func permutations(set []int, size int) (result [][]int) {
	if len(set) == 1 {
		result = make([][]int, 0)
		result = append(result, []int{set[0]})
		return result
	} else {
		pivot := set[0]
		a := make([]int, len(set)-1)
		copy(a, set[1:])
		b := permutations(a, size-1)
		result = make([][]int, 0)
		result = append(result, []int{pivot})
		for _, v := range b {
			result = append(result, v)
			if len(v) <= mainSize {
				result = append(result, append(v, pivot))
			}
		}
		return result
	}
}

func opPermutations(set []string, size int) (result [][]string) {
	if len(set) == 1 {
		result = make([][]string, 0)
		result = append(result, []string{set[0]})
		return result
	} else {
		pivot := set[0]
		a := make([]string, len(set)-1)
		copy(a, set[1:])
		b := opPermutations(a, size-1)
		result = make([][]string, 0)
		result = append(result, []string{pivot})
		for _, v := range b {
			result = append(result, v)
			if len(v) <= mainSize-1 {
				result = append(result, append(v, pivot))
			}
		}
		return result
	}
}
