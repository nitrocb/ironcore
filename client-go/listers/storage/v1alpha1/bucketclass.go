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

// BucketClassLister helps list BucketClasses.
// All objects returned here must be treated as read-only.
type BucketClassLister interface {
	// List lists all BucketClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BucketClass, err error)
	// Get retrieves the BucketClass from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BucketClass, error)
	BucketClassListerExpansion
}

// bucketClassLister implements the BucketClassLister interface.
type bucketClassLister struct {
	indexer cache.Indexer
}

// NewBucketClassLister returns a new BucketClassLister.
func NewBucketClassLister(indexer cache.Indexer) BucketClassLister {
	return &bucketClassLister{indexer: indexer}
}

// List lists all BucketClasses in the indexer.
func (s *bucketClassLister) List(selector labels.Selector) (ret []*v1alpha1.BucketClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketClass))
	})
	return ret, err
}

// Get retrieves the BucketClass from the index for a given name.
func (s *bucketClassLister) Get(name string) (*v1alpha1.BucketClass, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bucketclass"), name)
	}
	return obj.(*v1alpha1.BucketClass), nil
}
