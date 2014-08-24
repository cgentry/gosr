package gosr

type ResponseInterface interface {
	Verify( secret []byte , timeWindow int ) error
	CalculateSignature( secret []byte ) string
	GetSignature() string
	Sign() *ResponseInterface
}
