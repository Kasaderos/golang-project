// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/loms/v1/service.proto

package loms

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on OrderCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrderCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderCreateRequestMultiError, or nil if none found.
func (m *OrderCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() <= 0 {
		err := OrderCreateRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OrderCreateRequestValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OrderCreateRequestValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrderCreateRequestValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OrderCreateRequestMultiError(errors)
	}

	return nil
}

// OrderCreateRequestMultiError is an error wrapping multiple validation errors
// returned by OrderCreateRequest.ValidateAll() if the designated constraints
// aren't met.
type OrderCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderCreateRequestMultiError) AllErrors() []error { return m }

// OrderCreateRequestValidationError is the validation error returned by
// OrderCreateRequest.Validate if the designated constraints aren't met.
type OrderCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderCreateRequestValidationError) ErrorName() string {
	return "OrderCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OrderCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderCreateRequestValidationError{}

// Validate checks the field values on OrderInfoItem with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderInfoItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderInfoItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderInfoItemMultiError, or
// nil if none found.
func (m *OrderInfoItem) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderInfoItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sku

	// no validation rules for Count

	if len(errors) > 0 {
		return OrderInfoItemMultiError(errors)
	}

	return nil
}

// OrderInfoItemMultiError is an error wrapping multiple validation errors
// returned by OrderInfoItem.ValidateAll() if the designated constraints
// aren't met.
type OrderInfoItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderInfoItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderInfoItemMultiError) AllErrors() []error { return m }

// OrderInfoItemValidationError is the validation error returned by
// OrderInfoItem.Validate if the designated constraints aren't met.
type OrderInfoItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderInfoItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderInfoItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderInfoItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderInfoItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderInfoItemValidationError) ErrorName() string { return "OrderInfoItemValidationError" }

// Error satisfies the builtin error interface
func (e OrderInfoItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderInfoItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderInfoItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderInfoItemValidationError{}

// Validate checks the field values on OrderCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrderCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderCreateResponseMultiError, or nil if none found.
func (m *OrderCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return OrderCreateResponseMultiError(errors)
	}

	return nil
}

// OrderCreateResponseMultiError is an error wrapping multiple validation
// errors returned by OrderCreateResponse.ValidateAll() if the designated
// constraints aren't met.
type OrderCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderCreateResponseMultiError) AllErrors() []error { return m }

// OrderCreateResponseValidationError is the validation error returned by
// OrderCreateResponse.Validate if the designated constraints aren't met.
type OrderCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderCreateResponseValidationError) ErrorName() string {
	return "OrderCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OrderCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderCreateResponseValidationError{}

// Validate checks the field values on GetStockInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetStockInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStockInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStockInfoRequestMultiError, or nil if none found.
func (m *GetStockInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStockInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetSku() <= 0 {
		err := GetStockInfoRequestValidationError{
			field:  "Sku",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetStockInfoRequestMultiError(errors)
	}

	return nil
}

// GetStockInfoRequestMultiError is an error wrapping multiple validation
// errors returned by GetStockInfoRequest.ValidateAll() if the designated
// constraints aren't met.
type GetStockInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStockInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStockInfoRequestMultiError) AllErrors() []error { return m }

// GetStockInfoRequestValidationError is the validation error returned by
// GetStockInfoRequest.Validate if the designated constraints aren't met.
type GetStockInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStockInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStockInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStockInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStockInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStockInfoRequestValidationError) ErrorName() string {
	return "GetStockInfoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetStockInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStockInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStockInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStockInfoRequestValidationError{}

// Validate checks the field values on GetStockInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetStockInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStockInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStockInfoResponseMultiError, or nil if none found.
func (m *GetStockInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStockInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Count

	if len(errors) > 0 {
		return GetStockInfoResponseMultiError(errors)
	}

	return nil
}

// GetStockInfoResponseMultiError is an error wrapping multiple validation
// errors returned by GetStockInfoResponse.ValidateAll() if the designated
// constraints aren't met.
type GetStockInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStockInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStockInfoResponseMultiError) AllErrors() []error { return m }

// GetStockInfoResponseValidationError is the validation error returned by
// GetStockInfoResponse.Validate if the designated constraints aren't met.
type GetStockInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStockInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStockInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStockInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStockInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStockInfoResponseValidationError) ErrorName() string {
	return "GetStockInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetStockInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStockInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStockInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStockInfoResponseValidationError{}

// Validate checks the field values on CreateOrderErrorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderErrorResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderErrorResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderErrorResponseMultiError, or nil if none found.
func (m *CreateOrderErrorResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderErrorResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return CreateOrderErrorResponseMultiError(errors)
	}

	return nil
}

// CreateOrderErrorResponseMultiError is an error wrapping multiple validation
// errors returned by CreateOrderErrorResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateOrderErrorResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderErrorResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderErrorResponseMultiError) AllErrors() []error { return m }

// CreateOrderErrorResponseValidationError is the validation error returned by
// CreateOrderErrorResponse.Validate if the designated constraints aren't met.
type CreateOrderErrorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderErrorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderErrorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderErrorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderErrorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderErrorResponseValidationError) ErrorName() string {
	return "CreateOrderErrorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderErrorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderErrorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderErrorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderErrorResponseValidationError{}

// Validate checks the field values on GetStockInfoErrorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetStockInfoErrorResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStockInfoErrorResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStockInfoErrorResponseMultiError, or nil if none found.
func (m *GetStockInfoErrorResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStockInfoErrorResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return GetStockInfoErrorResponseMultiError(errors)
	}

	return nil
}

// GetStockInfoErrorResponseMultiError is an error wrapping multiple validation
// errors returned by GetStockInfoErrorResponse.ValidateAll() if the
// designated constraints aren't met.
type GetStockInfoErrorResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStockInfoErrorResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStockInfoErrorResponseMultiError) AllErrors() []error { return m }

// GetStockInfoErrorResponseValidationError is the validation error returned by
// GetStockInfoErrorResponse.Validate if the designated constraints aren't met.
type GetStockInfoErrorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStockInfoErrorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStockInfoErrorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStockInfoErrorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStockInfoErrorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStockInfoErrorResponseValidationError) ErrorName() string {
	return "GetStockInfoErrorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetStockInfoErrorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStockInfoErrorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStockInfoErrorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStockInfoErrorResponseValidationError{}

// Validate checks the field values on CancelOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CancelOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CancelOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CancelOrderRequestMultiError, or nil if none found.
func (m *CancelOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CancelOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetOrderId() <= 0 {
		err := CancelOrderRequestValidationError{
			field:  "OrderId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CancelOrderRequestMultiError(errors)
	}

	return nil
}

// CancelOrderRequestMultiError is an error wrapping multiple validation errors
// returned by CancelOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type CancelOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CancelOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CancelOrderRequestMultiError) AllErrors() []error { return m }

// CancelOrderRequestValidationError is the validation error returned by
// CancelOrderRequest.Validate if the designated constraints aren't met.
type CancelOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CancelOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CancelOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CancelOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CancelOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CancelOrderRequestValidationError) ErrorName() string {
	return "CancelOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CancelOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCancelOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CancelOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CancelOrderRequestValidationError{}

// Validate checks the field values on GetOrderInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetOrderInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrderInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrderInfoRequestMultiError, or nil if none found.
func (m *GetOrderInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrderInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetOrderId() <= 0 {
		err := GetOrderInfoRequestValidationError{
			field:  "OrderId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetOrderInfoRequestMultiError(errors)
	}

	return nil
}

// GetOrderInfoRequestMultiError is an error wrapping multiple validation
// errors returned by GetOrderInfoRequest.ValidateAll() if the designated
// constraints aren't met.
type GetOrderInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrderInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrderInfoRequestMultiError) AllErrors() []error { return m }

// GetOrderInfoRequestValidationError is the validation error returned by
// GetOrderInfoRequest.Validate if the designated constraints aren't met.
type GetOrderInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderInfoRequestValidationError) ErrorName() string {
	return "GetOrderInfoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetOrderInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderInfoRequestValidationError{}

// Validate checks the field values on GetOrderInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetOrderInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrderInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrderInfoResponseMultiError, or nil if none found.
func (m *GetOrderInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrderInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for User

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetOrderInfoResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetOrderInfoResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetOrderInfoResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetOrderInfoResponseMultiError(errors)
	}

	return nil
}

// GetOrderInfoResponseMultiError is an error wrapping multiple validation
// errors returned by GetOrderInfoResponse.ValidateAll() if the designated
// constraints aren't met.
type GetOrderInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrderInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrderInfoResponseMultiError) AllErrors() []error { return m }

// GetOrderInfoResponseValidationError is the validation error returned by
// GetOrderInfoResponse.Validate if the designated constraints aren't met.
type GetOrderInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderInfoResponseValidationError) ErrorName() string {
	return "GetOrderInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetOrderInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderInfoResponseValidationError{}

// Validate checks the field values on OrderPayRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OrderPayRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderPayRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderPayRequestMultiError, or nil if none found.
func (m *OrderPayRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderPayRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetOrderId() <= 0 {
		err := OrderPayRequestValidationError{
			field:  "OrderId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OrderPayRequestMultiError(errors)
	}

	return nil
}

// OrderPayRequestMultiError is an error wrapping multiple validation errors
// returned by OrderPayRequest.ValidateAll() if the designated constraints
// aren't met.
type OrderPayRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderPayRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderPayRequestMultiError) AllErrors() []error { return m }

// OrderPayRequestValidationError is the validation error returned by
// OrderPayRequest.Validate if the designated constraints aren't met.
type OrderPayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderPayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderPayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderPayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderPayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderPayRequestValidationError) ErrorName() string { return "OrderPayRequestValidationError" }

// Error satisfies the builtin error interface
func (e OrderPayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderPayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderPayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderPayRequestValidationError{}
