package main

func main() {
	// целые числа
	var i int = 25 // int8, int16, int32, int64
	var autoInt = -22
	var bigInt int64 = 1<<32 - 1 // bit-shift, 2^32 - 1
	var unsignedInt uint = 778899
	var unsignedBigInt uint64 = 1<<64 - 1 // uint8, unit16, uint32, unit64
	println("integers", i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	// С плавующей точкой
	var f32 float32 = 3.145  // Pi
	var f64 float64 = 3.1456 // Pi
	println("float: ", f32, f64)

	// Булевые значения
	var b = true
	println("flag is", b)

	// Строки
	var hello string = "Hi\n\t"
	var world = "World"
	println(hello, world)

	// Перменные имеют значения по умлочанию
	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool
	println("default values: ", defaultInt+1, defaultFloat, defaultString, defaultBool)

	var rawBinary byte = '\x27' //бинарные данные
	println("rawBinary", string(rawBinary), rawBinary)

	meaningOfLive := 42 // краткое объявление
	println("Meaning of life is ", meaningOfLive)

	// приведение типов
	println("float to int ", int(f32))
	println("int to string ", string(48))

	// комплексные числа
	z := 2 + 3i
	println("complex number: ", z)

	// операции со строками
	s1 := "Denis"
	s2 := "Volkov"
	fullName := s1 + s2
	println("My name is: ", fullName, len(fullName))

	// Многострочное объявление
	escaping := `Hello\r\n
	World`
	println("<pre>: ", escaping, "</pre>")

	// несколько переменных
	var v1, v2 string = "v1", "v2"
	println(v1, v2)

	var (
		m0 int = 12
		m2     = "string"
		m3     = 23
	)

	println(m0, m2, m3)
}
