package interfaces

type IApiServer interface {
	Start(port uint16)
	Stop()
}
