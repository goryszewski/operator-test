package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

func (in *Cluster) DeepCopyInto() {}
func (in *Cluster) DeepCopy()     {}
func (in *Cluster) DeepCopyObject() runtime.Object {
	return nil
}

func (in *ClusterList) DeepCopyInto() {}
func (in *ClusterList) DeepCopy()     {}
func (in *ClusterList) DeepCopyObject() runtime.Object {
	return nil
}
