package main

import "fmt"

func main() {

	var a int = 10
	var p *int = &a

	fmt.Println("Valor de a:", a)
	fmt.Println("Endereço de a:", &a)
	fmt.Println("Valor de p (endereço de a):", p)
	fmt.Println("Valor apontado por p (valor de a):", *p)

	p1 := Pessoa{
		Nome:  "João",
		Idade: 30,
	}

	fmt.Println("Nome:", p1.Nome)
	fmt.Println("Idade:", p1.Idade)
	fmt.Println("Telefone:", p1.Telefone())

	p1.SetIdade(35)
	fmt.Println("Idade após atualização:", p1.Idade)

	p1.SetTelefone("123456789")
	fmt.Println("Telefone após atualização:", p1.Telefone())

}

type Pessoa struct {
	Nome     string
	Idade    int
	telefone *string
}

func (p Pessoa) Telefone() string {
	if p.telefone == nil {
		return ""
	}
	return *p.telefone
}

// func (p Pessoa) SetIdade(idade int) { // não funciona porque o método recebe uma cópia da struct, então a alteração não afeta o objeto original
// 	p.Idade = idade
// }

func (p *Pessoa) SetIdade(idade int) { // funciona porque o método recebe um ponteiro para a struct, então a alteração afeta o objeto original
	p.Idade = idade
}

func (p *Pessoa) SetTelefone(telefone string) {
	p.telefone = &telefone
}
