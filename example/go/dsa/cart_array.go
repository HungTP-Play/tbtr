package dsa

type Product struct {
	ID       string
	Name     string
	Price    float32
	Quantity int32
}

type Cart struct {
	arr []*Product
}

func NewCart() *Cart {
	return &Cart{
		arr: make([]*Product, 0),
	}
}

func (c *Cart) Get(index int) Product {
	return *c.arr[index]
}

func (c *Cart) Set(index int, p *Product) {
	c.arr[index] = p
}

func (c *Cart) Add(p *Product) {
	c.arr = append(c.arr, p)
}

func (c *Cart) Delete(index int) {
	for i := index; i < len(c.arr)-1; i++ {
		c.arr[index] = c.arr[index+1]
	}

	c.arr[len(c.arr)-1] = nil
}

func (c *Cart) Filter(predicate func(p *Product, index int) bool) []*Product {
	var filteredProducts []*Product = make([]*Product, 0)
	for i, val := range c.arr {
		ok := predicate(val, i)
		if ok {
			filteredProducts = append(filteredProducts, val)
		}
	}

	return filteredProducts
}

func (c *Cart) ForEach(iterateFunc func(p *Product, index int)) {
	for i, val := range c.arr {
		iterateFunc(val, i)
	}
}

func (c *Cart) Find(productId string) *Product {
	for _, val := range c.arr {
		if val.ID == productId {
			return val
		}
	}

	return nil
}

func (c *Cart) Total() float64 {
	sum := float64(0)
	for _, val := range c.arr {
		sum += (float64(val.Price) * float64(val.Quantity))
	}
	return sum
}
