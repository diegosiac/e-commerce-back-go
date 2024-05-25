package payment

import (
	"os"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/charge"
)

func Init() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}

func CreateCharge(amount int64, currency, description string) (*stripe.Charge, error) {
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(amount),
		Currency:    stripe.String(currency),
		Description: stripe.String(description),
	}
	return charge.New(params)
}
