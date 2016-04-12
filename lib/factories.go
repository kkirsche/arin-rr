package irr

// NewARINRouteEntry represents the data necessary to communicate an update to ARIN.
func NewARINRouteEntry(email Email, entry RouteRegistryEntry) *ArinRouteEntry {
	return &ArinRouteEntry{
		Email: email,
		Entry: entry,
	}
}

// Flatten removes the layers withihn the structure.
func (e *ArinRouteEntry) Flatten() *ArinFlatRouteEntry {
	return &ArinFlatRouteEntry{
		To:           e.Email.To,
		From:         e.Email.From,
		Subject:      e.Email.Subject,
		SMTPServer:   e.Email.SMTPServer,
		Route:        e.Entry.Route,
		Description:  e.Entry.Description,
		ASN:          e.Entry.ASN,
		NotifyEmail:  e.Entry.NotifyEmail,
		MaintainedBy: e.Entry.MaintainedBy,
		ChangedEmail: e.Entry.ChangedEmail,
		Source:       e.Entry.Source,
	}
}

// NewEmail returns a new instance of an Email
func NewEmail(from, to, subject, smtpServer string) Email {
	return Email{
		From:       from,
		To:         to,
		Subject:    subject,
		SMTPServer: smtpServer,
	}
}

// NewRouteRegistryEntry returns a new ARIN route registry entry
func NewRouteRegistryEntry(route, desc string, asn int, holes, memberOf,
	inject, aggrMtd, aggrBndry, exportComps, components, remarks, notifyEmail,
	maintBy, changeEmail, source string) RouteRegistryEntry {
	return RouteRegistryEntry{
		Route:        route,
		Description:  desc,
		ASN:          asn,
		Holes:        holes,
		MemberOf:     memberOf,
		Inject:       inject,
		AggrMtd:      aggrMtd,
		AggrBndry:    aggrBndry,
		ExportComps:  exportComps,
		Components:   components,
		Remarks:      remarks,
		NotifyEmail:  notifyEmail,
		MaintainedBy: maintBy,
		ChangedEmail: changeEmail,
		Source:       source,
	}
}
