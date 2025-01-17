// Code generated by go-swagger; DO NOT EDIT.

package image_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new image config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for image config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateImage(params *CreateImageParams, authInfo runtime.ClientAuthInfoWriter) (*CreateImageNoContent, *CreateImageBadRequest, *CreateImageUnauthorized, *CreateImageConflict, *CreateImageInternalServerError, error)

	DeleteImage(params *DeleteImageParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteImageNoContent, *DeleteImageBadRequest, *DeleteImageUnauthorized, *DeleteImageNotFound, *DeleteImageUnprocessableEntity, *DeleteImageInternalServerError, error)

	ExportImages(params *ExportImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ExportImagesOK, *ExportImagesUnauthorized, *ExportImagesForbidden, *ExportImagesNotFound, *ExportImagesInternalServerError, error)

	GetImageDetail(params *GetImageDetailParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageDetailOK, *GetImageDetailUnauthorized, *GetImageDetailNotFound, *GetImageDetailInternalServerError, error)

	GetImageLimit(params *GetImageLimitParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageLimitOK, *GetImageLimitBadRequest, *GetImageLimitUnauthorized, *GetImageLimitInternalServerError, error)

	ImportImages(params *ImportImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ImportImagesOK, *ImportImagesBadRequest, *ImportImagesUnauthorized, *ImportImagesForbidden, *ImportImagesInternalServerError, error)

	ListImages(params *ListImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ListImagesOK, *ListImagesBadRequest, *ListImagesUnauthorized, *ListImagesInternalServerError, error)

	UpdateImage(params *UpdateImageParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateImageNoContent, *UpdateImageBadRequest, *UpdateImageUnauthorized, *UpdateImageInternalServerError, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateImage creates image

  ```
Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [CREATE]
Required scope: social

This endpoint will create image.

Sample image:
{
	"namespace":"dewa",
	"version":"1.0.0",
	"image":"144436415367.dkr.ecr.us-west-2.amazonaws.com/dewa:1.0.0",
	"persistent":false
}
```
*/
func (a *Client) CreateImage(params *CreateImageParams, authInfo runtime.ClientAuthInfoWriter) (*CreateImageNoContent, *CreateImageBadRequest, *CreateImageUnauthorized, *CreateImageConflict, *CreateImageInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateImageParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateImage",
		Method:             "POST",
		PathPattern:        "/dsmcontroller/admin/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateImageNoContent:
		return v, nil, nil, nil, nil, nil
	case *CreateImageBadRequest:
		return nil, v, nil, nil, nil, nil
	case *CreateImageUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *CreateImageConflict:
		return nil, nil, nil, v, nil, nil
	case *CreateImageInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteImage deletes an image

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [UPDATE]

Required scope: social

This endpoint will delete an image that specified in the request parameter
*/
func (a *Client) DeleteImage(params *DeleteImageParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteImageNoContent, *DeleteImageBadRequest, *DeleteImageUnauthorized, *DeleteImageNotFound, *DeleteImageUnprocessableEntity, *DeleteImageInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImageParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteImage",
		Method:             "DELETE",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteImageNoContent:
		return v, nil, nil, nil, nil, nil, nil
	case *DeleteImageBadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *DeleteImageUnauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *DeleteImageNotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *DeleteImageUnprocessableEntity:
		return nil, nil, nil, nil, v, nil, nil
	case *DeleteImageInternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ExportImages exports d s m controller images for a namespace

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [READ]

Required scope: social

This endpoint export a dedicated servers images in a namespace.

*/
func (a *Client) ExportImages(params *ExportImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ExportImagesOK, *ExportImagesUnauthorized, *ExportImagesForbidden, *ExportImagesNotFound, *ExportImagesInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportImagesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ExportImages",
		Method:             "GET",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/images/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExportImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *ExportImagesOK:
		return v, nil, nil, nil, nil, nil
	case *ExportImagesUnauthorized:
		return nil, v, nil, nil, nil, nil
	case *ExportImagesForbidden:
		return nil, nil, v, nil, nil, nil
	case *ExportImagesNotFound:
		return nil, nil, nil, v, nil, nil
	case *ExportImagesInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetImageDetail ds s image detail

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [READ]

Required scope: social

This endpoint get specific version of dedicated servers images.
*/
func (a *Client) GetImageDetail(params *GetImageDetailParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageDetailOK, *GetImageDetailUnauthorized, *GetImageDetailNotFound, *GetImageDetailInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageDetailParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetImageDetail",
		Method:             "GET",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/images/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetImageDetailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetImageDetailOK:
		return v, nil, nil, nil, nil
	case *GetImageDetailUnauthorized:
		return nil, v, nil, nil, nil
	case *GetImageDetailNotFound:
		return nil, nil, v, nil, nil
	case *GetImageDetailInternalServerError:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetImageLimit ds s image limit

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [READ]

Required scope: social

This endpoint get ds image limit for specific namespace
*/
func (a *Client) GetImageLimit(params *GetImageLimitParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageLimitOK, *GetImageLimitBadRequest, *GetImageLimitUnauthorized, *GetImageLimitInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageLimitParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetImageLimit",
		Method:             "GET",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/images/limit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetImageLimitReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetImageLimitOK:
		return v, nil, nil, nil, nil
	case *GetImageLimitBadRequest:
		return nil, v, nil, nil, nil
	case *GetImageLimitUnauthorized:
		return nil, nil, v, nil, nil
	case *GetImageLimitInternalServerError:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ImportImages imports images for a namespace

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [CREATE]

Required scope: social

This endpoint import a dedicated servers images in a namespace.

The image will be upsert. Existing version will be replaced with imported image, will create new one if not found.

Example data inside imported file
[
  {
	"namespace": "dewa",
	"image": "123456789.dkr.ecr.us-west-2.amazonaws.com/ds-dewa:0.0.1-alpha",
	"version": "0.0.1",
	"persistent": true
  }
]

*/
func (a *Client) ImportImages(params *ImportImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ImportImagesOK, *ImportImagesBadRequest, *ImportImagesUnauthorized, *ImportImagesForbidden, *ImportImagesInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportImagesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ImportImages",
		Method:             "POST",
		PathPattern:        "/dsmcontroller/admin/images/import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ImportImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *ImportImagesOK:
		return v, nil, nil, nil, nil, nil
	case *ImportImagesBadRequest:
		return nil, v, nil, nil, nil, nil
	case *ImportImagesUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *ImportImagesForbidden:
		return nil, nil, nil, v, nil, nil
	case *ImportImagesInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ListImages lists all d s images

  Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [READ]

Required scope: social

This endpoint lists all of dedicated servers images.
*/
func (a *Client) ListImages(params *ListImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ListImagesOK, *ListImagesBadRequest, *ListImagesUnauthorized, *ListImagesInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListImagesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListImages",
		Method:             "GET",
		PathPattern:        "/dsmcontroller/admin/namespaces/{namespace}/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *ListImagesOK:
		return v, nil, nil, nil, nil
	case *ListImagesBadRequest:
		return nil, v, nil, nil, nil
	case *ListImagesUnauthorized:
		return nil, nil, v, nil, nil
	case *ListImagesInternalServerError:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateImage updates image

  ```
Required permission: ADMIN:NAMESPACE:{namespace}:DSM:CONFIG [UPDATE]
Required scope: social

This endpoint will update an image name and/or image persistent flag.

Sample image:
{
	"namespace":"dewa",
	"version":"1.0.0",
	"image":"144436415367.dkr.ecr.us-west-2.amazonaws.com/dewa:1.0.0",
	"persistent":false
}
```
*/
func (a *Client) UpdateImage(params *UpdateImageParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateImageNoContent, *UpdateImageBadRequest, *UpdateImageUnauthorized, *UpdateImageInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateImageParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateImage",
		Method:             "PUT",
		PathPattern:        "/dsmcontroller/admin/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateImageNoContent:
		return v, nil, nil, nil, nil
	case *UpdateImageBadRequest:
		return nil, v, nil, nil, nil
	case *UpdateImageUnauthorized:
		return nil, nil, v, nil, nil
	case *UpdateImageInternalServerError:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
