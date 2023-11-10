package util

const (
	USD = "USD"
	EUR = "EUR"
	TRY = "TRY"
)

// IsSupportedCurrency returns true if given currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, TRY:
		return true
	}
	return false
}
