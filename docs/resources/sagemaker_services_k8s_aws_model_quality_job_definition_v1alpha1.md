---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_sagemaker_services_k8s_aws_model_quality_job_definition_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "sagemaker.services.k8s.aws/v1alpha1"
description: |-
  ModelQualityJobDefinition is the Schema for the ModelQualityJobDefinitions API
---

# k8s_sagemaker_services_k8s_aws_model_quality_job_definition_v1alpha1 (Resource)

ModelQualityJobDefinition is the Schema for the ModelQualityJobDefinitions API

## Example Usage

```terraform
resource "k8s_sagemaker_services_k8s_aws_model_quality_job_definition_v1alpha1" "minimal" {
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

- `spec` (Attributes) ModelQualityJobDefinitionSpec defines the desired state of ModelQualityJobDefinition. (see [below for nested schema](#nestedatt--spec))

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

- `job_definition_name` (String) The name of the monitoring job definition.
- `job_resources` (Attributes) Identifies the resources to deploy for a monitoring job. (see [below for nested schema](#nestedatt--spec--job_resources))
- `model_quality_app_specification` (Attributes) The container that runs the monitoring job. (see [below for nested schema](#nestedatt--spec--model_quality_app_specification))
- `model_quality_job_input` (Attributes) A list of the inputs that are monitored. Currently endpoints are supported. (see [below for nested schema](#nestedatt--spec--model_quality_job_input))
- `model_quality_job_output_config` (Attributes) The output configuration for monitoring jobs. (see [below for nested schema](#nestedatt--spec--model_quality_job_output_config))
- `role_arn` (String) The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.

Optional:

- `model_quality_baseline_config` (Attributes) Specifies the constraints and baselines for the monitoring job. (see [below for nested schema](#nestedatt--spec--model_quality_baseline_config))
- `network_config` (Attributes) Specifies the network configuration for the monitoring job. (see [below for nested schema](#nestedatt--spec--network_config))
- `stopping_condition` (Attributes) A time limit for how long the monitoring job is allowed to run before stopping. (see [below for nested schema](#nestedatt--spec--stopping_condition))
- `tags` (Attributes List) (Optional) An array of key-value pairs. For more information, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL) in the Amazon Web Services Billing and Cost Management User Guide. (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--job_resources"></a>
### Nested Schema for `spec.job_resources`

Optional:

- `cluster_config` (Attributes) Configuration for the cluster used to run model monitoring jobs. (see [below for nested schema](#nestedatt--spec--job_resources--cluster_config))

<a id="nestedatt--spec--job_resources--cluster_config"></a>
### Nested Schema for `spec.job_resources.cluster_config`

Optional:

- `instance_count` (Number)
- `instance_type` (String)
- `volume_kms_key_id` (String)
- `volume_size_in_gb` (Number)



<a id="nestedatt--spec--model_quality_app_specification"></a>
### Nested Schema for `spec.model_quality_app_specification`

Optional:

- `container_arguments` (List of String)
- `container_entrypoint` (List of String)
- `environment` (Map of String)
- `image_uri` (String)
- `post_analytics_processor_source_uri` (String)
- `problem_type` (String)
- `record_preprocessor_source_uri` (String)


<a id="nestedatt--spec--model_quality_job_input"></a>
### Nested Schema for `spec.model_quality_job_input`

Optional:

- `endpoint_input` (Attributes) Input object for the endpoint (see [below for nested schema](#nestedatt--spec--model_quality_job_input--endpoint_input))
- `ground_truth_s3_input` (Attributes) The ground truth labels for the dataset used for the monitoring job. (see [below for nested schema](#nestedatt--spec--model_quality_job_input--ground_truth_s3_input))

<a id="nestedatt--spec--model_quality_job_input--endpoint_input"></a>
### Nested Schema for `spec.model_quality_job_input.endpoint_input`

Optional:

- `end_time_offset` (String)
- `endpoint_name` (String)
- `features_attribute` (String)
- `inference_attribute` (String)
- `local_path` (String)
- `probability_attribute` (String)
- `probability_threshold_attribute` (Number)
- `s3_data_distribution_type` (String)
- `s3_input_mode` (String)
- `start_time_offset` (String)


<a id="nestedatt--spec--model_quality_job_input--ground_truth_s3_input"></a>
### Nested Schema for `spec.model_quality_job_input.ground_truth_s3_input`

Optional:

- `s3_uri` (String)



<a id="nestedatt--spec--model_quality_job_output_config"></a>
### Nested Schema for `spec.model_quality_job_output_config`

Optional:

- `kms_key_id` (String)
- `monitoring_outputs` (Attributes List) (see [below for nested schema](#nestedatt--spec--model_quality_job_output_config--monitoring_outputs))

<a id="nestedatt--spec--model_quality_job_output_config--monitoring_outputs"></a>
### Nested Schema for `spec.model_quality_job_output_config.monitoring_outputs`

Optional:

- `s3_output` (Attributes) Information about where and how you want to store the results of a monitoring job. (see [below for nested schema](#nestedatt--spec--model_quality_job_output_config--monitoring_outputs--s3_output))

<a id="nestedatt--spec--model_quality_job_output_config--monitoring_outputs--s3_output"></a>
### Nested Schema for `spec.model_quality_job_output_config.monitoring_outputs.s3_output`

Optional:

- `local_path` (String)
- `s3_upload_mode` (String)
- `s3_uri` (String)




<a id="nestedatt--spec--model_quality_baseline_config"></a>
### Nested Schema for `spec.model_quality_baseline_config`

Optional:

- `baselining_job_name` (String)
- `constraints_resource` (Attributes) The constraints resource for a monitoring job. (see [below for nested schema](#nestedatt--spec--model_quality_baseline_config--constraints_resource))

<a id="nestedatt--spec--model_quality_baseline_config--constraints_resource"></a>
### Nested Schema for `spec.model_quality_baseline_config.constraints_resource`

Optional:

- `s3_uri` (String)



<a id="nestedatt--spec--network_config"></a>
### Nested Schema for `spec.network_config`

Optional:

- `enable_inter_container_traffic_encryption` (Boolean)
- `enable_network_isolation` (Boolean)
- `vpc_config` (Attributes) Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html). (see [below for nested schema](#nestedatt--spec--network_config--vpc_config))

<a id="nestedatt--spec--network_config--vpc_config"></a>
### Nested Schema for `spec.network_config.vpc_config`

Optional:

- `security_group_i_ds` (List of String)
- `subnets` (List of String)



<a id="nestedatt--spec--stopping_condition"></a>
### Nested Schema for `spec.stopping_condition`

Optional:

- `max_runtime_in_seconds` (Number)


<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)

