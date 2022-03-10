package fungsi

import "fmt"

type Ortu struct {
	Name string
}
type Data struct {
	Ortu
	name string
	age  int
}

func Introduce(name string) {
	fmt.Println("Hy my name is ", name)
}

func Variadic(num ...int) {
	fmt.Println((num))
}

func Main2() {
	// z := Ortu{Name: "reza"}
	datadiri := Data{name: "Arum", age: 21, Ortu: Ortu{Name: "reza"}}
	fmt.Println(datadiri)
}
