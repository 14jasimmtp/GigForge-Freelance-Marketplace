package paypal

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pkg/domain"
	"github.com/plutov/paypal/v4"
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
				Value:    fmt.Sprintf("%.2f", order.FreelancerFee),
			},
			Payee: &paypal.PayeeForOrders{
				EmailAddress: freelancerEmail,
			},
			ReferenceID: "1",
		},
		{
			Amount: &paypal.PurchaseUnitAmount{
				Currency: "USD",
				Value:    fmt.Sprintf("%.2f", order.MarketplaceFee),
			},
			Payee: &paypal.PayeeForOrders{
				EmailAddress: "sb-aeymg30598091@business.example.com",
			},
			ReferenceID: "2",
		},
	}

	// Debug information to print the purchase units
	fmt.Printf("Purchase Units: %+v\n", purchaseUnits)
	var err error
	paypalClient, err = paypal.NewClient("ARh_tJkDPzL6OkIUdjKEMyxg8t_ZKiZ_sYm7Sapv4x9NTsPxjQqKAGyEUcpsyT_7_MdZeYUTM40o7oLl", "EDAsi6PgoWAgCqiEO-hhki2mmICwFM7q0z8K1fMylnQdkrH8jvUQyFPPc24Ie8j_Vmb6UCYOcUYa2ToT", paypal.APIBaseSandBox)
	if err != nil {
		log.Fatal(err)
	}
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
	merchantIDs := []string{freelancerEmail, "sb-aeymg30598091@business.example.com"}

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
