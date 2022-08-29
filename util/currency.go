package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	YEN = "YEN"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, YEN, CAD:
		return true
	}
	return false
}
