package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClusterState defines states for a cluster.
const (
	ClusterStateCreating         = "Creating"
	ClusterStateRunning          = "Running"
	ClusterStateReconciling      = "Reconciling"
	ClusterStateStopping         = "Stopping"
	ClusterStatePartiallyStopped = "PartiallyStopped"
	ClusterStateStopped          = "Stopped"
)

// ComponentState defines states for a cluster component.
const (
	ComponentStateNotReady = "NotReady"
	ComponentStateReady    = "Ready"
	ComponentStateDeleted  = "Deleted"
)

// JobState defines states for a Flink job.
const (
	JobStatePending   = "Pending"
	JobStateRunning   = "Running"
	JobStateSucceeded = "Succeeded"
	JobStateFailed    = "Failed"
	JobStateCancelled = "Cancelled"
	JobStateUnknown   = "Unknown"
)

// JobRestartPolicy defines the restart policy when a job fails.
type JobRestartPolicy = string

const (
	// JobRestartPolicyNever - never restarts a failed job.
	JobRestartPolicyNever = "Never"

	// JobRestartPolicyFromSavepointOnFailure - restart the job from the latest
	// savepoint if available, otherwise do not restart.
	JobRestartPolicyFromSavepointOnFailure = "FromSavepointOnFailure"
)

// User requested control
const (
	// control annotation key
	ControlAnnotation = "flinkoperator.k8s.io/user-control"

	// control name
	ControlNameCancel    = "cancel"
	ControlNameSavepoint = "savepoint"

	// control state
	ControlStateProgressing = "Progressing"
	ControlStateSucceeded   = "Succeeded"
	ControlStateFailed      = "Failed"
)

// ImageSpec defines Flink image of JobManager and TaskManager containers.
type ImageSpec struct {
	// Flink image name.
	Name string `json:"name"`

	// Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always
	// if :latest tag is specified, or IfNotPresent otherwise.
	PullPolicy corev1.PullPolicy `json:"pullPolicy,omitempty"`

	// Secrets for image pull.
	PullSecrets []corev1.LocalObjectReference `json:"pullSecrets,omitempty"`
}

// JobManagerPorts defines ports of JobManager.
type JobManagerPorts struct {
	// RPC port, default: 6123.
	RPC *int32 `json:"rpc,omitempty"`

	// Blob port, default: 6124.
	Blob *int32 `json:"blob,omitempty"`

	// Query port, default: 6125.
	Query *int32 `json:"query,omitempty"`

	// UI port, default: 8081.
	UI *int32 `json:"ui,omitempty"`
}

// JobManagerSpec defines properties of JobManager.
type JobManagerSpec struct {
	// The number of replicas.
	Replicas *int32 `json:"replicas,omitempty"`

	// Ports.
	Ports JobManagerPorts `json:"ports,omitempty"`

	// Compute resources required by each JobManager container.
	// If omitted, a default value will be used.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// TODO: Memory calculation would be change. Let's watch the issue FLINK-13980.

	// Volumes in the JobManager pod.
	Volumes []corev1.Volume `json:"volumes,omitempty"`

	// Volume mounts in the JobManager container.
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`

	// Selector which must match a node's labels for the JobManager pod to be
	// scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Sidecar containers running alongside with the JobManager container in the
	// pod.
	Sidecars []corev1.Container `json:"sidecars,omitempty"`
}

// TaskManagerPorts defines ports of TaskManager.
type TaskManagerPorts struct {
	// Data port, default: 6121.
	Data *int32 `json:"data,omitempty"`

	// RPC port, default: 6122.
	RPC *int32 `json:"rpc,omitempty"`

	// Query port.
	Query *int32 `json:"query,omitempty"`
}

// TaskManagerSpec defines properties of TaskManager.
type TaskManagerSpec struct {
	// The number of replicas.
	Replicas int32 `json:"replicas"`

	// Ports.
	Ports TaskManagerPorts `json:"ports,omitempty"`

	// Compute resources required by each TaskManager container.
	// If omitted, a default value will be used.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// TODO: Memory calculation would be change. Let's watch the issue FLINK-13980.

	// Volumes in the TaskManager pods.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes/
	Volumes []corev1.Volume `json:"volumes,omitempty"`

	// Volume mounts in the TaskManager containers.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes/
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`

	// Selector which must match a node's labels for the TaskManager pod to be
	// scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Sidecar containers running alongside with the TaskManager container in the
	// pod.
	Sidecars []corev1.Container `json:"sidecars,omitempty"`
}

// CleanupAction defines the action to take after job finishes.
type CleanupAction string

const (
	// CleanupActionKeepCluster - keep the entire cluster.
	CleanupActionKeepCluster = "KeepCluster"
	// CleanupActionDeleteCluster - delete the entire cluster.
	CleanupActionDeleteCluster = "DeleteCluster"
	// CleanupActionDeleteTaskManager - delete task manager, keep job manager.
	CleanupActionDeleteTaskManager = "DeleteTaskManager"
)

// CleanupPolicy defines the action to take after job finishes.
type CleanupPolicy struct {
	// Action to take after job succeeds.
	AfterJobSucceeds CleanupAction `json:"afterJobSucceeds,omitempty"`
	// Action to take after job fails.
	AfterJobFails CleanupAction `json:"afterJobFails,omitempty"`
	// Action to take after job is cancelled.
	AfterJobCancelled CleanupAction `json:"afterJobCancelled,omitempty"`
}

// JobSpec defines properties of a Flink job.
type JobSpec struct {
	// Args of the job.
	Args []string `json:"args,omitempty"`

	// FromSavepoint where to restore the job from (e.g., gs://my-savepoint/1234).
	FromSavepoint *string `json:"fromSavepoint,omitempty"`

	// Allow non-restored state, default: false.
	AllowNonRestoredState *bool `json:"allowNonRestoredState,omitempty"`

	// Savepoints dir where to store savepoints of the job.
	SavepointsDir *string `json:"savepointsDir,omitempty"`

	// Automatically take a savepoint to the `savepointsDir` every n seconds.
	AutoSavepointSeconds *int32 `json:"autoSavepointSeconds,omitempty"`

	// Update this field to `jobStatus.savepointGeneration + 1` for a running job
	// cluster to trigger a new savepoint to `savepointsDir` on demand.
	SavepointGeneration int32 `json:"savepointGeneration,omitempty"`

	// Job parallelism, default: 1.
	Parallelism *int32 `json:"parallelism,omitempty"`

	// No logging output to STDOUT, default: false.
	NoLoggingToStdout *bool `json:"noLoggingToStdout,omitempty"`

	// Volumes in the Job pod.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes/
	Volumes []corev1.Volume `json:"volumes,omitempty"`

	// Volume mounts in the Job container.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes/
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`

	// Init containers of the Job pod. A typical use case could be using an init
	// container to download a remote job jar to a local path which is
	// referenced by the `jarFile` property.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	InitContainers []corev1.Container `json:"initContainers,omitempty"`

	// Restart policy when the job fails, "Never" or "FromSavepointOnFailure",
	// default: "Never".
	//
	// "Never" means the operator will never try to restart a failed job, manual
	// cleanup and restart is required.
	//
	// "FromSavepointOnFailure" means the operator will try to restart the failed
	// job from the savepoint recorded in the job status if available; otherwise,
	// the job will stay in failed state. This option is usually used together
	// with `autoSavepointSeconds` and `savepointsDir`.
	RestartPolicy *JobRestartPolicy `json:"restartPolicy"`

	// The action to take after job finishes.
	CleanupPolicy *CleanupPolicy `json:"cleanupPolicy,omitempty"`

	// Request the job to be cancelled. Only applies to running jobs. If
	// `savePointsDir` is provided, a savepoint will be taken before stopping the
	// job.
	CancelRequested *bool `json:"cancelRequested,omitempty"`
}

// FlinkClusterSpec defines the desired state of FlinkCluster
type FlinkClusterSpec struct {
	// Flink image spec for the cluster's components.
	Image ImageSpec `json:"image"`

	// Flink JobManager spec.
	JobManager JobManagerSpec `json:"jobManager"`

	// Flink TaskManager spec.
	TaskManager TaskManagerSpec `json:"taskManager"`

	// (Optional) Job spec. If specified, this cluster is an ephemeral Job
	// Cluster, which will be automatically terminated after the job finishes;
	// otherwise, it is a long-running Session Cluster.
	Job *JobSpec `json:"job,omitempty"`

	// Environment variables shared by all JobManager, TaskManager and job
	// containers.
	EnvVars []corev1.EnvVar `json:"envVars,omitempty"`

	// Flink properties which are appened to flink-conf.yaml.
	FlinkProperties map[string]string `json:"flinkProperties,omitempty"`

	// Config for Hadoop.
	HadoopConfig *HadoopConfig `json:"hadoopConfig,omitempty"`

	// Config for GCP.
	AzureBlobConfig *AzureBlobConfig `json:"azureBlobConfig,omitempty"`
}

// HadoopConfig defines configs for Hadoop.
type HadoopConfig struct {
	// The name of the ConfigMap which contains the Hadoop config files.
	// The ConfigMap must be in the same namespace as the FlinkCluster.
	ConfigMapName string `json:"configMapName,omitempty"`

	// The path where to mount the Volume of the ConfigMap.
	MountPath string `json:"mountPath,omitempty"`
}

// AzureBlobConfig defines configs for Azure Storage Account.
type AzureBlobConfig struct {
	// GCP service account.
	SecretName *string `json:"secretName"`
}

// FlinkClusterComponentState defines the observed state of a component
// of a FlinkCluster.
type FlinkClusterComponentState struct {
	// The resource name of the component.
	Name string `json:"name"`

	// The state of the component.
	State string `json:"state"`
}

// FlinkClusterComponentsStatus defines the observed status of the
// components of a FlinkCluster.
type FlinkClusterComponentsStatus struct {
	// The state of configMap.
	ConfigMap FlinkClusterComponentState `json:"configMap"`

	// The state of JobManager deployment.
	JobManagerDeployment FlinkClusterComponentState `json:"jobManagerDeployment"`

	// The state of JobManager service.
	JobManagerService JobManagerServiceStatus `json:"jobManagerService"`

	// The state of TaskManager deployment.
	TaskManagerDeployment FlinkClusterComponentState `json:"taskManagerDeployment"`

	// The status of the job, available only when JobSpec is provided.
	Job *JobStatus `json:"job,omitempty"`
}

// FlinkClusterControlState defines the user control state
type FlinkClusterControlState struct {
	// Control name
	Name string `json:"name"`

	// Control data
	Details map[string]string `json:"details,omitempty"`

	// State
	State string `json:"state"`

	// Message
	Message string `json:"message,omitempty"`

	// State update time
	UpdateTime string `json:"updateTime"`
}

// JobStatus defines the status of a job.
type JobStatus struct {
	// The name of the Kubernetes job resource.
	Name string `json:"name"`

	// The ID of the Flink job.
	ID string `json:"id"`

	// The state of the Kubernetes job.
	State string `json:"state"`

	// The actual savepoint from which this job started.
	// In case of restart, it might be different from the savepoint in the job
	// spec.
	FromSavepoint string `json:"fromSavepoint,omitempty"`

	// The generation of the savepoint in `savepointsDir` taken by the operator.
	// The value starts from 0 when there is no savepoint and increases by 1 for
	// each successful savepoint.
	SavepointGeneration int32 `json:"savepointGeneration,omitempty"`

	// Savepoint location.
	SavepointLocation string `json:"savepointLocation,omitempty"`

	// Last savepoint trigger ID.
	LastSavepointTriggerID string `json:"lastSavepointTriggerID,omitempty"`

	// Last successful or failed savepoint operation timestamp.
	LastSavepointTime string `json:"lastSavepointTime,omitempty"`

	// The number of restarts.
	RestartCount int32 `json:"restartCount,omitempty"`
}

// JobManagerServiceStatus defines the observed state of FlinkCluster
type JobManagerServiceStatus struct {
	// The name of the Kubernetes jobManager service.
	Name string `json:"name"`

	// The state of the component.
	State string `json:"state"`

	// (Optional) The node port, present when `accessScope` is `NodePort`.
	NodePort int32 `json:"nodePort,omitempty"`
}

// FlinkClusterStatus defines the observed state of FlinkCluster
type FlinkClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The overall state of the Flink cluster.
	State string `json:"state"`

	// The status of the components.
	Components FlinkClusterComponentsStatus `json:"components"`

	// The status of control requested by user
	Control *FlinkClusterControlState `json:"control,omitempty"`

	// Last update timestamp for this status.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FlinkCluster is the Schema for the flinkclusters API
// +kubebuilder:subresource:status
// +kubebuilder:storageversion:v1alpha1
// +kubebuilder:resource:path=flinkclusters,scope=Namespaced
type FlinkCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FlinkClusterSpec   `json:"spec,omitempty"`
	Status FlinkClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FlinkClusterList contains a list of FlinkCluster
type FlinkClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlinkCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FlinkCluster{}, &FlinkClusterList{})
}
