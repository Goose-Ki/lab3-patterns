package builder

import "strings"

type Computer struct {
	CPU      string
	RAM      string
	Storage  string
	GPU      string
	Monitor  string
	Keyboard string
	Mouse    string
}

func (c *Computer) String() string {
	var builder strings.Builder
	builder.WriteString("Computer configuration:\n")

	if c.CPU != "" {
		builder.WriteString("  CPU: " + c.CPU + "\n")
	}
	if c.RAM != "" {
		builder.WriteString("  RAM: " + c.RAM + "\n")
	}
	if c.Storage != "" {
		builder.WriteString("  Storage: " + c.Storage + "\n")
	}
	if c.GPU != "" {
		builder.WriteString("  GPU: " + c.GPU + "\n")
	}
	if c.Monitor != "" {
		builder.WriteString("  Monitor: " + c.Monitor + "\n")
	}
	if c.Keyboard != "" {
		builder.WriteString("  Keyboard: " + c.Keyboard + "\n")
	}
	if c.Mouse != "" {
		builder.WriteString("  Mouse: " + c.Mouse + "\n")
	}

	return builder.String()
}

type ComputerBuilder interface {
	SetCPU(string) ComputerBuilder
	SetRAM(string) ComputerBuilder
	SetStorage(string) ComputerBuilder
	SetGPU(string) ComputerBuilder
	SetMonitor(string) ComputerBuilder
	SetKeyboard(string) ComputerBuilder
	SetMouse(string) ComputerBuilder
	Build() *Computer
}

type GamingComputerBuilder struct {
	computer *Computer
}

func NewGamingComputerBuilder() *GamingComputerBuilder {
	return &GamingComputerBuilder{
		computer: &Computer{},
	}
}

func (gcb *GamingComputerBuilder) SetCPU(cpu string) ComputerBuilder {
	gcb.computer.CPU = cpu
	return gcb
}

func (gcb *GamingComputerBuilder) SetRAM(ram string) ComputerBuilder {
	gcb.computer.RAM = ram
	return gcb
}

func (gcb *GamingComputerBuilder) SetStorage(storage string) ComputerBuilder {
	gcb.computer.Storage = storage
	return gcb
}

func (gcb *GamingComputerBuilder) SetGPU(gpu string) ComputerBuilder {
	gcb.computer.GPU = gpu
	return gcb
}

func (gcb *GamingComputerBuilder) SetMonitor(monitor string) ComputerBuilder {
	gcb.computer.Monitor = monitor
	return gcb
}

func (gcb *GamingComputerBuilder) SetKeyboard(keyboard string) ComputerBuilder {
	gcb.computer.Keyboard = keyboard
	return gcb
}

func (gcb *GamingComputerBuilder) SetMouse(mouse string) ComputerBuilder {
	gcb.computer.Mouse = mouse
	return gcb
}

func (gcb *GamingComputerBuilder) Build() *Computer {
	return gcb.computer
}

type ComputerDirector struct {
	builder ComputerBuilder
}

func NewComputerDirector(builder ComputerBuilder) *ComputerDirector {
	return &ComputerDirector{builder: builder}
}

func (cd *ComputerDirector) ConstructGamingPC() *Computer {
	return cd.builder.
		SetCPU("Intel Core i9").
		SetRAM("32GB DDR5").
		SetStorage("2TB NVMe SSD").
		SetGPU("NVIDIA RTX 4090").
		SetMonitor("27-inch 4K 144Hz").
		SetKeyboard("Mechanical Gaming Keyboard").
		SetMouse("Gaming Mouse 16000 DPI").
		Build()
}

func (cd *ComputerDirector) ConstructOfficePC() *Computer {
	return cd.builder.
		SetCPU("Intel Core i5").
		SetRAM("16GB DDR4").
		SetStorage("512GB SSD").
		SetMonitor("24-inch Full HD").
		Build()
}
