package dsa

import "fmt"

type Inventory struct {
	products    []*Product
	WarningQuan int
}

func NewInventory(warningQuan int) *Inventory {
	return &Inventory{
		products:    make([]*Product, 0),
		WarningQuan: warningQuan,
	}
}

func (i *Inventory) Add(p *Product) {
	i.products = append(i.products, p)
}

func (i *Inventory) WarningProducts() []*Product {
	wp := make([]*Product, 0)

	for _, p := range i.products {
		if p.Quantity < int32(i.WarningQuan) {
			wp = append(wp, p)
		}
	}

	return wp
}

func (i *Inventory) UpdateQuan(ID string, newQuan int32) (bool, error) {
	found := false
	for _, p := range i.products {
		if p.ID == ID {
			p.Quantity = newQuan
			found = true
			break
		}
	}

	if found {
		return found, nil
	} else {
		return found, fmt.Errorf("cannot find matching product")
	}

}

func (i *Inventory) FindId(productId string) *Product {
	for _, p := range i.products {
		if p.ID == productId {
			return p
		}
	}
	return nil
}

func (i *Inventory) FindName(name string) *Product {
	for _, p := range i.products {
		if p.Name == name {
			return p
		}
	}

	return nil
}
