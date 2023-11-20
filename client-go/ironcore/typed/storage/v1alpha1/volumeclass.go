// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/ironcore-dev/ironcore/api/storage/v1alpha1"
	storagev1alpha1 "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/storage/v1alpha1"
	scheme "github.com/ironcore-dev/ironcore/client-go/ironcore/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumeClassesGetter has a method to return a VolumeClassInterface.
// A group's client should implement this interface.
type VolumeClassesGetter interface {
	VolumeClasses() VolumeClassInterface
}

// VolumeClassInterface has methods to work with VolumeClass resources.
type VolumeClassInterface interface {
	Create(ctx context.Context, volumeClass *v1alpha1.VolumeClass, opts v1.CreateOptions) (*v1alpha1.VolumeClass, error)
	Update(ctx context.Context, volumeClass *v1alpha1.VolumeClass, opts v1.UpdateOptions) (*v1alpha1.VolumeClass, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VolumeClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VolumeClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeClass, err error)
	Apply(ctx context.Context, volumeClass *storagev1alpha1.VolumeClassApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VolumeClass, err error)
	VolumeClassExpansion
}

// volumeClasses implements VolumeClassInterface
type volumeClasses struct {
	client rest.Interface
}

// newVolumeClasses returns a VolumeClasses
func newVolumeClasses(c *StorageV1alpha1Client) *volumeClasses {
	return &volumeClasses{
		client: c.RESTClient(),
	}
}

// Get takes name of the volumeClass, and returns the corresponding volumeClass object, and an error if there is any.
func (c *volumeClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VolumeClass, err error) {
	result = &v1alpha1.VolumeClass{}
	err = c.client.Get().
		Resource("volumeclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeClasses that match those selectors.
func (c *volumeClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VolumeClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VolumeClassList{}
	err = c.client.Get().
		Resource("volumeclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeClasses.
func (c *volumeClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("volumeclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a volumeClass and creates it.  Returns the server's representation of the volumeClass, and an error, if there is any.
func (c *volumeClasses) Create(ctx context.Context, volumeClass *v1alpha1.VolumeClass, opts v1.CreateOptions) (result *v1alpha1.VolumeClass, err error) {
	result = &v1alpha1.VolumeClass{}
	err = c.client.Post().
		Resource("volumeclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a volumeClass and updates it. Returns the server's representation of the volumeClass, and an error, if there is any.
func (c *volumeClasses) Update(ctx context.Context, volumeClass *v1alpha1.VolumeClass, opts v1.UpdateOptions) (result *v1alpha1.VolumeClass, err error) {
	result = &v1alpha1.VolumeClass{}
	err = c.client.Put().
		Resource("volumeclasses").
		Name(volumeClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the volumeClass and deletes it. Returns an error if one occurs.
func (c *volumeClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("volumeclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("volumeclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched volumeClass.
func (c *volumeClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeClass, err error) {
	result = &v1alpha1.VolumeClass{}
	err = c.client.Patch(pt).
		Resource("volumeclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied volumeClass.
func (c *volumeClasses) Apply(ctx context.Context, volumeClass *storagev1alpha1.VolumeClassApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VolumeClass, err error) {
	if volumeClass == nil {
		return nil, fmt.Errorf("volumeClass provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(volumeClass)
	if err != nil {
		return nil, err
	}
	name := volumeClass.Name
	if name == nil {
		return nil, fmt.Errorf("volumeClass.Name must be provided to Apply")
	}
	result = &v1alpha1.VolumeClass{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("volumeclasses").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
