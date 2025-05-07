package main

import "fmt"

func main() {
	var a int = 32
	b := int32(a) //32
	c := float64(b) //32
	fmt.Println(a, b, c)

	e := 3.14
	f := int(e)	//3
	fmt.Println(e,f)

	//Type(value)
	g := "Hello @ こんにちは 🧑 привет"
	var h []byte
	h = []byte(g)
	fmt.Println(h)
	i := []byte{255, 120, 72}
	j := string(i)
	fmt.Println(j)
}
