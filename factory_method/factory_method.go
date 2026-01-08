package factory_method

import "fmt"

type Transport interface {
	Deliver() string
}

type Truck struct{}

func (t *Truck) Deliver() string {
	return "Delivering by land in a box"
}

type Ship struct{}

func (s *Ship) Deliver() string {
	return "Delivering by sea in a container"
}

type Plane struct{}

func (p *Plane) Deliver() string {
	return "Delivering by air in a cargo hold"
}

type Logistics interface {
	CreateTransport() Transport
	PlanDelivery() string
}

type RoadLogistics struct{}

func (rl *RoadLogistics) CreateTransport() Transport {
	return &Truck{}
}

func (rl *RoadLogistics) PlanDelivery() string {
	transport := rl.CreateTransport()
	return fmt.Sprintf("Road logistics: %s", transport.Deliver())
}

type SeaLogistics struct{}

func (sl *SeaLogistics) CreateTransport() Transport {
	return &Ship{}
}

func (sl *SeaLogistics) PlanDelivery() string {
	transport := sl.CreateTransport()
	return fmt.Sprintf("Sea logistics: %s", transport.Deliver())
}

type AirLogistics struct{}

func (al *AirLogistics) CreateTransport() Transport {
	return &Plane{}
}

func (al *AirLogistics) PlanDelivery() string {
	transport := al.CreateTransport()
	return fmt.Sprintf("Air logistics: %s", transport.Deliver())
}
