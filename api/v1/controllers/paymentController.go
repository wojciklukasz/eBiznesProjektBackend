package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetPaymentRouting(e *echo.Group) {
	g := e.Group("/payment")
	g.POST("/:id", func(c echo.Context) error {
		id := c.Param("id")
		res, err := GeneratePaymentLink(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to handle payment"})
		}
		return c.JSON(http.StatusOK, map[string]string{"clientSecret": res})
	})
}

func CalculateOrderCost(orderID string) (int64, error) {
	total := 1400
	order, err := GetOrderById(orderID)
	if err != nil {
		return -1, err
	}

	itemOrders := order.Items
	itemOrdersArr := strings.Split(itemOrders, ",")

	for it, item := range itemOrdersArr {
		if it == len(itemOrdersArr)-1 {
			break
		}

		id, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			return -1, err
		}

		itemOrder, err := GetItemOrderByID(int(id))
		if err != nil {
			return -1, err
		}

		product, err := GetProductByID(itemOrder.ProductID)
		if err != nil {
			return -1, err
		}

		total += int(product.Price*100) * itemOrder.Count
	}

	return int64(total), nil
}

func GeneratePaymentLink(orderID string) (string, error) {
	amount, err := CalculateOrderCost(orderID)
	if err != nil {
		return "", err
	}

	stripe.Key = os.Getenv("STRIPE_SECRET")
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(string(stripe.CurrencyPLN)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return "", err
	}

	return pi.ClientSecret, nil
}
