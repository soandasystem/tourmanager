package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type FlowAPI struct {
	SecretKey string
	APIURL    string
	APIKey    string
}

func NewFlowAPI(apiKey, secretKey, apiURL string) *FlowAPI {
	return &FlowAPI{
		APIKey:    apiKey,
		SecretKey: secretKey,
		APIURL:    apiURL,
	}
}

func (api *FlowAPI) SetAPIKey(apiKey string) {
	api.APIKey = apiKey
}

func (api *FlowAPI) SetSecretKey(secretKey string) {
	api.SecretKey = secretKey
}

func (api *FlowAPI) SetAPIURL(apiURL string) {
	api.APIURL = apiURL
}

func (api *FlowAPI) Send(service string, params map[string]string, method string) (map[string]interface{}, error) {
	method = strings.ToUpper(method)
	if method == "" {
		method = "GET"
	}
	requestURL := fmt.Sprintf("%s/%s", api.APIURL, service)

	if params == nil {
		params = make(map[string]string)
	}
	// Copy params to avoid modifying the original map
	requestParams := make(map[string]string)
	for k, v := range params {
		requestParams[k] = v
	}
	requestParams["apiKey"] = api.APIKey

	data := api.getPack(requestParams, method)
	sign := api.sign(requestParams)

	var responseData []byte
	var httpCode int
	var err error

	if method == "GET" {
		responseData, httpCode, err = api.httpGet(requestURL, data, sign)
	} else {
		responseData, httpCode, err = api.httpPost(requestURL, data, sign)
	}

	if err != nil {
		return nil, err
	}

	var body map[string]interface{}
	if err := json.Unmarshal(responseData, &body); err != nil {
		return nil, fmt.Errorf("error decoding response: %v. Output: %s", err, string(responseData))
	}

	if httpCode == 200 {
		return body, nil
	} else if httpCode == 400 || httpCode == 401 {
		message := "Unknown error"
		code := float64(0)
		if msg, ok := body["message"].(string); ok {
			message = msg
		}
		if c, ok := body["code"].(float64); ok {
			code = c
		}
		return nil, fmt.Errorf("error %v: %s", code, message)
	}

	return nil, fmt.Errorf("unexpected error occurred. HTTP_CODE: %d", httpCode)
}

func (api *FlowAPI) getPack(params map[string]string, method string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var data strings.Builder
	for _, k := range keys {
		if method == "GET" {
			// Using QueryEscape to emulate PHP's rawurlencode behavior
			data.WriteString("&" + url.QueryEscape(k) + "=" + url.QueryEscape(params[k]))
		} else {
			data.WriteString("&" + k + "=" + params[k])
		}
	}

	res := data.String()
	if len(res) > 0 {
		return res[1:]
	}
	return res
}

func (api *FlowAPI) sign(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var toSign strings.Builder
	for _, k := range keys {
		toSign.WriteString("&" + k + "=" + params[k])
	}

	signStr := toSign.String()
	if len(signStr) > 0 {
		signStr = signStr[1:]
	}

	h := hmac.New(sha256.New, []byte(api.SecretKey))
	h.Write([]byte(signStr))
	return hex.EncodeToString(h.Sum(nil))
}

func (api *FlowAPI) httpGet(reqURL, data, sign string) ([]byte, int, error) {
	fullURL := fmt.Sprintf("%s?%s&s=%s", reqURL, data, sign)
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}

func (api *FlowAPI) httpPost(reqURL, data, sign string) ([]byte, int, error) {
	payload := fmt.Sprintf("%s&s=%s", data, sign)
	req, err := http.NewRequest("POST", reqURL, strings.NewReader(payload))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}
