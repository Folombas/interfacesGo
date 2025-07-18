package main

import "fmt"

// Инерфейс для всех поваров
type Chef interface {
	Cook(dish string) string
	Speciality() string
}

// Интерфейс для оборудования
type Appliace interface {
	Use() string
}

// Повар-грильщик
type GrillChef struct {
	name 		string
	grillType 	string
}

func (gc GrillChef) Cook(dish string) string {
	return fmt.Sprintf("%s готовит на %s: %s", gc.name, gc.grillType, dish)
}

func (g GrillChef) Speciality() string {
	return "стейки и барбекю"
}

func (g GrillChef) Use() string {
	return fmt.Sprintf("Разжигае %s", g.grillType)
}

// Кондитер 
type PastChef struct {
	name 		string
	ovenTemp 	int
}

