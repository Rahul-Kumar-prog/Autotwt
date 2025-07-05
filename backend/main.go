package main

import (
	"bufio"
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

// TwitterConfig holds the OAuth 1.0a credentials
type TwitterConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

// TweetRequest represents the structure for posting a tweet
type TweetRequest struct {
	Text string `json:"text"`
}

// TweetResponse represents the response from Twitter API
type TweetResponse struct {
	Data struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

// generateNonce creates a random nonce for OAuth
func generateNonce() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

// generateOAuthSignature creates the OAuth 1.0a signature
func generateOAuthSignature(method, baseURL string, params map[string]string, consumerSecret, tokenSecret string) string {
	// Sort parameters
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Create parameter string
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

// postToTwitter posts a tweet to Twitter using OAuth 1.0a
func postToTwitter(content string, config TwitterConfig) error {
	apiURL := "https://api.twitter.com/2/tweets"

	// Create request body
	tweetReq := TweetRequest{Text: content}
	jsonData, err := json.Marshal(tweetReq)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	// OAuth parameters
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

	// Generate signature
	signature := generateOAuthSignature("POST", apiURL, oauthParams, config.ConsumerSecret, config.AccessSecret)
	oauthParams["oauth_signature"] = signature

	// Create Authorization header
	var authPairs []string
	for k, v := range oauthParams {
		authPairs = append(authPairs, fmt.Sprintf(`%s="%s"`, k, url.QueryEscape(v)))
	}
	authHeader := "OAuth " + strings.Join(authPairs, ", ")

	// Create HTTP request
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

	// Check for errors
	if resp.StatusCode != 201 {
		if len(tweetResp.Errors) > 0 {
			return fmt.Errorf("Twitter API error: %s", tweetResp.Errors[0].Message)
		}
		return fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
	}

	return nil
}

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get Twitter OAuth credentials from environment variables
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		fmt.Println("Please set the following environment variables in your .env file:")
		fmt.Println("TWITTER_CONSUMER_KEY=your_consumer_key")
		fmt.Println("TWITTER_CONSUMER_SECRET=your_consumer_secret")
		fmt.Println("TWITTER_ACCESS_TOKEN=your_access_token")
		fmt.Println("TWITTER_ACCESS_SECRET=your_access_secret")
		return
	}

	config := TwitterConfig{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		AccessToken:    accessToken,
		AccessSecret:   accessSecret,
	}

	scanner := bufio.NewScanner(os.Stdin)

	// Ask for the post content
	fmt.Print("What do you want to post? ")
	scanner.Scan()
	postContent := scanner.Text()

	// Ask for confirmation
	fmt.Print("Do you want to post this? (yes/no): ")
	scanner.Scan()
	confirmation := strings.ToLower(strings.TrimSpace(scanner.Text()))

	// Check confirmation and act accordingly
	if confirmation == "yes" || confirmation == "y" {
		// Actually post to Twitter
		fmt.Println("Posting to Twitter...")
		err := postToTwitter(postContent, config)
		if err != nil {
			fmt.Printf("Failed to post: %v\n", err)
		} else {
			fmt.Println("Successfully posted to Twitter!")
		}
	} else {
		fmt.Println("Posting cancelled")
	}
}
