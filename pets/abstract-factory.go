package pets

import (
	"errors"
	"fmt"
	"rahuljsaliaan/go-breeders/models"
)

type AnimalInterface interface {
	Show() string
}

// Dog Factory
type DogFromFactory struct {
	Pet *models.Dog
}

func (dff *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", dff.Pet.Breed.Breed)
}

// Cat Factory
type CatFromFactory struct {
	Pet *models.Cat
}

func (dff *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", dff.Pet.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
}

// Dog Abstract Factory
type DogAbstractFactory struct{}

func (df *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

// Cat Abstract Factory
type CatAbstractFactory struct{}

func (df *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

// Returns the abstract factory of that species
func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil
	default:
		return nil, errors.New("invalid species supplied")
	}
}
