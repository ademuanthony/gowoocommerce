package woocommerce

type AuthTokener interface {
	AuthToken(key, secret string) string
}

