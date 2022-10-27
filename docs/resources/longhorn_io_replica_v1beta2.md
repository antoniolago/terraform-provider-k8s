---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_longhorn_io_replica_v1beta2 Resource - terraform-provider-k8s"
subcategory: "longhorn.io"
description: |-
  Replica is where Longhorn stores replica object.
---

# k8s_longhorn_io_replica_v1beta2 (Resource)

Replica is where Longhorn stores replica object.

## Example Usage

```terraform
resource "k8s_longhorn_io_replica_v1beta2" "minimal" {
  metadata = {
    name = "test"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) ReplicaSpec defines the desired state of the Longhorn replica (see [below for nested schema](#nestedatt--spec))

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

Optional:

- `active` (Boolean)
- `backing_image` (String)
- `base_image` (String) Deprecated. Rename to BackingImage
- `data_directory_name` (String)
- `data_path` (String) Deprecated
- `desire_state` (String)
- `disk_id` (String)
- `disk_path` (String)
- `engine_image` (String)
- `engine_name` (String)
- `failed_at` (String)
- `hard_node_affinity` (String)
- `healthy_at` (String)
- `log_requested` (Boolean)
- `node_id` (String)
- `rebuild_retry_count` (Number)
- `revision_counter_disabled` (Boolean)
- `salvage_requested` (Boolean)
- `volume_name` (String)
- `volume_size` (String)

