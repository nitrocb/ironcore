// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ironcore/api/storage/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BucketLister helps list Buckets.
// All objects returned here must be treated as read-only.
type BucketLister interface {
	// List lists all Buckets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Bucket, err error)
	// Buckets returns an object that can list and get Buckets.
	Buckets(namespace string) BucketNamespaceLister
	BucketListerExpansion
}

// bucketLister implements the BucketLister interface.
type bucketLister struct {
	indexer cache.Indexer
}

// NewBucketLister returns a new BucketLister.
func NewBucketLister(indexer cache.Indexer) BucketLister {
	return &bucketLister{indexer: indexer}
}

// List lists all Buckets in the indexer.
func (s *bucketLister) List(selector labels.Selector) (ret []*v1alpha1.Bucket, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Bucket))
	})
	return ret, err
}

// Buckets returns an object that can list and get Buckets.
func (s *bucketLister) Buckets(namespace string) BucketNamespaceLister {
	return bucketNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BucketNamespaceLister helps list and get Buckets.
// All objects returned here must be treated as read-only.
type BucketNamespaceLister interface {
	// List lists all Buckets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Bucket, err error)
	// Get retrieves the Bucket from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Bucket, error)
	BucketNamespaceListerExpansion
}

// bucketNamespaceLister implements the BucketNamespaceLister
// interface.
type bucketNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Buckets in the indexer for a given namespace.
func (s bucketNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Bucket, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Bucket))
	})
	return ret, err
}

// Get retrieves the Bucket from the indexer for a given namespace and name.
func (s bucketNamespaceLister) Get(name string) (*v1alpha1.Bucket, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bucket"), name)
	}
	return obj.(*v1alpha1.Bucket), nil
}
