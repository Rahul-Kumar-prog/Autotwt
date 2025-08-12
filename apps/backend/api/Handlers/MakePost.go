package handlers

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type TwitterConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

type TweetRequest struct {
	Text string `json:"text"`
}

type TweetResponse struct {
	Data struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
	Error []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

func generateNonce() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err) // This should never happen with crypto/rand
	}
	return base64.StdEncoding.EncodeToString(b)
}

func MakePost(Msg string) error {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get Twitter OAuth credentials from environment variables
	consumerKey := os.Getenv("X_CONSUMER_KEY")
	consumerSecret := os.Getenv("X_CONSUMER_SECRET")
	accessToken := os.Getenv("X_ACCESS_TOKEN")
	accessSecret := os.Getenv("X_ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		return fmt.Errorf("missing Twitter API credentials. Please set the following environment variables in your .env file: X_CONSUMER_KEY, X_CONSUMER_SECRET, X_ACCESS_TOKEN, X_ACCESS_SECRET")
	}

	config := TwitterConfig{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		AccessToken:    accessToken,
		AccessSecret:   accessSecret,
	}

	apiURL := "https://api.twitter.com/2/tweets"

	tweetReq := TweetRequest{Text: Msg}
	jsonData, err := json.Marshal(tweetReq)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := generateNonce()

	oauthParams := map[string]string{
		"oauth_consumer_key":     config.ConsumerKey,
		"oauth_nonce":            nonce,
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        timestamp,
		"oauth_token":            config.AccessToken,
		"oauth_version":          "1.0",
	}

	signature := generateOauthSignature("POST", apiURL, oauthParams, config.ConsumerSecret, config.AccessSecret)
	oauthParams["oauth_signature"] = signature

	var authPairs []string
	for k, v := range oauthParams {
		authPairs = append(authPairs, fmt.Sprintf(`%s="%s"`, k, url.QueryEscape(v)))
	}
	authHeader := "OAuth " + strings.Join(authPairs, ", ")

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	// Parse response
	var tweetResp TweetResponse
	if err := json.Unmarshal(body, &tweetResp); err != nil {
		return fmt.Errorf("error parsing response: %v", err)
	}

	// Check for errors in response
	if len(tweetResp.Error) > 0 {
		return fmt.Errorf("Twitter API error: %s", tweetResp.Error[0].Message)
	}

	// Log successful response
	fmt.Printf("Tweet posted successfully! Tweet ID: %s\n", tweetResp.Data.ID)
	return nil
}

func generateOauthSignature(method, baseURL string, params map[string]string, consumerSecret, tokenSecret string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var paramPairs []string
	for _, k := range keys {
		paramPairs = append(paramPairs, fmt.Sprintf("%s=%s", url.QueryEscape(k), url.QueryEscape(params[k])))
	}
	paramString := strings.Join(paramPairs, "&")

	// Create signature base string
	signatureBaseString := fmt.Sprintf("%s&%s&%s",
		method,
		url.QueryEscape(baseURL),
		url.QueryEscape(paramString))

	// Create signing key
	signingKey := fmt.Sprintf("%s&%s",
		url.QueryEscape(consumerSecret),
		url.QueryEscape(tokenSecret))

	// Generate signature
	h := hmac.New(sha1.New, []byte(signingKey))
	h.Write([]byte(signatureBaseString))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}
