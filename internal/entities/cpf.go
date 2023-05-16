package entities

import (
	"errors"
	"regexp"
	"strconv"
)

type CPF struct {
	Value string `json:"cpf"`
}

var (
	DIGIT_1_FACTOR int = 10
	DIGIT_2_FACTOR int = 11
)

func NewCPF(cpf string) (*CPF, error) {
	c := &CPF{Value: cpf}

	return c, nil
}

func Validate(cpf string) error {
	if len(cpf) == 0 {
		return errors.New("CPF is required")
	}

	cleanCPF := regexp.MustCompile(`\D+`).ReplaceAllString(cpf, "")
	cpfSlice := convertCPFToIntSlice(cleanCPF)

	if !isValidLength(cleanCPF) {
		return errors.New("invalid cpf")
	}

	if areDigitsEqual(cleanCPF) {
		return errors.New("invalid cpf")
	}

	dg1 := calculateDigit(cpfSlice[:9], DIGIT_1_FACTOR)
	dg2 := calculateDigit(cpfSlice[:10], DIGIT_2_FACTOR)

	checkDigit := cpfSlice[9:11]
	checkDigitString := strconv.Itoa(checkDigit[0]) + strconv.Itoa(checkDigit[1])
	calculatedCheckDigit := strconv.Itoa(dg1) + strconv.Itoa(dg2)

	if checkDigitString != calculatedCheckDigit {
		return errors.New("invalid cpf")
	}
	return nil
}

func calculateDigit(cpf []int, factor int) int {
	sum := 0
	for _, number := range cpf {
		if factor > 1 {
			sum += number * factor
		}
		factor--
	}

	rest := sum % 11

	if rest < 2 {
		return 0
	} else {
		return 11 - rest
	}
}

func convertCPFToIntSlice(cpf string) []int {
	intSlice := make([]int, len(cpf))
	for i, number := range cpf {
		intSlice[i] = int(number - '0')
	}
	return intSlice
}

func isValidLength(cpf string) bool {
	return len(cpf) == 11
}

func areDigitsEqual(cpf string) bool {
	firstDigit := cpf[0]
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != firstDigit {
			return false
		}
	}
	return true
}
