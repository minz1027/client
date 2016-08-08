package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	// "github.com/go-openapi/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*GetBuild get build

swagger:model Get Build
*/
type GetBuild struct {

	/* Reason why this build started

	Required: true
	*/
	Cause string `json:"cause"`

	/* Container this build is running in
	 */
	Container *string `json:"container,omitempty"`

	/* When this build was created

	Required: true
	*/
	CreateTime string `json:"createTime"`

	/* When this build stopped running
	 */
	EndTime *string `json:"endTime,omitempty"`

	/* Identifier of this build

	Required: true
	*/
	ID string `json:"id"`

	/* Identifier of the Job

	Required: true
	*/
	JobID string `json:"jobId"`

	/* Key=>Value information from the build itself
	 */
	Meta interface{} `json:"meta,omitempty"`

	/* Incrementing number of a Job

	Required: true
	*/
	Number float64 `json:"number"`

	/* Input parameters that defined this build
	 */
	Parameters interface{} `json:"parameters,omitempty"`

	/* Parent build in the case of matrix jobs
	 */
	ParentBuildID *string `json:"parentBuildId,omitempty"`

	/* SHA this project was built on
	 */
	Sha *string `json:"sha,omitempty"`

	/* When this build started on a build machine
	 */
	StartTime *strfmt.Date `json:"startTime,omitempty"`

	/* Current status of the build

	Required: true
	*/
	Status string `json:"status"`
}

// Validate validates this get build
func (m *GetBuild) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreateTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateJobID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBuild) validateCause(formats strfmt.Registry) error {

	if err := validate.RequiredString("cause", "body", string(m.Cause)); err != nil {
		return err
	}

	return nil
}

func (m *GetBuild) validateCreateTime(formats strfmt.Registry) error {

	if err := validate.Required("createTime", "body", string(m.CreateTime)); err != nil {
		return err
	}

	return nil
}

func (m *GetBuild) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *GetBuild) validateJobID(formats strfmt.Registry) error {

	if err := validate.RequiredString("jobId", "body", string(m.JobID)); err != nil {
		return err
	}

	return nil
}

func (m *GetBuild) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", float64(m.Number)); err != nil {
		return err
	}

	return nil
}

var getBuildTypeStatusPropEnum []interface{}

func (m *GetBuild) validateStatusEnum(path, location string, value string) error {
	if getBuildTypeStatusPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["SUCCESS","FAILURE","QUEUED","ABORTED","RUNNING"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			getBuildTypeStatusPropEnum = append(getBuildTypeStatusPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, getBuildTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetBuild) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}