/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	authv1 "tkestack.io/tke/api/auth/v1"
	versioned "tkestack.io/tke/api/client/clientset/versioned"
	internalinterfaces "tkestack.io/tke/api/client/informers/externalversions/internalinterfaces"
	v1 "tkestack.io/tke/api/client/listers/auth/v1"
)

// ProjectPolicyBindingInformer provides access to a shared informer and lister for
// ProjectPolicyBindings.
type ProjectPolicyBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ProjectPolicyBindingLister
}

type projectPolicyBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewProjectPolicyBindingInformer constructs a new informer for ProjectPolicyBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProjectPolicyBindingInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProjectPolicyBindingInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredProjectPolicyBindingInformer constructs a new informer for ProjectPolicyBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProjectPolicyBindingInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AuthV1().ProjectPolicyBindings().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AuthV1().ProjectPolicyBindings().Watch(options)
			},
		},
		&authv1.ProjectPolicyBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *projectPolicyBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProjectPolicyBindingInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *projectPolicyBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&authv1.ProjectPolicyBinding{}, f.defaultInformer)
}

func (f *projectPolicyBindingInformer) Lister() v1.ProjectPolicyBindingLister {
	return v1.NewProjectPolicyBindingLister(f.Informer().GetIndexer())
}
