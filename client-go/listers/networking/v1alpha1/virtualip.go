// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ironcore/api/networking/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VirtualIPLister helps list VirtualIPs.
// All objects returned here must be treated as read-only.
type VirtualIPLister interface {
	// List lists all VirtualIPs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualIP, err error)
	// VirtualIPs returns an object that can list and get VirtualIPs.
	VirtualIPs(namespace string) VirtualIPNamespaceLister
	VirtualIPListerExpansion
}

// virtualIPLister implements the VirtualIPLister interface.
type virtualIPLister struct {
	indexer cache.Indexer
}

// NewVirtualIPLister returns a new VirtualIPLister.
func NewVirtualIPLister(indexer cache.Indexer) VirtualIPLister {
	return &virtualIPLister{indexer: indexer}
}

// List lists all VirtualIPs in the indexer.
func (s *virtualIPLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualIP, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualIP))
	})
	return ret, err
}

// VirtualIPs returns an object that can list and get VirtualIPs.
func (s *virtualIPLister) VirtualIPs(namespace string) VirtualIPNamespaceLister {
	return virtualIPNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualIPNamespaceLister helps list and get VirtualIPs.
// All objects returned here must be treated as read-only.
type VirtualIPNamespaceLister interface {
	// List lists all VirtualIPs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualIP, err error)
	// Get retrieves the VirtualIP from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VirtualIP, error)
	VirtualIPNamespaceListerExpansion
}

// virtualIPNamespaceLister implements the VirtualIPNamespaceLister
// interface.
type virtualIPNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualIPs in the indexer for a given namespace.
func (s virtualIPNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualIP, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualIP))
	})
	return ret, err
}

// Get retrieves the VirtualIP from the indexer for a given namespace and name.
func (s virtualIPNamespaceLister) Get(name string) (*v1alpha1.VirtualIP, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("virtualip"), name)
	}
	return obj.(*v1alpha1.VirtualIP), nil
}
