package main
// Аналог: Связка ключей — разные замки открываются разными ключами, но принцип один.
import "fmt"

// Всё, что можно открыть ключом
type Lock interface {
	Open(key string) string
}

// Дверь дома
type HouseDoor struct {}

func (d HouseDoor) Open(key string) string {
	if key == "Домашний ключ" {
		return "Дверь открыта!"
	}
	return "Неверный ключ для двери дома!"
}

// Сейф
type Safe struct {
	Code int
}

func (s Safe) Open(key string) string {
	if key == "Секретный код" {
		return fmt.Sprintf("Сейф с кодом %d открыт!", s.Code)
	}
	return "Неверный ключ для сейфа!"
}

func TryOpen(lock Lock, key string) {
	fmt.Println(lock.Open(key))
}

func main() {
	TryOpen(HouseDoor{}, "Домашний ключ")
	TryOpen(Safe{Code: 1234}, "пароль")
	TryOpen(Safe{Code: 5678}, "Секретный код")
}