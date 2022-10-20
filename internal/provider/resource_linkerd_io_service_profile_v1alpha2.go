/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type LinkerdIoServiceProfileV1Alpha2Resource struct{}

var (
	_ resource.Resource = (*LinkerdIoServiceProfileV1Alpha2Resource)(nil)
)

type LinkerdIoServiceProfileV1Alpha2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type LinkerdIoServiceProfileV1Alpha2GoModel struct {
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
		DstOverrides *[]struct {
			Authority *string `tfsdk:"authority" yaml:"authority,omitempty"`

			Weight utilities.IntOrString `tfsdk:"weight" yaml:"weight,omitempty"`
		} `tfsdk:"dst_overrides" yaml:"dstOverrides,omitempty"`

		OpaquePorts *[]string `tfsdk:"opaque_ports" yaml:"opaquePorts,omitempty"`

		RetryBudget *struct {
			MinRetriesPerSecond *int64 `tfsdk:"min_retries_per_second" yaml:"minRetriesPerSecond,omitempty"`

			RetryRatio *float64 `tfsdk:"retry_ratio" yaml:"retryRatio,omitempty"`

			Ttl *string `tfsdk:"ttl" yaml:"ttl,omitempty"`
		} `tfsdk:"retry_budget" yaml:"retryBudget,omitempty"`

		Routes *[]struct {
			Condition *struct {
				All *[]map[string]string `tfsdk:"all" yaml:"all,omitempty"`

				Any *[]map[string]string `tfsdk:"any" yaml:"any,omitempty"`

				Method *string `tfsdk:"method" yaml:"method,omitempty"`

				Not *[]map[string]string `tfsdk:"not" yaml:"not,omitempty"`

				PathRegex *string `tfsdk:"path_regex" yaml:"pathRegex,omitempty"`
			} `tfsdk:"condition" yaml:"condition,omitempty"`

			IsRetryable *bool `tfsdk:"is_retryable" yaml:"isRetryable,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			ResponseClasses *[]struct {
				Condition *struct {
					All *[]map[string]string `tfsdk:"all" yaml:"all,omitempty"`

					Any *[]map[string]string `tfsdk:"any" yaml:"any,omitempty"`

					Not *[]map[string]string `tfsdk:"not" yaml:"not,omitempty"`

					Status *struct {
						Max *int64 `tfsdk:"max" yaml:"max,omitempty"`

						Min *int64 `tfsdk:"min" yaml:"min,omitempty"`
					} `tfsdk:"status" yaml:"status,omitempty"`
				} `tfsdk:"condition" yaml:"condition,omitempty"`

				IsFailure *bool `tfsdk:"is_failure" yaml:"isFailure,omitempty"`
			} `tfsdk:"response_classes" yaml:"responseClasses,omitempty"`

			Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
		} `tfsdk:"routes" yaml:"routes,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewLinkerdIoServiceProfileV1Alpha2Resource() resource.Resource {
	return &LinkerdIoServiceProfileV1Alpha2Resource{}
}

func (r *LinkerdIoServiceProfileV1Alpha2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_linkerd_io_service_profile_v1alpha2"
}

func (r *LinkerdIoServiceProfileV1Alpha2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				Description:         "Spec is the custom resource spec",
				MarkdownDescription: "Spec is the custom resource spec",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"dst_overrides": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"authority": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"weight": {
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

					"opaque_ports": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"retry_budget": {
						Description:         "RetryBudget describes the maximum number of retries that should be issued to this service.",
						MarkdownDescription: "RetryBudget describes the maximum number of retries that should be issued to this service.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"min_retries_per_second": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"retry_ratio": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Float64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"ttl": {
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

					"routes": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"condition": {
								Description:         "RequestMatch describes the conditions under which to match a Route.",
								MarkdownDescription: "RequestMatch describes the conditions under which to match a Route.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"all": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"any": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"method": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"not": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path_regex": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"is_retryable": {
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

								Required: true,
								Optional: false,
								Computed: false,
							},

							"response_classes": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"condition": {
										Description:         "ResponseMatch describes the conditions under which to classify a response.",
										MarkdownDescription: "ResponseMatch describes the conditions under which to classify a response.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"all": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"any": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"not": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"status": {
												Description:         "Range describes a range of integers (e.g. status codes).",
												MarkdownDescription: "Range describes a range of integers (e.g. status codes).",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"max": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"min": {
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

										Required: true,
										Optional: false,
										Computed: false,
									},

									"is_failure": {
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

							"timeout": {
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
		},
	}, nil
}

func (r *LinkerdIoServiceProfileV1Alpha2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_linkerd_io_service_profile_v1alpha2")

	var state LinkerdIoServiceProfileV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LinkerdIoServiceProfileV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("linkerd.io/v1alpha2")
	goModel.Kind = utilities.Ptr("ServiceProfile")

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

func (r *LinkerdIoServiceProfileV1Alpha2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_linkerd_io_service_profile_v1alpha2")
	// NO-OP: All data is already in Terraform state
}

func (r *LinkerdIoServiceProfileV1Alpha2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_linkerd_io_service_profile_v1alpha2")

	var state LinkerdIoServiceProfileV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LinkerdIoServiceProfileV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("linkerd.io/v1alpha2")
	goModel.Kind = utilities.Ptr("ServiceProfile")

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

func (r *LinkerdIoServiceProfileV1Alpha2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_linkerd_io_service_profile_v1alpha2")
	// NO-OP: Terraform removes the state automatically for us
}
