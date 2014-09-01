package gosr

type ServiceInterface interface {
	Verify( secret []byte , timeWindow int ) * Error
	CalculateSignature( secret []byte )( string, * Error)
	GetSignature() string
	GetUser() string


}
