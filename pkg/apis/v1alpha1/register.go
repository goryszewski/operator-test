package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemaGroupVersion = schema.GroupVersion{
	Group:   "testG",
	Version: "v1alpha1",
}

var SchemaBuilder runtime.SchemeBuilder

func init() {
	SchemaBuilder.Register(addKnownTypes)
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemaGroupVersion, &Cluster{}, &ClusterList{})
	metav1.AddToGroupVersion(scheme, SchemaGroupVersion)
	return nil
}
