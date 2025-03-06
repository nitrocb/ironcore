//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	storagev1alpha1 "github.com/ironcore-dev/ironcore/api/storage/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&storagev1alpha1.Bucket{}, func(obj interface{}) { SetObjectDefaults_Bucket(obj.(*storagev1alpha1.Bucket)) })
	scheme.AddTypeDefaultingFunc(&storagev1alpha1.BucketList{}, func(obj interface{}) { SetObjectDefaults_BucketList(obj.(*storagev1alpha1.BucketList)) })
	scheme.AddTypeDefaultingFunc(&storagev1alpha1.Volume{}, func(obj interface{}) { SetObjectDefaults_Volume(obj.(*storagev1alpha1.Volume)) })
	scheme.AddTypeDefaultingFunc(&storagev1alpha1.VolumeClass{}, func(obj interface{}) { SetObjectDefaults_VolumeClass(obj.(*storagev1alpha1.VolumeClass)) })
	scheme.AddTypeDefaultingFunc(&storagev1alpha1.VolumeClassList{}, func(obj interface{}) { SetObjectDefaults_VolumeClassList(obj.(*storagev1alpha1.VolumeClassList)) })
	scheme.AddTypeDefaultingFunc(&storagev1alpha1.VolumeList{}, func(obj interface{}) { SetObjectDefaults_VolumeList(obj.(*storagev1alpha1.VolumeList)) })
	return nil
}

func SetObjectDefaults_Bucket(in *storagev1alpha1.Bucket) {
	SetDefaults_BucketStatus(&in.Status)
}

func SetObjectDefaults_BucketList(in *storagev1alpha1.BucketList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Bucket(a)
	}
}

func SetObjectDefaults_Volume(in *storagev1alpha1.Volume) {
	SetDefaults_VolumeStatus(&in.Status)
}

func SetObjectDefaults_VolumeClass(in *storagev1alpha1.VolumeClass) {
	SetDefaults_VolumeClass(in)
}

func SetObjectDefaults_VolumeClassList(in *storagev1alpha1.VolumeClassList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_VolumeClass(a)
	}
}

func SetObjectDefaults_VolumeList(in *storagev1alpha1.VolumeList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Volume(a)
	}
}
