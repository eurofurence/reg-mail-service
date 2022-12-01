package mailctl

import (
	"context"
	"encoding/json"
	"github.com/eurofurence/reg-mail-service/docs"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/mail"
	"net/url"
	"reflect"
	"testing"
)

func tstCreateValidMail() mail.MailSendDto {
	return mail.MailSendDto{
		CommonID:  "test",
		Lang:      "de-DE",
		To:        []string{"test@example.com"},
		Cc:        nil,
		Bcc:       []string{"admin@example.com"},
		Variables: map[string]string{"nickname": "test_nick"},
	}
}

func TestValidateSuccess(t *testing.T) {
	docs.Description("a valid mail reports no validation errors")
	a := tstCreateValidMail()
	expected := url.Values{}
	performValidationTest(t, &a, expected)
}

func performValidationTest(t *testing.T, a *mail.MailSendDto, expectedErrors url.Values) {
	actualErrors := validate(context.TODO(), a)

	prettyPrintedActualErrors, _ := json.MarshalIndent(actualErrors, "", "  ")
	prettyPrintedExpectedErrors, _ := json.MarshalIndent(expectedErrors, "", "  ")

	if !reflect.DeepEqual(actualErrors, expectedErrors) {
		t.Errorf("Errors were not as expected.\nActual:\n%v\nExpected:\n%v\n", string(prettyPrintedActualErrors), string(prettyPrintedExpectedErrors))
	}
}
