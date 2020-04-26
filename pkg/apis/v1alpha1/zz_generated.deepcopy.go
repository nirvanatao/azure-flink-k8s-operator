// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureBlobConfig) DeepCopyInto(out *AzureBlobConfig) {
	*out = *in
	if in.SecretName != nil {
		in, out := &in.SecretName, &out.SecretName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureBlobConfig.
func (in *AzureBlobConfig) DeepCopy() *AzureBlobConfig {
	if in == nil {
		return nil
	}
	out := new(AzureBlobConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CleanupPolicy) DeepCopyInto(out *CleanupPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CleanupPolicy.
func (in *CleanupPolicy) DeepCopy() *CleanupPolicy {
	if in == nil {
		return nil
	}
	out := new(CleanupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkCluster) DeepCopyInto(out *FlinkCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkCluster.
func (in *FlinkCluster) DeepCopy() *FlinkCluster {
	if in == nil {
		return nil
	}
	out := new(FlinkCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlinkCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkClusterComponentState) DeepCopyInto(out *FlinkClusterComponentState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkClusterComponentState.
func (in *FlinkClusterComponentState) DeepCopy() *FlinkClusterComponentState {
	if in == nil {
		return nil
	}
	out := new(FlinkClusterComponentState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkClusterComponentsStatus) DeepCopyInto(out *FlinkClusterComponentsStatus) {
	*out = *in
	out.ConfigMap = in.ConfigMap
	out.JobManagerDeployment = in.JobManagerDeployment
	out.JobManagerService = in.JobManagerService
	out.TaskManagerDeployment = in.TaskManagerDeployment
	if in.Job != nil {
		in, out := &in.Job, &out.Job
		*out = new(JobStatus)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkClusterComponentsStatus.
func (in *FlinkClusterComponentsStatus) DeepCopy() *FlinkClusterComponentsStatus {
	if in == nil {
		return nil
	}
	out := new(FlinkClusterComponentsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkClusterControlState) DeepCopyInto(out *FlinkClusterControlState) {
	*out = *in
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkClusterControlState.
func (in *FlinkClusterControlState) DeepCopy() *FlinkClusterControlState {
	if in == nil {
		return nil
	}
	out := new(FlinkClusterControlState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkClusterList) DeepCopyInto(out *FlinkClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlinkCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkClusterList.
func (in *FlinkClusterList) DeepCopy() *FlinkClusterList {
	if in == nil {
		return nil
	}
	out := new(FlinkClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlinkClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkClusterSpec) DeepCopyInto(out *FlinkClusterSpec) {
	*out = *in
	in.Image.DeepCopyInto(&out.Image)
	in.JobManager.DeepCopyInto(&out.JobManager)
	in.TaskManager.DeepCopyInto(&out.TaskManager)
	if in.Job != nil {
		in, out := &in.Job, &out.Job
		*out = new(JobSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FlinkProperties != nil {
		in, out := &in.FlinkProperties, &out.FlinkProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HadoopConfig != nil {
		in, out := &in.HadoopConfig, &out.HadoopConfig
		*out = new(HadoopConfig)
		**out = **in
	}
	if in.AzureBlobConfig != nil {
		in, out := &in.AzureBlobConfig, &out.AzureBlobConfig
		*out = new(AzureBlobConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkClusterSpec.
func (in *FlinkClusterSpec) DeepCopy() *FlinkClusterSpec {
	if in == nil {
		return nil
	}
	out := new(FlinkClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkClusterStatus) DeepCopyInto(out *FlinkClusterStatus) {
	*out = *in
	in.Components.DeepCopyInto(&out.Components)
	if in.Control != nil {
		in, out := &in.Control, &out.Control
		*out = new(FlinkClusterControlState)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkClusterStatus.
func (in *FlinkClusterStatus) DeepCopy() *FlinkClusterStatus {
	if in == nil {
		return nil
	}
	out := new(FlinkClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HadoopConfig) DeepCopyInto(out *HadoopConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HadoopConfig.
func (in *HadoopConfig) DeepCopy() *HadoopConfig {
	if in == nil {
		return nil
	}
	out := new(HadoopConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.PullSecrets != nil {
		in, out := &in.PullSecrets, &out.PullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobManagerPorts) DeepCopyInto(out *JobManagerPorts) {
	*out = *in
	if in.RPC != nil {
		in, out := &in.RPC, &out.RPC
		*out = new(int32)
		**out = **in
	}
	if in.Blob != nil {
		in, out := &in.Blob, &out.Blob
		*out = new(int32)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(int32)
		**out = **in
	}
	if in.UI != nil {
		in, out := &in.UI, &out.UI
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobManagerPorts.
func (in *JobManagerPorts) DeepCopy() *JobManagerPorts {
	if in == nil {
		return nil
	}
	out := new(JobManagerPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobManagerServiceStatus) DeepCopyInto(out *JobManagerServiceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobManagerServiceStatus.
func (in *JobManagerServiceStatus) DeepCopy() *JobManagerServiceStatus {
	if in == nil {
		return nil
	}
	out := new(JobManagerServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobManagerSpec) DeepCopyInto(out *JobManagerSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Ports.DeepCopyInto(&out.Ports)
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobManagerSpec.
func (in *JobManagerSpec) DeepCopy() *JobManagerSpec {
	if in == nil {
		return nil
	}
	out := new(JobManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSpec) DeepCopyInto(out *JobSpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FromSavepoint != nil {
		in, out := &in.FromSavepoint, &out.FromSavepoint
		*out = new(string)
		**out = **in
	}
	if in.AllowNonRestoredState != nil {
		in, out := &in.AllowNonRestoredState, &out.AllowNonRestoredState
		*out = new(bool)
		**out = **in
	}
	if in.SavepointsDir != nil {
		in, out := &in.SavepointsDir, &out.SavepointsDir
		*out = new(string)
		**out = **in
	}
	if in.AutoSavepointSeconds != nil {
		in, out := &in.AutoSavepointSeconds, &out.AutoSavepointSeconds
		*out = new(int32)
		**out = **in
	}
	if in.Parallelism != nil {
		in, out := &in.Parallelism, &out.Parallelism
		*out = new(int32)
		**out = **in
	}
	if in.NoLoggingToStdout != nil {
		in, out := &in.NoLoggingToStdout, &out.NoLoggingToStdout
		*out = new(bool)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RestartPolicy != nil {
		in, out := &in.RestartPolicy, &out.RestartPolicy
		*out = new(string)
		**out = **in
	}
	if in.CleanupPolicy != nil {
		in, out := &in.CleanupPolicy, &out.CleanupPolicy
		*out = new(CleanupPolicy)
		**out = **in
	}
	if in.CancelRequested != nil {
		in, out := &in.CancelRequested, &out.CancelRequested
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSpec.
func (in *JobSpec) DeepCopy() *JobSpec {
	if in == nil {
		return nil
	}
	out := new(JobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobStatus) DeepCopyInto(out *JobStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobStatus.
func (in *JobStatus) DeepCopy() *JobStatus {
	if in == nil {
		return nil
	}
	out := new(JobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskManagerPorts) DeepCopyInto(out *TaskManagerPorts) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(int32)
		**out = **in
	}
	if in.RPC != nil {
		in, out := &in.RPC, &out.RPC
		*out = new(int32)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskManagerPorts.
func (in *TaskManagerPorts) DeepCopy() *TaskManagerPorts {
	if in == nil {
		return nil
	}
	out := new(TaskManagerPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskManagerSpec) DeepCopyInto(out *TaskManagerSpec) {
	*out = *in
	in.Ports.DeepCopyInto(&out.Ports)
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskManagerSpec.
func (in *TaskManagerSpec) DeepCopy() *TaskManagerSpec {
	if in == nil {
		return nil
	}
	out := new(TaskManagerSpec)
	in.DeepCopyInto(out)
	return out
}