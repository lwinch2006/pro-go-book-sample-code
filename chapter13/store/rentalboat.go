package store

type Crew struct {
	Captain, FirstOfficer string
}

func newCrew(captain, firstOfficer string) *Crew {
	return &Crew{captain, firstOfficer}
}

type RentalBoat struct {
	*Boat
	IncludeCrew bool
	*Crew
}

func NewRentalBoat(name, category string, price float64, capacity int, motorized, includeCrew bool, captain, firstOfficer string) *RentalBoat {
	return &RentalBoat{NewBoat(name, category, price, capacity, motorized), includeCrew, newCrew(captain, firstOfficer)}
}
