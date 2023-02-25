// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCollections(params *GetCollectionsParams, opts ...ClientOption) (*GetCollectionsOK, error)

	GetCollectionsOuID(params *GetCollectionsOuIDParams, opts ...ClientOption) (*GetCollectionsOuIDOK, error)

	GetCollectionsProjectID(params *GetCollectionsProjectIDParams, opts ...ClientOption) (*GetCollectionsProjectIDOK, error)

	GetPing(params *GetPingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPingOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetCollections gets metadata of all collections
*/
func (a *Client) GetCollections(params *GetCollectionsParams, opts ...ClientOption) (*GetCollectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCollectionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCollections",
		Method:             "GET",
		PathPattern:        "/collections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCollectionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCollectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCollections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCollectionsOuID gets metadata of collections belonging to the organizational unit
*/
func (a *Client) GetCollectionsOuID(params *GetCollectionsOuIDParams, opts ...ClientOption) (*GetCollectionsOuIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCollectionsOuIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCollectionsOuID",
		Method:             "GET",
		PathPattern:        "/collections/ou/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCollectionsOuIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCollectionsOuIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCollectionsOuID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCollectionsProjectID gets metadata of collections associated with the project id
*/
func (a *Client) GetCollectionsProjectID(params *GetCollectionsProjectIDParams, opts ...ClientOption) (*GetCollectionsProjectIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCollectionsProjectIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCollectionsProjectID",
		Method:             "GET",
		PathPattern:        "/collections/project/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCollectionsProjectIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCollectionsProjectIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCollectionsProjectID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPing endpoints for API server health check
*/
func (a *Client) GetPing(params *GetPingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPing",
		Method:             "GET",
		PathPattern:        "/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPing: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
