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

type GetambassadorIoLogServiceV3Alpha1Resource struct{}

var (
	_ resource.Resource = (*GetambassadorIoLogServiceV3Alpha1Resource)(nil)
)

type GetambassadorIoLogServiceV3Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GetambassadorIoLogServiceV3Alpha1GoModel struct {
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
		Ambassador_id *[]string `tfsdk:"ambassador_id" yaml:"ambassador_id,omitempty"`

		Driver *string `tfsdk:"driver" yaml:"driver,omitempty"`

		Driver_config *struct {
			Additional_log_headers *[]struct {
				During_request *bool `tfsdk:"during_request" yaml:"during_request,omitempty"`

				During_response *bool `tfsdk:"during_response" yaml:"during_response,omitempty"`

				During_trailer *bool `tfsdk:"during_trailer" yaml:"during_trailer,omitempty"`

				Header_name *string `tfsdk:"header_name" yaml:"header_name,omitempty"`
			} `tfsdk:"additional_log_headers" yaml:"additional_log_headers,omitempty"`
		} `tfsdk:"driver_config" yaml:"driver_config,omitempty"`

		Flush_interval_byte_size *int64 `tfsdk:"flush_interval_byte_size" yaml:"flush_interval_byte_size,omitempty"`

		Flush_interval_time *int64 `tfsdk:"flush_interval_time" yaml:"flush_interval_time,omitempty"`

		Grpc *bool `tfsdk:"grpc" yaml:"grpc,omitempty"`

		Protocol_version *string `tfsdk:"protocol_version" yaml:"protocol_version,omitempty"`

		Service *string `tfsdk:"service" yaml:"service,omitempty"`

		Stats_name *string `tfsdk:"stats_name" yaml:"stats_name,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGetambassadorIoLogServiceV3Alpha1Resource() resource.Resource {
	return &GetambassadorIoLogServiceV3Alpha1Resource{}
}

func (r *GetambassadorIoLogServiceV3Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_getambassador_io_log_service_v3alpha1"
}

func (r *GetambassadorIoLogServiceV3Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "LogService is the Schema for the logservices API",
		MarkdownDescription: "LogService is the Schema for the logservices API",
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
				Description:         "LogServiceSpec defines the desired state of LogService",
				MarkdownDescription: "LogServiceSpec defines the desired state of LogService",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"ambassador_id": {
						Description:         "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  	ambassador_id: 	- 'default'  TODO(lukeshu): In v3alpha2, consider renaming all of the 'ambassador_id' (singular) fields to 'ambassador_ids' (plural).",
						MarkdownDescription: "AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  	ambassador_id: 	- 'default'  TODO(lukeshu): In v3alpha2, consider renaming all of the 'ambassador_id' (singular) fields to 'ambassador_ids' (plural).",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"driver": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("tcp", "http"),
						},
					},

					"driver_config": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"additional_log_headers": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"during_request": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"during_response": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"during_trailer": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"flush_interval_byte_size": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"flush_interval_time": {
						Description:         "TODO(lukeshu): In v3alpha2, change all of the '{foo}s'/'SecondDuration' fields to '{foo}'/'metav1.Duration'.",
						MarkdownDescription: "TODO(lukeshu): In v3alpha2, change all of the '{foo}s'/'SecondDuration' fields to '{foo}'/'metav1.Duration'.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"grpc": {
						Description:         "TODO(lukeshu): In v3alpha2, drop this LogService.spec.grpc.  Due to sloppy implementation it is required to be present, and required to be 'true'.  It is silly to have a required field with only one valid value, we should just remove the thing.",
						MarkdownDescription: "TODO(lukeshu): In v3alpha2, drop this LogService.spec.grpc.  Due to sloppy implementation it is required to be present, and required to be 'true'.  It is silly to have a required field with only one valid value, we should just remove the thing.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"protocol_version": {
						Description:         "ProtocolVersion is the envoy api transport protocol version",
						MarkdownDescription: "ProtocolVersion is the envoy api transport protocol version",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("v2", "v3"),
						},
					},

					"service": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"stats_name": {
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
		},
	}, nil
}

func (r *GetambassadorIoLogServiceV3Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_getambassador_io_log_service_v3alpha1")

	var state GetambassadorIoLogServiceV3Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GetambassadorIoLogServiceV3Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("getambassador.io/v3alpha1")
	goModel.Kind = utilities.Ptr("LogService")

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

func (r *GetambassadorIoLogServiceV3Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_log_service_v3alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *GetambassadorIoLogServiceV3Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_getambassador_io_log_service_v3alpha1")

	var state GetambassadorIoLogServiceV3Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GetambassadorIoLogServiceV3Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("getambassador.io/v3alpha1")
	goModel.Kind = utilities.Ptr("LogService")

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

func (r *GetambassadorIoLogServiceV3Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_getambassador_io_log_service_v3alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
