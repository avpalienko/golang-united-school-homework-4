package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands  = errors.New("expecting two operands, but received more or less")
	errorInvalidOperand  = errors.New("invalid operand")
	errorInvalidOpertion = errors.New("invalid operation")
	errorExtraSymbols    = errors.New("extra symbols")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

const (
	stSignOrDig  = 0
	stWaitFirst  = 1
	stWaitSecond = 2
	stWaitOp     = 3
	stEnd        = 4
)

func StringSum(input string) (output string, err error) {
	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	inp := strings.Replace(input, " ", "", -1)
	token := ""
	output = ""
	op1 := int64(0)
	op2 := int64(0)
	oper := ""
	state := stWaitFirst
	i := 0
	for ; i < len(inp); i++ {
		if state == stWaitFirst || state == stWaitSecond {
			if inp[i] == '+' || inp[i] == '-' {
				token += inp[i : i+1]
				i++
			}
			for ; i < len(inp) && isDig(inp[i:i+1]); i++ {
				token += inp[i : i+1]
			}
			i--
			if op, er := strconv.ParseInt(token, 10, 64); er != nil {
				return "", fmt.Errorf("%w: %s", errorInvalidOperand, token)
			} else {
				if state == stWaitFirst {
					op1 = op
					state = stWaitOp
				} else {
					op2 = op
					state = stEnd
				}
				token = ""
			}

		} else if state == stWaitOp {
			switch inp[i] {
			case '+':
				oper = "+"
			case '-':
				oper = "-"
			default:
				return "", fmt.Errorf("%w: %s", errorInvalidOpertion, inp[i:i+1])
			}
			state = stWaitSecond
		} else if state == stEnd {
			return "", fmt.Errorf("%w: %s", errorExtraSymbols, inp[i:])
		}
	}
	if state != stEnd {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	if oper == "+" {
		return strconv.FormatInt(op1+op2, 10), nil
	} else {
		return strconv.FormatInt(op1-op2, 10), nil
	}
}

func isDig(input string) bool {
	if m, _ := regexp.MatchString("[0-9]", input); m {
		return true
	}
	return false
}
