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
	"net"
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func IsPodContainerReady(ctx context.Context, restConfig *rest.Config, pod *corev1.Pod, container string) error {
	for _, status := range pod.Status.ContainerStatuses {
		if status.Name == container {
			if !status.Ready {
				key := client.ObjectKeyFromObject(pod)
				return fmt.Errorf("pod %s container %s not ready: %s", key, status.Name, containerStateString(&status.State))
			}
			return nil
		}
	}
	key := client.ObjectKeyFromObject(pod)
	return fmt.Errorf("no container named %s found in pod %s", container, key)
}

func WaitForPodContainerReady(ctx context.Context, t testing.TB, restConfig *rest.Config, kubeClient client.Client, pod *corev1.Pod, container string) error {
	return waitForResourceReady(ctx, kubeClient, pod, 1*time.Minute, func(pod *corev1.Pod) error {
		return IsPodContainerReady(ctx, restConfig, pod, container)
	})
}

func IsPodReady(ctx context.Context, restConfig *rest.Config, pod *corev1.Pod) error {
	var errs []error
	for _, status := range pod.Status.ContainerStatuses {
		if !status.Ready {
			key := client.ObjectKeyFromObject(pod)
			errs = append(errs, fmt.Errorf("pod %s container %s not ready: %s", key, status.Name, containerStateString(&status.State)))
		}
	}
	return errors.Join(errs...)
}

func WaitForPodReady(ctx context.Context, t *testing.T, restConfig *rest.Config, kubeClient client.Client, pod *corev1.Pod) error {
	return waitForResourceReady(ctx, kubeClient, pod, 30*time.Second, func(pod *corev1.Pod) error {
		return IsPodReady(ctx, restConfig, pod)
	})
}

func PodByIP(ctx context.Context, kubeClient client.Client, ip net.IP) (*corev1.Pod, error) {
	var pods corev1.PodList
	if err := kubeClient.List(ctx, &pods, &client.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("status.podIP", ip.String()),
	}); err != nil {
		return nil, err
	}
	if len(pods.Items) != 1 {
		return nil, fmt.Errorf("expected 1 pod with IP %s, got %d", ip.String(), len(pods.Items))
	}
	return &pods.Items[0], nil
}

func PodByAddr(ctx context.Context, kubeClient client.Client, addr *net.TCPAddr) (*corev1.Pod, string, error) {
	pod, err := PodByIP(ctx, kubeClient, addr.IP)
	if err != nil {
		return nil, "", err
	}
	for _, container := range pod.Spec.Containers {
		for _, port := range container.Ports {
			if int(port.ContainerPort) == addr.Port {
				return pod, container.Name, nil
			}
		}
	}
	key := client.ObjectKeyFromObject(pod)
	return nil, "", fmt.Errorf("unable to find port %d in pod %s", addr.Port, key)
}

func selectorPods(ctx context.Context, kubeClient client.Client, selector *metav1.LabelSelector) ([]corev1.Pod, error) {
	var podList corev1.PodList
	requirements, err := labels.ParseToRequirements(metav1.FormatLabelSelector(selector))
	if err != nil {
		return nil, err
	}
	if err := kubeClient.List(ctx, &podList, &client.MatchingLabelsSelector{
		Selector: labels.NewSelector().Add(requirements...),
	}); err != nil {
		return nil, err
	}
	return podList.Items, nil
}
