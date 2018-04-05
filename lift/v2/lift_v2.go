// Package liftv2 intended for serve better
// exterience than first version
package liftv2

// Instance is liftv2 instance
type Instance instance

// Route is navigational point
type Route struct {
	Path   string
	Method string
}

// instance real type
type instance struct {
	routes map[string]string
}
