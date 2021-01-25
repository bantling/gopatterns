// SPDX-License-Identifier: Apache-2.0

package main

import "fmt"

func main() {
	for _, output := range []Output{
		Normalize(NewShipNormalizer(Ship{Name: "The Marianna", Type: "Fishing Vessel"})),
		Normalize(NewShipNormalizer(Ship{Name: "Selena", Type: "Dingy"})),
		Normalize(NewVehicleNormalizer(Vehicle{Make: "Jeep", Model: "Renegade"})),
		Normalize(NewVehicleNormalizer(Vehicle{Make: "Dodge", Model: "Nitro"})),
	} {
		fmt.Printf("%+v\n", output)
	}
}
