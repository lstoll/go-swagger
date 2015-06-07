package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/casualjim/go-swagger/errors"
	"github.com/casualjim/go-swagger/middleware"
	"github.com/casualjim/go-swagger/strfmt"
)

// DeleteOrderParams contains all the bound params for the delete order operation
// typically these are obtained from a http.Request
type DeleteOrderParams struct {
	// ID of the order that needs to be deleted
	OrderID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteOrderParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	if err := o.bindOrderID(route.Params.Get("orderId"), route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteOrderParams) bindOrderID(raw string, formats strfmt.Registry) error {

	o.OrderID = raw

	if err := o.validateOrderID(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteOrderParams) validateOrderID(formats strfmt.Registry) error {

	return nil
}
