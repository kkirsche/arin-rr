package irr

import (
	"fmt"
	"html/template"
	"net/smtp"
)

// SendArinEmail is used to send an Email using a template to Arin containing the
// route update(s).
func SendArinEmail(entry *ArinFlatRouteEntry, t *template.Template) error {
	if entry == nil {
		return fmt.Errorf("Null entry provided. Cannot send email using nil entry.")
	}

	if t == nil {
		return fmt.Errorf("Null template provided. Cannot send email using nil template.")
	}

	client, err := smtp.Dial(entry.SMTPServer)
	if err != nil {
		return fmt.Errorf("Failed to connect to the SMTP server with error: %s", err.Error())
	}
	defer client.Close()
	defer client.Quit()

	err = client.Mail(entry.From)
	if err != nil {
		return fmt.Errorf("Failed to set 'From' address with error: %s.", err.Error())
	}

	err = client.Rcpt(entry.To)
	if err != nil {
		return fmt.Errorf("Failed to set 'To' address with error: %s.", err.Error())
	}

	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("Failed to set 'data' / 'message' with error: %s.", err.Error())
	}

	err = t.Execute(writer, entry)
	if err != nil {
		return fmt.Errorf("Failed to execute template with error: %s.", err.Error())
	}

	return nil
}
