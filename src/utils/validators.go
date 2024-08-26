package utils

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/mvrilo/go-cpf"
)

// checks if a field is a string and the string is a valid CPF (no numbers)
var IsCpf validator.Func = func(fl validator.FieldLevel) bool {
	fieldValue, fieldIsNumeric := fieldIsNumericString(fl)
	if !fieldIsNumeric {
		return false
	}

	valid, _ := cpf.Valid(fieldValue)

	return valid
}

// checks if a field is a string and the string is a valid CNPJ (no numbers)
var IsCnpj validator.Func = func(fl validator.FieldLevel) bool {
	fieldValue, fieldIsNumeric := fieldIsNumericString(fl)
	if !fieldIsNumeric {
		return false
	}

	return isCnpj(fieldValue)
}

var (
	cnpjFirstDigitTable  = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjSecondDigitTable = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

// returns the field value and if it contains a string convertable to a int
func fieldIsNumericString(fl validator.FieldLevel) (string, bool) {
	fieldValue, ok := fl.Field().Interface().(string)
	if !ok {
		return fieldValue, false
	}

	if _, err := strconv.Atoi(fieldValue); err != nil {
		return fieldValue, false
	}

	return fieldValue, true
}

func sumDigit(s string, table []int) int {
	if len(s) != len(table) {
		return 0
	}

	sum := 0

	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}

	return sum
}

// validates CNPJ without punctuation
func isCnpj(cnpj string) bool {
	if len(cnpj) != 14 {
		return false
	}

	firstPart := cnpj[:12]
	sum1 := sumDigit(firstPart, cnpjFirstDigitTable)
	rest1 := sum1 % 11
	d1 := 0

	if rest1 >= 2 {
		d1 = 11 - rest1
	}

	secondPart := fmt.Sprintf("%s%d", firstPart, d1)
	sum2 := sumDigit(secondPart, cnpjSecondDigitTable)
	rest2 := sum2 % 11
	d2 := 0

	if rest2 >= 2 {
		d2 = 11 - rest2
	}

	finalPart := fmt.Sprintf("%s%d", secondPart, d2)
	return finalPart == cnpj
}
