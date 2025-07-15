package main

// Зарядка электро-устройств
import "fmt"

// Всё, что можно заряжать
type Charger interface {
	Charge() string
}

// Телефон
type Phone struct{}

func (p Phone) Charge() string {
	return "Телефон заряжается 🔋"
}

// Ноутбук
type Laptop struct{}

func (l Laptop) Charge() string {
	return "Ноутбук заряжается 💻"
}

// Электроавтомобиль
type EV struct {
	Model string
}

func (e EV) Charge() string {
	return fmt.Sprintf("%s заряжается 🚗⚡", e.Model)
}

// Зарядная станция
func ChargeDevice(device Charger) {
	fmt.Println(device.Charge())
}

func main() {
	ChargeDevice(Phone{})
	ChargeDevice(Laptop{})
	ChargeDevice(EV{"Tesla Model S"})
}
