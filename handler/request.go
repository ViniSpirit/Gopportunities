package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamIsRequired(param string, typ string) error {

	return fmt.Errorf("param %s (type %s) is required", param, typ)
}

func (r *CreateOpeningRequest) Validate(ctx *gin.Context) error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate(ctx *gin.Context) error {
	// if any field is provided, validation is ok
	if r.Role == "" || r.Company == "" || r.Location == "" || r.Link == "" || r.Remote == nil || r.Salary <= 0 {
		return nil
	}
	// if none of the fields are provided, return an error
	return fmt.Errorf("at least one field is required to update an opening")
}
