package store

type BundleOffer struct {
	*Product
	*SpecialDeal
}

func NewBundleOffer(specialDealName string, specialDealPrice float64, productName, productCategory string, productPrice float64) *BundleOffer {
	newSpecialDeal := NewSpeciaDeal(specialDealName, specialDealPrice, productName, productCategory, productPrice)
	return &BundleOffer{newSpecialDeal.Product, newSpecialDeal}
}
