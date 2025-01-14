---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "applicationautoscaling.services.k8s.aws"
description: |-
  ScalingPolicy is the Schema for the ScalingPolicies API
---

# k8s_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1 (Resource)

ScalingPolicy is the Schema for the ScalingPolicies API

## Example Usage

```terraform
resource "k8s_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1" "minimal" {
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

- `spec` (Attributes) ScalingPolicySpec defines the desired state of ScalingPolicy.  Represents a scaling policy to use with Application Auto Scaling.  For more information about configuring scaling policies for a specific service, see Getting started with Application Auto Scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/getting-started.html) in the Application Auto Scaling User Guide. (see [below for nested schema](#nestedatt--spec))

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

- `policy_name` (String) The name of the scaling policy.
- `resource_id` (String) The identifier of the resource associated with the scaling policy. This string consists of the resource type and unique identifier.  * ECS service - The resource type is service and the unique identifier is the cluster name and service name. Example: service/default/sample-webapp.  * Spot Fleet - The resource type is spot-fleet-request and the unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.  * EMR cluster - The resource type is instancegroup and the unique identifier is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.  * AppStream 2.0 fleet - The resource type is fleet and the unique identifier is the fleet name. Example: fleet/sample-fleet.  * DynamoDB table - The resource type is table and the unique identifier is the table name. Example: table/my-table.  * DynamoDB global secondary index - The resource type is index and the unique identifier is the index name. Example: table/my-table/index/my-table-index.  * Aurora DB cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:my-db-cluster.  * SageMaker endpoint variant - The resource type is variant and the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.  * Custom resources are not supported with a resource type. This parameter must specify the OutputValue from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).  * Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.  * Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE.  * Lambda provisioned concurrency - The resource type is function and the unique identifier is the function name with a function version or alias name suffix that is not $LATEST. Example: function:my-function:prod or function:my-function:1.  * Amazon Keyspaces table - The resource type is table and the unique identifier is the table name. Example: keyspace/mykeyspace/table/mytable.  * Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5.  * Amazon ElastiCache replication group - The resource type is replication-group and the unique identifier is the replication group name. Example: replication-group/mycluster.  * Neptune cluster - The resource type is cluster and the unique identifier is the cluster name. Example: cluster:mycluster.
- `scalable_dimension` (String) The scalable dimension. This string consists of the service namespace, resource type, and scaling property.  * ecs:service:DesiredCount - The desired task count of an ECS service.  * elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR Instance Group.  * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet.  * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream 2.0 fleet.  * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a DynamoDB table.  * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a DynamoDB table.  * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB global secondary index.  * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a DynamoDB global secondary index.  * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.  * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for an SageMaker model endpoint variant.  * custom-resource:ResourceType:Property - The scalable dimension for a custom resource provided by your own application or service.  * comprehend:document-classifier-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend document classification endpoint.  * comprehend:entity-recognizer-endpoint:DesiredInferenceUnits - The number of inference units for an Amazon Comprehend entity recognizer endpoint.  * lambda:function:ProvisionedConcurrency - The provisioned concurrency for a Lambda function.  * cassandra:table:ReadCapacityUnits - The provisioned read capacity for an Amazon Keyspaces table.  * cassandra:table:WriteCapacityUnits - The provisioned write capacity for an Amazon Keyspaces table.  * kafka:broker-storage:VolumeSize - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.  * elasticache:replication-group:NodeGroups - The number of node groups for an Amazon ElastiCache replication group.  * elasticache:replication-group:Replicas - The number of replicas per node group for an Amazon ElastiCache replication group.  * neptune:cluster:ReadReplicaCount - The count of read replicas in an Amazon Neptune DB cluster.
- `service_namespace` (String) The namespace of the Amazon Web Services service that provides the resource. For a resource provided by your own application or service, use custom-resource instead.

Optional:

- `policy_type` (String) The policy type. This parameter is required if you are creating a scaling policy.  The following policy types are supported:  TargetTrackingScaling—Not supported for Amazon EMR  StepScaling—Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.  For more information, see Target tracking scaling policies (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and Step scaling policies (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) in the Application Auto Scaling User Guide.
- `step_scaling_policy_configuration` (Attributes) A step scaling policy.  This parameter is required if you are creating a policy and the policy type is StepScaling. (see [below for nested schema](#nestedatt--spec--step_scaling_policy_configuration))
- `target_tracking_scaling_policy_configuration` (Attributes) A target tracking scaling policy. Includes support for predefined or customized metrics.  This parameter is required if you are creating a policy and the policy type is TargetTrackingScaling. (see [below for nested schema](#nestedatt--spec--target_tracking_scaling_policy_configuration))

<a id="nestedatt--spec--step_scaling_policy_configuration"></a>
### Nested Schema for `spec.step_scaling_policy_configuration`

Optional:

- `adjustment_type` (String)
- `cooldown` (Number)
- `metric_aggregation_type` (String)
- `min_adjustment_magnitude` (Number)
- `step_adjustments` (Attributes List) (see [below for nested schema](#nestedatt--spec--step_scaling_policy_configuration--step_adjustments))

<a id="nestedatt--spec--step_scaling_policy_configuration--step_adjustments"></a>
### Nested Schema for `spec.step_scaling_policy_configuration.step_adjustments`

Optional:

- `metric_interval_lower_bound` (Number)
- `metric_interval_upper_bound` (Number)
- `scaling_adjustment` (Number)



<a id="nestedatt--spec--target_tracking_scaling_policy_configuration"></a>
### Nested Schema for `spec.target_tracking_scaling_policy_configuration`

Optional:

- `customized_metric_specification` (Attributes) Represents a CloudWatch metric of your choosing for a target tracking scaling policy to use with Application Auto Scaling.  For information about the available metrics for a service, see Amazon Web Services Services That Publish CloudWatch Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html) in the Amazon CloudWatch User Guide.  To create your customized metric specification:  * Add values for each required parameter from CloudWatch. You can use an existing metric, or a new metric that you create. To use your own metric, you must first publish the metric to CloudWatch. For more information, see Publish Custom Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html) in the Amazon CloudWatch User Guide.  * Choose a metric that changes proportionally with capacity. The value of the metric should increase or decrease in inverse proportion to the number of capacity units. That is, the value of the metric should decrease when capacity increases, and increase when capacity decreases.  For more information about CloudWatch, see Amazon CloudWatch Concepts (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html). (see [below for nested schema](#nestedatt--spec--target_tracking_scaling_policy_configuration--customized_metric_specification))
- `disable_scale_in` (Boolean)
- `predefined_metric_specification` (Attributes) Represents a predefined metric for a target tracking scaling policy to use with Application Auto Scaling.  Only the Amazon Web Services that you're using send metrics to Amazon CloudWatch. To determine whether a desired metric already exists by looking up its namespace and dimension using the CloudWatch metrics dashboard in the console, follow the procedure in Building dashboards with CloudWatch (https://docs.aws.amazon.com/autoscaling/application/userguide/monitoring-cloudwatch.html) in the Application Auto Scaling User Guide. (see [below for nested schema](#nestedatt--spec--target_tracking_scaling_policy_configuration--predefined_metric_specification))
- `scale_in_cooldown` (Number)
- `scale_out_cooldown` (Number)
- `target_value` (Number)

<a id="nestedatt--spec--target_tracking_scaling_policy_configuration--customized_metric_specification"></a>
### Nested Schema for `spec.target_tracking_scaling_policy_configuration.customized_metric_specification`

Optional:

- `dimensions` (Attributes List) (see [below for nested schema](#nestedatt--spec--target_tracking_scaling_policy_configuration--customized_metric_specification--dimensions))
- `metric_name` (String)
- `namespace` (String)
- `statistic` (String)
- `unit` (String)

<a id="nestedatt--spec--target_tracking_scaling_policy_configuration--customized_metric_specification--dimensions"></a>
### Nested Schema for `spec.target_tracking_scaling_policy_configuration.customized_metric_specification.unit`

Optional:

- `name` (String)
- `value` (String)



<a id="nestedatt--spec--target_tracking_scaling_policy_configuration--predefined_metric_specification"></a>
### Nested Schema for `spec.target_tracking_scaling_policy_configuration.predefined_metric_specification`

Optional:

- `predefined_metric_type` (String)
- `resource_label` (String)


