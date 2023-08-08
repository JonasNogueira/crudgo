package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOpening

type CreateStudentRequest struct {
	Name   string `json:"name"`
	Period string `json:"period"`
	Lab    string `json:"lab"`
}

func (r *CreateStudentRequest) Validate() error {
	if r.Name == "" && r.Period == "" && r.Lab == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Period == "" {
		return errParamIsRequired("period", "string")
	}
	if r.Lab == "" {
		return errParamIsRequired("lab", "string")
	}

	return nil
}

// UpdateOpening

type UpdateStudentRequest struct {
	Name   string `json:"name"`
	Period string `json:"period"`
	Lab    string `json:"Lab"`
}

func (r *UpdateStudentRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.Name != "" || r.Period != "" || r.Lab != "" {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
