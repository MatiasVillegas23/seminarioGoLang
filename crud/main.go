package main

import (
	"fmt"
)

func main(){

	a:= NewAgencia()
	a.Add(Auto{0, "Ford", "Mustang"})
	a.Add(Auto{1, "Chevrolet", "Impala"})
	a.Add(Auto{2, "Hyundai ", "Tucson"})

	a.Print()

	ID := 0
	c1 := a.FindByID(ID)
	if c1 != nil {
		fmt.Println("se encontro el auto con el ID: ", ID)
	}

	a.Delete(ID)
	a.Print()

	a.Update(Auto{1, "Chevrolet", "Silverado"})
	a.Print()
}