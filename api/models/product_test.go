package models

import "testing"

func TestProductModel(t *testing.T) {

	p := &Product{}
	p.Name = ""
	p.Price = 1600.0
	p.Quantity = 10
	p.Status = ProductStatus_Available

	expected := "product.name can't be empty"



	if err := p.Validate(); err.Error() != expected {
		t.Errorf("expected: %s, got: %s", expected, err.Error())
	}

}