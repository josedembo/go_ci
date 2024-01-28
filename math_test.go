package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	valor_esperado := 30

	if total != valor_esperado {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado %d.", total, valor_esperado)
	}
}
