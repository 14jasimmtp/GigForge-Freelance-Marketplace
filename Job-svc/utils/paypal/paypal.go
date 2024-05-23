package paypal

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/domain"
	"github.com/plutov/paypal/v4"
)

var paypalClient *paypal.Client

type Order struct {
	OrderID    string
	MerchantID []string
}

func CreatePayment(invoice *domain.Invoice, freelancerEmail string) (*Order, error) {
fmt.Println(invoice.Freelancer_fee,invoice.MarketPlace_fee)
	purchaseUnits := []paypal.PurchaseUnitRequest{
		{
			Amount: &paypal.PurchaseUnitAmount{
				Currency: "USD",
				Value:    fmt.Sprintf("%.2f", invoice.Freelancer_fee),
			},
			Payee: &paypal.PayeeForOrders{
				EmailAddress: freelancerEmail,
			},
			ReferenceID: "1",
		},
		{
			Amount: &paypal.PurchaseUnitAmount{
				Currency: "USD",
				Value:    fmt.Sprintf("%.2f", invoice.MarketPlace_fee),
			},
			Payee: &paypal.PayeeForOrders{
				EmailAddress: "sb-aeymg30598091@business.example.com",
			},
			ReferenceID: "2",
		},
	}

	// Debug information to print the purchase units
	fmt.Printf("Purchase Units: %+v\n", purchaseUnits)
	fmt.Println(paypalClient)
	var err error
	paypalClient, err = paypal.NewClient("ARh_tJkDPzL6OkIUdjKEMyxg8t_ZKiZ_sYm7Sapv4x9NTsPxjQqKAGyEUcpsyT_7_MdZeYUTM40o7oLl", "EDAsi6PgoWAgCqiEO-hhki2mmICwFM7q0z8K1fMylnQdkrH8jvUQyFPPc24Ie8j_Vmb6UCYOcUYa2ToT", paypal.APIBaseSandBox)
	if err != nil {
		log.Fatal(err)
	}
	order, err := paypalClient.CreateOrder(context.Background(), paypal.OrderIntentCapture, purchaseUnits, &paypal.CreateOrderPayer{
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

	// fmt.Println(order)
	merchantIDs := []string{freelancerEmail, "sb-aeymg30598091@business.example.com"}

	return &Order{OrderID: order.ID, MerchantID: merchantIDs}, nil
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

// func CreateOrder(accessToken, itemName, itemPrice, currency, auth_assertion_header string) (string, error) {
// 	url := "https://api-m.sandbox.paypal.com/v2/checkout/orders"

// 	itemTotal := itemPrice
// 	purchaseUnits := []map[string]interface{}{
// 		{
// 			"amount": map[string]interface{}{
// 				"currency_code": currency,
// 				"value":         itemTotal,
// 			},
// 			"items": []map[string]interface{}{
// 				{
// 					"name": itemName,
// 					"unit_amount": map[string]interface{}{
// 						"currency_code": currency,
// 						"value":         itemPrice,
// 					},
// 				},
// 			},
// 		},
// 	}

// 	orderData := map[string]interface{}{
// 		"intent":         "CAPTURE",
// 		"purchase_units": purchaseUnits,
// 	}

// 	jsonData, err := json.Marshal(orderData)
// 	if err != nil {
// 		fmt.Println("Error marshaling order data:", err)
// 		return "", err
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return "", err
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+accessToken)
// 	req.Header.Set("PayPal - Auth - Assertion", auth_assertion_header)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response:", err)
// 		return "", err
// 	}

// 	var data map[string]interface{}
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		fmt.Println("Error unmarshaling response:", err)
// 		return "", err
// 	}

// 	orderId := data["id"].(string)
// 	return orderId, nil
// }

// func CapturePayment(accessToken, orderId string) error {
// 	url := "https://api-m.sandbox.paypal.com/v2/checkout/orders/" + orderId + "/capture"

// 	req, err := http.NewRequest("POST", url, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return err
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+accessToken)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		fmt.Println("Payment captured successfully.")
// 		return nil
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response:", err)
// 		return err
// 	}

// 	fmt.Println("Error capturing payment:", string(body))
// 	return fmt.Errorf("payment capture failed")
// }

// func GetAuthAssertionValue(clientID, sellerPayerID string) string {
// 	header := map[string]string{
// 		"alg": "none",
// 	}
// 	encodedHeader := base64URLEncode(header)
// 	payload := map[string]string{
// 		"iss":      clientID,
// 		"payer_id": sellerPayerID,
// 	}
// 	encodedPayload := base64URLEncode(payload)
// 	return fmt.Sprintf("%s.%s.", encodedHeader, encodedPayload)
// }

// func base64URLEncode(data map[string]string) string {
// 	jsonData, _ := json.Marshal(data)
// 	encoded := base64.RawURLEncoding.EncodeToString(jsonData)
// 	return encoded
// }
