/*
Copyright 2022 The Knative Authors

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

package v1alpha1

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"knative.dev/operator/pkg/apis/operator/base"
	"knative.dev/operator/pkg/apis/operator/v1beta1"
	util "knative.dev/operator/pkg/reconciler/common/testing"
)

func TestKnativeServingConvertTo(t *testing.T) {
	tests := []struct {
		Name           string
		Input          *KnativeServing
		ExpectedOutput *v1beta1.KnativeServing
	}{{
		Name: "Knative Serving conversion",
		Input: &KnativeServing{
			Spec: KnativeServingSpec{
				CommonSpec: base.CommonSpec{
					Version: "1.2",
					DeprecatedResources: []base.ResourceRequirementsOverride{{
						Container: "webhook",
						ResourceRequirements: corev1.ResourceRequirements{
							Limits: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("999m"),
								corev1.ResourceMemory: resource.MustParse("999Mi")},
							Requests: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("999m"),
								corev1.ResourceMemory: resource.MustParse("999Mi")},
						},
					}},
				},
				Ingress: &IngressConfigs{
					Istio: base.IstioIngressConfiguration{
						Enabled: true,
					},
					Contour: base.ContourIngressConfiguration{
						Enabled: false,
					},
				},
				ControllerCustomCerts: base.CustomCerts{
					Type: "test-type",
					Name: "test-name",
				},
			},
		},
		ExpectedOutput: &v1beta1.KnativeServing{
			Spec: v1beta1.KnativeServingSpec{
				CommonSpec: base.CommonSpec{
					Version: "1.2",
					DeploymentOverride: []base.DeploymentOverride{
						{
							Name: "webhook",
							Resources: []base.ResourceRequirementsOverride{{
								Container: "webhook",
								ResourceRequirements: corev1.ResourceRequirements{
									Limits: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("999m"),
										corev1.ResourceMemory: resource.MustParse("999Mi")},
									Requests: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("999m"),
										corev1.ResourceMemory: resource.MustParse("999Mi")},
								},
							}},
						},
					},
				},
				Ingress: &v1beta1.IngressConfigs{
					Istio: base.IstioIngressConfiguration{
						Enabled: true,
					},
					Contour: base.ContourIngressConfiguration{
						Enabled: false,
					},
					Kourier: base.KourierIngressConfiguration{
						Enabled: false,
					},
				},
				ControllerCustomCerts: base.CustomCerts{
					Type: "test-type",
					Name: "test-name",
				},
			},
		},
	}, {
		Name: "Knative Serving conversion with no ingresses",
		Input: &KnativeServing{
			Spec: KnativeServingSpec{
				CommonSpec: base.CommonSpec{
					Version: "1.2",
					DeprecatedResources: []base.ResourceRequirementsOverride{{
						Container: "webhook",
						ResourceRequirements: corev1.ResourceRequirements{
							Limits: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("999m"),
								corev1.ResourceMemory: resource.MustParse("999Mi")},
							Requests: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("999m"),
								corev1.ResourceMemory: resource.MustParse("999Mi")},
						},
					}},
				},
				ControllerCustomCerts: base.CustomCerts{
					Type: "test-type",
					Name: "test-name",
				},
			},
		},
		ExpectedOutput: &v1beta1.KnativeServing{
			Spec: v1beta1.KnativeServingSpec{
				CommonSpec: base.CommonSpec{
					Version: "1.2",
					DeploymentOverride: []base.DeploymentOverride{
						{
							Name: "webhook",
							Resources: []base.ResourceRequirementsOverride{{
								Container: "webhook",
								ResourceRequirements: corev1.ResourceRequirements{
									Limits: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("999m"),
										corev1.ResourceMemory: resource.MustParse("999Mi")},
									Requests: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("999m"),
										corev1.ResourceMemory: resource.MustParse("999Mi")},
								},
							}},
						},
					},
				},
				ControllerCustomCerts: base.CustomCerts{
					Type: "test-type",
					Name: "test-name",
				},
			},
		},
	}}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			sink := &v1beta1.KnativeServing{}
			err := test.Input.ConvertTo(context.Background(), sink)
			util.AssertEqual(t, err, nil)
			util.AssertDeepEqual(t, sink, test.ExpectedOutput)
		})
	}
}

func TestTestKnativeServingConvertFrom(t *testing.T) {
	ke := &KnativeServing{}
	source := &v1beta1.KnativeServing{
		Spec: v1beta1.KnativeServingSpec{
			CommonSpec: base.CommonSpec{
				Version: "1.2",
			},
			Ingress: &v1beta1.IngressConfigs{
				Istio: base.IstioIngressConfiguration{
					Enabled: true,
				},
				Contour: base.ContourIngressConfiguration{
					Enabled: false,
				},
			},
			ControllerCustomCerts: base.CustomCerts{
				Type: "test-type",
				Name: "test-name",
			},
		},
	}
	err := ke.ConvertFrom(context.Background(), source)
	util.AssertEqual(t, err, nil)
	util.AssertEqual(t, ke.GetSpec().GetVersion(), "1.2")
	util.AssertEqual(t, ke.Spec.ControllerCustomCerts.Name, "test-name")
	util.AssertEqual(t, ke.Spec.ControllerCustomCerts.Type, "test-type")
	util.AssertEqual(t, ke.Spec.Ingress.Istio.Enabled, true)
	util.AssertEqual(t, ke.Spec.Ingress.Contour.Enabled, false)
}
