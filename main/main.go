package main

import "fmt"

// Интерфейс для всех поваров
type Chef interface {
    Cook(dish string) string
    Specialty() string
}

// Интерфейс для оборудования
type Appliance interface {
    Use() string
}

// Повар-грильщик
type GrillChef struct {
    name      string
    grillType string
}

func (g GrillChef) Cook(dish string) string {
    return fmt.Sprintf("%s готовит на %s: %s", g.name, g.grillType, dish)
}

func (g GrillChef) Specialty() string {
    return "стейки и барбекю"
}

func (g GrillChef) Use() string {
    return fmt.Sprintf("Разжигает %s", g.grillType)
}

// Кондитер
type PastryChef struct {
    name    string
    ovenTemp int
}

func (p PastryChef) Cook(dish string) string {
    return fmt.Sprintf("%s выпекает при %d°C: %s", p.name, p.ovenTemp, dish)
}

func (p PastryChef) Specialty() string {
    return "десерты и выпечка"
}

func (p PastryChef) Preheat() string {
    return fmt.Sprintf("Разогревает духовку до %d°C", p.ovenTemp)
}

// Су-шеф, использующий приборы
type SousChef struct {
    name string
    tool Appliance
}

func (s SousChef) Cook(dish string) string {
    return fmt.Sprintf("%s использует %s для приготовления %s", s.name, s.tool.Use(), dish)
}

func (s SousChef) Specialty() string {
    return "техника приготовления"
}

// Функция работы кухни (использует интерфейс Chef)
func kitchenService(chef Chef, dish string) {
    if dish == "" {
        dish = chef.Specialty()
    }
    fmt.Println("== Заказ принят ==")
    fmt.Println(chef.Cook(dish))
    fmt.Printf("Специализация: %s\n\n", chef.Specialty())
}

func main() {
    // Создаем поваров
    grillMaster := GrillChef{"Антонио", "угольный гриль"}
    dessertExpert := PastryChef{"Мари", 180}
    
    // Создаем су-шефа с прибором
    mixer := struct {
        brand string
    }{brand: "KitchenAid"}
    
    sousChef := SousChef{
        name: "Жан",
        tool: ApplianceFunc(func() string {
            return fmt.Sprintf("миксер %s", mixer.brand)
        }),
    }

    // Демонстрация работы через интерфейс Chef
    kitchenService(grillMaster, "стейк Рибай")
    kitchenService(dessertExpert, "крем-брюле")
    kitchenService(sousChef, "безе")

    // Дополнительная демонстрация полиморфизма
    chefs := []Chef{grillMaster, dessertExpert, sousChef}
    fmt.Println("=== Все повара за работой ===")
    for _, chef := range chefs {
        kitchenService(chef, "")
    }
}

// Вспомогательный тип для реализации Appliance
type ApplianceFunc func() string

func (f ApplianceFunc) Use() string {
    return f()
}