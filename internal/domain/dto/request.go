package dto

import "strings"

type Key string

const RequestContextKey Key = "requestContext"

type Request struct {
	Method  string
	URL     string
	Headers []*Header
	Body    any
	Entity  any
}

type Header struct {
	Name  string
	Value interface{}
}

type RequestContext struct {
	Tid   string
	Roles []string
}

func (rc *RequestContext) HasRole(role string) bool {
	for _, r := range rc.Roles {
		if strings.EqualFold(role, r) {
			return true
		}
	}
	return false
}

func (rc *RequestContext) HasAnyRole(roles ...string) bool {
	for _, r := range roles {
		if rc.HasRole(r) {
			return true
		}
	}
	return false
}
