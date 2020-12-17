// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deployable

import (
	"testing"

	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	//	. "github.com/onsi/gomega"
)

/*
In our call, structure went Pod > ReplicaSet > Deployment
How to implement a Deployment here?
*/
var (
	mcPod = &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "testPod",
			Namespace: mcName,
			OwnerReferences: []metav1.OwnerReference{
				{
					Kind: "Deployment",
					Name: "testDeployment",
				},
			},
		},
	}
)

func TestOwnerTraversal(t *testing.T) {

	// This test should test:
	// Passing a deployment (returns itself)
	// Passing a pod (returns associated deployment)

	g := NewWithT(t)

	metaNew, err := meta.Accessor(mcPod)
	if err != nil {
		klog.Error("Failed to access object metadata for sync with error: ", err)
		return
	}

	obj, err := findRootResource(metaNew)

	g.Expect(err).To(BeNil())
	g.Expect(obj).To(Equal(mcPod))
}
