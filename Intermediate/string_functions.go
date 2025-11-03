package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello, World!"

	fmt.Println(len(str))

	fmt.Println(str[7])

	fmt.Println(str[0:5])

	fruits := "apple,banana,cherry"

	parts := strings.Split(fruits, `,`)
	fmt.Println(parts)

	countries := []string{"USA", "Canada", "Mexico"}
	result := strings.Join(countries, `, `)
	fmt.Println(result)

	fmt.Println(strings.Contains(str, "World"))

	fmt.Println(strings.Replace(str, "World", "Go", 1))

	strwspace := " Hello, World! "
	fmt.Println(strings.TrimSpace(strwspace))

	fmt.Println(strings.ToUpper(str))

	fmt.Println(strings.Count("Hello", `l`))
	fmt.Println(strings.HasPrefix("Hello", "he"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	str6 := "aabccudgcfygfcycyha"
	fmt.Println(utf8.RuneCountInString(str6))

	fmt.Println(strings.Repeat("Go! ", 3))

	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("Builder!")
	fmt.Println(builder.String())
}
