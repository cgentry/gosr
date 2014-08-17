package http

type HttpInterface interface {
	CalculateSignature( secret string ) string
	Verify( secret []byte , timeWindow int ) error
	GetUser() string
	GetSignature() string
	GetContentType() string
	GetContentSignature() string
}
