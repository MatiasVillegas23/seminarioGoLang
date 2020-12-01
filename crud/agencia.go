package main

import (
	"fmt"
)

type Agencia struct {
	autos map[int] *Auto
}

type Auto struct {
	ID int
	marca string
	modelo string
}

func NewAgencia() Agencia {
	autos := make(map[int]*Auto)
	return Agencia{
		autos,
	}
}

func (a Agencia) Print(){
	for _,v := range a.autos {
		fmt.Println(v.ID, v.marca, v.modelo)
	}
}

func (a Agencia) Add(c Auto) {
	a.autos[c.ID] = &c
}

func (a Agencia) FindByID(ID int) *Auto{
	return a.autos[ID]
}

func (a Agencia) Delete(ID int) {
	delete(a.autos, ID)
}

func (a Agencia) Update(c Auto) {
	a.autos[c.ID] = &c
}