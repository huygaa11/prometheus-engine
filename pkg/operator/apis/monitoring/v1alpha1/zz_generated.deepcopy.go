//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2021 Google LLC
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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertingSpec) DeepCopyInto(out *AlertingSpec) {
	*out = *in
	if in.Alertmanagers != nil {
		in, out := &in.Alertmanagers, &out.Alertmanagers
		*out = make([]AlertmanagerEndpoints, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertingSpec.
func (in *AlertingSpec) DeepCopy() *AlertingSpec {
	if in == nil {
		return nil
	}
	out := new(AlertingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerEndpoints) DeepCopyInto(out *AlertmanagerEndpoints) {
	*out = *in
	out.Port = in.Port
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(SafeAuthorization)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerEndpoints.
func (in *AlertmanagerEndpoints) DeepCopy() *AlertmanagerEndpoints {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerEndpoints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelMapping) DeepCopyInto(out *LabelMapping) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelMapping.
func (in *LabelMapping) DeepCopy() *LabelMapping {
	if in == nil {
		return nil
	}
	out := new(LabelMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringCondition) DeepCopyInto(out *MonitoringCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringCondition.
func (in *MonitoringCondition) DeepCopy() *MonitoringCondition {
	if in == nil {
		return nil
	}
	out := new(MonitoringCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedConfigMapKeySelector) DeepCopyInto(out *NamespacedConfigMapKeySelector) {
	*out = *in
	in.ConfigMapKeySelector.DeepCopyInto(&out.ConfigMapKeySelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedConfigMapKeySelector.
func (in *NamespacedConfigMapKeySelector) DeepCopy() *NamespacedConfigMapKeySelector {
	if in == nil {
		return nil
	}
	out := new(NamespacedConfigMapKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedSecretKeySelector) DeepCopyInto(out *NamespacedSecretKeySelector) {
	*out = *in
	in.SecretKeySelector.DeepCopyInto(&out.SecretKeySelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedSecretKeySelector.
func (in *NamespacedSecretKeySelector) DeepCopy() *NamespacedSecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(NamespacedSecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedSecretOrConfigMap) DeepCopyInto(out *NamespacedSecretOrConfigMap) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(NamespacedSecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(NamespacedConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedSecretOrConfigMap.
func (in *NamespacedSecretOrConfigMap) DeepCopy() *NamespacedSecretOrConfigMap {
	if in == nil {
		return nil
	}
	out := new(NamespacedSecretOrConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorConfig) DeepCopyInto(out *OperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Rules.DeepCopyInto(&out.Rules)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorConfig.
func (in *OperatorConfig) DeepCopy() *OperatorConfig {
	if in == nil {
		return nil
	}
	out := new(OperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorConfigList) DeepCopyInto(out *OperatorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OperatorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorConfigList.
func (in *OperatorConfigList) DeepCopy() *OperatorConfigList {
	if in == nil {
		return nil
	}
	out := new(OperatorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodMonitoring) DeepCopyInto(out *PodMonitoring) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodMonitoring.
func (in *PodMonitoring) DeepCopy() *PodMonitoring {
	if in == nil {
		return nil
	}
	out := new(PodMonitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodMonitoring) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodMonitoringList) DeepCopyInto(out *PodMonitoringList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodMonitoring, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodMonitoringList.
func (in *PodMonitoringList) DeepCopy() *PodMonitoringList {
	if in == nil {
		return nil
	}
	out := new(PodMonitoringList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodMonitoringList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodMonitoringSpec) DeepCopyInto(out *PodMonitoringSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]ScrapeEndpoint, len(*in))
		copy(*out, *in)
	}
	in.TargetLabels.DeepCopyInto(&out.TargetLabels)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodMonitoringSpec.
func (in *PodMonitoringSpec) DeepCopy() *PodMonitoringSpec {
	if in == nil {
		return nil
	}
	out := new(PodMonitoringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodMonitoringStatus) DeepCopyInto(out *PodMonitoringStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MonitoringCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodMonitoringStatus.
func (in *PodMonitoringStatus) DeepCopy() *PodMonitoringStatus {
	if in == nil {
		return nil
	}
	out := new(PodMonitoringStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleEvaluatorSpec) DeepCopyInto(out *RuleEvaluatorSpec) {
	*out = *in
	in.Alerting.DeepCopyInto(&out.Alerting)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleEvaluatorSpec.
func (in *RuleEvaluatorSpec) DeepCopy() *RuleEvaluatorSpec {
	if in == nil {
		return nil
	}
	out := new(RuleEvaluatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleGroup) DeepCopyInto(out *RuleGroup) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleGroup.
func (in *RuleGroup) DeepCopy() *RuleGroup {
	if in == nil {
		return nil
	}
	out := new(RuleGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rules) DeepCopyInto(out *Rules) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rules.
func (in *Rules) DeepCopy() *Rules {
	if in == nil {
		return nil
	}
	out := new(Rules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Rules) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RulesList) DeepCopyInto(out *RulesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Rules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RulesList.
func (in *RulesList) DeepCopy() *RulesList {
	if in == nil {
		return nil
	}
	out := new(RulesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RulesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RulesSpec) DeepCopyInto(out *RulesSpec) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]RuleGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RulesSpec.
func (in *RulesSpec) DeepCopy() *RulesSpec {
	if in == nil {
		return nil
	}
	out := new(RulesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RulesStatus) DeepCopyInto(out *RulesStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RulesStatus.
func (in *RulesStatus) DeepCopy() *RulesStatus {
	if in == nil {
		return nil
	}
	out := new(RulesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SafeAuthorization) DeepCopyInto(out *SafeAuthorization) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(NamespacedSecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SafeAuthorization.
func (in *SafeAuthorization) DeepCopy() *SafeAuthorization {
	if in == nil {
		return nil
	}
	out := new(SafeAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScrapeEndpoint) DeepCopyInto(out *ScrapeEndpoint) {
	*out = *in
	out.Port = in.Port
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScrapeEndpoint.
func (in *ScrapeEndpoint) DeepCopy() *ScrapeEndpoint {
	if in == nil {
		return nil
	}
	out := new(ScrapeEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	in.CA.DeepCopyInto(&out.CA)
	in.Cert.DeepCopyInto(&out.Cert)
	if in.KeySecret != nil {
		in, out := &in.KeySecret, &out.KeySecret
		*out = new(NamespacedSecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetLabels) DeepCopyInto(out *TargetLabels) {
	*out = *in
	if in.FromPod != nil {
		in, out := &in.FromPod, &out.FromPod
		*out = make([]LabelMapping, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetLabels.
func (in *TargetLabels) DeepCopy() *TargetLabels {
	if in == nil {
		return nil
	}
	out := new(TargetLabels)
	in.DeepCopyInto(out)
	return out
}
