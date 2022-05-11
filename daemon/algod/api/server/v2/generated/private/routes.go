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
