package main

import (
	"fmt"

	"./model"
)

func main() {
	user := model.NewObject(1, map[string]interface{}{
		"name":  "Pedro Oliveira",
		"age":   30,
		"email": "@.com",
		"CPF":   "000",
	})

	user.SetAttribute("endereco", model.NewObject(2, map[string]interface{}{
		"rua":         "Rua",
		"numero":      531,
		"complemento": "A803",
		"CEP":         551133,
	}))

	user.SetAttribute("endereco.telefone", 991547900)

	fmt.Print("User Name: ")
	fmt.Println(user.GetAttribute("name"))
	fmt.Print("CPF: ")
	fmt.Println(user.GetAttribute("CPF"))
	fmt.Print("E-mail: ")
	fmt.Println(user.GetAttribute("email"))

	fmt.Println("Endereço: ")
	fmt.Print("------Rua: ")
	fmt.Println(user.GetAttribute("endereco.rua"))
	fmt.Print("------Número: ")
	fmt.Println(user.GetAttribute("endereco.numero"))
	fmt.Print("------CEP: ")
	fmt.Println(user.GetAttribute("endereco.CEP"))
	fmt.Print("------Complemento: ")
	fmt.Println(user.GetAttribute("endereco.complemento"))
	fmt.Print("------Telefone: ")
	fmt.Println(user.GetAttribute("endereco.telefone"))
}
