package gosr

type ServiceInterface interface {
	CalculateSignature( secret string ) string
	Verify( secret []byte , timeWindow int ) error
	GetUser() string
	GetSignature() string
	GetContentType() string
	GetContentSignature() string
}
