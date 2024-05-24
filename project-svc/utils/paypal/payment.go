package paypal

import (
	"context"
	"errors"
	"fmt"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pkg/domain"
	"github.com/plutov/paypal/v4"
	"github.com/spf13/viper"
)

var paypalClient *paypal.Client

type Order struct {
	OrderID    string
	MerchantID []string
}

func CreatePayment(order *domain.ProjectOrders, freelancerEmail string) (*Order, error) {
	purchaseUnits := []paypal.PurchaseUnitRequest{
		{
			Amount: &paypal.PurchaseUnitAmount{
				Currency: "USD",
				Value:    fmt.Sprintf("%f", order.FreelancerFee),
			},
			Payee: &paypal.PayeeForOrders{
				EmailAddress: freelancerEmail,
			},
			ReferenceID: "1",
		},
		{
			Amount: &paypal.PurchaseUnitAmount{
				Currency: "USD",
				Value:    fmt.Sprintf("%f", order.MarketplaceFee),
			},
			Payee: &paypal.PayeeForOrders{
				EmailAddress: viper.GetString("MarketPlacePaypalEmail"),
			},
			ReferenceID: "2",
		},
	}

	// Debug information to print the purchase units
	fmt.Printf("Purchase Units: %+v\n", purchaseUnits)

	orders, err := paypalClient.CreateOrder(context.Background(), paypal.OrderIntentCapture, purchaseUnits, &paypal.CreateOrderPayer{
		Name: &paypal.CreateOrderPayerName{
			GivenName: "John",
			Surname:   "Doe",
		},
	}, &paypal.ApplicationContext{
		BrandName:          "GigForge Marketplace",
		LandingPage:        "BILLING",
		UserAction:         "PAY_NOW",
		ShippingPreference: "NO_SHIPPING",
	})

	if err != nil {
		fmt.Println("Error creating order:", err)
		return nil, errors.New(`failed to create order`)
	}

	fmt.Println(order)
	merchantIDs := []string{freelancerEmail, viper.GetString("MarketPlacePaypalEmail")}

	return &Order{OrderID: orders.ID, MerchantID: merchantIDs}, nil

	// Print full order response for debugging
	// fmt.Printf("Full Order Response: %+v\n", order)

	// // Check each part of the order
	// if order.Status != "CREATED" {
	//     fmt.Println("Order was not created successfully")
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
	//     return
	// }

	// if len(order.PurchaseUnits) == 0 {
	//     fmt.Println("No purchase units in order")
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
	//     return
	// }

	// purchaseUnit := order.PurchaseUnits[0]
	// if purchaseUnit.Amount == nil {
	//     fmt.Println("Purchase unit amount is nil")
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
	//     return
	// }

	// if purchaseUnit.PaymentInstruction == nil || len(purchaseUnit.PaymentInstruction.PlatformFees) == 0 {
	//     fmt.Println("Platform fees are not included")
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
	//     return
	// }

	// // Log the platform fee details
	// platformFee := purchaseUnit.PaymentInstruction.PlatformFees[0]
	// fmt.Printf("Platform Fee Details: %+v\n", platformFee)
}

func CapturePayment(PaymentID string) (string, error) {
	fmt.Println("capturing order...")

	captureResponse, err := paypalClient.CaptureOrder(context.Background(), PaymentID, paypal.CaptureOrderRequest{})
	if err != nil {
		fmt.Println(err)
		return "", errors.New(`failed to capture order`)
	}

	return captureResponse.Payer.Name.GivenName, nil
}
