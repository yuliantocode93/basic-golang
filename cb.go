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
// func main () {
// 	var name1 = "kwon"
// 	var name2 = "kwon"

// 	var result = name1 == name2
// 	var result2 = name1 != name2

// 	fmt.Println(result)// true
// 	fmt.Println(result2)// false

// 	var a = "a"
// 	var b = "b"

// 	var hasil = a < b

// 	fmt.Println(hasil)// true
// }

//* Operasi Boolean
// func main () {
// var nilaiAkhir = 90
// var absensi = 80

// var lulusNilaiAkhir = nilaiAkhir > 80
// var lulusAbsensi = absensi > 80

// var lulus = lulusNilaiAkhir && lulusAbsensi

// fmt.Println(lulus)// false

// 	var nilaiAkhir = 90
// 	var absensi = 81

// 	var lulusNilaiAkhir = nilaiAkhir > 80
// 	var lulusAbsensi = absensi > 80

// 	var lulus = lulusNilaiAkhir && lulusAbsensi

// 	fmt.Println(lulus)// true
// }

//* Tipe Data Array
// func main () {
// 	var names [3]string

// 	names[0] = "kwon"
// 	names[1] = "yuli"
// 	names[2] = "anto"

// 	fmt.Println(names[0])
// 	fmt.Println(names[1])
// 	fmt.Println(names[2])

// 	var value = [3]int{90, 80, 95}
// 	fmt.Println(value)

// 	var values = [...] int {
// 		90,
// 		80,
// 		95,
// 		100,
// 		110,
// 	}

// 	fmt.Println(values)
// 	fmt.Println(values[0])
// 	fmt.Println(values[1])
// 	fmt.Println(values[2])
// 	fmt.Println(len(values))
// }

//* Tipe Data Slice

// func main () {
// 	names := [...]string{"Kwon","yuli","anto","echo","echo1", "echo2"}

// 	slice1 := names[4:6]
// 	fmt.Println(slice1)

// 	slice2 := names[:3]
// 	fmt.Println(slice2)

// 	slice3 := names[3:]
// 	fmt.Println(slice3)

// 	var slice4 = names[:]
// 	fmt.Println(slice4)

// 	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

// 	daySlice1 := days[5:] // sabtu, minggu
// 	fmt.Println(daySlice1)

// 	daySlice1[0] = "sabtu baru"
// 	daySlice1[1] = "minggu baru"
// 	fmt.Println(daySlice1)
// 	fmt.Println(days)

// 	daySlice2 := append(daySlice1, "Libur Baru")
// 	daySlice2[0] = "Sabtu Lama"

// 	fmt.Println(daySlice1)
// 	fmt.Println(daySlice2)
// 	fmt.Println(days)

// 	var newSlice = make([]string, 2, 5)
// 	newSlice[0] = "echo1"
// 	newSlice[1] = "echo2"

// 	fmt.Println(newSlice)
// 	fmt.Println(len(newSlice))
// 	fmt.Println(cap(newSlice))

// 	// menambah array newslice
// 	newSlice2 := append(newSlice, "echo3")

// 	fmt.Println(newSlice2)
// 	fmt.Println(len(newSlice2))
// 	fmt.Println(cap(newSlice2))

// 	newSlice2[0] = "kwon"
// 	fmt.Println(newSlice2)
// 	fmt.Println(newSlice)

// 	fromSlice := days[:]
// 	toSlice := make([]string, len(fromSlice), cap(fromSlice))

// 	copy(toSlice, fromSlice)

// 	fmt.Println(fromSlice)
// 	fmt.Println(toSlice)

// 	iniArray := [...]int{1,2,3,4,5}
// 	iniSlice := []int{1,2,3,4,5}

// 	fmt.Println(iniArray)
// 	fmt.Println(iniSlice)
// }

//* Tipe data Map
func main () {
	person := map[string]string{
		"name": "Kwon",
		"address" : "Guandong",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Kwon"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}