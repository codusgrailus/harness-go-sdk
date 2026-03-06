package nextgen

// IacmExecutionConfig contains the container image tags used for IACM pipeline steps
type IacmExecutionConfig struct {
	AddonTag                            string `json:"addonTag,omitempty"`
	LiteEngineTag                       string `json:"liteEngineTag,omitempty"`
	AddonTagRootless                    string `json:"addonTagRootless,omitempty"`
	LiteEngineTagRootless               string `json:"liteEngineTagRootless,omitempty"`
	GitCloneTag                         string `json:"gitCloneTag,omitempty"`
	BuildAndPushDockerRegistryTag       string `json:"buildAndPushDockerRegistryTag,omitempty"`
	BuildAndPushECRTag                  string `json:"buildAndPushECRTag,omitempty"`
	BuildAndPushACRTag                  string `json:"buildAndPushACRTag,omitempty"`
	BuildAndPushGCRTag                  string `json:"buildAndPushGCRTag,omitempty"`
	BuildAndPushGARTag                  string `json:"buildAndPushGARTag,omitempty"`
	BuildAndPushBuildxDockerRegistryTag string `json:"buildAndPushBuildxDockerRegistryTag,omitempty"`
	BuildAndPushBuildxECRTag            string `json:"buildAndPushBuildxECRTag,omitempty"`
	BuildAndPushBuildxGARTag            string `json:"buildAndPushBuildxGARTag,omitempty"`
	BuildAndPushBuildxACRTag            string `json:"buildAndPushBuildxACRTag,omitempty"`
	GcsUploadTag                        string `json:"gcsUploadTag,omitempty"`
	S3UploadTag                         string `json:"s3UploadTag,omitempty"`
	ArtifactoryUploadTag                string `json:"artifactoryUploadTag,omitempty"`
	HarUploadTag                        string `json:"harUploadTag,omitempty"`
	CacheGCSTag                         string `json:"cacheGCSTag,omitempty"`
	CacheS3Tag                          string `json:"cacheS3Tag,omitempty"`
	CacheAzureTag                       string `json:"cacheAzureTag,omitempty"`
	CacheTag                            string `json:"cacheTag,omitempty"`
	CacheProxyImage                     string `json:"cacheProxyImage,omitempty"`
	SecurityTag                         string `json:"securityTag,omitempty"`
	SscaOrchestrationTag                string `json:"sscaOrchestrationTag,omitempty"`
	SscaEnforcementTag                  string `json:"sscaEnforcementTag,omitempty"`
	SscaArtifactSigningTag              string `json:"sscaArtifactSigningTag,omitempty"`
	SscaArtifactVerificationTag         string `json:"sscaArtifactVerificationTag,omitempty"`
	SscaCdxgenOrchestrationTag          string `json:"sscaCdxgenOrchestrationTag,omitempty"`
	ProvenanceTag                       string `json:"provenanceTag,omitempty"`
	SlsaVerificationTag                 string `json:"slsaVerificationTag,omitempty"`
	SscaComplianceTag                   string `json:"sscaComplianceTag,omitempty"`
	IacmTerraform                       string `json:"iacmTerraform,omitempty"`
	IacmTerragrunt                      string `json:"iacmTerragrunt,omitempty"`
	IacmAwsCdk                          string `json:"iacmAwsCdk,omitempty"`
	IacmAnsible                         string `json:"iacmAnsible,omitempty"`
	IacmOpenTofu                        string `json:"iacmOpenTofu,omitempty"`
	IacmCheckov                         string `json:"iacmCheckov,omitempty"`
	IacmTFCompliance                    string `json:"iacmTFCompliance,omitempty"`
	IacmTFLint                          string `json:"iacmTFLint,omitempty"`
	IacmTFSec                           string `json:"iacmTFSec,omitempty"`
	IacmModuleTest                      string `json:"iacmModuleTest,omitempty"`
	CookieCutter                        string `json:"cookieCutter,omitempty"`
	CreateRepo                          string `json:"createRepo,omitempty"`
	DirectPush                          string `json:"directPush,omitempty"`
	RegisterCatalog                     string `json:"registerCatalog,omitempty"`
	CreateCatalog                       string `json:"createCatalog,omitempty"`
	SlackNotify                         string `json:"slackNotify,omitempty"`
	CreateOrganisation                  string `json:"createOrganisation,omitempty"`
	CreateProject                       string `json:"createProject,omitempty"`
	CreateResource                      string `json:"createResource,omitempty"`
	UpdateCatalogProperty               string `json:"updateCatalogProperty,omitempty"`
}

// IacmExecutionConfigResponse is the API response wrapper for the IACM execution config endpoint
type IacmExecutionConfigResponse struct {
	Status        string               `json:"status,omitempty"`
	Data          *IacmExecutionConfig `json:"data,omitempty"`
	MetaData      interface{}          `json:"metaData,omitempty"`
	CorrelationId string               `json:"correlationId,omitempty"`
}

// IacmExecutionConfigUpdate represents a single image field update for the update-config endpoint
type IacmExecutionConfigUpdate struct {
	Field string `json:"field"`
	Value string `json:"value,omitempty"`
}
