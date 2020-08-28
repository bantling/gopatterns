package main

import (
	"fmt"
)

// ClaimsProcess is the interface that describes the required template methods.
type ClaimsProcess interface {
	CreateClaim()
	PayClaim()
}

// ClaimsProcessValidate defines the optional claims process validation method
type ClaimsProcessValidate interface {
	ValidateClaim() error
}

// ClaimsProcessor is a template for making claims on different types of insurance by calling the template methods.
type ClaimsProcessor struct{}

func (p *ClaimsProcessor) ProcessClaim(cp ClaimsProcess) error {
	// Create the claim
	cp.CreateClaim()

	// Validate the claim if supported, stopping if it is not valid
	if cpv, isa := cp.(ClaimsProcessValidate); isa {
		if err := cpv.ValidateClaim(); err != nil {
			return err
		}
	}

	// Pay the claim
	cp.PayClaim()

	return nil
}

// LifeInsuranceClaimsProcess is for life insurance
type LifeInsuranceClaimsProcess struct{}

func (c *LifeInsuranceClaimsProcess) CreateClaim() {
	fmt.Println("Life insurance claim created")
}

func (c *LifeInsuranceClaimsProcess) ValidateClaim() error {
	fmt.Println("Life insurance claim validated")
	return nil
}

func (c *LifeInsuranceClaimsProcess) PayClaim() {
	fmt.Println("Life insurance claim paid")
}

// PropertyInsuranceClaimsProcess is for property insurance
type PropertyInsuranceClaimsProcess struct{}

func (c *PropertyInsuranceClaimsProcess) CreateClaim() {
	fmt.Println("Property insurance claim created")
}

func (c *PropertyInsuranceClaimsProcess) PayClaim() {
	fmt.Println("Property insurance claim paid")
}

func main() {
	allClaimsProcesses := []ClaimsProcess{
		&LifeInsuranceClaimsProcess{},
		&PropertyInsuranceClaimsProcess{},
	}

	p := &ClaimsProcessor{}
	for _, cp := range allClaimsProcesses {
		p.ProcessClaim(cp)
	}
}
