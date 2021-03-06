// Code generated by go-swagger; DO NOT EDIT.

package operations

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDescribeResourceParams creates a new DescribeResourceParams object
// with the default values initialized.
func NewDescribeResourceParams() *DescribeResourceParams {
	var (
		pageDefault     = int64(1)
		pageSizeDefault = int64(25)
	)
	return &DescribeResourceParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeResourceParamsWithTimeout creates a new DescribeResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeResourceParamsWithTimeout(timeout time.Duration) *DescribeResourceParams {
	var (
		pageDefault     = int64(1)
		pageSizeDefault = int64(25)
	)
	return &DescribeResourceParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,

		timeout: timeout,
	}
}

// NewDescribeResourceParamsWithContext creates a new DescribeResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeResourceParamsWithContext(ctx context.Context) *DescribeResourceParams {
	var (
		pageDefault     = int64(1)
		pageSizeDefault = int64(25)
	)
	return &DescribeResourceParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,

		Context: ctx,
	}
}

// NewDescribeResourceParamsWithHTTPClient creates a new DescribeResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeResourceParamsWithHTTPClient(client *http.Client) *DescribeResourceParams {
	var (
		pageDefault     = int64(1)
		pageSizeDefault = int64(25)
	)
	return &DescribeResourceParams{
		Page:       &pageDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*DescribeResourceParams contains all the parameters to send to the API endpoint
for the describe resource operation typically these are written to a http.Request
*/
type DescribeResourceParams struct {

	/*Page
	  Which page of results to retrieve

	*/
	Page *int64
	/*PageSize
	  Number of items in each page of results

	*/
	PageSize *int64
	/*ResourceID
	  URL-encoded unique resource ID

	*/
	ResourceID string
	/*Severity
	  Limit entries to those whose policies have this severity

	*/
	Severity *string
	/*Status
	  Limit entries to those with a specific compliance status

	*/
	Status *string
	/*Suppressed
	  Limit entries to those which are/are not suppressed

	*/
	Suppressed *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe resource params
func (o *DescribeResourceParams) WithTimeout(timeout time.Duration) *DescribeResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe resource params
func (o *DescribeResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe resource params
func (o *DescribeResourceParams) WithContext(ctx context.Context) *DescribeResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe resource params
func (o *DescribeResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe resource params
func (o *DescribeResourceParams) WithHTTPClient(client *http.Client) *DescribeResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe resource params
func (o *DescribeResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the describe resource params
func (o *DescribeResourceParams) WithPage(page *int64) *DescribeResourceParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the describe resource params
func (o *DescribeResourceParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the describe resource params
func (o *DescribeResourceParams) WithPageSize(pageSize *int64) *DescribeResourceParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the describe resource params
func (o *DescribeResourceParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithResourceID adds the resourceID to the describe resource params
func (o *DescribeResourceParams) WithResourceID(resourceID string) *DescribeResourceParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the describe resource params
func (o *DescribeResourceParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WithSeverity adds the severity to the describe resource params
func (o *DescribeResourceParams) WithSeverity(severity *string) *DescribeResourceParams {
	o.SetSeverity(severity)
	return o
}

// SetSeverity adds the severity to the describe resource params
func (o *DescribeResourceParams) SetSeverity(severity *string) {
	o.Severity = severity
}

// WithStatus adds the status to the describe resource params
func (o *DescribeResourceParams) WithStatus(status *string) *DescribeResourceParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe resource params
func (o *DescribeResourceParams) SetStatus(status *string) {
	o.Status = status
}

// WithSuppressed adds the suppressed to the describe resource params
func (o *DescribeResourceParams) WithSuppressed(suppressed *bool) *DescribeResourceParams {
	o.SetSuppressed(suppressed)
	return o
}

// SetSuppressed adds the suppressed to the describe resource params
func (o *DescribeResourceParams) SetSuppressed(suppressed *bool) {
	o.Suppressed = suppressed
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	// query param resourceId
	qrResourceID := o.ResourceID
	qResourceID := qrResourceID
	if qResourceID != "" {
		if err := r.SetQueryParam("resourceId", qResourceID); err != nil {
			return err
		}
	}

	if o.Severity != nil {

		// query param severity
		var qrSeverity string
		if o.Severity != nil {
			qrSeverity = *o.Severity
		}
		qSeverity := qrSeverity
		if qSeverity != "" {
			if err := r.SetQueryParam("severity", qSeverity); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Suppressed != nil {

		// query param suppressed
		var qrSuppressed bool
		if o.Suppressed != nil {
			qrSuppressed = *o.Suppressed
		}
		qSuppressed := swag.FormatBool(qrSuppressed)
		if qSuppressed != "" {
			if err := r.SetQueryParam("suppressed", qSuppressed); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
