package main

// Цель - показать, как интерфейс заставляет разные объекты
// делать одно и то же действие по-своему.
import "fmt"

// 1. Объявляем интерфейс - "контракт" для всех говорящих существ
type Speaker interface {
	Speak() string // Всё, что умеет говорить, должно иметь метод Speak()
}

// 2. Создаём разные типы животных (каждый - отдельный тип)

// Собака
type Dog struct {
	Name string
}

// Метод Speak() для собаки
func (d Dog) Speak() string {
	return "Гав! Меня зовут " + d.Name
}

// Кот
type Cat struct {
	Name string
}

// Метод Speak() для кота (своя реализация)
func (c Cat) Speak() string {
	return "Мяу! Меня зовут " + c.Name
}

// Утка
type Duck struct {
	Name string
}

// Метод Speak() для утки (своя реализация)
func (d Duck) Speak() string {
	return "Кря! Меня зовут " + d.Name
}

// Лиса
type Fox struct {}

func (f Fox) Speak() string {
	return "Фыр! Что сказала лиса?"
}

// Корова
type Cow struct {
	Name string
}

func (c Cow) Speak() string {
	return "Муууу! Меня зовут " + c.Name
}

// 3. Главная магия: функция, работающая с ЛЮБЫМ объектом, который реализует интерфейс Speaker
func MakeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	// Создаём животных
	rex := Dog{Name: "Рекс"}
	murka := Cat{Name: "Мурка"}
	donald := Duck{Name: "Дональд"}
	burenka := Cow{Name: "Бурёнкаё"}

	// 4. Собираем их в один список через интерфейс
	zoo := []Speaker{rex, murka, donald, Fox{}, burenka}

	// 5. Заставляем всех говорить одной командой!
	for _, animal := range zoo {
		MakeSpeak(animal)
	}
}
