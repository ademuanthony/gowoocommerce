package woocommerce

type client struct {
	key string
	secret string

}

func NewClient(key, secret string) client {
	return client{key, secret}
}

func (this client) authToken() string {
	
}