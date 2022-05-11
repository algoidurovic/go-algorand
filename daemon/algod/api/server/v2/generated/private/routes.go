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
	"AAA=",
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
