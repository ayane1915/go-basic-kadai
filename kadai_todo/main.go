package main

import "fmt"

type ToDos []ToDo
type ToDo struct {
	ID        int
	Title     string
	Completed bool
}

func (p *ToDo) Complete() ToDo {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", p.ID)
	p.Completed = true
	return *p
}

func (p ToDos) printTodos() {
	for _, todo := range p {
		fmt.Printf("[%t]（ID:%d）%s\n", todo.Completed, todo.ID, todo.Title)
	}
}

func main() {
	todo := ToDos{
		{
			ID:        1,
			Title:     "学習計画",
			Completed: false,
		},
		{
			ID:        2,
			Title:     "環境構築",
			Completed: false,
		},
		{
			ID:        3,
			Title:     "基礎文法",
			Completed: false,
		},
	}

	todo.printTodos()
	for i := range todo {
		if todo[i].ID == 1 || todo[i].ID == 2 {
			todo[i].Complete()
		}
	}
	todo.printTodos()
}
