// Code generated by go-swagger; DO NOT EDIT.

package immutable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new immutable API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for immutable API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateImmuRule(params *CreateImmuRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateImmuRuleCreated, error)

	DeleteImmuRule(params *DeleteImmuRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteImmuRuleOK, error)

	ListImmuRules(params *ListImmuRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListImmuRulesOK, error)

	UpdateImmuRule(params *UpdateImmuRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateImmuRuleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateImmuRule adds an immutable tag rule to current project

  This endpoint add an immutable tag rule to the project

*/
func (a *Client) CreateImmuRule(params *CreateImmuRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateImmuRuleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateImmuRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateImmuRule",
		Method:             "POST",
		PathPattern:        "/projects/{project_name_or_id}/immutabletagrules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateImmuRuleReader{formats: a.formats},
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
	success, ok := result.(*CreateImmuRuleCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateImmuRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteImmuRule deletes the immutable tag rule
*/
func (a *Client) DeleteImmuRule(params *DeleteImmuRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteImmuRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImmuRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteImmuRule",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteImmuRuleReader{formats: a.formats},
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
	success, ok := result.(*DeleteImmuRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteImmuRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListImmuRules lists all immutable tag rules of current project

  This endpoint returns the immutable tag rules of a project

*/
func (a *Client) ListImmuRules(params *ListImmuRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListImmuRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListImmuRulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListImmuRules",
		Method:             "GET",
		PathPattern:        "/projects/{project_name_or_id}/immutabletagrules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListImmuRulesReader{formats: a.formats},
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
	success, ok := result.(*ListImmuRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListImmuRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateImmuRule updates the immutable tag rule or enable or disable the rule
*/
func (a *Client) UpdateImmuRule(params *UpdateImmuRuleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateImmuRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateImmuRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateImmuRule",
		Method:             "PUT",
		PathPattern:        "/projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateImmuRuleReader{formats: a.formats},
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
	success, ok := result.(*UpdateImmuRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateImmuRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
