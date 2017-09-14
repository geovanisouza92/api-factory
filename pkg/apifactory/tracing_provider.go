package apifactory

type TracingProvider interface {
	Configure(*ProviderConfig) error
	Send(TracingEvent)
}

type TracingEvent struct {
	RequestID  string // User? %u
	Scheme     string
	Host       string // %h
	Method     string // %r
	RequestURI string // %r
	Proto      string // %r
	RemoteAddr string // %l
	Status     int    // %s
	Bytes      int    // %b
	Duration   int64
	// TODO: Time // %t
}
