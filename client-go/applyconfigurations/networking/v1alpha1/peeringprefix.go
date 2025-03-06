// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	commonv1alpha1 "github.com/ironcore-dev/ironcore/api/common/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

// PeeringPrefixApplyConfiguration represents a declarative configuration of the PeeringPrefix type for use
// with apply.
type PeeringPrefixApplyConfiguration struct {
	Name      *string                  `json:"name,omitempty"`
	Prefix    *commonv1alpha1.IPPrefix `json:"prefix,omitempty"`
	PrefixRef *v1.LocalObjectReference `json:"prefixRef,omitempty"`
}

// PeeringPrefixApplyConfiguration constructs a declarative configuration of the PeeringPrefix type for use with
// apply.
func PeeringPrefix() *PeeringPrefixApplyConfiguration {
	return &PeeringPrefixApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PeeringPrefixApplyConfiguration) WithName(value string) *PeeringPrefixApplyConfiguration {
	b.Name = &value
	return b
}

// WithPrefix sets the Prefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Prefix field is set to the value of the last call.
func (b *PeeringPrefixApplyConfiguration) WithPrefix(value commonv1alpha1.IPPrefix) *PeeringPrefixApplyConfiguration {
	b.Prefix = &value
	return b
}

// WithPrefixRef sets the PrefixRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PrefixRef field is set to the value of the last call.
func (b *PeeringPrefixApplyConfiguration) WithPrefixRef(value v1.LocalObjectReference) *PeeringPrefixApplyConfiguration {
	b.PrefixRef = &value
	return b
}
