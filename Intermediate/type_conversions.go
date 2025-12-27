package main

func main() {
	var a int = 42
	b := int32(a)
	c := float64(b)
	println(a, b, c)

	e := 3.14
	f := int(e)
	println(e, f)

	g := "Hello"
	h := []byte(g)
	println(g, h)

	i := []byte{32, 120, 72}
	j := string(i)
	println(i, j)
}
