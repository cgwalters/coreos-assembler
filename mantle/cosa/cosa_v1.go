package cosa

// generated by "schematyper ../src/schema/v1.json -o cosa/cosa_v1.go.tmp --package=cosa --root-type=Build --ptr-for-omit" -- DO NOT EDIT

type AliyunImage struct {
	ImageID string `json:"id"`
	Region  string `json:"name"`
}

type Amis struct {
	Hvm      string `json:"hvm"`
	Region   string `json:"name"`
	Snapshot string `json:"snapshot"`
}

type Artifact struct {
	Path               string `json:"path"`
	Sha256             string `json:"sha256"`
	SizeInBytes        int    `json:"size,omitempty"`
	UncompressedSha256 string `json:"uncompressed-sha256,omitempty"`
	UncompressedSize   int    `json:"uncompressed-size,omitempty"`
}

type Build struct {
	AlibabaAliyunUploads      []AliyunImage           `json:"aliyun,omitempty"`
	Amis                      []Amis                  `json:"amis,omitempty"`
	Architecture              string                  `json:"coreos-assembler.basearch,omitempty"`
	Azure                     *Cloudartifact          `json:"azure,omitempty"`
	BuildArtifacts            *BuildArtifacts         `json:"images,omitempty"`
	BuildID                   string                  `json:"buildid"`
	BuildRef                  string                  `json:"ref,omitempty"`
	BuildSummary              string                  `json:"summary"`
	BuildTimeStamp            string                  `json:"coreos-assembler.build-timestamp,omitempty"`
	BuildURL                  string                  `json:"build-url,omitempty"`
	ConfigGitRev              string                  `json:"coreos-assembler.config-gitrev,omitempty"`
	ContainerConfigGit        *Git                    `json:"coreos-assembler.container-config-git,omitempty"`
	CoreOsSource              string                  `json:"coreos-assembler.code-source,omitempty"`
	CosaContainerImageGit     *Git                    `json:"coreos-assembler.container-image-git,omitempty"`
	CosaImageChecksum         string                  `json:"coreos-assembler.image-config-checksum,omitempty"`
	CosaImageVersion          int                     `json:"coreos-assembler.image-genver,omitempty"`
	FedoraCoreOsParentCommit  string                  `json:"fedora-coreos.parent-commit,omitempty"`
	FedoraCoreOsParentVersion string                  `json:"fedora-coreos.parent-version,omitempty"`
	Gcp                       *Cloudartifact          `json:"gcp,omitempty"`
	GitDirty                  string                  `json:"coreos-assembler.config-dirty,omitempty"`
	ImageInputChecksum        string                  `json:"coreos-assembler.image-input-checksum,omitempty"`
	InputHasOfTheRpmOstree    string                  `json:"rpm-ostree-inputhash"`
	Name                      string                  `json:"name"`
	Oscontainer               *Image                  `json:"oscontainer,omitempty"`
	OstreeCommit              string                  `json:"ostree-commit"`
	OstreeContentBytesWritten int                     `json:"ostree-content-bytes-written"`
	OstreeContentChecksum     string                  `json:"ostree-content-checksum"`
	OstreeNCacheHits          int                     `json:"ostree-n-cache-hits"`
	OstreeNContentTotal       int                     `json:"ostree-n-content-total"`
	OstreeNContentWritten     int                     `json:"ostree-n-content-written"`
	OstreeNMetadataTotal      int                     `json:"ostree-n-metadata-total"`
	OstreeNMetadataWritten    int                     `json:"ostree-n-metadata-written"`
	OstreeTimestamp           string                  `json:"ostree-timestamp"`
	OstreeVersion             string                  `json:"ostree-version"`
	OverridesActive           bool                    `json:"coreos-assembler.overrides-active,omitempty"`
	PkgdiffBetweenBuilds      []PackageSetDifferences `json:"pkgdiff,omitempty"`
}

type BuildArtifacts struct {
	Aliyun        *Artifact `json:"aliyun,omitempty"`
	Aws           *Artifact `json:"aws,omitempty"`
	Azure         *Artifact `json:"azure,omitempty"`
	Dasd          *Artifact `json:"dasd,omitempty"`
	Exoscale      *Artifact `json:"exoscale,omitempty"`
	Gcp           *Artifact `json:"gcp,omitempty"`
	Initramfs     *Artifact `json:"initramfs,omitempty"`
	Iso           *Artifact `json:"iso,omitempty"`
	Kernel        *Artifact `json:"kernel,omitempty"`
	LiveInitramfs *Artifact `json:"live-initramfs,omitempty"`
	LiveIso       *Artifact `json:"live-iso,omitempty"`
	LiveKernel    *Artifact `json:"live-kernel,omitempty"`
	Metal         *Artifact `json:"metal,omitempty"`
	Metal4KNative *Artifact `json:"metal4k,omitempty"`
	OpenStack     *Artifact `json:"openstack,omitempty"`
	Ostree        Artifact  `json:"ostree"`
	Qemu          *Artifact `json:"qemu,omitempty"`
	Vmware        *Artifact `json:"vmware,omitempty"`
}

type Cloudartifact struct {
	Image string `json:"image"`
	URL   string `json:"url"`
}

type Git struct {
	Branch string `json:"branch,omitempty"`
	Commit string `json:"commit"`
	Dirty  string `json:"dirty,omitempty"`
	Origin string `json:"origin"`
}

type Image struct {
	Digest string `json:"digest"`
	Image  string `json:"image"`
}

type Items interface{}

type PackageSetDifferences []Items
