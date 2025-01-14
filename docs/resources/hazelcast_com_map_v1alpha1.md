---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_hazelcast_com_map_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "hazelcast.com"
description: |-
  Map is the Schema for the maps API
---

# k8s_hazelcast_com_map_v1alpha1 (Resource)

Map is the Schema for the maps API

## Example Usage

```terraform
resource "k8s_hazelcast_com_map_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    hazelcast_resource_name = "some-name"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) MapSpec defines the desired state of Hazelcast Map Config (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `api_version` (String) APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
- `id` (Number) The timestamp of the last change to this resource.
- `kind` (String) Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Required:

- `hazelcast_resource_name` (String) HazelcastResourceName defines the name of the Hazelcast resource. It cannot be updated after map config is created successfully.

Optional:

- `backup_count` (Number) Count of synchronous backups. It cannot be updated after map config is created successfully.
- `entry_listeners` (Attributes List) EntryListeners contains the configuration for the map-level or entry-based events listeners provided by the Hazelcast’s eventing framework. You can learn more at https://docs.hazelcast.com/hazelcast/latest/events/object-events. (see [below for nested schema](#nestedatt--spec--entry_listeners))
- `eviction` (Attributes) Configuration for removing data from the map when it reaches its max size. It can be updated. (see [below for nested schema](#nestedatt--spec--eviction))
- `in_memory_format` (String) InMemoryFormat specifies in which format data will be stored in your map
- `indexes` (Attributes List) Indexes to be created for the map data. You can learn more at https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps. It cannot be updated after map config is created successfully. (see [below for nested schema](#nestedatt--spec--indexes))
- `map_store` (Attributes) Configuration options when you want to load/store the map entries from/to a persistent data store such as a relational database You can learn more at https://docs.hazelcast.com/hazelcast/latest/data-structures/working-with-external-data (see [below for nested schema](#nestedatt--spec--map_store))
- `max_idle_seconds` (Number) Maximum time in seconds for each entry to stay idle in the map. Entries that are idle for more than this time are evicted automatically. It can be updated.
- `name` (String) Name of the map config to be created. If empty, CR name will be used. It cannot be updated after map config is created successfully.
- `persistence_enabled` (Boolean) When enabled, map data will be persisted. It cannot be updated after map config is created successfully.
- `time_to_live_seconds` (Number) Maximum time in seconds for each entry to stay in the map. If it is not 0, entries that are older than this time and not updated for this time are evicted automatically. It can be updated.

<a id="nestedatt--spec--entry_listeners"></a>
### Nested Schema for `spec.entry_listeners`

Required:

- `class_name` (String) ClassName is the fully qualified name of the class that implements any of the Listener interface.

Optional:

- `include_values` (Boolean) IncludeValues is an optional attribute that indicates whether the event will contain the map value. Defaults to true.
- `local` (Boolean) Local is an optional attribute that indicates whether the map on the local member can be listened to. Defaults to false.


<a id="nestedatt--spec--eviction"></a>
### Nested Schema for `spec.eviction`

Optional:

- `eviction_policy` (String) Eviction policy to be applied when map reaches its max size according to the max size policy.
- `max_size` (Number) Max size of the map.
- `max_size_policy` (String) Policy for deciding if the maxSize is reached.


<a id="nestedatt--spec--indexes"></a>
### Nested Schema for `spec.indexes`

Required:

- `attributes` (List of String) Attributes of the index.
- `type` (String) Type of the index.

Optional:

- `bit_map_index_options` (Attributes) Options for 'BITMAP' index type. (see [below for nested schema](#nestedatt--spec--indexes--bit_map_index_options))
- `name` (String) Name of the index config.

<a id="nestedatt--spec--indexes--bit_map_index_options"></a>
### Nested Schema for `spec.indexes.bit_map_index_options`

Required:

- `unique_key` (String)
- `unique_key_transition` (String)



<a id="nestedatt--spec--map_store"></a>
### Nested Schema for `spec.map_store`

Required:

- `class_name` (String) Name of your class implementing MapLoader and/or MapStore interface.

Optional:

- `initial_mode` (String) Sets the initial entry loading mode.
- `properties_secret_name` (String) Properties can be used for giving information to the MapStore implementation
- `write_batch_size` (Number) Used to create batches when writing to map store.
- `write_coealescing` (Boolean) It is meaningful if you are using write behind in MapStore. When it is set to true, only the latest store operation on a key during the write-delay-seconds will be reflected to MapStore.
- `write_delay_seconds` (Number) Number of seconds to delay the storing of entries.


