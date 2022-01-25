//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	*out = *in
	in.ImpersonateConfig.DeepCopyInto(&out.ImpersonateConfig)
	out.IAMKeySecret = in.IAMKeySecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMKeySecret) DeepCopyInto(out *IAMKeySecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMKeySecret.
func (in *IAMKeySecret) DeepCopy() *IAMKeySecret {
	if in == nil {
		return nil
	}
	out := new(IAMKeySecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImpersonateConfig) DeepCopyInto(out *ImpersonateConfig) {
	*out = *in
	if in.Delegates != nil {
		in, out := &in.Delegates, &out.Delegates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImpersonateConfig.
func (in *ImpersonateConfig) DeepCopy() *ImpersonateConfig {
	if in == nil {
		return nil
	}
	out := new(ImpersonateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleConfig) DeepCopyInto(out *ScaleConfig) {
	*out = *in
	out.Nodes = in.Nodes
	out.ProcessingUnits = in.ProcessingUnits
	out.TargetCPUUtilization = in.TargetCPUUtilization
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleConfig.
func (in *ScaleConfig) DeepCopy() *ScaleConfig {
	if in == nil {
		return nil
	}
	out := new(ScaleConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleConfigNodes) DeepCopyInto(out *ScaleConfigNodes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleConfigNodes.
func (in *ScaleConfigNodes) DeepCopy() *ScaleConfigNodes {
	if in == nil {
		return nil
	}
	out := new(ScaleConfigNodes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleConfigPUs) DeepCopyInto(out *ScaleConfigPUs) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleConfigPUs.
func (in *ScaleConfigPUs) DeepCopy() *ScaleConfigPUs {
	if in == nil {
		return nil
	}
	out := new(ScaleConfigPUs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerAutoscaler) DeepCopyInto(out *SpannerAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerAutoscaler.
func (in *SpannerAutoscaler) DeepCopy() *SpannerAutoscaler {
	if in == nil {
		return nil
	}
	out := new(SpannerAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpannerAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerAutoscalerList) DeepCopyInto(out *SpannerAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpannerAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerAutoscalerList.
func (in *SpannerAutoscalerList) DeepCopy() *SpannerAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(SpannerAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpannerAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerAutoscalerSpec) DeepCopyInto(out *SpannerAutoscalerSpec) {
	*out = *in
	out.TargetInstance = in.TargetInstance
	in.Authentication.DeepCopyInto(&out.Authentication)
	out.ScaleConfig = in.ScaleConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerAutoscalerSpec.
func (in *SpannerAutoscalerSpec) DeepCopy() *SpannerAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(SpannerAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpannerAutoscalerStatus) DeepCopyInto(out *SpannerAutoscalerStatus) {
	*out = *in
	if in.Schedules != nil {
		in, out := &in.Schedules, &out.Schedules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CurrentlyActiveSchedules != nil {
		in, out := &in.CurrentlyActiveSchedules, &out.CurrentlyActiveSchedules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.LastScaleTime.DeepCopyInto(&out.LastScaleTime)
	in.LastSyncTime.DeepCopyInto(&out.LastSyncTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpannerAutoscalerStatus.
func (in *SpannerAutoscalerStatus) DeepCopy() *SpannerAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(SpannerAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetCPUUtilization) DeepCopyInto(out *TargetCPUUtilization) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetCPUUtilization.
func (in *TargetCPUUtilization) DeepCopy() *TargetCPUUtilization {
	if in == nil {
		return nil
	}
	out := new(TargetCPUUtilization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetInstance) DeepCopyInto(out *TargetInstance) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetInstance.
func (in *TargetInstance) DeepCopy() *TargetInstance {
	if in == nil {
		return nil
	}
	out := new(TargetInstance)
	in.DeepCopyInto(out)
	return out
}