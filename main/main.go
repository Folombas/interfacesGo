package main

import "fmt"

// Объявляем интерфейс Flyer с методом Fly()
type Flyer interface {
	Fly() string
}

// Структура Bird с собственным методом Fly()
type Bird struct {
	Name  string
	Wings int
}

func (b Bird) Fly() string {
	return fmt.Sprintf("%s летит, махая %d крыльями", b.Name, b.Wings)
}

// Структура Airplane с собственным методом Fly()
type Airplane struct {
	Model string
}

func (a Airplane) Fly() string {
	return fmt.Sprintf("Самолет %s летит на реактивной тяге", a.Model)
}

// Структура Superman (демонстрация неявной реализации)
type Superman struct {
	PowerLevel int
}

func (s Superman) Fly() string {
	return fmt.Sprintf("Супермен летит с силой %d%%", s.PowerLevel)
}

// Функция принимающая ЛЮБОЙ объект, реализующий Flyer
func StartFlight(f Flyer) {
	fmt.Println(f.Fly())
}

func main() {
	// Создаем объекты разных типов
	sparrow := Bird{"Воробей", 2}
	boeing := Airplane{"Boeing 747"}
	clark := Superman{100}

	// Вызываем метод Fly() напрямую
	fmt.Println(sparrow.Fly())
	fmt.Println(boeing.Fly())
	fmt.Println(clark.Fly())

	fmt.Println("\nЧерез интерфейс:")
	// Используем интерфейсный тип
	var flyer Flyer

	flyer = sparrow // Bird реализует Flyer
	StartFlight(flyer)

	flyer = boeing // Airplane реализует Flyer
	StartFlight(flyer)

	flyer = clark // Superman реализует Flyer
	StartFlight(flyer)

	// Демонстрация полиморфизма
	fmt.Println("\nПолиморфизм в срезе:")
	flyingObjects := []Flyer{
		Bird{"Орел", 2},
		Airplane{"Сухой Суперджет"},
		Superman{85},
	}

	for _, obj := range flyingObjects {
		fmt.Println(obj.Fly())
	}
}

// Пример показывает, как совершенно разные сущности (птицы, самолеты, супергерои)
// могут быть объединены общим поведением (умением летать) через интерфейс.
