// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// Athlete defines model for Athlete.
type Athlete struct {
	BadgeTypeId   *int       `json:"badge_type_id,omitempty"`
	Bio           *string    `json:"bio,omitempty"`
	City          *string    `json:"city,omitempty"`
	Country       *string    `json:"country,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	Firstname     *string    `json:"firstname,omitempty"`
	Follower      *int       `json:"follower,omitempty"`
	Friend        *int       `json:"friend,omitempty"`
	Id            *int       `json:"id,omitempty"`
	Lastname      *string    `json:"lastname,omitempty"`
	Premium       *bool      `json:"premium,omitempty"`
	Profile       *string    `json:"profile,omitempty"`
	ProfileMedium *string    `json:"profile_medium,omitempty"`
	ResourceState *int       `json:"resource_state,omitempty"`
	Sex           *string    `json:"sex,omitempty"`
	State         *string    `json:"state,omitempty"`
	Summit        *bool      `json:"summit,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	Username      *string    `json:"username,omitempty"`
	Weight        *float32   `json:"weight,omitempty"`
}

// AthleteStats defines model for AthleteStats.
type AthleteStats struct {
	AllRideTotals             *Totals  `json:"all_ride_totals,omitempty"`
	AllRunTotals              *Totals  `json:"all_run_totals,omitempty"`
	AllSwimTotals             *Totals  `json:"all_swim_totals,omitempty"`
	BiggestClimbElevationGain *float32 `json:"biggest_climb_elevation_gain,omitempty"`
	BiggestRideDistance       *float32 `json:"biggest_ride_distance,omitempty"`
	RecentRideTotals          *Totals  `json:"recent_ride_totals,omitempty"`
	RecentRunTotals           *Totals  `json:"recent_run_totals,omitempty"`
	RecentSwimTotals          *Totals  `json:"recent_swim_totals,omitempty"`
	YtdRideTotals             *Totals  `json:"ytd_ride_totals,omitempty"`
	YtdRunTotals              *Totals  `json:"ytd_run_totals,omitempty"`
	YtdSwimTotals             *Totals  `json:"ytd_swim_totals,omitempty"`
}

// Problem defines model for Problem.
type Problem struct {
	Detail string `json:"detail"`
	Status int    `json:"status"`
	Title  string `json:"title"`
}

// TotalDistance defines model for TotalDistance.
type TotalDistance struct {
	TotalDistance *float32 `json:"total_distance,omitempty"`
}

// Totals defines model for Totals.
type Totals struct {
	AchievementCount *int     `json:"achievement_count,omitempty"`
	Count            *int     `json:"count,omitempty"`
	Distance         *float32 `json:"distance,omitempty"`
	ElapsedTime      *int     `json:"elapsed_time,omitempty"`
	ElevationGain    *float32 `json:"elevation_gain,omitempty"`
	MovingTime       *int     `json:"moving_time,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get athlete stats
	// (GET /ahlete/stats)
	GetAthleteStats(w http.ResponseWriter, r *http.Request)
	// Get athlete
	// (GET /athlete)
	GetAthlete(w http.ResponseWriter, r *http.Request)
	// Get athlete stats totals
	// (GET /athlete/stats/totals)
	GetAthleteStatsTotals(w http.ResponseWriter, r *http.Request)
	// Get auth
	// (GET /auth)
	Auth(w http.ResponseWriter, r *http.Request)
	// Callback
	// (GET /callback)
	Callback(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Get athlete stats
// (GET /ahlete/stats)
func (_ Unimplemented) GetAthleteStats(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get athlete
// (GET /athlete)
func (_ Unimplemented) GetAthlete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get athlete stats totals
// (GET /athlete/stats/totals)
func (_ Unimplemented) GetAthleteStatsTotals(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get auth
// (GET /auth)
func (_ Unimplemented) Auth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Callback
// (GET /callback)
func (_ Unimplemented) Callback(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetAthleteStats operation middleware
func (siw *ServerInterfaceWrapper) GetAthleteStats(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAthleteStats(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetAthlete operation middleware
func (siw *ServerInterfaceWrapper) GetAthlete(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAthlete(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetAthleteStatsTotals operation middleware
func (siw *ServerInterfaceWrapper) GetAthleteStatsTotals(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAthleteStatsTotals(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// Auth operation middleware
func (siw *ServerInterfaceWrapper) Auth(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Auth(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// Callback operation middleware
func (siw *ServerInterfaceWrapper) Callback(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Callback(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/ahlete/stats", wrapper.GetAthleteStats)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/athlete", wrapper.GetAthlete)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/athlete/stats/totals", wrapper.GetAthleteStatsTotals)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/auth", wrapper.Auth)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/callback", wrapper.Callback)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xWT4/bthP9KgJ/v6MauckeAt0WKVAELZqg6S1YCCNyJDOhSJYceWME/u4FSdmybGo3",
	"MZqeTHP45s/jmxG/Mm4GazRq8qz+yjzf4gBxeU9bhYRhaZ2x6EhiNLQgemxob7GRImyEJauZ1IQ9OnYo",
	"WSvNmcGTk7oP+1zSPm8woya3YnMIhKIBCubOuCGsmADCn0gOyMprTCedJw0DZj12RinziC6ffOck6pXC",
	"1gpW8EQ463CQ43Bma41RCDoZTSfVGjDamgHFEj8fcejN6Dg2noAwn5zHL1nsJeTMMg6DpHzCoxXffR2j",
	"R7dKzyPKfnseTI9DGzI/nDyZ9hNyCocnVX4gSIJdShOUapwU2JAhUHHr/w47VrP/VbPQq0nl1V/p1KFM",
	"wFHfhPOPcvhuYCv7Hj01XMmhbVDhDkga3fQgdYaKGRDLE9ITaI7Zkw45arqNhyP2Biom6E1s7Enclm8E",
	"3pBswN2QaU6R751pFQ7XYhRIIFVY4RcYbOjxU7sW2lDRmVGLXLuExhz9Anm3uSszrU2S0uyYQ/yx7jne",
	"0t+jdChY/XECn8KVx4wfMlVGBn45k92y1kjjU7I8rPnMdTHfStzhEOQUPwz5sfaE6cn+QAXWo2jisMrC",
	"v6EdB7OTul/1cV1u2JK6M0kZnjtpQwRWsw/kYAfF/fu37HShy80dOp/O/vxi82IT4huLGqxkNXsVt0pm",
	"gbaRvgriiKz8cUb2GGkKFMeq3gpWs1+RFrM0fkqs0T7dwcvNJvxwowkTy2Ctkjw6qD55o+fnwnPds4gT",
	"iVgS8O63UNHd5u6am1nMh+mzBOGVELIvILkt/JQ/Qe+DrKd9zx4CpoL5HfMMEf8BB/9++c8WnnRQzYPu",
	"W+QwteYPJGQ5UH6QKgo61rHG0UjbVU7ug/GCglebl9f5/IlCOuQxXjF1bmdcEbyjpomgXLIpwim58Dcl",
	"xkGpFvjns+SWMd+AUiiKdn8MCB3hZciysM7spEBfAOfoQ4KfURegReGwc+i3aYeVF7W/OcbPS+D6qha1",
	"naEva4vvURcmGqs/Xjr63XBQhcAdKmPD/C/SWVay0SlWsy2RratKhXNb46l+vXm9YYeHU5xLj/ec5E7S",
	"vjiVF9SQ3qIMkjF8dg7lFXISUhZ4VNE17N3peJLA5OVCCrOnQMvh4fBPAAAA//8Uv0ueiQ0AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
