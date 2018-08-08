package gowoocommerce

type AuthTokener interface {
	AuthToken(key, secret string) string
}
