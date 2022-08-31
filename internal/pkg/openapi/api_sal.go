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
	//"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// SalApiController binds http requests to an api service and writes the service results to the http response
type SalApiController struct {
	service      SalApiServicer
	errorHandler ErrorHandler
}

// SalApiOption for how the controller is set up.
type SalApiOption func(*SalApiController)

// WithSalApiErrorHandler inject ErrorHandler into controller
func WithSalApiErrorHandler(h ErrorHandler) SalApiOption {
	return func(c *SalApiController) {
		c.errorHandler = h
	}
}

// NewSalApiController creates a default api controller
func NewSalApiController(s SalApiServicer, opts ...SalApiOption) Router {
	controller := &SalApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SalApiController
func (c *SalApiController) Routes() Routes {
	return Routes{
		{
			"DeleteSal",
			strings.ToUpper("Delete"),
			"/api/v1/sal/{empId}",
			c.DeleteSal,
		},
		{
			"GetSalByEmpId",
			strings.ToUpper("Get"),
			"/api/v1/sal/{empId}",
			c.GetSalByEmpId,
		},
		{
			"GetWageSumByMgrId",
			strings.ToUpper("Get"),
			"/api/v1/sal/wagesum/{empId}",
			c.GetWageSumByMgrId,
		},
		{
			"UpdateSalWithForm",
			strings.ToUpper("Post"),
			"/api/v1/sal/{empId}",
			c.UpdateSalWithForm,
		},
	}
}

// DeleteSal - Deletes a sal
func (c *SalApiController) DeleteSal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	empIdParam, err := parseInt64Parameter(params["empId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.DeleteSal(r.Context(), empIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetSalByEmpId - Find sal by ID
func (c *SalApiController) GetSalByEmpId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	empIdParam, err := parseInt64Parameter(params["empId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetSalByEmpId(r.Context(), empIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetWageSumByMgrId - Find sum sal by manager ID
func (c *SalApiController) GetWageSumByMgrId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	empIdParam, err := parseInt64Parameter(params["empId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetWageSumByMgrId(r.Context(), empIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateSalWithForm - Updates a sal in the store with form data
func (c *SalApiController) UpdateSalWithForm(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	empIdParam, err := parseInt64Parameter(params["empId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	valueParam := query.Get("value")
	result, err := c.service.UpdateSalWithForm(r.Context(), empIdParam, valueParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
