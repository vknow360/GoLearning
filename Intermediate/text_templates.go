package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	// tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?")
	// if err != nil {
	// 	panic(err)
	// }

	// data := map[string]interface{}{
	// 	"name": "John",
	// }

	// err = tmpl.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	templates := map[string]string{
		"welcome":  "Welcome, {{.name}}! We're glad to have you here.\n",
		"farewell": "Goodbye, {{.name}}! Hope to see you again soon.\n",
		"birthday": "Happy Birthday, {{.name}}! Wishing you a wonderful year ahead.\n",
	}

	// for key, tmplStr := range templates {
	// 	tmpl, err := template.New(key).Parse(tmplStr)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	data := map[string]interface{}{
	// 		"name": name,
	// 	}
	// 	err = tmpl.Execute(os.Stdout, data)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	parsedTemplates := make(map[string]*template.Template)

	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		fmt.Print("Enter template name (welcome, farewell, birthday) or 'exit' to quit: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "exit" {
			break
		}
		tmpl, exists := parsedTemplates[choice]
		if !exists {
			fmt.Println("Template not found. Please try again.")
			continue
		}
		err := tmpl.Execute(os.Stdout, map[string]interface{}{
			"name": name,
		})
		if err != nil {
			panic(err)
		}
	}

}
