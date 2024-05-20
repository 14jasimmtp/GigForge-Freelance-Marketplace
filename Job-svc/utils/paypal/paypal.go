package paypal

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateOrder(accessToken, itemName, itemPrice, currency, auth_assertion_header string) (string, error) {
	url := "https://api-m.sandbox.paypal.com/v2/checkout/orders"

	itemTotal := itemPrice
	purchaseUnits := []map[string]interface{}{
		{
			"amount": map[string]interface{}{
				"currency_code": currency,
				"value":         itemTotal,
			},
			"items": []map[string]interface{}{
				{
					"name": itemName,
					"unit_amount": map[string]interface{}{
						"currency_code": currency,
						"value":         itemPrice,
					},
				},
			},
		},
	}

	orderData := map[string]interface{}{
		"intent":         "CAPTURE",
		"purchase_units": purchaseUnits,
	}

	jsonData, err := json.Marshal(orderData)
	if err != nil {
		fmt.Println("Error marshaling order data:", err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("PayPal - Auth - Assertion", auth_assertion_header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return "", err
	}

	orderId := data["id"].(string)
	return orderId, nil
}

func CapturePayment(accessToken, orderId string) error {
	url := "https://api-m.sandbox.paypal.com/v2/checkout/orders/" + orderId + "/capture"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Payment captured successfully.")
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}

	fmt.Println("Error capturing payment:", string(body))
	return fmt.Errorf("payment capture failed")
}

func GetAuthAssertionValue(clientID, sellerPayerID string) string {
	header := map[string]string{
		"alg": "none",
	}
	encodedHeader := base64URLEncode(header)
	payload := map[string]string{
		"iss":      clientID,
		"payer_id": sellerPayerID,
	}
	encodedPayload := base64URLEncode(payload)
	return fmt.Sprintf("%s.%s.", encodedHeader, encodedPayload)
}

func base64URLEncode(data map[string]string) string {
	jsonData, _ := json.Marshal(data)
	encoded := base64.RawURLEncoding.EncodeToString(jsonData)
	return encoded
}
