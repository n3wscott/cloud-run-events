/*
Copyright 2019 Google LLC

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/google/knative-gcp/pkg/apis/messaging/v1alpha1"
	scheme "github.com/google/knative-gcp/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DecoratorsGetter has a method to return a DecoratorInterface.
// A group's client should implement this interface.
type DecoratorsGetter interface {
	Decorators(namespace string) DecoratorInterface
}

// DecoratorInterface has methods to work with Decorator resources.
type DecoratorInterface interface {
	Create(*v1alpha1.Decorator) (*v1alpha1.Decorator, error)
	Update(*v1alpha1.Decorator) (*v1alpha1.Decorator, error)
	UpdateStatus(*v1alpha1.Decorator) (*v1alpha1.Decorator, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Decorator, error)
	List(opts v1.ListOptions) (*v1alpha1.DecoratorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Decorator, err error)
	DecoratorExpansion
}

// decorators implements DecoratorInterface
type decorators struct {
	client rest.Interface
	ns     string
}

// newDecorators returns a Decorators
func newDecorators(c *MessagingV1alpha1Client, namespace string) *decorators {
	return &decorators{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the decorator, and returns the corresponding decorator object, and an error if there is any.
func (c *decorators) Get(name string, options v1.GetOptions) (result *v1alpha1.Decorator, err error) {
	result = &v1alpha1.Decorator{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("decorators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Decorators that match those selectors.
func (c *decorators) List(opts v1.ListOptions) (result *v1alpha1.DecoratorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DecoratorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("decorators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested decorators.
func (c *decorators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("decorators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a decorator and creates it.  Returns the server's representation of the decorator, and an error, if there is any.
func (c *decorators) Create(decorator *v1alpha1.Decorator) (result *v1alpha1.Decorator, err error) {
	result = &v1alpha1.Decorator{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("decorators").
		Body(decorator).
		Do().
		Into(result)
	return
}

// Update takes the representation of a decorator and updates it. Returns the server's representation of the decorator, and an error, if there is any.
func (c *decorators) Update(decorator *v1alpha1.Decorator) (result *v1alpha1.Decorator, err error) {
	result = &v1alpha1.Decorator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("decorators").
		Name(decorator.Name).
		Body(decorator).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *decorators) UpdateStatus(decorator *v1alpha1.Decorator) (result *v1alpha1.Decorator, err error) {
	result = &v1alpha1.Decorator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("decorators").
		Name(decorator.Name).
		SubResource("status").
		Body(decorator).
		Do().
		Into(result)
	return
}

// Delete takes name of the decorator and deletes it. Returns an error if one occurs.
func (c *decorators) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("decorators").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *decorators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("decorators").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched decorator.
func (c *decorators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Decorator, err error) {
	result = &v1alpha1.Decorator{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("decorators").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
