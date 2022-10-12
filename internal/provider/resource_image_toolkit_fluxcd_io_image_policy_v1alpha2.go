/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type ImageToolkitFluxcdIoImagePolicyV1Alpha2Resource struct{}

var (
	_ resource.Resource = (*ImageToolkitFluxcdIoImagePolicyV1Alpha2Resource)(nil)
)

type ImageToolkitFluxcdIoImagePolicyV1Alpha2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ImageToolkitFluxcdIoImagePolicyV1Alpha2GoModel struct {
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
		FilterTags *struct {
			Extract *string `tfsdk:"extract" yaml:"extract,omitempty"`

			Pattern *string `tfsdk:"pattern" yaml:"pattern,omitempty"`
		} `tfsdk:"filter_tags" yaml:"filterTags,omitempty"`

		ImageRepositoryRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"image_repository_ref" yaml:"imageRepositoryRef,omitempty"`

		Policy *struct {
			Alphabetical *struct {
				Order *string `tfsdk:"order" yaml:"order,omitempty"`
			} `tfsdk:"alphabetical" yaml:"alphabetical,omitempty"`

			Numerical *struct {
				Order *string `tfsdk:"order" yaml:"order,omitempty"`
			} `tfsdk:"numerical" yaml:"numerical,omitempty"`

			Semver *struct {
				Range *string `tfsdk:"range" yaml:"range,omitempty"`
			} `tfsdk:"semver" yaml:"semver,omitempty"`
		} `tfsdk:"policy" yaml:"policy,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewImageToolkitFluxcdIoImagePolicyV1Alpha2Resource() resource.Resource {
	return &ImageToolkitFluxcdIoImagePolicyV1Alpha2Resource{}
}

func (r *ImageToolkitFluxcdIoImagePolicyV1Alpha2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_image_toolkit_fluxcd_io_image_policy_v1alpha2"
}

func (r *ImageToolkitFluxcdIoImagePolicyV1Alpha2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ImagePolicy is the Schema for the imagepolicies API",
		MarkdownDescription: "ImagePolicy is the Schema for the imagepolicies API",
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
				Description:         "ImagePolicySpec defines the parameters for calculating the ImagePolicy",
				MarkdownDescription: "ImagePolicySpec defines the parameters for calculating the ImagePolicy",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"filter_tags": {
						Description:         "FilterTags enables filtering for only a subset of tags based on a set of rules. If no rules are provided, all the tags from the repository will be ordered and compared.",
						MarkdownDescription: "FilterTags enables filtering for only a subset of tags based on a set of rules. If no rules are provided, all the tags from the repository will be ordered and compared.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"extract": {
								Description:         "Extract allows a capture group to be extracted from the specified regular expression pattern, useful before tag evaluation.",
								MarkdownDescription: "Extract allows a capture group to be extracted from the specified regular expression pattern, useful before tag evaluation.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pattern": {
								Description:         "Pattern specifies a regular expression pattern used to filter for image tags.",
								MarkdownDescription: "Pattern specifies a regular expression pattern used to filter for image tags.",

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

					"image_repository_ref": {
						Description:         "ImageRepositoryRef points at the object specifying the image being scanned",
						MarkdownDescription: "ImageRepositoryRef points at the object specifying the image being scanned",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

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

					"policy": {
						Description:         "Policy gives the particulars of the policy to be followed in selecting the most recent image",
						MarkdownDescription: "Policy gives the particulars of the policy to be followed in selecting the most recent image",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"alphabetical": {
								Description:         "Alphabetical set of rules to use for alphabetical ordering of the tags.",
								MarkdownDescription: "Alphabetical set of rules to use for alphabetical ordering of the tags.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"order": {
										Description:         "Order specifies the sorting order of the tags. Given the letters of the alphabet as tags, ascending order would select Z, and descending order would select A.",
										MarkdownDescription: "Order specifies the sorting order of the tags. Given the letters of the alphabet as tags, ascending order would select Z, and descending order would select A.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("asc", "desc"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"numerical": {
								Description:         "Numerical set of rules to use for numerical ordering of the tags.",
								MarkdownDescription: "Numerical set of rules to use for numerical ordering of the tags.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"order": {
										Description:         "Order specifies the sorting order of the tags. Given the integer values from 0 to 9 as tags, ascending order would select 9, and descending order would select 0.",
										MarkdownDescription: "Order specifies the sorting order of the tags. Given the integer values from 0 to 9 as tags, ascending order would select 9, and descending order would select 0.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("asc", "desc"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"semver": {
								Description:         "SemVer gives a semantic version range to check against the tags available.",
								MarkdownDescription: "SemVer gives a semantic version range to check against the tags available.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"range": {
										Description:         "Range gives a semver range for the image tag; the highest version within the range that's a tag yields the latest image.",
										MarkdownDescription: "Range gives a semver range for the image tag; the highest version within the range that's a tag yields the latest image.",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *ImageToolkitFluxcdIoImagePolicyV1Alpha2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_image_toolkit_fluxcd_io_image_policy_v1alpha2")

	var state ImageToolkitFluxcdIoImagePolicyV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ImageToolkitFluxcdIoImagePolicyV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("image.toolkit.fluxcd.io/v1alpha2")
	goModel.Kind = utilities.Ptr("ImagePolicy")

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

func (r *ImageToolkitFluxcdIoImagePolicyV1Alpha2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_image_toolkit_fluxcd_io_image_policy_v1alpha2")
	// NO-OP: All data is already in Terraform state
}

func (r *ImageToolkitFluxcdIoImagePolicyV1Alpha2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_image_toolkit_fluxcd_io_image_policy_v1alpha2")

	var state ImageToolkitFluxcdIoImagePolicyV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ImageToolkitFluxcdIoImagePolicyV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("image.toolkit.fluxcd.io/v1alpha2")
	goModel.Kind = utilities.Ptr("ImagePolicy")

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

func (r *ImageToolkitFluxcdIoImagePolicyV1Alpha2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_image_toolkit_fluxcd_io_image_policy_v1alpha2")
	// NO-OP: Terraform removes the state automatically for us
}
