package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
	Age    int // Иногда в шаблонах есть возраст, если его нет в твоем шаблоне, удали эту строку
}

func (p Personal) Print() {
	// Обрати внимание на %.2f для веса
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}