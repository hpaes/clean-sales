package entities

import (
	"clean-sales/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenValidCPF_ShoulReturnCPF(t *testing.T) {
	cases := []struct {
		name string
		cpf  string
	}{
		{"CPF 1", "029.496.970-54"},
		{"CPF 2", "093.926.420-08"},
		{"CPF 3", "896.762.430-19"},
		{"CPF 4", "972.049.230-90"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cpf, err := entities.NewCPF(c.cpf)
			assert.Nil(t, err)
			assert.Equal(t, c.cpf, cpf.Value)
		})
	}
}

func TestGivenCPFWithAllDigitsEqual_ThenShouldThrowError(t *testing.T) {
	cases := []struct {
		name string
		cpf  string
	}{
		{"CPF 1", "111.111.111-11"},
		{"CPF 2", "222.222.222-22"},
		{"CPF 3", "333.333.333-33"},
		{"CPF 4", "444.444.444-44"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := entities.Validate(c.cpf)
			assert.NotNil(t, err)
			assert.EqualError(t, err, "invalid cpf")
		})
	}
}

func TestGivenInvalidCPF_ThenShouldThrowError(t *testing.T) {
	cases := []struct {
		name string
		cpf  string
	}{
		{"CPF 1", "029.496.970-53"},
		{"CPF 2", "093.926.420-07"},
		{"CPF 3", "896.762.430-18"},
		{"CPF 4", "972.049.230-91"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := entities.Validate(c.cpf)
			assert.NotNil(t, err)
			assert.EqualError(t, err, "invalid cpf")
		})
	}
}
