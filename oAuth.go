package gowoocommerce

import (
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type oAuth struct {
	url            string
	consumerKey    string
	comsumerSecret string
	apiVersion     string
	method         string
	params         map[string]string
	timestamp      string
}

func NewOAuth(url string, consumerKey string, consumerSecret string, apiVersion string, method string, params map[string]string, timestamp string) oAuth {
	oAuth := oAuth{
		url, consumerKey, consumerSecret, apiVersion, method, params, timestamp,
	}

	return oAuth
}

func (this oAuth) encode(value string) string {
	value = url.QueryEscape(value)
	value = strings.Replace(value, "+", " ", -1)
	value = strings.Replace(value, "%7E", "~", -1)
	return value
}

func (this oAuth) encodeArr(value []string) []string {
	result := []string{}
	for _, v := range value {
		result = append(result, this.encode(v))
	}
	return result
}

func (this oAuth) normalize(param map[string]string) map[string]string {
	normalized := make(map[string]string)
	for key, val := range param {
		normalized[this.encode(key)] = this.encode(val)
	}
	return normalized
}

func (this oAuth) sort(params map[string]string) map[string]string {
	var keys []string
	for key, _ := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	sorted := make(map[string]string)
	for _, key := range keys {
		sorted[key] = params[key]
	}
	return sorted
}

func (this oAuth) generateOauthSignature(params map[string]string) string {
	params = this.normalize(params)
	params = this.sort(params)
	paramString := strings.Join(this.joinWithEqualsSign(params), "&")
	signatureBaseString := this.method + "&" + url.QueryEscape(this.url) + "&" + url.QueryEscape(paramString)

	return getSha256(this.comsumerSecret+"&", signatureBaseString)
}
func (this oAuth) joinWithEqualsSign(params map[string]string) []string {
	var result []string
	for key, val := range params {
		result = append(result, key+"="+val)
	}
	return result
}

func (this oAuth) GetParameters() map[string]string {
	params := map[string]string{
		"oauth_consumer_key":     this.consumerKey,
		"oauth_timestamp":        this.timestamp,
		"oauth_nonce":            strconv.FormatInt(time.Now().Unix(), 10),
		"oauth_signature_method": "HMAC-SHA256",
	}
	params["oauth_signature"] = this.generateOauthSignature(params)

	return params
}
