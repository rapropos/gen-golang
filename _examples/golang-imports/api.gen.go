// example-api-service v1.0.0 cae4e128f4fb4c938bfe1ea312deeea3dfd6b6af
// --
// Code generated by webrpc-gen@v0.23.2 with ../../../gen-golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=./proto/api.ridl -target=../../../gen-golang -out=./api.gen.go -pkg=main -server -client -fmt=false
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

)

const WebrpcHeader = "Webrpc"

const WebrpcHeaderValue = "webrpc@v0.23.2;gen-golang@unknown;example-api-service@v1.0.0"

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v1.0.0"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "cae4e128f4fb4c938bfe1ea312deeea3dfd6b6af"
}

type WebrpcGenVersions struct {
    WebrpcGenVersion string
    CodeGenName string
    CodeGenVersion string
    SchemaName string
    SchemaVersion string
}

func VersionFromHeader(h http.Header) (*WebrpcGenVersions, error) {
    if h.Get(WebrpcHeader) == "" {
        return nil, fmt.Errorf("header is empty or missing")
    }

    versions, err := parseWebrpcGenVersions(h.Get(WebrpcHeader))
    if err != nil {
        return nil, fmt.Errorf("webrpc header is invalid: %w", err)
    }

    return versions, nil
}

func parseWebrpcGenVersions(header string) (*WebrpcGenVersions, error) {
    versions := strings.Split(header, ";")
    if len(versions) < 3 {
        return nil, fmt.Errorf("expected at least 3 parts while parsing webrpc header: %v", header)
    }

    _, webrpcGenVersion, ok := strings.Cut(versions[0], "@")
    if !ok {
        return nil, fmt.Errorf("webrpc gen version could not be parsed from: %s", versions[0])
    }

    tmplTarget, tmplVersion, ok := strings.Cut(versions[1], "@")
    if !ok {
        return nil, fmt.Errorf("tmplTarget and tmplVersion could not be parsed from: %s", versions[1])
    }

    schemaName, schemaVersion, ok := strings.Cut(versions[2], "@")
    if !ok {
        return nil, fmt.Errorf("schema name and schema version could not be parsed from: %s", versions[2])
    }

    return &WebrpcGenVersions{
        WebrpcGenVersion: webrpcGenVersion,
        CodeGenName: tmplTarget,
        CodeGenVersion: tmplVersion,
        SchemaName: schemaName,
        SchemaVersion: schemaVersion,
    }, nil
}

//
// Common types
//

type User struct {
		Username string `json:"username"`
		Age uint32 `json:"age"`
}

type Location uint32

const (
	Location_TORONTO Location = 0
	Location_NEW_YORK Location = 1
)

var Location_name = map[uint32]string{
	0: "TORONTO",
	1: "NEW_YORK",
}

var Location_value = map[string]uint32{
	"TORONTO": 0,
	"NEW_YORK": 1,
}

func (x Location) String() string {
	return Location_name[uint32(x)]
}

func (x Location) MarshalText() ([]byte, error) {
	return []byte(Location_name[uint32(x)]), nil
}

func (x *Location) UnmarshalText(b []byte) error {
	*x = Location(Location_value[string(b)])
	return nil
}

func (x *Location) Is(values ...Location) bool {
	if x == nil {
		return false
	}
	for _, v := range values {
		if *x == v {
			return true
		}
	}
	return false
}

var methods = map[string]method{
	"/rpc/ExampleAPI/Ping": {
		Name: "Ping",
		Service: "ExampleAPI",
		Annotations: map[string]string{},
	},
	"/rpc/ExampleAPI/Status": {
		Name: "Status",
		Service: "ExampleAPI",
		Annotations: map[string]string{},
	},
	"/rpc/ExampleAPI/GetUsers": {
		Name: "GetUsers",
		Service: "ExampleAPI",
		Annotations: map[string]string{},
	},
}

func WebrpcMethods() map[string]method {
	res := make(map[string]method, len(methods))
	for k, v := range methods {
		res[k] = v
	}

	return res
}

var WebRPCServices = map[string][]string{
	"ExampleAPI": {
		"Ping",
		"Status",
		"GetUsers",
	},
}

//
// Server types
//

type ExampleAPI interface {
	Ping(ctx context.Context) error
	Status(ctx context.Context) (bool, error)
	GetUsers(ctx context.Context) ([]*User, Location, error)
}




//
// Client types
//

type ExampleAPIClient interface {
	Ping(ctx context.Context) error
	Status(ctx context.Context) (bool, error)
	GetUsers(ctx context.Context) ([]*User, Location, error)
}




//
// Server
//

type WebRPCServer interface {
	http.Handler
}

type exampleAPIServer struct {
	ExampleAPI
	OnError func(r *http.Request, rpcErr *WebRPCError)
	OnRequest func(w http.ResponseWriter, r *http.Request) error
}

func NewExampleAPIServer(svc ExampleAPI) *exampleAPIServer {
	return &exampleAPIServer{
		ExampleAPI: svc,
	}
}

func (s *exampleAPIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// In case of a panic, serve a HTTP 500 error and then panic.
		if rr := recover(); rr != nil {
			s.sendErrorJSON(w, r, ErrWebrpcServerPanic.WithCausef("%v", rr))
			panic(rr)
		}
	}()

	w.Header().Set(WebrpcHeader, WebrpcHeaderValue)

	ctx := r.Context()
	ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
	ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)
	ctx = context.WithValue(ctx, ServiceNameCtxKey, "ExampleAPI")

	r = r.WithContext(ctx)

	var handler func(ctx context.Context, w http.ResponseWriter, r *http.Request)
	switch r.URL.Path {
	case "/rpc/ExampleAPI/Ping":
		handler = s.servePingJSON
	case "/rpc/ExampleAPI/Status":
		handler = s.serveStatusJSON
	case "/rpc/ExampleAPI/GetUsers":
		handler = s.serveGetUsersJSON
	default:
		err := ErrWebrpcBadRoute.WithCausef("no webrpc method defined for path %v", r.URL.Path)
		s.sendErrorJSON(w, r, err)
		return
	}

	if r.Method != "POST" {
		w.Header().Add("Allow", "POST") // RFC 9110.
		err := ErrWebrpcBadMethod.WithCausef("unsupported HTTP method %v (only POST is allowed)", r.Method)
		s.sendErrorJSON(w, r, err)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if i := strings.Index(contentType, ";"); i >= 0 {
		contentType = contentType[:i]
	}
	contentType = strings.TrimSpace(strings.ToLower(contentType))

	switch contentType {
	case "application/json":
		if s.OnRequest != nil {
			if err := s.OnRequest(w, r); err != nil {
				rpcErr, ok := err.(WebRPCError)
				if !ok {
					rpcErr = ErrWebrpcEndpoint.WithCause(err)
				}
				s.sendErrorJSON(w, r, rpcErr)
				return
			}
		}

		handler(ctx, w, r)
	default:
		err := ErrWebrpcBadRequest.WithCausef("unsupported Content-Type %q (only application/json is allowed)", r.Header.Get("Content-Type"))
		s.sendErrorJSON(w, r, err)
	}
}

func (s *exampleAPIServer) servePingJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Ping")

	

	// Call service method implementation.
	err := s.ExampleAPI.Ping(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *exampleAPIServer) serveStatusJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Status")

	

	// Call service method implementation.
	ret0, err := s.ExampleAPI.Status(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 bool `json:"status"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCausef("failed to marshal json response: %w", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *exampleAPIServer) serveGetUsersJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetUsers")

	

	// Call service method implementation.
	ret0, ret1, err := s.ExampleAPI.GetUsers(ctx)
	if err != nil {
		rpcErr, ok := err.(WebRPCError)
		if !ok {
			rpcErr = ErrWebrpcEndpoint.WithCause(err)
		}
		s.sendErrorJSON(w, r, rpcErr)
		return
	}

	respPayload := struct {
		Ret0 []*User `json:"users"`
		Ret1 Location `json:"location"`
	}{ret0, ret1}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		s.sendErrorJSON(w, r, ErrWebrpcBadResponse.WithCausef("failed to marshal json response: %w", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}


func (s *exampleAPIServer) sendErrorJSON(w http.ResponseWriter, r *http.Request, rpcErr WebRPCError) {
	if s.OnError != nil {
		s.OnError(r, &rpcErr)
	}

	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rpcErr.HTTPStatus)

	respBody, _ := json.Marshal(rpcErr)
	w.Write(respBody)
}

func RespondWithError(w http.ResponseWriter, err error) {
	rpcErr, ok := err.(WebRPCError)
	if !ok {
		rpcErr = ErrWebrpcEndpoint.WithCause(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rpcErr.HTTPStatus)

	respBody, _ := json.Marshal(rpcErr)
	w.Write(respBody)
}



//
// Client
//

const ExampleAPIPathPrefix = "/rpc/ExampleAPI/"

type exampleAPIClient struct {
	client HTTPClient
	urls	 [3]string
}

func NewExampleAPIClient(addr string, client HTTPClient) ExampleAPIClient {
	prefix := urlBase(addr) + ExampleAPIPathPrefix
	urls := [3]string{
		prefix + "Ping",
		prefix + "Status",
		prefix + "GetUsers",
	}
	return &exampleAPIClient{
		client: client,
		urls:	 urls,
	}
}

func (c *exampleAPIClient) Ping(ctx context.Context) error {

	resp, err := doHTTPRequest(ctx, c.client, c.urls[0], nil, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return err
}

func (c *exampleAPIClient) Status(ctx context.Context) (bool, error) {
	out := struct {
		Ret0 bool `json:"status"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[1], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, err
}

func (c *exampleAPIClient) GetUsers(ctx context.Context) ([]*User, Location, error) {
	out := struct {
		Ret0 []*User `json:"users"`
		Ret1 Location `json:"location"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[2], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, out.Ret1, err
}

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
		req.Header.Set(WebrpcHeader, WebrpcHeaderValue)
	if headers, ok := HTTPRequestHeaders(ctx); ok {
		for k := range headers {
			for _, v := range headers[k] {
				req.Header.Add(k, v)
			}
		}
	}
	return req, nil
}

// doHTTPRequest is common code to make a request to the remote service.
func doHTTPRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) (*http.Response, error) {
	reqBody, err := json.Marshal(in)
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCausef("failed to marshal JSON body: %w", err)
	}
	if err = ctx.Err(); err != nil {
		return nil, ErrWebrpcRequestFailed.WithCausef("aborted because context was done: %w", err)
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCausef("could not build request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(err)
	}

	if resp.StatusCode != 200 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCausef("failed to read server error response body: %w", err)
		}

		var rpcErr WebRPCError
		if err := json.Unmarshal(respBody, &rpcErr); err != nil {
			return nil, ErrWebrpcBadResponse.WithCausef("failed to unmarshal server error: %w", err)
		}
		if rpcErr.Cause != "" {
			rpcErr.cause = errors.New(rpcErr.Cause)
		}
		return nil, rpcErr
	}

	if out != nil {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCausef("failed to read response body: %w", err)
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCausef("failed to unmarshal JSON response body: %w", err)
		}
	}

	return resp, nil
}

func WithHTTPRequestHeaders(ctx context.Context, h http.Header) (context.Context, error) {
	if _, ok := h["Accept"]; ok {
		return nil, errors.New("provided header cannot set Accept")
	}
	if _, ok := h["Content-Type"]; ok {
		return nil, errors.New("provided header cannot set Content-Type")
	}

	copied := make(http.Header, len(h))
	for k, vv := range h {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}

	return context.WithValue(ctx, HTTPClientRequestHeadersCtxKey, copied), nil
}

func HTTPRequestHeaders(ctx context.Context) (http.Header, bool) {
	h, ok := ctx.Value(HTTPClientRequestHeadersCtxKey).(http.Header)
	return h, ok
}

//
// Helpers
//

type method struct  {
	Name string
	Service string
	Annotations map[string]string
}

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "webrpc context value " + k.name
}

var (
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}
	HTTPResponseWriterCtxKey = &contextKey{"HTTPResponseWriter"}

	HTTPRequestCtxKey = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)

func ServiceNameFromContext(ctx context.Context) string {
	service, _ := ctx.Value(ServiceNameCtxKey).(string)
	return service
}

func MethodNameFromContext(ctx context.Context) string {
	method, _ := ctx.Value(MethodNameCtxKey).(string)
	return method
}

func RequestFromContext(ctx context.Context) *http.Request {
	r, _ := ctx.Value(HTTPRequestCtxKey).(*http.Request)
	return r
}

func MethodCtx(ctx context.Context) (method, bool) {
	req := RequestFromContext(ctx)
	if req == nil {
		return method{}, false
	}

	m, ok := methods[req.URL.Path]
	if !ok {
		return method{}, false
	}

	return m, true
}


func ResponseWriterFromContext(ctx context.Context) http.ResponseWriter {
	w, _ := ctx.Value(HTTPResponseWriterCtxKey).(http.ResponseWriter)
	return w
}

//
// Errors
//

type WebRPCError struct {
	Name       string `json:"error"`
	Code       int    `json:"code"`
	Message    string `json:"msg"`
	Cause      string `json:"cause,omitempty"`
	HTTPStatus int    `json:"status"`
	cause      error
}

var _ error = WebRPCError{}

func (e WebRPCError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s %d: %s: %v", e.Name, e.Code, e.Message, e.cause)
	}
	return fmt.Sprintf("%s %d: %s", e.Name, e.Code, e.Message)
}

func (e WebRPCError) Is(target error) bool {
	if target == nil {
		return false
	}
	if rpcErr, ok := target.(WebRPCError); ok {
		return rpcErr.Code == e.Code
	}
	return errors.Is(e.cause, target)
}

func (e WebRPCError) Unwrap() error {
	return e.cause
}

func (e WebRPCError) WithCause(cause error) WebRPCError {
	err := e
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

func (e WebRPCError) WithCausef(format string, args ...interface{}) WebRPCError {
	cause := fmt.Errorf(format, args...)
	err := e
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

// Deprecated: Use .WithCause() method on WebRPCError.
func ErrorWithCause(rpcErr WebRPCError, cause error) WebRPCError {
	return rpcErr.WithCause(cause)
}

// Webrpc errors
var (
	ErrWebrpcEndpoint = WebRPCError{Code: 0, Name: "WebrpcEndpoint", Message: "endpoint error", HTTPStatus: 400}
	ErrWebrpcRequestFailed = WebRPCError{Code: -1, Name: "WebrpcRequestFailed", Message: "request failed", HTTPStatus: 400}
	ErrWebrpcBadRoute = WebRPCError{Code: -2, Name: "WebrpcBadRoute", Message: "bad route", HTTPStatus: 404}
	ErrWebrpcBadMethod = WebRPCError{Code: -3, Name: "WebrpcBadMethod", Message: "bad method", HTTPStatus: 405}
	ErrWebrpcBadRequest = WebRPCError{Code: -4, Name: "WebrpcBadRequest", Message: "bad request", HTTPStatus: 400}
	ErrWebrpcBadResponse = WebRPCError{Code: -5, Name: "WebrpcBadResponse", Message: "bad response", HTTPStatus: 500}
	ErrWebrpcServerPanic = WebRPCError{Code: -6, Name: "WebrpcServerPanic", Message: "server panic", HTTPStatus: 500}
	ErrWebrpcInternalError = WebRPCError{Code: -7, Name: "WebrpcInternalError", Message: "internal error", HTTPStatus: 500}
	ErrWebrpcClientDisconnected = WebRPCError{Code: -8, Name: "WebrpcClientDisconnected", Message: "client disconnected", HTTPStatus: 400}
	ErrWebrpcStreamLost = WebRPCError{Code: -9, Name: "WebrpcStreamLost", Message: "stream lost", HTTPStatus: 400}
	ErrWebrpcStreamFinished = WebRPCError{Code: -10, Name: "WebrpcStreamFinished", Message: "stream finished", HTTPStatus: 200}
)