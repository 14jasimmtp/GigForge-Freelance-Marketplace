package paypal

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GenerateAccessToken() (string, error) {
	paypalClientID := "ARh_tJkDPzL6OkIUdjKEMyxg8t_ZKiZ_sYm7Sapv4x9NTsPxjQqKAGyEUcpsyT_7_MdZeYUTM40o7oLl"
	paypalClientSecret := "EDAsi6PgoWAgCqiEO-hhki2mmICwFM7q0z8K1fMylnQdkrH8jvUQyFPPc24Ie8j_Vmb6UCYOcUYa2ToT"

	url := "https://api-m.sandbox.paypal.com/v1/oauth2/token"

	// Encode client ID and secret using Base64
	auth := paypalClientID + ":" + paypalClientSecret
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	payload := "grant_type=client_credentials"

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+encodedAuth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()
	data := make(map[string]interface{})
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}
	accessToken := data["access_token"].(string)

	fmt.Println("Response:", accessToken)
	return string(accessToken), nil
}

func OnboardFreelancer(user_id, access_token string) (string, error) {
	url := "https://api-m.sandbox.paypal.com/v2/customer/partner-referrals"
	trackingID := user_id

	jsonData := []byte(fmt.Sprintf(`{
        "tracking_id": "%s",
        "operations": [{
            "operation": "API_INTEGRATION",
            "api_integration_preference": {
                "rest_api_integration": {
                    "integration_method": "PAYPAL",
                    "integration_type": "THIRD_PARTY",
                    "third_party_details": {
                        "features": [
                            "PAYMENT",
                            "REFUND"
                        ]
                    }
                }
            }
        }],
        "products": [
            "EXPRESS_CHECKOUT"
        ],
        "legal_consents": [{
            "type": "SHARE_DATA_CONSENT",
            "granted": true
        }]
    }`, trackingID))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ access_token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "",err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "",err
	}

	return string(body),nil
}
