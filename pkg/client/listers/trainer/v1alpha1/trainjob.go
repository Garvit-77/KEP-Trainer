// Copyright 2024 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	trainerv1alpha1 "github.com/kubeflow/trainer/pkg/apis/trainer/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// TrainJobLister helps list TrainJobs.
// All objects returned here must be treated as read-only.
type TrainJobLister interface {
	// List lists all TrainJobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*trainerv1alpha1.TrainJob, err error)
	// TrainJobs returns an object that can list and get TrainJobs.
	TrainJobs(namespace string) TrainJobNamespaceLister
	TrainJobListerExpansion
}

// trainJobLister implements the TrainJobLister interface.
type trainJobLister struct {
	listers.ResourceIndexer[*trainerv1alpha1.TrainJob]
}

// NewTrainJobLister returns a new TrainJobLister.
func NewTrainJobLister(indexer cache.Indexer) TrainJobLister {
	return &trainJobLister{listers.New[*trainerv1alpha1.TrainJob](indexer, trainerv1alpha1.Resource("trainjob"))}
}

// TrainJobs returns an object that can list and get TrainJobs.
func (s *trainJobLister) TrainJobs(namespace string) TrainJobNamespaceLister {
	return trainJobNamespaceLister{listers.NewNamespaced[*trainerv1alpha1.TrainJob](s.ResourceIndexer, namespace)}
}

// TrainJobNamespaceLister helps list and get TrainJobs.
// All objects returned here must be treated as read-only.
type TrainJobNamespaceLister interface {
	// List lists all TrainJobs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*trainerv1alpha1.TrainJob, err error)
	// Get retrieves the TrainJob from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*trainerv1alpha1.TrainJob, error)
	TrainJobNamespaceListerExpansion
}

// trainJobNamespaceLister implements the TrainJobNamespaceLister
// interface.
type trainJobNamespaceLister struct {
	listers.ResourceIndexer[*trainerv1alpha1.TrainJob]
}
