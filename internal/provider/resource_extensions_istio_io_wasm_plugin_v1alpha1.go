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

type ExtensionsIstioIoWasmPluginV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ExtensionsIstioIoWasmPluginV1Alpha1Resource)(nil)
)

type ExtensionsIstioIoWasmPluginV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ExtensionsIstioIoWasmPluginV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

		ImagePullSecret *string `tfsdk:"image_pull_secret" yaml:"imagePullSecret,omitempty"`

		Match *[]struct {
			Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

			Ports *[]struct {
				Number *int64 `tfsdk:"number" yaml:"number,omitempty"`
			} `tfsdk:"ports" yaml:"ports,omitempty"`
		} `tfsdk:"match" yaml:"match,omitempty"`

		Phase *string `tfsdk:"phase" yaml:"phase,omitempty"`

		PluginConfig *map[string]string `tfsdk:"plugin_config" yaml:"pluginConfig,omitempty"`

		PluginName *string `tfsdk:"plugin_name" yaml:"pluginName,omitempty"`

		Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

		Selector *struct {
			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"selector" yaml:"selector,omitempty"`

		Sha256 *string `tfsdk:"sha256" yaml:"sha256,omitempty"`

		Url *string `tfsdk:"url" yaml:"url,omitempty"`

		VerificationKey *string `tfsdk:"verification_key" yaml:"verificationKey,omitempty"`

		VmConfig *struct {
			Env *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *string `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"env" yaml:"env,omitempty"`
		} `tfsdk:"vm_config" yaml:"vmConfig,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewExtensionsIstioIoWasmPluginV1Alpha1Resource() resource.Resource {
	return &ExtensionsIstioIoWasmPluginV1Alpha1Resource{}
}

func (r *ExtensionsIstioIoWasmPluginV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_extensions_istio_io_wasm_plugin_v1alpha1"
}

func (r *ExtensionsIstioIoWasmPluginV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html",
				MarkdownDescription: "Extend the functionality provided by the Istio proxy through WebAssembly filters. See more details at: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"image_pull_policy": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_pull_secret": {
						Description:         "Credentials to use for OCI image pulling.",
						MarkdownDescription: "Credentials to use for OCI image pulling.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"match": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"mode": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ports": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"number": {
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

					"phase": {
						Description:         "Determines where in the filter chain this 'WasmPlugin' is to be injected.",
						MarkdownDescription: "Determines where in the filter chain this 'WasmPlugin' is to be injected.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"plugin_config": {
						Description:         "The configuration that will be passed on to the plugin.",
						MarkdownDescription: "The configuration that will be passed on to the plugin.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"plugin_name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"priority": {
						Description:         "Determines ordering of 'WasmPlugins' in the same 'phase'.",
						MarkdownDescription: "Determines ordering of 'WasmPlugins' in the same 'phase'.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_labels": {
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

					"sha256": {
						Description:         "SHA256 checksum that will be used to verify Wasm module or OCI container.",
						MarkdownDescription: "SHA256 checksum that will be used to verify Wasm module or OCI container.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"url": {
						Description:         "URL of a Wasm module or OCI container.",
						MarkdownDescription: "URL of a Wasm module or OCI container.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"verification_key": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vm_config": {
						Description:         "Configuration for a Wasm VM.",
						MarkdownDescription: "Configuration for a Wasm VM.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"env": {
								Description:         "Specifies environment variables to be injected to this VM.",
								MarkdownDescription: "Specifies environment variables to be injected to this VM.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Value for the environment variable.",
										MarkdownDescription: "Value for the environment variable.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
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
		},
	}, nil
}

func (r *ExtensionsIstioIoWasmPluginV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_extensions_istio_io_wasm_plugin_v1alpha1")

	var state ExtensionsIstioIoWasmPluginV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ExtensionsIstioIoWasmPluginV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("extensions.istio.io/v1alpha1")
	goModel.Kind = utilities.Ptr("WasmPlugin")

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

func (r *ExtensionsIstioIoWasmPluginV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_extensions_istio_io_wasm_plugin_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ExtensionsIstioIoWasmPluginV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_extensions_istio_io_wasm_plugin_v1alpha1")

	var state ExtensionsIstioIoWasmPluginV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ExtensionsIstioIoWasmPluginV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("extensions.istio.io/v1alpha1")
	goModel.Kind = utilities.Ptr("WasmPlugin")

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

func (r *ExtensionsIstioIoWasmPluginV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_extensions_istio_io_wasm_plugin_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
