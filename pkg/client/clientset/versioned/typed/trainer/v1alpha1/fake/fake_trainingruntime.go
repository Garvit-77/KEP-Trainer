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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kubeflow/trainer/pkg/apis/trainer/v1alpha1"
	trainerv1alpha1 "github.com/kubeflow/trainer/pkg/client/applyconfiguration/trainer/v1alpha1"
	typedtrainerv1alpha1 "github.com/kubeflow/trainer/pkg/client/clientset/versioned/typed/trainer/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeTrainingRuntimes implements TrainingRuntimeInterface
type fakeTrainingRuntimes struct {
	*gentype.FakeClientWithListAndApply[*v1alpha1.TrainingRuntime, *v1alpha1.TrainingRuntimeList, *trainerv1alpha1.TrainingRuntimeApplyConfiguration]
	Fake *FakeTrainerV1alpha1
}

func newFakeTrainingRuntimes(fake *FakeTrainerV1alpha1, namespace string) typedtrainerv1alpha1.TrainingRuntimeInterface {
	return &fakeTrainingRuntimes{
		gentype.NewFakeClientWithListAndApply[*v1alpha1.TrainingRuntime, *v1alpha1.TrainingRuntimeList, *trainerv1alpha1.TrainingRuntimeApplyConfiguration](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("trainingruntimes"),
			v1alpha1.SchemeGroupVersion.WithKind("TrainingRuntime"),
			func() *v1alpha1.TrainingRuntime { return &v1alpha1.TrainingRuntime{} },
			func() *v1alpha1.TrainingRuntimeList { return &v1alpha1.TrainingRuntimeList{} },
			func(dst, src *v1alpha1.TrainingRuntimeList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.TrainingRuntimeList) []*v1alpha1.TrainingRuntime {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.TrainingRuntimeList, items []*v1alpha1.TrainingRuntime) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
