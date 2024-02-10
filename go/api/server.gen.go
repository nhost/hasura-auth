// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version (devel) DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /signin/email-password)
	PostSigninEmailPassword(c *gin.Context)
	// Signup with email and password
	// (POST /signup/email-password)
	PostSignupEmailPassword(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostSigninEmailPassword operation middleware
func (siw *ServerInterfaceWrapper) PostSigninEmailPassword(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostSigninEmailPassword(c)
}

// PostSignupEmailPassword operation middleware
func (siw *ServerInterfaceWrapper) PostSignupEmailPassword(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostSignupEmailPassword(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/signin/email-password", wrapper.PostSigninEmailPassword)
	router.POST(options.BaseURL+"/signup/email-password", wrapper.PostSignupEmailPassword)
}

type PostSigninEmailPasswordRequestObject struct {
	Body *PostSigninEmailPasswordJSONRequestBody
}

type PostSigninEmailPasswordResponseObject interface {
	VisitPostSigninEmailPasswordResponse(w http.ResponseWriter) error
}

type PostSigninEmailPassword200JSONResponse SessionPayload

func (response PostSigninEmailPassword200JSONResponse) VisitPostSigninEmailPasswordResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostSignupEmailPasswordRequestObject struct {
	Body *PostSignupEmailPasswordJSONRequestBody
}

type PostSignupEmailPasswordResponseObject interface {
	VisitPostSignupEmailPasswordResponse(w http.ResponseWriter) error
}

type PostSignupEmailPassword200JSONResponse SessionPayload

func (response PostSignupEmailPassword200JSONResponse) VisitPostSignupEmailPasswordResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /signin/email-password)
	PostSigninEmailPassword(ctx context.Context, request PostSigninEmailPasswordRequestObject) (PostSigninEmailPasswordResponseObject, error)
	// Signup with email and password
	// (POST /signup/email-password)
	PostSignupEmailPassword(ctx context.Context, request PostSignupEmailPasswordRequestObject) (PostSignupEmailPasswordResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// PostSigninEmailPassword operation middleware
func (sh *strictHandler) PostSigninEmailPassword(ctx *gin.Context) {
	var request PostSigninEmailPasswordRequestObject

	var body PostSigninEmailPasswordJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostSigninEmailPassword(ctx, request.(PostSigninEmailPasswordRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostSigninEmailPassword")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostSigninEmailPasswordResponseObject); ok {
		if err := validResponse.VisitPostSigninEmailPasswordResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostSignupEmailPassword operation middleware
func (sh *strictHandler) PostSignupEmailPassword(ctx *gin.Context) {
	var request PostSignupEmailPasswordRequestObject

	var body PostSignupEmailPasswordJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostSignupEmailPassword(ctx, request.(PostSignupEmailPasswordRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostSignupEmailPassword")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostSignupEmailPasswordResponseObject); ok {
		if err := validResponse.VisitPostSignupEmailPasswordResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RXX2/bOBL/KgR7h3uRLNfJ3TV+2uwiC6Ro06BOdh8aP4ylkcVWIrnkKI7h1XdfkJJl",
	"SZb7Byi6AfbJFuf//GaGwx2PVaGVREmWz3fcxhkW4P++/fXS/WijNBoS6A9JxJ+Q/L+tRj7nloyQa15V",
	"wf5ErT5iTLwK+AKtFUoea4E4Rmvv1CeUI6qCLv3qSQuD9tozpsoUQHzO01wB8dakLIsVGidpMDVos1Z1",
	"gjY2QpN3g7+vqYwcmSWls8egpAwliRgcF1OGbTKUrNHkOChD9vr3Ox5wfIJC587kLD777+p/6VkYn68u",
	"wvNXeBZe/P8VhMl5Mk1fJucznJ3zgGsgQuNsPzysPkzDCwjT5e5V9fCwCtvP8+rk/67Uy5kTO0R9yFdp",
	"0bho/2Uw5XP+IjqAGjWIRveOp/Ip+qMUBhM+/9AD4kTaBzldnsb5Fra5guQY7iKFL3nniq0KuD1UzOe4",
	"94U1XnViLa/lVQEivwVrN8okCy937Bg6puMyuWSPkIuE1eQu6h9VJie2EJT9JDNlaSIUDw5luRc4Akg3",
	"jnijrbYFmalceydfKJOEF+d//vtYegDZ3karcnkiB/f6eeVAedV1/yeJcB+Q33Z8SSG3GAxHRZ6rDSbv",
	"VV5/J5hCmZPLRIG8Kf1lx7/BuSAsvKCS+C7l8w87jrIsGr5lFRy+a5FqeSgqMAa2zvfGqvOi50QtE3RV",
	"jri05zpKSSKszmF7AwX2K+O1yiRbuByPieUqhqEjvoGHENJGhXEGBmJCY1kj1/XMixXw9AblmjI+nwW8",
	"ELLzdWS8QIIECL6IY+vbrupY3PFUGEt1yD5OHvAc2pM66LG+NpgIgzHdqX7kGZG28ygqtiFoPYlV0Yvw",
	"mBzFQHEW7vW5ZAVfc511m3iY6T2NqZQVQoqiLNgZO+S+59KJvu8k/uy7TIH75lr4lnaLSTzi2xTuvLId",
	"B7kdtE3HVFsVByop0iMcy2Fxvi1zEmEKMSkzvIKd7IRdMlnmuRtDJbIEvWdAaJm7K7rp9Cbd1vAIBObe",
	"5KM7RWwQCJNL6i0SCRCGJHzfHrfns+r6HzaoPeE3NCIVmPQCb0qmkVgplSNIJyJGmuLaN4Pbnfbpep7L",
	"k7CXUsltoUr7dcH+M8evzpTEm3rRHuuvDv3basf8nVf7YKwehkS/TbvDpS2AoG2hbgn1x0YHsGAwXIeN",
	"1k/xeEL32Tqe9i4SIVN/N8ZKEsTU2e64LbVWhrqDQNaI37iT/1i2qDlcft0Ebe/NVqIaFvgCzaOIkZFi",
	"l4cB7gLLRYzS+nJrrFxqiDNks8n0yMBms5mAJ0+UWUeNrI3eXP9ydbO4CmeT6SSjIvdViaaw79LGcqNk",
	"HkV2A+s1molQkWeJXKMLytsAvYc84I9o6gcGfzmZTqb1XooStOBzfuaP/NzJfGFFVqylkJHPYti9/7Wy",
	"dDzy3NrNhGQgE4ZPwpJ7PbraY6V1f70eT211efvG33vXCZ/zW2Vp4a32tnde1yla+lkl2z3IKL0ToHXe",
	"3J3RR1s/n2y77n/2IXXyrVT1W4NMif7AauXQcYpn0+n3c6T/hPTWx3O7Acts6V+qaZnXmxqsreveGq19",
	"Y4UgkwNmS8fo8Sz1Z/AcB6PUPwqM8Ufb8wOjBSDfMoNrYQkNJhN241a25hXPCgRpm6J/9COsWfCEZRpl",
	"sl9sbVkUYLYNyKVmG0HZeLN0oS71Sair6q8AAAD//yvHCX/ZEgAA",
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
