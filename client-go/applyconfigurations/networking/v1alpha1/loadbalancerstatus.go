// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ironcore/api/common/v1alpha1"
)

// LoadBalancerStatusApplyConfiguration represents an declarative configuration of the LoadBalancerStatus type for use
// with apply.
type LoadBalancerStatusApplyConfiguration struct {
	IPs []v1alpha1.IP `json:"ips,omitempty"`
}

// LoadBalancerStatusApplyConfiguration constructs an declarative configuration of the LoadBalancerStatus type for use with
// apply.
func LoadBalancerStatus() *LoadBalancerStatusApplyConfiguration {
	return &LoadBalancerStatusApplyConfiguration{}
}

// WithIPs adds the given value to the IPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IPs field.
func (b *LoadBalancerStatusApplyConfiguration) WithIPs(values ...v1alpha1.IP) *LoadBalancerStatusApplyConfiguration {
	for i := range values {
		b.IPs = append(b.IPs, values[i])
	}
	return b
}
