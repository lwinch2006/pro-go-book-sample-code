package store

const defaultTaxRate = 0.2
const minThreshold = 10

var categoryMaxPrices = map[string]float64{
	"Milk":  25,
	"Bread": 50,
}

func init() {
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + price*defaultTaxRate
	}
}

type taxRate struct {
	rate, threshold float64
}

func newTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}

	if threshold < minThreshold {
		threshold = minThreshold
	}

	return &taxRate{rate, threshold}
}

func (t *taxRate) calcTax(p *Product) (totalPrice float64) {
	totalPrice = p.price

	if p.price > t.threshold {
		totalPrice += p.price * t.rate
	}

	if maxPrice, found := categoryMaxPrices[p.Name]; found && totalPrice > maxPrice {
		totalPrice = maxPrice
	}

	return
}
