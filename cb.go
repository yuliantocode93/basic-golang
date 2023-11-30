package main

import "fmt"

//* Number
// func main() {
// 	fmt.Println("satu =", 1)
// 	fmt.Println("dua =", 2)
// 	fmt.Println("tiga koma lima =", 3.5)
// }

//* boolean
// func main() {
// 	fmt.Println("benar =", true)
// 	fmt.Println("salah =", false)
// }

//* tipe data string
// func main() {
// 	fmt.Println(" ini tipe data string")
// 	fmt.Println(" ini tipe data string kedua")
// }
// func main() {
// 	fmt.Println(len("satu"))// 4
// 	fmt.Println(" ini tipe data string kedua"[0])// 32
// 	fmt.Println(" ini tipe data string kedua"[1])// 105
// }

//* Variable
// func main() {
// 	var name string

// 	name = "Variable"
// 	fmt.Println((name))

// 	name = "Variable yang baru"
// 	fmt.Println((name))
// }

// func main() {
// 	var name = "Variable"
// 	fmt.Println((name))

// 	name = "Variable yang baru"
// 	fmt.Println((name))
// }

// func main() {
// 	name := "Variable"
// 	fmt.Println((name))

// 	name = "Variable yang baru"
// 	fmt.Println((name))

// 	var (
// 		firstName ="Kwon"
// 		lastName = "yuli"
// 	)
// 	fmt.Println(firstName)
// 	fmt.Println(lastName)
// }

//* Constant
// func main() {
// 	const firstName ="kwon"
// 	const lastName = "yuli"

// 	fmt.Println(firstName)
// 	fmt.Println(lastName)
// }

//* konversi tipe data
// func main(){
// 	var nilai32 int32 = 32768
// 	var nilai64 int64 = int64(nilai32)
// 	var nilai16 int16 = int16(nilai32)

// 	fmt.Println(nilai32)// 32768
// 	fmt.Println(nilai64)// 32768
// 	fmt.Println(nilai16)// -32768

// 	var name = "kwon"
// 	var e = name[0]
// 	var eString = string(e)

// 	fmt.Println(name)
// 	fmt.Println(e)
// 	fmt.Println(eString)
// }

//* Type Declaration
// func main()  {
// 	type NoKTP string

// 	var ktpKwon NoKTP = "1234566"
// 	var contohKtp = "1234567890"
// 	fmt.Println(ktpKwon)
// 	fmt.Println(contohKtp)
// }

//* Operasi matematika
// func main() {
// 	var a = 10
// 	var b = 10
// 	var c = a + b
// 	fmt.Println(c)
// }

//* Augmented Assignments
// func main () {
// 	var i = 10
// 	i += 10 //10s
// 	fmt.Println(i)//20
// 	i += 5
// 	fmt.Println(i)//25
// }

//* Unary Operator
// func main () {
// 	var j = 1
// 	j++ // j = j + 1 // 2
// 	fmt.Println(j)
// 	j++ // j = j + 1 // 3
// 	fmt.Println(j)

// 	j-- // 2
// 	fmt.Println(j)
// 	j-- // 1
// 	fmt.Println(j)
// }

//* Operasi Perbandingan
func main () {
	var name1 = "kwon"
	var name2 = "kwon"

	var result = name1 == name2
	var result2 = name1 != name2

	fmt.Println(result)// true
	fmt.Println(result2)// false

	var a = "a"
	var b = "b"

	var hasil = a < b

	fmt.Println(hasil)// true
}