// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package types

// Example defines model for example.
type Example []MyItem

// MyItem defines model for my_item.
type MyItem struct {
	Age  *int    `json:"age,omitempty"`
	Name *string `json:"name,omitempty"`
}