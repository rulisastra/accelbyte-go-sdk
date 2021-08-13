// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new session API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for session API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ClaimServer(params *ClaimServerParams, authInfo runtime.ClientAuthInfoWriter) (*ClaimServerNoContent, error)

	CreateSession(params *CreateSessionParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSessionOK, error)

	GetSession(params *GetSessionParams, authInfo runtime.ClientAuthInfoWriter) (*GetSessionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ClaimServer claims a d s for a game session

  Required permission: NAMESPACE:{namespace}:DSM:SESSION [UPDATE]

Required scope: social

This endpoint is intended to be called by game session manager (matchmaker, lobby, etc.) to claim a dedicated server. The dedicated server cannot be claimed unless the status is READY
*/
func (a *Client) ClaimServer(params *ClaimServerParams, authInfo runtime.ClientAuthInfoWriter) (*ClaimServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClaimServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ClaimServer",
		Method:             "POST",
		PathPattern:        "/dsmcontroller/namespaces/{namespace}/sessions/claim",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ClaimServerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClaimServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ClaimServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateSession registers a new game session

  Required permission: NAMESPACE:{namespace}:DSM:SESSION [CREATE]

Required scope: social

This endpoint is intended to be called by game session manager (matchmaker, lobby, etc.) to get a dedicated server for a game session.

If a dedicated server is available, it will respond with a dedicated server details ready to be used.

Otherwise it will trigger new dedicated server creation and respond with a server status CREATING. The game session manager then expected to wait and query the server readiness with GET /namespaces/{namespace}/sessions/{sessionID} endpoint until the serverstatus is READY

Specify pod_name with name of local DS in the request to create a session using the registered local DS
*/
func (a *Client) CreateSession(params *CreateSessionParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateSession",
		Method:             "POST",
		PathPattern:        "/dsmcontroller/namespaces/{namespace}/sessions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSessionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSessionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSession queries specified session

  Required permission: NAMESPACE:{namespace}:DSM:SESSION [READ]

Required scope: social

This endpoint is intended to be called by game session manager (matchmaker, lobby, etc.) to query the status of dedicated server that is created for the session.

The server is ready to use when the status is READY. At which point, the game session manager can claim the server using the GET /namespaces/{namespace}/sessions/{sessionID}/claim endpoint
*/
func (a *Client) GetSession(params *GetSessionParams, authInfo runtime.ClientAuthInfoWriter) (*GetSessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSession",
		Method:             "GET",
		PathPattern:        "/dsmcontroller/namespaces/{namespace}/sessions/{sessionID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSessionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSessionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}