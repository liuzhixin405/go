package registry

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName
	ServiceUpdateURL string
	HeartbeatURL     string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	ProtalService  = ServiceName("Portald")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}
type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
