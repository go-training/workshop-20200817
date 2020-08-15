package bmw

// Car ...
type Car struct {
	Price float64
	total float64
}

// Process ...
func (c *Car) Process() {
	c.total = c.Price * 0.5
}

// Total ...
func (c *Car) Total() float64 {
	return c.total
}

// New car
func New(price float64) *Car {
	return &Car{
		Price: price,
	}
}
