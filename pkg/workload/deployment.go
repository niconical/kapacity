/*
 Copyright 2023 The Kapacity Authors.

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

package workload

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Deployment represents behaviors of a Kubernetes Deployment.
type Deployment struct {
	client.Client
	Namespace string
	Selector  labels.Selector
}

func (w *Deployment) Sort(ctx context.Context, pods []*corev1.Pod) ([]*corev1.Pod, error) {
	rs := &ReplicaSet{
		Client:    w.Client,
		Namespace: w.Namespace,
		Selector:  w.Selector,
	}
	return rs.Sort(ctx, pods)
}

func (*Deployment) CanSelectPodsToScaleDown(context.Context) bool {
	return false
}

func (*Deployment) SelectPodsToScaleDown(context.Context, []*corev1.Pod) error {
	return fmt.Errorf("Deployment does not support this operation")
}
