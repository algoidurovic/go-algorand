// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error
	// Return a list of participation keys
	// (GET /v2/participation)
	GetParticipationKeys(ctx echo.Context) error
	// Add a participation key to the node
	// (POST /v2/participation)
	AddParticipationKey(ctx echo.Context) error
	// Delete a given participation key by ID
	// (DELETE /v2/participation/{participation-id})
	DeleteParticipationKeyByID(ctx echo.Context, participationId string) error
	// Get participation key info given a participation ID
	// (GET /v2/participation/{participation-id})
	GetParticipationKeyByID(ctx echo.Context, participationId string) error
	// Append state proof keys to a participation key
	// (POST /v2/participation/{participation-id})
	AppendKeys(ctx echo.Context, participationId string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":  true,
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.GET("/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST("/v2/participation", wrapper.AddParticipationKey, m...)
	router.DELETE("/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET("/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST("/v2/participation/:participation-id", wrapper.AppendKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"qf+5PnndKlx9uvr/AQAA//8Mn0i4r7EAAA==",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
