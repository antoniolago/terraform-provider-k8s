/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type HazelcastComMapV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*HazelcastComMapV1Alpha1Resource)(nil)
)

type HazelcastComMapV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HazelcastComMapV1Alpha1GoModel struct {
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
		BackupCount *int64 `tfsdk:"backup_count" yaml:"backupCount,omitempty"`

		EntryListeners *[]struct {
			ClassName *string `tfsdk:"class_name" yaml:"className,omitempty"`

			IncludeValues *bool `tfsdk:"include_values" yaml:"includeValues,omitempty"`

			Local *bool `tfsdk:"local" yaml:"local,omitempty"`
		} `tfsdk:"entry_listeners" yaml:"entryListeners,omitempty"`

		Eviction *struct {
			EvictionPolicy *string `tfsdk:"eviction_policy" yaml:"evictionPolicy,omitempty"`

			MaxSize *int64 `tfsdk:"max_size" yaml:"maxSize,omitempty"`

			MaxSizePolicy *string `tfsdk:"max_size_policy" yaml:"maxSizePolicy,omitempty"`
		} `tfsdk:"eviction" yaml:"eviction,omitempty"`

		HazelcastResourceName *string `tfsdk:"hazelcast_resource_name" yaml:"hazelcastResourceName,omitempty"`

		InMemoryFormat *string `tfsdk:"in_memory_format" yaml:"inMemoryFormat,omitempty"`

		Indexes *[]struct {
			Attributes *[]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

			BitMapIndexOptions *struct {
				UniqueKey *string `tfsdk:"unique_key" yaml:"uniqueKey,omitempty"`

				UniqueKeyTransition *string `tfsdk:"unique_key_transition" yaml:"uniqueKeyTransition,omitempty"`
			} `tfsdk:"bit_map_index_options" yaml:"bitMapIndexOptions,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"indexes" yaml:"indexes,omitempty"`

		MapStore *struct {
			ClassName *string `tfsdk:"class_name" yaml:"className,omitempty"`

			InitialMode *string `tfsdk:"initial_mode" yaml:"initialMode,omitempty"`

			PropertiesSecretName *string `tfsdk:"properties_secret_name" yaml:"propertiesSecretName,omitempty"`

			WriteBatchSize *int64 `tfsdk:"write_batch_size" yaml:"writeBatchSize,omitempty"`

			WriteCoealescing *bool `tfsdk:"write_coealescing" yaml:"writeCoealescing,omitempty"`

			WriteDelaySeconds *int64 `tfsdk:"write_delay_seconds" yaml:"writeDelaySeconds,omitempty"`
		} `tfsdk:"map_store" yaml:"mapStore,omitempty"`

		MaxIdleSeconds *int64 `tfsdk:"max_idle_seconds" yaml:"maxIdleSeconds,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		PersistenceEnabled *bool `tfsdk:"persistence_enabled" yaml:"persistenceEnabled,omitempty"`

		TimeToLiveSeconds *int64 `tfsdk:"time_to_live_seconds" yaml:"timeToLiveSeconds,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHazelcastComMapV1Alpha1Resource() resource.Resource {
	return &HazelcastComMapV1Alpha1Resource{}
}

func (r *HazelcastComMapV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hazelcast_com_map_v1alpha1"
}

func (r *HazelcastComMapV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Map is the Schema for the maps API",
		MarkdownDescription: "Map is the Schema for the maps API",
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
				Description:         "MapSpec defines the desired state of Hazelcast Map Config",
				MarkdownDescription: "MapSpec defines the desired state of Hazelcast Map Config",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"backup_count": {
						Description:         "Count of synchronous backups. It cannot be updated after map config is created successfully.",
						MarkdownDescription: "Count of synchronous backups. It cannot be updated after map config is created successfully.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"entry_listeners": {
						Description:         "EntryListeners contains the configuration for the map-level or entry-based events listeners provided by the Hazelcast’s eventing framework. You can learn more at https://docs.hazelcast.com/hazelcast/latest/events/object-events.",
						MarkdownDescription: "EntryListeners contains the configuration for the map-level or entry-based events listeners provided by the Hazelcast’s eventing framework. You can learn more at https://docs.hazelcast.com/hazelcast/latest/events/object-events.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"class_name": {
								Description:         "ClassName is the fully qualified name of the class that implements any of the Listener interface.",
								MarkdownDescription: "ClassName is the fully qualified name of the class that implements any of the Listener interface.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtLeast(1),
								},
							},

							"include_values": {
								Description:         "IncludeValues is an optional attribute that indicates whether the event will contain the map value. Defaults to true.",
								MarkdownDescription: "IncludeValues is an optional attribute that indicates whether the event will contain the map value. Defaults to true.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"local": {
								Description:         "Local is an optional attribute that indicates whether the map on the local member can be listened to. Defaults to false.",
								MarkdownDescription: "Local is an optional attribute that indicates whether the map on the local member can be listened to. Defaults to false.",

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

					"eviction": {
						Description:         "Configuration for removing data from the map when it reaches its max size. It can be updated.",
						MarkdownDescription: "Configuration for removing data from the map when it reaches its max size. It can be updated.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"eviction_policy": {
								Description:         "Eviction policy to be applied when map reaches its max size according to the max size policy.",
								MarkdownDescription: "Eviction policy to be applied when map reaches its max size according to the max size policy.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("NONE", "LRU", "LFU", "RANDOM"),
								},
							},

							"max_size": {
								Description:         "Max size of the map.",
								MarkdownDescription: "Max size of the map.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_size_policy": {
								Description:         "Policy for deciding if the maxSize is reached.",
								MarkdownDescription: "Policy for deciding if the maxSize is reached.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("PER_NODE", "PER_PARTITION", "USED_HEAP_SIZE", "USED_HEAP_PERCENTAGE", "FREE_HEAP_SIZE", "FREE_HEAP_PERCENTAGE", "USED_NATIVE_MEMORY_SIZE", "USED_NATIVE_MEMORY_PERCENTAGE", "FREE_NATIVE_MEMORY_SIZE", "FREE_NATIVE_MEMORY_PERCENTAGE"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"hazelcast_resource_name": {
						Description:         "HazelcastResourceName defines the name of the Hazelcast resource. It cannot be updated after map config is created successfully.",
						MarkdownDescription: "HazelcastResourceName defines the name of the Hazelcast resource. It cannot be updated after map config is created successfully.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtLeast(1),
						},
					},

					"in_memory_format": {
						Description:         "InMemoryFormat specifies in which format data will be stored in your map",
						MarkdownDescription: "InMemoryFormat specifies in which format data will be stored in your map",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("BINARY", "OBJECT"),
						},
					},

					"indexes": {
						Description:         "Indexes to be created for the map data. You can learn more at https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps. It cannot be updated after map config is created successfully.",
						MarkdownDescription: "Indexes to be created for the map data. You can learn more at https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps. It cannot be updated after map config is created successfully.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"attributes": {
								Description:         "Attributes of the index.",
								MarkdownDescription: "Attributes of the index.",

								Type: types.ListType{ElemType: types.StringType},

								Required: true,
								Optional: false,
								Computed: false,
							},

							"bit_map_index_options": {
								Description:         "Options for 'BITMAP' index type.",
								MarkdownDescription: "Options for 'BITMAP' index type.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"unique_key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"unique_key_transition": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("OBJECT", "LONG", "RAW"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the index config.",
								MarkdownDescription: "Name of the index config.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "Type of the index.",
								MarkdownDescription: "Type of the index.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("SORTED", "HASH", "BITMAP"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"map_store": {
						Description:         "Configuration options when you want to load/store the map entries from/to a persistent data store such as a relational database You can learn more at https://docs.hazelcast.com/hazelcast/latest/data-structures/working-with-external-data",
						MarkdownDescription: "Configuration options when you want to load/store the map entries from/to a persistent data store such as a relational database You can learn more at https://docs.hazelcast.com/hazelcast/latest/data-structures/working-with-external-data",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"class_name": {
								Description:         "Name of your class implementing MapLoader and/or MapStore interface.",
								MarkdownDescription: "Name of your class implementing MapLoader and/or MapStore interface.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"initial_mode": {
								Description:         "Sets the initial entry loading mode.",
								MarkdownDescription: "Sets the initial entry loading mode.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("LAZY", "EAGER"),
								},
							},

							"properties_secret_name": {
								Description:         "Properties can be used for giving information to the MapStore implementation",
								MarkdownDescription: "Properties can be used for giving information to the MapStore implementation",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"write_batch_size": {
								Description:         "Used to create batches when writing to map store.",
								MarkdownDescription: "Used to create batches when writing to map store.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"write_coealescing": {
								Description:         "It is meaningful if you are using write behind in MapStore. When it is set to true, only the latest store operation on a key during the write-delay-seconds will be reflected to MapStore.",
								MarkdownDescription: "It is meaningful if you are using write behind in MapStore. When it is set to true, only the latest store operation on a key during the write-delay-seconds will be reflected to MapStore.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"write_delay_seconds": {
								Description:         "Number of seconds to delay the storing of entries.",
								MarkdownDescription: "Number of seconds to delay the storing of entries.",

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

					"max_idle_seconds": {
						Description:         "Maximum time in seconds for each entry to stay idle in the map. Entries that are idle for more than this time are evicted automatically. It can be updated.",
						MarkdownDescription: "Maximum time in seconds for each entry to stay idle in the map. Entries that are idle for more than this time are evicted automatically. It can be updated.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": {
						Description:         "Name of the map config to be created. If empty, CR name will be used. It cannot be updated after map config is created successfully.",
						MarkdownDescription: "Name of the map config to be created. If empty, CR name will be used. It cannot be updated after map config is created successfully.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"persistence_enabled": {
						Description:         "When enabled, map data will be persisted. It cannot be updated after map config is created successfully.",
						MarkdownDescription: "When enabled, map data will be persisted. It cannot be updated after map config is created successfully.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"time_to_live_seconds": {
						Description:         "Maximum time in seconds for each entry to stay in the map. If it is not 0, entries that are older than this time and not updated for this time are evicted automatically. It can be updated.",
						MarkdownDescription: "Maximum time in seconds for each entry to stay in the map. If it is not 0, entries that are older than this time and not updated for this time are evicted automatically. It can be updated.",

						Type: types.Int64Type,

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

func (r *HazelcastComMapV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hazelcast_com_map_v1alpha1")

	var state HazelcastComMapV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HazelcastComMapV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hazelcast.com/v1alpha1")
	goModel.Kind = utilities.Ptr("Map")

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

func (r *HazelcastComMapV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hazelcast_com_map_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *HazelcastComMapV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hazelcast_com_map_v1alpha1")

	var state HazelcastComMapV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HazelcastComMapV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hazelcast.com/v1alpha1")
	goModel.Kind = utilities.Ptr("Map")

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

func (r *HazelcastComMapV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hazelcast_com_map_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
