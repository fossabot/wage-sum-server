/*
 * Employee wage sum
 *
 * Wage sum - demo application with GO language
 *
 * API version: 0.2.0
 * Contact: lsmhun@github
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// EmpApiController binds http requests to an api service and writes the service results to the http response
type EmpApiController struct {
	service      EmpApiServicer
	errorHandler ErrorHandler
}

// EmpApiOption for how the controller is set up.
type EmpApiOption func(*EmpApiController)

// WithEmpApiErrorHandler inject ErrorHandler into controller
func WithEmpApiErrorHandler(h ErrorHandler) EmpApiOption {
	return func(c *EmpApiController) {
		c.errorHandler = h
	}
}

// NewEmpApiController creates a default api controller
func NewEmpApiController(s EmpApiServicer, opts ...EmpApiOption) Router {
	controller := &EmpApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the EmpApiController
func (c *EmpApiController) Routes() Routes {
	return Routes{
		{
			"AddEmp",
			strings.ToUpper("Post"),
			"/api/v1/emp",
			c.AddEmp,
		},
		{
			"DeleteEmp",
			strings.ToUpper("Delete"),
			"/api/v1/emp/{empId}",
			c.DeleteEmp,
		},
		{
			"FindEmpsByType",
			strings.ToUpper("Get"),
			"/api/v1/emp/findByType",
			c.FindEmpsByType,
		},
		{
			"GetEmpById",
			strings.ToUpper("Get"),
			"/api/v1/emp/{empId}",
			c.GetEmpById,
		},
		{
			"UpdateEmp",
			strings.ToUpper("Put"),
			"/api/v1/emp",
			c.UpdateEmp,
		},
		{
			"UpdateEmpWithForm",
			strings.ToUpper("Post"),
			"/api/v1/emp/{empId}",
			c.UpdateEmpWithForm,
		},
	}
}

// AddEmp - Add a new emp to the store
func (c *EmpApiController) AddEmp(w http.ResponseWriter, r *http.Request) {
	empParam := Emp{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&empParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEmpRequired(empParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddEmp(r.Context(), empParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	err = EncodeJSONResponse(result.Body, &result.Code, w)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
}

// DeleteEmp - Deletes a emp
func (c *EmpApiController) DeleteEmp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	empIdParam, err := parseInt64Parameter(params["empId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.DeleteEmp(r.Context(), empIdParam, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	err = EncodeJSONResponse(result.Body, &result.Code, w)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
}

// FindEmpsByType - Finds emps by type
func (c *EmpApiController) FindEmpsByType(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	type_Param := query.Get("type")
	result, err := c.service.FindEmpsByType(r.Context(), type_Param)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	err = EncodeJSONResponse(result.Body, &result.Code, w)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
}

// GetEmpById - Find emp by ID
func (c *EmpApiController) GetEmpById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	empIdParam, err := parseInt64Parameter(params["empId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetEmpById(r.Context(), empIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	err = EncodeJSONResponse(result.Body, &result.Code, w)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
}

// UpdateEmp - Update an existing emp
func (c *EmpApiController) UpdateEmp(w http.ResponseWriter, r *http.Request) {
	empParam := Emp{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&empParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEmpRequired(empParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateEmp(r.Context(), empParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	err = EncodeJSONResponse(result.Body, &result.Code, w)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
}

// UpdateEmpWithForm - Updates a emp in the store with form data
func (c *EmpApiController) UpdateEmpWithForm(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	empIdParam, err := parseInt64Parameter(params["empId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	nameParam := query.Get("name")
	statusParam := query.Get("status")
	result, err := c.service.UpdateEmpWithForm(r.Context(), empIdParam, nameParam, statusParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	err = EncodeJSONResponse(result.Body, &result.Code, w)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
}
