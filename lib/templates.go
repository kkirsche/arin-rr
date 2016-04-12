package irr

const (
	// DefaultToEmail is the e-mail address to use when emailing ARIN about the RR
	DefaultToEmail = "rr@arin.net"
	// DefaultSubject is the subject line for the ARIN email
	DefaultSubject = "route"
	// DefaultSource The source line for the ARIN email
	DefaultSource = "ARIN"

	// ArinRouteEntryTemplate is the template to execute when sending a new route
	// to ARIN
	ArinRouteEntryTemplate = "To: {{ .To }}\r\n" +
		"From: {{ .From }}\r\n" +
		"Subject: {{ .Subject }}\r\n" +
		"\r\n" +
		"route: {{ .Route }}\r\n" +
		"descr: {{ .Description }}\r\n" +
		"origin: AS{{ .ASN }}\r\n" +
		"holes: {{.Holes }}\r\n" +
		"member-of: {{ .MemberOf }}\r\n" +
		"inject: {{ .Inject }}\r\n" +
		"aggr-mtd: {{ .AggrMtd }}\r\n" +
		"aggr-bndry: {{ .AggrBndry }}\r\n" +
		"export-comps: {{ .ExportComps }}\r\n" +
		"components: {{ .Components }}\r\n" +
		"remarks: {{ .Remarks }}\r\n" +
		"notify: {{ .NotifyEmail }}\r\n" +
		"mnt-by: {{ .MaintainedBy }}\r\n" +
		"changed: {{ .ChangedEmail }}\r\n" +
		"source: {{ .Source }}\r\n"
)
