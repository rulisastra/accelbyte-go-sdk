// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewQuerySessionParams creates a new QuerySessionParams object
// with the default values initialized.
func NewQuerySessionParams() *QuerySessionParams {
	var ()
	return &QuerySessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuerySessionParamsWithTimeout creates a new QuerySessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuerySessionParamsWithTimeout(timeout time.Duration) *QuerySessionParams {
	var ()
	return &QuerySessionParams{

		timeout: timeout,
	}
}

// NewQuerySessionParamsWithContext creates a new QuerySessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuerySessionParamsWithContext(ctx context.Context) *QuerySessionParams {
	var ()
	return &QuerySessionParams{

		Context: ctx,
	}
}

// NewQuerySessionParamsWithHTTPClient creates a new QuerySessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuerySessionParamsWithHTTPClient(client *http.Client) *QuerySessionParams {
	var ()
	return &QuerySessionParams{
		HTTPClient: client,
	}
}

/*QuerySessionParams contains all the parameters to send to the API endpoint
for the query session operation typically these are written to a http.Request
*/
type QuerySessionParams struct {

	/*GameMode
	  Game mode of the session to query

	*/
	GameMode *string
	/*GameVersion
	  Game version of the session to query

	*/
	GameVersion *string
	/*Joinable
	  filter session if joinable or not, accept 'true' or 'false'

	*/
	Joinable *string
	/*Limit
	  max item to be returned

	*/
	Limit *string
	/*MatchID
	  filter session by matchmaking ID

	*/
	MatchID *string
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Offset
	  skip some item(s), offset=3, will skip 3 first items

	*/
	Offset *string
	/*ServerStatus
	  filter session by server status, accept CREATING,READY,BUSY,REMOVING,UNREACHABLE

	*/
	ServerStatus *string
	/*SessionType
	  Session type to query, value is 'p2p' or 'dedicated'

	*/
	SessionType string
	/*UserID
	  Owner of the P2P session. Dedicated server does not have user_id

	*/
	UserID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query session params
func (o *QuerySessionParams) WithTimeout(timeout time.Duration) *QuerySessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query session params
func (o *QuerySessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query session params
func (o *QuerySessionParams) WithContext(ctx context.Context) *QuerySessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query session params
func (o *QuerySessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query session params
func (o *QuerySessionParams) WithHTTPClient(client *http.Client) *QuerySessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query session params
func (o *QuerySessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGameMode adds the gameMode to the query session params
func (o *QuerySessionParams) WithGameMode(gameMode *string) *QuerySessionParams {
	o.SetGameMode(gameMode)
	return o
}

// SetGameMode adds the gameMode to the query session params
func (o *QuerySessionParams) SetGameMode(gameMode *string) {
	o.GameMode = gameMode
}

// WithGameVersion adds the gameVersion to the query session params
func (o *QuerySessionParams) WithGameVersion(gameVersion *string) *QuerySessionParams {
	o.SetGameVersion(gameVersion)
	return o
}

// SetGameVersion adds the gameVersion to the query session params
func (o *QuerySessionParams) SetGameVersion(gameVersion *string) {
	o.GameVersion = gameVersion
}

// WithJoinable adds the joinable to the query session params
func (o *QuerySessionParams) WithJoinable(joinable *string) *QuerySessionParams {
	o.SetJoinable(joinable)
	return o
}

// SetJoinable adds the joinable to the query session params
func (o *QuerySessionParams) SetJoinable(joinable *string) {
	o.Joinable = joinable
}

// WithLimit adds the limit to the query session params
func (o *QuerySessionParams) WithLimit(limit *string) *QuerySessionParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query session params
func (o *QuerySessionParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithMatchID adds the matchID to the query session params
func (o *QuerySessionParams) WithMatchID(matchID *string) *QuerySessionParams {
	o.SetMatchID(matchID)
	return o
}

// SetMatchID adds the matchId to the query session params
func (o *QuerySessionParams) SetMatchID(matchID *string) {
	o.MatchID = matchID
}

// WithNamespace adds the namespace to the query session params
func (o *QuerySessionParams) WithNamespace(namespace string) *QuerySessionParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the query session params
func (o *QuerySessionParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the query session params
func (o *QuerySessionParams) WithOffset(offset *string) *QuerySessionParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query session params
func (o *QuerySessionParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithServerStatus adds the serverStatus to the query session params
func (o *QuerySessionParams) WithServerStatus(serverStatus *string) *QuerySessionParams {
	o.SetServerStatus(serverStatus)
	return o
}

// SetServerStatus adds the serverStatus to the query session params
func (o *QuerySessionParams) SetServerStatus(serverStatus *string) {
	o.ServerStatus = serverStatus
}

// WithSessionType adds the sessionType to the query session params
func (o *QuerySessionParams) WithSessionType(sessionType string) *QuerySessionParams {
	o.SetSessionType(sessionType)
	return o
}

// SetSessionType adds the sessionType to the query session params
func (o *QuerySessionParams) SetSessionType(sessionType string) {
	o.SessionType = sessionType
}

// WithUserID adds the userID to the query session params
func (o *QuerySessionParams) WithUserID(userID *string) *QuerySessionParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the query session params
func (o *QuerySessionParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *QuerySessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GameMode != nil {

		// query param game_mode
		var qrGameMode string
		if o.GameMode != nil {
			qrGameMode = *o.GameMode
		}
		qGameMode := qrGameMode
		if qGameMode != "" {
			if err := r.SetQueryParam("game_mode", qGameMode); err != nil {
				return err
			}
		}

	}

	if o.GameVersion != nil {

		// query param game_version
		var qrGameVersion string
		if o.GameVersion != nil {
			qrGameVersion = *o.GameVersion
		}
		qGameVersion := qrGameVersion
		if qGameVersion != "" {
			if err := r.SetQueryParam("game_version", qGameVersion); err != nil {
				return err
			}
		}

	}

	if o.Joinable != nil {

		// query param joinable
		var qrJoinable string
		if o.Joinable != nil {
			qrJoinable = *o.Joinable
		}
		qJoinable := qrJoinable
		if qJoinable != "" {
			if err := r.SetQueryParam("joinable", qJoinable); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.MatchID != nil {

		// query param match_id
		var qrMatchID string
		if o.MatchID != nil {
			qrMatchID = *o.MatchID
		}
		qMatchID := qrMatchID
		if qMatchID != "" {
			if err := r.SetQueryParam("match_id", qMatchID); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.ServerStatus != nil {

		// query param server_status
		var qrServerStatus string
		if o.ServerStatus != nil {
			qrServerStatus = *o.ServerStatus
		}
		qServerStatus := qrServerStatus
		if qServerStatus != "" {
			if err := r.SetQueryParam("server_status", qServerStatus); err != nil {
				return err
			}
		}

	}

	// query param session_type
	qrSessionType := o.SessionType
	qSessionType := qrSessionType
	if qSessionType != "" {
		if err := r.SetQueryParam("session_type", qSessionType); err != nil {
			return err
		}
	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {
			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}