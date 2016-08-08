package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*BuildWebhook build webhook

swagger:model Build Webhook
*/
type BuildWebhook struct {

	/* Key=>Value information from the build itself
	 */
	Meta interface{} `json:"meta,omitempty"`

	/* Current status of the build

	Required: true
	*/
	Status string `json:"status"`
}

// Validate validates this build webhook
func (m *BuildWebhook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var buildWebhookTypeStatusPropEnum []interface{}

func (m *BuildWebhook) validateStatusEnum(path, location string, value string) error {
	if buildWebhookTypeStatusPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["SUCCESS","FAILURE","QUEUED","ABORTED","RUNNING"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			buildWebhookTypeStatusPropEnum = append(buildWebhookTypeStatusPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, buildWebhookTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BuildWebhook) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}