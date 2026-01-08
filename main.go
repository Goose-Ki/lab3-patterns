package main

import (
	"fmt"

	"lab3/abstract_factory"
	"lab3/builder"
	"lab3/factory_method"
	"lab3/singleton"
)

func main() {
	fmt.Println("=== Лаба 3 - Порождающие паттерны проектирования ===")

	fmt.Println("\n1. Singleton:")
	db1 := singleton.GetInstance()
	db2 := singleton.GetInstance()
	fmt.Println(db1.Query("SELECT * FROM users"))
	fmt.Printf("Are instances the same? %v\n", db1 == db2)

	fmt.Println("\n2. Factory Method:")
	roadLogistics := &factory_method.RoadLogistics{}
	seaLogistics := &factory_method.SeaLogistics{}
	airLogistics := &factory_method.AirLogistics{}

	fmt.Println(roadLogistics.PlanDelivery())
	fmt.Println(seaLogistics.PlanDelivery())
	fmt.Println(airLogistics.PlanDelivery())

	fmt.Println("\n3. Abstract Factory:")
	fmt.Println("Windows UI:")
	abstract_factory.CreateUI(&abstract_factory.WindowsFactory{})

	fmt.Println("\nMacOS UI:")
	abstract_factory.CreateUI(&abstract_factory.MacOSFactory{})

	fmt.Println("\n4. Builder:")
	builderInstance := builder.NewGamingComputerBuilder()
	director := builder.NewComputerDirector(builderInstance)

	gamingPC := director.ConstructGamingPC()
	fmt.Println(gamingPC.String())

	officePC := director.ConstructOfficePC()
	fmt.Println(officePC.String())

	// Построение без директора
	customPC := builder.NewGamingComputerBuilder().
		SetCPU("AMD Ryzen 7").
		SetRAM("64GB").
		SetStorage("4TB SSD").
		SetGPU("AMD RX 7900 XT").
		Build()
	fmt.Println(customPC.String())
}
