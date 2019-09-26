// Code generated by kubetype-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "istio.io/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RbacConfig implements the ClusterRbaConfig Custom Resource Definition for controlling Istio RBAC behavior.
// The ClusterRbaConfig Custom Resource is a singleton where only one ClusterRbaConfig should be created
// globally in the mesh and the namespace should be the same to other Istio components, which usually is `istio-system`.
//
// Below is an example of an `ClusterRbacConfig` resource called `istio-rbac-config` which enables Istio RBAC for all
// services in the default namespace.
//
// ```yaml
// apiVersion: "rbac.istio.io/v1alpha1"
// kind: ClusterRbacConfig
// metadata:
//   name: default
//   namespace: istio-system
// spec:
//   mode: ON_WITH_INCLUSION
//   inclusion:
//     namespaces: [ "default" ]
// ```
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=rbac.istio.io/v1alpha1
// +kubetype-gen:kubeType=RbacConfig
// +kubetype-gen:kubeType=ClusterRbacConfig
// +kubetype-gen:ClusterRbacConfig:tag=genclient:nonNamespaced
// +genclient
// +k8s:deepcopy-gen=true
// -->
type RbacConfig struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1alpha1.RbacConfig `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RbacConfigList is a collection of RbacConfigs.
type RbacConfigList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []RbacConfig `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RbacConfig implements the ClusterRbaConfig Custom Resource Definition for controlling Istio RBAC behavior.
// The ClusterRbaConfig Custom Resource is a singleton where only one ClusterRbaConfig should be created
// globally in the mesh and the namespace should be the same to other Istio components, which usually is `istio-system`.
//
// Below is an example of an `ClusterRbacConfig` resource called `istio-rbac-config` which enables Istio RBAC for all
// services in the default namespace.
//
// ```yaml
// apiVersion: "rbac.istio.io/v1alpha1"
// kind: ClusterRbacConfig
// metadata:
//   name: default
//   namespace: istio-system
// spec:
//   mode: ON_WITH_INCLUSION
//   inclusion:
//     namespaces: [ "default" ]
// ```
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=rbac.istio.io/v1alpha1
// +kubetype-gen:kubeType=RbacConfig
// +kubetype-gen:kubeType=ClusterRbacConfig
// +kubetype-gen:ClusterRbacConfig:tag=genclient:nonNamespaced
// +genclient
// +k8s:deepcopy-gen=true
// -->
type ClusterRbacConfig struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1alpha1.RbacConfig `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterRbacConfigList is a collection of ClusterRbacConfigs.
type ClusterRbacConfigList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []ClusterRbacConfig `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceRole specification contains a list of access rules (permissions).
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=rbac.istio.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type ServiceRole struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1alpha1.ServiceRole `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceRoleList is a collection of ServiceRoles.
type ServiceRoleList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []ServiceRole `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceRoleBinding assigns a ServiceRole to a list of subjects.
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=rbac.istio.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type ServiceRoleBinding struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1alpha1.ServiceRoleBinding `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceRoleBindingList is a collection of ServiceRoleBindings.
type ServiceRoleBindingList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []ServiceRoleBinding `json:"items" protobuf:"bytes,2,rep,name=items"`
}
