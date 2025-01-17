// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package e2e contains tests that validate the behavior of gmp-operator against a cluster.
// To make tests simple and fast, the test suite runs the operator internally. The CRDs
// are expected to be installed out of band (along with the operator deployment itself in
// a real world setup).
package kubeutil

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func isDeploymentReady(deployment *appsv1.Deployment) error {
	// Set to default replicas value.
	expected := int32(1)
	if deployment.Spec.Replicas != nil {
		expected = *deployment.Spec.Replicas
	}
	if deployment.Status.ReadyReplicas != expected {
		return errors.New("replicas unavailable")
	}
	return nil
}

func WaitForDeploymentReady(ctx context.Context, kubeClient client.Client, namespace, name string) error {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	return waitForResourceReady(ctx, kubeClient, deployment, 2*time.Minute, isDeploymentReady)
}

func DeploymentContainer(deployment *appsv1.Deployment, name string) (*corev1.Container, error) {
	for i := range deployment.Spec.Template.Spec.Containers {
		container := &deployment.Spec.Template.Spec.Containers[i]
		if container.Name == name {
			return container, nil
		}
	}
	return nil, fmt.Errorf("unable to find container %q", name)
}

func DeploymentPods(ctx context.Context, kubeClient client.Client, deployment *appsv1.Deployment) ([]corev1.Pod, error) {
	return selectorPods(ctx, kubeClient, deployment.Spec.Selector)
}

func DeploymentDebug(t *testing.T, ctx context.Context, restConfig *rest.Config, kubeClient client.Client, namespace, name string) {
	deployment := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
	}
	containerDebug(t, ctx, restConfig, kubeClient, schema.GroupVersionKind{
		Version: "v1",
		Kind:    "Deployment",
	}, &deployment, "deployment")
}
