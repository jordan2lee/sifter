// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// PostPlaybookIDHandlerFunc turns a function with the right signature into a post playbook ID handler
type PostPlaybookIDHandlerFunc func(PostPlaybookIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostPlaybookIDHandlerFunc) Handle(params PostPlaybookIDParams) middleware.Responder {
	return fn(params)
}

// PostPlaybookIDHandler interface for that can handle valid post playbook ID params
type PostPlaybookIDHandler interface {
	Handle(PostPlaybookIDParams) middleware.Responder
}

// NewPostPlaybookID creates a new http.Handler for the post playbook ID operation
func NewPostPlaybookID(ctx *middleware.Context, handler PostPlaybookIDHandler) *PostPlaybookID {
	return &PostPlaybookID{Context: ctx, Handler: handler}
}

/*PostPlaybookID swagger:route POST /playbook/{id} postPlaybookId

Post a run request

*/
type PostPlaybookID struct {
	Context *middleware.Context
	Handler PostPlaybookIDHandler
}

func (o *PostPlaybookID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostPlaybookIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostPlaybookIDOKBody post playbook ID o k body
// swagger:model PostPlaybookIDOKBody
type PostPlaybookIDOKBody struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this post playbook ID o k body
func (o *PostPlaybookIDOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostPlaybookIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPlaybookIDOKBody) UnmarshalBinary(b []byte) error {
	var res PostPlaybookIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}