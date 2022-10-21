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

type ChaosMeshOrgAWSChaosV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ChaosMeshOrgAWSChaosV1Alpha1Resource)(nil)
)

type ChaosMeshOrgAWSChaosV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ChaosMeshOrgAWSChaosV1Alpha1GoModel struct {
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
		Action *string `tfsdk:"action" yaml:"action,omitempty"`

		AwsRegion *string `tfsdk:"aws_region" yaml:"awsRegion,omitempty"`

		DeviceName *string `tfsdk:"device_name" yaml:"deviceName,omitempty"`

		Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

		Ec2Instance *string `tfsdk:"ec2_instance" yaml:"ec2Instance,omitempty"`

		Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

		SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

		VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewChaosMeshOrgAWSChaosV1Alpha1Resource() resource.Resource {
	return &ChaosMeshOrgAWSChaosV1Alpha1Resource{}
}

func (r *ChaosMeshOrgAWSChaosV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chaos_mesh_org_aws_chaos_v1alpha1"
}

func (r *ChaosMeshOrgAWSChaosV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "AWSChaos is the Schema for the awschaos API",
		MarkdownDescription: "AWSChaos is the Schema for the awschaos API",
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
				Description:         "AWSChaosSpec is the content of the specification for an AWSChaos",
				MarkdownDescription: "AWSChaosSpec is the content of the specification for an AWSChaos",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"action": {
						Description:         "Action defines the specific aws chaos action. Supported action: ec2-stop / ec2-restart / detach-volume Default action: ec2-stop",
						MarkdownDescription: "Action defines the specific aws chaos action. Supported action: ec2-stop / ec2-restart / detach-volume Default action: ec2-stop",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("ec2-stop", "ec2-restart", "detach-volume"),
						},
					},

					"aws_region": {
						Description:         "AWSRegion defines the region of aws.",
						MarkdownDescription: "AWSRegion defines the region of aws.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"device_name": {
						Description:         "DeviceName indicates the name of the device. Needed in detach-volume.",
						MarkdownDescription: "DeviceName indicates the name of the device. Needed in detach-volume.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"duration": {
						Description:         "Duration represents the duration of the chaos action.",
						MarkdownDescription: "Duration represents the duration of the chaos action.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ec2_instance": {
						Description:         "Ec2Instance indicates the ID of the ec2 instance.",
						MarkdownDescription: "Ec2Instance indicates the ID of the ec2 instance.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"endpoint": {
						Description:         "Endpoint indicates the endpoint of the aws server. Just used it in test now.",
						MarkdownDescription: "Endpoint indicates the endpoint of the aws server. Just used it in test now.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_name": {
						Description:         "SecretName defines the name of kubernetes secret.",
						MarkdownDescription: "SecretName defines the name of kubernetes secret.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"volume_id": {
						Description:         "EbsVolume indicates the ID of the EBS volume. Needed in detach-volume.",
						MarkdownDescription: "EbsVolume indicates the ID of the EBS volume. Needed in detach-volume.",

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
		},
	}, nil
}

func (r *ChaosMeshOrgAWSChaosV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_chaos_mesh_org_aws_chaos_v1alpha1")

	var state ChaosMeshOrgAWSChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgAWSChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("AWSChaos")

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

func (r *ChaosMeshOrgAWSChaosV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_aws_chaos_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ChaosMeshOrgAWSChaosV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_chaos_mesh_org_aws_chaos_v1alpha1")

	var state ChaosMeshOrgAWSChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgAWSChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("AWSChaos")

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

func (r *ChaosMeshOrgAWSChaosV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_chaos_mesh_org_aws_chaos_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
