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

type DevicesKubeedgeIoDeviceModelV1Alpha2Resource struct{}

var (
	_ resource.Resource = (*DevicesKubeedgeIoDeviceModelV1Alpha2Resource)(nil)
)

type DevicesKubeedgeIoDeviceModelV1Alpha2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type DevicesKubeedgeIoDeviceModelV1Alpha2GoModel struct {
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
		Properties *[]struct {
			Description *string `tfsdk:"description" yaml:"description,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Type *struct {
				Boolean *struct {
					AccessMode *string `tfsdk:"access_mode" yaml:"accessMode,omitempty"`

					DefaultValue *bool `tfsdk:"default_value" yaml:"defaultValue,omitempty"`
				} `tfsdk:"boolean" yaml:"boolean,omitempty"`

				Bytes *struct {
					AccessMode *string `tfsdk:"access_mode" yaml:"accessMode,omitempty"`
				} `tfsdk:"bytes" yaml:"bytes,omitempty"`

				Double *struct {
					AccessMode *string `tfsdk:"access_mode" yaml:"accessMode,omitempty"`

					DefaultValue utilities.DynamicNumber `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

					Maximum utilities.DynamicNumber `tfsdk:"maximum" yaml:"maximum,omitempty"`

					Minimum utilities.DynamicNumber `tfsdk:"minimum" yaml:"minimum,omitempty"`

					Unit *string `tfsdk:"unit" yaml:"unit,omitempty"`
				} `tfsdk:"double" yaml:"double,omitempty"`

				Float *struct {
					AccessMode *string `tfsdk:"access_mode" yaml:"accessMode,omitempty"`

					DefaultValue utilities.DynamicNumber `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

					Maximum utilities.DynamicNumber `tfsdk:"maximum" yaml:"maximum,omitempty"`

					Minimum utilities.DynamicNumber `tfsdk:"minimum" yaml:"minimum,omitempty"`

					Unit *string `tfsdk:"unit" yaml:"unit,omitempty"`
				} `tfsdk:"float" yaml:"float,omitempty"`

				Int *struct {
					AccessMode *string `tfsdk:"access_mode" yaml:"accessMode,omitempty"`

					DefaultValue *int64 `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

					Maximum *int64 `tfsdk:"maximum" yaml:"maximum,omitempty"`

					Minimum *int64 `tfsdk:"minimum" yaml:"minimum,omitempty"`

					Unit *string `tfsdk:"unit" yaml:"unit,omitempty"`
				} `tfsdk:"int" yaml:"int,omitempty"`

				String *struct {
					AccessMode *string `tfsdk:"access_mode" yaml:"accessMode,omitempty"`

					DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`
				} `tfsdk:"string" yaml:"string,omitempty"`
			} `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"properties" yaml:"properties,omitempty"`

		Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewDevicesKubeedgeIoDeviceModelV1Alpha2Resource() resource.Resource {
	return &DevicesKubeedgeIoDeviceModelV1Alpha2Resource{}
}

func (r *DevicesKubeedgeIoDeviceModelV1Alpha2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_kubeedge_io_device_model_v1alpha2"
}

func (r *DevicesKubeedgeIoDeviceModelV1Alpha2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "DeviceModel is the Schema for the device model API",
		MarkdownDescription: "DeviceModel is the Schema for the device model API",
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
				Description:         "DeviceModelSpec defines the model / template for a device.It is a blueprint which describes the device capabilities and access mechanism via property visitors.",
				MarkdownDescription: "DeviceModelSpec defines the model / template for a device.It is a blueprint which describes the device capabilities and access mechanism via property visitors.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"properties": {
						Description:         "Required: List of device properties.",
						MarkdownDescription: "Required: List of device properties.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"description": {
								Description:         "The device property description.",
								MarkdownDescription: "The device property description.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Required: The device property name.",
								MarkdownDescription: "Required: The device property name.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "Required: PropertyType represents the type and data validation of the property.",
								MarkdownDescription: "Required: PropertyType represents the type and data validation of the property.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"boolean": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"access_mode": {
												Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
												MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("ReadWrite", "ReadOnly"),
												},
											},

											"default_value": {
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

									"bytes": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"access_mode": {
												Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
												MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("ReadWrite", "ReadOnly"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"double": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"access_mode": {
												Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
												MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("ReadWrite", "ReadOnly"),
												},
											},

											"default_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicNumberType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"maximum": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicNumberType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"minimum": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicNumberType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"unit": {
												Description:         "The unit of the property",
												MarkdownDescription: "The unit of the property",

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

									"float": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"access_mode": {
												Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
												MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("ReadWrite", "ReadOnly"),
												},
											},

											"default_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicNumberType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"maximum": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicNumberType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"minimum": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicNumberType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"unit": {
												Description:         "The unit of the property",
												MarkdownDescription: "The unit of the property",

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

									"int": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"access_mode": {
												Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
												MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("ReadWrite", "ReadOnly"),
												},
											},

											"default_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"maximum": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"minimum": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"unit": {
												Description:         "The unit of the property",
												MarkdownDescription: "The unit of the property",

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

									"string": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"access_mode": {
												Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
												MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("ReadWrite", "ReadOnly"),
												},
											},

											"default_value": {
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

					"protocol": {
						Description:         "Required for DMI: Protocol name used by the device.",
						MarkdownDescription: "Required for DMI: Protocol name used by the device.",

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

func (r *DevicesKubeedgeIoDeviceModelV1Alpha2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_devices_kubeedge_io_device_model_v1alpha2")

	var state DevicesKubeedgeIoDeviceModelV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DevicesKubeedgeIoDeviceModelV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("devices.kubeedge.io/v1alpha2")
	goModel.Kind = utilities.Ptr("DeviceModel")

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

func (r *DevicesKubeedgeIoDeviceModelV1Alpha2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_devices_kubeedge_io_device_model_v1alpha2")
	// NO-OP: All data is already in Terraform state
}

func (r *DevicesKubeedgeIoDeviceModelV1Alpha2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_devices_kubeedge_io_device_model_v1alpha2")

	var state DevicesKubeedgeIoDeviceModelV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DevicesKubeedgeIoDeviceModelV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("devices.kubeedge.io/v1alpha2")
	goModel.Kind = utilities.Ptr("DeviceModel")

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

func (r *DevicesKubeedgeIoDeviceModelV1Alpha2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_devices_kubeedge_io_device_model_v1alpha2")
	// NO-OP: Terraform removes the state automatically for us
}