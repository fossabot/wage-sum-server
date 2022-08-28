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
	"context"
	"net/http"
)



// EmpApiRouter defines the required methods for binding the api requests to a responses for the EmpApi
// The EmpApiRouter implementation should parse necessary information from the http request,
// pass the data to a EmpApiServicer to perform the required actions, then write the service results to the http response.
type EmpApiRouter interface { 
	AddEmp(http.ResponseWriter, *http.Request)
	DeleteEmp(http.ResponseWriter, *http.Request)
	FindEmpsByType(http.ResponseWriter, *http.Request)
	GetEmpById(http.ResponseWriter, *http.Request)
	UpdateEmp(http.ResponseWriter, *http.Request)
	UpdateEmpWithForm(http.ResponseWriter, *http.Request)
}
// SalApiRouter defines the required methods for binding the api requests to a responses for the SalApi
// The SalApiRouter implementation should parse necessary information from the http request,
// pass the data to a SalApiServicer to perform the required actions, then write the service results to the http response.
type SalApiRouter interface { 
	GetSalByEmpId(http.ResponseWriter, *http.Request)
	GetWageSumByMgrId(http.ResponseWriter, *http.Request)
	UpdateSalWithForm(http.ResponseWriter, *http.Request)
}


// EmpApiServicer defines the api actions for the EmpApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type EmpApiServicer interface { 
	AddEmp(context.Context, Emp) (ImplResponse, error)
	DeleteEmp(context.Context, int64, string) (ImplResponse, error)
	FindEmpsByType(context.Context, string) (ImplResponse, error)
	GetEmpById(context.Context, int64) (ImplResponse, error)
	UpdateEmp(context.Context, Emp) (ImplResponse, error)
	UpdateEmpWithForm(context.Context, int64, string, string) (ImplResponse, error)
}


// SalApiServicer defines the api actions for the SalApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type SalApiServicer interface { 
	GetSalByEmpId(context.Context, string) (ImplResponse, error)
	GetWageSumByMgrId(context.Context, int64) (ImplResponse, error)
	UpdateSalWithForm(context.Context, int64, string) (ImplResponse, error)
}
