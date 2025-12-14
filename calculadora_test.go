package main

import (
	"testing"
)

func TestSomarShouldCorrect(t *testing.T) {
	// Arrange
	soma := somar(5, 2, 1, 98)

	// Act
	esperado := 106

	// Assert
	if soma != esperado {
		t.Error("Esperado: ", esperado, "Retornado: ", soma)
	}
}

func TestSomarShouldIncorrect(t *testing.T) {
	soma := somar(2, 1, 4, 3)

	esperado := 9

	if soma != esperado {
		t.Error("Esperado: ", esperado, "Retornado: ", soma)
	}
}

func TestMultiplicarShouldCorrect(t *testing.T) {
	produto := multiplicar(2, 5, 2, 9)

	esperado := 180

	if produto != esperado {
		t.Error("Esperado: ", esperado, "Retornado: ", produto)
	}
}

func TestMultiplicarShouldIncorrect(t *testing.T) {
	produto := multiplicar(5, 2, 1)

	esperado := 15

	if produto != esperado {
		t.Error("Esperado: ", esperado, "Retornado: ", produto)
	}
}

func TestSubtrairShouldCorrect(t *testing.T) {
	resto := subtrair(5, 1, 25)

	esperado := -21

	if resto != esperado {
		t.Error("Esperado: ", esperado, "Retornado: ", resto)
	}
}

func TestSubtrairShouldIncorrect(t *testing.T) {
	resto := subtrair(5, 2, 42)

	esperado := -50

	if resto != esperado {
		t.Error("Esperado: ", esperado, "Retornado: ", resto)
	}
}

func TestDividirShouldCorrect(t *testing.T) {
	quociente := dividir(80, 2, 4)

	esperado := dividir(10)

	if quociente != esperado {
		t.Error("Esperado: ", esperado, "Retornado: ", quociente)
	}
}

func TestDividirShouldIncorrect(t *testing.T) {
	quociente := dividir(40, 2, 8)

	esperado := 12

	if quociente != esperado {
		t.Error("Esperado: ", esperado, "Retornado: ", quociente)
	}
}
