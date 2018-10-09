package client

const (
	StorageSpecType                     = "storageSpec"
	StorageSpecFieldClass               = "class"
	StorageSpecFieldEmptyDir            = "emptyDir"
	StorageSpecFieldResources           = "resources"
	StorageSpecFieldSelector            = "selector"
	StorageSpecFieldVolumeClaimTemplate = "volumeClaimTemplate"
)

type StorageSpec struct {
	Class               string                 `json:"class,omitempty" yaml:"class,omitempty"`
	EmptyDir            *EmptyDirVolumeSource  `json:"emptyDir,omitempty" yaml:"emptyDir,omitempty"`
	Resources           *ResourceRequirements  `json:"resources,omitempty" yaml:"resources,omitempty"`
	Selector            *LabelSelector         `json:"selector,omitempty" yaml:"selector,omitempty"`
	VolumeClaimTemplate *PersistentVolumeClaim `json:"volumeClaimTemplate,omitempty" yaml:"volumeClaimTemplate,omitempty"`
}
