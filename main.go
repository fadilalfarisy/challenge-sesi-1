package main

import "fmt"

func main() {
	var i int = 21

	//menampilkan nilai i
	fmt.Printf("%d \n", i)

	//menampilkan tipe data dari variabel i
	fmt.Printf("%T \n", i)

	//menampilkan tanda %
	fmt.Printf("%s \n", "%")

	var j bool = true

	//menampilkan nilai boolean j
	fmt.Printf("%t \n", j)

	//menampilkan nilai base 2 dari i
	fmt.Printf("%b \n", i)

	////menampilkan nilai base 10 dari i
	fmt.Printf("%d \n", i)

	//menampilkan nilai base 8 dari i
	fmt.Printf("%o \n", i)

	//menampilkan nilai base 16 dari f
	fmt.Printf("%x \n", 15)

	//menampilkan nilai base 16 dari F
	fmt.Printf("%X \n", 15)

	uniChar := 'Я'

	//menampilkan unicode karakter Я
	fmt.Printf("%U \n", uniChar)

	var k float64 = 123.456

	//menampilkan float
	fmt.Printf("%f \n", k)

	//menampilkan float scientific
	fmt.Printf("%E \n", k)
}