/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type CertManagerIoClusterIssuerV1Resource struct{}

var (
	_ resource.Resource = (*CertManagerIoClusterIssuerV1Resource)(nil)
)

type CertManagerIoClusterIssuerV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CertManagerIoClusterIssuerV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Acme *struct {
			DisableAccountKeyGeneration *bool `tfsdk:"disable_account_key_generation" yaml:"disableAccountKeyGeneration,omitempty"`

			Email *string `tfsdk:"email" yaml:"email,omitempty"`

			EnableDurationFeature *bool `tfsdk:"enable_duration_feature" yaml:"enableDurationFeature,omitempty"`

			ExternalAccountBinding *struct {
				KeyAlgorithm *string `tfsdk:"key_algorithm" yaml:"keyAlgorithm,omitempty"`

				KeyID *string `tfsdk:"key_id" yaml:"keyID,omitempty"`

				KeySecretRef *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"key_secret_ref" yaml:"keySecretRef,omitempty"`
			} `tfsdk:"external_account_binding" yaml:"externalAccountBinding,omitempty"`

			PreferredChain *string `tfsdk:"preferred_chain" yaml:"preferredChain,omitempty"`

			PrivateKeySecretRef *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"private_key_secret_ref" yaml:"privateKeySecretRef,omitempty"`

			Server *string `tfsdk:"server" yaml:"server,omitempty"`

			SkipTLSVerify *bool `tfsdk:"skip_tls_verify" yaml:"skipTLSVerify,omitempty"`

			Solvers *[]struct {
				Dns01 *struct {
					AcmeDNS *struct {
						AccountSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"account_secret_ref" yaml:"accountSecretRef,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`
					} `tfsdk:"acme_dns" yaml:"acmeDNS,omitempty"`

					Akamai *struct {
						AccessTokenSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"access_token_secret_ref" yaml:"accessTokenSecretRef,omitempty"`

						ClientSecretSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"client_secret_secret_ref" yaml:"clientSecretSecretRef,omitempty"`

						ClientTokenSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"client_token_secret_ref" yaml:"clientTokenSecretRef,omitempty"`

						ServiceConsumerDomain *string `tfsdk:"service_consumer_domain" yaml:"serviceConsumerDomain,omitempty"`
					} `tfsdk:"akamai" yaml:"akamai,omitempty"`

					AzureDNS *struct {
						ClientID *string `tfsdk:"client_id" yaml:"clientID,omitempty"`

						ClientSecretSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"client_secret_secret_ref" yaml:"clientSecretSecretRef,omitempty"`

						Environment *string `tfsdk:"environment" yaml:"environment,omitempty"`

						HostedZoneName *string `tfsdk:"hosted_zone_name" yaml:"hostedZoneName,omitempty"`

						ManagedIdentity *struct {
							ClientID *string `tfsdk:"client_id" yaml:"clientID,omitempty"`

							ResourceID *string `tfsdk:"resource_id" yaml:"resourceID,omitempty"`
						} `tfsdk:"managed_identity" yaml:"managedIdentity,omitempty"`

						ResourceGroupName *string `tfsdk:"resource_group_name" yaml:"resourceGroupName,omitempty"`

						SubscriptionID *string `tfsdk:"subscription_id" yaml:"subscriptionID,omitempty"`

						TenantID *string `tfsdk:"tenant_id" yaml:"tenantID,omitempty"`
					} `tfsdk:"azure_dns" yaml:"azureDNS,omitempty"`

					CloudDNS *struct {
						HostedZoneName *string `tfsdk:"hosted_zone_name" yaml:"hostedZoneName,omitempty"`

						Project *string `tfsdk:"project" yaml:"project,omitempty"`

						ServiceAccountSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"service_account_secret_ref" yaml:"serviceAccountSecretRef,omitempty"`
					} `tfsdk:"cloud_dns" yaml:"cloudDNS,omitempty"`

					Cloudflare *struct {
						ApiKeySecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"api_key_secret_ref" yaml:"apiKeySecretRef,omitempty"`

						ApiTokenSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"api_token_secret_ref" yaml:"apiTokenSecretRef,omitempty"`

						Email *string `tfsdk:"email" yaml:"email,omitempty"`
					} `tfsdk:"cloudflare" yaml:"cloudflare,omitempty"`

					CnameStrategy *string `tfsdk:"cname_strategy" yaml:"cnameStrategy,omitempty"`

					Digitalocean *struct {
						TokenSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"token_secret_ref" yaml:"tokenSecretRef,omitempty"`
					} `tfsdk:"digitalocean" yaml:"digitalocean,omitempty"`

					Rfc2136 *struct {
						Nameserver *string `tfsdk:"nameserver" yaml:"nameserver,omitempty"`

						TsigAlgorithm *string `tfsdk:"tsig_algorithm" yaml:"tsigAlgorithm,omitempty"`

						TsigKeyName *string `tfsdk:"tsig_key_name" yaml:"tsigKeyName,omitempty"`

						TsigSecretSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"tsig_secret_secret_ref" yaml:"tsigSecretSecretRef,omitempty"`
					} `tfsdk:"rfc2136" yaml:"rfc2136,omitempty"`

					Route53 *struct {
						AccessKeyID *string `tfsdk:"access_key_id" yaml:"accessKeyID,omitempty"`

						AccessKeyIDSecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"access_key_id_secret_ref" yaml:"accessKeyIDSecretRef,omitempty"`

						HostedZoneID *string `tfsdk:"hosted_zone_id" yaml:"hostedZoneID,omitempty"`

						Region *string `tfsdk:"region" yaml:"region,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						SecretAccessKeySecretRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_access_key_secret_ref" yaml:"secretAccessKeySecretRef,omitempty"`
					} `tfsdk:"route53" yaml:"route53,omitempty"`

					Webhook *struct {
						Config utilities.Dynamic `tfsdk:"config" yaml:"config,omitempty"`

						GroupName *string `tfsdk:"group_name" yaml:"groupName,omitempty"`

						SolverName *string `tfsdk:"solver_name" yaml:"solverName,omitempty"`
					} `tfsdk:"webhook" yaml:"webhook,omitempty"`
				} `tfsdk:"dns01" yaml:"dns01,omitempty"`

				Http01 *struct {
					GatewayHTTPRoute *struct {
						Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

						ParentRefs *[]struct {
							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							SectionName *string `tfsdk:"section_name" yaml:"sectionName,omitempty"`
						} `tfsdk:"parent_refs" yaml:"parentRefs,omitempty"`

						ServiceType *string `tfsdk:"service_type" yaml:"serviceType,omitempty"`
					} `tfsdk:"gateway_http_route" yaml:"gatewayHTTPRoute,omitempty"`

					Ingress *struct {
						Class *string `tfsdk:"class" yaml:"class,omitempty"`

						IngressTemplate *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`
						} `tfsdk:"ingress_template" yaml:"ingressTemplate,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						PodTemplate *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

								Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
							} `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								Affinity *struct {
									NodeAffinity *struct {
										PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
											Preference *struct {
												MatchExpressions *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

													Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
												} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

												MatchFields *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

													Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
												} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
											} `tfsdk:"preference" yaml:"preference,omitempty"`

											Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
										} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

										RequiredDuringSchedulingIgnoredDuringExecution *struct {
											NodeSelectorTerms *[]struct {
												MatchExpressions *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

													Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
												} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

												MatchFields *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

													Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
												} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
											} `tfsdk:"node_selector_terms" yaml:"nodeSelectorTerms,omitempty"`
										} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
									} `tfsdk:"node_affinity" yaml:"nodeAffinity,omitempty"`

									PodAffinity *struct {
										PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
											PodAffinityTerm *struct {
												LabelSelector *struct {
													MatchExpressions *[]struct {
														Key *string `tfsdk:"key" yaml:"key,omitempty"`

														Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

														Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
													} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

													MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
												} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

												NamespaceSelector *struct {
													MatchExpressions *[]struct {
														Key *string `tfsdk:"key" yaml:"key,omitempty"`

														Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

														Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
													} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

													MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
												} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

												Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

												TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
											} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

											Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
										} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

										RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
											LabelSelector *struct {
												MatchExpressions *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

													Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
												} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

												MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
											} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

											NamespaceSelector *struct {
												MatchExpressions *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

													Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
												} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

												MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
											} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

											Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

											TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
										} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
									} `tfsdk:"pod_affinity" yaml:"podAffinity,omitempty"`

									PodAntiAffinity *struct {
										PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
											PodAffinityTerm *struct {
												LabelSelector *struct {
													MatchExpressions *[]struct {
														Key *string `tfsdk:"key" yaml:"key,omitempty"`

														Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

														Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
													} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

													MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
												} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

												NamespaceSelector *struct {
													MatchExpressions *[]struct {
														Key *string `tfsdk:"key" yaml:"key,omitempty"`

														Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

														Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
													} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

													MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
												} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

												Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

												TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
											} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

											Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
										} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

										RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
											LabelSelector *struct {
												MatchExpressions *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

													Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
												} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

												MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
											} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

											NamespaceSelector *struct {
												MatchExpressions *[]struct {
													Key *string `tfsdk:"key" yaml:"key,omitempty"`

													Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

													Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
												} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

												MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
											} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

											Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

											TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
										} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
									} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`
								} `tfsdk:"affinity" yaml:"affinity,omitempty"`

								NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

								PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

								ServiceAccountName *string `tfsdk:"service_account_name" yaml:"serviceAccountName,omitempty"`

								Tolerations *[]struct {
									Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"pod_template" yaml:"podTemplate,omitempty"`

						ServiceType *string `tfsdk:"service_type" yaml:"serviceType,omitempty"`
					} `tfsdk:"ingress" yaml:"ingress,omitempty"`
				} `tfsdk:"http01" yaml:"http01,omitempty"`

				Selector *struct {
					DnsNames *[]string `tfsdk:"dns_names" yaml:"dnsNames,omitempty"`

					DnsZones *[]string `tfsdk:"dns_zones" yaml:"dnsZones,omitempty"`

					MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`
			} `tfsdk:"solvers" yaml:"solvers,omitempty"`
		} `tfsdk:"acme" yaml:"acme,omitempty"`

		Ca *struct {
			CrlDistributionPoints *[]string `tfsdk:"crl_distribution_points" yaml:"crlDistributionPoints,omitempty"`

			OcspServers *[]string `tfsdk:"ocsp_servers" yaml:"ocspServers,omitempty"`

			SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
		} `tfsdk:"ca" yaml:"ca,omitempty"`

		SelfSigned *struct {
			CrlDistributionPoints *[]string `tfsdk:"crl_distribution_points" yaml:"crlDistributionPoints,omitempty"`
		} `tfsdk:"self_signed" yaml:"selfSigned,omitempty"`

		Vault *struct {
			Auth *struct {
				AppRole *struct {
					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					RoleId *string `tfsdk:"role_id" yaml:"roleId,omitempty"`

					SecretRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
				} `tfsdk:"app_role" yaml:"appRole,omitempty"`

				Kubernetes *struct {
					MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

					Role *string `tfsdk:"role" yaml:"role,omitempty"`

					SecretRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
				} `tfsdk:"kubernetes" yaml:"kubernetes,omitempty"`

				TokenSecretRef *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"token_secret_ref" yaml:"tokenSecretRef,omitempty"`
			} `tfsdk:"auth" yaml:"auth,omitempty"`

			CaBundle *string `tfsdk:"ca_bundle" yaml:"caBundle,omitempty"`

			CaBundleSecretRef *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"ca_bundle_secret_ref" yaml:"caBundleSecretRef,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Server *string `tfsdk:"server" yaml:"server,omitempty"`
		} `tfsdk:"vault" yaml:"vault,omitempty"`

		Venafi *struct {
			Cloud *struct {
				ApiTokenSecretRef *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"api_token_secret_ref" yaml:"apiTokenSecretRef,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"cloud" yaml:"cloud,omitempty"`

			Tpp *struct {
				CaBundle *string `tfsdk:"ca_bundle" yaml:"caBundle,omitempty"`

				CredentialsRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"credentials_ref" yaml:"credentialsRef,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"tpp" yaml:"tpp,omitempty"`

			Zone *string `tfsdk:"zone" yaml:"zone,omitempty"`
		} `tfsdk:"venafi" yaml:"venafi,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCertManagerIoClusterIssuerV1Resource() resource.Resource {
	return &CertManagerIoClusterIssuerV1Resource{}
}

func (r *CertManagerIoClusterIssuerV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cert_manager_io_cluster_issuer_v1"
}

func (r *CertManagerIoClusterIssuerV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "A ClusterIssuer represents a certificate issuing authority which can be referenced as part of 'issuerRef' fields. It is similar to an Issuer, however it is cluster-scoped and therefore can be referenced by resources that exist in *any* namespace, not just the same namespace as the referent.",
		MarkdownDescription: "A ClusterIssuer represents a certificate issuing authority which can be referenced as part of 'issuerRef' fields. It is similar to an Issuer, however it is cluster-scoped and therefore can be referenced by resources that exist in *any* namespace, not just the same namespace as the referent.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "Desired state of the ClusterIssuer resource.",
				MarkdownDescription: "Desired state of the ClusterIssuer resource.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"acme": {
						Description:         "ACME configures this issuer to communicate with a RFC8555 (ACME) server to obtain signed x509 certificates.",
						MarkdownDescription: "ACME configures this issuer to communicate with a RFC8555 (ACME) server to obtain signed x509 certificates.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"disable_account_key_generation": {
								Description:         "Enables or disables generating a new ACME account key. If true, the Issuer resource will *not* request a new account but will expect the account key to be supplied via an existing secret. If false, the cert-manager system will generate a new ACME account key for the Issuer. Defaults to false.",
								MarkdownDescription: "Enables or disables generating a new ACME account key. If true, the Issuer resource will *not* request a new account but will expect the account key to be supplied via an existing secret. If false, the cert-manager system will generate a new ACME account key for the Issuer. Defaults to false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"email": {
								Description:         "Email is the email address to be associated with the ACME account. This field is optional, but it is strongly recommended to be set. It will be used to contact you in case of issues with your account or certificates, including expiry notification emails. This field may be updated after the account is initially registered.",
								MarkdownDescription: "Email is the email address to be associated with the ACME account. This field is optional, but it is strongly recommended to be set. It will be used to contact you in case of issues with your account or certificates, including expiry notification emails. This field may be updated after the account is initially registered.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_duration_feature": {
								Description:         "Enables requesting a Not After date on certificates that matches the duration of the certificate. This is not supported by all ACME servers like Let's Encrypt. If set to true when the ACME server does not support it it will create an error on the Order. Defaults to false.",
								MarkdownDescription: "Enables requesting a Not After date on certificates that matches the duration of the certificate. This is not supported by all ACME servers like Let's Encrypt. If set to true when the ACME server does not support it it will create an error on the Order. Defaults to false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"external_account_binding": {
								Description:         "ExternalAccountBinding is a reference to a CA external account of the ACME server. If set, upon registration cert-manager will attempt to associate the given external account credentials with the registered ACME account.",
								MarkdownDescription: "ExternalAccountBinding is a reference to a CA external account of the ACME server. If set, upon registration cert-manager will attempt to associate the given external account credentials with the registered ACME account.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key_algorithm": {
										Description:         "Deprecated: keyAlgorithm field exists for historical compatibility reasons and should not be used. The algorithm is now hardcoded to HS256 in golang/x/crypto/acme.",
										MarkdownDescription: "Deprecated: keyAlgorithm field exists for historical compatibility reasons and should not be used. The algorithm is now hardcoded to HS256 in golang/x/crypto/acme.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("HS256", "HS384", "HS512"),
										},
									},

									"key_id": {
										Description:         "keyID is the ID of the CA key that the External Account is bound to.",
										MarkdownDescription: "keyID is the ID of the CA key that the External Account is bound to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"key_secret_ref": {
										Description:         "keySecretRef is a Secret Key Selector referencing a data item in a Kubernetes Secret which holds the symmetric MAC key of the External Account Binding. The 'key' is the index string that is paired with the key data in the Secret and should not be confused with the key data itself, or indeed with the External Account Binding keyID above. The secret key stored in the Secret **must** be un-padded, base64 URL encoded data.",
										MarkdownDescription: "keySecretRef is a Secret Key Selector referencing a data item in a Kubernetes Secret which holds the symmetric MAC key of the External Account Binding. The 'key' is the index string that is paired with the key data in the Secret and should not be confused with the key data itself, or indeed with the External Account Binding keyID above. The secret key stored in the Secret **must** be un-padded, base64 URL encoded data.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"preferred_chain": {
								Description:         "PreferredChain is the chain to use if the ACME server outputs multiple. PreferredChain is no guarantee that this one gets delivered by the ACME endpoint. For example, for Let's Encrypt's DST crosssign you would use: 'DST Root CA X3' or 'ISRG Root X1' for the newer Let's Encrypt root CA. This value picks the first certificate bundle in the ACME alternative chains that has a certificate with this value as its issuer's CN",
								MarkdownDescription: "PreferredChain is the chain to use if the ACME server outputs multiple. PreferredChain is no guarantee that this one gets delivered by the ACME endpoint. For example, for Let's Encrypt's DST crosssign you would use: 'DST Root CA X3' or 'ISRG Root X1' for the newer Let's Encrypt root CA. This value picks the first certificate bundle in the ACME alternative chains that has a certificate with this value as its issuer's CN",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtMost(64),
								},
							},

							"private_key_secret_ref": {
								Description:         "PrivateKey is the name of a Kubernetes Secret resource that will be used to store the automatically generated ACME account private key. Optionally, a 'key' may be specified to select a specific entry within the named Secret resource. If 'key' is not specified, a default of 'tls.key' will be used.",
								MarkdownDescription: "PrivateKey is the name of a Kubernetes Secret resource that will be used to store the automatically generated ACME account private key. Optionally, a 'key' may be specified to select a specific entry within the named Secret resource. If 'key' is not specified, a default of 'tls.key' will be used.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
										MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"server": {
								Description:         "Server is the URL used to access the ACME server's 'directory' endpoint. For example, for Let's Encrypt's staging endpoint, you would use: 'https://acme-staging-v02.api.letsencrypt.org/directory'. Only ACME v2 endpoints (i.e. RFC 8555) are supported.",
								MarkdownDescription: "Server is the URL used to access the ACME server's 'directory' endpoint. For example, for Let's Encrypt's staging endpoint, you would use: 'https://acme-staging-v02.api.letsencrypt.org/directory'. Only ACME v2 endpoints (i.e. RFC 8555) are supported.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"skip_tls_verify": {
								Description:         "Enables or disables validation of the ACME server TLS certificate. If true, requests to the ACME server will not have their TLS certificate validated (i.e. insecure connections will be allowed). Only enable this option in development environments. The cert-manager system installed roots will be used to verify connections to the ACME server if this is false. Defaults to false.",
								MarkdownDescription: "Enables or disables validation of the ACME server TLS certificate. If true, requests to the ACME server will not have their TLS certificate validated (i.e. insecure connections will be allowed). Only enable this option in development environments. The cert-manager system installed roots will be used to verify connections to the ACME server if this is false. Defaults to false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"solvers": {
								Description:         "Solvers is a list of challenge solvers that will be used to solve ACME challenges for the matching domains. Solver configurations must be provided in order to obtain certificates from an ACME server. For more information, see: https://cert-manager.io/docs/configuration/acme/",
								MarkdownDescription: "Solvers is a list of challenge solvers that will be used to solve ACME challenges for the matching domains. Solver configurations must be provided in order to obtain certificates from an ACME server. For more information, see: https://cert-manager.io/docs/configuration/acme/",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"dns01": {
										Description:         "Configures cert-manager to attempt to complete authorizations by performing the DNS01 challenge flow.",
										MarkdownDescription: "Configures cert-manager to attempt to complete authorizations by performing the DNS01 challenge flow.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"acme_dns": {
												Description:         "Use the 'ACME DNS' (https://github.com/joohoi/acme-dns) API to manage DNS01 challenge records.",
												MarkdownDescription: "Use the 'ACME DNS' (https://github.com/joohoi/acme-dns) API to manage DNS01 challenge records.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"account_secret_ref": {
														Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"host": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"akamai": {
												Description:         "Use the Akamai DNS zone management API to manage DNS01 challenge records.",
												MarkdownDescription: "Use the Akamai DNS zone management API to manage DNS01 challenge records.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"access_token_secret_ref": {
														Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"client_secret_secret_ref": {
														Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"client_token_secret_ref": {
														Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"service_consumer_domain": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"azure_dns": {
												Description:         "Use the Microsoft Azure DNS API to manage DNS01 challenge records.",
												MarkdownDescription: "Use the Microsoft Azure DNS API to manage DNS01 challenge records.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_id": {
														Description:         "if both this and ClientSecret are left unset MSI will be used",
														MarkdownDescription: "if both this and ClientSecret are left unset MSI will be used",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"client_secret_secret_ref": {
														Description:         "if both this and ClientID are left unset MSI will be used",
														MarkdownDescription: "if both this and ClientID are left unset MSI will be used",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"environment": {
														Description:         "name of the Azure environment (default AzurePublicCloud)",
														MarkdownDescription: "name of the Azure environment (default AzurePublicCloud)",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("AzurePublicCloud", "AzureChinaCloud", "AzureGermanCloud", "AzureUSGovernmentCloud"),
														},
													},

													"hosted_zone_name": {
														Description:         "name of the DNS zone that should be used",
														MarkdownDescription: "name of the DNS zone that should be used",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"managed_identity": {
														Description:         "managed identity configuration, can not be used at the same time as clientID, clientSecretSecretRef or tenantID",
														MarkdownDescription: "managed identity configuration, can not be used at the same time as clientID, clientSecretSecretRef or tenantID",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"client_id": {
																Description:         "client ID of the managed identity, can not be used at the same time as resourceID",
																MarkdownDescription: "client ID of the managed identity, can not be used at the same time as resourceID",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"resource_id": {
																Description:         "resource ID of the managed identity, can not be used at the same time as clientID",
																MarkdownDescription: "resource ID of the managed identity, can not be used at the same time as clientID",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource_group_name": {
														Description:         "resource group the DNS zone is located in",
														MarkdownDescription: "resource group the DNS zone is located in",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"subscription_id": {
														Description:         "ID of the Azure subscription",
														MarkdownDescription: "ID of the Azure subscription",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"tenant_id": {
														Description:         "when specifying ClientID and ClientSecret then this field is also needed",
														MarkdownDescription: "when specifying ClientID and ClientSecret then this field is also needed",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cloud_dns": {
												Description:         "Use the Google Cloud DNS API to manage DNS01 challenge records.",
												MarkdownDescription: "Use the Google Cloud DNS API to manage DNS01 challenge records.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"hosted_zone_name": {
														Description:         "HostedZoneName is an optional field that tells cert-manager in which Cloud DNS zone the challenge record has to be created. If left empty cert-manager will automatically choose a zone.",
														MarkdownDescription: "HostedZoneName is an optional field that tells cert-manager in which Cloud DNS zone the challenge record has to be created. If left empty cert-manager will automatically choose a zone.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"project": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"service_account_secret_ref": {
														Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cloudflare": {
												Description:         "Use the Cloudflare API to manage DNS01 challenge records.",
												MarkdownDescription: "Use the Cloudflare API to manage DNS01 challenge records.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_key_secret_ref": {
														Description:         "API key to use to authenticate with Cloudflare. Note: using an API token to authenticate is now the recommended method as it allows greater control of permissions.",
														MarkdownDescription: "API key to use to authenticate with Cloudflare. Note: using an API token to authenticate is now the recommended method as it allows greater control of permissions.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"api_token_secret_ref": {
														Description:         "API token used to authenticate with Cloudflare.",
														MarkdownDescription: "API token used to authenticate with Cloudflare.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"email": {
														Description:         "Email of the account, only required when using API key based authentication.",
														MarkdownDescription: "Email of the account, only required when using API key based authentication.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cname_strategy": {
												Description:         "CNAMEStrategy configures how the DNS01 provider should handle CNAME records when found in DNS zones.",
												MarkdownDescription: "CNAMEStrategy configures how the DNS01 provider should handle CNAME records when found in DNS zones.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("None", "Follow"),
												},
											},

											"digitalocean": {
												Description:         "Use the DigitalOcean DNS API to manage DNS01 challenge records.",
												MarkdownDescription: "Use the DigitalOcean DNS API to manage DNS01 challenge records.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"token_secret_ref": {
														Description:         "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",
														MarkdownDescription: "A reference to a specific 'key' within a Secret resource. In some instances, 'key' is a required field.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rfc2136": {
												Description:         "Use RFC2136 ('Dynamic Updates in the Domain Name System') (https://datatracker.ietf.org/doc/rfc2136/) to manage DNS01 challenge records.",
												MarkdownDescription: "Use RFC2136 ('Dynamic Updates in the Domain Name System') (https://datatracker.ietf.org/doc/rfc2136/) to manage DNS01 challenge records.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"nameserver": {
														Description:         "The IP address or hostname of an authoritative DNS server supporting RFC2136 in the form host:port. If the host is an IPv6 address it must be enclosed in square brackets (e.g [2001:db8::1]) ; port is optional. This field is required.",
														MarkdownDescription: "The IP address or hostname of an authoritative DNS server supporting RFC2136 in the form host:port. If the host is an IPv6 address it must be enclosed in square brackets (e.g [2001:db8::1]) ; port is optional. This field is required.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"tsig_algorithm": {
														Description:         "The TSIG Algorithm configured in the DNS supporting RFC2136. Used only when ''tsigSecretSecretRef'' and ''tsigKeyName'' are defined. Supported values are (case-insensitive): ''HMACMD5'' (default), ''HMACSHA1'', ''HMACSHA256'' or ''HMACSHA512''.",
														MarkdownDescription: "The TSIG Algorithm configured in the DNS supporting RFC2136. Used only when ''tsigSecretSecretRef'' and ''tsigKeyName'' are defined. Supported values are (case-insensitive): ''HMACMD5'' (default), ''HMACSHA1'', ''HMACSHA256'' or ''HMACSHA512''.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tsig_key_name": {
														Description:         "The TSIG Key name configured in the DNS. If ''tsigSecretSecretRef'' is defined, this field is required.",
														MarkdownDescription: "The TSIG Key name configured in the DNS. If ''tsigSecretSecretRef'' is defined, this field is required.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tsig_secret_secret_ref": {
														Description:         "The name of the secret containing the TSIG value. If ''tsigKeyName'' is defined, this field is required.",
														MarkdownDescription: "The name of the secret containing the TSIG value. If ''tsigKeyName'' is defined, this field is required.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"route53": {
												Description:         "Use the AWS Route53 API to manage DNS01 challenge records.",
												MarkdownDescription: "Use the AWS Route53 API to manage DNS01 challenge records.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"access_key_id": {
														Description:         "The AccessKeyID is used for authentication. Cannot be set when SecretAccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
														MarkdownDescription: "The AccessKeyID is used for authentication. Cannot be set when SecretAccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"access_key_id_secret_ref": {
														Description:         "The SecretAccessKey is used for authentication. If set, pull the AWS access key ID from a key within a Kubernetes Secret. Cannot be set when AccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
														MarkdownDescription: "The SecretAccessKey is used for authentication. If set, pull the AWS access key ID from a key within a Kubernetes Secret. Cannot be set when AccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"hosted_zone_id": {
														Description:         "If set, the provider will manage only this zone in Route53 and will not do an lookup using the route53:ListHostedZonesByName api call.",
														MarkdownDescription: "If set, the provider will manage only this zone in Route53 and will not do an lookup using the route53:ListHostedZonesByName api call.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"region": {
														Description:         "Always set the region when using AccessKeyID and SecretAccessKey",
														MarkdownDescription: "Always set the region when using AccessKeyID and SecretAccessKey",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"role": {
														Description:         "Role is a Role ARN which the Route53 provider will assume using either the explicit credentials AccessKeyID/SecretAccessKey or the inferred credentials from environment variables, shared credentials file or AWS Instance metadata",
														MarkdownDescription: "Role is a Role ARN which the Route53 provider will assume using either the explicit credentials AccessKeyID/SecretAccessKey or the inferred credentials from environment variables, shared credentials file or AWS Instance metadata",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_access_key_secret_ref": {
														Description:         "The SecretAccessKey is used for authentication. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
														MarkdownDescription: "The SecretAccessKey is used for authentication. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"webhook": {
												Description:         "Configure an external webhook based DNS01 challenge solver to manage DNS01 challenge records.",
												MarkdownDescription: "Configure an external webhook based DNS01 challenge solver to manage DNS01 challenge records.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"config": {
														Description:         "Additional configuration that should be passed to the webhook apiserver when challenges are processed. This can contain arbitrary JSON data. Secret values should not be specified in this stanza. If secret values are needed (e.g. credentials for a DNS service), you should use a SecretKeySelector to reference a Secret resource. For details on the schema of this field, consult the webhook provider implementation's documentation.",
														MarkdownDescription: "Additional configuration that should be passed to the webhook apiserver when challenges are processed. This can contain arbitrary JSON data. Secret values should not be specified in this stanza. If secret values are needed (e.g. credentials for a DNS service), you should use a SecretKeySelector to reference a Secret resource. For details on the schema of this field, consult the webhook provider implementation's documentation.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"group_name": {
														Description:         "The API group name that should be used when POSTing ChallengePayload resources to the webhook apiserver. This should be the same as the GroupName specified in the webhook provider implementation.",
														MarkdownDescription: "The API group name that should be used when POSTing ChallengePayload resources to the webhook apiserver. This should be the same as the GroupName specified in the webhook provider implementation.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"solver_name": {
														Description:         "The name of the solver to use, as defined in the webhook provider implementation. This will typically be the name of the provider, e.g. 'cloudflare'.",
														MarkdownDescription: "The name of the solver to use, as defined in the webhook provider implementation. This will typically be the name of the provider, e.g. 'cloudflare'.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http01": {
										Description:         "Configures cert-manager to attempt to complete authorizations by performing the HTTP01 challenge flow. It is not possible to obtain certificates for wildcard domain names (e.g. '*.example.com') using the HTTP01 challenge mechanism.",
										MarkdownDescription: "Configures cert-manager to attempt to complete authorizations by performing the HTTP01 challenge flow. It is not possible to obtain certificates for wildcard domain names (e.g. '*.example.com') using the HTTP01 challenge mechanism.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"gateway_http_route": {
												Description:         "The Gateway API is a sig-network community API that models service networking in Kubernetes (https://gateway-api.sigs.k8s.io/). The Gateway solver will create HTTPRoutes with the specified labels in the same namespace as the challenge. This solver is experimental, and fields / behaviour may change in the future.",
												MarkdownDescription: "The Gateway API is a sig-network community API that models service networking in Kubernetes (https://gateway-api.sigs.k8s.io/). The Gateway solver will create HTTPRoutes with the specified labels in the same namespace as the challenge. This solver is experimental, and fields / behaviour may change in the future.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"labels": {
														Description:         "Custom labels that will be applied to HTTPRoutes created by cert-manager while solving HTTP-01 challenges.",
														MarkdownDescription: "Custom labels that will be applied to HTTPRoutes created by cert-manager while solving HTTP-01 challenges.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"parent_refs": {
														Description:         "When solving an HTTP-01 challenge, cert-manager creates an HTTPRoute. cert-manager needs to know which parentRefs should be used when creating the HTTPRoute. Usually, the parentRef references a Gateway. See: https://gateway-api.sigs.k8s.io/v1alpha2/api-types/httproute/#attaching-to-gateways",
														MarkdownDescription: "When solving an HTTP-01 challenge, cert-manager creates an HTTPRoute. cert-manager needs to know which parentRefs should be used when creating the HTTPRoute. Usually, the parentRef references a Gateway. See: https://gateway-api.sigs.k8s.io/v1alpha2/api-types/httproute/#attaching-to-gateways",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"group": {
																Description:         "Group is the group of the referent.  Support: Core",
																MarkdownDescription: "Group is the group of the referent.  Support: Core",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.LengthAtMost(253),

																	stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																},
															},

															"kind": {
																Description:         "Kind is kind of the referent.  Support: Core (Gateway)  Support: Custom (Other Resources)",
																MarkdownDescription: "Kind is kind of the referent.  Support: Core (Gateway)  Support: Custom (Other Resources)",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.LengthAtLeast(1),

																	stringvalidator.LengthAtMost(63),

																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
																},
															},

															"name": {
																Description:         "Name is the name of the referent.  Support: Core",
																MarkdownDescription: "Name is the name of the referent.  Support: Core",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.LengthAtLeast(1),

																	stringvalidator.LengthAtMost(253),
																},
															},

															"namespace": {
																Description:         "Namespace is the namespace of the referent. When unspecified (or empty string), this refers to the local namespace of the Route.  Support: Core",
																MarkdownDescription: "Namespace is the namespace of the referent. When unspecified (or empty string), this refers to the local namespace of the Route.  Support: Core",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.LengthAtLeast(1),

																	stringvalidator.LengthAtMost(63),

																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"port": {
																Description:         "Port is the network port this Route targets. It can be interpreted differently based on the type of parent resource.  When the parent resource is a Gateway, this targets all listeners listening on the specified port that also support this kind of Route(and select this Route). It's not recommended to set 'Port' unless the networking behaviors specified in a Route must apply to a specific port as opposed to a listener(s) whose port(s) may be changed. When both Port and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support other parent resources. Implementations supporting other types of parent resources MUST clearly document how/if Port is interpreted.  For the purpose of status, an attachment is considered successful as long as the parent resource accepts it partially. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Extended  <gateway:experimental>",
																MarkdownDescription: "Port is the network port this Route targets. It can be interpreted differently based on the type of parent resource.  When the parent resource is a Gateway, this targets all listeners listening on the specified port that also support this kind of Route(and select this Route). It's not recommended to set 'Port' unless the networking behaviors specified in a Route must apply to a specific port as opposed to a listener(s) whose port(s) may be changed. When both Port and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support other parent resources. Implementations supporting other types of parent resources MUST clearly document how/if Port is interpreted.  For the purpose of status, an attachment is considered successful as long as the parent resource accepts it partially. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Extended  <gateway:experimental>",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	int64validator.AtLeast(1),

																	int64validator.AtMost(65535),
																},
															},

															"section_name": {
																Description:         "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",
																MarkdownDescription: "SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values.  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.LengthAtLeast(1),

																	stringvalidator.LengthAtMost(253),

																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																},
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_type": {
														Description:         "Optional service type for Kubernetes solver service. Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.",
														MarkdownDescription: "Optional service type for Kubernetes solver service. Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ingress": {
												Description:         "The ingress based HTTP01 challenge solver will solve challenges by creating or modifying Ingress resources in order to route requests for '/.well-known/acme-challenge/XYZ' to 'challenge solver' pods that are provisioned by cert-manager for each Challenge to be completed.",
												MarkdownDescription: "The ingress based HTTP01 challenge solver will solve challenges by creating or modifying Ingress resources in order to route requests for '/.well-known/acme-challenge/XYZ' to 'challenge solver' pods that are provisioned by cert-manager for each Challenge to be completed.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"class": {
														Description:         "The ingress class to use when creating Ingress resources to solve ACME challenges that use this challenge solver. Only one of 'class' or 'name' may be specified.",
														MarkdownDescription: "The ingress class to use when creating Ingress resources to solve ACME challenges that use this challenge solver. Only one of 'class' or 'name' may be specified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ingress_template": {
														Description:         "Optional ingress template used to configure the ACME challenge solver ingress used for HTTP01 challenges.",
														MarkdownDescription: "Optional ingress template used to configure the ACME challenge solver ingress used for HTTP01 challenges.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "ObjectMeta overrides for the ingress used to solve HTTP01 challenges. Only the 'labels' and 'annotations' fields may be set. If labels or annotations overlap with in-built values, the values here will override the in-built values.",
																MarkdownDescription: "ObjectMeta overrides for the ingress used to solve HTTP01 challenges. Only the 'labels' and 'annotations' fields may be set. If labels or annotations overlap with in-built values, the values here will override the in-built values.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "Annotations that should be added to the created ACME HTTP01 solver ingress.",
																		MarkdownDescription: "Annotations that should be added to the created ACME HTTP01 solver ingress.",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "Labels that should be added to the created ACME HTTP01 solver ingress.",
																		MarkdownDescription: "Labels that should be added to the created ACME HTTP01 solver ingress.",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "The name of the ingress resource that should have ACME challenge solving routes inserted into it in order to solve HTTP01 challenges. This is typically used in conjunction with ingress controllers like ingress-gce, which maintains a 1:1 mapping between external IPs and ingress resources.",
														MarkdownDescription: "The name of the ingress resource that should have ACME challenge solving routes inserted into it in order to solve HTTP01 challenges. This is typically used in conjunction with ingress controllers like ingress-gce, which maintains a 1:1 mapping between external IPs and ingress resources.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_template": {
														Description:         "Optional pod template used to configure the ACME challenge solver pods used for HTTP01 challenges.",
														MarkdownDescription: "Optional pod template used to configure the ACME challenge solver pods used for HTTP01 challenges.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "ObjectMeta overrides for the pod used to solve HTTP01 challenges. Only the 'labels' and 'annotations' fields may be set. If labels or annotations overlap with in-built values, the values here will override the in-built values.",
																MarkdownDescription: "ObjectMeta overrides for the pod used to solve HTTP01 challenges. Only the 'labels' and 'annotations' fields may be set. If labels or annotations overlap with in-built values, the values here will override the in-built values.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"annotations": {
																		Description:         "Annotations that should be added to the create ACME HTTP01 solver pods.",
																		MarkdownDescription: "Annotations that should be added to the create ACME HTTP01 solver pods.",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"labels": {
																		Description:         "Labels that should be added to the created ACME HTTP01 solver pods.",
																		MarkdownDescription: "Labels that should be added to the created ACME HTTP01 solver pods.",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"spec": {
																Description:         "PodSpec defines overrides for the HTTP01 challenge solver pod. Only the 'priorityClassName', 'nodeSelector', 'affinity', 'serviceAccountName' and 'tolerations' fields are supported currently. All other fields will be ignored.",
																MarkdownDescription: "PodSpec defines overrides for the HTTP01 challenge solver pod. Only the 'priorityClassName', 'nodeSelector', 'affinity', 'serviceAccountName' and 'tolerations' fields are supported currently. All other fields will be ignored.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"affinity": {
																		Description:         "If specified, the pod's scheduling constraints",
																		MarkdownDescription: "If specified, the pod's scheduling constraints",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"node_affinity": {
																				Description:         "Describes node affinity scheduling rules for the pod.",
																				MarkdownDescription: "Describes node affinity scheduling rules for the pod.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"preferred_during_scheduling_ignored_during_execution": {
																						Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
																						MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"preference": {
																								Description:         "A node selector term, associated with the corresponding weight.",
																								MarkdownDescription: "A node selector term, associated with the corresponding weight.",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"match_expressions": {
																										Description:         "A list of node selector requirements by node's labels.",
																										MarkdownDescription: "A list of node selector requirements by node's labels.",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "The label key that the selector applies to.",
																												MarkdownDescription: "The label key that the selector applies to.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"operator": {
																												Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																												MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"values": {
																												Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																												Type: types.ListType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},
																										}),

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"match_fields": {
																										Description:         "A list of node selector requirements by node's fields.",
																										MarkdownDescription: "A list of node selector requirements by node's fields.",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "The label key that the selector applies to.",
																												MarkdownDescription: "The label key that the selector applies to.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"operator": {
																												Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																												MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"values": {
																												Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																												Type: types.ListType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},
																										}),

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},
																								}),

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"weight": {
																								Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																								MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",

																								Type: types.Int64Type,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},
																						}),

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"required_during_scheduling_ignored_during_execution": {
																						Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
																						MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"node_selector_terms": {
																								Description:         "Required. A list of node selector terms. The terms are ORed.",
																								MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"match_expressions": {
																										Description:         "A list of node selector requirements by node's labels.",
																										MarkdownDescription: "A list of node selector requirements by node's labels.",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "The label key that the selector applies to.",
																												MarkdownDescription: "The label key that the selector applies to.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"operator": {
																												Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																												MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"values": {
																												Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																												Type: types.ListType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},
																										}),

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"match_fields": {
																										Description:         "A list of node selector requirements by node's fields.",
																										MarkdownDescription: "A list of node selector requirements by node's fields.",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "The label key that the selector applies to.",
																												MarkdownDescription: "The label key that the selector applies to.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"operator": {
																												Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																												MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"values": {
																												Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																												Type: types.ListType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},
																										}),

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},
																								}),

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},
																						}),

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"pod_affinity": {
																				Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
																				MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"preferred_during_scheduling_ignored_during_execution": {
																						Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																						MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"pod_affinity_term": {
																								Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																								MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"label_selector": {
																										Description:         "A label query over a set of resources, in this case pods.",
																										MarkdownDescription: "A label query over a set of resources, in this case pods.",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"match_expressions": {
																												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																													"key": {
																														Description:         "key is the label key that the selector applies to.",
																														MarkdownDescription: "key is the label key that the selector applies to.",

																														Type: types.StringType,

																														Required: true,
																														Optional: false,
																														Computed: false,
																													},

																													"operator": {
																														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																														Type: types.StringType,

																														Required: true,
																														Optional: false,
																														Computed: false,
																													},

																													"values": {
																														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																														Type: types.ListType{ElemType: types.StringType},

																														Required: false,
																														Optional: true,
																														Computed: false,
																													},
																												}),

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"match_labels": {
																												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																												Type: types.MapType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},
																										}),

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"namespace_selector": {
																										Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																										MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"match_expressions": {
																												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																													"key": {
																														Description:         "key is the label key that the selector applies to.",
																														MarkdownDescription: "key is the label key that the selector applies to.",

																														Type: types.StringType,

																														Required: true,
																														Optional: false,
																														Computed: false,
																													},

																													"operator": {
																														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																														Type: types.StringType,

																														Required: true,
																														Optional: false,
																														Computed: false,
																													},

																													"values": {
																														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																														Type: types.ListType{ElemType: types.StringType},

																														Required: false,
																														Optional: true,
																														Computed: false,
																													},
																												}),

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"match_labels": {
																												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																												Type: types.MapType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},
																										}),

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"namespaces": {
																										Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																										MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",

																										Type: types.ListType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"topology_key": {
																										Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																										MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},
																								}),

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"weight": {
																								Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																								MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",

																								Type: types.Int64Type,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},
																						}),

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"required_during_scheduling_ignored_during_execution": {
																						Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																						MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"label_selector": {
																								Description:         "A label query over a set of resources, in this case pods.",
																								MarkdownDescription: "A label query over a set of resources, in this case pods.",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"match_expressions": {
																										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "key is the label key that the selector applies to.",
																												MarkdownDescription: "key is the label key that the selector applies to.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"operator": {
																												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"values": {
																												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																												Type: types.ListType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},
																										}),

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"match_labels": {
																										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																										Type: types.MapType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},
																								}),

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"namespace_selector": {
																								Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																								MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"match_expressions": {
																										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "key is the label key that the selector applies to.",
																												MarkdownDescription: "key is the label key that the selector applies to.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"operator": {
																												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"values": {
																												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																												Type: types.ListType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},
																										}),

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"match_labels": {
																										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																										Type: types.MapType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},
																								}),

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"namespaces": {
																								Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																								MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"topology_key": {
																								Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																								MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},
																						}),

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"pod_anti_affinity": {
																				Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
																				MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"preferred_during_scheduling_ignored_during_execution": {
																						Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																						MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"pod_affinity_term": {
																								Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																								MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"label_selector": {
																										Description:         "A label query over a set of resources, in this case pods.",
																										MarkdownDescription: "A label query over a set of resources, in this case pods.",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"match_expressions": {
																												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																													"key": {
																														Description:         "key is the label key that the selector applies to.",
																														MarkdownDescription: "key is the label key that the selector applies to.",

																														Type: types.StringType,

																														Required: true,
																														Optional: false,
																														Computed: false,
																													},

																													"operator": {
																														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																														Type: types.StringType,

																														Required: true,
																														Optional: false,
																														Computed: false,
																													},

																													"values": {
																														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																														Type: types.ListType{ElemType: types.StringType},

																														Required: false,
																														Optional: true,
																														Computed: false,
																													},
																												}),

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"match_labels": {
																												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																												Type: types.MapType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},
																										}),

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"namespace_selector": {
																										Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																										MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",

																										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																											"match_expressions": {
																												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																													"key": {
																														Description:         "key is the label key that the selector applies to.",
																														MarkdownDescription: "key is the label key that the selector applies to.",

																														Type: types.StringType,

																														Required: true,
																														Optional: false,
																														Computed: false,
																													},

																													"operator": {
																														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																														Type: types.StringType,

																														Required: true,
																														Optional: false,
																														Computed: false,
																													},

																													"values": {
																														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																														Type: types.ListType{ElemType: types.StringType},

																														Required: false,
																														Optional: true,
																														Computed: false,
																													},
																												}),

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},

																											"match_labels": {
																												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																												Type: types.MapType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},
																										}),

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"namespaces": {
																										Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																										MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",

																										Type: types.ListType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"topology_key": {
																										Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																										MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},
																								}),

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"weight": {
																								Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																								MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",

																								Type: types.Int64Type,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},
																						}),

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"required_during_scheduling_ignored_during_execution": {
																						Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																						MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"label_selector": {
																								Description:         "A label query over a set of resources, in this case pods.",
																								MarkdownDescription: "A label query over a set of resources, in this case pods.",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"match_expressions": {
																										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "key is the label key that the selector applies to.",
																												MarkdownDescription: "key is the label key that the selector applies to.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"operator": {
																												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"values": {
																												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																												Type: types.ListType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},
																										}),

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"match_labels": {
																										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																										Type: types.MapType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},
																								}),

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"namespace_selector": {
																								Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																								MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",

																								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																									"match_expressions": {
																										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																											"key": {
																												Description:         "key is the label key that the selector applies to.",
																												MarkdownDescription: "key is the label key that the selector applies to.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"operator": {
																												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																												Type: types.StringType,

																												Required: true,
																												Optional: false,
																												Computed: false,
																											},

																											"values": {
																												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																												Type: types.ListType{ElemType: types.StringType},

																												Required: false,
																												Optional: true,
																												Computed: false,
																											},
																										}),

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"match_labels": {
																										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

																										Type: types.MapType{ElemType: types.StringType},

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},
																								}),

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"namespaces": {
																								Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																								MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"topology_key": {
																								Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																								MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},
																						}),

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"node_selector": {
																		Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
																		MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"priority_class_name": {
																		Description:         "If specified, the pod's priorityClassName.",
																		MarkdownDescription: "If specified, the pod's priorityClassName.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"service_account_name": {
																		Description:         "If specified, the pod's service account",
																		MarkdownDescription: "If specified, the pod's service account",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"tolerations": {
																		Description:         "If specified, the pod's tolerations.",
																		MarkdownDescription: "If specified, the pod's tolerations.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"effect": {
																				Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																				MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"key": {
																				Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																				MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"operator": {
																				Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
																				MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"toleration_seconds": {
																				Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
																				MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
																				MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_type": {
														Description:         "Optional service type for Kubernetes solver service. Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.",
														MarkdownDescription: "Optional service type for Kubernetes solver service. Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"selector": {
										Description:         "Selector selects a set of DNSNames on the Certificate resource that should be solved using this challenge solver. If not specified, the solver will be treated as the 'default' solver with the lowest priority, i.e. if any other solver has a more specific match, it will be used instead.",
										MarkdownDescription: "Selector selects a set of DNSNames on the Certificate resource that should be solved using this challenge solver. If not specified, the solver will be treated as the 'default' solver with the lowest priority, i.e. if any other solver has a more specific match, it will be used instead.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dns_names": {
												Description:         "List of DNSNames that this solver will be used to solve. If specified and a match is found, a dnsNames selector will take precedence over a dnsZones selector. If multiple solvers match with the same dnsNames value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.",
												MarkdownDescription: "List of DNSNames that this solver will be used to solve. If specified and a match is found, a dnsNames selector will take precedence over a dnsZones selector. If multiple solvers match with the same dnsNames value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dns_zones": {
												Description:         "List of DNSZones that this solver will be used to solve. The most specific DNS zone match specified here will take precedence over other DNS zone matches, so a solver specifying sys.example.com will be selected over one specifying example.com for the domain www.sys.example.com. If multiple solvers match with the same dnsZones value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.",
												MarkdownDescription: "List of DNSZones that this solver will be used to solve. The most specific DNS zone match specified here will take precedence over other DNS zone matches, so a solver specifying sys.example.com will be selected over one specifying example.com for the domain www.sys.example.com. If multiple solvers match with the same dnsZones value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": {
												Description:         "A label selector that is used to refine the set of certificate's that this challenge solver will apply to.",
												MarkdownDescription: "A label selector that is used to refine the set of certificate's that this challenge solver will apply to.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ca": {
						Description:         "CA configures this issuer to sign certificates using a signing CA keypair stored in a Secret resource. This is used to build internal PKIs that are managed by cert-manager.",
						MarkdownDescription: "CA configures this issuer to sign certificates using a signing CA keypair stored in a Secret resource. This is used to build internal PKIs that are managed by cert-manager.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"crl_distribution_points": {
								Description:         "The CRL distribution points is an X.509 v3 certificate extension which identifies the location of the CRL from which the revocation of this certificate can be checked. If not set, certificates will be issued without distribution points set.",
								MarkdownDescription: "The CRL distribution points is an X.509 v3 certificate extension which identifies the location of the CRL from which the revocation of this certificate can be checked. If not set, certificates will be issued without distribution points set.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ocsp_servers": {
								Description:         "The OCSP server list is an X.509 v3 extension that defines a list of URLs of OCSP responders. The OCSP responders can be queried for the revocation status of an issued certificate. If not set, the certificate will be issued with no OCSP servers set. For example, an OCSP server URL could be 'http://ocsp.int-x3.letsencrypt.org'.",
								MarkdownDescription: "The OCSP server list is an X.509 v3 extension that defines a list of URLs of OCSP responders. The OCSP responders can be queried for the revocation status of an issued certificate. If not set, the certificate will be issued with no OCSP servers set. For example, an OCSP server URL could be 'http://ocsp.int-x3.letsencrypt.org'.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret_name": {
								Description:         "SecretName is the name of the secret used to sign Certificates issued by this Issuer.",
								MarkdownDescription: "SecretName is the name of the secret used to sign Certificates issued by this Issuer.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"self_signed": {
						Description:         "SelfSigned configures this issuer to 'self sign' certificates using the private key used to create the CertificateRequest object.",
						MarkdownDescription: "SelfSigned configures this issuer to 'self sign' certificates using the private key used to create the CertificateRequest object.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"crl_distribution_points": {
								Description:         "The CRL distribution points is an X.509 v3 certificate extension which identifies the location of the CRL from which the revocation of this certificate can be checked. If not set certificate will be issued without CDP. Values are strings.",
								MarkdownDescription: "The CRL distribution points is an X.509 v3 certificate extension which identifies the location of the CRL from which the revocation of this certificate can be checked. If not set certificate will be issued without CDP. Values are strings.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vault": {
						Description:         "Vault configures this issuer to sign certificates using a HashiCorp Vault PKI backend.",
						MarkdownDescription: "Vault configures this issuer to sign certificates using a HashiCorp Vault PKI backend.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"auth": {
								Description:         "Auth configures how cert-manager authenticates with the Vault server.",
								MarkdownDescription: "Auth configures how cert-manager authenticates with the Vault server.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"app_role": {
										Description:         "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",
										MarkdownDescription: "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"path": {
												Description:         "Path where the App Role authentication backend is mounted in Vault, e.g: 'approle'",
												MarkdownDescription: "Path where the App Role authentication backend is mounted in Vault, e.g: 'approle'",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"role_id": {
												Description:         "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
												MarkdownDescription: "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"secret_ref": {
												Description:         "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
												MarkdownDescription: "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role secret.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kubernetes": {
										Description:         "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
										MarkdownDescription: "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"mount_path": {
												Description:         "The Vault mountPath here is the mount path to use when authenticating with Vault. For example, setting a value to '/v1/auth/foo', will use the path '/v1/auth/foo/login' to authenticate with Vault. If unspecified, the default value '/v1/auth/kubernetes' will be used.",
												MarkdownDescription: "The Vault mountPath here is the mount path to use when authenticating with Vault. For example, setting a value to '/v1/auth/foo', will use the path '/v1/auth/foo/login' to authenticate with Vault. If unspecified, the default value '/v1/auth/kubernetes' will be used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"role": {
												Description:         "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
												MarkdownDescription: "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"secret_ref": {
												Description:         "The required Secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. Use of 'ambient credentials' is not supported.",
												MarkdownDescription: "The required Secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. Use of 'ambient credentials' is not supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"token_secret_ref": {
										Description:         "TokenSecretRef authenticates with Vault by presenting a token.",
										MarkdownDescription: "TokenSecretRef authenticates with Vault by presenting a token.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"ca_bundle": {
								Description:         "PEM-encoded CA bundle (base64-encoded) used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection. Mutually exclusive with CABundleSecretRef. If neither CABundle nor CABundleSecretRef are defined, the cert-manager controller system root certificates are used to validate the TLS connection.",
								MarkdownDescription: "PEM-encoded CA bundle (base64-encoded) used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection. Mutually exclusive with CABundleSecretRef. If neither CABundle nor CABundleSecretRef are defined, the cert-manager controller system root certificates are used to validate the TLS connection.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									validators.Base64Validator(),
								},
							},

							"ca_bundle_secret_ref": {
								Description:         "CABundleSecretRef is a reference to a Secret which contains the CABundle which will be used when connecting to Vault when using HTTPS. Mutually exclusive with CABundle. If neither CABundleSecretRef nor CABundle are defined, the cert-manager controller system root certificates are used to validate the TLS connection. If no key for the Secret is specified, cert-manager will default to 'ca.crt'.",
								MarkdownDescription: "CABundleSecretRef is a reference to a Secret which contains the CABundle which will be used when connecting to Vault when using HTTPS. Mutually exclusive with CABundle. If neither CABundleSecretRef nor CABundle are defined, the cert-manager controller system root certificates are used to validate the TLS connection. If no key for the Secret is specified, cert-manager will default to 'ca.crt'.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
										MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1' More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
								MarkdownDescription: "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1' More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": {
								Description:         "Path is the mount path of the Vault PKI backend's 'sign' endpoint, e.g: 'my_pki_mount/sign/my-role-name'.",
								MarkdownDescription: "Path is the mount path of the Vault PKI backend's 'sign' endpoint, e.g: 'my_pki_mount/sign/my-role-name'.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"server": {
								Description:         "Server is the connection address for the Vault server, e.g: 'https://vault.example.com:8200'.",
								MarkdownDescription: "Server is the connection address for the Vault server, e.g: 'https://vault.example.com:8200'.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"venafi": {
						Description:         "Venafi configures this issuer to sign certificates using a Venafi TPP or Venafi Cloud policy zone.",
						MarkdownDescription: "Venafi configures this issuer to sign certificates using a Venafi TPP or Venafi Cloud policy zone.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cloud": {
								Description:         "Cloud specifies the Venafi cloud configuration settings. Only one of TPP or Cloud may be specified.",
								MarkdownDescription: "Cloud specifies the Venafi cloud configuration settings. Only one of TPP or Cloud may be specified.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"api_token_secret_ref": {
										Description:         "APITokenSecretRef is a secret key selector for the Venafi Cloud API token.",
										MarkdownDescription: "APITokenSecretRef is a secret key selector for the Venafi Cloud API token.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "The key of the entry in the Secret resource's 'data' field to be used. Some instances of this field may be defaulted, in others it may be required.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"url": {
										Description:         "URL is the base URL for Venafi Cloud. Defaults to 'https://api.venafi.cloud/v1'.",
										MarkdownDescription: "URL is the base URL for Venafi Cloud. Defaults to 'https://api.venafi.cloud/v1'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tpp": {
								Description:         "TPP specifies Trust Protection Platform configuration settings. Only one of TPP or Cloud may be specified.",
								MarkdownDescription: "TPP specifies Trust Protection Platform configuration settings. Only one of TPP or Cloud may be specified.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca_bundle": {
										Description:         "CABundle is a PEM encoded TLS certificate to use to verify connections to the TPP instance. If specified, system roots will not be used and the issuing CA for the TPP instance must be verifiable using the provided root. If not specified, the connection will be verified using the cert-manager system root certificates.",
										MarkdownDescription: "CABundle is a PEM encoded TLS certificate to use to verify connections to the TPP instance. If specified, system roots will not be used and the issuing CA for the TPP instance must be verifiable using the provided root. If not specified, the connection will be verified using the cert-manager system root certificates.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											validators.Base64Validator(),
										},
									},

									"credentials_ref": {
										Description:         "CredentialsRef is a reference to a Secret containing the username and password for the TPP server. The secret must contain two keys, 'username' and 'password'.",
										MarkdownDescription: "CredentialsRef is a reference to a Secret containing the username and password for the TPP server. The secret must contain two keys, 'username' and 'password'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the resource being referred to. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"url": {
										Description:         "URL is the base URL for the vedsdk endpoint of the Venafi TPP instance, for example: 'https://tpp.example.com/vedsdk'.",
										MarkdownDescription: "URL is the base URL for the vedsdk endpoint of the Venafi TPP instance, for example: 'https://tpp.example.com/vedsdk'.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"zone": {
								Description:         "Zone is the Venafi Policy Zone to use for this issuer. All requests made to the Venafi platform will be restricted by the named zone policy. This field is required.",
								MarkdownDescription: "Zone is the Venafi Policy Zone to use for this issuer. All requests made to the Venafi platform will be restricted by the named zone policy. This field is required.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *CertManagerIoClusterIssuerV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_cert_manager_io_cluster_issuer_v1")

	var state CertManagerIoClusterIssuerV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CertManagerIoClusterIssuerV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cert-manager.io/v1")
	goModel.Kind = utilities.Ptr("ClusterIssuer")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CertManagerIoClusterIssuerV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cert_manager_io_cluster_issuer_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CertManagerIoClusterIssuerV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_cert_manager_io_cluster_issuer_v1")

	var state CertManagerIoClusterIssuerV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CertManagerIoClusterIssuerV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cert-manager.io/v1")
	goModel.Kind = utilities.Ptr("ClusterIssuer")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CertManagerIoClusterIssuerV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_cert_manager_io_cluster_issuer_v1")
	// NO-OP: Terraform removes the state automatically for us
}
