// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	models "github.com/MontFerret/ferret-server/server/http/api/models"
)

// NewFerretServerAPI creates a new FerretServer instance
func NewFerretServerAPI(spec *loads.Document) *FerretServerAPI {
	return &FerretServerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		CreateExecutionHandler: CreateExecutionHandlerFunc(func(params CreateExecutionParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateExecution has not yet been implemented")
		}),
		CreateProjectHandler: CreateProjectHandlerFunc(func(params CreateProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateProject has not yet been implemented")
		}),
		CreateScriptHandler: CreateScriptHandlerFunc(func(params CreateScriptParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateScript has not yet been implemented")
		}),
		CreateUserHandler: CreateUserHandlerFunc(func(params CreateUserParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation CreateUser has not yet been implemented")
		}),
		DeleteExecutionHandler: DeleteExecutionHandlerFunc(func(params DeleteExecutionParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteExecution has not yet been implemented")
		}),
		DeleteProjectHandler: DeleteProjectHandlerFunc(func(params DeleteProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteProject has not yet been implemented")
		}),
		DeleteScriptHandler: DeleteScriptHandlerFunc(func(params DeleteScriptParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteScript has not yet been implemented")
		}),
		DeleteScriptDataHandler: DeleteScriptDataHandlerFunc(func(params DeleteScriptDataParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteScriptData has not yet been implemented")
		}),
		FindExecutionsHandler: FindExecutionsHandlerFunc(func(params FindExecutionsParams) middleware.Responder {
			return middleware.NotImplemented("operation FindExecutions has not yet been implemented")
		}),
		FindProjectDataHandler: FindProjectDataHandlerFunc(func(params FindProjectDataParams) middleware.Responder {
			return middleware.NotImplemented("operation FindProjectData has not yet been implemented")
		}),
		FindProjectsHandler: FindProjectsHandlerFunc(func(params FindProjectsParams) middleware.Responder {
			return middleware.NotImplemented("operation FindProjects has not yet been implemented")
		}),
		FindScriptDataHandler: FindScriptDataHandlerFunc(func(params FindScriptDataParams) middleware.Responder {
			return middleware.NotImplemented("operation FindScriptData has not yet been implemented")
		}),
		FindScriptsHandler: FindScriptsHandlerFunc(func(params FindScriptsParams) middleware.Responder {
			return middleware.NotImplemented("operation FindScripts has not yet been implemented")
		}),
		GetExecutionHandler: GetExecutionHandlerFunc(func(params GetExecutionParams) middleware.Responder {
			return middleware.NotImplemented("operation GetExecution has not yet been implemented")
		}),
		GetProjectHandler: GetProjectHandlerFunc(func(params GetProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation GetProject has not yet been implemented")
		}),
		GetScriptHandler: GetScriptHandlerFunc(func(params GetScriptParams) middleware.Responder {
			return middleware.NotImplemented("operation GetScript has not yet been implemented")
		}),
		GetScriptDataHandler: GetScriptDataHandlerFunc(func(params GetScriptDataParams) middleware.Responder {
			return middleware.NotImplemented("operation GetScriptData has not yet been implemented")
		}),
		TokenByCredentialsHandler: TokenByCredentialsHandlerFunc(func(params TokenByCredentialsParams) middleware.Responder {
			return middleware.NotImplemented("operation TokenByCredentials has not yet been implemented")
		}),
		UpdateProjectHandler: UpdateProjectHandlerFunc(func(params UpdateProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateProject has not yet been implemented")
		}),
		UpdateScriptHandler: UpdateScriptHandlerFunc(func(params UpdateScriptParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateScript has not yet been implemented")
		}),
		UpdateScriptDataHandler: UpdateScriptDataHandlerFunc(func(params UpdateScriptDataParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateScriptData has not yet been implemented")
		}),

		// Applies when the "x-auth-token" header is set
		XAuthTokenAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (X-Auth-Token) x-auth-token from header param [x-auth-token] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*FerretServerAPI API of Ferret Server */
type FerretServerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// XAuthTokenAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key x-auth-token provided in the header
	XAuthTokenAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// CreateExecutionHandler sets the operation handler for the create execution operation
	CreateExecutionHandler CreateExecutionHandler
	// CreateProjectHandler sets the operation handler for the create project operation
	CreateProjectHandler CreateProjectHandler
	// CreateScriptHandler sets the operation handler for the create script operation
	CreateScriptHandler CreateScriptHandler
	// CreateUserHandler sets the operation handler for the create user operation
	CreateUserHandler CreateUserHandler
	// DeleteExecutionHandler sets the operation handler for the delete execution operation
	DeleteExecutionHandler DeleteExecutionHandler
	// DeleteProjectHandler sets the operation handler for the delete project operation
	DeleteProjectHandler DeleteProjectHandler
	// DeleteScriptHandler sets the operation handler for the delete script operation
	DeleteScriptHandler DeleteScriptHandler
	// DeleteScriptDataHandler sets the operation handler for the delete script data operation
	DeleteScriptDataHandler DeleteScriptDataHandler
	// FindExecutionsHandler sets the operation handler for the find executions operation
	FindExecutionsHandler FindExecutionsHandler
	// FindProjectDataHandler sets the operation handler for the find project data operation
	FindProjectDataHandler FindProjectDataHandler
	// FindProjectsHandler sets the operation handler for the find projects operation
	FindProjectsHandler FindProjectsHandler
	// FindScriptDataHandler sets the operation handler for the find script data operation
	FindScriptDataHandler FindScriptDataHandler
	// FindScriptsHandler sets the operation handler for the find scripts operation
	FindScriptsHandler FindScriptsHandler
	// GetExecutionHandler sets the operation handler for the get execution operation
	GetExecutionHandler GetExecutionHandler
	// GetProjectHandler sets the operation handler for the get project operation
	GetProjectHandler GetProjectHandler
	// GetScriptHandler sets the operation handler for the get script operation
	GetScriptHandler GetScriptHandler
	// GetScriptDataHandler sets the operation handler for the get script data operation
	GetScriptDataHandler GetScriptDataHandler
	// TokenByCredentialsHandler sets the operation handler for the token by credentials operation
	TokenByCredentialsHandler TokenByCredentialsHandler
	// UpdateProjectHandler sets the operation handler for the update project operation
	UpdateProjectHandler UpdateProjectHandler
	// UpdateScriptHandler sets the operation handler for the update script operation
	UpdateScriptHandler UpdateScriptHandler
	// UpdateScriptDataHandler sets the operation handler for the update script data operation
	UpdateScriptDataHandler UpdateScriptDataHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *FerretServerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *FerretServerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *FerretServerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *FerretServerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *FerretServerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *FerretServerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *FerretServerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the FerretServerAPI
func (o *FerretServerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.XAuthTokenAuth == nil {
		unregistered = append(unregistered, "XAuthTokenAuth")
	}

	if o.CreateExecutionHandler == nil {
		unregistered = append(unregistered, "CreateExecutionHandler")
	}

	if o.CreateProjectHandler == nil {
		unregistered = append(unregistered, "CreateProjectHandler")
	}

	if o.CreateScriptHandler == nil {
		unregistered = append(unregistered, "CreateScriptHandler")
	}

	if o.CreateUserHandler == nil {
		unregistered = append(unregistered, "CreateUserHandler")
	}

	if o.DeleteExecutionHandler == nil {
		unregistered = append(unregistered, "DeleteExecutionHandler")
	}

	if o.DeleteProjectHandler == nil {
		unregistered = append(unregistered, "DeleteProjectHandler")
	}

	if o.DeleteScriptHandler == nil {
		unregistered = append(unregistered, "DeleteScriptHandler")
	}

	if o.DeleteScriptDataHandler == nil {
		unregistered = append(unregistered, "DeleteScriptDataHandler")
	}

	if o.FindExecutionsHandler == nil {
		unregistered = append(unregistered, "FindExecutionsHandler")
	}

	if o.FindProjectDataHandler == nil {
		unregistered = append(unregistered, "FindProjectDataHandler")
	}

	if o.FindProjectsHandler == nil {
		unregistered = append(unregistered, "FindProjectsHandler")
	}

	if o.FindScriptDataHandler == nil {
		unregistered = append(unregistered, "FindScriptDataHandler")
	}

	if o.FindScriptsHandler == nil {
		unregistered = append(unregistered, "FindScriptsHandler")
	}

	if o.GetExecutionHandler == nil {
		unregistered = append(unregistered, "GetExecutionHandler")
	}

	if o.GetProjectHandler == nil {
		unregistered = append(unregistered, "GetProjectHandler")
	}

	if o.GetScriptHandler == nil {
		unregistered = append(unregistered, "GetScriptHandler")
	}

	if o.GetScriptDataHandler == nil {
		unregistered = append(unregistered, "GetScriptDataHandler")
	}

	if o.TokenByCredentialsHandler == nil {
		unregistered = append(unregistered, "TokenByCredentialsHandler")
	}

	if o.UpdateProjectHandler == nil {
		unregistered = append(unregistered, "UpdateProjectHandler")
	}

	if o.UpdateScriptHandler == nil {
		unregistered = append(unregistered, "UpdateScriptHandler")
	}

	if o.UpdateScriptDataHandler == nil {
		unregistered = append(unregistered, "UpdateScriptDataHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *FerretServerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *FerretServerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "X-Auth-Token":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.XAuthTokenAuth(token)
			})

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *FerretServerAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *FerretServerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *FerretServerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *FerretServerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the ferret server API
func (o *FerretServerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *FerretServerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/projects/{projectID}/exec"] = NewCreateExecution(o.context, o.CreateExecutionHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/projects"] = NewCreateProject(o.context, o.CreateProjectHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/projects/{projectID}/scripts"] = NewCreateScript(o.context, o.CreateScriptHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users"] = NewCreateUser(o.context, o.CreateUserHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/projects/{projectID}/exec/{jobID}"] = NewDeleteExecution(o.context, o.DeleteExecutionHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/projects/{projectID}"] = NewDeleteProject(o.context, o.DeleteProjectHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/projects/{projectID}/scripts/{scriptID}"] = NewDeleteScript(o.context, o.DeleteScriptHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/projects/{projectID}/data/{scriptID}/{dataId}"] = NewDeleteScriptData(o.context, o.DeleteScriptDataHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectID}/exec"] = NewFindExecutions(o.context, o.FindExecutionsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectID}/data"] = NewFindProjectData(o.context, o.FindProjectDataHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects"] = NewFindProjects(o.context, o.FindProjectsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectID}/data/{scriptID}"] = NewFindScriptData(o.context, o.FindScriptDataHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectID}/scripts"] = NewFindScripts(o.context, o.FindScriptsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectID}/exec/{jobID}"] = NewGetExecution(o.context, o.GetExecutionHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectID}"] = NewGetProject(o.context, o.GetProjectHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectID}/scripts/{scriptID}"] = NewGetScript(o.context, o.GetScriptHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectID}/data/{scriptID}/{dataId}"] = NewGetScriptData(o.context, o.GetScriptDataHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/credentials"] = NewTokenByCredentials(o.context, o.TokenByCredentialsHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/projects/{projectID}"] = NewUpdateProject(o.context, o.UpdateProjectHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/projects/{projectID}/scripts/{scriptID}"] = NewUpdateScript(o.context, o.UpdateScriptHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/projects/{projectID}/data/{scriptID}/{dataId}"] = NewUpdateScriptData(o.context, o.UpdateScriptDataHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *FerretServerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *FerretServerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *FerretServerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *FerretServerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
