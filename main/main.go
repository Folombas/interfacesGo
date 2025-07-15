package main

import "fmt"

// Интерфейс - всё, что можно "нажать"
type Button interface {
	Press() string
}

// Реальная кнопка
type PhysicalButton struct{}

func (b PhysicalButton) Press() string {
	return "Щёлк!"
}

// Экранная кнопка
type TouchButton struct{}

func (b TouchButton) Press() string {
	return "Тач!"
}

// Кнопка на клавиатуре
type KeyboardButton struct {
	Key string
}

func (b KeyboardButton) Press() string {
	return fmt.Sprintf("Нажата клавиша %s", b.Key)
}

func main() {
	buttons := []Button{
		PhysicalButton{},
		TouchButton{},
		KeyboardButton{Key: "Enter"},
		KeyboardButton{Key: "Shift"},
		KeyboardButton{Key: "Tab"},
	}

	for _, button := range buttons {
		fmt.Println(button.Press())
	}
}
