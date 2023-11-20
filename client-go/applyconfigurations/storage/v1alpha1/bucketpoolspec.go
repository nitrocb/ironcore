// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/common/v1alpha1"
)

// BucketPoolSpecApplyConfiguration represents an declarative configuration of the BucketPoolSpec type for use
// with apply.
type BucketPoolSpecApplyConfiguration struct {
	ProviderID *string                            `json:"providerID,omitempty"`
	Taints     []v1alpha1.TaintApplyConfiguration `json:"taints,omitempty"`
}

// BucketPoolSpecApplyConfiguration constructs an declarative configuration of the BucketPoolSpec type for use with
// apply.
func BucketPoolSpec() *BucketPoolSpecApplyConfiguration {
	return &BucketPoolSpecApplyConfiguration{}
}

// WithProviderID sets the ProviderID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProviderID field is set to the value of the last call.
func (b *BucketPoolSpecApplyConfiguration) WithProviderID(value string) *BucketPoolSpecApplyConfiguration {
	b.ProviderID = &value
	return b
}

// WithTaints adds the given value to the Taints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Taints field.
func (b *BucketPoolSpecApplyConfiguration) WithTaints(values ...*v1alpha1.TaintApplyConfiguration) *BucketPoolSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTaints")
		}
		b.Taints = append(b.Taints, *values[i])
	}
	return b
}
