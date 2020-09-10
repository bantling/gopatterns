package main

// Normalizer describes a process that produces an outout for some unknown type of input
type Normalizer interface {
	// HandleSource produces the Source value
	HandleSource() string

	// HandleName produces the Name value
	HandleName() string

	// HandleCategory produces the Category value
	HandleCategory() string
}

// Normalize produces an Output given an input
// It is a visitor for all the fields of Output
func Normalize(n Normalizer) Output {
	output := Output{}
	output.Source = n.HandleSource()
	output.Name = n.HandleName()
	output.Category = n.HandleCategory()

	return output
}

// ShipNormalizer normalizes a Ship
type ShipNormalizer struct {
	Normalizer
	ship Ship
}

// NewShipNormalizer constructs a ShipNormalizer
func NewShipNormalizer(ship Ship) *ShipNormalizer {
	return &ShipNormalizer{ship: ship}
}

func (n ShipNormalizer) HandleSource() string {
	return "Ship"
}

func (n ShipNormalizer) HandleName() string {
	return n.ship.Name
}

func (n ShipNormalizer) HandleCategory() string {
	switch n.ship.Type {
	case "Fishing Vessel":
		return "Commercial"
	}

	return "Pleasure"
}

// VehicleNormalizer normalizes a Vehicle
type VehicleNormalizer struct {
	Normalizer
	vehicle Vehicle
}

// NewVehicleNormalizer constructs a VehicleNormalizer
func NewVehicleNormalizer(vehicle Vehicle) *VehicleNormalizer {
	return &VehicleNormalizer{vehicle: vehicle}
}

func (n VehicleNormalizer) HandleSource() string {
	return "Vehicle"
}

func (n VehicleNormalizer) HandleName() string {
	return n.vehicle.Model
}

func (n VehicleNormalizer) HandleCategory() string {
	return n.vehicle.Make
}
