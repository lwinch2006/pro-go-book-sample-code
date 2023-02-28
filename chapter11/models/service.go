package models

type ServiceModel struct {
	Description    string
	DurationMonths int
	MonthlyFee     float64
}

func NewService(description string, durationMonths int, monthlyFee float64) *ServiceModel {
	return &ServiceModel{description, durationMonths, monthlyFee}
}

func (service *ServiceModel) GetName() string {
	return service.Description
}

func (service *ServiceModel) GetCost(annual bool) float64 {
	if annual {
		return service.MonthlyFee * float64(service.DurationMonths)
	}

	return service.MonthlyFee
}
