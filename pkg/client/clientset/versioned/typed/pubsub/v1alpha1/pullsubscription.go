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
	v1alpha1 "github.com/googlecloudplatform/cloud-run-events/pkg/apis/pubsub/v1alpha1"
	scheme "github.com/googlecloudplatform/cloud-run-events/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PullSubscriptionsGetter has a method to return a PullSubscriptionInterface.
// A group's client should implement this interface.
type PullSubscriptionsGetter interface {
	PullSubscriptions(namespace string) PullSubscriptionInterface
}

// PullSubscriptionInterface has methods to work with PullSubscription resources.
type PullSubscriptionInterface interface {
	Create(*v1alpha1.PullSubscription) (*v1alpha1.PullSubscription, error)
	Update(*v1alpha1.PullSubscription) (*v1alpha1.PullSubscription, error)
	UpdateStatus(*v1alpha1.PullSubscription) (*v1alpha1.PullSubscription, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PullSubscription, error)
	List(opts v1.ListOptions) (*v1alpha1.PullSubscriptionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PullSubscription, err error)
	PullSubscriptionExpansion
}

// pullSubscriptions implements PullSubscriptionInterface
type pullSubscriptions struct {
	client rest.Interface
	ns     string
}

// newPullSubscriptions returns a PullSubscriptions
func newPullSubscriptions(c *PubsubV1alpha1Client, namespace string) *pullSubscriptions {
	return &pullSubscriptions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pullSubscription, and returns the corresponding pullSubscription object, and an error if there is any.
func (c *pullSubscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.PullSubscription, err error) {
	result = &v1alpha1.PullSubscription{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pullsubscriptions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PullSubscriptions that match those selectors.
func (c *pullSubscriptions) List(opts v1.ListOptions) (result *v1alpha1.PullSubscriptionList, err error) {
	result = &v1alpha1.PullSubscriptionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pullsubscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pullSubscriptions.
func (c *pullSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pullsubscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a pullSubscription and creates it.  Returns the server's representation of the pullSubscription, and an error, if there is any.
func (c *pullSubscriptions) Create(pullSubscription *v1alpha1.PullSubscription) (result *v1alpha1.PullSubscription, err error) {
	result = &v1alpha1.PullSubscription{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pullsubscriptions").
		Body(pullSubscription).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pullSubscription and updates it. Returns the server's representation of the pullSubscription, and an error, if there is any.
func (c *pullSubscriptions) Update(pullSubscription *v1alpha1.PullSubscription) (result *v1alpha1.PullSubscription, err error) {
	result = &v1alpha1.PullSubscription{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pullsubscriptions").
		Name(pullSubscription.Name).
		Body(pullSubscription).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pullSubscriptions) UpdateStatus(pullSubscription *v1alpha1.PullSubscription) (result *v1alpha1.PullSubscription, err error) {
	result = &v1alpha1.PullSubscription{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pullsubscriptions").
		Name(pullSubscription.Name).
		SubResource("status").
		Body(pullSubscription).
		Do().
		Into(result)
	return
}

// Delete takes name of the pullSubscription and deletes it. Returns an error if one occurs.
func (c *pullSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pullsubscriptions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pullSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pullsubscriptions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pullSubscription.
func (c *pullSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PullSubscription, err error) {
	result = &v1alpha1.PullSubscription{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pullsubscriptions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
