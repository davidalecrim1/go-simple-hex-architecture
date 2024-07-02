package internal

import "github.com/google/uuid"

type Tax struct {
	ID  string
	Tax float64
}

func NewTax(id string) *Tax {
	return &Tax{ID: id}
}

func (t *Tax) Calculate(price float64, rate float64) {
	t.Tax = price + (price * rate)
}

// it's good practice for consumer (service)
// to define the producer interface (repository)
type TaxRepository interface {
	Save(id string, tax float64) error
}

type TaxService struct {
	repo TaxRepository
}

// good to know: passing by value (repo TaxRepository) not pointer because when you pass an interface by value,
// you are passing a copy of this pair of values, but both the original
// and the copy refer to the same underlying concrete implementation.
func NewTaxService(repo TaxRepository) *TaxService {
	return &TaxService{
		repo: repo,
	}
}

func (s *TaxService) Calculate(price float64, rate float64) (float64, error) {
	tax := NewTax(uuid.NewString())
	tax.Calculate(price, rate)
	return tax.Tax, s.repo.Save(tax.ID, tax.Tax)
}
