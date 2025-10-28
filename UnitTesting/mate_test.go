package main

import "testing"

func TestSumar(t *testing.T) {
	// result := Sumar(2, 3)
	// expected := 5

	// if result != expected {
	// 	t.Errorf("Sumar(2, 3) = %d; want %d", result, expected)
	// }
	casos := []struct {
		a, b, e int
	}{
		{3, 2, 5},
		{0, 2, 2},
		{3, -2, 1},
	}

	for _, caso := range casos {
		r := Sumar(caso.a, caso.b)

		if r != caso.e {
			t.Errorf("Error al sumar: se esperaba %d, pero se obtuvo %d", caso.e, r)
		}
	}
}

func TestMayor(t *testing.T) {
	casos := []struct {
		a, b, e int
	}{
		{1, 2, 2},
		{6, 9, 9},
		{8, 0, 8},
	}

	for _, caso := range casos {
		r := Mayor(caso.a, caso.b)

		if r != caso.e {
			t.Errorf("Error al comparar: se esperaba %d, pero se obtuvo %d", caso.e, r)
		}
	}
}

func TestFibonacci(t *testing.T) {
	casos := []struct {
		n, e int
	}{
		{0, 0},
		{1, 1},
		{7, 13},
		{47, 2971215073},
	}

	for _, caso := range casos {
		r := Fibonacci(caso.n)

		if r != caso.e {
			t.Errorf("Error al calcular Fibonacci: se esperaba %d, pero se obtuvo %d", caso.e, r)
		}
	}
}
