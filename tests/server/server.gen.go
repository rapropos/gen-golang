// Test v0.10.0 249dca96f024d73e0d3e7a762aa32065e228c5fa
// --
// Code generated by webrpc-gen@v0.10.0 with golang generator. DO NOT EDIT.
//
// webrpc-gen@v0.10.0.darwin-arm64 -schema=./schema/test.ridl -target=golang -pkg=server -server -out=./server/server.gen.go
package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v0.10.0"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "249dca96f024d73e0d3e7a762aa32065e228c5fa"
}

//
// Types
//

type Status uint32

const (
	Status_AVAILABLE     Status = 0
	Status_NOT_AVAILABLE Status = 1
)

var Status_name = map[uint32]string{
	0: "AVAILABLE",
	1: "NOT_AVAILABLE",
}

var Status_value = map[string]uint32{
	"AVAILABLE":     0,
	"NOT_AVAILABLE": 1,
}

func (x Status) String() string {
	return Status_name[uint32(x)]
}

func (x Status) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(Status_name[uint32(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *Status) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = Status(Status_value[j])
	return nil
}

type Simple struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID       uint64 `json:"id" db:"id"`
	Username string `json:"USERNAME" db:"username"`
	Role     string `json:"role" db:"-"`
}

type Complex struct {
	Meta              map[string]interface{}       `json:"meta"`
	MetaNestedExample map[string]map[string]uint32 `json:"metaNestedExample"`
	NamesList         []string                     `json:"namesList"`
	NumsList          []int64                      `json:"numsList"`
	DoubleArray       [][]string                   `json:"doubleArray"`
	ListOfMaps        []map[string]uint32          `json:"listOfMaps"`
	ListOfUsers       []*User                      `json:"listOfUsers"`
	MapOfUsers        map[string]*User             `json:"mapOfUsers"`
	User              *User                        `json:"user"`
	Enum              *Status                      `json:"enum"`
}

type TestApi interface {
	GetEmpty(ctx context.Context) error
	GetError(ctx context.Context) error
	GetOne(ctx context.Context) (*Simple, error)
	SendOne(ctx context.Context, one *Simple) error
	GetMulti(ctx context.Context) (*Simple, *Simple, *Simple, error)
	SendMulti(ctx context.Context, one *Simple, two *Simple, three *Simple) error
	GetComplex(ctx context.Context) (*Complex, error)
	SendComplex(ctx context.Context, complex *Complex) error
}

var WebRPCServices = map[string][]string{
	"TestApi": {
		"GetEmpty",
		"GetError",
		"GetOne",
		"SendOne",
		"GetMulti",
		"SendMulti",
		"GetComplex",
		"SendComplex",
	},
}

//
// Server
//

type WebRPCServer interface {
	http.Handler
}

type testApiServer struct {
	TestApi
}

func NewTestApiServer(svc TestApi) WebRPCServer {
	return &testApiServer{
		TestApi: svc,
	}
}

func (s *testApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
	ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)
	ctx = context.WithValue(ctx, ServiceNameCtxKey, "TestApi")

	if r.Method != "POST" {
		err := Errorf(ErrBadRoute, "unsupported method %q (only POST is allowed)", r.Method)
		RespondWithError(w, err)
		return
	}

	switch r.URL.Path {
	case "/rpc/TestApi/GetEmpty":
		s.serveGetEmpty(ctx, w, r)
		return
	case "/rpc/TestApi/GetError":
		s.serveGetError(ctx, w, r)
		return
	case "/rpc/TestApi/GetOne":
		s.serveGetOne(ctx, w, r)
		return
	case "/rpc/TestApi/SendOne":
		s.serveSendOne(ctx, w, r)
		return
	case "/rpc/TestApi/GetMulti":
		s.serveGetMulti(ctx, w, r)
		return
	case "/rpc/TestApi/SendMulti":
		s.serveSendMulti(ctx, w, r)
		return
	case "/rpc/TestApi/GetComplex":
		s.serveGetComplex(ctx, w, r)
		return
	case "/rpc/TestApi/SendComplex":
		s.serveSendComplex(ctx, w, r)
		return
	default:
		err := Errorf(ErrBadRoute, "no handler for path %q", r.URL.Path)
		RespondWithError(w, err)
		return
	}
}

func (s *testApiServer) serveGetEmpty(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetEmptyJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *testApiServer) serveGetEmptyJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetEmpty")

	// Call service method
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		err = s.TestApi.GetEmpty(ctx)
	}()

	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *testApiServer) serveGetError(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetErrorJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *testApiServer) serveGetErrorJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetError")

	// Call service method
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		err = s.TestApi.GetError(ctx)
	}()

	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *testApiServer) serveGetOne(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetOneJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *testApiServer) serveGetOneJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetOne")

	// Call service method
	var ret0 *Simple
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		ret0, err = s.TestApi.GetOne(ctx)
	}()
	respContent := struct {
		Ret0 *Simple `json:"one"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to marshal json response")
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *testApiServer) serveSendOne(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSendOneJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *testApiServer) serveSendOneJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "SendOne")
	reqContent := struct {
		Arg0 *Simple `json:"one"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to read request data")
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = WrapError(ErrInvalidArgument, err, "failed to unmarshal request data")
		RespondWithError(w, err)
		return
	}

	// Call service method
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		err = s.TestApi.SendOne(ctx, reqContent.Arg0)
	}()

	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *testApiServer) serveGetMulti(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetMultiJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *testApiServer) serveGetMultiJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetMulti")

	// Call service method
	var ret0 *Simple
	var ret1 *Simple
	var ret2 *Simple
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		ret0, ret1, ret2, err = s.TestApi.GetMulti(ctx)
	}()
	respContent := struct {
		Ret0 *Simple `json:"one"`
		Ret1 *Simple `json:"two"`
		Ret2 *Simple `json:"three"`
	}{ret0, ret1, ret2}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to marshal json response")
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *testApiServer) serveSendMulti(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSendMultiJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *testApiServer) serveSendMultiJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "SendMulti")
	reqContent := struct {
		Arg0 *Simple `json:"one"`
		Arg1 *Simple `json:"two"`
		Arg2 *Simple `json:"three"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to read request data")
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = WrapError(ErrInvalidArgument, err, "failed to unmarshal request data")
		RespondWithError(w, err)
		return
	}

	// Call service method
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		err = s.TestApi.SendMulti(ctx, reqContent.Arg0, reqContent.Arg1, reqContent.Arg2)
	}()

	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *testApiServer) serveGetComplex(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetComplexJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *testApiServer) serveGetComplexJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetComplex")

	// Call service method
	var ret0 *Complex
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		ret0, err = s.TestApi.GetComplex(ctx)
	}()
	respContent := struct {
		Ret0 *Complex `json:"complex"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to marshal json response")
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *testApiServer) serveSendComplex(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSendComplexJSON(ctx, w, r)
	default:
		err := Errorf(ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		RespondWithError(w, err)
	}
}

func (s *testApiServer) serveSendComplexJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "SendComplex")
	reqContent := struct {
		Arg0 *Complex `json:"complex"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to read request data")
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = WrapError(ErrInvalidArgument, err, "failed to unmarshal request data")
		RespondWithError(w, err)
		return
	}

	// Call service method
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		err = s.TestApi.SendComplex(ctx, reqContent.Arg0)
	}()

	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func RespondWithError(w http.ResponseWriter, err error) {
	rpcErr, ok := err.(Error)
	if !ok {
		rpcErr = WrapError(ErrInternal, err, "webrpc error")
	}

	statusCode := HTTPStatusFromErrorCode(rpcErr.Code())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	respBody, _ := json.Marshal(rpcErr.Payload())
	w.Write(respBody)
}

//
// Helpers
//

type ErrorPayload struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Cause  string `json:"cause,omitempty"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}

type Error interface {
	// Code is of the valid error codes
	Code() ErrorCode

	// Msg returns a human-readable, unstructured messages describing the error
	Msg() string

	// Cause is reason for the error
	Cause() error

	// Error returns a string of the form "webrpc error <Code>: <Msg>"
	Error() string

	// Error response payload
	Payload() ErrorPayload
}

func Errorf(code ErrorCode, msgf string, args ...interface{}) Error {
	msg := fmt.Sprintf(msgf, args...)
	if IsValidErrorCode(code) {
		return &rpcErr{code: code, msg: msg}
	}
	return &rpcErr{code: ErrInternal, msg: "invalid error type " + string(code)}
}

func WrapError(code ErrorCode, cause error, format string, args ...interface{}) Error {
	msg := fmt.Sprintf(format, args...)
	if IsValidErrorCode(code) {
		return &rpcErr{code: code, msg: msg, cause: cause}
	}
	return &rpcErr{code: ErrInternal, msg: "invalid error type " + string(code), cause: cause}
}

func Failf(format string, args ...interface{}) Error {
	return Errorf(ErrFail, format, args...)
}

func WrapFailf(cause error, format string, args ...interface{}) Error {
	return WrapError(ErrFail, cause, format, args...)
}

func ErrorNotFound(format string, args ...interface{}) Error {
	return Errorf(ErrNotFound, format, args...)
}

func ErrorInvalidArgument(argument string, validationMsg string) Error {
	return Errorf(ErrInvalidArgument, argument+" "+validationMsg)
}

func ErrorRequiredArgument(argument string) Error {
	return ErrorInvalidArgument(argument, "is required")
}

func ErrorInternal(format string, args ...interface{}) Error {
	return Errorf(ErrInternal, format, args...)
}

type ErrorCode string

const (
	// Unknown error. For example when handling errors raised by APIs that do not
	// return enough error information.
	ErrUnknown ErrorCode = "unknown"

	// Fail error. General failure error type.
	ErrFail ErrorCode = "fail"

	// Canceled indicates the operation was cancelled (typically by the caller).
	ErrCanceled ErrorCode = "canceled"

	// InvalidArgument indicates client specified an invalid argument. It
	// indicates arguments that are problematic regardless of the state of the
	// system (i.e. a malformed file name, required argument, number out of range,
	// etc.).
	ErrInvalidArgument ErrorCode = "invalid argument"

	// DeadlineExceeded means operation expired before completion. For operations
	// that change the state of the system, this error may be returned even if the
	// operation has completed successfully (timeout).
	ErrDeadlineExceeded ErrorCode = "deadline exceeded"

	// NotFound means some requested entity was not found.
	ErrNotFound ErrorCode = "not found"

	// BadRoute means that the requested URL path wasn't routable to a webrpc
	// service and method. This is returned by the generated server, and usually
	// shouldn't be returned by applications. Instead, applications should use
	// NotFound or Unimplemented.
	ErrBadRoute ErrorCode = "bad route"

	// AlreadyExists means an attempt to create an entity failed because one
	// already exists.
	ErrAlreadyExists ErrorCode = "already exists"

	// PermissionDenied indicates the caller does not have permission to execute
	// the specified operation. It must not be used if the caller cannot be
	// identified (Unauthenticated).
	ErrPermissionDenied ErrorCode = "permission denied"

	// Unauthenticated indicates the request does not have valid authentication
	// credentials for the operation.
	ErrUnauthenticated ErrorCode = "unauthenticated"

	// ResourceExhausted indicates some resource has been exhausted, perhaps a
	// per-user quota, or perhaps the entire file system is out of space.
	ErrResourceExhausted ErrorCode = "resource exhausted"

	// FailedPrecondition indicates operation was rejected because the system is
	// not in a state required for the operation's execution. For example, doing
	// an rmdir operation on a directory that is non-empty, or on a non-directory
	// object, or when having conflicting read-modify-write on the same resource.
	ErrFailedPrecondition ErrorCode = "failed precondition"

	// Aborted indicates the operation was aborted, typically due to a concurrency
	// issue like sequencer check failures, transaction aborts, etc.
	ErrAborted ErrorCode = "aborted"

	// OutOfRange means operation was attempted past the valid range. For example,
	// seeking or reading past end of a paginated collection.
	//
	// Unlike InvalidArgument, this error indicates a problem that may be fixed if
	// the system state changes (i.e. adding more items to the collection).
	//
	// There is a fair bit of overlap between FailedPrecondition and OutOfRange.
	// We recommend using OutOfRange (the more specific error) when it applies so
	// that callers who are iterating through a space can easily look for an
	// OutOfRange error to detect when they are done.
	ErrOutOfRange ErrorCode = "out of range"

	// Unimplemented indicates operation is not implemented or not
	// supported/enabled in this service.
	ErrUnimplemented ErrorCode = "unimplemented"

	// Internal errors. When some invariants expected by the underlying system
	// have been broken. In other words, something bad happened in the library or
	// backend service. Do not confuse with HTTP Internal Server Error; an
	// Internal error could also happen on the client code, i.e. when parsing a
	// server response.
	ErrInternal ErrorCode = "internal"

	// Unavailable indicates the service is currently unavailable. This is a most
	// likely a transient condition and may be corrected by retrying with a
	// backoff.
	ErrUnavailable ErrorCode = "unavailable"

	// DataLoss indicates unrecoverable data loss or corruption.
	ErrDataLoss ErrorCode = "data loss"

	// ErrNone is the zero-value, is considered an empty error and should not be
	// used.
	ErrNone ErrorCode = ""
)

func HTTPStatusFromErrorCode(code ErrorCode) int {
	switch code {
	case ErrCanceled:
		return 408 // RequestTimeout
	case ErrUnknown:
		return 400 // Bad Request
	case ErrFail:
		return 422 // Unprocessable Entity
	case ErrInvalidArgument:
		return 400 // BadRequest
	case ErrDeadlineExceeded:
		return 408 // RequestTimeout
	case ErrNotFound:
		return 404 // Not Found
	case ErrBadRoute:
		return 404 // Not Found
	case ErrAlreadyExists:
		return 409 // Conflict
	case ErrPermissionDenied:
		return 403 // Forbidden
	case ErrUnauthenticated:
		return 401 // Unauthorized
	case ErrResourceExhausted:
		return 403 // Forbidden
	case ErrFailedPrecondition:
		return 412 // Precondition Failed
	case ErrAborted:
		return 409 // Conflict
	case ErrOutOfRange:
		return 400 // Bad Request
	case ErrUnimplemented:
		return 501 // Not Implemented
	case ErrInternal:
		return 500 // Internal Server Error
	case ErrUnavailable:
		return 503 // Service Unavailable
	case ErrDataLoss:
		return 500 // Internal Server Error
	case ErrNone:
		return 200 // OK
	default:
		return 0 // Invalid!
	}
}

func IsErrorCode(err error, code ErrorCode) bool {
	if rpcErr, ok := err.(Error); ok {
		if rpcErr.Code() == code {
			return true
		}
	}
	return false
}

func IsValidErrorCode(code ErrorCode) bool {
	return HTTPStatusFromErrorCode(code) != 0
}

type rpcErr struct {
	code  ErrorCode
	msg   string
	cause error
}

func (e *rpcErr) Code() ErrorCode {
	return e.code
}

func (e *rpcErr) Msg() string {
	return e.msg
}

func (e *rpcErr) Cause() error {
	return e.cause
}

func (e *rpcErr) Error() string {
	if e.cause != nil && e.cause.Error() != "" {
		if e.msg != "" {
			return fmt.Sprintf("webrpc %s error: %s -- %s", e.code, e.cause.Error(), e.msg)
		} else {
			return fmt.Sprintf("webrpc %s error: %s", e.code, e.cause.Error())
		}
	} else {
		return fmt.Sprintf("webrpc %s error: %s", e.code, e.msg)
	}
}

func (e *rpcErr) Payload() ErrorPayload {
	statusCode := HTTPStatusFromErrorCode(e.Code())
	errPayload := ErrorPayload{
		Status: statusCode,
		Code:   string(e.Code()),
		Msg:    e.Msg(),
		Error:  e.Error(),
	}
	if e.Cause() != nil {
		errPayload.Cause = e.Cause().Error()
	}
	return errPayload
}

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "webrpc context value " + k.name
}

var (
	// For Client
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}

	// For Server
	HTTPResponseWriterCtxKey = &contextKey{"HTTPResponseWriter"}

	HTTPRequestCtxKey = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)
