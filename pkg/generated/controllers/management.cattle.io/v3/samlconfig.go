/*
Copyright 2020 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type SamlConfigHandler func(string, *v3.SamlConfig) (*v3.SamlConfig, error)

type SamlConfigController interface {
	generic.ControllerMeta
	SamlConfigClient

	OnChange(ctx context.Context, name string, sync SamlConfigHandler)
	OnRemove(ctx context.Context, name string, sync SamlConfigHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() SamlConfigCache
}

type SamlConfigClient interface {
	Create(*v3.SamlConfig) (*v3.SamlConfig, error)
	Update(*v3.SamlConfig) (*v3.SamlConfig, error)

	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v3.SamlConfig, error)
	List(opts metav1.ListOptions) (*v3.SamlConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.SamlConfig, err error)
}

type SamlConfigCache interface {
	Get(name string) (*v3.SamlConfig, error)
	List(selector labels.Selector) ([]*v3.SamlConfig, error)

	AddIndexer(indexName string, indexer SamlConfigIndexer)
	GetByIndex(indexName, key string) ([]*v3.SamlConfig, error)
}

type SamlConfigIndexer func(obj *v3.SamlConfig) ([]string, error)

type samlConfigController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewSamlConfigController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) SamlConfigController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &samlConfigController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromSamlConfigHandlerToHandler(sync SamlConfigHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.SamlConfig
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.SamlConfig))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *samlConfigController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.SamlConfig))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateSamlConfigDeepCopyOnChange(client SamlConfigClient, obj *v3.SamlConfig, handler func(obj *v3.SamlConfig) (*v3.SamlConfig, error)) (*v3.SamlConfig, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *samlConfigController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *samlConfigController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *samlConfigController) OnChange(ctx context.Context, name string, sync SamlConfigHandler) {
	c.AddGenericHandler(ctx, name, FromSamlConfigHandlerToHandler(sync))
}

func (c *samlConfigController) OnRemove(ctx context.Context, name string, sync SamlConfigHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromSamlConfigHandlerToHandler(sync)))
}

func (c *samlConfigController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *samlConfigController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *samlConfigController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *samlConfigController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *samlConfigController) Cache() SamlConfigCache {
	return &samlConfigCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *samlConfigController) Create(obj *v3.SamlConfig) (*v3.SamlConfig, error) {
	result := &v3.SamlConfig{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *samlConfigController) Update(obj *v3.SamlConfig) (*v3.SamlConfig, error) {
	result := &v3.SamlConfig{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *samlConfigController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *samlConfigController) Get(name string, options metav1.GetOptions) (*v3.SamlConfig, error) {
	result := &v3.SamlConfig{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *samlConfigController) List(opts metav1.ListOptions) (*v3.SamlConfigList, error) {
	result := &v3.SamlConfigList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *samlConfigController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *samlConfigController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v3.SamlConfig, error) {
	result := &v3.SamlConfig{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type samlConfigCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *samlConfigCache) Get(name string) (*v3.SamlConfig, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.SamlConfig), nil
}

func (c *samlConfigCache) List(selector labels.Selector) (ret []*v3.SamlConfig, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.SamlConfig))
	})

	return ret, err
}

func (c *samlConfigCache) AddIndexer(indexName string, indexer SamlConfigIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.SamlConfig))
		},
	}))
}

func (c *samlConfigCache) GetByIndex(indexName, key string) (result []*v3.SamlConfig, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.SamlConfig, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.SamlConfig))
	}
	return result, nil
}
