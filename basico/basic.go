package main

import (
	"fmt"
	"errors"
)

//funciones retornan mas de un valor
func suma(a,b int)(int, error){
 return a + b, nil
}

//si la funcion empieza con mayuscula es publica
func SumaGlobal(a,b int)(int, error){
	if a < b {
		return 0, errors.New("el valor a es menor a b")
	}
	return a + b, nil
}

//uso de estructuras for slice y map
func estructuras(){

	var arr[2] int
	arr[0] = 1
	arr[1] = 2

	for i,v := range arr{
		fmt.Println(i, v)
	}

	fmt.Println("-------")
	//var l []int
	l := make([]int, 3) 
	l = append(l, 10)
	fmt.Printf("%p\n", l)
	l = append(l, 100)
	fmt.Printf("%p\n", l)
	l = append(l, 1000)
	fmt.Printf("%p\n", l)

	for i,v := range l{
		fmt.Println(i, v)
	}

	fmt.Println("-------")
	m := make(map[int]string)
	m[0] = "a"
	m[1] = "b"
	
	for k,v := range m{
		fmt.Println(k, v)
	}
}

//Printable ..
type Printable interface{
	print()
}

//"objetos"
type person struct{
	name string
}

func (p *person) print(){
	fmt.Println("[person]", p.name)
}

/*func (p *person) clean(){
	p.name = ""
}*/

type figure struct{
	name string
}

func (f *figure) print(){
	fmt.Println("[figure]", f.name)
}

func invokePrint(p Printable){
	p.print()
}

//tipo funcion
type close func(s string)

func invokeClose(c close){
	c("Hola tipo funcion")
}

//startPoint
func main(){
	fmt.Println("hello world")

	fmt.Println(suma(1,2))
	r, err := SumaGlobal(1,2)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(r)
	}

	estructuras()
	
	p:= &person{name: "Juan"}
	f:= &figure{name: "Cube"}
	//p.print()
	//p.clean()
	//p.print()
	invokePrint(p)
	invokePrint(f)
	invokeClose(func(s string){
		fmt.Println(s)
	})

}