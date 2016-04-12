package irr

// ArinRouteEntry constitutes an entry into ARIN's Route Registry
type ArinRouteEntry struct {
	Email Email
	Entry RouteRegistryEntry
}

// ArinFlatRouteEntry is the flat version of ArinRouteEntry
type ArinFlatRouteEntry struct {
	From         string
	To           string
	Subject      string
	SMTPServer   string
	Route        string
	Description  string
	ASN          int
	NotifyEmail  string
	MaintainedBy string
	ChangedEmail string
	Source       string
}

// RouteRegistryEntry is the Route Registry entry itself
type RouteRegistryEntry struct {
	Route        string
	Description  string
	ASN          int
	Holes        string
	MemberOf     string
	Inject       string
	AggrMtd      string
	AggrBndry    string
	ExportComps  string
	Components   string
	Remarks      string
	NotifyEmail  string
	MaintainedBy string
	ChangedEmail string
	Source       string
}

// Email is the data for sending / receiving the email.
type Email struct {
	From       string
	To         string
	Subject    string
	SMTPServer string
}

// Logger represents a logging object for use by IRR.
type Logger struct {
	Verbose bool
}
