package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Cluster struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec ClusterSpec
}

type ClusterSpec struct {
	Name    string
	Version string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterList struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	items []Cluster
}
