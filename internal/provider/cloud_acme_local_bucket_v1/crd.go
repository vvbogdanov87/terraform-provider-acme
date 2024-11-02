package cloud_acme_local_bucket_v1

import (
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type K8sCR struct {
	metav1.TypeMeta `tfsdk:"-" json:",inline"`
	Metadata        metav1.ObjectMeta `tfsdk:"-" json:"metadata,omitempty"`

	Name            types.String   `tfsdk:"name" json:"-"`
	Timeouts        timeouts.Value `tfsdk:"timeouts" json:"-"`
	ResourceVersion types.String   `tfsdk:"resource_version" json:"-"`

	Spec   *K8sSpec   `tfsdk:"spec" json:"spec,omitempty"`
	Status *K8sStatus `tfsdk:"status" json:"status"`
}

type K8sSpec struct {
	Region string `tfsdk:"region" json:"region"`
}

type K8sStatus struct {
	Arn *string `tfsdk:"arn" json:"arn"`

	Conditions *[]struct {
		Type   *string `tfsdk:"-" json:"type"`
		Status *string `tfsdk:"-" json:"status"`
	} `tfsdk:"-" json:"conditions"`
}
