package gowoocommerce

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Version int

func (version Version) ToString() string {
	switch version {
	case Versions.Legacy:
		return "1"
	case Versions.Version1:
		return "2"
	case Versions.Version2:
		return "3"
	case Versions.ThirdPartyPlugins:
		return "99"
	}
	return ""
}

var Versions = struct {
	Unknown           Version
	Legacy            Version
	Version1          Version
	Version2          Version
	ThirdPartyPlugins Version
}{
	Unknown:           0,
	Legacy:            1,
	Version1:          2,
	Version2:          3,
	ThirdPartyPlugins: 99,
}

type RequestMethod int

func (this RequestMethod) ToString() string {
	switch this {
	case 1:
		return "HEAD"
	case 2:
		return "GET"
	case 3:
		return "POST"
	case 4:
		return "PUT"
	case 5:
		return "PATCH"
	case 6:
		return "DELETE"
	}
	return "GET"
}

var RequestMethods = struct {
	HEAD   RequestMethod
	GET    RequestMethod
	POST   RequestMethod
	PUT    RequestMethod
	PATCH  RequestMethod
	DELETE RequestMethod
}{
	HEAD:   1,
	GET:    2,
	POST:   3,
	PUT:    4,
	PATCH:  5,
	DELETE: 6,
}

type restfulClient struct {
	wc_url    string
	wc_key    string
	wc_secret string

	Version Version
}

func NewRestfulClient(url, key, secret string) (client *restfulClient) {
	if url == "" {
		panic("Please use a valid WooCommerce Restful API url.")
	}

	client = &restfulClient{}

	urlLower := strings.TrimRight(strings.ToLower(strings.TrimSpace(url)), "/")
	if strings.HasSuffix(urlLower, "wc-api/v1") || strings.HasSuffix(urlLower, "wc-api/v2") || strings.HasSuffix(urlLower, "wc-api/v3") {
		client.Version = Versions.Legacy
	} else if strings.HasSuffix(urlLower, "wp-json/wc/v1") {
		client.Version = Versions.Version1
	} else if strings.HasSuffix(urlLower, "wp-json/wc/v2") {
		client.Version = Versions.Version2
	} else if strings.HasSuffix(urlLower, "wp-json/wc-") {
		client.Version = Versions.ThirdPartyPlugins
	} else {
		panic("Unknow WooCommerce Restful API version.")
	}

	if strings.HasSuffix(url, "/") {
		client.wc_url = url
	} else {
		client.wc_url = url + "/"
	}

	client.wc_key = key
	client.wc_secret = secret

	return client
}

func (this restfulClient) IsLegacyVersion() bool {
	return this.Version == Versions.Legacy
}

func (this restfulClient) Url() string {
	return this.wc_url
}

//Make restful call
func (this restfulClient) request(endpoint string, method RequestMethod, body interface{}, params map[string]string) (string, error) {
	var respBody string

	if params == nil {
		params = make(map[string]string)
	}
	if _, ok := params["consumer_key"]; !ok {
		params["consumer_key"] = this.wc_key
	}
	if _, ok := params["consumer_secret"]; !ok {
		params["consumer_secret"] = this.wc_secret
	}

	fullUrl := this.wc_url + this.getOAuthEndPoint(endpoint, method.ToString(), params)

	//println(fullUrl)

	switch method {
	case RequestMethods.GET:
		resp, err := http.Get(fullUrl)
		if err != nil {
			return respBody, err
		}

		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		respBody = string(bodyBytes)

		return respBody, nil
	}

	panic("Request Method Not Supported")
}

func (this restfulClient) getOAuthEndPoint(endpoint string, method string, params map[string]string) string {
	auth := NewOAuth(this.wc_url+endpoint, this.wc_key, this.wc_secret, this.Version.ToString(),
		method, params, GetUnixTime(false))
	parameter := auth.GetParameters()

	paramstr := ""
	for key, val := range parameter {
		paramstr += key + "=" + val + "&"
	}

	return endpoint + "?" + strings.TrimRight(paramstr, "&")

}
