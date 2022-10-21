/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

type K8sProvider struct{}

var (
	_ provider.Provider             = (*K8sProvider)(nil)
	_ provider.ProviderWithMetadata = (*K8sProvider)(nil)
)

func New() provider.Provider {
	return &K8sProvider{}
}

func (p *K8sProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "k8s"
}

func (p *K8sProvider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Provider for Kubernetes resources. Requires Terraform 1.0 or later.",
		MarkdownDescription: "Provider for [Kubernetes](https://kubernetes.io/) resources. Requires Terraform 1.0 or later.",
	}, nil
}

func (p *K8sProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
	// NO-OP: provider requires no configuration
}

func (p *K8sProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *K8sProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAcidZalanDoOperatorConfigurationV1Resource,
		NewAcidZalanDoPostgresTeamV1Resource,
		NewAcidZalanDoPostgresqlV1Resource,
		NewAcmeCertManagerIoChallengeV1Resource,
		NewAcmeCertManagerIoOrderV1Resource,
		NewAgentK8SElasticCoAgentV1Alpha1Resource,
		NewApicodegenApimaticIoAPIMaticV1Beta1Resource,
		NewApigatewayv2ServicesK8SAwsAPIV1Alpha1Resource,
		NewApigatewayv2ServicesK8SAwsAuthorizerV1Alpha1Resource,
		NewApigatewayv2ServicesK8SAwsDeploymentV1Alpha1Resource,
		NewApigatewayv2ServicesK8SAwsIntegrationV1Alpha1Resource,
		NewApigatewayv2ServicesK8SAwsRouteV1Alpha1Resource,
		NewApigatewayv2ServicesK8SAwsStageV1Alpha1Resource,
		NewApigatewayv2ServicesK8SAwsVPCLinkV1Alpha1Resource,
		NewApmK8SElasticCoApmServerV1Resource,
		NewAppKiegroupOrgKogitoBuildV1Beta1Resource,
		NewAppKiegroupOrgKogitoInfraV1Beta1Resource,
		NewAppKiegroupOrgKogitoRuntimeV1Beta1Resource,
		NewAppKiegroupOrgKogitoSupportingServiceV1Beta1Resource,
		NewAppLightbendComAkkaClusterV1Alpha1Resource,
		NewApplicationautoscalingServicesK8SAwsScalableTargetV1Alpha1Resource,
		NewApplicationautoscalingServicesK8SAwsScalingPolicyV1Alpha1Resource,
		NewApps3ScaleNetAPIcastV1Alpha1Resource,
		NewAppsGitlabComGitLabV1Beta1Resource,
		NewAppsGitlabComRunnerV1Beta2Resource,
		NewAppsM88IIoNexusV1Alpha1Resource,
		NewAppsRedhatComClusterImpairmentV1Alpha1Resource,
		NewAquasecurityGithubIoAquaStarboardV1Alpha1Resource,
		NewArgoprojIoAppProjectV1Alpha1Resource,
		NewArgoprojIoApplicationSetV1Alpha1Resource,
		NewArgoprojIoApplicationV1Alpha1Resource,
		NewArgoprojIoArgoCDExportV1Alpha1Resource,
		NewArgoprojIoArgoCDV1Alpha1Resource,
		NewAsdbAerospikeComAerospikeClusterV1Beta1Resource,
		NewAutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource,
		NewBeatK8SElasticCoBeatV1Beta1Resource,
		NewCamelApacheOrgBuildV1Resource,
		NewCamelApacheOrgCamelCatalogV1Resource,
		NewCamelApacheOrgIntegrationKitV1Resource,
		NewCamelApacheOrgIntegrationPlatformV1Resource,
		NewCamelApacheOrgIntegrationV1Resource,
		NewCamelApacheOrgKameletBindingV1Alpha1Resource,
		NewCamelApacheOrgKameletV1Alpha1Resource,
		NewCertManagerIoCertificateRequestV1Resource,
		NewCertManagerIoCertificateV1Resource,
		NewCertManagerIoClusterIssuerV1Resource,
		NewCertManagerIoIssuerV1Resource,
		NewChartsFlagsmithComFlagsmithV1Alpha1Resource,
		NewChartsHelmK8SIoSnykMonitorV1Alpha1Resource,
		NewChartsOpdevIoSynapseV1Alpha1Resource,
		NewChartsOperatorhubIoCockroachdbV1Alpha1Resource,
		NewCheEclipseOrgKubernetesImagePullerV1Alpha1Resource,
		NewCiliumIoCiliumBGPLoadBalancerIPPoolV2Alpha1Resource,
		NewCiliumIoCiliumBGPPeeringPolicyV2Alpha1Resource,
		NewCiliumIoCiliumClusterwideEnvoyConfigV2Resource,
		NewCiliumIoCiliumClusterwideNetworkPolicyV2Resource,
		NewCiliumIoCiliumEgressGatewayPolicyV2Resource,
		NewCiliumIoCiliumEgressNATPolicyV2Alpha1Resource,
		NewCiliumIoCiliumEndpointSliceV2Alpha1Resource,
		NewCiliumIoCiliumEnvoyConfigV2Resource,
		NewCiliumIoCiliumExternalWorkloadV2Resource,
		NewCiliumIoCiliumIdentityV2Resource,
		NewCiliumIoCiliumLocalRedirectPolicyV2Resource,
		NewCiliumIoCiliumNetworkPolicyV2Resource,
		NewCiliumIoCiliumNodeV2Resource,
		NewConfigGatekeeperShConfigV1Alpha1Resource,
		NewCoreStrimziIoStrimziPodSetV1Beta2Resource,
		NewCouchbaseComCouchbaseAutoscalerV2Resource,
		NewCouchbaseComCouchbaseBackupRestoreV2Resource,
		NewCouchbaseComCouchbaseBackupV2Resource,
		NewCouchbaseComCouchbaseBucketV2Resource,
		NewCouchbaseComCouchbaseClusterV2Resource,
		NewCouchbaseComCouchbaseCollectionGroupV2Resource,
		NewCouchbaseComCouchbaseCollectionV2Resource,
		NewCouchbaseComCouchbaseEphemeralBucketV2Resource,
		NewCouchbaseComCouchbaseGroupV2Resource,
		NewCouchbaseComCouchbaseMemcachedBucketV2Resource,
		NewCouchbaseComCouchbaseMigrationReplicationV2Resource,
		NewCouchbaseComCouchbaseReplicationV2Resource,
		NewCouchbaseComCouchbaseRoleBindingV2Resource,
		NewCouchbaseComCouchbaseScopeGroupV2Resource,
		NewCouchbaseComCouchbaseScopeV2Resource,
		NewCouchbaseComCouchbaseUserV2Resource,
		NewCrdProjectcalicoOrgBGPConfigurationV1Resource,
		NewCrdProjectcalicoOrgBGPPeerV1Resource,
		NewCrdProjectcalicoOrgBlockAffinityV1Resource,
		NewCrdProjectcalicoOrgCalicoNodeStatusV1Resource,
		NewCrdProjectcalicoOrgClusterInformationV1Resource,
		NewCrdProjectcalicoOrgFelixConfigurationV1Resource,
		NewCrdProjectcalicoOrgGlobalNetworkPolicyV1Resource,
		NewCrdProjectcalicoOrgGlobalNetworkSetV1Resource,
		NewCrdProjectcalicoOrgHostEndpointV1Resource,
		NewCrdProjectcalicoOrgIPAMBlockV1Resource,
		NewCrdProjectcalicoOrgIPAMConfigV1Resource,
		NewCrdProjectcalicoOrgIPAMHandleV1Resource,
		NewCrdProjectcalicoOrgIPPoolV1Resource,
		NewCrdProjectcalicoOrgIPReservationV1Resource,
		NewCrdProjectcalicoOrgKubeControllersConfigurationV1Resource,
		NewCrdProjectcalicoOrgNetworkPolicyV1Resource,
		NewCrdProjectcalicoOrgNetworkSetV1Resource,
		NewDynamodbServicesK8SAwsBackupV1Alpha1Resource,
		NewDynamodbServicesK8SAwsGlobalTableV1Alpha1Resource,
		NewDynamodbServicesK8SAwsTableV1Alpha1Resource,
		NewEc2ServicesK8SAwsDHCPOptionsV1Alpha1Resource,
		NewEc2ServicesK8SAwsElasticIPAddressV1Alpha1Resource,
		NewEc2ServicesK8SAwsInstanceV1Alpha1Resource,
		NewEc2ServicesK8SAwsInternetGatewayV1Alpha1Resource,
		NewEc2ServicesK8SAwsNATGatewayV1Alpha1Resource,
		NewEc2ServicesK8SAwsRouteTableV1Alpha1Resource,
		NewEc2ServicesK8SAwsSecurityGroupV1Alpha1Resource,
		NewEc2ServicesK8SAwsSubnetV1Alpha1Resource,
		NewEc2ServicesK8SAwsTransitGatewayV1Alpha1Resource,
		NewEc2ServicesK8SAwsVPCEndpointV1Alpha1Resource,
		NewEc2ServicesK8SAwsVPCV1Alpha1Resource,
		NewEcrServicesK8SAwsPullThroughCacheRuleV1Alpha1Resource,
		NewEcrServicesK8SAwsRepositoryV1Alpha1Resource,
		NewEksServicesK8SAwsAddonV1Alpha1Resource,
		NewEksServicesK8SAwsClusterV1Alpha1Resource,
		NewEksServicesK8SAwsFargateProfileV1Alpha1Resource,
		NewEksServicesK8SAwsNodegroupV1Alpha1Resource,
		NewElasticacheServicesK8SAwsCacheParameterGroupV1Alpha1Resource,
		NewElasticacheServicesK8SAwsCacheSubnetGroupV1Alpha1Resource,
		NewElasticacheServicesK8SAwsReplicationGroupV1Alpha1Resource,
		NewElasticacheServicesK8SAwsSnapshotV1Alpha1Resource,
		NewElasticacheServicesK8SAwsUserGroupV1Alpha1Resource,
		NewElasticacheServicesK8SAwsUserV1Alpha1Resource,
		NewElasticsearchK8SElasticCoElasticsearchV1Resource,
		NewEmrcontainersServicesK8SAwsJobRunV1Alpha1Resource,
		NewEmrcontainersServicesK8SAwsVirtualClusterV1Alpha1Resource,
		NewEnterprisesearchK8SElasticCoEnterpriseSearchV1Resource,
		NewExecutionFurikoIoJobConfigV1Alpha1Resource,
		NewExecutionFurikoIoJobV1Alpha1Resource,
		NewExtensionsIstioIoWasmPluginV1Alpha1Resource,
		NewExternalSecretsIoClusterExternalSecretV1Beta1Resource,
		NewExternalSecretsIoClusterSecretStoreV1Alpha1Resource,
		NewExternalSecretsIoClusterSecretStoreV1Beta1Resource,
		NewExternalSecretsIoExternalSecretV1Alpha1Resource,
		NewExternalSecretsIoExternalSecretV1Beta1Resource,
		NewExternalSecretsIoSecretStoreV1Alpha1Resource,
		NewExternalSecretsIoSecretStoreV1Beta1Resource,
		NewExternaldataGatekeeperShProviderV1Alpha1Resource,
		NewExternaldnsK8SIoDNSEndpointV1Alpha1Resource,
		NewFlaggerAppAlertProviderV1Beta1Resource,
		NewFlaggerAppCanaryV1Beta1Resource,
		NewFlaggerAppMetricTemplateV1Beta1Resource,
		NewFlinkApacheOrgFlinkDeploymentV1Beta1Resource,
		NewFlinkApacheOrgFlinkSessionJobV1Beta1Resource,
		NewFossulIoBackupConfigV1Resource,
		NewFossulIoBackupScheduleV1Resource,
		NewFossulIoBackupV1Resource,
		NewFossulIoFossulV1Resource,
		NewFossulIoRestoreV1Resource,
		NewGatewayNetworkingK8SIoGatewayClassV1Alpha2Resource,
		NewGatewayNetworkingK8SIoGatewayV1Alpha2Resource,
		NewGatewayNetworkingK8SIoHTTPRouteV1Alpha2Resource,
		NewGatewayNetworkingK8SIoTCPRouteV1Alpha2Resource,
		NewGatewayNetworkingK8SIoTLSRouteV1Alpha2Resource,
		NewGetambassadorIoAuthServiceV2Resource,
		NewGetambassadorIoAuthServiceV3Alpha1Resource,
		NewGetambassadorIoConsulResolverV2Resource,
		NewGetambassadorIoConsulResolverV3Alpha1Resource,
		NewGetambassadorIoDevPortalV2Resource,
		NewGetambassadorIoDevPortalV3Alpha1Resource,
		NewGetambassadorIoHostV2Resource,
		NewGetambassadorIoHostV3Alpha1Resource,
		NewGetambassadorIoKubernetesEndpointResolverV2Resource,
		NewGetambassadorIoKubernetesEndpointResolverV3Alpha1Resource,
		NewGetambassadorIoKubernetesServiceResolverV2Resource,
		NewGetambassadorIoKubernetesServiceResolverV3Alpha1Resource,
		NewGetambassadorIoListenerV3Alpha1Resource,
		NewGetambassadorIoLogServiceV2Resource,
		NewGetambassadorIoLogServiceV3Alpha1Resource,
		NewGetambassadorIoMappingV2Resource,
		NewGetambassadorIoMappingV3Alpha1Resource,
		NewGetambassadorIoModuleV2Resource,
		NewGetambassadorIoModuleV3Alpha1Resource,
		NewGetambassadorIoRateLimitServiceV2Resource,
		NewGetambassadorIoRateLimitServiceV3Alpha1Resource,
		NewGetambassadorIoTCPMappingV2Resource,
		NewGetambassadorIoTCPMappingV3Alpha1Resource,
		NewGetambassadorIoTLSContextV2Resource,
		NewGetambassadorIoTLSContextV3Alpha1Resource,
		NewGetambassadorIoTracingServiceV2Resource,
		NewGetambassadorIoTracingServiceV3Alpha1Resource,
		NewHazelcastComCronHotBackupV1Alpha1Resource,
		NewHazelcastComHazelcastV1Alpha1Resource,
		NewHazelcastComHotBackupV1Alpha1Resource,
		NewHazelcastComManagementCenterV1Alpha1Resource,
		NewHazelcastComMapV1Alpha1Resource,
		NewHazelcastComWanReplicationV1Alpha1Resource,
		NewHelmSigstoreDevRekorV1Alpha1Resource,
		NewHelmToolkitFluxcdIoHelmReleaseV2Beta1Resource,
		NewHiveOpenshiftIoCheckpointV1Resource,
		NewHiveOpenshiftIoClusterClaimV1Resource,
		NewHiveOpenshiftIoClusterDeploymentCustomizationV1Resource,
		NewHiveOpenshiftIoClusterDeploymentV1Resource,
		NewHiveOpenshiftIoClusterDeprovisionV1Resource,
		NewHiveOpenshiftIoClusterImageSetV1Resource,
		NewHiveOpenshiftIoClusterPoolV1Resource,
		NewHiveOpenshiftIoClusterProvisionV1Resource,
		NewHiveOpenshiftIoClusterRelocateV1Resource,
		NewHiveOpenshiftIoClusterStateV1Resource,
		NewHiveOpenshiftIoDNSZoneV1Resource,
		NewHiveOpenshiftIoHiveConfigV1Resource,
		NewHiveOpenshiftIoMachinePoolNameLeaseV1Resource,
		NewHiveOpenshiftIoMachinePoolV1Resource,
		NewHiveOpenshiftIoSelectorSyncIdentityProviderV1Resource,
		NewHiveOpenshiftIoSelectorSyncSetV1Resource,
		NewHiveOpenshiftIoSyncIdentityProviderV1Resource,
		NewHiveOpenshiftIoSyncSetV1Resource,
		NewHiveinternalOpenshiftIoClusterSyncLeaseV1Alpha1Resource,
		NewHiveinternalOpenshiftIoClusterSyncV1Alpha1Resource,
		NewHiveinternalOpenshiftIoFakeClusterInstallV1Alpha1Resource,
		NewHyperfoilIoHorreumV1Alpha1Resource,
		NewHyperfoilIoHyperfoilV1Alpha2Resource,
		NewIamServicesK8SAwsGroupV1Alpha1Resource,
		NewIamServicesK8SAwsPolicyV1Alpha1Resource,
		NewIamServicesK8SAwsRoleV1Alpha1Resource,
		NewIbmcloudIbmComComposableV1Alpha1Resource,
		NewImageToolkitFluxcdIoImagePolicyV1Alpha1Resource,
		NewImageToolkitFluxcdIoImagePolicyV1Alpha2Resource,
		NewImageToolkitFluxcdIoImagePolicyV1Beta1Resource,
		NewImageToolkitFluxcdIoImageRepositoryV1Alpha1Resource,
		NewImageToolkitFluxcdIoImageRepositoryV1Alpha2Resource,
		NewImageToolkitFluxcdIoImageRepositoryV1Beta1Resource,
		NewImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1Resource,
		NewImageToolkitFluxcdIoImageUpdateAutomationV1Alpha2Resource,
		NewImageToolkitFluxcdIoImageUpdateAutomationV1Beta1Resource,
		NewImagingIngestionAlvearieOrgDicomEventBridgeV1Alpha1Resource,
		NewImagingIngestionAlvearieOrgDicomEventDrivenIngestionV1Alpha1Resource,
		NewImagingIngestionAlvearieOrgDicomInstanceBindingV1Alpha1Resource,
		NewImagingIngestionAlvearieOrgDicomStudyBindingV1Alpha1Resource,
		NewImagingIngestionAlvearieOrgDicomwebIngestionServiceV1Alpha1Resource,
		NewImagingIngestionAlvearieOrgDimseIngestionServiceV1Alpha1Resource,
		NewImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource,
		NewInfinispanOrgBackupV2Alpha1Resource,
		NewInfinispanOrgBatchV2Alpha1Resource,
		NewInfinispanOrgCacheV2Alpha1Resource,
		NewInfinispanOrgInfinispanV1Resource,
		NewInfinispanOrgRestoreV2Alpha1Resource,
		NewInstallationMattermostComMattermostV1Beta1Resource,
		NewIntegreatlyOrgGrafanaDashboardV1Alpha1Resource,
		NewIntegreatlyOrgGrafanaDataSourceV1Alpha1Resource,
		NewIntegreatlyOrgGrafanaFolderV1Alpha1Resource,
		NewIntegreatlyOrgGrafanaNotificationChannelV1Alpha1Resource,
		NewIntegreatlyOrgGrafanaV1Alpha1Resource,
		NewIotEclipseOrgDittoV1Alpha1Resource,
		NewIotEclipseOrgHawkbitV1Alpha1Resource,
		NewJaegertracingIoJaegerV1Resource,
		NewK8GbAbsaOssGslbV1Beta1Resource,
		NewKafkaStrimziIoKafkaBridgeV1Beta2Resource,
		NewKafkaStrimziIoKafkaConnectV1Beta2Resource,
		NewKafkaStrimziIoKafkaConnectorV1Beta2Resource,
		NewKafkaStrimziIoKafkaMirrorMaker2V1Beta2Resource,
		NewKafkaStrimziIoKafkaMirrorMakerV1Beta2Resource,
		NewKafkaStrimziIoKafkaRebalanceV1Beta2Resource,
		NewKafkaStrimziIoKafkaTopicV1Beta2Resource,
		NewKafkaStrimziIoKafkaUserV1Beta2Resource,
		NewKafkaStrimziIoKafkaV1Beta2Resource,
		NewKeycloakOrgKeycloakBackupV1Alpha1Resource,
		NewKeycloakOrgKeycloakClientV1Alpha1Resource,
		NewKeycloakOrgKeycloakRealmV1Alpha1Resource,
		NewKeycloakOrgKeycloakUserV1Alpha1Resource,
		NewKeycloakOrgKeycloakV1Alpha1Resource,
		NewKialiIoKialiV1Alpha1Resource,
		NewKibanaK8SElasticCoKibanaV1Resource,
		NewKmsServicesK8SAwsAliasV1Alpha1Resource,
		NewKmsServicesK8SAwsGrantV1Alpha1Resource,
		NewKmsServicesK8SAwsKeyV1Alpha1Resource,
		NewKustomizeToolkitFluxcdIoKustomizationV1Beta1Resource,
		NewKustomizeToolkitFluxcdIoKustomizationV1Beta2Resource,
		NewKyvernoIoAdmissionReportV1Alpha2Resource,
		NewKyvernoIoBackgroundScanReportV1Alpha2Resource,
		NewKyvernoIoClusterAdmissionReportV1Alpha2Resource,
		NewKyvernoIoClusterBackgroundScanReportV1Alpha2Resource,
		NewKyvernoIoClusterPolicyV1Resource,
		NewKyvernoIoClusterPolicyV2Beta1Resource,
		NewKyvernoIoGenerateRequestV1Resource,
		NewKyvernoIoPolicyV1Resource,
		NewKyvernoIoPolicyV2Beta1Resource,
		NewKyvernoIoUpdateRequestV1Beta1Resource,
		NewLambdaServicesK8SAwsAliasV1Alpha1Resource,
		NewLambdaServicesK8SAwsCodeSigningConfigV1Alpha1Resource,
		NewLambdaServicesK8SAwsEventSourceMappingV1Alpha1Resource,
		NewLambdaServicesK8SAwsFunctionURLConfigV1Alpha1Resource,
		NewLambdaServicesK8SAwsFunctionV1Alpha1Resource,
		NewLinkerdIoServiceProfileV1Alpha1Resource,
		NewLinkerdIoServiceProfileV1Alpha2Resource,
		NewLitmuschaosIoChaosEngineV1Alpha1Resource,
		NewLitmuschaosIoChaosExperimentV1Alpha1Resource,
		NewLitmuschaosIoChaosResultV1Alpha1Resource,
		NewMapsK8SElasticCoElasticMapsServerV1Alpha1Resource,
		NewMattermostComClusterInstallationV1Alpha1Resource,
		NewMattermostComMattermostRestoreDBV1Alpha1Resource,
		NewMinioMinIoTenantV1Resource,
		NewMinioMinIoTenantV2Resource,
		NewMonitoringCoreosComAlertmanagerConfigV1Alpha1Resource,
		NewMonitoringCoreosComAlertmanagerV1Resource,
		NewMonitoringCoreosComPodMonitorV1Resource,
		NewMonitoringCoreosComProbeV1Resource,
		NewMonitoringCoreosComPrometheusRuleV1Resource,
		NewMonitoringCoreosComPrometheusV1Resource,
		NewMonitoringCoreosComServiceMonitorV1Resource,
		NewMonitoringCoreosComThanosRulerV1Resource,
		NewMqServicesK8SAwsBrokerV1Alpha1Resource,
		NewMutationsGatekeeperShAssignMetadataV1Alpha1Resource,
		NewMutationsGatekeeperShAssignMetadataV1Beta1Resource,
		NewMutationsGatekeeperShAssignV1Alpha1Resource,
		NewMutationsGatekeeperShAssignV1Beta1Resource,
		NewMutationsGatekeeperShModifySetV1Alpha1Resource,
		NewMutationsGatekeeperShModifySetV1Beta1Resource,
		NewNetworkingIstioIoDestinationRuleV1Alpha3Resource,
		NewNetworkingIstioIoEnvoyFilterV1Alpha3Resource,
		NewNetworkingIstioIoGatewayV1Alpha3Resource,
		NewNetworkingIstioIoGatewayV1Beta1Resource,
		NewNetworkingIstioIoProxyConfigV1Beta1Resource,
		NewNetworkingIstioIoServiceEntryV1Alpha3Resource,
		NewNetworkingIstioIoServiceEntryV1Beta1Resource,
		NewNetworkingIstioIoSidecarV1Alpha3Resource,
		NewNetworkingIstioIoSidecarV1Beta1Resource,
		NewNetworkingIstioIoVirtualServiceV1Alpha3Resource,
		NewNetworkingIstioIoVirtualServiceV1Beta1Resource,
		NewNetworkingIstioIoWorkloadEntryV1Alpha3Resource,
		NewNetworkingIstioIoWorkloadEntryV1Beta1Resource,
		NewNetworkingIstioIoWorkloadGroupV1Alpha3Resource,
		NewNetworkingIstioIoWorkloadGroupV1Beta1Resource,
		NewNotificationToolkitFluxcdIoAlertV1Beta1Resource,
		NewNotificationToolkitFluxcdIoProviderV1Beta1Resource,
		NewNotificationToolkitFluxcdIoReceiverV1Beta1Resource,
		NewOpensearchserviceServicesK8SAwsDomainV1Alpha1Resource,
		NewOpentelemetryIoInstrumentationV1Alpha1Resource,
		NewOpentelemetryIoOpenTelemetryCollectorV1Alpha1Resource,
		NewOperatorAquasecComAquaCspV1Alpha1Resource,
		NewOperatorAquasecComAquaDatabaseV1Alpha1Resource,
		NewOperatorAquasecComAquaEnforcerV1Alpha1Resource,
		NewOperatorAquasecComAquaGatewayV1Alpha1Resource,
		NewOperatorAquasecComAquaKubeEnforcerV1Alpha1Resource,
		NewOperatorAquasecComAquaScannerV1Alpha1Resource,
		NewOperatorAquasecComAquaServerV1Alpha1Resource,
		NewOperatorCryostatIoCryostatV1Beta1Resource,
		NewOperatorKnativeDevKnativeEventingV1Beta1Resource,
		NewOperatorKnativeDevKnativeServingV1Beta1Resource,
		NewOperatorOpenClusterManagementIoClusterManagerV1Resource,
		NewOperatorOpenClusterManagementIoKlusterletV1Resource,
		NewOperatorTektonDevTektonResultV1Alpha1Resource,
		NewOperatorTigeraIoAPIServerV1Resource,
		NewOperatorTigeraIoImageSetV1Resource,
		NewOperatorTigeraIoInstallationV1Resource,
		NewOperatorTigeraIoTigeraStatusV1Resource,
		NewPolicyLinkerdIoAuthorizationPolicyV1Alpha1Resource,
		NewPolicyLinkerdIoHTTPRouteV1Alpha1Resource,
		NewPolicyLinkerdIoHTTPRouteV1Beta1Resource,
		NewPolicyLinkerdIoMeshTLSAuthenticationV1Alpha1Resource,
		NewPolicyLinkerdIoNetworkAuthenticationV1Alpha1Resource,
		NewPolicyLinkerdIoServerAuthorizationV1Alpha1Resource,
		NewPolicyLinkerdIoServerAuthorizationV1Beta1Resource,
		NewPolicyLinkerdIoServerV1Alpha1Resource,
		NewPolicyLinkerdIoServerV1Beta1Resource,
		NewPostgresOperatorCrunchydataComPostgresClusterV1Beta1Resource,
		NewPrometheusserviceServicesK8SAwsAlertManagerDefinitionV1Alpha1Resource,
		NewPrometheusserviceServicesK8SAwsRuleGroupsNamespaceV1Alpha1Resource,
		NewPrometheusserviceServicesK8SAwsWorkspaceV1Alpha1Resource,
		NewQuayRedhatComQuayRegistryV1Resource,
		NewRdsServicesK8SAwsDBClusterParameterGroupV1Alpha1Resource,
		NewRdsServicesK8SAwsDBClusterV1Alpha1Resource,
		NewRdsServicesK8SAwsDBInstanceV1Alpha1Resource,
		NewRdsServicesK8SAwsDBParameterGroupV1Alpha1Resource,
		NewRdsServicesK8SAwsDBProxyV1Alpha1Resource,
		NewRdsServicesK8SAwsDBSubnetGroupV1Alpha1Resource,
		NewRdsServicesK8SAwsGlobalClusterV1Alpha1Resource,
		NewRedhatcopRedhatIoQuayEcosystemV1Alpha1Resource,
		NewRegistryApicurIoApicurioRegistryV1Resource,
		NewRipsawCloudbulldozerIoBenchmarkV1Alpha1Resource,
		NewRocketmqApacheOrgBrokerV1Alpha1Resource,
		NewRocketmqApacheOrgConsoleV1Alpha1Resource,
		NewRocketmqApacheOrgNameServiceV1Alpha1Resource,
		NewRocketmqApacheOrgTopicTransferV1Alpha1Resource,
		NewS3ServicesK8SAwsBucketV1Alpha1Resource,
		NewSagemakerServicesK8SAwsAppV1Alpha1Resource,
		NewSagemakerServicesK8SAwsDataQualityJobDefinitionV1Alpha1Resource,
		NewSagemakerServicesK8SAwsDomainV1Alpha1Resource,
		NewSagemakerServicesK8SAwsEndpointConfigV1Alpha1Resource,
		NewSagemakerServicesK8SAwsEndpointV1Alpha1Resource,
		NewSagemakerServicesK8SAwsFeatureGroupV1Alpha1Resource,
		NewSagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Resource,
		NewSagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1Resource,
		NewSagemakerServicesK8SAwsModelExplainabilityJobDefinitionV1Alpha1Resource,
		NewSagemakerServicesK8SAwsModelPackageGroupV1Alpha1Resource,
		NewSagemakerServicesK8SAwsModelPackageV1Alpha1Resource,
		NewSagemakerServicesK8SAwsModelQualityJobDefinitionV1Alpha1Resource,
		NewSagemakerServicesK8SAwsModelV1Alpha1Resource,
		NewSagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Resource,
		NewSagemakerServicesK8SAwsNotebookInstanceLifecycleConfigV1Alpha1Resource,
		NewSagemakerServicesK8SAwsNotebookInstanceV1Alpha1Resource,
		NewSagemakerServicesK8SAwsProcessingJobV1Alpha1Resource,
		NewSagemakerServicesK8SAwsTrainingJobV1Alpha1Resource,
		NewSagemakerServicesK8SAwsTransformJobV1Alpha1Resource,
		NewSagemakerServicesK8SAwsUserProfileV1Alpha1Resource,
		NewScyllaScylladbComNodeConfigV1Alpha1Resource,
		NewScyllaScylladbComScyllaClusterV1Resource,
		NewScyllaScylladbComScyllaOperatorConfigV1Alpha1Resource,
		NewSecscanQuayRedhatComImageManifestVulnV1Alpha1Resource,
		NewSecurityIstioIoAuthorizationPolicyV1Beta1Resource,
		NewSecurityIstioIoPeerAuthenticationV1Beta1Resource,
		NewSecurityIstioIoRequestAuthenticationV1Beta1Resource,
		NewSecurityProfilesOperatorXK8SIoAppArmorProfileV1Alpha1Resource,
		NewSecurityProfilesOperatorXK8SIoProfileBindingV1Alpha1Resource,
		NewSecurityProfilesOperatorXK8SIoProfileRecordingV1Alpha1Resource,
		NewSecurityProfilesOperatorXK8SIoRawSelinuxProfileV1Alpha2Resource,
		NewSecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1Resource,
		NewSecurityProfilesOperatorXK8SIoSecurityProfileNodeStatusV1Alpha1Resource,
		NewSecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Resource,
		NewSecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2Resource,
		NewServicesK8SAwsAdoptedResourceV1Alpha1Resource,
		NewServicesK8SAwsFieldExportV1Alpha1Resource,
		NewSfnServicesK8SAwsActivityV1Alpha1Resource,
		NewSfnServicesK8SAwsStateMachineV1Alpha1Resource,
		NewSourceToolkitFluxcdIoBucketV1Beta1Resource,
		NewSourceToolkitFluxcdIoBucketV1Beta2Resource,
		NewSourceToolkitFluxcdIoGitRepositoryV1Beta1Resource,
		NewSourceToolkitFluxcdIoGitRepositoryV1Beta2Resource,
		NewSourceToolkitFluxcdIoHelmChartV1Beta1Resource,
		NewSourceToolkitFluxcdIoHelmChartV1Beta2Resource,
		NewSourceToolkitFluxcdIoHelmRepositoryV1Beta1Resource,
		NewSourceToolkitFluxcdIoHelmRepositoryV1Beta2Resource,
		NewSourceToolkitFluxcdIoOCIRepositoryV1Beta2Resource,
		NewSparkoperatorK8SIoScheduledSparkApplicationV1Beta2Resource,
		NewSparkoperatorK8SIoSparkApplicationV1Beta2Resource,
		NewTelemetryIstioIoTelemetryV1Alpha1Resource,
		NewTemplatesGatekeeperShConstraintTemplateV1Alpha1Resource,
		NewTemplatesGatekeeperShConstraintTemplateV1Beta1Resource,
		NewTemplatesGatekeeperShConstraintTemplateV1Resource,
		NewTraefikContainoUsIngressRouteTCPV1Alpha1Resource,
		NewTraefikContainoUsIngressRouteUDPV1Alpha1Resource,
		NewTraefikContainoUsIngressRouteV1Alpha1Resource,
		NewTraefikContainoUsMiddlewareTCPV1Alpha1Resource,
		NewTraefikContainoUsMiddlewareV1Alpha1Resource,
		NewTraefikContainoUsServersTransportV1Alpha1Resource,
		NewTraefikContainoUsTLSOptionV1Alpha1Resource,
		NewTraefikContainoUsTLSStoreV1Alpha1Resource,
		NewTraefikContainoUsTraefikServiceV1Alpha1Resource,
		NewWgpolicyk8SIoClusterPolicyReportV1Alpha2Resource,
		NewWgpolicyk8SIoPolicyReportV1Alpha2Resource,
		NewWildflyOrgWildFlyServerV1Alpha1Resource,
		NewAdmissionregistrationK8SIoMutatingWebhookConfigurationV1Resource,
		NewAdmissionregistrationK8SIoValidatingWebhookConfigurationV1Resource,
		NewAppsDaemonSetV1Resource,
		NewAppsDeploymentV1Resource,
		NewAppsReplicaSetV1Resource,
		NewAppsStatefulSetV1Resource,
		NewAutoscalingHorizontalPodAutoscalerV1Resource,
		NewAutoscalingHorizontalPodAutoscalerV2Resource,
		NewBatchCronJobV1Resource,
		NewBatchJobV1Resource,
		NewCertificatesK8SIoCertificateSigningRequestV1Resource,
		NewConfigMapV1Resource,
		NewDiscoveryK8SIoEndpointSliceV1Resource,
		NewEndpointsV1Resource,
		NewEventsK8SIoEventV1Resource,
		NewFlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource,
		NewFlowcontrolApiserverK8SIoFlowSchemaV1Beta3Resource,
		NewFlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2Resource,
		NewFlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta3Resource,
		NewLimitRangeV1Resource,
		NewNamespaceV1Resource,
		NewNetworkingK8SIoIngressClassV1Resource,
		NewNetworkingK8SIoIngressV1Resource,
		NewNetworkingK8SIoNetworkPolicyV1Resource,
		NewPersistentVolumeClaimV1Resource,
		NewPersistentVolumeV1Resource,
		NewPodV1Resource,
		NewPolicyPodDisruptionBudgetV1Resource,
		NewRbacAuthorizationK8SIoClusterRoleBindingV1Resource,
		NewRbacAuthorizationK8SIoClusterRoleV1Resource,
		NewRbacAuthorizationK8SIoRoleBindingV1Resource,
		NewRbacAuthorizationK8SIoRoleV1Resource,
		NewReplicationControllerV1Resource,
		NewSchedulingK8SIoPriorityClassV1Resource,
		NewSecretV1Resource,
		NewServiceAccountV1Resource,
		NewServiceV1Resource,
		NewStorageK8SIoCSIDriverV1Resource,
		NewStorageK8SIoCSINodeV1Resource,
		NewStorageK8SIoStorageClassV1Resource,
		NewStorageK8SIoVolumeAttachmentV1Resource,
	}
}
