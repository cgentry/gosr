package gosr

type ServiceInterface interface {
	Verify( secret []byte , timeWindow int ) error
	CalculateSignature( secret []byte )( string, error)
	GetSignature() string
	GetUser() string


}
