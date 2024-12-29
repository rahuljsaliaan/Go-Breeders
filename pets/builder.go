package pets

import "errors"

// Interface for the Pet struct that defines the methods that can be used to set the fields of the struct. A builder pattern is used to construct the struct.
type PetInterface interface {
	SetSpecies(string) *Pet
	SetBreed(string) *Pet
	SetMinWeight(int) *Pet
	SetMaxWeight(int) *Pet
	SetWeight(int) *Pet
	SetDescription(string) *Pet
	SetLifeSpan(int) *Pet
	SetGeographicOrigin(string) *Pet
	SetColor(string) *Pet
	SetAge(int) *Pet
	SetAgeEstimated(bool) *Pet
}

func NewPetBuilder() *Pet {
	return &Pet{}
}

func (p *Pet) SetSpecies(species string) *Pet {
	p.Species = species
	return p
}

func (p *Pet) SetBreed(breed string) *Pet {
	p.Breed = breed
	return p
}

func (p *Pet) SetMinWeight(minWeight int) *Pet {
	p.MinWeight = minWeight
	return p
}

func (p *Pet) SetMaxWeight(maxWeight int) *Pet {
	p.MaxWeight = maxWeight
	return p
}

func (p *Pet) SetWeight(weight int) *Pet {
	p.Weight = weight
	return p
}

func (p *Pet) SetDescription(description string) *Pet {
	p.Description = description
	return p
}

func (p *Pet) SetLifeSpan(lifeSpan int) *Pet {
	p.LifeSpan = lifeSpan
	return p
}

func (p *Pet) SetGeographicOrigin(geographicOrigin string) *Pet {
	p.GeographicOrigin = geographicOrigin
	return p
}

func (p *Pet) SetColor(color string) *Pet {
	p.Color = color
	return p
}

func (p *Pet) SetAge(age int) *Pet {
	p.Age = age
	return p
}

func (p *Pet) SetAgeEstimated(ageEstimated bool) *Pet {
	p.AgeEstimated = ageEstimated
	return p
}

// Build uses various fields of the Pet struct to calculate the average weight of the pet. If the minimum weight is greater than the maximum weight, an error is returned.
func (p *Pet) Build() (*Pet, error) {
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("minimum weight must be less than maximum weight")
	}

	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2

	return p, nil
}
