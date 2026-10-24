package payment

type PaymentConfig struct {
	DefaultCurrency         string   `json:"default_currency"`          // Default currency for payments
	XApiKey                 string   `json:"-"`                         // Api key for external payment service
	SupportedCurrencies     []string `json:"supported_currencies"`      // List of supported currencies for payments
	SupportedPaymentMethods []string `json:"supported_payment_methods"` // List of supported payment methods
}
