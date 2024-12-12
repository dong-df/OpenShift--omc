//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha2

import (
	unsafe "unsafe"

	coordinationv1 "k8s.io/api/coordination/v1"
	coordinationv1alpha2 "k8s.io/api/coordination/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	coordination "k8s.io/kubernetes/pkg/apis/coordination"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*coordinationv1alpha2.LeaseCandidate)(nil), (*coordination.LeaseCandidate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_LeaseCandidate_To_coordination_LeaseCandidate(a.(*coordinationv1alpha2.LeaseCandidate), b.(*coordination.LeaseCandidate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coordination.LeaseCandidate)(nil), (*coordinationv1alpha2.LeaseCandidate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coordination_LeaseCandidate_To_v1alpha2_LeaseCandidate(a.(*coordination.LeaseCandidate), b.(*coordinationv1alpha2.LeaseCandidate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coordinationv1alpha2.LeaseCandidateList)(nil), (*coordination.LeaseCandidateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_LeaseCandidateList_To_coordination_LeaseCandidateList(a.(*coordinationv1alpha2.LeaseCandidateList), b.(*coordination.LeaseCandidateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coordination.LeaseCandidateList)(nil), (*coordinationv1alpha2.LeaseCandidateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coordination_LeaseCandidateList_To_v1alpha2_LeaseCandidateList(a.(*coordination.LeaseCandidateList), b.(*coordinationv1alpha2.LeaseCandidateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coordinationv1alpha2.LeaseCandidateSpec)(nil), (*coordination.LeaseCandidateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_LeaseCandidateSpec_To_coordination_LeaseCandidateSpec(a.(*coordinationv1alpha2.LeaseCandidateSpec), b.(*coordination.LeaseCandidateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coordination.LeaseCandidateSpec)(nil), (*coordinationv1alpha2.LeaseCandidateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coordination_LeaseCandidateSpec_To_v1alpha2_LeaseCandidateSpec(a.(*coordination.LeaseCandidateSpec), b.(*coordinationv1alpha2.LeaseCandidateSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_LeaseCandidate_To_coordination_LeaseCandidate(in *coordinationv1alpha2.LeaseCandidate, out *coordination.LeaseCandidate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha2_LeaseCandidateSpec_To_coordination_LeaseCandidateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_LeaseCandidate_To_coordination_LeaseCandidate is an autogenerated conversion function.
func Convert_v1alpha2_LeaseCandidate_To_coordination_LeaseCandidate(in *coordinationv1alpha2.LeaseCandidate, out *coordination.LeaseCandidate, s conversion.Scope) error {
	return autoConvert_v1alpha2_LeaseCandidate_To_coordination_LeaseCandidate(in, out, s)
}

func autoConvert_coordination_LeaseCandidate_To_v1alpha2_LeaseCandidate(in *coordination.LeaseCandidate, out *coordinationv1alpha2.LeaseCandidate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_coordination_LeaseCandidateSpec_To_v1alpha2_LeaseCandidateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_coordination_LeaseCandidate_To_v1alpha2_LeaseCandidate is an autogenerated conversion function.
func Convert_coordination_LeaseCandidate_To_v1alpha2_LeaseCandidate(in *coordination.LeaseCandidate, out *coordinationv1alpha2.LeaseCandidate, s conversion.Scope) error {
	return autoConvert_coordination_LeaseCandidate_To_v1alpha2_LeaseCandidate(in, out, s)
}

func autoConvert_v1alpha2_LeaseCandidateList_To_coordination_LeaseCandidateList(in *coordinationv1alpha2.LeaseCandidateList, out *coordination.LeaseCandidateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]coordination.LeaseCandidate)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha2_LeaseCandidateList_To_coordination_LeaseCandidateList is an autogenerated conversion function.
func Convert_v1alpha2_LeaseCandidateList_To_coordination_LeaseCandidateList(in *coordinationv1alpha2.LeaseCandidateList, out *coordination.LeaseCandidateList, s conversion.Scope) error {
	return autoConvert_v1alpha2_LeaseCandidateList_To_coordination_LeaseCandidateList(in, out, s)
}

func autoConvert_coordination_LeaseCandidateList_To_v1alpha2_LeaseCandidateList(in *coordination.LeaseCandidateList, out *coordinationv1alpha2.LeaseCandidateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]coordinationv1alpha2.LeaseCandidate)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_coordination_LeaseCandidateList_To_v1alpha2_LeaseCandidateList is an autogenerated conversion function.
func Convert_coordination_LeaseCandidateList_To_v1alpha2_LeaseCandidateList(in *coordination.LeaseCandidateList, out *coordinationv1alpha2.LeaseCandidateList, s conversion.Scope) error {
	return autoConvert_coordination_LeaseCandidateList_To_v1alpha2_LeaseCandidateList(in, out, s)
}

func autoConvert_v1alpha2_LeaseCandidateSpec_To_coordination_LeaseCandidateSpec(in *coordinationv1alpha2.LeaseCandidateSpec, out *coordination.LeaseCandidateSpec, s conversion.Scope) error {
	out.LeaseName = in.LeaseName
	out.PingTime = (*v1.MicroTime)(unsafe.Pointer(in.PingTime))
	out.RenewTime = (*v1.MicroTime)(unsafe.Pointer(in.RenewTime))
	out.BinaryVersion = in.BinaryVersion
	out.EmulationVersion = in.EmulationVersion
	out.Strategy = coordination.CoordinatedLeaseStrategy(in.Strategy)
	return nil
}

// Convert_v1alpha2_LeaseCandidateSpec_To_coordination_LeaseCandidateSpec is an autogenerated conversion function.
func Convert_v1alpha2_LeaseCandidateSpec_To_coordination_LeaseCandidateSpec(in *coordinationv1alpha2.LeaseCandidateSpec, out *coordination.LeaseCandidateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha2_LeaseCandidateSpec_To_coordination_LeaseCandidateSpec(in, out, s)
}

func autoConvert_coordination_LeaseCandidateSpec_To_v1alpha2_LeaseCandidateSpec(in *coordination.LeaseCandidateSpec, out *coordinationv1alpha2.LeaseCandidateSpec, s conversion.Scope) error {
	out.LeaseName = in.LeaseName
	out.PingTime = (*v1.MicroTime)(unsafe.Pointer(in.PingTime))
	out.RenewTime = (*v1.MicroTime)(unsafe.Pointer(in.RenewTime))
	out.BinaryVersion = in.BinaryVersion
	out.EmulationVersion = in.EmulationVersion
	out.Strategy = coordinationv1.CoordinatedLeaseStrategy(in.Strategy)
	return nil
}

// Convert_coordination_LeaseCandidateSpec_To_v1alpha2_LeaseCandidateSpec is an autogenerated conversion function.
func Convert_coordination_LeaseCandidateSpec_To_v1alpha2_LeaseCandidateSpec(in *coordination.LeaseCandidateSpec, out *coordinationv1alpha2.LeaseCandidateSpec, s conversion.Scope) error {
	return autoConvert_coordination_LeaseCandidateSpec_To_v1alpha2_LeaseCandidateSpec(in, out, s)
}
