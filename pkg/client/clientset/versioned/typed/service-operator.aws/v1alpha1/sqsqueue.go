/*
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
	v1alpha1 "github.com/awslabs/aws-service-operator/pkg/apis/service-operator.aws/v1alpha1"
	scheme "github.com/awslabs/aws-service-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SQSQueuesGetter has a method to return a SQSQueueInterface.
// A group's client should implement this interface.
type SQSQueuesGetter interface {
	SQSQueues(namespace string) SQSQueueInterface
}

// SQSQueueInterface has methods to work with SQSQueue resources.
type SQSQueueInterface interface {
	Create(*v1alpha1.SQSQueue) (*v1alpha1.SQSQueue, error)
	Update(*v1alpha1.SQSQueue) (*v1alpha1.SQSQueue, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SQSQueue, error)
	List(opts v1.ListOptions) (*v1alpha1.SQSQueueList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SQSQueue, err error)
	SQSQueueExpansion
}

// sQSQueues implements SQSQueueInterface
type sQSQueues struct {
	client rest.Interface
	ns     string
}

// newSQSQueues returns a SQSQueues
func newSQSQueues(c *ServiceoperatorV1alpha1Client, namespace string) *sQSQueues {
	return &sQSQueues{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sQSQueue, and returns the corresponding sQSQueue object, and an error if there is any.
func (c *sQSQueues) Get(name string, options v1.GetOptions) (result *v1alpha1.SQSQueue, err error) {
	result = &v1alpha1.SQSQueue{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sqsqueues").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SQSQueues that match those selectors.
func (c *sQSQueues) List(opts v1.ListOptions) (result *v1alpha1.SQSQueueList, err error) {
	result = &v1alpha1.SQSQueueList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sqsqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sQSQueues.
func (c *sQSQueues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sqsqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a sQSQueue and creates it.  Returns the server's representation of the sQSQueue, and an error, if there is any.
func (c *sQSQueues) Create(sQSQueue *v1alpha1.SQSQueue) (result *v1alpha1.SQSQueue, err error) {
	result = &v1alpha1.SQSQueue{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sqsqueues").
		Body(sQSQueue).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sQSQueue and updates it. Returns the server's representation of the sQSQueue, and an error, if there is any.
func (c *sQSQueues) Update(sQSQueue *v1alpha1.SQSQueue) (result *v1alpha1.SQSQueue, err error) {
	result = &v1alpha1.SQSQueue{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sqsqueues").
		Name(sQSQueue.Name).
		Body(sQSQueue).
		Do().
		Into(result)
	return
}

// Delete takes name of the sQSQueue and deletes it. Returns an error if one occurs.
func (c *sQSQueues) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sqsqueues").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sQSQueues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sqsqueues").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sQSQueue.
func (c *sQSQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SQSQueue, err error) {
	result = &v1alpha1.SQSQueue{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sqsqueues").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
