package q4

import "fmt"

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	// Verifica se o preço base é zero ou negativo
	if basePrice <= 0 {
		return 0, fmt.Errorf("O preço base deve ser maior que zero")
	}
	if taxCode < 1 || taxCode > 3 {
		return 0, fmt.Errorf("A opção escolhida é inválida")
	}
	var TaxRate float64
	switch taxCode {
	case 1:
		TaxRate = 0.05
	case 2:
		TaxRate = 0.1
	case 3:
		TaxRate = 0.15
	}
	var freightRate float64
	switch state {
	case "SP":
		freightRate = 0.1
	case "RJ":
		freightRate = 0.15
	case "MG":
		freightRate = 0.2
	case "ES":
		freightRate = 0.25
	default:
		freightRate = 0.3
	}

	// Calcula o preço final
	finalPrice := basePrice + basePrice*TaxRate + basePrice*freightRate

	return finalPrice, nil
}
