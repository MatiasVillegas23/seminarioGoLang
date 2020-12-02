package main

import "testing"

func TestAgenciaAdd(t *testing.T) {

	a:= NewAgencia()
	c0 := a.FindByID(0)
	if c0 != nil {
		t.Error("El auto con el ID=0 ya existe")
	}

	a.Add(Auto{0, "Audi", "R8"})
	c0 = a.FindByID(0)

	if c0 == nil {
		t.Error("El auto con el ID=0 no fue agregado")
	}

	if c0.marca != "Audi" && c0.modelo != "R8" {
		t.Error("El auto con el ID=0 no tiene los datos correctos")
	}
}