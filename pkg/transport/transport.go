package transport

// Server server
type Server interface {
	Endpoint() ([]string, error)
	Start() error
	Stop() error
}
