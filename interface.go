package perfmon

type Service interface {
	Start()
	PushData(name string, value int)
	Shutdown()
}
