/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type GatewaySoloIoVirtualHostOptionV1Resource struct{}

var (
	_ resource.Resource = (*GatewaySoloIoVirtualHostOptionV1Resource)(nil)
)

type GatewaySoloIoVirtualHostOptionV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GatewaySoloIoVirtualHostOptionV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		NamespacedStatuses *struct {
			Statuses utilities.Dynamic `tfsdk:"statuses" yaml:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" yaml:"namespacedStatuses,omitempty"`

		Options *struct {
			BufferPerRoute *struct {
				Buffer *struct {
					MaxRequestBytes *int64 `tfsdk:"max_request_bytes" yaml:"maxRequestBytes,omitempty"`
				} `tfsdk:"buffer" yaml:"buffer,omitempty"`

				Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`
			} `tfsdk:"buffer_per_route" yaml:"bufferPerRoute,omitempty"`

			Cors *struct {
				AllowCredentials *bool `tfsdk:"allow_credentials" yaml:"allowCredentials,omitempty"`

				AllowHeaders *[]string `tfsdk:"allow_headers" yaml:"allowHeaders,omitempty"`

				AllowMethods *[]string `tfsdk:"allow_methods" yaml:"allowMethods,omitempty"`

				AllowOrigin *[]string `tfsdk:"allow_origin" yaml:"allowOrigin,omitempty"`

				AllowOriginRegex *[]string `tfsdk:"allow_origin_regex" yaml:"allowOriginRegex,omitempty"`

				DisableForRoute *bool `tfsdk:"disable_for_route" yaml:"disableForRoute,omitempty"`

				ExposeHeaders *[]string `tfsdk:"expose_headers" yaml:"exposeHeaders,omitempty"`

				MaxAge *string `tfsdk:"max_age" yaml:"maxAge,omitempty"`
			} `tfsdk:"cors" yaml:"cors,omitempty"`

			Csrf *struct {
				AdditionalOrigins *[]struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					IgnoreCase *bool `tfsdk:"ignore_case" yaml:"ignoreCase,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					SafeRegex *struct {
						GoogleRe2 *struct {
							MaxProgramSize *int64 `tfsdk:"max_program_size" yaml:"maxProgramSize,omitempty"`
						} `tfsdk:"google_re2" yaml:"googleRe2,omitempty"`

						Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
					} `tfsdk:"safe_regex" yaml:"safeRegex,omitempty"`

					Suffix *string `tfsdk:"suffix" yaml:"suffix,omitempty"`
				} `tfsdk:"additional_origins" yaml:"additionalOrigins,omitempty"`

				FilterEnabled *struct {
					DefaultValue *struct {
						Denominator utilities.IntOrString `tfsdk:"denominator" yaml:"denominator,omitempty"`

						Numerator *int64 `tfsdk:"numerator" yaml:"numerator,omitempty"`
					} `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

					RuntimeKey *string `tfsdk:"runtime_key" yaml:"runtimeKey,omitempty"`
				} `tfsdk:"filter_enabled" yaml:"filterEnabled,omitempty"`

				ShadowEnabled *struct {
					DefaultValue *struct {
						Denominator utilities.IntOrString `tfsdk:"denominator" yaml:"denominator,omitempty"`

						Numerator *int64 `tfsdk:"numerator" yaml:"numerator,omitempty"`
					} `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

					RuntimeKey *string `tfsdk:"runtime_key" yaml:"runtimeKey,omitempty"`
				} `tfsdk:"shadow_enabled" yaml:"shadowEnabled,omitempty"`
			} `tfsdk:"csrf" yaml:"csrf,omitempty"`

			Dlp *struct {
				Actions *[]struct {
					ActionType utilities.IntOrString `tfsdk:"action_type" yaml:"actionType,omitempty"`

					CustomAction *struct {
						MaskChar *string `tfsdk:"mask_char" yaml:"maskChar,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Percent *struct {
							Value utilities.DynamicNumber `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"percent" yaml:"percent,omitempty"`

						Regex *[]string `tfsdk:"regex" yaml:"regex,omitempty"`

						RegexActions *[]struct {
							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

							Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
						} `tfsdk:"regex_actions" yaml:"regexActions,omitempty"`
					} `tfsdk:"custom_action" yaml:"customAction,omitempty"`

					KeyValueAction *struct {
						KeyToMask *string `tfsdk:"key_to_mask" yaml:"keyToMask,omitempty"`

						MaskChar *string `tfsdk:"mask_char" yaml:"maskChar,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Percent *struct {
							Value utilities.DynamicNumber `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"percent" yaml:"percent,omitempty"`
					} `tfsdk:"key_value_action" yaml:"keyValueAction,omitempty"`

					Shadow *bool `tfsdk:"shadow" yaml:"shadow,omitempty"`
				} `tfsdk:"actions" yaml:"actions,omitempty"`

				EnabledFor utilities.IntOrString `tfsdk:"enabled_for" yaml:"enabledFor,omitempty"`
			} `tfsdk:"dlp" yaml:"dlp,omitempty"`

			Extauth *struct {
				ConfigRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"config_ref" yaml:"configRef,omitempty"`

				CustomAuth *struct {
					ContextExtensions *map[string]string `tfsdk:"context_extensions" yaml:"contextExtensions,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"custom_auth" yaml:"customAuth,omitempty"`

				Disable *bool `tfsdk:"disable" yaml:"disable,omitempty"`
			} `tfsdk:"extauth" yaml:"extauth,omitempty"`

			Extensions *struct {
				Configs utilities.Dynamic `tfsdk:"configs" yaml:"configs,omitempty"`
			} `tfsdk:"extensions" yaml:"extensions,omitempty"`

			HeaderManipulation *struct {
				RequestHeadersToAdd *[]struct {
					Append *bool `tfsdk:"append" yaml:"append,omitempty"`

					Header *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"header" yaml:"header,omitempty"`

					HeaderSecretRef *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"header_secret_ref" yaml:"headerSecretRef,omitempty"`
				} `tfsdk:"request_headers_to_add" yaml:"requestHeadersToAdd,omitempty"`

				RequestHeadersToRemove *[]string `tfsdk:"request_headers_to_remove" yaml:"requestHeadersToRemove,omitempty"`

				ResponseHeadersToAdd *[]struct {
					Append *bool `tfsdk:"append" yaml:"append,omitempty"`

					Header *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"header" yaml:"header,omitempty"`
				} `tfsdk:"response_headers_to_add" yaml:"responseHeadersToAdd,omitempty"`

				ResponseHeadersToRemove *[]string `tfsdk:"response_headers_to_remove" yaml:"responseHeadersToRemove,omitempty"`
			} `tfsdk:"header_manipulation" yaml:"headerManipulation,omitempty"`

			IncludeAttemptCountInResponse *bool `tfsdk:"include_attempt_count_in_response" yaml:"includeAttemptCountInResponse,omitempty"`

			IncludeRequestAttemptCount *bool `tfsdk:"include_request_attempt_count" yaml:"includeRequestAttemptCount,omitempty"`

			Jwt *struct {
				AllowMissingOrFailedJwt *bool `tfsdk:"allow_missing_or_failed_jwt" yaml:"allowMissingOrFailedJwt,omitempty"`

				Providers *struct {
					Audiences *[]string `tfsdk:"audiences" yaml:"audiences,omitempty"`

					ClaimsToHeaders *[]struct {
						Append *bool `tfsdk:"append" yaml:"append,omitempty"`

						Claim *string `tfsdk:"claim" yaml:"claim,omitempty"`

						Header *string `tfsdk:"header" yaml:"header,omitempty"`
					} `tfsdk:"claims_to_headers" yaml:"claimsToHeaders,omitempty"`

					Issuer *string `tfsdk:"issuer" yaml:"issuer,omitempty"`

					Jwks *struct {
						Local *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`
						} `tfsdk:"local" yaml:"local,omitempty"`

						Remote *struct {
							AsyncFetch *struct {
								FastListener *bool `tfsdk:"fast_listener" yaml:"fastListener,omitempty"`
							} `tfsdk:"async_fetch" yaml:"asyncFetch,omitempty"`

							CacheDuration *string `tfsdk:"cache_duration" yaml:"cacheDuration,omitempty"`

							UpstreamRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"upstream_ref" yaml:"upstreamRef,omitempty"`

							Url *string `tfsdk:"url" yaml:"url,omitempty"`
						} `tfsdk:"remote" yaml:"remote,omitempty"`
					} `tfsdk:"jwks" yaml:"jwks,omitempty"`

					KeepToken *bool `tfsdk:"keep_token" yaml:"keepToken,omitempty"`

					TokenSource *struct {
						Headers *[]struct {
							Header *string `tfsdk:"header" yaml:"header,omitempty"`

							Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`
						} `tfsdk:"headers" yaml:"headers,omitempty"`

						QueryParams *[]string `tfsdk:"query_params" yaml:"queryParams,omitempty"`
					} `tfsdk:"token_source" yaml:"tokenSource,omitempty"`
				} `tfsdk:"providers" yaml:"providers,omitempty"`
			} `tfsdk:"jwt" yaml:"jwt,omitempty"`

			JwtStaged *struct {
				AfterExtAuth *struct {
					AllowMissingOrFailedJwt *bool `tfsdk:"allow_missing_or_failed_jwt" yaml:"allowMissingOrFailedJwt,omitempty"`

					Providers *struct {
						Audiences *[]string `tfsdk:"audiences" yaml:"audiences,omitempty"`

						ClaimsToHeaders *[]struct {
							Append *bool `tfsdk:"append" yaml:"append,omitempty"`

							Claim *string `tfsdk:"claim" yaml:"claim,omitempty"`

							Header *string `tfsdk:"header" yaml:"header,omitempty"`
						} `tfsdk:"claims_to_headers" yaml:"claimsToHeaders,omitempty"`

						Issuer *string `tfsdk:"issuer" yaml:"issuer,omitempty"`

						Jwks *struct {
							Local *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`
							} `tfsdk:"local" yaml:"local,omitempty"`

							Remote *struct {
								AsyncFetch *struct {
									FastListener *bool `tfsdk:"fast_listener" yaml:"fastListener,omitempty"`
								} `tfsdk:"async_fetch" yaml:"asyncFetch,omitempty"`

								CacheDuration *string `tfsdk:"cache_duration" yaml:"cacheDuration,omitempty"`

								UpstreamRef *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"upstream_ref" yaml:"upstreamRef,omitempty"`

								Url *string `tfsdk:"url" yaml:"url,omitempty"`
							} `tfsdk:"remote" yaml:"remote,omitempty"`
						} `tfsdk:"jwks" yaml:"jwks,omitempty"`

						KeepToken *bool `tfsdk:"keep_token" yaml:"keepToken,omitempty"`

						TokenSource *struct {
							Headers *[]struct {
								Header *string `tfsdk:"header" yaml:"header,omitempty"`

								Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`

							QueryParams *[]string `tfsdk:"query_params" yaml:"queryParams,omitempty"`
						} `tfsdk:"token_source" yaml:"tokenSource,omitempty"`
					} `tfsdk:"providers" yaml:"providers,omitempty"`
				} `tfsdk:"after_ext_auth" yaml:"afterExtAuth,omitempty"`

				BeforeExtAuth *struct {
					AllowMissingOrFailedJwt *bool `tfsdk:"allow_missing_or_failed_jwt" yaml:"allowMissingOrFailedJwt,omitempty"`

					Providers *struct {
						Audiences *[]string `tfsdk:"audiences" yaml:"audiences,omitempty"`

						ClaimsToHeaders *[]struct {
							Append *bool `tfsdk:"append" yaml:"append,omitempty"`

							Claim *string `tfsdk:"claim" yaml:"claim,omitempty"`

							Header *string `tfsdk:"header" yaml:"header,omitempty"`
						} `tfsdk:"claims_to_headers" yaml:"claimsToHeaders,omitempty"`

						Issuer *string `tfsdk:"issuer" yaml:"issuer,omitempty"`

						Jwks *struct {
							Local *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`
							} `tfsdk:"local" yaml:"local,omitempty"`

							Remote *struct {
								AsyncFetch *struct {
									FastListener *bool `tfsdk:"fast_listener" yaml:"fastListener,omitempty"`
								} `tfsdk:"async_fetch" yaml:"asyncFetch,omitempty"`

								CacheDuration *string `tfsdk:"cache_duration" yaml:"cacheDuration,omitempty"`

								UpstreamRef *struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"upstream_ref" yaml:"upstreamRef,omitempty"`

								Url *string `tfsdk:"url" yaml:"url,omitempty"`
							} `tfsdk:"remote" yaml:"remote,omitempty"`
						} `tfsdk:"jwks" yaml:"jwks,omitempty"`

						KeepToken *bool `tfsdk:"keep_token" yaml:"keepToken,omitempty"`

						TokenSource *struct {
							Headers *[]struct {
								Header *string `tfsdk:"header" yaml:"header,omitempty"`

								Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`

							QueryParams *[]string `tfsdk:"query_params" yaml:"queryParams,omitempty"`
						} `tfsdk:"token_source" yaml:"tokenSource,omitempty"`
					} `tfsdk:"providers" yaml:"providers,omitempty"`
				} `tfsdk:"before_ext_auth" yaml:"beforeExtAuth,omitempty"`
			} `tfsdk:"jwt_staged" yaml:"jwtStaged,omitempty"`

			RateLimitConfigs *struct {
				Refs *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"refs" yaml:"refs,omitempty"`
			} `tfsdk:"rate_limit_configs" yaml:"rateLimitConfigs,omitempty"`

			RateLimitEarlyConfigs *struct {
				Refs *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"refs" yaml:"refs,omitempty"`
			} `tfsdk:"rate_limit_early_configs" yaml:"rateLimitEarlyConfigs,omitempty"`

			RateLimitRegularConfigs *struct {
				Refs *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"refs" yaml:"refs,omitempty"`
			} `tfsdk:"rate_limit_regular_configs" yaml:"rateLimitRegularConfigs,omitempty"`

			Ratelimit *struct {
				RateLimits *[]struct {
					Actions *[]struct {
						DestinationCluster *map[string]string `tfsdk:"destination_cluster" yaml:"destinationCluster,omitempty"`

						GenericKey *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`
						} `tfsdk:"generic_key" yaml:"genericKey,omitempty"`

						HeaderValueMatch *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`

							ExpectMatch *bool `tfsdk:"expect_match" yaml:"expectMatch,omitempty"`

							Headers *[]struct {
								ExactMatch *string `tfsdk:"exact_match" yaml:"exactMatch,omitempty"`

								InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								PrefixMatch *string `tfsdk:"prefix_match" yaml:"prefixMatch,omitempty"`

								PresentMatch *bool `tfsdk:"present_match" yaml:"presentMatch,omitempty"`

								RangeMatch *struct {
									End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

									Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
								} `tfsdk:"range_match" yaml:"rangeMatch,omitempty"`

								RegexMatch *string `tfsdk:"regex_match" yaml:"regexMatch,omitempty"`

								SuffixMatch *string `tfsdk:"suffix_match" yaml:"suffixMatch,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`
						} `tfsdk:"header_value_match" yaml:"headerValueMatch,omitempty"`

						Metadata *struct {
							DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

							DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

							MetadataKey *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Path *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`
								} `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"metadata_key" yaml:"metadataKey,omitempty"`

							Source utilities.IntOrString `tfsdk:"source" yaml:"source,omitempty"`
						} `tfsdk:"metadata" yaml:"metadata,omitempty"`

						RemoteAddress *map[string]string `tfsdk:"remote_address" yaml:"remoteAddress,omitempty"`

						RequestHeaders *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

							HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`
						} `tfsdk:"request_headers" yaml:"requestHeaders,omitempty"`

						SourceCluster *map[string]string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`
					} `tfsdk:"actions" yaml:"actions,omitempty"`

					SetActions *[]struct {
						DestinationCluster *map[string]string `tfsdk:"destination_cluster" yaml:"destinationCluster,omitempty"`

						GenericKey *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`
						} `tfsdk:"generic_key" yaml:"genericKey,omitempty"`

						HeaderValueMatch *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`

							ExpectMatch *bool `tfsdk:"expect_match" yaml:"expectMatch,omitempty"`

							Headers *[]struct {
								ExactMatch *string `tfsdk:"exact_match" yaml:"exactMatch,omitempty"`

								InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								PrefixMatch *string `tfsdk:"prefix_match" yaml:"prefixMatch,omitempty"`

								PresentMatch *bool `tfsdk:"present_match" yaml:"presentMatch,omitempty"`

								RangeMatch *struct {
									End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

									Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
								} `tfsdk:"range_match" yaml:"rangeMatch,omitempty"`

								RegexMatch *string `tfsdk:"regex_match" yaml:"regexMatch,omitempty"`

								SuffixMatch *string `tfsdk:"suffix_match" yaml:"suffixMatch,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`
						} `tfsdk:"header_value_match" yaml:"headerValueMatch,omitempty"`

						Metadata *struct {
							DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

							DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

							MetadataKey *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Path *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`
								} `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"metadata_key" yaml:"metadataKey,omitempty"`

							Source utilities.IntOrString `tfsdk:"source" yaml:"source,omitempty"`
						} `tfsdk:"metadata" yaml:"metadata,omitempty"`

						RemoteAddress *map[string]string `tfsdk:"remote_address" yaml:"remoteAddress,omitempty"`

						RequestHeaders *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

							HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`
						} `tfsdk:"request_headers" yaml:"requestHeaders,omitempty"`

						SourceCluster *map[string]string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`
					} `tfsdk:"set_actions" yaml:"setActions,omitempty"`
				} `tfsdk:"rate_limits" yaml:"rateLimits,omitempty"`
			} `tfsdk:"ratelimit" yaml:"ratelimit,omitempty"`

			RatelimitBasic *struct {
				AnonymousLimits *struct {
					RequestsPerUnit *int64 `tfsdk:"requests_per_unit" yaml:"requestsPerUnit,omitempty"`

					Unit utilities.IntOrString `tfsdk:"unit" yaml:"unit,omitempty"`
				} `tfsdk:"anonymous_limits" yaml:"anonymousLimits,omitempty"`

				AuthorizedLimits *struct {
					RequestsPerUnit *int64 `tfsdk:"requests_per_unit" yaml:"requestsPerUnit,omitempty"`

					Unit utilities.IntOrString `tfsdk:"unit" yaml:"unit,omitempty"`
				} `tfsdk:"authorized_limits" yaml:"authorizedLimits,omitempty"`
			} `tfsdk:"ratelimit_basic" yaml:"ratelimitBasic,omitempty"`

			RatelimitEarly *struct {
				RateLimits *[]struct {
					Actions *[]struct {
						DestinationCluster *map[string]string `tfsdk:"destination_cluster" yaml:"destinationCluster,omitempty"`

						GenericKey *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`
						} `tfsdk:"generic_key" yaml:"genericKey,omitempty"`

						HeaderValueMatch *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`

							ExpectMatch *bool `tfsdk:"expect_match" yaml:"expectMatch,omitempty"`

							Headers *[]struct {
								ExactMatch *string `tfsdk:"exact_match" yaml:"exactMatch,omitempty"`

								InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								PrefixMatch *string `tfsdk:"prefix_match" yaml:"prefixMatch,omitempty"`

								PresentMatch *bool `tfsdk:"present_match" yaml:"presentMatch,omitempty"`

								RangeMatch *struct {
									End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

									Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
								} `tfsdk:"range_match" yaml:"rangeMatch,omitempty"`

								RegexMatch *string `tfsdk:"regex_match" yaml:"regexMatch,omitempty"`

								SuffixMatch *string `tfsdk:"suffix_match" yaml:"suffixMatch,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`
						} `tfsdk:"header_value_match" yaml:"headerValueMatch,omitempty"`

						Metadata *struct {
							DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

							DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

							MetadataKey *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Path *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`
								} `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"metadata_key" yaml:"metadataKey,omitempty"`

							Source utilities.IntOrString `tfsdk:"source" yaml:"source,omitempty"`
						} `tfsdk:"metadata" yaml:"metadata,omitempty"`

						RemoteAddress *map[string]string `tfsdk:"remote_address" yaml:"remoteAddress,omitempty"`

						RequestHeaders *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

							HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`
						} `tfsdk:"request_headers" yaml:"requestHeaders,omitempty"`

						SourceCluster *map[string]string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`
					} `tfsdk:"actions" yaml:"actions,omitempty"`

					SetActions *[]struct {
						DestinationCluster *map[string]string `tfsdk:"destination_cluster" yaml:"destinationCluster,omitempty"`

						GenericKey *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`
						} `tfsdk:"generic_key" yaml:"genericKey,omitempty"`

						HeaderValueMatch *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`

							ExpectMatch *bool `tfsdk:"expect_match" yaml:"expectMatch,omitempty"`

							Headers *[]struct {
								ExactMatch *string `tfsdk:"exact_match" yaml:"exactMatch,omitempty"`

								InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								PrefixMatch *string `tfsdk:"prefix_match" yaml:"prefixMatch,omitempty"`

								PresentMatch *bool `tfsdk:"present_match" yaml:"presentMatch,omitempty"`

								RangeMatch *struct {
									End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

									Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
								} `tfsdk:"range_match" yaml:"rangeMatch,omitempty"`

								RegexMatch *string `tfsdk:"regex_match" yaml:"regexMatch,omitempty"`

								SuffixMatch *string `tfsdk:"suffix_match" yaml:"suffixMatch,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`
						} `tfsdk:"header_value_match" yaml:"headerValueMatch,omitempty"`

						Metadata *struct {
							DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

							DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

							MetadataKey *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Path *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`
								} `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"metadata_key" yaml:"metadataKey,omitempty"`

							Source utilities.IntOrString `tfsdk:"source" yaml:"source,omitempty"`
						} `tfsdk:"metadata" yaml:"metadata,omitempty"`

						RemoteAddress *map[string]string `tfsdk:"remote_address" yaml:"remoteAddress,omitempty"`

						RequestHeaders *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

							HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`
						} `tfsdk:"request_headers" yaml:"requestHeaders,omitempty"`

						SourceCluster *map[string]string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`
					} `tfsdk:"set_actions" yaml:"setActions,omitempty"`
				} `tfsdk:"rate_limits" yaml:"rateLimits,omitempty"`
			} `tfsdk:"ratelimit_early" yaml:"ratelimitEarly,omitempty"`

			RatelimitRegular *struct {
				RateLimits *[]struct {
					Actions *[]struct {
						DestinationCluster *map[string]string `tfsdk:"destination_cluster" yaml:"destinationCluster,omitempty"`

						GenericKey *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`
						} `tfsdk:"generic_key" yaml:"genericKey,omitempty"`

						HeaderValueMatch *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`

							ExpectMatch *bool `tfsdk:"expect_match" yaml:"expectMatch,omitempty"`

							Headers *[]struct {
								ExactMatch *string `tfsdk:"exact_match" yaml:"exactMatch,omitempty"`

								InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								PrefixMatch *string `tfsdk:"prefix_match" yaml:"prefixMatch,omitempty"`

								PresentMatch *bool `tfsdk:"present_match" yaml:"presentMatch,omitempty"`

								RangeMatch *struct {
									End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

									Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
								} `tfsdk:"range_match" yaml:"rangeMatch,omitempty"`

								RegexMatch *string `tfsdk:"regex_match" yaml:"regexMatch,omitempty"`

								SuffixMatch *string `tfsdk:"suffix_match" yaml:"suffixMatch,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`
						} `tfsdk:"header_value_match" yaml:"headerValueMatch,omitempty"`

						Metadata *struct {
							DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

							DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

							MetadataKey *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Path *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`
								} `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"metadata_key" yaml:"metadataKey,omitempty"`

							Source utilities.IntOrString `tfsdk:"source" yaml:"source,omitempty"`
						} `tfsdk:"metadata" yaml:"metadata,omitempty"`

						RemoteAddress *map[string]string `tfsdk:"remote_address" yaml:"remoteAddress,omitempty"`

						RequestHeaders *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

							HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`
						} `tfsdk:"request_headers" yaml:"requestHeaders,omitempty"`

						SourceCluster *map[string]string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`
					} `tfsdk:"actions" yaml:"actions,omitempty"`

					SetActions *[]struct {
						DestinationCluster *map[string]string `tfsdk:"destination_cluster" yaml:"destinationCluster,omitempty"`

						GenericKey *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`
						} `tfsdk:"generic_key" yaml:"genericKey,omitempty"`

						HeaderValueMatch *struct {
							DescriptorValue *string `tfsdk:"descriptor_value" yaml:"descriptorValue,omitempty"`

							ExpectMatch *bool `tfsdk:"expect_match" yaml:"expectMatch,omitempty"`

							Headers *[]struct {
								ExactMatch *string `tfsdk:"exact_match" yaml:"exactMatch,omitempty"`

								InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								PrefixMatch *string `tfsdk:"prefix_match" yaml:"prefixMatch,omitempty"`

								PresentMatch *bool `tfsdk:"present_match" yaml:"presentMatch,omitempty"`

								RangeMatch *struct {
									End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

									Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
								} `tfsdk:"range_match" yaml:"rangeMatch,omitempty"`

								RegexMatch *string `tfsdk:"regex_match" yaml:"regexMatch,omitempty"`

								SuffixMatch *string `tfsdk:"suffix_match" yaml:"suffixMatch,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`
						} `tfsdk:"header_value_match" yaml:"headerValueMatch,omitempty"`

						Metadata *struct {
							DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

							DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

							MetadataKey *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Path *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`
								} `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"metadata_key" yaml:"metadataKey,omitempty"`

							Source utilities.IntOrString `tfsdk:"source" yaml:"source,omitempty"`
						} `tfsdk:"metadata" yaml:"metadata,omitempty"`

						RemoteAddress *map[string]string `tfsdk:"remote_address" yaml:"remoteAddress,omitempty"`

						RequestHeaders *struct {
							DescriptorKey *string `tfsdk:"descriptor_key" yaml:"descriptorKey,omitempty"`

							HeaderName *string `tfsdk:"header_name" yaml:"headerName,omitempty"`
						} `tfsdk:"request_headers" yaml:"requestHeaders,omitempty"`

						SourceCluster *map[string]string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`
					} `tfsdk:"set_actions" yaml:"setActions,omitempty"`
				} `tfsdk:"rate_limits" yaml:"rateLimits,omitempty"`
			} `tfsdk:"ratelimit_regular" yaml:"ratelimitRegular,omitempty"`

			Rbac *struct {
				Disable *bool `tfsdk:"disable" yaml:"disable,omitempty"`

				Policies *struct {
					NestedClaimDelimiter *string `tfsdk:"nested_claim_delimiter" yaml:"nestedClaimDelimiter,omitempty"`

					Permissions *struct {
						Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

						PathPrefix *string `tfsdk:"path_prefix" yaml:"pathPrefix,omitempty"`
					} `tfsdk:"permissions" yaml:"permissions,omitempty"`

					Principals *[]struct {
						JwtPrincipal *struct {
							Claims *map[string]string `tfsdk:"claims" yaml:"claims,omitempty"`

							Matcher utilities.IntOrString `tfsdk:"matcher" yaml:"matcher,omitempty"`

							Provider *string `tfsdk:"provider" yaml:"provider,omitempty"`
						} `tfsdk:"jwt_principal" yaml:"jwtPrincipal,omitempty"`
					} `tfsdk:"principals" yaml:"principals,omitempty"`
				} `tfsdk:"policies" yaml:"policies,omitempty"`
			} `tfsdk:"rbac" yaml:"rbac,omitempty"`

			Retries *struct {
				NumRetries *int64 `tfsdk:"num_retries" yaml:"numRetries,omitempty"`

				PerTryTimeout *string `tfsdk:"per_try_timeout" yaml:"perTryTimeout,omitempty"`

				RetryOn *string `tfsdk:"retry_on" yaml:"retryOn,omitempty"`
			} `tfsdk:"retries" yaml:"retries,omitempty"`

			StagedTransformations *struct {
				Early *struct {
					RequestTransforms *[]struct {
						ClearRouteCache *bool `tfsdk:"clear_route_cache" yaml:"clearRouteCache,omitempty"`

						Matcher *struct {
							CaseSensitive *bool `tfsdk:"case_sensitive" yaml:"caseSensitive,omitempty"`

							Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

							Headers *[]struct {
								InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`

							Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

							Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

							QueryParameters *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"query_parameters" yaml:"queryParameters,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
						} `tfsdk:"matcher" yaml:"matcher,omitempty"`

						RequestTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

								Body *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"body" yaml:"body,omitempty"`

								DynamicMetadataValues *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

								Extractors *struct {
									Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

									Header *string `tfsdk:"header" yaml:"header,omitempty"`

									Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

									Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
								} `tfsdk:"extractors" yaml:"extractors,omitempty"`

								Headers *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`

								HeadersToAppend *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

								HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

								IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

								ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

								Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

							XsltTransformation *struct {
								NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

								SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

								Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
						} `tfsdk:"request_transformation" yaml:"requestTransformation,omitempty"`

						ResponseTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

								Body *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"body" yaml:"body,omitempty"`

								DynamicMetadataValues *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

								Extractors *struct {
									Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

									Header *string `tfsdk:"header" yaml:"header,omitempty"`

									Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

									Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
								} `tfsdk:"extractors" yaml:"extractors,omitempty"`

								Headers *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`

								HeadersToAppend *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

								HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

								IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

								ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

								Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

							XsltTransformation *struct {
								NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

								SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

								Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
						} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
					} `tfsdk:"request_transforms" yaml:"requestTransforms,omitempty"`

					ResponseTransforms *[]struct {
						Matchers *[]struct {
							InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"matchers" yaml:"matchers,omitempty"`

						ResponseCodeDetails *string `tfsdk:"response_code_details" yaml:"responseCodeDetails,omitempty"`

						ResponseTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

								Body *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"body" yaml:"body,omitempty"`

								DynamicMetadataValues *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

								Extractors *struct {
									Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

									Header *string `tfsdk:"header" yaml:"header,omitempty"`

									Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

									Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
								} `tfsdk:"extractors" yaml:"extractors,omitempty"`

								Headers *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`

								HeadersToAppend *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

								HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

								IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

								ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

								Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

							XsltTransformation *struct {
								NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

								SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

								Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
						} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
					} `tfsdk:"response_transforms" yaml:"responseTransforms,omitempty"`
				} `tfsdk:"early" yaml:"early,omitempty"`

				InheritTransformation *bool `tfsdk:"inherit_transformation" yaml:"inheritTransformation,omitempty"`

				Regular *struct {
					RequestTransforms *[]struct {
						ClearRouteCache *bool `tfsdk:"clear_route_cache" yaml:"clearRouteCache,omitempty"`

						Matcher *struct {
							CaseSensitive *bool `tfsdk:"case_sensitive" yaml:"caseSensitive,omitempty"`

							Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

							Headers *[]struct {
								InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"headers" yaml:"headers,omitempty"`

							Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

							Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

							QueryParameters *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"query_parameters" yaml:"queryParameters,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
						} `tfsdk:"matcher" yaml:"matcher,omitempty"`

						RequestTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

								Body *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"body" yaml:"body,omitempty"`

								DynamicMetadataValues *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

								Extractors *struct {
									Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

									Header *string `tfsdk:"header" yaml:"header,omitempty"`

									Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

									Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
								} `tfsdk:"extractors" yaml:"extractors,omitempty"`

								Headers *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`

								HeadersToAppend *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

								HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

								IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

								ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

								Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

							XsltTransformation *struct {
								NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

								SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

								Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
						} `tfsdk:"request_transformation" yaml:"requestTransformation,omitempty"`

						ResponseTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

								Body *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"body" yaml:"body,omitempty"`

								DynamicMetadataValues *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

								Extractors *struct {
									Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

									Header *string `tfsdk:"header" yaml:"header,omitempty"`

									Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

									Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
								} `tfsdk:"extractors" yaml:"extractors,omitempty"`

								Headers *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`

								HeadersToAppend *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

								HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

								IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

								ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

								Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

							XsltTransformation *struct {
								NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

								SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

								Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
						} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
					} `tfsdk:"request_transforms" yaml:"requestTransforms,omitempty"`

					ResponseTransforms *[]struct {
						Matchers *[]struct {
							InvertMatch *bool `tfsdk:"invert_match" yaml:"invertMatch,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"matchers" yaml:"matchers,omitempty"`

						ResponseCodeDetails *string `tfsdk:"response_code_details" yaml:"responseCodeDetails,omitempty"`

						ResponseTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

								Body *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"body" yaml:"body,omitempty"`

								DynamicMetadataValues *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

								Extractors *struct {
									Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

									Header *string `tfsdk:"header" yaml:"header,omitempty"`

									Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

									Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
								} `tfsdk:"extractors" yaml:"extractors,omitempty"`

								Headers *struct {
									Text *string `tfsdk:"text" yaml:"text,omitempty"`
								} `tfsdk:"headers" yaml:"headers,omitempty"`

								HeadersToAppend *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Value *struct {
										Text *string `tfsdk:"text" yaml:"text,omitempty"`
									} `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

								HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

								IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

								ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

								Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

							XsltTransformation *struct {
								NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

								SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

								Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
						} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
					} `tfsdk:"response_transforms" yaml:"responseTransforms,omitempty"`
				} `tfsdk:"regular" yaml:"regular,omitempty"`
			} `tfsdk:"staged_transformations" yaml:"stagedTransformations,omitempty"`

			Stats *struct {
				VirtualClusters *[]struct {
					Method *string `tfsdk:"method" yaml:"method,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Pattern *string `tfsdk:"pattern" yaml:"pattern,omitempty"`
				} `tfsdk:"virtual_clusters" yaml:"virtualClusters,omitempty"`
			} `tfsdk:"stats" yaml:"stats,omitempty"`

			Transformations *struct {
				ClearRouteCache *bool `tfsdk:"clear_route_cache" yaml:"clearRouteCache,omitempty"`

				RequestTransformation *struct {
					HeaderBodyTransform *struct {
						AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
					} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

					TransformationTemplate *struct {
						AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

						Body *struct {
							Text *string `tfsdk:"text" yaml:"text,omitempty"`
						} `tfsdk:"body" yaml:"body,omitempty"`

						DynamicMetadataValues *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

							Value *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

						Extractors *struct {
							Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

							Header *string `tfsdk:"header" yaml:"header,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

							Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
						} `tfsdk:"extractors" yaml:"extractors,omitempty"`

						Headers *struct {
							Text *string `tfsdk:"text" yaml:"text,omitempty"`
						} `tfsdk:"headers" yaml:"headers,omitempty"`

						HeadersToAppend *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Value *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

						HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

						IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

						MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

						ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

						Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
					} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

					XsltTransformation *struct {
						NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

						SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

						Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
					} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
				} `tfsdk:"request_transformation" yaml:"requestTransformation,omitempty"`

				ResponseTransformation *struct {
					HeaderBodyTransform *struct {
						AddRequestMetadata *bool `tfsdk:"add_request_metadata" yaml:"addRequestMetadata,omitempty"`
					} `tfsdk:"header_body_transform" yaml:"headerBodyTransform,omitempty"`

					TransformationTemplate *struct {
						AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

						Body *struct {
							Text *string `tfsdk:"text" yaml:"text,omitempty"`
						} `tfsdk:"body" yaml:"body,omitempty"`

						DynamicMetadataValues *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

							Value *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

						Extractors *struct {
							Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

							Header *string `tfsdk:"header" yaml:"header,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

							Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
						} `tfsdk:"extractors" yaml:"extractors,omitempty"`

						Headers *struct {
							Text *string `tfsdk:"text" yaml:"text,omitempty"`
						} `tfsdk:"headers" yaml:"headers,omitempty"`

						HeadersToAppend *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Value *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

						HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

						IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

						MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

						ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

						Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
					} `tfsdk:"transformation_template" yaml:"transformationTemplate,omitempty"`

					XsltTransformation *struct {
						NonXmlTransform *bool `tfsdk:"non_xml_transform" yaml:"nonXmlTransform,omitempty"`

						SetContentType *string `tfsdk:"set_content_type" yaml:"setContentType,omitempty"`

						Xslt *string `tfsdk:"xslt" yaml:"xslt,omitempty"`
					} `tfsdk:"xslt_transformation" yaml:"xsltTransformation,omitempty"`
				} `tfsdk:"response_transformation" yaml:"responseTransformation,omitempty"`
			} `tfsdk:"transformations" yaml:"transformations,omitempty"`

			Waf *struct {
				AuditLogging *struct {
					Action utilities.IntOrString `tfsdk:"action" yaml:"action,omitempty"`

					Location utilities.IntOrString `tfsdk:"location" yaml:"location,omitempty"`
				} `tfsdk:"audit_logging" yaml:"auditLogging,omitempty"`

				ConfigMapRuleSets *[]struct {
					ConfigMapRef *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

					DataMapKeys *[]string `tfsdk:"data_map_keys" yaml:"dataMapKeys,omitempty"`
				} `tfsdk:"config_map_rule_sets" yaml:"configMapRuleSets,omitempty"`

				CoreRuleSet *struct {
					CustomSettingsFile *string `tfsdk:"custom_settings_file" yaml:"customSettingsFile,omitempty"`

					CustomSettingsString *string `tfsdk:"custom_settings_string" yaml:"customSettingsString,omitempty"`
				} `tfsdk:"core_rule_set" yaml:"coreRuleSet,omitempty"`

				CustomInterventionMessage *string `tfsdk:"custom_intervention_message" yaml:"customInterventionMessage,omitempty"`

				Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`

				RequestHeadersOnly *bool `tfsdk:"request_headers_only" yaml:"requestHeadersOnly,omitempty"`

				ResponseHeadersOnly *bool `tfsdk:"response_headers_only" yaml:"responseHeadersOnly,omitempty"`

				RuleSets *[]struct {
					Directory *string `tfsdk:"directory" yaml:"directory,omitempty"`

					Files *[]string `tfsdk:"files" yaml:"files,omitempty"`

					RuleStr *string `tfsdk:"rule_str" yaml:"ruleStr,omitempty"`
				} `tfsdk:"rule_sets" yaml:"ruleSets,omitempty"`
			} `tfsdk:"waf" yaml:"waf,omitempty"`
		} `tfsdk:"options" yaml:"options,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGatewaySoloIoVirtualHostOptionV1Resource() resource.Resource {
	return &GatewaySoloIoVirtualHostOptionV1Resource{}
}

func (r *GatewaySoloIoVirtualHostOptionV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_solo_io_virtual_host_option_v1"
}

func (r *GatewaySoloIoVirtualHostOptionV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
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
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"namespaced_statuses": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"statuses": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"options": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"buffer_per_route": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"buffer": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max_request_bytes": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(4.294967295e+09),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disabled": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cors": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"allow_credentials": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allow_headers": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allow_methods": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allow_origin": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allow_origin_regex": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disable_for_route": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"expose_headers": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_age": {
										Description:         "",
										MarkdownDescription: "",

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

							"csrf": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"additional_origins": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ignore_case": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"safe_regex": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"google_re2": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"max_program_size": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	int64validator.AtLeast(0),

																	int64validator.AtMost(4.294967295e+09),
																},
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"regex": {
														Description:         "",
														MarkdownDescription: "",

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

											"suffix": {
												Description:         "",
												MarkdownDescription: "",

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

									"filter_enabled": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"default_value": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"denominator": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"numerator": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"runtime_key": {
												Description:         "",
												MarkdownDescription: "",

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

									"shadow_enabled": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"default_value": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"denominator": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"numerator": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"runtime_key": {
												Description:         "",
												MarkdownDescription: "",

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

							"dlp": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"actions": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"action_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"custom_action": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"mask_char": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"percent": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"value": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.DynamicNumberType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"regex": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"regex_actions": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"regex": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"subgroup": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

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

											"key_value_action": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key_to_mask": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"mask_char": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"percent": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"value": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.DynamicNumberType{},

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

											"shadow": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled_for": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.IntOrStringType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"extauth": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_ref": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

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

									"custom_auth": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"context_extensions": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

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

									"disable": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"extensions": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configs": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"header_manipulation": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"request_headers_to_add": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"append": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"header": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "",
														MarkdownDescription: "",

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

											"header_secret_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "",
														MarkdownDescription: "",

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

									"request_headers_to_remove": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"response_headers_to_add": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"append": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"header": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "",
														MarkdownDescription: "",

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

									"response_headers_to_remove": {
										Description:         "",
										MarkdownDescription: "",

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

							"include_attempt_count_in_response": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"include_request_attempt_count": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"jwt": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"allow_missing_or_failed_jwt": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"providers": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"audiences": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"claims_to_headers": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"append": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"claim": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"header": {
														Description:         "",
														MarkdownDescription: "",

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

											"issuer": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"jwks": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"local": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

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

													"remote": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"async_fetch": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"fast_listener": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"cache_duration": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"upstream_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"namespace": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"url": {
																Description:         "",
																MarkdownDescription: "",

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

											"keep_token": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"token_source": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"header": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"prefix": {
																Description:         "",
																MarkdownDescription: "",

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

													"query_params": {
														Description:         "",
														MarkdownDescription: "",

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

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"jwt_staged": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"after_ext_auth": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allow_missing_or_failed_jwt": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"providers": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"audiences": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"claims_to_headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"append": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"claim": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header": {
																Description:         "",
																MarkdownDescription: "",

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

													"issuer": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"jwks": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"local": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"remote": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"async_fetch": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"fast_listener": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"cache_duration": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"upstream_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"namespace": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"url": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"keep_token": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_source": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"header": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prefix": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"query_params": {
																Description:         "",
																MarkdownDescription: "",

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

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"before_ext_auth": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allow_missing_or_failed_jwt": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"providers": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"audiences": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"claims_to_headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"append": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"claim": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header": {
																Description:         "",
																MarkdownDescription: "",

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

													"issuer": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"jwks": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"local": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"remote": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"async_fetch": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"fast_listener": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"cache_duration": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"upstream_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"namespace": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"url": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"keep_token": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_source": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"header": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prefix": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"query_params": {
																Description:         "",
																MarkdownDescription: "",

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

							"rate_limit_configs": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"refs": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

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

							"rate_limit_early_configs": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"refs": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

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

							"rate_limit_regular_configs": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"refs": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

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

							"ratelimit": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"rate_limits": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"actions": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"destination_cluster": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"generic_key": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_value": {
																Description:         "",
																MarkdownDescription: "",

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

													"header_value_match": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"expect_match": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"exact_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"invert_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prefix_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"present_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"range_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"end": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"start": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"suffix_match": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"metadata": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"default_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"descriptor_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"metadata_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"source": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"remote_address": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"request_headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header_name": {
																Description:         "",
																MarkdownDescription: "",

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

													"source_cluster": {
														Description:         "",
														MarkdownDescription: "",

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

											"set_actions": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"destination_cluster": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"generic_key": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_value": {
																Description:         "",
																MarkdownDescription: "",

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

													"header_value_match": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"expect_match": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"exact_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"invert_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prefix_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"present_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"range_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"end": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"start": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"suffix_match": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"metadata": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"default_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"descriptor_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"metadata_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"source": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"remote_address": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"request_headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header_name": {
																Description:         "",
																MarkdownDescription: "",

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

													"source_cluster": {
														Description:         "",
														MarkdownDescription: "",

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

							"ratelimit_basic": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"anonymous_limits": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"requests_per_unit": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"unit": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authorized_limits": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"requests_per_unit": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"unit": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

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

							"ratelimit_early": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"rate_limits": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"actions": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"destination_cluster": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"generic_key": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_value": {
																Description:         "",
																MarkdownDescription: "",

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

													"header_value_match": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"expect_match": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"exact_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"invert_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prefix_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"present_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"range_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"end": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"start": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"suffix_match": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"metadata": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"default_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"descriptor_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"metadata_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"source": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"remote_address": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"request_headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header_name": {
																Description:         "",
																MarkdownDescription: "",

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

													"source_cluster": {
														Description:         "",
														MarkdownDescription: "",

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

											"set_actions": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"destination_cluster": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"generic_key": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_value": {
																Description:         "",
																MarkdownDescription: "",

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

													"header_value_match": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"expect_match": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"exact_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"invert_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prefix_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"present_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"range_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"end": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"start": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"suffix_match": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"metadata": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"default_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"descriptor_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"metadata_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"source": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"remote_address": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"request_headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header_name": {
																Description:         "",
																MarkdownDescription: "",

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

													"source_cluster": {
														Description:         "",
														MarkdownDescription: "",

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

							"ratelimit_regular": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"rate_limits": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"actions": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"destination_cluster": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"generic_key": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_value": {
																Description:         "",
																MarkdownDescription: "",

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

													"header_value_match": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"expect_match": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"exact_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"invert_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prefix_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"present_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"range_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"end": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"start": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"suffix_match": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"metadata": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"default_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"descriptor_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"metadata_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"source": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"remote_address": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"request_headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header_name": {
																Description:         "",
																MarkdownDescription: "",

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

													"source_cluster": {
														Description:         "",
														MarkdownDescription: "",

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

											"set_actions": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"destination_cluster": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"generic_key": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_value": {
																Description:         "",
																MarkdownDescription: "",

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

													"header_value_match": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"expect_match": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"exact_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"invert_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"prefix_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"present_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"range_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"end": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"start": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"suffix_match": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"metadata": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"default_value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"descriptor_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"metadata_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

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

															"source": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"remote_address": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"request_headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"descriptor_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header_name": {
																Description:         "",
																MarkdownDescription: "",

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

													"source_cluster": {
														Description:         "",
														MarkdownDescription: "",

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

							"rbac": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"disable": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"policies": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"nested_claim_delimiter": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"permissions": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"methods": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path_prefix": {
														Description:         "",
														MarkdownDescription: "",

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

											"principals": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"jwt_principal": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"claims": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"matcher": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"provider": {
																Description:         "",
																MarkdownDescription: "",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"retries": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"num_retries": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"per_try_timeout": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"retry_on": {
										Description:         "",
										MarkdownDescription: "",

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

							"staged_transformations": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"early": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"request_transforms": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"clear_route_cache": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"matcher": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"case_sensitive": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"exact": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"invert_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"methods": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"prefix": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"query_parameters": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"regex": {
																Description:         "",
																MarkdownDescription: "",

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

													"request_transformation": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"header_body_transform": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"add_request_metadata": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"transformation_template": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"advanced_templates": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"dynamic_metadata_values": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"metadata_namespace": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"extractors": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"header": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"subgroup": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"headers_to_append": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"headers_to_remove": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"ignore_error_on_parse": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"merge_extractors_to_body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parse_body_behavior": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"passthrough": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"xslt_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"non_xml_transform": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"set_content_type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"xslt": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"response_transformation": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"header_body_transform": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"add_request_metadata": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"transformation_template": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"advanced_templates": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"dynamic_metadata_values": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"metadata_namespace": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"extractors": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"header": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"subgroup": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"headers_to_append": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"headers_to_remove": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"ignore_error_on_parse": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"merge_extractors_to_body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parse_body_behavior": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"passthrough": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"xslt_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"non_xml_transform": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"set_content_type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"xslt": {
																		Description:         "",
																		MarkdownDescription: "",

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

											"response_transforms": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"matchers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"invert_match": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

													"response_code_details": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"response_transformation": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"header_body_transform": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"add_request_metadata": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"transformation_template": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"advanced_templates": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"dynamic_metadata_values": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"metadata_namespace": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"extractors": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"header": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"subgroup": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"headers_to_append": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"headers_to_remove": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"ignore_error_on_parse": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"merge_extractors_to_body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parse_body_behavior": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"passthrough": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"xslt_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"non_xml_transform": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"set_content_type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"xslt": {
																		Description:         "",
																		MarkdownDescription: "",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"inherit_transformation": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"regular": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"request_transforms": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"clear_route_cache": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"matcher": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"case_sensitive": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"exact": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"headers": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"invert_match": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"methods": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"prefix": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"query_parameters": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"regex": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"regex": {
																Description:         "",
																MarkdownDescription: "",

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

													"request_transformation": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"header_body_transform": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"add_request_metadata": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"transformation_template": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"advanced_templates": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"dynamic_metadata_values": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"metadata_namespace": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"extractors": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"header": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"subgroup": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"headers_to_append": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"headers_to_remove": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"ignore_error_on_parse": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"merge_extractors_to_body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parse_body_behavior": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"passthrough": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"xslt_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"non_xml_transform": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"set_content_type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"xslt": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"response_transformation": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"header_body_transform": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"add_request_metadata": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"transformation_template": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"advanced_templates": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"dynamic_metadata_values": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"metadata_namespace": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"extractors": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"header": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"subgroup": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"headers_to_append": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"headers_to_remove": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"ignore_error_on_parse": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"merge_extractors_to_body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parse_body_behavior": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"passthrough": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"xslt_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"non_xml_transform": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"set_content_type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"xslt": {
																		Description:         "",
																		MarkdownDescription: "",

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

											"response_transforms": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"matchers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"invert_match": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

													"response_code_details": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"response_transformation": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"header_body_transform": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"add_request_metadata": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"transformation_template": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"advanced_templates": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"dynamic_metadata_values": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"metadata_namespace": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"extractors": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"body": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"header": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"regex": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"subgroup": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"text": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"headers_to_append": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"text": {
																						Description:         "",
																						MarkdownDescription: "",

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

																	"headers_to_remove": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"ignore_error_on_parse": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"merge_extractors_to_body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"parse_body_behavior": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"passthrough": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"xslt_transformation": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"non_xml_transform": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"set_content_type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"xslt": {
																		Description:         "",
																		MarkdownDescription: "",

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

							"stats": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"virtual_clusters": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"method": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pattern": {
												Description:         "",
												MarkdownDescription: "",

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

							"transformations": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"clear_route_cache": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"request_transformation": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"header_body_transform": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"add_request_metadata": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"transformation_template": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"advanced_templates": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"body": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"text": {
																Description:         "",
																MarkdownDescription: "",

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

													"dynamic_metadata_values": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"metadata_namespace": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"extractors": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"body": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"subgroup": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"text": {
																Description:         "",
																MarkdownDescription: "",

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

													"headers_to_append": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"headers_to_remove": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ignore_error_on_parse": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"merge_extractors_to_body": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"parse_body_behavior": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"passthrough": {
														Description:         "",
														MarkdownDescription: "",

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

											"xslt_transformation": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"non_xml_transform": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"set_content_type": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"xslt": {
														Description:         "",
														MarkdownDescription: "",

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

									"response_transformation": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"header_body_transform": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"add_request_metadata": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"transformation_template": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"advanced_templates": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"body": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"text": {
																Description:         "",
																MarkdownDescription: "",

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

													"dynamic_metadata_values": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"metadata_namespace": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"extractors": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"body": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"subgroup": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"text": {
																Description:         "",
																MarkdownDescription: "",

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

													"headers_to_append": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
																		Description:         "",
																		MarkdownDescription: "",

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

													"headers_to_remove": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ignore_error_on_parse": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"merge_extractors_to_body": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"parse_body_behavior": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"passthrough": {
														Description:         "",
														MarkdownDescription: "",

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

											"xslt_transformation": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"non_xml_transform": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"set_content_type": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"xslt": {
														Description:         "",
														MarkdownDescription: "",

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

							"waf": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"audit_logging": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"location": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"config_map_rule_sets": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "",
														MarkdownDescription: "",

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

											"data_map_keys": {
												Description:         "",
												MarkdownDescription: "",

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

									"core_rule_set": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"custom_settings_file": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"custom_settings_string": {
												Description:         "",
												MarkdownDescription: "",

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

									"custom_intervention_message": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disabled": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"request_headers_only": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"response_headers_only": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rule_sets": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"directory": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"files": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rule_str": {
												Description:         "",
												MarkdownDescription: "",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *GatewaySoloIoVirtualHostOptionV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_gateway_solo_io_virtual_host_option_v1")

	var state GatewaySoloIoVirtualHostOptionV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GatewaySoloIoVirtualHostOptionV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gateway.solo.io/v1")
	goModel.Kind = utilities.Ptr("VirtualHostOption")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GatewaySoloIoVirtualHostOptionV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_solo_io_virtual_host_option_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *GatewaySoloIoVirtualHostOptionV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_gateway_solo_io_virtual_host_option_v1")

	var state GatewaySoloIoVirtualHostOptionV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GatewaySoloIoVirtualHostOptionV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gateway.solo.io/v1")
	goModel.Kind = utilities.Ptr("VirtualHostOption")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GatewaySoloIoVirtualHostOptionV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_gateway_solo_io_virtual_host_option_v1")
	// NO-OP: Terraform removes the state automatically for us
}
