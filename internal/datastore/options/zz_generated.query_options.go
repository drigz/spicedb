// Code generated by github.com/ecordell/optgen. DO NOT EDIT.
package options

import v0 "github.com/authzed/authzed-go/proto/authzed/api/v0"

type QueryOptionsOption func(q *QueryOptions)

// NewQueryOptionsWithOptions creates a new QueryOptions with the passed in options set
func NewQueryOptionsWithOptions(opts ...QueryOptionsOption) *QueryOptions {
	q := &QueryOptions{}
	for _, o := range opts {
		o(q)
	}
	return q
}

// ToOption returns a new QueryOptionsOption that sets the values from the passed in QueryOptions
func (q *QueryOptions) ToOption() QueryOptionsOption {
	return func(to *QueryOptions) {
		to.Limit = q.Limit
		to.Usersets = q.Usersets
	}
}

// QueryOptionsWithOptions configures an existing QueryOptions with the passed in options set
func QueryOptionsWithOptions(q *QueryOptions, opts ...QueryOptionsOption) *QueryOptions {
	for _, o := range opts {
		o(q)
	}
	return q
}

// WithLimit returns an option that can set Limit on a QueryOptions
func WithLimit(limit *uint64) QueryOptionsOption {
	return func(q *QueryOptions) {
		q.Limit = limit
	}
}

// WithUsersets returns an option that can append Usersetss to QueryOptions.Usersets
func WithUsersets(usersets *v0.ObjectAndRelation) QueryOptionsOption {
	return func(q *QueryOptions) {
		q.Usersets = append(q.Usersets, usersets)
	}
}

// SetUsersets returns an option that can set Usersets on a QueryOptions
func SetUsersets(usersets []*v0.ObjectAndRelation) QueryOptionsOption {
	return func(q *QueryOptions) {
		q.Usersets = usersets
	}
}

type ReverseQueryOptionsOption func(r *ReverseQueryOptions)

// NewReverseQueryOptionsWithOptions creates a new ReverseQueryOptions with the passed in options set
func NewReverseQueryOptionsWithOptions(opts ...ReverseQueryOptionsOption) *ReverseQueryOptions {
	r := &ReverseQueryOptions{}
	for _, o := range opts {
		o(r)
	}
	return r
}

// ToOption returns a new ReverseQueryOptionsOption that sets the values from the passed in ReverseQueryOptions
func (r *ReverseQueryOptions) ToOption() ReverseQueryOptionsOption {
	return func(to *ReverseQueryOptions) {
		to.ReverseLimit = r.ReverseLimit
		to.ResRelation = r.ResRelation
	}
}

// ReverseQueryOptionsWithOptions configures an existing ReverseQueryOptions with the passed in options set
func ReverseQueryOptionsWithOptions(r *ReverseQueryOptions, opts ...ReverseQueryOptionsOption) *ReverseQueryOptions {
	for _, o := range opts {
		o(r)
	}
	return r
}

// WithReverseLimit returns an option that can set ReverseLimit on a ReverseQueryOptions
func WithReverseLimit(reverseLimit *uint64) ReverseQueryOptionsOption {
	return func(r *ReverseQueryOptions) {
		r.ReverseLimit = reverseLimit
	}
}

// WithResRelation returns an option that can set ResRelation on a ReverseQueryOptions
func WithResRelation(resRelation *ResourceRelation) ReverseQueryOptionsOption {
	return func(r *ReverseQueryOptions) {
		r.ResRelation = resRelation
	}
}
