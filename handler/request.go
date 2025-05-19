package handler

import "fmt"

func errParamIsRequied(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOpening

type CreateOpeningRequest struct {
	Role      	string    `json:"role"`
    Company   	string    `json:"company"`
    Location  	string    `json:"location"`
    Remote    	*bool     `json:"remote"`
    Link      	string    `json:"link"`
    Salary    	int64     `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil  && r.Link =="" && r.Salary <= 0{
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Role == "" {
		return errParamIsRequied("role", "string")
	}

	if r.Company == "" {
		return errParamIsRequied("company", "string")
	}

	if r.Location == "" {
		return errParamIsRequied("location", "string")
	}

	if r.Link == "" {
		return errParamIsRequied("link", "string")
	}

	if r.Remote == nil{
		return errParamIsRequied("remote", "bool")
	}

	if r.Salary <= 0 {
		return errParamIsRequied("salary", "int64")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role      	string    `json:"role"`
    Company   	string    `json:"company"`
    Location  	string    `json:"location"`
    Remote    	*bool     `json:"remote"`
    Link      	string    `json:"link"`
    Salary    	int64     `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// if any field is provided, validation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}

	// if none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}
