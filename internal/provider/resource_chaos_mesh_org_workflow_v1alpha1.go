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

type ChaosMeshOrgWorkflowV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ChaosMeshOrgWorkflowV1Alpha1Resource)(nil)
)

type ChaosMeshOrgWorkflowV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ChaosMeshOrgWorkflowV1Alpha1GoModel struct {
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
		Entry *string `tfsdk:"entry" yaml:"entry,omitempty"`

		Templates *[]struct {
			AbortWithStatusCheck *bool `tfsdk:"abort_with_status_check" yaml:"abortWithStatusCheck,omitempty"`

			AwsChaos *struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				AwsRegion *string `tfsdk:"aws_region" yaml:"awsRegion,omitempty"`

				DeviceName *string `tfsdk:"device_name" yaml:"deviceName,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				Ec2Instance *string `tfsdk:"ec2_instance" yaml:"ec2Instance,omitempty"`

				Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

				VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
			} `tfsdk:"aws_chaos" yaml:"awsChaos,omitempty"`

			AzureChaos *struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				DiskName *string `tfsdk:"disk_name" yaml:"diskName,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				Lun *int64 `tfsdk:"lun" yaml:"lun,omitempty"`

				ResourceGroupName *string `tfsdk:"resource_group_name" yaml:"resourceGroupName,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

				SubscriptionID *string `tfsdk:"subscription_id" yaml:"subscriptionID,omitempty"`

				VmName *string `tfsdk:"vm_name" yaml:"vmName,omitempty"`
			} `tfsdk:"azure_chaos" yaml:"azureChaos,omitempty"`

			BlockChaos *struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

				Delay *struct {
					Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

					Jitter *string `tfsdk:"jitter" yaml:"jitter,omitempty"`

					Latency *string `tfsdk:"latency" yaml:"latency,omitempty"`
				} `tfsdk:"delay" yaml:"delay,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				Selector *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

					ExpressionSelectors *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

					FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

					LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

					Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

					PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

					Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`
			} `tfsdk:"block_chaos" yaml:"blockChaos,omitempty"`

			Children *[]string `tfsdk:"children" yaml:"children,omitempty"`

			ConditionalBranches *[]struct {
				Expression *string `tfsdk:"expression" yaml:"expression,omitempty"`

				Target *string `tfsdk:"target" yaml:"target,omitempty"`
			} `tfsdk:"conditional_branches" yaml:"conditionalBranches,omitempty"`

			Deadline *string `tfsdk:"deadline" yaml:"deadline,omitempty"`

			DnsChaos *struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				Patterns *[]string `tfsdk:"patterns" yaml:"patterns,omitempty"`

				Selector *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

					ExpressionSelectors *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

					FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

					LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

					Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

					PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

					Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"dns_chaos" yaml:"dnsChaos,omitempty"`

			GcpChaos *struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				DeviceNames *[]string `tfsdk:"device_names" yaml:"deviceNames,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				Instance *string `tfsdk:"instance" yaml:"instance,omitempty"`

				Project *string `tfsdk:"project" yaml:"project,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

				Zone *string `tfsdk:"zone" yaml:"zone,omitempty"`
			} `tfsdk:"gcp_chaos" yaml:"gcpChaos,omitempty"`

			HttpChaos *struct {
				Abort *bool `tfsdk:"abort" yaml:"abort,omitempty"`

				Code *int64 `tfsdk:"code" yaml:"code,omitempty"`

				Delay *string `tfsdk:"delay" yaml:"delay,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				Method *string `tfsdk:"method" yaml:"method,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				Patch *struct {
					Body *struct {
						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"body" yaml:"body,omitempty"`

					Headers *[]string `tfsdk:"headers" yaml:"headers,omitempty"`

					Queries *[]string `tfsdk:"queries" yaml:"queries,omitempty"`
				} `tfsdk:"patch" yaml:"patch,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				Replace *struct {
					Body *string `tfsdk:"body" yaml:"body,omitempty"`

					Code *int64 `tfsdk:"code" yaml:"code,omitempty"`

					Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

					Method *string `tfsdk:"method" yaml:"method,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Queries *map[string]string `tfsdk:"queries" yaml:"queries,omitempty"`
				} `tfsdk:"replace" yaml:"replace,omitempty"`

				Request_headers *map[string]string `tfsdk:"request_headers" yaml:"request_headers,omitempty"`

				Response_headers *map[string]string `tfsdk:"response_headers" yaml:"response_headers,omitempty"`

				Selector *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

					ExpressionSelectors *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

					FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

					LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

					Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

					PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

					Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				Target *string `tfsdk:"target" yaml:"target,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"http_chaos" yaml:"httpChaos,omitempty"`

			IoChaos *struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				Attr *struct {
					Atime *struct {
						Nsec *int64 `tfsdk:"nsec" yaml:"nsec,omitempty"`

						Sec *int64 `tfsdk:"sec" yaml:"sec,omitempty"`
					} `tfsdk:"atime" yaml:"atime,omitempty"`

					Blocks *int64 `tfsdk:"blocks" yaml:"blocks,omitempty"`

					Ctime *struct {
						Nsec *int64 `tfsdk:"nsec" yaml:"nsec,omitempty"`

						Sec *int64 `tfsdk:"sec" yaml:"sec,omitempty"`
					} `tfsdk:"ctime" yaml:"ctime,omitempty"`

					Gid *int64 `tfsdk:"gid" yaml:"gid,omitempty"`

					Ino *int64 `tfsdk:"ino" yaml:"ino,omitempty"`

					Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

					Mtime *struct {
						Nsec *int64 `tfsdk:"nsec" yaml:"nsec,omitempty"`

						Sec *int64 `tfsdk:"sec" yaml:"sec,omitempty"`
					} `tfsdk:"mtime" yaml:"mtime,omitempty"`

					Nlink *int64 `tfsdk:"nlink" yaml:"nlink,omitempty"`

					Perm *int64 `tfsdk:"perm" yaml:"perm,omitempty"`

					Rdev *int64 `tfsdk:"rdev" yaml:"rdev,omitempty"`

					Size *int64 `tfsdk:"size" yaml:"size,omitempty"`

					Uid *int64 `tfsdk:"uid" yaml:"uid,omitempty"`
				} `tfsdk:"attr" yaml:"attr,omitempty"`

				ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

				Delay *string `tfsdk:"delay" yaml:"delay,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				Errno *int64 `tfsdk:"errno" yaml:"errno,omitempty"`

				Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

				Mistake *struct {
					Filling *string `tfsdk:"filling" yaml:"filling,omitempty"`

					MaxLength *int64 `tfsdk:"max_length" yaml:"maxLength,omitempty"`

					MaxOccurrences *int64 `tfsdk:"max_occurrences" yaml:"maxOccurrences,omitempty"`
				} `tfsdk:"mistake" yaml:"mistake,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Percent *int64 `tfsdk:"percent" yaml:"percent,omitempty"`

				Selector *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

					ExpressionSelectors *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

					FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

					LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

					Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

					PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

					Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				VolumePath *string `tfsdk:"volume_path" yaml:"volumePath,omitempty"`
			} `tfsdk:"io_chaos" yaml:"ioChaos,omitempty"`

			JvmChaos *struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				Class *string `tfsdk:"class" yaml:"class,omitempty"`

				ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

				CpuCount *int64 `tfsdk:"cpu_count" yaml:"cpuCount,omitempty"`

				Database *string `tfsdk:"database" yaml:"database,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				Exception *string `tfsdk:"exception" yaml:"exception,omitempty"`

				Latency *int64 `tfsdk:"latency" yaml:"latency,omitempty"`

				MemType *string `tfsdk:"mem_type" yaml:"memType,omitempty"`

				Method *string `tfsdk:"method" yaml:"method,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				MysqlConnectorVersion *string `tfsdk:"mysql_connector_version" yaml:"mysqlConnectorVersion,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				RuleData *string `tfsdk:"rule_data" yaml:"ruleData,omitempty"`

				Selector *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

					ExpressionSelectors *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

					FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

					LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

					Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

					PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

					Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				SqlType *string `tfsdk:"sql_type" yaml:"sqlType,omitempty"`

				Table *string `tfsdk:"table" yaml:"table,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"jvm_chaos" yaml:"jvmChaos,omitempty"`

			KernelChaos *struct {
				ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				FailKernRequest *struct {
					Callchain *[]struct {
						Funcname *string `tfsdk:"funcname" yaml:"funcname,omitempty"`

						Parameters *string `tfsdk:"parameters" yaml:"parameters,omitempty"`

						Predicate *string `tfsdk:"predicate" yaml:"predicate,omitempty"`
					} `tfsdk:"callchain" yaml:"callchain,omitempty"`

					Failtype *int64 `tfsdk:"failtype" yaml:"failtype,omitempty"`

					Headers *[]string `tfsdk:"headers" yaml:"headers,omitempty"`

					Probability *int64 `tfsdk:"probability" yaml:"probability,omitempty"`

					Times *int64 `tfsdk:"times" yaml:"times,omitempty"`
				} `tfsdk:"fail_kern_request" yaml:"failKernRequest,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				Selector *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

					ExpressionSelectors *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

					FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

					LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

					Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

					PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

					Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"kernel_chaos" yaml:"kernelChaos,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			NetworkChaos *struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				Bandwidth *struct {
					Buffer *int64 `tfsdk:"buffer" yaml:"buffer,omitempty"`

					Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`

					Minburst *int64 `tfsdk:"minburst" yaml:"minburst,omitempty"`

					Peakrate *int64 `tfsdk:"peakrate" yaml:"peakrate,omitempty"`

					Rate *string `tfsdk:"rate" yaml:"rate,omitempty"`
				} `tfsdk:"bandwidth" yaml:"bandwidth,omitempty"`

				Corrupt *struct {
					Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

					Corrupt *string `tfsdk:"corrupt" yaml:"corrupt,omitempty"`
				} `tfsdk:"corrupt" yaml:"corrupt,omitempty"`

				Delay *struct {
					Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

					Jitter *string `tfsdk:"jitter" yaml:"jitter,omitempty"`

					Latency *string `tfsdk:"latency" yaml:"latency,omitempty"`

					Reorder *struct {
						Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

						Gap *int64 `tfsdk:"gap" yaml:"gap,omitempty"`

						Reorder *string `tfsdk:"reorder" yaml:"reorder,omitempty"`
					} `tfsdk:"reorder" yaml:"reorder,omitempty"`
				} `tfsdk:"delay" yaml:"delay,omitempty"`

				Device *string `tfsdk:"device" yaml:"device,omitempty"`

				Direction *string `tfsdk:"direction" yaml:"direction,omitempty"`

				Duplicate *struct {
					Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

					Duplicate *string `tfsdk:"duplicate" yaml:"duplicate,omitempty"`
				} `tfsdk:"duplicate" yaml:"duplicate,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				ExternalTargets *[]string `tfsdk:"external_targets" yaml:"externalTargets,omitempty"`

				Loss *struct {
					Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

					Loss *string `tfsdk:"loss" yaml:"loss,omitempty"`
				} `tfsdk:"loss" yaml:"loss,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				Selector *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

					ExpressionSelectors *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

					FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

					LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

					Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

					PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

					Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				Target *struct {
					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

						ExpressionSelectors *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

						FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

						LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

						Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

						PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

						Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"target" yaml:"target,omitempty"`

				TargetDevice *string `tfsdk:"target_device" yaml:"targetDevice,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"network_chaos" yaml:"networkChaos,omitempty"`

			PhysicalmachineChaos *struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				Address *[]string `tfsdk:"address" yaml:"address,omitempty"`

				Clock *struct {
					Clock_ids_slice *string `tfsdk:"clock_ids_slice" yaml:"clock-ids-slice,omitempty"`

					Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

					Time_offset *string `tfsdk:"time_offset" yaml:"time-offset,omitempty"`
				} `tfsdk:"clock" yaml:"clock,omitempty"`

				Disk_fill *struct {
					Fill_by_fallocate *bool `tfsdk:"fill_by_fallocate" yaml:"fill-by-fallocate,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Size *string `tfsdk:"size" yaml:"size,omitempty"`
				} `tfsdk:"disk_fill" yaml:"disk-fill,omitempty"`

				Disk_read_payload *struct {
					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Payload_process_num *int64 `tfsdk:"payload_process_num" yaml:"payload-process-num,omitempty"`

					Size *string `tfsdk:"size" yaml:"size,omitempty"`
				} `tfsdk:"disk_read_payload" yaml:"disk-read-payload,omitempty"`

				Disk_write_payload *struct {
					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Payload_process_num *int64 `tfsdk:"payload_process_num" yaml:"payload-process-num,omitempty"`

					Size *string `tfsdk:"size" yaml:"size,omitempty"`
				} `tfsdk:"disk_write_payload" yaml:"disk-write-payload,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				File_append *struct {
					Count *int64 `tfsdk:"count" yaml:"count,omitempty"`

					Data *string `tfsdk:"data" yaml:"data,omitempty"`

					File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`
				} `tfsdk:"file_append" yaml:"file-append,omitempty"`

				File_create *struct {
					Dir_name *string `tfsdk:"dir_name" yaml:"dir-name,omitempty"`

					File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`
				} `tfsdk:"file_create" yaml:"file-create,omitempty"`

				File_delete *struct {
					Dir_name *string `tfsdk:"dir_name" yaml:"dir-name,omitempty"`

					File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`
				} `tfsdk:"file_delete" yaml:"file-delete,omitempty"`

				File_modify *struct {
					File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`

					Privilege *int64 `tfsdk:"privilege" yaml:"privilege,omitempty"`
				} `tfsdk:"file_modify" yaml:"file-modify,omitempty"`

				File_rename *struct {
					Dest_file *string `tfsdk:"dest_file" yaml:"dest-file,omitempty"`

					Source_file *string `tfsdk:"source_file" yaml:"source-file,omitempty"`
				} `tfsdk:"file_rename" yaml:"file-rename,omitempty"`

				File_replace *struct {
					Dest_string *string `tfsdk:"dest_string" yaml:"dest-string,omitempty"`

					File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`

					Line *int64 `tfsdk:"line" yaml:"line,omitempty"`

					Origin_string *string `tfsdk:"origin_string" yaml:"origin-string,omitempty"`
				} `tfsdk:"file_replace" yaml:"file-replace,omitempty"`

				Http_abort *struct {
					Code *string `tfsdk:"code" yaml:"code,omitempty"`

					Method *string `tfsdk:"method" yaml:"method,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Proxy_ports *[]string `tfsdk:"proxy_ports" yaml:"proxy_ports,omitempty"`

					Target *string `tfsdk:"target" yaml:"target,omitempty"`
				} `tfsdk:"http_abort" yaml:"http-abort,omitempty"`

				Http_config *struct {
					File_path *string `tfsdk:"file_path" yaml:"file_path,omitempty"`
				} `tfsdk:"http_config" yaml:"http-config,omitempty"`

				Http_delay *struct {
					Code *string `tfsdk:"code" yaml:"code,omitempty"`

					Delay *string `tfsdk:"delay" yaml:"delay,omitempty"`

					Method *string `tfsdk:"method" yaml:"method,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Proxy_ports *[]string `tfsdk:"proxy_ports" yaml:"proxy_ports,omitempty"`

					Target *string `tfsdk:"target" yaml:"target,omitempty"`
				} `tfsdk:"http_delay" yaml:"http-delay,omitempty"`

				Http_request *struct {
					Count *int64 `tfsdk:"count" yaml:"count,omitempty"`

					Enable_conn_pool *bool `tfsdk:"enable_conn_pool" yaml:"enable-conn-pool,omitempty"`

					Url *string `tfsdk:"url" yaml:"url,omitempty"`
				} `tfsdk:"http_request" yaml:"http-request,omitempty"`

				Jvm_exception *struct {
					Class *string `tfsdk:"class" yaml:"class,omitempty"`

					Exception *string `tfsdk:"exception" yaml:"exception,omitempty"`

					Method *string `tfsdk:"method" yaml:"method,omitempty"`

					Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
				} `tfsdk:"jvm_exception" yaml:"jvm-exception,omitempty"`

				Jvm_gc *struct {
					Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
				} `tfsdk:"jvm_gc" yaml:"jvm-gc,omitempty"`

				Jvm_latency *struct {
					Class *string `tfsdk:"class" yaml:"class,omitempty"`

					Latency *int64 `tfsdk:"latency" yaml:"latency,omitempty"`

					Method *string `tfsdk:"method" yaml:"method,omitempty"`

					Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
				} `tfsdk:"jvm_latency" yaml:"jvm-latency,omitempty"`

				Jvm_mysql *struct {
					Database *string `tfsdk:"database" yaml:"database,omitempty"`

					Exception *string `tfsdk:"exception" yaml:"exception,omitempty"`

					Latency *int64 `tfsdk:"latency" yaml:"latency,omitempty"`

					MysqlConnectorVersion *string `tfsdk:"mysql_connector_version" yaml:"mysqlConnectorVersion,omitempty"`

					Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					SqlType *string `tfsdk:"sql_type" yaml:"sqlType,omitempty"`

					Table *string `tfsdk:"table" yaml:"table,omitempty"`
				} `tfsdk:"jvm_mysql" yaml:"jvm-mysql,omitempty"`

				Jvm_return *struct {
					Class *string `tfsdk:"class" yaml:"class,omitempty"`

					Method *string `tfsdk:"method" yaml:"method,omitempty"`

					Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"jvm_return" yaml:"jvm-return,omitempty"`

				Jvm_rule_data *struct {
					Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Rule_data *string `tfsdk:"rule_data" yaml:"rule-data,omitempty"`
				} `tfsdk:"jvm_rule_data" yaml:"jvm-rule-data,omitempty"`

				Jvm_stress *struct {
					Cpu_count *int64 `tfsdk:"cpu_count" yaml:"cpu-count,omitempty"`

					Mem_type *string `tfsdk:"mem_type" yaml:"mem-type,omitempty"`

					Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
				} `tfsdk:"jvm_stress" yaml:"jvm-stress,omitempty"`

				Kafka_fill *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					MaxBytes *int64 `tfsdk:"max_bytes" yaml:"maxBytes,omitempty"`

					MessageSize *int64 `tfsdk:"message_size" yaml:"messageSize,omitempty"`

					Password *string `tfsdk:"password" yaml:"password,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					ReloadCommand *string `tfsdk:"reload_command" yaml:"reloadCommand,omitempty"`

					Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`

					Username *string `tfsdk:"username" yaml:"username,omitempty"`
				} `tfsdk:"kafka_fill" yaml:"kafka-fill,omitempty"`

				Kafka_flood *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					MessageSize *int64 `tfsdk:"message_size" yaml:"messageSize,omitempty"`

					Password *string `tfsdk:"password" yaml:"password,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Threads *int64 `tfsdk:"threads" yaml:"threads,omitempty"`

					Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`

					Username *string `tfsdk:"username" yaml:"username,omitempty"`
				} `tfsdk:"kafka_flood" yaml:"kafka-flood,omitempty"`

				Kafka_io *struct {
					ConfigFile *string `tfsdk:"config_file" yaml:"configFile,omitempty"`

					NonReadable *bool `tfsdk:"non_readable" yaml:"nonReadable,omitempty"`

					NonWritable *bool `tfsdk:"non_writable" yaml:"nonWritable,omitempty"`

					Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`
				} `tfsdk:"kafka_io" yaml:"kafka-io,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				Network_bandwidth *struct {
					Buffer *int64 `tfsdk:"buffer" yaml:"buffer,omitempty"`

					Device *string `tfsdk:"device" yaml:"device,omitempty"`

					Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

					Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

					Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`

					Minburst *int64 `tfsdk:"minburst" yaml:"minburst,omitempty"`

					Peakrate *int64 `tfsdk:"peakrate" yaml:"peakrate,omitempty"`

					Rate *string `tfsdk:"rate" yaml:"rate,omitempty"`
				} `tfsdk:"network_bandwidth" yaml:"network-bandwidth,omitempty"`

				Network_corrupt *struct {
					Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

					Device *string `tfsdk:"device" yaml:"device,omitempty"`

					Egress_port *string `tfsdk:"egress_port" yaml:"egress-port,omitempty"`

					Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

					Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

					Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`

					Percent *string `tfsdk:"percent" yaml:"percent,omitempty"`

					Source_port *string `tfsdk:"source_port" yaml:"source-port,omitempty"`
				} `tfsdk:"network_corrupt" yaml:"network-corrupt,omitempty"`

				Network_delay *struct {
					Accept_tcp_flags *string `tfsdk:"accept_tcp_flags" yaml:"accept-tcp-flags,omitempty"`

					Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

					Device *string `tfsdk:"device" yaml:"device,omitempty"`

					Egress_port *string `tfsdk:"egress_port" yaml:"egress-port,omitempty"`

					Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

					Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

					Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`

					Jitter *string `tfsdk:"jitter" yaml:"jitter,omitempty"`

					Latency *string `tfsdk:"latency" yaml:"latency,omitempty"`

					Source_port *string `tfsdk:"source_port" yaml:"source-port,omitempty"`
				} `tfsdk:"network_delay" yaml:"network-delay,omitempty"`

				Network_dns *struct {
					Dns_domain_name *string `tfsdk:"dns_domain_name" yaml:"dns-domain-name,omitempty"`

					Dns_ip *string `tfsdk:"dns_ip" yaml:"dns-ip,omitempty"`

					Dns_server *string `tfsdk:"dns_server" yaml:"dns-server,omitempty"`
				} `tfsdk:"network_dns" yaml:"network-dns,omitempty"`

				Network_down *struct {
					Device *string `tfsdk:"device" yaml:"device,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`
				} `tfsdk:"network_down" yaml:"network-down,omitempty"`

				Network_duplicate *struct {
					Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

					Device *string `tfsdk:"device" yaml:"device,omitempty"`

					Egress_port *string `tfsdk:"egress_port" yaml:"egress-port,omitempty"`

					Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

					Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

					Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`

					Percent *string `tfsdk:"percent" yaml:"percent,omitempty"`

					Source_port *string `tfsdk:"source_port" yaml:"source-port,omitempty"`
				} `tfsdk:"network_duplicate" yaml:"network-duplicate,omitempty"`

				Network_flood *struct {
					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

					Parallel *int64 `tfsdk:"parallel" yaml:"parallel,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Rate *string `tfsdk:"rate" yaml:"rate,omitempty"`
				} `tfsdk:"network_flood" yaml:"network-flood,omitempty"`

				Network_loss *struct {
					Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

					Device *string `tfsdk:"device" yaml:"device,omitempty"`

					Egress_port *string `tfsdk:"egress_port" yaml:"egress-port,omitempty"`

					Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

					Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

					Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`

					Percent *string `tfsdk:"percent" yaml:"percent,omitempty"`

					Source_port *string `tfsdk:"source_port" yaml:"source-port,omitempty"`
				} `tfsdk:"network_loss" yaml:"network-loss,omitempty"`

				Network_partition *struct {
					Accept_tcp_flags *string `tfsdk:"accept_tcp_flags" yaml:"accept-tcp-flags,omitempty"`

					Device *string `tfsdk:"device" yaml:"device,omitempty"`

					Direction *string `tfsdk:"direction" yaml:"direction,omitempty"`

					Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

					Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

					Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`
				} `tfsdk:"network_partition" yaml:"network-partition,omitempty"`

				Process *struct {
					Process *string `tfsdk:"process" yaml:"process,omitempty"`

					RecoverCmd *string `tfsdk:"recover_cmd" yaml:"recoverCmd,omitempty"`

					Signal *int64 `tfsdk:"signal" yaml:"signal,omitempty"`
				} `tfsdk:"process" yaml:"process,omitempty"`

				Redis_cacheLimit *struct {
					Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

					CacheSize *string `tfsdk:"cache_size" yaml:"cacheSize,omitempty"`

					Password *string `tfsdk:"password" yaml:"password,omitempty"`

					Percent *string `tfsdk:"percent" yaml:"percent,omitempty"`
				} `tfsdk:"redis_cache_limit" yaml:"redis-cacheLimit,omitempty"`

				Redis_expiration *struct {
					Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

					Expiration *string `tfsdk:"expiration" yaml:"expiration,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Option *string `tfsdk:"option" yaml:"option,omitempty"`

					Password *string `tfsdk:"password" yaml:"password,omitempty"`
				} `tfsdk:"redis_expiration" yaml:"redis-expiration,omitempty"`

				Redis_penetration *struct {
					Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

					Password *string `tfsdk:"password" yaml:"password,omitempty"`

					RequestNum *int64 `tfsdk:"request_num" yaml:"requestNum,omitempty"`
				} `tfsdk:"redis_penetration" yaml:"redis-penetration,omitempty"`

				Redis_restart *struct {
					Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

					Conf *string `tfsdk:"conf" yaml:"conf,omitempty"`

					FlushConfig *bool `tfsdk:"flush_config" yaml:"flushConfig,omitempty"`

					Password *string `tfsdk:"password" yaml:"password,omitempty"`

					RedisPath *bool `tfsdk:"redis_path" yaml:"redisPath,omitempty"`
				} `tfsdk:"redis_restart" yaml:"redis-restart,omitempty"`

				Redis_stop *struct {
					Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

					Conf *string `tfsdk:"conf" yaml:"conf,omitempty"`

					FlushConfig *bool `tfsdk:"flush_config" yaml:"flushConfig,omitempty"`

					Password *string `tfsdk:"password" yaml:"password,omitempty"`

					RedisPath *bool `tfsdk:"redis_path" yaml:"redisPath,omitempty"`
				} `tfsdk:"redis_stop" yaml:"redis-stop,omitempty"`

				Selector *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

					ExpressionSelectors *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

					FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

					LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					PhysicalMachines *map[string][]string `tfsdk:"physical_machines" yaml:"physicalMachines,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				Stress_cpu *struct {
					Load *int64 `tfsdk:"load" yaml:"load,omitempty"`

					Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

					Workers *int64 `tfsdk:"workers" yaml:"workers,omitempty"`
				} `tfsdk:"stress_cpu" yaml:"stress-cpu,omitempty"`

				Stress_mem *struct {
					Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

					Size *string `tfsdk:"size" yaml:"size,omitempty"`
				} `tfsdk:"stress_mem" yaml:"stress-mem,omitempty"`

				Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`

				User_defined *struct {
					AttackCmd *string `tfsdk:"attack_cmd" yaml:"attackCmd,omitempty"`

					RecoverCmd *string `tfsdk:"recover_cmd" yaml:"recoverCmd,omitempty"`
				} `tfsdk:"user_defined" yaml:"user_defined,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				Vm *struct {
					Vm_name *string `tfsdk:"vm_name" yaml:"vm-name,omitempty"`
				} `tfsdk:"vm" yaml:"vm,omitempty"`
			} `tfsdk:"physicalmachine_chaos" yaml:"physicalmachineChaos,omitempty"`

			PodChaos *struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				GracePeriod *int64 `tfsdk:"grace_period" yaml:"gracePeriod,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				Selector *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

					ExpressionSelectors *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

					FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

					LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

					Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

					PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

					Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"pod_chaos" yaml:"podChaos,omitempty"`

			Schedule *struct {
				AwsChaos *struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					AwsRegion *string `tfsdk:"aws_region" yaml:"awsRegion,omitempty"`

					DeviceName *string `tfsdk:"device_name" yaml:"deviceName,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					Ec2Instance *string `tfsdk:"ec2_instance" yaml:"ec2Instance,omitempty"`

					Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

					VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
				} `tfsdk:"aws_chaos" yaml:"awsChaos,omitempty"`

				AzureChaos *struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					DiskName *string `tfsdk:"disk_name" yaml:"diskName,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					Lun *int64 `tfsdk:"lun" yaml:"lun,omitempty"`

					ResourceGroupName *string `tfsdk:"resource_group_name" yaml:"resourceGroupName,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

					SubscriptionID *string `tfsdk:"subscription_id" yaml:"subscriptionID,omitempty"`

					VmName *string `tfsdk:"vm_name" yaml:"vmName,omitempty"`
				} `tfsdk:"azure_chaos" yaml:"azureChaos,omitempty"`

				BlockChaos *struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

					Delay *struct {
						Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

						Jitter *string `tfsdk:"jitter" yaml:"jitter,omitempty"`

						Latency *string `tfsdk:"latency" yaml:"latency,omitempty"`
					} `tfsdk:"delay" yaml:"delay,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

						ExpressionSelectors *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

						FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

						LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

						Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

						PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

						Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`
				} `tfsdk:"block_chaos" yaml:"blockChaos,omitempty"`

				ConcurrencyPolicy *string `tfsdk:"concurrency_policy" yaml:"concurrencyPolicy,omitempty"`

				DnsChaos *struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					Patterns *[]string `tfsdk:"patterns" yaml:"patterns,omitempty"`

					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

						ExpressionSelectors *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

						FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

						LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

						Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

						PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

						Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"dns_chaos" yaml:"dnsChaos,omitempty"`

				GcpChaos *struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					DeviceNames *[]string `tfsdk:"device_names" yaml:"deviceNames,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					Instance *string `tfsdk:"instance" yaml:"instance,omitempty"`

					Project *string `tfsdk:"project" yaml:"project,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

					Zone *string `tfsdk:"zone" yaml:"zone,omitempty"`
				} `tfsdk:"gcp_chaos" yaml:"gcpChaos,omitempty"`

				HistoryLimit *int64 `tfsdk:"history_limit" yaml:"historyLimit,omitempty"`

				HttpChaos *struct {
					Abort *bool `tfsdk:"abort" yaml:"abort,omitempty"`

					Code *int64 `tfsdk:"code" yaml:"code,omitempty"`

					Delay *string `tfsdk:"delay" yaml:"delay,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					Method *string `tfsdk:"method" yaml:"method,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					Patch *struct {
						Body *struct {
							Type *string `tfsdk:"type" yaml:"type,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"body" yaml:"body,omitempty"`

						Headers *[]string `tfsdk:"headers" yaml:"headers,omitempty"`

						Queries *[]string `tfsdk:"queries" yaml:"queries,omitempty"`
					} `tfsdk:"patch" yaml:"patch,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Replace *struct {
						Body *string `tfsdk:"body" yaml:"body,omitempty"`

						Code *int64 `tfsdk:"code" yaml:"code,omitempty"`

						Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

						Method *string `tfsdk:"method" yaml:"method,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Queries *map[string]string `tfsdk:"queries" yaml:"queries,omitempty"`
					} `tfsdk:"replace" yaml:"replace,omitempty"`

					Request_headers *map[string]string `tfsdk:"request_headers" yaml:"request_headers,omitempty"`

					Response_headers *map[string]string `tfsdk:"response_headers" yaml:"response_headers,omitempty"`

					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

						ExpressionSelectors *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

						FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

						LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

						Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

						PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

						Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`

					Target *string `tfsdk:"target" yaml:"target,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"http_chaos" yaml:"httpChaos,omitempty"`

				IoChaos *struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					Attr *struct {
						Atime *struct {
							Nsec *int64 `tfsdk:"nsec" yaml:"nsec,omitempty"`

							Sec *int64 `tfsdk:"sec" yaml:"sec,omitempty"`
						} `tfsdk:"atime" yaml:"atime,omitempty"`

						Blocks *int64 `tfsdk:"blocks" yaml:"blocks,omitempty"`

						Ctime *struct {
							Nsec *int64 `tfsdk:"nsec" yaml:"nsec,omitempty"`

							Sec *int64 `tfsdk:"sec" yaml:"sec,omitempty"`
						} `tfsdk:"ctime" yaml:"ctime,omitempty"`

						Gid *int64 `tfsdk:"gid" yaml:"gid,omitempty"`

						Ino *int64 `tfsdk:"ino" yaml:"ino,omitempty"`

						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						Mtime *struct {
							Nsec *int64 `tfsdk:"nsec" yaml:"nsec,omitempty"`

							Sec *int64 `tfsdk:"sec" yaml:"sec,omitempty"`
						} `tfsdk:"mtime" yaml:"mtime,omitempty"`

						Nlink *int64 `tfsdk:"nlink" yaml:"nlink,omitempty"`

						Perm *int64 `tfsdk:"perm" yaml:"perm,omitempty"`

						Rdev *int64 `tfsdk:"rdev" yaml:"rdev,omitempty"`

						Size *int64 `tfsdk:"size" yaml:"size,omitempty"`

						Uid *int64 `tfsdk:"uid" yaml:"uid,omitempty"`
					} `tfsdk:"attr" yaml:"attr,omitempty"`

					ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

					Delay *string `tfsdk:"delay" yaml:"delay,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					Errno *int64 `tfsdk:"errno" yaml:"errno,omitempty"`

					Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

					Mistake *struct {
						Filling *string `tfsdk:"filling" yaml:"filling,omitempty"`

						MaxLength *int64 `tfsdk:"max_length" yaml:"maxLength,omitempty"`

						MaxOccurrences *int64 `tfsdk:"max_occurrences" yaml:"maxOccurrences,omitempty"`
					} `tfsdk:"mistake" yaml:"mistake,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Percent *int64 `tfsdk:"percent" yaml:"percent,omitempty"`

					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

						ExpressionSelectors *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

						FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

						LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

						Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

						PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

						Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					VolumePath *string `tfsdk:"volume_path" yaml:"volumePath,omitempty"`
				} `tfsdk:"io_chaos" yaml:"ioChaos,omitempty"`

				JvmChaos *struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					Class *string `tfsdk:"class" yaml:"class,omitempty"`

					ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

					CpuCount *int64 `tfsdk:"cpu_count" yaml:"cpuCount,omitempty"`

					Database *string `tfsdk:"database" yaml:"database,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					Exception *string `tfsdk:"exception" yaml:"exception,omitempty"`

					Latency *int64 `tfsdk:"latency" yaml:"latency,omitempty"`

					MemType *string `tfsdk:"mem_type" yaml:"memType,omitempty"`

					Method *string `tfsdk:"method" yaml:"method,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					MysqlConnectorVersion *string `tfsdk:"mysql_connector_version" yaml:"mysqlConnectorVersion,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					RuleData *string `tfsdk:"rule_data" yaml:"ruleData,omitempty"`

					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

						ExpressionSelectors *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

						FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

						LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

						Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

						PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

						Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`

					SqlType *string `tfsdk:"sql_type" yaml:"sqlType,omitempty"`

					Table *string `tfsdk:"table" yaml:"table,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"jvm_chaos" yaml:"jvmChaos,omitempty"`

				KernelChaos *struct {
					ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					FailKernRequest *struct {
						Callchain *[]struct {
							Funcname *string `tfsdk:"funcname" yaml:"funcname,omitempty"`

							Parameters *string `tfsdk:"parameters" yaml:"parameters,omitempty"`

							Predicate *string `tfsdk:"predicate" yaml:"predicate,omitempty"`
						} `tfsdk:"callchain" yaml:"callchain,omitempty"`

						Failtype *int64 `tfsdk:"failtype" yaml:"failtype,omitempty"`

						Headers *[]string `tfsdk:"headers" yaml:"headers,omitempty"`

						Probability *int64 `tfsdk:"probability" yaml:"probability,omitempty"`

						Times *int64 `tfsdk:"times" yaml:"times,omitempty"`
					} `tfsdk:"fail_kern_request" yaml:"failKernRequest,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

						ExpressionSelectors *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

						FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

						LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

						Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

						PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

						Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"kernel_chaos" yaml:"kernelChaos,omitempty"`

				NetworkChaos *struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					Bandwidth *struct {
						Buffer *int64 `tfsdk:"buffer" yaml:"buffer,omitempty"`

						Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`

						Minburst *int64 `tfsdk:"minburst" yaml:"minburst,omitempty"`

						Peakrate *int64 `tfsdk:"peakrate" yaml:"peakrate,omitempty"`

						Rate *string `tfsdk:"rate" yaml:"rate,omitempty"`
					} `tfsdk:"bandwidth" yaml:"bandwidth,omitempty"`

					Corrupt *struct {
						Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

						Corrupt *string `tfsdk:"corrupt" yaml:"corrupt,omitempty"`
					} `tfsdk:"corrupt" yaml:"corrupt,omitempty"`

					Delay *struct {
						Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

						Jitter *string `tfsdk:"jitter" yaml:"jitter,omitempty"`

						Latency *string `tfsdk:"latency" yaml:"latency,omitempty"`

						Reorder *struct {
							Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

							Gap *int64 `tfsdk:"gap" yaml:"gap,omitempty"`

							Reorder *string `tfsdk:"reorder" yaml:"reorder,omitempty"`
						} `tfsdk:"reorder" yaml:"reorder,omitempty"`
					} `tfsdk:"delay" yaml:"delay,omitempty"`

					Device *string `tfsdk:"device" yaml:"device,omitempty"`

					Direction *string `tfsdk:"direction" yaml:"direction,omitempty"`

					Duplicate *struct {
						Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

						Duplicate *string `tfsdk:"duplicate" yaml:"duplicate,omitempty"`
					} `tfsdk:"duplicate" yaml:"duplicate,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					ExternalTargets *[]string `tfsdk:"external_targets" yaml:"externalTargets,omitempty"`

					Loss *struct {
						Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

						Loss *string `tfsdk:"loss" yaml:"loss,omitempty"`
					} `tfsdk:"loss" yaml:"loss,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

						ExpressionSelectors *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

						FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

						LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

						Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

						PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

						Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`

					Target *struct {
						Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

						Selector *struct {
							AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

							ExpressionSelectors *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

							FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

							LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

							Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

							NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

							Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

							PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

							Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
						} `tfsdk:"selector" yaml:"selector,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"target" yaml:"target,omitempty"`

					TargetDevice *string `tfsdk:"target_device" yaml:"targetDevice,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"network_chaos" yaml:"networkChaos,omitempty"`

				PhysicalmachineChaos *struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					Address *[]string `tfsdk:"address" yaml:"address,omitempty"`

					Clock *struct {
						Clock_ids_slice *string `tfsdk:"clock_ids_slice" yaml:"clock-ids-slice,omitempty"`

						Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

						Time_offset *string `tfsdk:"time_offset" yaml:"time-offset,omitempty"`
					} `tfsdk:"clock" yaml:"clock,omitempty"`

					Disk_fill *struct {
						Fill_by_fallocate *bool `tfsdk:"fill_by_fallocate" yaml:"fill-by-fallocate,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Size *string `tfsdk:"size" yaml:"size,omitempty"`
					} `tfsdk:"disk_fill" yaml:"disk-fill,omitempty"`

					Disk_read_payload *struct {
						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Payload_process_num *int64 `tfsdk:"payload_process_num" yaml:"payload-process-num,omitempty"`

						Size *string `tfsdk:"size" yaml:"size,omitempty"`
					} `tfsdk:"disk_read_payload" yaml:"disk-read-payload,omitempty"`

					Disk_write_payload *struct {
						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Payload_process_num *int64 `tfsdk:"payload_process_num" yaml:"payload-process-num,omitempty"`

						Size *string `tfsdk:"size" yaml:"size,omitempty"`
					} `tfsdk:"disk_write_payload" yaml:"disk-write-payload,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					File_append *struct {
						Count *int64 `tfsdk:"count" yaml:"count,omitempty"`

						Data *string `tfsdk:"data" yaml:"data,omitempty"`

						File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`
					} `tfsdk:"file_append" yaml:"file-append,omitempty"`

					File_create *struct {
						Dir_name *string `tfsdk:"dir_name" yaml:"dir-name,omitempty"`

						File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`
					} `tfsdk:"file_create" yaml:"file-create,omitempty"`

					File_delete *struct {
						Dir_name *string `tfsdk:"dir_name" yaml:"dir-name,omitempty"`

						File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`
					} `tfsdk:"file_delete" yaml:"file-delete,omitempty"`

					File_modify *struct {
						File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`

						Privilege *int64 `tfsdk:"privilege" yaml:"privilege,omitempty"`
					} `tfsdk:"file_modify" yaml:"file-modify,omitempty"`

					File_rename *struct {
						Dest_file *string `tfsdk:"dest_file" yaml:"dest-file,omitempty"`

						Source_file *string `tfsdk:"source_file" yaml:"source-file,omitempty"`
					} `tfsdk:"file_rename" yaml:"file-rename,omitempty"`

					File_replace *struct {
						Dest_string *string `tfsdk:"dest_string" yaml:"dest-string,omitempty"`

						File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`

						Line *int64 `tfsdk:"line" yaml:"line,omitempty"`

						Origin_string *string `tfsdk:"origin_string" yaml:"origin-string,omitempty"`
					} `tfsdk:"file_replace" yaml:"file-replace,omitempty"`

					Http_abort *struct {
						Code *string `tfsdk:"code" yaml:"code,omitempty"`

						Method *string `tfsdk:"method" yaml:"method,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

						Proxy_ports *[]string `tfsdk:"proxy_ports" yaml:"proxy_ports,omitempty"`

						Target *string `tfsdk:"target" yaml:"target,omitempty"`
					} `tfsdk:"http_abort" yaml:"http-abort,omitempty"`

					Http_config *struct {
						File_path *string `tfsdk:"file_path" yaml:"file_path,omitempty"`
					} `tfsdk:"http_config" yaml:"http-config,omitempty"`

					Http_delay *struct {
						Code *string `tfsdk:"code" yaml:"code,omitempty"`

						Delay *string `tfsdk:"delay" yaml:"delay,omitempty"`

						Method *string `tfsdk:"method" yaml:"method,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

						Proxy_ports *[]string `tfsdk:"proxy_ports" yaml:"proxy_ports,omitempty"`

						Target *string `tfsdk:"target" yaml:"target,omitempty"`
					} `tfsdk:"http_delay" yaml:"http-delay,omitempty"`

					Http_request *struct {
						Count *int64 `tfsdk:"count" yaml:"count,omitempty"`

						Enable_conn_pool *bool `tfsdk:"enable_conn_pool" yaml:"enable-conn-pool,omitempty"`

						Url *string `tfsdk:"url" yaml:"url,omitempty"`
					} `tfsdk:"http_request" yaml:"http-request,omitempty"`

					Jvm_exception *struct {
						Class *string `tfsdk:"class" yaml:"class,omitempty"`

						Exception *string `tfsdk:"exception" yaml:"exception,omitempty"`

						Method *string `tfsdk:"method" yaml:"method,omitempty"`

						Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"jvm_exception" yaml:"jvm-exception,omitempty"`

					Jvm_gc *struct {
						Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"jvm_gc" yaml:"jvm-gc,omitempty"`

					Jvm_latency *struct {
						Class *string `tfsdk:"class" yaml:"class,omitempty"`

						Latency *int64 `tfsdk:"latency" yaml:"latency,omitempty"`

						Method *string `tfsdk:"method" yaml:"method,omitempty"`

						Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"jvm_latency" yaml:"jvm-latency,omitempty"`

					Jvm_mysql *struct {
						Database *string `tfsdk:"database" yaml:"database,omitempty"`

						Exception *string `tfsdk:"exception" yaml:"exception,omitempty"`

						Latency *int64 `tfsdk:"latency" yaml:"latency,omitempty"`

						MysqlConnectorVersion *string `tfsdk:"mysql_connector_version" yaml:"mysqlConnectorVersion,omitempty"`

						Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

						SqlType *string `tfsdk:"sql_type" yaml:"sqlType,omitempty"`

						Table *string `tfsdk:"table" yaml:"table,omitempty"`
					} `tfsdk:"jvm_mysql" yaml:"jvm-mysql,omitempty"`

					Jvm_return *struct {
						Class *string `tfsdk:"class" yaml:"class,omitempty"`

						Method *string `tfsdk:"method" yaml:"method,omitempty"`

						Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"jvm_return" yaml:"jvm-return,omitempty"`

					Jvm_rule_data *struct {
						Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

						Rule_data *string `tfsdk:"rule_data" yaml:"rule-data,omitempty"`
					} `tfsdk:"jvm_rule_data" yaml:"jvm-rule-data,omitempty"`

					Jvm_stress *struct {
						Cpu_count *int64 `tfsdk:"cpu_count" yaml:"cpu-count,omitempty"`

						Mem_type *string `tfsdk:"mem_type" yaml:"mem-type,omitempty"`

						Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"jvm_stress" yaml:"jvm-stress,omitempty"`

					Kafka_fill *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						MaxBytes *int64 `tfsdk:"max_bytes" yaml:"maxBytes,omitempty"`

						MessageSize *int64 `tfsdk:"message_size" yaml:"messageSize,omitempty"`

						Password *string `tfsdk:"password" yaml:"password,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

						ReloadCommand *string `tfsdk:"reload_command" yaml:"reloadCommand,omitempty"`

						Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`

						Username *string `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"kafka_fill" yaml:"kafka-fill,omitempty"`

					Kafka_flood *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						MessageSize *int64 `tfsdk:"message_size" yaml:"messageSize,omitempty"`

						Password *string `tfsdk:"password" yaml:"password,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

						Threads *int64 `tfsdk:"threads" yaml:"threads,omitempty"`

						Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`

						Username *string `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"kafka_flood" yaml:"kafka-flood,omitempty"`

					Kafka_io *struct {
						ConfigFile *string `tfsdk:"config_file" yaml:"configFile,omitempty"`

						NonReadable *bool `tfsdk:"non_readable" yaml:"nonReadable,omitempty"`

						NonWritable *bool `tfsdk:"non_writable" yaml:"nonWritable,omitempty"`

						Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`
					} `tfsdk:"kafka_io" yaml:"kafka-io,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					Network_bandwidth *struct {
						Buffer *int64 `tfsdk:"buffer" yaml:"buffer,omitempty"`

						Device *string `tfsdk:"device" yaml:"device,omitempty"`

						Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

						Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

						Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`

						Minburst *int64 `tfsdk:"minburst" yaml:"minburst,omitempty"`

						Peakrate *int64 `tfsdk:"peakrate" yaml:"peakrate,omitempty"`

						Rate *string `tfsdk:"rate" yaml:"rate,omitempty"`
					} `tfsdk:"network_bandwidth" yaml:"network-bandwidth,omitempty"`

					Network_corrupt *struct {
						Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

						Device *string `tfsdk:"device" yaml:"device,omitempty"`

						Egress_port *string `tfsdk:"egress_port" yaml:"egress-port,omitempty"`

						Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

						Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

						Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`

						Percent *string `tfsdk:"percent" yaml:"percent,omitempty"`

						Source_port *string `tfsdk:"source_port" yaml:"source-port,omitempty"`
					} `tfsdk:"network_corrupt" yaml:"network-corrupt,omitempty"`

					Network_delay *struct {
						Accept_tcp_flags *string `tfsdk:"accept_tcp_flags" yaml:"accept-tcp-flags,omitempty"`

						Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

						Device *string `tfsdk:"device" yaml:"device,omitempty"`

						Egress_port *string `tfsdk:"egress_port" yaml:"egress-port,omitempty"`

						Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

						Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

						Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`

						Jitter *string `tfsdk:"jitter" yaml:"jitter,omitempty"`

						Latency *string `tfsdk:"latency" yaml:"latency,omitempty"`

						Source_port *string `tfsdk:"source_port" yaml:"source-port,omitempty"`
					} `tfsdk:"network_delay" yaml:"network-delay,omitempty"`

					Network_dns *struct {
						Dns_domain_name *string `tfsdk:"dns_domain_name" yaml:"dns-domain-name,omitempty"`

						Dns_ip *string `tfsdk:"dns_ip" yaml:"dns-ip,omitempty"`

						Dns_server *string `tfsdk:"dns_server" yaml:"dns-server,omitempty"`
					} `tfsdk:"network_dns" yaml:"network-dns,omitempty"`

					Network_down *struct {
						Device *string `tfsdk:"device" yaml:"device,omitempty"`

						Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`
					} `tfsdk:"network_down" yaml:"network-down,omitempty"`

					Network_duplicate *struct {
						Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

						Device *string `tfsdk:"device" yaml:"device,omitempty"`

						Egress_port *string `tfsdk:"egress_port" yaml:"egress-port,omitempty"`

						Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

						Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

						Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`

						Percent *string `tfsdk:"percent" yaml:"percent,omitempty"`

						Source_port *string `tfsdk:"source_port" yaml:"source-port,omitempty"`
					} `tfsdk:"network_duplicate" yaml:"network-duplicate,omitempty"`

					Network_flood *struct {
						Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

						Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

						Parallel *int64 `tfsdk:"parallel" yaml:"parallel,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Rate *string `tfsdk:"rate" yaml:"rate,omitempty"`
					} `tfsdk:"network_flood" yaml:"network-flood,omitempty"`

					Network_loss *struct {
						Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

						Device *string `tfsdk:"device" yaml:"device,omitempty"`

						Egress_port *string `tfsdk:"egress_port" yaml:"egress-port,omitempty"`

						Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

						Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

						Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`

						Percent *string `tfsdk:"percent" yaml:"percent,omitempty"`

						Source_port *string `tfsdk:"source_port" yaml:"source-port,omitempty"`
					} `tfsdk:"network_loss" yaml:"network-loss,omitempty"`

					Network_partition *struct {
						Accept_tcp_flags *string `tfsdk:"accept_tcp_flags" yaml:"accept-tcp-flags,omitempty"`

						Device *string `tfsdk:"device" yaml:"device,omitempty"`

						Direction *string `tfsdk:"direction" yaml:"direction,omitempty"`

						Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

						Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

						Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`
					} `tfsdk:"network_partition" yaml:"network-partition,omitempty"`

					Process *struct {
						Process *string `tfsdk:"process" yaml:"process,omitempty"`

						RecoverCmd *string `tfsdk:"recover_cmd" yaml:"recoverCmd,omitempty"`

						Signal *int64 `tfsdk:"signal" yaml:"signal,omitempty"`
					} `tfsdk:"process" yaml:"process,omitempty"`

					Redis_cacheLimit *struct {
						Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

						CacheSize *string `tfsdk:"cache_size" yaml:"cacheSize,omitempty"`

						Password *string `tfsdk:"password" yaml:"password,omitempty"`

						Percent *string `tfsdk:"percent" yaml:"percent,omitempty"`
					} `tfsdk:"redis_cache_limit" yaml:"redis-cacheLimit,omitempty"`

					Redis_expiration *struct {
						Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

						Expiration *string `tfsdk:"expiration" yaml:"expiration,omitempty"`

						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Option *string `tfsdk:"option" yaml:"option,omitempty"`

						Password *string `tfsdk:"password" yaml:"password,omitempty"`
					} `tfsdk:"redis_expiration" yaml:"redis-expiration,omitempty"`

					Redis_penetration *struct {
						Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

						Password *string `tfsdk:"password" yaml:"password,omitempty"`

						RequestNum *int64 `tfsdk:"request_num" yaml:"requestNum,omitempty"`
					} `tfsdk:"redis_penetration" yaml:"redis-penetration,omitempty"`

					Redis_restart *struct {
						Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

						Conf *string `tfsdk:"conf" yaml:"conf,omitempty"`

						FlushConfig *bool `tfsdk:"flush_config" yaml:"flushConfig,omitempty"`

						Password *string `tfsdk:"password" yaml:"password,omitempty"`

						RedisPath *bool `tfsdk:"redis_path" yaml:"redisPath,omitempty"`
					} `tfsdk:"redis_restart" yaml:"redis-restart,omitempty"`

					Redis_stop *struct {
						Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

						Conf *string `tfsdk:"conf" yaml:"conf,omitempty"`

						FlushConfig *bool `tfsdk:"flush_config" yaml:"flushConfig,omitempty"`

						Password *string `tfsdk:"password" yaml:"password,omitempty"`

						RedisPath *bool `tfsdk:"redis_path" yaml:"redisPath,omitempty"`
					} `tfsdk:"redis_stop" yaml:"redis-stop,omitempty"`

					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

						ExpressionSelectors *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

						FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

						LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						PhysicalMachines *map[string][]string `tfsdk:"physical_machines" yaml:"physicalMachines,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`

					Stress_cpu *struct {
						Load *int64 `tfsdk:"load" yaml:"load,omitempty"`

						Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

						Workers *int64 `tfsdk:"workers" yaml:"workers,omitempty"`
					} `tfsdk:"stress_cpu" yaml:"stress-cpu,omitempty"`

					Stress_mem *struct {
						Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

						Size *string `tfsdk:"size" yaml:"size,omitempty"`
					} `tfsdk:"stress_mem" yaml:"stress-mem,omitempty"`

					Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`

					User_defined *struct {
						AttackCmd *string `tfsdk:"attack_cmd" yaml:"attackCmd,omitempty"`

						RecoverCmd *string `tfsdk:"recover_cmd" yaml:"recoverCmd,omitempty"`
					} `tfsdk:"user_defined" yaml:"user_defined,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					Vm *struct {
						Vm_name *string `tfsdk:"vm_name" yaml:"vm-name,omitempty"`
					} `tfsdk:"vm" yaml:"vm,omitempty"`
				} `tfsdk:"physicalmachine_chaos" yaml:"physicalmachineChaos,omitempty"`

				PodChaos *struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					GracePeriod *int64 `tfsdk:"grace_period" yaml:"gracePeriod,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

						ExpressionSelectors *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

						FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

						LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

						Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

						PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

						Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"pod_chaos" yaml:"podChaos,omitempty"`

				Schedule *string `tfsdk:"schedule" yaml:"schedule,omitempty"`

				StartingDeadlineSeconds *int64 `tfsdk:"starting_deadline_seconds" yaml:"startingDeadlineSeconds,omitempty"`

				StressChaos *struct {
					ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

						ExpressionSelectors *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

						FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

						LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

						Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

						PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

						Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`

					StressngStressors *string `tfsdk:"stressng_stressors" yaml:"stressngStressors,omitempty"`

					Stressors *struct {
						Cpu *struct {
							Load *int64 `tfsdk:"load" yaml:"load,omitempty"`

							Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

							Workers *int64 `tfsdk:"workers" yaml:"workers,omitempty"`
						} `tfsdk:"cpu" yaml:"cpu,omitempty"`

						Memory *struct {
							OomScoreAdj *int64 `tfsdk:"oom_score_adj" yaml:"oomScoreAdj,omitempty"`

							Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

							Size *string `tfsdk:"size" yaml:"size,omitempty"`

							Workers *int64 `tfsdk:"workers" yaml:"workers,omitempty"`
						} `tfsdk:"memory" yaml:"memory,omitempty"`
					} `tfsdk:"stressors" yaml:"stressors,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"stress_chaos" yaml:"stressChaos,omitempty"`

				TimeChaos *struct {
					ClockIds *[]string `tfsdk:"clock_ids" yaml:"clockIds,omitempty"`

					ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

					Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					Selector *struct {
						AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

						ExpressionSelectors *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

						FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

						LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

						Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

						PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

						Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`

					TimeOffset *string `tfsdk:"time_offset" yaml:"timeOffset,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"time_chaos" yaml:"timeChaos,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"schedule" yaml:"schedule,omitempty"`

			StatusCheck *struct {
				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				Http *struct {
					Body *string `tfsdk:"body" yaml:"body,omitempty"`

					Criteria *struct {
						StatusCode *string `tfsdk:"status_code" yaml:"statusCode,omitempty"`
					} `tfsdk:"criteria" yaml:"criteria,omitempty"`

					Headers *map[string][]string `tfsdk:"headers" yaml:"headers,omitempty"`

					Method *string `tfsdk:"method" yaml:"method,omitempty"`

					Url *string `tfsdk:"url" yaml:"url,omitempty"`
				} `tfsdk:"http" yaml:"http,omitempty"`

				IntervalSeconds *int64 `tfsdk:"interval_seconds" yaml:"intervalSeconds,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				RecordsHistoryLimit *int64 `tfsdk:"records_history_limit" yaml:"recordsHistoryLimit,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"status_check" yaml:"statusCheck,omitempty"`

			StressChaos *struct {
				ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				Selector *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

					ExpressionSelectors *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

					FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

					LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

					Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

					PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

					Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				StressngStressors *string `tfsdk:"stressng_stressors" yaml:"stressngStressors,omitempty"`

				Stressors *struct {
					Cpu *struct {
						Load *int64 `tfsdk:"load" yaml:"load,omitempty"`

						Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

						Workers *int64 `tfsdk:"workers" yaml:"workers,omitempty"`
					} `tfsdk:"cpu" yaml:"cpu,omitempty"`

					Memory *struct {
						OomScoreAdj *int64 `tfsdk:"oom_score_adj" yaml:"oomScoreAdj,omitempty"`

						Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

						Size *string `tfsdk:"size" yaml:"size,omitempty"`

						Workers *int64 `tfsdk:"workers" yaml:"workers,omitempty"`
					} `tfsdk:"memory" yaml:"memory,omitempty"`
				} `tfsdk:"stressors" yaml:"stressors,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"stress_chaos" yaml:"stressChaos,omitempty"`

			Task *struct {
				Container *struct {
					Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`

						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

							FieldRef *struct {
								ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

								FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
							} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

							ResourceFieldRef *struct {
								ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

								Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

								Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
							} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

							SecretKeyRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

					EnvFrom *[]struct {
						ConfigMapRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

						Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
					} `tfsdk:"env_from" yaml:"envFrom,omitempty"`

					Image *string `tfsdk:"image" yaml:"image,omitempty"`

					ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

					Lifecycle *struct {
						PostStart *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
							} `tfsdk:"exec" yaml:"exec,omitempty"`

							HttpGet *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								HttpHeaders *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

								Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
							} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
						} `tfsdk:"post_start" yaml:"postStart,omitempty"`

						PreStop *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
							} `tfsdk:"exec" yaml:"exec,omitempty"`

							HttpGet *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								HttpHeaders *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

								Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
							} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
						} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
					} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

					LivenessProbe *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`

						FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

						Grpc *struct {
							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

						HttpGet *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							HttpHeaders *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

						PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

						SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

						TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

						TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
					} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Ports *[]struct {
						ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`

						HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

						HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
					} `tfsdk:"ports" yaml:"ports,omitempty"`

					ReadinessProbe *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`

						FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

						Grpc *struct {
							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

						HttpGet *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							HttpHeaders *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

						PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

						SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

						TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

						TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
					} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

					Resources *struct {
						Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

						Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
					} `tfsdk:"resources" yaml:"resources,omitempty"`

					SecurityContext *struct {
						AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

						Capabilities *struct {
							Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

							Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
						} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

						Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

						ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

						ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

						RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

						RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

						RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

						SeLinuxOptions *struct {
							Level *string `tfsdk:"level" yaml:"level,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Type *string `tfsdk:"type" yaml:"type,omitempty"`

							User *string `tfsdk:"user" yaml:"user,omitempty"`
						} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

						SeccompProfile *struct {
							LocalhostProfile *string `tfsdk:"localhost_profile" yaml:"localhostProfile,omitempty"`

							Type *string `tfsdk:"type" yaml:"type,omitempty"`
						} `tfsdk:"seccomp_profile" yaml:"seccompProfile,omitempty"`

						WindowsOptions *struct {
							GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

							GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

							HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

							RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
						} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
					} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

					StartupProbe *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`

						FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

						Grpc *struct {
							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

						HttpGet *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							HttpHeaders *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

						PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

						SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

						TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

						TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
					} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

					Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

					StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

					TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`

					TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

					Tty *bool `tfsdk:"tty" yaml:"tty,omitempty"`

					VolumeDevices *[]struct {
						DevicePath *string `tfsdk:"device_path" yaml:"devicePath,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"volume_devices" yaml:"volumeDevices,omitempty"`

					VolumeMounts *[]struct {
						MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

						MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

						SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
					} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

					WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`
				} `tfsdk:"container" yaml:"container,omitempty"`

				Volumes *[]struct {
					AwsElasticBlockStore *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Partition *int64 `tfsdk:"partition" yaml:"partition,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
					} `tfsdk:"aws_elastic_block_store" yaml:"awsElasticBlockStore,omitempty"`

					AzureDisk *struct {
						CachingMode *string `tfsdk:"caching_mode" yaml:"cachingMode,omitempty"`

						DiskName *string `tfsdk:"disk_name" yaml:"diskName,omitempty"`

						DiskURI *string `tfsdk:"disk_uri" yaml:"diskURI,omitempty"`

						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
					} `tfsdk:"azure_disk" yaml:"azureDisk,omitempty"`

					AzureFile *struct {
						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

						ShareName *string `tfsdk:"share_name" yaml:"shareName,omitempty"`
					} `tfsdk:"azure_file" yaml:"azureFile,omitempty"`

					Cephfs *struct {
						Monitors *[]string `tfsdk:"monitors" yaml:"monitors,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SecretFile *string `tfsdk:"secret_file" yaml:"secretFile,omitempty"`

						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`
					} `tfsdk:"cephfs" yaml:"cephfs,omitempty"`

					Cinder *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
					} `tfsdk:"cinder" yaml:"cinder,omitempty"`

					ConfigMap *struct {
						DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`

						Items *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"items" yaml:"items,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Csi *struct {
						Driver *string `tfsdk:"driver" yaml:"driver,omitempty"`

						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						NodePublishSecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"node_publish_secret_ref" yaml:"nodePublishSecretRef,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						VolumeAttributes *map[string]string `tfsdk:"volume_attributes" yaml:"volumeAttributes,omitempty"`
					} `tfsdk:"csi" yaml:"csi,omitempty"`

					DownwardAPI *struct {
						DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`

						Items *[]struct {
							FieldRef *struct {
								ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

								FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
							} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

							Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							ResourceFieldRef *struct {
								ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

								Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

								Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
							} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`
						} `tfsdk:"items" yaml:"items,omitempty"`
					} `tfsdk:"downward_api" yaml:"downwardAPI,omitempty"`

					EmptyDir *struct {
						Medium *string `tfsdk:"medium" yaml:"medium,omitempty"`

						SizeLimit utilities.IntOrString `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`
					} `tfsdk:"empty_dir" yaml:"emptyDir,omitempty"`

					Ephemeral *struct {
						VolumeClaimTemplate *struct {
							Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

							Spec *struct {
								AccessModes *[]string `tfsdk:"access_modes" yaml:"accessModes,omitempty"`

								DataSource *struct {
									ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`
								} `tfsdk:"data_source" yaml:"dataSource,omitempty"`

								DataSourceRef *struct {
									ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

									Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`
								} `tfsdk:"data_source_ref" yaml:"dataSourceRef,omitempty"`

								Resources *struct {
									Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

									Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
								} `tfsdk:"resources" yaml:"resources,omitempty"`

								Selector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"selector" yaml:"selector,omitempty"`

								StorageClassName *string `tfsdk:"storage_class_name" yaml:"storageClassName,omitempty"`

								VolumeMode *string `tfsdk:"volume_mode" yaml:"volumeMode,omitempty"`

								VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`
							} `tfsdk:"spec" yaml:"spec,omitempty"`
						} `tfsdk:"volume_claim_template" yaml:"volumeClaimTemplate,omitempty"`
					} `tfsdk:"ephemeral" yaml:"ephemeral,omitempty"`

					Fc *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Lun *int64 `tfsdk:"lun" yaml:"lun,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						TargetWWNs *[]string `tfsdk:"target_ww_ns" yaml:"targetWWNs,omitempty"`

						Wwids *[]string `tfsdk:"wwids" yaml:"wwids,omitempty"`
					} `tfsdk:"fc" yaml:"fc,omitempty"`

					FlexVolume *struct {
						Driver *string `tfsdk:"driver" yaml:"driver,omitempty"`

						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Options *map[string]string `tfsdk:"options" yaml:"options,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
					} `tfsdk:"flex_volume" yaml:"flexVolume,omitempty"`

					Flocker *struct {
						DatasetName *string `tfsdk:"dataset_name" yaml:"datasetName,omitempty"`

						DatasetUUID *string `tfsdk:"dataset_uuid" yaml:"datasetUUID,omitempty"`
					} `tfsdk:"flocker" yaml:"flocker,omitempty"`

					GcePersistentDisk *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Partition *int64 `tfsdk:"partition" yaml:"partition,omitempty"`

						PdName *string `tfsdk:"pd_name" yaml:"pdName,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
					} `tfsdk:"gce_persistent_disk" yaml:"gcePersistentDisk,omitempty"`

					GitRepo *struct {
						Directory *string `tfsdk:"directory" yaml:"directory,omitempty"`

						Repository *string `tfsdk:"repository" yaml:"repository,omitempty"`

						Revision *string `tfsdk:"revision" yaml:"revision,omitempty"`
					} `tfsdk:"git_repo" yaml:"gitRepo,omitempty"`

					Glusterfs *struct {
						Endpoints *string `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
					} `tfsdk:"glusterfs" yaml:"glusterfs,omitempty"`

					HostPath *struct {
						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"host_path" yaml:"hostPath,omitempty"`

					Iscsi *struct {
						ChapAuthDiscovery *bool `tfsdk:"chap_auth_discovery" yaml:"chapAuthDiscovery,omitempty"`

						ChapAuthSession *bool `tfsdk:"chap_auth_session" yaml:"chapAuthSession,omitempty"`

						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						InitiatorName *string `tfsdk:"initiator_name" yaml:"initiatorName,omitempty"`

						Iqn *string `tfsdk:"iqn" yaml:"iqn,omitempty"`

						IscsiInterface *string `tfsdk:"iscsi_interface" yaml:"iscsiInterface,omitempty"`

						Lun *int64 `tfsdk:"lun" yaml:"lun,omitempty"`

						Portals *[]string `tfsdk:"portals" yaml:"portals,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						TargetPortal *string `tfsdk:"target_portal" yaml:"targetPortal,omitempty"`
					} `tfsdk:"iscsi" yaml:"iscsi,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Nfs *struct {
						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						Server *string `tfsdk:"server" yaml:"server,omitempty"`
					} `tfsdk:"nfs" yaml:"nfs,omitempty"`

					PersistentVolumeClaim *struct {
						ClaimName *string `tfsdk:"claim_name" yaml:"claimName,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
					} `tfsdk:"persistent_volume_claim" yaml:"persistentVolumeClaim,omitempty"`

					PhotonPersistentDisk *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						PdID *string `tfsdk:"pd_id" yaml:"pdID,omitempty"`
					} `tfsdk:"photon_persistent_disk" yaml:"photonPersistentDisk,omitempty"`

					PortworxVolume *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
					} `tfsdk:"portworx_volume" yaml:"portworxVolume,omitempty"`

					Projected *struct {
						DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`

						Sources *[]struct {
							ConfigMap *struct {
								Items *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"items" yaml:"items,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"config_map" yaml:"configMap,omitempty"`

							DownwardAPI *struct {
								Items *[]struct {
									FieldRef *struct {
										ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

										FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
									} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

									Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									ResourceFieldRef *struct {
										ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

										Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

										Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
									} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`
								} `tfsdk:"items" yaml:"items,omitempty"`
							} `tfsdk:"downward_api" yaml:"downwardAPI,omitempty"`

							Secret *struct {
								Items *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"items" yaml:"items,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccountToken *struct {
								Audience *string `tfsdk:"audience" yaml:"audience,omitempty"`

								ExpirationSeconds *int64 `tfsdk:"expiration_seconds" yaml:"expirationSeconds,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"service_account_token" yaml:"serviceAccountToken,omitempty"`
						} `tfsdk:"sources" yaml:"sources,omitempty"`
					} `tfsdk:"projected" yaml:"projected,omitempty"`

					Quobyte *struct {
						Group *string `tfsdk:"group" yaml:"group,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						Registry *string `tfsdk:"registry" yaml:"registry,omitempty"`

						Tenant *string `tfsdk:"tenant" yaml:"tenant,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`

						Volume *string `tfsdk:"volume" yaml:"volume,omitempty"`
					} `tfsdk:"quobyte" yaml:"quobyte,omitempty"`

					Rbd *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Image *string `tfsdk:"image" yaml:"image,omitempty"`

						Keyring *string `tfsdk:"keyring" yaml:"keyring,omitempty"`

						Monitors *[]string `tfsdk:"monitors" yaml:"monitors,omitempty"`

						Pool *string `tfsdk:"pool" yaml:"pool,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`
					} `tfsdk:"rbd" yaml:"rbd,omitempty"`

					ScaleIO *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Gateway *string `tfsdk:"gateway" yaml:"gateway,omitempty"`

						ProtectionDomain *string `tfsdk:"protection_domain" yaml:"protectionDomain,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						SslEnabled *bool `tfsdk:"ssl_enabled" yaml:"sslEnabled,omitempty"`

						StorageMode *string `tfsdk:"storage_mode" yaml:"storageMode,omitempty"`

						StoragePool *string `tfsdk:"storage_pool" yaml:"storagePool,omitempty"`

						System *string `tfsdk:"system" yaml:"system,omitempty"`

						VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`
					} `tfsdk:"scale_io" yaml:"scaleIO,omitempty"`

					Secret *struct {
						DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`

						Items *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"items" yaml:"items,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`

					Storageos *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`

						VolumeNamespace *string `tfsdk:"volume_namespace" yaml:"volumeNamespace,omitempty"`
					} `tfsdk:"storageos" yaml:"storageos,omitempty"`

					VsphereVolume *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						StoragePolicyID *string `tfsdk:"storage_policy_id" yaml:"storagePolicyID,omitempty"`

						StoragePolicyName *string `tfsdk:"storage_policy_name" yaml:"storagePolicyName,omitempty"`

						VolumePath *string `tfsdk:"volume_path" yaml:"volumePath,omitempty"`
					} `tfsdk:"vsphere_volume" yaml:"vsphereVolume,omitempty"`
				} `tfsdk:"volumes" yaml:"volumes,omitempty"`
			} `tfsdk:"task" yaml:"task,omitempty"`

			TemplateType *string `tfsdk:"template_type" yaml:"templateType,omitempty"`

			TimeChaos *struct {
				ClockIds *[]string `tfsdk:"clock_ids" yaml:"clockIds,omitempty"`

				ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

				Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				Selector *struct {
					AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

					ExpressionSelectors *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

					FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

					LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

					Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

					NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

					Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

					PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

					Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
				} `tfsdk:"selector" yaml:"selector,omitempty"`

				TimeOffset *string `tfsdk:"time_offset" yaml:"timeOffset,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"time_chaos" yaml:"timeChaos,omitempty"`
		} `tfsdk:"templates" yaml:"templates,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewChaosMeshOrgWorkflowV1Alpha1Resource() resource.Resource {
	return &ChaosMeshOrgWorkflowV1Alpha1Resource{}
}

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chaos_mesh_org_workflow_v1alpha1"
}

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				Description:         "Spec defines the behavior of a workflow",
				MarkdownDescription: "Spec defines the behavior of a workflow",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"entry": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"templates": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"abort_with_status_check": {
								Description:         "AbortWithStatusCheck describe whether to abort the workflow when the failure threshold of StatusCheck is exceeded. Only used when Type is TypeStatusCheck.",
								MarkdownDescription: "AbortWithStatusCheck describe whether to abort the workflow when the failure threshold of StatusCheck is exceeded. Only used when Type is TypeStatusCheck.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"aws_chaos": {
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

								Required: false,
								Optional: true,
								Computed: false,
							},

							"azure_chaos": {
								Description:         "AzureChaosSpec is the content of the specification for an AzureChaos",
								MarkdownDescription: "AzureChaosSpec is the content of the specification for an AzureChaos",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"action": {
										Description:         "Action defines the specific azure chaos action. Supported action: vm-stop / vm-restart / disk-detach Default action: vm-stop",
										MarkdownDescription: "Action defines the specific azure chaos action. Supported action: vm-stop / vm-restart / disk-detach Default action: vm-stop",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("vm-stop", "vm-restart", "disk-detach"),
										},
									},

									"disk_name": {
										Description:         "DiskName indicates the name of the disk. Needed in disk-detach.",
										MarkdownDescription: "DiskName indicates the name of the disk. Needed in disk-detach.",

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

									"lun": {
										Description:         "LUN indicates the Logical Unit Number of the data disk. Needed in disk-detach.",
										MarkdownDescription: "LUN indicates the Logical Unit Number of the data disk. Needed in disk-detach.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resource_group_name": {
										Description:         "ResourceGroupName defines the name of ResourceGroup",
										MarkdownDescription: "ResourceGroupName defines the name of ResourceGroup",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret_name": {
										Description:         "SecretName defines the name of kubernetes secret. It is used for Azure credentials.",
										MarkdownDescription: "SecretName defines the name of kubernetes secret. It is used for Azure credentials.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subscription_id": {
										Description:         "SubscriptionID defines the id of Azure subscription.",
										MarkdownDescription: "SubscriptionID defines the id of Azure subscription.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"vm_name": {
										Description:         "VMName defines the name of Virtual Machine",
										MarkdownDescription: "VMName defines the name of Virtual Machine",

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

							"block_chaos": {
								Description:         "BlockChaosSpec is the content of the specification for a BlockChaos",
								MarkdownDescription: "BlockChaosSpec is the content of the specification for a BlockChaos",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"action": {
										Description:         "Action defines the specific block chaos action. Supported action: delay",
										MarkdownDescription: "Action defines the specific block chaos action. Supported action: delay",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("delay"),
										},
									},

									"container_names": {
										Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
										MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"delay": {
										Description:         "Delay defines the delay distribution.",
										MarkdownDescription: "Delay defines the delay distribution.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"correlation": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"jitter": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"latency": {
												Description:         "Latency defines the latency of every io request.",
												MarkdownDescription: "Latency defines the latency of every io request.",

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

									"duration": {
										Description:         "Duration represents the duration of the chaos action.",
										MarkdownDescription: "Duration represents the duration of the chaos action.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mode": {
										Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
										MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
										},
									},

									"selector": {
										Description:         "Selector is used to select pods that are used to inject chaos action.",
										MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expression_selectors": {
												Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
												MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"label_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "Namespaces is a set of namespace to which objects belong.",
												MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selectors": {
												Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
												MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nodes": {
												Description:         "Nodes is a set of node name and objects must belong to these nodes.",
												MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_phase_selectors": {
												Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
												MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pods": {
												Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
												MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
										MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_name": {
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

							"children": {
								Description:         "Children describes the children steps of serial or parallel node. Only used when Type is TypeSerial or TypeParallel.",
								MarkdownDescription: "Children describes the children steps of serial or parallel node. Only used when Type is TypeSerial or TypeParallel.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"conditional_branches": {
								Description:         "ConditionalBranches describes the conditional branches of custom tasks. Only used when Type is TypeTask.",
								MarkdownDescription: "ConditionalBranches describes the conditional branches of custom tasks. Only used when Type is TypeTask.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"expression": {
										Description:         "Expression is the expression for this conditional branch, expected type of result is boolean. If expression is empty, this branch will always be selected/the template will be spawned.",
										MarkdownDescription: "Expression is the expression for this conditional branch, expected type of result is boolean. If expression is empty, this branch will always be selected/the template will be spawned.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target": {
										Description:         "Target is the name of other template, if expression is evaluated as true, this template will be spawned.",
										MarkdownDescription: "Target is the name of other template, if expression is evaluated as true, this template will be spawned.",

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

							"deadline": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dns_chaos": {
								Description:         "DNSChaosSpec defines the desired state of DNSChaos",
								MarkdownDescription: "DNSChaosSpec defines the desired state of DNSChaos",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"action": {
										Description:         "Action defines the specific DNS chaos action. Supported action: error, random Default action: error",
										MarkdownDescription: "Action defines the specific DNS chaos action. Supported action: error, random Default action: error",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("error", "random"),
										},
									},

									"container_names": {
										Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
										MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"duration": {
										Description:         "Duration represents the duration of the chaos action",
										MarkdownDescription: "Duration represents the duration of the chaos action",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mode": {
										Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
										MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
										},
									},

									"patterns": {
										Description:         "Choose which domain names to take effect, support the placeholder ? and wildcard *, or the Specified domain name. Note:      1. The wildcard * must be at the end of the string. For example, chaos-*.org is invalid.      2. if the patterns is empty, will take effect on all the domain names. For example: 		The value is ['google.com', 'github.*', 'chaos-mes?.org'], 		will take effect on 'google.com', 'github.com' and 'chaos-mesh.org'",
										MarkdownDescription: "Choose which domain names to take effect, support the placeholder ? and wildcard *, or the Specified domain name. Note:      1. The wildcard * must be at the end of the string. For example, chaos-*.org is invalid.      2. if the patterns is empty, will take effect on all the domain names. For example: 		The value is ['google.com', 'github.*', 'chaos-mes?.org'], 		will take effect on 'google.com', 'github.com' and 'chaos-mesh.org'",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"selector": {
										Description:         "Selector is used to select pods that are used to inject chaos action.",
										MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expression_selectors": {
												Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
												MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"label_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "Namespaces is a set of namespace to which objects belong.",
												MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selectors": {
												Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
												MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nodes": {
												Description:         "Nodes is a set of node name and objects must belong to these nodes.",
												MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_phase_selectors": {
												Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
												MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pods": {
												Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
												MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
										MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

							"gcp_chaos": {
								Description:         "GCPChaosSpec is the content of the specification for a GCPChaos",
								MarkdownDescription: "GCPChaosSpec is the content of the specification for a GCPChaos",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"action": {
										Description:         "Action defines the specific gcp chaos action. Supported action: node-stop / node-reset / disk-loss Default action: node-stop",
										MarkdownDescription: "Action defines the specific gcp chaos action. Supported action: node-stop / node-reset / disk-loss Default action: node-stop",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("node-stop", "node-reset", "disk-loss"),
										},
									},

									"device_names": {
										Description:         "The device name of disks to detach. Needed in disk-loss.",
										MarkdownDescription: "The device name of disks to detach. Needed in disk-loss.",

										Type: types.ListType{ElemType: types.StringType},

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

									"instance": {
										Description:         "Instance defines the name of the instance",
										MarkdownDescription: "Instance defines the name of the instance",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"project": {
										Description:         "Project defines the ID of gcp project.",
										MarkdownDescription: "Project defines the ID of gcp project.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret_name": {
										Description:         "SecretName defines the name of kubernetes secret. It is used for GCP credentials.",
										MarkdownDescription: "SecretName defines the name of kubernetes secret. It is used for GCP credentials.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"zone": {
										Description:         "Zone defines the zone of gcp project.",
										MarkdownDescription: "Zone defines the zone of gcp project.",

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

							"http_chaos": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"abort": {
										Description:         "Abort is a rule to abort a http session.",
										MarkdownDescription: "Abort is a rule to abort a http session.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"code": {
										Description:         "Code is a rule to select target by http status code in response.",
										MarkdownDescription: "Code is a rule to select target by http status code in response.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"delay": {
										Description:         "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
										MarkdownDescription: "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

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

									"method": {
										Description:         "Method is a rule to select target by http method in request.",
										MarkdownDescription: "Method is a rule to select target by http method in request.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mode": {
										Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
										MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
										},
									},

									"patch": {
										Description:         "Patch is a rule to patch some contents in target.",
										MarkdownDescription: "Patch is a rule to patch some contents in target.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"body": {
												Description:         "Body is a rule to patch message body of target.",
												MarkdownDescription: "Body is a rule to patch message body of target.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"type": {
														Description:         "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",
														MarkdownDescription: "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Value is the patch contents.",
														MarkdownDescription: "Value is the patch contents.",

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

											"headers": {
												Description:         "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",
												MarkdownDescription: "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"queries": {
												Description:         "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",
												MarkdownDescription: "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path is a rule to select target by uri path in http request.",
										MarkdownDescription: "Path is a rule to select target by uri path in http request.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Port represents the target port to be proxy of.",
										MarkdownDescription: "Port represents the target port to be proxy of.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replace": {
										Description:         "Replace is a rule to replace some contents in target.",
										MarkdownDescription: "Replace is a rule to replace some contents in target.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"body": {
												Description:         "Body is a rule to replace http message body in target.",
												MarkdownDescription: "Body is a rule to replace http message body in target.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.Base64Validator(),
												},
											},

											"code": {
												Description:         "Code is a rule to replace http status code in response.",
												MarkdownDescription: "Code is a rule to replace http status code in response.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"headers": {
												Description:         "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",
												MarkdownDescription: "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"method": {
												Description:         "Method is a rule to replace http method in request.",
												MarkdownDescription: "Method is a rule to replace http method in request.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "Path is rule to to replace uri path in http request.",
												MarkdownDescription: "Path is rule to to replace uri path in http request.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"queries": {
												Description:         "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",
												MarkdownDescription: "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",

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

									"request_headers": {
										Description:         "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",
										MarkdownDescription: "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"response_headers": {
										Description:         "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",
										MarkdownDescription: "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"selector": {
										Description:         "Selector is used to select pods that are used to inject chaos action.",
										MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expression_selectors": {
												Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
												MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"label_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "Namespaces is a set of namespace to which objects belong.",
												MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selectors": {
												Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
												MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nodes": {
												Description:         "Nodes is a set of node name and objects must belong to these nodes.",
												MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_phase_selectors": {
												Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
												MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pods": {
												Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
												MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"target": {
										Description:         "Target is the object to be selected and injected.",
										MarkdownDescription: "Target is the object to be selected and injected.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Request", "Response"),
										},
									},

									"value": {
										Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
										MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

							"io_chaos": {
								Description:         "IOChaosSpec defines the desired state of IOChaos",
								MarkdownDescription: "IOChaosSpec defines the desired state of IOChaos",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"action": {
										Description:         "Action defines the specific pod chaos action. Supported action: latency / fault / attrOverride / mistake",
										MarkdownDescription: "Action defines the specific pod chaos action. Supported action: latency / fault / attrOverride / mistake",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("latency", "fault", "attrOverride", "mistake"),
										},
									},

									"attr": {
										Description:         "Attr defines the overrided attribution",
										MarkdownDescription: "Attr defines the overrided attribution",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"atime": {
												Description:         "Timespec represents a time",
												MarkdownDescription: "Timespec represents a time",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"nsec": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"sec": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"blocks": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ctime": {
												Description:         "Timespec represents a time",
												MarkdownDescription: "Timespec represents a time",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"nsec": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"sec": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"gid": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ino": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kind": {
												Description:         "FileType represents type of file",
												MarkdownDescription: "FileType represents type of file",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mtime": {
												Description:         "Timespec represents a time",
												MarkdownDescription: "Timespec represents a time",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"nsec": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"sec": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nlink": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"perm": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rdev": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"uid": {
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

									"container_names": {
										Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
										MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"delay": {
										Description:         "Delay defines the value of I/O chaos action delay. A delay string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
										MarkdownDescription: "Delay defines the value of I/O chaos action delay. A delay string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"duration": {
										Description:         "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
										MarkdownDescription: "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"errno": {
										Description:         "Errno defines the error code that returned by I/O action. refer to: https://www-numi.fnal.gov/offline_software/srt_public_context/WebDocs/Errors/unix_system_errors.html",
										MarkdownDescription: "Errno defines the error code that returned by I/O action. refer to: https://www-numi.fnal.gov/offline_software/srt_public_context/WebDocs/Errors/unix_system_errors.html",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"methods": {
										Description:         "Methods defines the I/O methods for injecting I/O chaos action. default: all I/O methods.",
										MarkdownDescription: "Methods defines the I/O methods for injecting I/O chaos action. default: all I/O methods.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mistake": {
										Description:         "Mistake defines what types of incorrectness are injected to IO operations",
										MarkdownDescription: "Mistake defines what types of incorrectness are injected to IO operations",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"filling": {
												Description:         "Filling determines what is filled in the mistake data.",
												MarkdownDescription: "Filling determines what is filled in the mistake data.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("zero", "random"),
												},
											},

											"max_length": {
												Description:         "Max length of each wrong data segment in bytes",
												MarkdownDescription: "Max length of each wrong data segment in bytes",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"max_occurrences": {
												Description:         "There will be [1, MaxOccurrences] segments of wrong data.",
												MarkdownDescription: "There will be [1, MaxOccurrences] segments of wrong data.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mode": {
										Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
										MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
										},
									},

									"path": {
										Description:         "Path defines the path of files for injecting I/O chaos action.",
										MarkdownDescription: "Path defines the path of files for injecting I/O chaos action.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"percent": {
										Description:         "Percent defines the percentage of injection errors and provides a number from 0-100. default: 100.",
										MarkdownDescription: "Percent defines the percentage of injection errors and provides a number from 0-100. default: 100.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"selector": {
										Description:         "Selector is used to select pods that are used to inject chaos action.",
										MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expression_selectors": {
												Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
												MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"label_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "Namespaces is a set of namespace to which objects belong.",
												MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selectors": {
												Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
												MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nodes": {
												Description:         "Nodes is a set of node name and objects must belong to these nodes.",
												MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_phase_selectors": {
												Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
												MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pods": {
												Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
												MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
										MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_path": {
										Description:         "VolumePath represents the mount path of injected volume",
										MarkdownDescription: "VolumePath represents the mount path of injected volume",

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

							"jvm_chaos": {
								Description:         "JVMChaosSpec defines the desired state of JVMChaos",
								MarkdownDescription: "JVMChaosSpec defines the desired state of JVMChaos",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"action": {
										Description:         "Action defines the specific jvm chaos action. Supported action: latency;return;exception;stress;gc;ruleData",
										MarkdownDescription: "Action defines the specific jvm chaos action. Supported action: latency;return;exception;stress;gc;ruleData",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("latency", "return", "exception", "stress", "gc", "ruleData", "mysql"),
										},
									},

									"class": {
										Description:         "Java class",
										MarkdownDescription: "Java class",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"container_names": {
										Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
										MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cpu_count": {
										Description:         "the CPU core number needs to use, only set it when action is stress",
										MarkdownDescription: "the CPU core number needs to use, only set it when action is stress",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"database": {
										Description:         "the match database default value is '', means match all database",
										MarkdownDescription: "the match database default value is '', means match all database",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"duration": {
										Description:         "Duration represents the duration of the chaos action",
										MarkdownDescription: "Duration represents the duration of the chaos action",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"exception": {
										Description:         "the exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
										MarkdownDescription: "the exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"latency": {
										Description:         "the latency duration for action 'latency', unit ms or the latency duration in action 'mysql'",
										MarkdownDescription: "the latency duration for action 'latency', unit ms or the latency duration in action 'mysql'",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mem_type": {
										Description:         "the memory type needs to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
										MarkdownDescription: "the memory type needs to locate, only set it when action is stress, the value can be 'stack' or 'heap'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"method": {
										Description:         "the method in Java class",
										MarkdownDescription: "the method in Java class",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mode": {
										Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
										MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
										},
									},

									"mysql_connector_version": {
										Description:         "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
										MarkdownDescription: "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "byteman rule name, should be unique, and will generate one if not set",
										MarkdownDescription: "byteman rule name, should be unique, and will generate one if not set",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pid": {
										Description:         "the pid of Java process which needs to attach",
										MarkdownDescription: "the pid of Java process which needs to attach",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "the port of agent server, default 9277",
										MarkdownDescription: "the port of agent server, default 9277",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rule_data": {
										Description:         "the byteman rule's data for action 'ruleData'",
										MarkdownDescription: "the byteman rule's data for action 'ruleData'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"selector": {
										Description:         "Selector is used to select pods that are used to inject chaos action.",
										MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expression_selectors": {
												Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
												MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"label_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "Namespaces is a set of namespace to which objects belong.",
												MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selectors": {
												Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
												MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nodes": {
												Description:         "Nodes is a set of node name and objects must belong to these nodes.",
												MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_phase_selectors": {
												Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
												MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pods": {
												Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
												MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"sql_type": {
										Description:         "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
										MarkdownDescription: "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"table": {
										Description:         "the match table default value is '', means match all table",
										MarkdownDescription: "the match table default value is '', means match all table",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
										MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

							"kernel_chaos": {
								Description:         "KernelChaosSpec defines the desired state of KernelChaos",
								MarkdownDescription: "KernelChaosSpec defines the desired state of KernelChaos",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"container_names": {
										Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
										MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"duration": {
										Description:         "Duration represents the duration of the chaos action",
										MarkdownDescription: "Duration represents the duration of the chaos action",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"fail_kern_request": {
										Description:         "FailKernRequest defines the request of kernel injection",
										MarkdownDescription: "FailKernRequest defines the request of kernel injection",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"callchain": {
												Description:         "Callchain indicate a special call chain, such as:     ext4_mount       -> mount_subtree          -> ...             -> should_failslab With an optional set of predicates and an optional set of parameters, which used with predicates. You can read call chan and predicate examples from https://github.com/chaos-mesh/bpfki/tree/develop/examples to learn more. If no special call chain, just keep Callchain empty, which means it will fail at any call chain with slab alloc (eg: kmalloc).",
												MarkdownDescription: "Callchain indicate a special call chain, such as:     ext4_mount       -> mount_subtree          -> ...             -> should_failslab With an optional set of predicates and an optional set of parameters, which used with predicates. You can read call chan and predicate examples from https://github.com/chaos-mesh/bpfki/tree/develop/examples to learn more. If no special call chain, just keep Callchain empty, which means it will fail at any call chain with slab alloc (eg: kmalloc).",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"funcname": {
														Description:         "Funcname can be find from kernel source or '/proc/kallsyms', such as 'ext4_mount'",
														MarkdownDescription: "Funcname can be find from kernel source or '/proc/kallsyms', such as 'ext4_mount'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"parameters": {
														Description:         "Parameters is used with predicate, for example, if you want to inject slab error in 'd_alloc_parallel(struct dentry *parent, const struct qstr *name)' with a special name 'bananas', you need to set it to 'struct dentry *parent, const struct qstr *name' otherwise omit it.",
														MarkdownDescription: "Parameters is used with predicate, for example, if you want to inject slab error in 'd_alloc_parallel(struct dentry *parent, const struct qstr *name)' with a special name 'bananas', you need to set it to 'struct dentry *parent, const struct qstr *name' otherwise omit it.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"predicate": {
														Description:         "Predicate will access the arguments of this Frame, example with Parameters's, you can set it to 'STRNCMP(name->name, 'bananas', 8)' to make inject only with it, or omit it to inject for all d_alloc_parallel call chain.",
														MarkdownDescription: "Predicate will access the arguments of this Frame, example with Parameters's, you can set it to 'STRNCMP(name->name, 'bananas', 8)' to make inject only with it, or omit it to inject for all d_alloc_parallel call chain.",

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

											"failtype": {
												Description:         "FailType indicates what to fail, can be set to '0' / '1' / '2' If '0', indicates slab to fail (should_failslab) If '1', indicates alloc_page to fail (should_fail_alloc_page) If '2', indicates bio to fail (should_fail_bio) You can read:   1. https://www.kernel.org/doc/html/latest/fault-injection/fault-injection.html   2. http://github.com/iovisor/bcc/blob/master/tools/inject_example.txt to learn more",
												MarkdownDescription: "FailType indicates what to fail, can be set to '0' / '1' / '2' If '0', indicates slab to fail (should_failslab) If '1', indicates alloc_page to fail (should_fail_alloc_page) If '2', indicates bio to fail (should_fail_bio) You can read:   1. https://www.kernel.org/doc/html/latest/fault-injection/fault-injection.html   2. http://github.com/iovisor/bcc/blob/master/tools/inject_example.txt to learn more",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(2),
												},
											},

											"headers": {
												Description:         "Headers indicates the appropriate kernel headers you need. Eg: 'linux/mmzone.h', 'linux/blkdev.h' and so on",
												MarkdownDescription: "Headers indicates the appropriate kernel headers you need. Eg: 'linux/mmzone.h', 'linux/blkdev.h' and so on",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"probability": {
												Description:         "Probability indicates the fails with probability. If you want 1%, please set this field with 1.",
												MarkdownDescription: "Probability indicates the fails with probability. If you want 1%, please set this field with 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"times": {
												Description:         "Times indicates the max times of fails.",
												MarkdownDescription: "Times indicates the max times of fails.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"mode": {
										Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
										MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
										},
									},

									"selector": {
										Description:         "Selector is used to select pods that are used to inject chaos action.",
										MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expression_selectors": {
												Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
												MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"label_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "Namespaces is a set of namespace to which objects belong.",
												MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selectors": {
												Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
												MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nodes": {
												Description:         "Nodes is a set of node name and objects must belong to these nodes.",
												MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_phase_selectors": {
												Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
												MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pods": {
												Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
												MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
										MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"network_chaos": {
								Description:         "NetworkChaosSpec defines the desired state of NetworkChaos",
								MarkdownDescription: "NetworkChaosSpec defines the desired state of NetworkChaos",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"action": {
										Description:         "Action defines the specific network chaos action. Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay",
										MarkdownDescription: "Action defines the specific network chaos action. Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("netem", "delay", "loss", "duplicate", "corrupt", "partition", "bandwidth"),
										},
									},

									"bandwidth": {
										Description:         "Bandwidth represents the detail about bandwidth control action",
										MarkdownDescription: "Bandwidth represents the detail about bandwidth control action",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"buffer": {
												Description:         "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",
												MarkdownDescription: "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"limit": {
												Description:         "Limit is the number of bytes that can be queued waiting for tokens to become available.",
												MarkdownDescription: "Limit is the number of bytes that can be queued waiting for tokens to become available.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"minburst": {
												Description:         "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",
												MarkdownDescription: "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"peakrate": {
												Description:         "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",
												MarkdownDescription: "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"rate": {
												Description:         "Rate is the speed knob. Allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",
												MarkdownDescription: "Rate is the speed knob. Allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",

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

									"corrupt": {
										Description:         "Corrupt represents the detail about corrupt action",
										MarkdownDescription: "Corrupt represents the detail about corrupt action",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"correlation": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"corrupt": {
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

									"delay": {
										Description:         "Delay represents the detail about delay action",
										MarkdownDescription: "Delay represents the detail about delay action",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"correlation": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"jitter": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"latency": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"reorder": {
												Description:         "ReorderSpec defines details of packet reorder.",
												MarkdownDescription: "ReorderSpec defines details of packet reorder.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"correlation": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"gap": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"reorder": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"device": {
										Description:         "Device represents the network device to be affected.",
										MarkdownDescription: "Device represents the network device to be affected.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"direction": {
										Description:         "Direction represents the direction, this applies on netem and network partition action",
										MarkdownDescription: "Direction represents the direction, this applies on netem and network partition action",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("to", "from", "both"),
										},
									},

									"duplicate": {
										Description:         "DuplicateSpec represents the detail about loss action",
										MarkdownDescription: "DuplicateSpec represents the detail about loss action",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"correlation": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"duplicate": {
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

									"duration": {
										Description:         "Duration represents the duration of the chaos action",
										MarkdownDescription: "Duration represents the duration of the chaos action",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"external_targets": {
										Description:         "ExternalTargets represents network targets outside k8s",
										MarkdownDescription: "ExternalTargets represents network targets outside k8s",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"loss": {
										Description:         "Loss represents the detail about loss action",
										MarkdownDescription: "Loss represents the detail about loss action",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"correlation": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"loss": {
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

									"mode": {
										Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
										MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
										},
									},

									"selector": {
										Description:         "Selector is used to select pods that are used to inject chaos action.",
										MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expression_selectors": {
												Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
												MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"label_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "Namespaces is a set of namespace to which objects belong.",
												MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selectors": {
												Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
												MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nodes": {
												Description:         "Nodes is a set of node name and objects must belong to these nodes.",
												MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_phase_selectors": {
												Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
												MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pods": {
												Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
												MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"target": {
										Description:         "Target represents network target, this applies on netem and network partition action",
										MarkdownDescription: "Target represents network target, this applies on netem and network partition action",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"mode": {
												Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
												MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
												},
											},

											"selector": {
												Description:         "Selector is used to select pods that are used to inject chaos action.",
												MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression_selectors": {
														Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
														MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "Namespaces is a set of namespace to which objects belong.",
														MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selectors": {
														Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
														MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nodes": {
														Description:         "Nodes is a set of node name and objects must belong to these nodes.",
														MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_phase_selectors": {
														Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
														MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pods": {
														Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
														MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
												MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

									"target_device": {
										Description:         "TargetDevice represents the network device to be affected in target scope.",
										MarkdownDescription: "TargetDevice represents the network device to be affected in target scope.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
										MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

							"physicalmachine_chaos": {
								Description:         "PhysicalMachineChaosSpec defines the desired state of PhysicalMachineChaos",
								MarkdownDescription: "PhysicalMachineChaosSpec defines the desired state of PhysicalMachineChaos",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"action": {
										Description:         "the subAction, generate automatically",
										MarkdownDescription: "the subAction, generate automatically",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("stress-cpu", "stress-mem", "disk-read-payload", "disk-write-payload", "disk-fill", "network-corrupt", "network-duplicate", "network-loss", "network-delay", "network-partition", "network-dns", "network-bandwidth", "network-flood", "network-down", "process", "jvm-exception", "jvm-gc", "jvm-latency", "jvm-return", "jvm-stress", "jvm-rule-data", "jvm-mysql", "clock", "redis-expiration", "redis-penetration", "redis-cacheLimit", "redis-restart", "redis-stop", "kafka-fill", "kafka-flood", "kafka-io", "file-create", "file-modify", "file-delete", "file-rename", "file-append", "file-replace", "vm", "user_defined"),
										},
									},

									"address": {
										Description:         "DEPRECATED: Use Selector instead. Only one of Address and Selector could be specified.",
										MarkdownDescription: "DEPRECATED: Use Selector instead. Only one of Address and Selector could be specified.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"clock": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"clock_ids_slice": {
												Description:         "the identifier of the particular clock on which to act. More clock description in linux kernel can be found in man page of clock_getres, clock_gettime, clock_settime. Muti clock ids should be split with ','",
												MarkdownDescription: "the identifier of the particular clock on which to act. More clock description in linux kernel can be found in man page of clock_getres, clock_gettime, clock_settime. Muti clock ids should be split with ','",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pid": {
												Description:         "the pid of target program.",
												MarkdownDescription: "the pid of target program.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"time_offset": {
												Description:         "specifies the length of time offset.",
												MarkdownDescription: "specifies the length of time offset.",

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

									"disk_fill": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fill_by_fallocate": {
												Description:         "fill disk by fallocate",
												MarkdownDescription: "fill disk by fallocate",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
												MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size": {
												Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
												MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",

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

									"disk_read_payload": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"path": {
												Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
												MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"payload_process_num": {
												Description:         "specifies the number of process work on writing, default 1, only 1-255 is valid value",
												MarkdownDescription: "specifies the number of process work on writing, default 1, only 1-255 is valid value",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size": {
												Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
												MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",

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

									"disk_write_payload": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"path": {
												Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
												MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"payload_process_num": {
												Description:         "specifies the number of process work on writing, default 1, only 1-255 is valid value",
												MarkdownDescription: "specifies the number of process work on writing, default 1, only 1-255 is valid value",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size": {
												Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
												MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",

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

									"duration": {
										Description:         "Duration represents the duration of the chaos action",
										MarkdownDescription: "Duration represents the duration of the chaos action",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"file_append": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"count": {
												Description:         "Count is the number of times to append the data.",
												MarkdownDescription: "Count is the number of times to append the data.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"data": {
												Description:         "Data is the data for append.",
												MarkdownDescription: "Data is the data for append.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"file_name": {
												Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
												MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

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

									"file_create": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dir_name": {
												Description:         "DirName is the directory name to create or delete.",
												MarkdownDescription: "DirName is the directory name to create or delete.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"file_name": {
												Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
												MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

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

									"file_delete": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dir_name": {
												Description:         "DirName is the directory name to create or delete.",
												MarkdownDescription: "DirName is the directory name to create or delete.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"file_name": {
												Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
												MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

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

									"file_modify": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"file_name": {
												Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
												MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"privilege": {
												Description:         "Privilege is the file privilege to be set.",
												MarkdownDescription: "Privilege is the file privilege to be set.",

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

									"file_rename": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dest_file": {
												Description:         "DestFile is the name to be renamed.",
												MarkdownDescription: "DestFile is the name to be renamed.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"source_file": {
												Description:         "SourceFile is the name need to be renamed.",
												MarkdownDescription: "SourceFile is the name need to be renamed.",

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

									"file_replace": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dest_string": {
												Description:         "DestStr is the destination string of the file.",
												MarkdownDescription: "DestStr is the destination string of the file.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"file_name": {
												Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
												MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"line": {
												Description:         "Line is the line number of the file to be replaced.",
												MarkdownDescription: "Line is the line number of the file to be replaced.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"origin_string": {
												Description:         "OriginStr is the origin string of the file.",
												MarkdownDescription: "OriginStr is the origin string of the file.",

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

									"http_abort": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"code": {
												Description:         "Code is a rule to select target by http status code in response",
												MarkdownDescription: "Code is a rule to select target by http status code in response",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"method": {
												Description:         "HTTP method",
												MarkdownDescription: "HTTP method",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "Match path of Uri with wildcard matches",
												MarkdownDescription: "Match path of Uri with wildcard matches",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "The TCP port that the target service listens on",
												MarkdownDescription: "The TCP port that the target service listens on",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy_ports": {
												Description:         "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
												MarkdownDescription: "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"target": {
												Description:         "HTTP target: Request or Response",
												MarkdownDescription: "HTTP target: Request or Response",

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

									"http_config": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"file_path": {
												Description:         "The config file path",
												MarkdownDescription: "The config file path",

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

									"http_delay": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"code": {
												Description:         "Code is a rule to select target by http status code in response",
												MarkdownDescription: "Code is a rule to select target by http status code in response",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"delay": {
												Description:         "Delay represents the delay of the target request/response",
												MarkdownDescription: "Delay represents the delay of the target request/response",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"method": {
												Description:         "HTTP method",
												MarkdownDescription: "HTTP method",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "Match path of Uri with wildcard matches",
												MarkdownDescription: "Match path of Uri with wildcard matches",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "The TCP port that the target service listens on",
												MarkdownDescription: "The TCP port that the target service listens on",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy_ports": {
												Description:         "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
												MarkdownDescription: "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"target": {
												Description:         "HTTP target: Request or Response",
												MarkdownDescription: "HTTP target: Request or Response",

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

									"http_request": {
										Description:         "used for HTTP request, now only support GET",
										MarkdownDescription: "used for HTTP request, now only support GET",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"count": {
												Description:         "The number of requests to send",
												MarkdownDescription: "The number of requests to send",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_conn_pool": {
												Description:         "Enable connection pool",
												MarkdownDescription: "Enable connection pool",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"url": {
												Description:         "Request to send'",
												MarkdownDescription: "Request to send'",

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

									"jvm_exception": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"class": {
												Description:         "Java class",
												MarkdownDescription: "Java class",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"exception": {
												Description:         "the exception which needs to throw for action 'exception'",
												MarkdownDescription: "the exception which needs to throw for action 'exception'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"method": {
												Description:         "the method in Java class",
												MarkdownDescription: "the method in Java class",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pid": {
												Description:         "the pid of Java process which needs to attach",
												MarkdownDescription: "the pid of Java process which needs to attach",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "the port of agent server, default 9277",
												MarkdownDescription: "the port of agent server, default 9277",

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

									"jvm_gc": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"pid": {
												Description:         "the pid of Java process which needs to attach",
												MarkdownDescription: "the pid of Java process which needs to attach",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "the port of agent server, default 9277",
												MarkdownDescription: "the port of agent server, default 9277",

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

									"jvm_latency": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"class": {
												Description:         "Java class",
												MarkdownDescription: "Java class",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"latency": {
												Description:         "the latency duration for action 'latency', unit ms",
												MarkdownDescription: "the latency duration for action 'latency', unit ms",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"method": {
												Description:         "the method in Java class",
												MarkdownDescription: "the method in Java class",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pid": {
												Description:         "the pid of Java process which needs to attach",
												MarkdownDescription: "the pid of Java process which needs to attach",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "the port of agent server, default 9277",
												MarkdownDescription: "the port of agent server, default 9277",

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

									"jvm_mysql": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"database": {
												Description:         "the match database default value is '', means match all database",
												MarkdownDescription: "the match database default value is '', means match all database",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"exception": {
												Description:         "The exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
												MarkdownDescription: "The exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"latency": {
												Description:         "The latency duration for action 'latency' or the latency duration in action 'mysql'",
												MarkdownDescription: "The latency duration for action 'latency' or the latency duration in action 'mysql'",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mysql_connector_version": {
												Description:         "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
												MarkdownDescription: "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pid": {
												Description:         "the pid of Java process which needs to attach",
												MarkdownDescription: "the pid of Java process which needs to attach",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "the port of agent server, default 9277",
												MarkdownDescription: "the port of agent server, default 9277",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sql_type": {
												Description:         "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
												MarkdownDescription: "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"table": {
												Description:         "the match table default value is '', means match all table",
												MarkdownDescription: "the match table default value is '', means match all table",

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

									"jvm_return": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"class": {
												Description:         "Java class",
												MarkdownDescription: "Java class",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"method": {
												Description:         "the method in Java class",
												MarkdownDescription: "the method in Java class",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pid": {
												Description:         "the pid of Java process which needs to attach",
												MarkdownDescription: "the pid of Java process which needs to attach",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "the port of agent server, default 9277",
												MarkdownDescription: "the port of agent server, default 9277",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "the return value for action 'return'",
												MarkdownDescription: "the return value for action 'return'",

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

									"jvm_rule_data": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"pid": {
												Description:         "the pid of Java process which needs to attach",
												MarkdownDescription: "the pid of Java process which needs to attach",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "the port of agent server, default 9277",
												MarkdownDescription: "the port of agent server, default 9277",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rule_data": {
												Description:         "RuleData used to save the rule file's data, will use it when recover",
												MarkdownDescription: "RuleData used to save the rule file's data, will use it when recover",

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

									"jvm_stress": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cpu_count": {
												Description:         "the CPU core number need to use, only set it when action is stress",
												MarkdownDescription: "the CPU core number need to use, only set it when action is stress",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mem_type": {
												Description:         "the memory type need to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
												MarkdownDescription: "the memory type need to locate, only set it when action is stress, the value can be 'stack' or 'heap'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pid": {
												Description:         "the pid of Java process which needs to attach",
												MarkdownDescription: "the pid of Java process which needs to attach",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "the port of agent server, default 9277",
												MarkdownDescription: "the port of agent server, default 9277",

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

									"kafka_fill": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "The host of kafka server",
												MarkdownDescription: "The host of kafka server",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_bytes": {
												Description:         "The max bytes to fill",
												MarkdownDescription: "The max bytes to fill",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"message_size": {
												Description:         "The size of each message",
												MarkdownDescription: "The size of each message",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"password": {
												Description:         "The password of kafka client",
												MarkdownDescription: "The password of kafka client",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "The port of kafka server",
												MarkdownDescription: "The port of kafka server",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"reload_command": {
												Description:         "The command to reload kafka config",
												MarkdownDescription: "The command to reload kafka config",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"topic": {
												Description:         "The topic to attack",
												MarkdownDescription: "The topic to attack",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"username": {
												Description:         "The username of kafka client",
												MarkdownDescription: "The username of kafka client",

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

									"kafka_flood": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "The host of kafka server",
												MarkdownDescription: "The host of kafka server",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"message_size": {
												Description:         "The size of each message",
												MarkdownDescription: "The size of each message",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"password": {
												Description:         "The password of kafka client",
												MarkdownDescription: "The password of kafka client",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "The port of kafka server",
												MarkdownDescription: "The port of kafka server",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"threads": {
												Description:         "The number of worker threads",
												MarkdownDescription: "The number of worker threads",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"topic": {
												Description:         "The topic to attack",
												MarkdownDescription: "The topic to attack",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"username": {
												Description:         "The username of kafka client",
												MarkdownDescription: "The username of kafka client",

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

									"kafka_io": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_file": {
												Description:         "The path of server config",
												MarkdownDescription: "The path of server config",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"non_readable": {
												Description:         "Make kafka cluster non-readable",
												MarkdownDescription: "Make kafka cluster non-readable",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"non_writable": {
												Description:         "Make kafka cluster non-writable",
												MarkdownDescription: "Make kafka cluster non-writable",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"topic": {
												Description:         "The topic to attack",
												MarkdownDescription: "The topic to attack",

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

									"mode": {
										Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
										MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
										},
									},

									"network_bandwidth": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"buffer": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"device": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"hostname": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_address": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"limit": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"minburst": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"peakrate": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rate": {
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

									"network_corrupt": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"correlation": {
												Description:         "correlation is percentage (10 is 10%)",
												MarkdownDescription: "correlation is percentage (10 is 10%)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"device": {
												Description:         "the network interface to impact",
												MarkdownDescription: "the network interface to impact",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"egress_port": {
												Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
												MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"hostname": {
												Description:         "only impact traffic to these hostnames",
												MarkdownDescription: "only impact traffic to these hostnames",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_address": {
												Description:         "only impact egress traffic to these IP addresses",
												MarkdownDescription: "only impact egress traffic to these IP addresses",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_protocol": {
												Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
												MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"percent": {
												Description:         "percentage of packets to corrupt (10 is 10%)",
												MarkdownDescription: "percentage of packets to corrupt (10 is 10%)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"source_port": {
												Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
												MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

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

									"network_delay": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"accept_tcp_flags": {
												Description:         "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
												MarkdownDescription: "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"correlation": {
												Description:         "correlation is percentage (10 is 10%)",
												MarkdownDescription: "correlation is percentage (10 is 10%)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"device": {
												Description:         "the network interface to impact",
												MarkdownDescription: "the network interface to impact",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"egress_port": {
												Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
												MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"hostname": {
												Description:         "only impact traffic to these hostnames",
												MarkdownDescription: "only impact traffic to these hostnames",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_address": {
												Description:         "only impact egress traffic to these IP addresses",
												MarkdownDescription: "only impact egress traffic to these IP addresses",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_protocol": {
												Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
												MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"jitter": {
												Description:         "jitter time, time units: ns, us (or µs), ms, s, m, h.",
												MarkdownDescription: "jitter time, time units: ns, us (or µs), ms, s, m, h.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"latency": {
												Description:         "delay egress time, time units: ns, us (or µs), ms, s, m, h.",
												MarkdownDescription: "delay egress time, time units: ns, us (or µs), ms, s, m, h.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"source_port": {
												Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
												MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

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

									"network_dns": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dns_domain_name": {
												Description:         "map this host to specified IP",
												MarkdownDescription: "map this host to specified IP",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dns_ip": {
												Description:         "map specified host to this IP address",
												MarkdownDescription: "map specified host to this IP address",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dns_server": {
												Description:         "update the DNS server in /etc/resolv.conf with this value",
												MarkdownDescription: "update the DNS server in /etc/resolv.conf with this value",

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

									"network_down": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"device": {
												Description:         "The network interface to impact",
												MarkdownDescription: "The network interface to impact",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"duration": {
												Description:         "NIC down time, time units: ns, us (or µs), ms, s, m, h.",
												MarkdownDescription: "NIC down time, time units: ns, us (or µs), ms, s, m, h.",

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

									"network_duplicate": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"correlation": {
												Description:         "correlation is percentage (10 is 10%)",
												MarkdownDescription: "correlation is percentage (10 is 10%)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"device": {
												Description:         "the network interface to impact",
												MarkdownDescription: "the network interface to impact",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"egress_port": {
												Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
												MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"hostname": {
												Description:         "only impact traffic to these hostnames",
												MarkdownDescription: "only impact traffic to these hostnames",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_address": {
												Description:         "only impact egress traffic to these IP addresses",
												MarkdownDescription: "only impact egress traffic to these IP addresses",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_protocol": {
												Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
												MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"percent": {
												Description:         "percentage of packets to duplicate (10 is 10%)",
												MarkdownDescription: "percentage of packets to duplicate (10 is 10%)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"source_port": {
												Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
												MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

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

									"network_flood": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"duration": {
												Description:         "The number of seconds to run the iperf test",
												MarkdownDescription: "The number of seconds to run the iperf test",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"ip_address": {
												Description:         "Generate traffic to this IP address",
												MarkdownDescription: "Generate traffic to this IP address",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"parallel": {
												Description:         "The number of iperf parallel client threads to run",
												MarkdownDescription: "The number of iperf parallel client threads to run",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Generate traffic to this port on the IP address",
												MarkdownDescription: "Generate traffic to this port on the IP address",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rate": {
												Description:         "The speed of network traffic, allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second",
												MarkdownDescription: "The speed of network traffic, allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second",

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

									"network_loss": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"correlation": {
												Description:         "correlation is percentage (10 is 10%)",
												MarkdownDescription: "correlation is percentage (10 is 10%)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"device": {
												Description:         "the network interface to impact",
												MarkdownDescription: "the network interface to impact",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"egress_port": {
												Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
												MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"hostname": {
												Description:         "only impact traffic to these hostnames",
												MarkdownDescription: "only impact traffic to these hostnames",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_address": {
												Description:         "only impact egress traffic to these IP addresses",
												MarkdownDescription: "only impact egress traffic to these IP addresses",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_protocol": {
												Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
												MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"percent": {
												Description:         "percentage of packets to loss (10 is 10%)",
												MarkdownDescription: "percentage of packets to loss (10 is 10%)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"source_port": {
												Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
												MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

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

									"network_partition": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"accept_tcp_flags": {
												Description:         "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
												MarkdownDescription: "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"device": {
												Description:         "the network interface to impact",
												MarkdownDescription: "the network interface to impact",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"direction": {
												Description:         "specifies the partition direction, values can be 'from', 'to'. 'from' means packets coming from the 'IPAddress' or 'Hostname' and going to your server, 'to' means packets originating from your server and going to the 'IPAddress' or 'Hostname'.",
												MarkdownDescription: "specifies the partition direction, values can be 'from', 'to'. 'from' means packets coming from the 'IPAddress' or 'Hostname' and going to your server, 'to' means packets originating from your server and going to the 'IPAddress' or 'Hostname'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"hostname": {
												Description:         "only impact traffic to these hostnames",
												MarkdownDescription: "only impact traffic to these hostnames",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_address": {
												Description:         "only impact egress traffic to these IP addresses",
												MarkdownDescription: "only impact egress traffic to these IP addresses",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_protocol": {
												Description:         "only impact egress traffic to these IP addresses",
												MarkdownDescription: "only impact egress traffic to these IP addresses",

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

									"process": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"process": {
												Description:         "the process name or the process ID",
												MarkdownDescription: "the process name or the process ID",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"recover_cmd": {
												Description:         "the command to be run when recovering experiment",
												MarkdownDescription: "the command to be run when recovering experiment",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"signal": {
												Description:         "the signal number to send",
												MarkdownDescription: "the signal number to send",

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

									"redis_cache_limit": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"addr": {
												Description:         "The adress of Redis server",
												MarkdownDescription: "The adress of Redis server",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cache_size": {
												Description:         "The size of 'maxmemory'",
												MarkdownDescription: "The size of 'maxmemory'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"password": {
												Description:         "The password of Redis server",
												MarkdownDescription: "The password of Redis server",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"percent": {
												Description:         "Specifies maxmemory as a percentage of the original value",
												MarkdownDescription: "Specifies maxmemory as a percentage of the original value",

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

									"redis_expiration": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"addr": {
												Description:         "The adress of Redis server",
												MarkdownDescription: "The adress of Redis server",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expiration": {
												Description:         "The expiration of the keys",
												MarkdownDescription: "The expiration of the keys",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "The keys to be expired",
												MarkdownDescription: "The keys to be expired",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"option": {
												Description:         "Additional options for 'expiration'",
												MarkdownDescription: "Additional options for 'expiration'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"password": {
												Description:         "The password of Redis server",
												MarkdownDescription: "The password of Redis server",

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

									"redis_penetration": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"addr": {
												Description:         "The adress of Redis server",
												MarkdownDescription: "The adress of Redis server",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"password": {
												Description:         "The password of Redis server",
												MarkdownDescription: "The password of Redis server",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"request_num": {
												Description:         "The number of requests to be sent",
												MarkdownDescription: "The number of requests to be sent",

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

									"redis_restart": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"addr": {
												Description:         "The adress of Redis server",
												MarkdownDescription: "The adress of Redis server",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"conf": {
												Description:         "The path of Sentinel conf",
												MarkdownDescription: "The path of Sentinel conf",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"flush_config": {
												Description:         "The control flag determines whether to flush config",
												MarkdownDescription: "The control flag determines whether to flush config",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"password": {
												Description:         "The password of Redis server",
												MarkdownDescription: "The password of Redis server",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"redis_path": {
												Description:         "The path of 'redis-server' command-line tool",
												MarkdownDescription: "The path of 'redis-server' command-line tool",

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

									"redis_stop": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"addr": {
												Description:         "The adress of Redis server",
												MarkdownDescription: "The adress of Redis server",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"conf": {
												Description:         "The path of Sentinel conf",
												MarkdownDescription: "The path of Sentinel conf",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"flush_config": {
												Description:         "The control flag determines whether to flush config",
												MarkdownDescription: "The control flag determines whether to flush config",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"password": {
												Description:         "The password of Redis server",
												MarkdownDescription: "The password of Redis server",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"redis_path": {
												Description:         "The path of 'redis-server' command-line tool",
												MarkdownDescription: "The path of 'redis-server' command-line tool",

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

									"selector": {
										Description:         "Selector is used to select physical machines that are used to inject chaos action.",
										MarkdownDescription: "Selector is used to select physical machines that are used to inject chaos action.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expression_selectors": {
												Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
												MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"label_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "Namespaces is a set of namespace to which objects belong.",
												MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"physical_machines": {
												Description:         "PhysicalMachines is a map of string keys and a set values that used to select physical machines. The key defines the namespace which physical machine belong, and each value is a set of physical machine names.",
												MarkdownDescription: "PhysicalMachines is a map of string keys and a set values that used to select physical machines. The key defines the namespace which physical machine belong, and each value is a set of physical machine names.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"stress_cpu": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"load": {
												Description:         "specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
												MarkdownDescription: "specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"options": {
												Description:         "extend stress-ng options",
												MarkdownDescription: "extend stress-ng options",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"workers": {
												Description:         "specifies N workers to apply the stressor.",
												MarkdownDescription: "specifies N workers to apply the stressor.",

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

									"stress_mem": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"options": {
												Description:         "extend stress-ng options",
												MarkdownDescription: "extend stress-ng options",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size": {
												Description:         "specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB..",
												MarkdownDescription: "specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB..",

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

									"uid": {
										Description:         "the experiment ID",
										MarkdownDescription: "the experiment ID",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user_defined": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"attack_cmd": {
												Description:         "The command to be executed when attack",
												MarkdownDescription: "The command to be executed when attack",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"recover_cmd": {
												Description:         "The command to be executed when recover",
												MarkdownDescription: "The command to be executed when recover",

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

									"value": {
										Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of physical machines to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of physical machines the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
										MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of physical machines to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of physical machines the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"vm": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"vm_name": {
												Description:         "The name of the VM to be injected",
												MarkdownDescription: "The name of the VM to be injected",

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

							"pod_chaos": {
								Description:         "PodChaosSpec defines the attributes that a user creates on a chaos experiment about pods.",
								MarkdownDescription: "PodChaosSpec defines the attributes that a user creates on a chaos experiment about pods.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"action": {
										Description:         "Action defines the specific pod chaos action. Supported action: pod-kill / pod-failure / container-kill Default action: pod-kill",
										MarkdownDescription: "Action defines the specific pod chaos action. Supported action: pod-kill / pod-failure / container-kill Default action: pod-kill",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("pod-kill", "pod-failure", "container-kill"),
										},
									},

									"container_names": {
										Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
										MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"duration": {
										Description:         "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
										MarkdownDescription: "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grace_period": {
										Description:         "GracePeriod is used in pod-kill action. It represents the duration in seconds before the pod should be deleted. Value must be non-negative integer. The default value is zero that indicates delete immediately.",
										MarkdownDescription: "GracePeriod is used in pod-kill action. It represents the duration in seconds before the pod should be deleted. Value must be non-negative integer. The default value is zero that indicates delete immediately.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"mode": {
										Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
										MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
										},
									},

									"selector": {
										Description:         "Selector is used to select pods that are used to inject chaos action.",
										MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expression_selectors": {
												Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
												MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"label_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "Namespaces is a set of namespace to which objects belong.",
												MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selectors": {
												Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
												MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nodes": {
												Description:         "Nodes is a set of node name and objects must belong to these nodes.",
												MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_phase_selectors": {
												Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
												MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pods": {
												Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
												MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
										MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

							"schedule": {
								Description:         "Schedule describe the Schedule(describing scheduled chaos) to be injected with chaos nodes. Only used when Type is TypeSchedule.",
								MarkdownDescription: "Schedule describe the Schedule(describing scheduled chaos) to be injected with chaos nodes. Only used when Type is TypeSchedule.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"aws_chaos": {
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

										Required: false,
										Optional: true,
										Computed: false,
									},

									"azure_chaos": {
										Description:         "AzureChaosSpec is the content of the specification for an AzureChaos",
										MarkdownDescription: "AzureChaosSpec is the content of the specification for an AzureChaos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "Action defines the specific azure chaos action. Supported action: vm-stop / vm-restart / disk-detach Default action: vm-stop",
												MarkdownDescription: "Action defines the specific azure chaos action. Supported action: vm-stop / vm-restart / disk-detach Default action: vm-stop",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("vm-stop", "vm-restart", "disk-detach"),
												},
											},

											"disk_name": {
												Description:         "DiskName indicates the name of the disk. Needed in disk-detach.",
												MarkdownDescription: "DiskName indicates the name of the disk. Needed in disk-detach.",

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

											"lun": {
												Description:         "LUN indicates the Logical Unit Number of the data disk. Needed in disk-detach.",
												MarkdownDescription: "LUN indicates the Logical Unit Number of the data disk. Needed in disk-detach.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_group_name": {
												Description:         "ResourceGroupName defines the name of ResourceGroup",
												MarkdownDescription: "ResourceGroupName defines the name of ResourceGroup",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"secret_name": {
												Description:         "SecretName defines the name of kubernetes secret. It is used for Azure credentials.",
												MarkdownDescription: "SecretName defines the name of kubernetes secret. It is used for Azure credentials.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"subscription_id": {
												Description:         "SubscriptionID defines the id of Azure subscription.",
												MarkdownDescription: "SubscriptionID defines the id of Azure subscription.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"vm_name": {
												Description:         "VMName defines the name of Virtual Machine",
												MarkdownDescription: "VMName defines the name of Virtual Machine",

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

									"block_chaos": {
										Description:         "BlockChaosSpec is the content of the specification for a BlockChaos",
										MarkdownDescription: "BlockChaosSpec is the content of the specification for a BlockChaos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "Action defines the specific block chaos action. Supported action: delay",
												MarkdownDescription: "Action defines the specific block chaos action. Supported action: delay",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("delay"),
												},
											},

											"container_names": {
												Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
												MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"delay": {
												Description:         "Delay defines the delay distribution.",
												MarkdownDescription: "Delay defines the delay distribution.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"correlation": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"jitter": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"latency": {
														Description:         "Latency defines the latency of every io request.",
														MarkdownDescription: "Latency defines the latency of every io request.",

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

											"duration": {
												Description:         "Duration represents the duration of the chaos action.",
												MarkdownDescription: "Duration represents the duration of the chaos action.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mode": {
												Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
												MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
												},
											},

											"selector": {
												Description:         "Selector is used to select pods that are used to inject chaos action.",
												MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression_selectors": {
														Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
														MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "Namespaces is a set of namespace to which objects belong.",
														MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selectors": {
														Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
														MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nodes": {
														Description:         "Nodes is a set of node name and objects must belong to these nodes.",
														MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_phase_selectors": {
														Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
														MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pods": {
														Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
														MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
												MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_name": {
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

									"concurrency_policy": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Forbid", "Allow"),
										},
									},

									"dns_chaos": {
										Description:         "DNSChaosSpec defines the desired state of DNSChaos",
										MarkdownDescription: "DNSChaosSpec defines the desired state of DNSChaos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "Action defines the specific DNS chaos action. Supported action: error, random Default action: error",
												MarkdownDescription: "Action defines the specific DNS chaos action. Supported action: error, random Default action: error",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("error", "random"),
												},
											},

											"container_names": {
												Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
												MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"duration": {
												Description:         "Duration represents the duration of the chaos action",
												MarkdownDescription: "Duration represents the duration of the chaos action",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mode": {
												Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
												MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
												},
											},

											"patterns": {
												Description:         "Choose which domain names to take effect, support the placeholder ? and wildcard *, or the Specified domain name. Note:      1. The wildcard * must be at the end of the string. For example, chaos-*.org is invalid.      2. if the patterns is empty, will take effect on all the domain names. For example: 		The value is ['google.com', 'github.*', 'chaos-mes?.org'], 		will take effect on 'google.com', 'github.com' and 'chaos-mesh.org'",
												MarkdownDescription: "Choose which domain names to take effect, support the placeholder ? and wildcard *, or the Specified domain name. Note:      1. The wildcard * must be at the end of the string. For example, chaos-*.org is invalid.      2. if the patterns is empty, will take effect on all the domain names. For example: 		The value is ['google.com', 'github.*', 'chaos-mes?.org'], 		will take effect on 'google.com', 'github.com' and 'chaos-mesh.org'",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "Selector is used to select pods that are used to inject chaos action.",
												MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression_selectors": {
														Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
														MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "Namespaces is a set of namespace to which objects belong.",
														MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selectors": {
														Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
														MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nodes": {
														Description:         "Nodes is a set of node name and objects must belong to these nodes.",
														MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_phase_selectors": {
														Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
														MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pods": {
														Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
														MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
												MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

									"gcp_chaos": {
										Description:         "GCPChaosSpec is the content of the specification for a GCPChaos",
										MarkdownDescription: "GCPChaosSpec is the content of the specification for a GCPChaos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "Action defines the specific gcp chaos action. Supported action: node-stop / node-reset / disk-loss Default action: node-stop",
												MarkdownDescription: "Action defines the specific gcp chaos action. Supported action: node-stop / node-reset / disk-loss Default action: node-stop",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("node-stop", "node-reset", "disk-loss"),
												},
											},

											"device_names": {
												Description:         "The device name of disks to detach. Needed in disk-loss.",
												MarkdownDescription: "The device name of disks to detach. Needed in disk-loss.",

												Type: types.ListType{ElemType: types.StringType},

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

											"instance": {
												Description:         "Instance defines the name of the instance",
												MarkdownDescription: "Instance defines the name of the instance",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"project": {
												Description:         "Project defines the ID of gcp project.",
												MarkdownDescription: "Project defines the ID of gcp project.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"secret_name": {
												Description:         "SecretName defines the name of kubernetes secret. It is used for GCP credentials.",
												MarkdownDescription: "SecretName defines the name of kubernetes secret. It is used for GCP credentials.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"zone": {
												Description:         "Zone defines the zone of gcp project.",
												MarkdownDescription: "Zone defines the zone of gcp project.",

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

									"history_limit": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"http_chaos": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"abort": {
												Description:         "Abort is a rule to abort a http session.",
												MarkdownDescription: "Abort is a rule to abort a http session.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"code": {
												Description:         "Code is a rule to select target by http status code in response.",
												MarkdownDescription: "Code is a rule to select target by http status code in response.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"delay": {
												Description:         "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
												MarkdownDescription: "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

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

											"method": {
												Description:         "Method is a rule to select target by http method in request.",
												MarkdownDescription: "Method is a rule to select target by http method in request.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mode": {
												Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
												MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
												},
											},

											"patch": {
												Description:         "Patch is a rule to patch some contents in target.",
												MarkdownDescription: "Patch is a rule to patch some contents in target.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"body": {
														Description:         "Body is a rule to patch message body of target.",
														MarkdownDescription: "Body is a rule to patch message body of target.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"type": {
																Description:         "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",
																MarkdownDescription: "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "Value is the patch contents.",
																MarkdownDescription: "Value is the patch contents.",

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

													"headers": {
														Description:         "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",
														MarkdownDescription: "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"queries": {
														Description:         "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",
														MarkdownDescription: "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "Path is a rule to select target by uri path in http request.",
												MarkdownDescription: "Path is a rule to select target by uri path in http request.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Port represents the target port to be proxy of.",
												MarkdownDescription: "Port represents the target port to be proxy of.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"replace": {
												Description:         "Replace is a rule to replace some contents in target.",
												MarkdownDescription: "Replace is a rule to replace some contents in target.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"body": {
														Description:         "Body is a rule to replace http message body in target.",
														MarkdownDescription: "Body is a rule to replace http message body in target.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															validators.Base64Validator(),
														},
													},

													"code": {
														Description:         "Code is a rule to replace http status code in response.",
														MarkdownDescription: "Code is a rule to replace http status code in response.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers": {
														Description:         "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",
														MarkdownDescription: "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"method": {
														Description:         "Method is a rule to replace http method in request.",
														MarkdownDescription: "Method is a rule to replace http method in request.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "Path is rule to to replace uri path in http request.",
														MarkdownDescription: "Path is rule to to replace uri path in http request.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"queries": {
														Description:         "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",
														MarkdownDescription: "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",

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

											"request_headers": {
												Description:         "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",
												MarkdownDescription: "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"response_headers": {
												Description:         "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",
												MarkdownDescription: "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "Selector is used to select pods that are used to inject chaos action.",
												MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression_selectors": {
														Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
														MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "Namespaces is a set of namespace to which objects belong.",
														MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selectors": {
														Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
														MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nodes": {
														Description:         "Nodes is a set of node name and objects must belong to these nodes.",
														MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_phase_selectors": {
														Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
														MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pods": {
														Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
														MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"target": {
												Description:         "Target is the object to be selected and injected.",
												MarkdownDescription: "Target is the object to be selected and injected.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Request", "Response"),
												},
											},

											"value": {
												Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
												MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

									"io_chaos": {
										Description:         "IOChaosSpec defines the desired state of IOChaos",
										MarkdownDescription: "IOChaosSpec defines the desired state of IOChaos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "Action defines the specific pod chaos action. Supported action: latency / fault / attrOverride / mistake",
												MarkdownDescription: "Action defines the specific pod chaos action. Supported action: latency / fault / attrOverride / mistake",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("latency", "fault", "attrOverride", "mistake"),
												},
											},

											"attr": {
												Description:         "Attr defines the overrided attribution",
												MarkdownDescription: "Attr defines the overrided attribution",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"atime": {
														Description:         "Timespec represents a time",
														MarkdownDescription: "Timespec represents a time",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"nsec": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"sec": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"blocks": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ctime": {
														Description:         "Timespec represents a time",
														MarkdownDescription: "Timespec represents a time",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"nsec": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"sec": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"gid": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ino": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "FileType represents type of file",
														MarkdownDescription: "FileType represents type of file",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"mtime": {
														Description:         "Timespec represents a time",
														MarkdownDescription: "Timespec represents a time",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"nsec": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"sec": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nlink": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"perm": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"rdev": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"size": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"uid": {
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

											"container_names": {
												Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
												MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"delay": {
												Description:         "Delay defines the value of I/O chaos action delay. A delay string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
												MarkdownDescription: "Delay defines the value of I/O chaos action delay. A delay string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"duration": {
												Description:         "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
												MarkdownDescription: "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"errno": {
												Description:         "Errno defines the error code that returned by I/O action. refer to: https://www-numi.fnal.gov/offline_software/srt_public_context/WebDocs/Errors/unix_system_errors.html",
												MarkdownDescription: "Errno defines the error code that returned by I/O action. refer to: https://www-numi.fnal.gov/offline_software/srt_public_context/WebDocs/Errors/unix_system_errors.html",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"methods": {
												Description:         "Methods defines the I/O methods for injecting I/O chaos action. default: all I/O methods.",
												MarkdownDescription: "Methods defines the I/O methods for injecting I/O chaos action. default: all I/O methods.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mistake": {
												Description:         "Mistake defines what types of incorrectness are injected to IO operations",
												MarkdownDescription: "Mistake defines what types of incorrectness are injected to IO operations",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"filling": {
														Description:         "Filling determines what is filled in the mistake data.",
														MarkdownDescription: "Filling determines what is filled in the mistake data.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("zero", "random"),
														},
													},

													"max_length": {
														Description:         "Max length of each wrong data segment in bytes",
														MarkdownDescription: "Max length of each wrong data segment in bytes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(1),
														},
													},

													"max_occurrences": {
														Description:         "There will be [1, MaxOccurrences] segments of wrong data.",
														MarkdownDescription: "There will be [1, MaxOccurrences] segments of wrong data.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(1),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mode": {
												Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
												MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
												},
											},

											"path": {
												Description:         "Path defines the path of files for injecting I/O chaos action.",
												MarkdownDescription: "Path defines the path of files for injecting I/O chaos action.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"percent": {
												Description:         "Percent defines the percentage of injection errors and provides a number from 0-100. default: 100.",
												MarkdownDescription: "Percent defines the percentage of injection errors and provides a number from 0-100. default: 100.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "Selector is used to select pods that are used to inject chaos action.",
												MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression_selectors": {
														Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
														MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "Namespaces is a set of namespace to which objects belong.",
														MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selectors": {
														Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
														MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nodes": {
														Description:         "Nodes is a set of node name and objects must belong to these nodes.",
														MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_phase_selectors": {
														Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
														MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pods": {
														Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
														MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
												MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_path": {
												Description:         "VolumePath represents the mount path of injected volume",
												MarkdownDescription: "VolumePath represents the mount path of injected volume",

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

									"jvm_chaos": {
										Description:         "JVMChaosSpec defines the desired state of JVMChaos",
										MarkdownDescription: "JVMChaosSpec defines the desired state of JVMChaos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "Action defines the specific jvm chaos action. Supported action: latency;return;exception;stress;gc;ruleData",
												MarkdownDescription: "Action defines the specific jvm chaos action. Supported action: latency;return;exception;stress;gc;ruleData",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("latency", "return", "exception", "stress", "gc", "ruleData", "mysql"),
												},
											},

											"class": {
												Description:         "Java class",
												MarkdownDescription: "Java class",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"container_names": {
												Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
												MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cpu_count": {
												Description:         "the CPU core number needs to use, only set it when action is stress",
												MarkdownDescription: "the CPU core number needs to use, only set it when action is stress",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"database": {
												Description:         "the match database default value is '', means match all database",
												MarkdownDescription: "the match database default value is '', means match all database",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"duration": {
												Description:         "Duration represents the duration of the chaos action",
												MarkdownDescription: "Duration represents the duration of the chaos action",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"exception": {
												Description:         "the exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
												MarkdownDescription: "the exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"latency": {
												Description:         "the latency duration for action 'latency', unit ms or the latency duration in action 'mysql'",
												MarkdownDescription: "the latency duration for action 'latency', unit ms or the latency duration in action 'mysql'",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mem_type": {
												Description:         "the memory type needs to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
												MarkdownDescription: "the memory type needs to locate, only set it when action is stress, the value can be 'stack' or 'heap'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"method": {
												Description:         "the method in Java class",
												MarkdownDescription: "the method in Java class",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mode": {
												Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
												MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
												},
											},

											"mysql_connector_version": {
												Description:         "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
												MarkdownDescription: "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "byteman rule name, should be unique, and will generate one if not set",
												MarkdownDescription: "byteman rule name, should be unique, and will generate one if not set",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pid": {
												Description:         "the pid of Java process which needs to attach",
												MarkdownDescription: "the pid of Java process which needs to attach",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "the port of agent server, default 9277",
												MarkdownDescription: "the port of agent server, default 9277",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"rule_data": {
												Description:         "the byteman rule's data for action 'ruleData'",
												MarkdownDescription: "the byteman rule's data for action 'ruleData'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "Selector is used to select pods that are used to inject chaos action.",
												MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression_selectors": {
														Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
														MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "Namespaces is a set of namespace to which objects belong.",
														MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selectors": {
														Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
														MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nodes": {
														Description:         "Nodes is a set of node name and objects must belong to these nodes.",
														MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_phase_selectors": {
														Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
														MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pods": {
														Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
														MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"sql_type": {
												Description:         "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
												MarkdownDescription: "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"table": {
												Description:         "the match table default value is '', means match all table",
												MarkdownDescription: "the match table default value is '', means match all table",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
												MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

									"kernel_chaos": {
										Description:         "KernelChaosSpec defines the desired state of KernelChaos",
										MarkdownDescription: "KernelChaosSpec defines the desired state of KernelChaos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"container_names": {
												Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
												MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"duration": {
												Description:         "Duration represents the duration of the chaos action",
												MarkdownDescription: "Duration represents the duration of the chaos action",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"fail_kern_request": {
												Description:         "FailKernRequest defines the request of kernel injection",
												MarkdownDescription: "FailKernRequest defines the request of kernel injection",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"callchain": {
														Description:         "Callchain indicate a special call chain, such as:     ext4_mount       -> mount_subtree          -> ...             -> should_failslab With an optional set of predicates and an optional set of parameters, which used with predicates. You can read call chan and predicate examples from https://github.com/chaos-mesh/bpfki/tree/develop/examples to learn more. If no special call chain, just keep Callchain empty, which means it will fail at any call chain with slab alloc (eg: kmalloc).",
														MarkdownDescription: "Callchain indicate a special call chain, such as:     ext4_mount       -> mount_subtree          -> ...             -> should_failslab With an optional set of predicates and an optional set of parameters, which used with predicates. You can read call chan and predicate examples from https://github.com/chaos-mesh/bpfki/tree/develop/examples to learn more. If no special call chain, just keep Callchain empty, which means it will fail at any call chain with slab alloc (eg: kmalloc).",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"funcname": {
																Description:         "Funcname can be find from kernel source or '/proc/kallsyms', such as 'ext4_mount'",
																MarkdownDescription: "Funcname can be find from kernel source or '/proc/kallsyms', such as 'ext4_mount'",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"parameters": {
																Description:         "Parameters is used with predicate, for example, if you want to inject slab error in 'd_alloc_parallel(struct dentry *parent, const struct qstr *name)' with a special name 'bananas', you need to set it to 'struct dentry *parent, const struct qstr *name' otherwise omit it.",
																MarkdownDescription: "Parameters is used with predicate, for example, if you want to inject slab error in 'd_alloc_parallel(struct dentry *parent, const struct qstr *name)' with a special name 'bananas', you need to set it to 'struct dentry *parent, const struct qstr *name' otherwise omit it.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"predicate": {
																Description:         "Predicate will access the arguments of this Frame, example with Parameters's, you can set it to 'STRNCMP(name->name, 'bananas', 8)' to make inject only with it, or omit it to inject for all d_alloc_parallel call chain.",
																MarkdownDescription: "Predicate will access the arguments of this Frame, example with Parameters's, you can set it to 'STRNCMP(name->name, 'bananas', 8)' to make inject only with it, or omit it to inject for all d_alloc_parallel call chain.",

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

													"failtype": {
														Description:         "FailType indicates what to fail, can be set to '0' / '1' / '2' If '0', indicates slab to fail (should_failslab) If '1', indicates alloc_page to fail (should_fail_alloc_page) If '2', indicates bio to fail (should_fail_bio) You can read:   1. https://www.kernel.org/doc/html/latest/fault-injection/fault-injection.html   2. http://github.com/iovisor/bcc/blob/master/tools/inject_example.txt to learn more",
														MarkdownDescription: "FailType indicates what to fail, can be set to '0' / '1' / '2' If '0', indicates slab to fail (should_failslab) If '1', indicates alloc_page to fail (should_fail_alloc_page) If '2', indicates bio to fail (should_fail_bio) You can read:   1. https://www.kernel.org/doc/html/latest/fault-injection/fault-injection.html   2. http://github.com/iovisor/bcc/blob/master/tools/inject_example.txt to learn more",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),

															int64validator.AtMost(2),
														},
													},

													"headers": {
														Description:         "Headers indicates the appropriate kernel headers you need. Eg: 'linux/mmzone.h', 'linux/blkdev.h' and so on",
														MarkdownDescription: "Headers indicates the appropriate kernel headers you need. Eg: 'linux/mmzone.h', 'linux/blkdev.h' and so on",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"probability": {
														Description:         "Probability indicates the fails with probability. If you want 1%, please set this field with 1.",
														MarkdownDescription: "Probability indicates the fails with probability. If you want 1%, please set this field with 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),

															int64validator.AtMost(100),
														},
													},

													"times": {
														Description:         "Times indicates the max times of fails.",
														MarkdownDescription: "Times indicates the max times of fails.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),
														},
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"mode": {
												Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
												MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
												},
											},

											"selector": {
												Description:         "Selector is used to select pods that are used to inject chaos action.",
												MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression_selectors": {
														Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
														MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "Namespaces is a set of namespace to which objects belong.",
														MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selectors": {
														Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
														MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nodes": {
														Description:         "Nodes is a set of node name and objects must belong to these nodes.",
														MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_phase_selectors": {
														Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
														MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pods": {
														Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
														MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
												MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

									"network_chaos": {
										Description:         "NetworkChaosSpec defines the desired state of NetworkChaos",
										MarkdownDescription: "NetworkChaosSpec defines the desired state of NetworkChaos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "Action defines the specific network chaos action. Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay",
												MarkdownDescription: "Action defines the specific network chaos action. Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("netem", "delay", "loss", "duplicate", "corrupt", "partition", "bandwidth"),
												},
											},

											"bandwidth": {
												Description:         "Bandwidth represents the detail about bandwidth control action",
												MarkdownDescription: "Bandwidth represents the detail about bandwidth control action",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"buffer": {
														Description:         "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",
														MarkdownDescription: "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(1),
														},
													},

													"limit": {
														Description:         "Limit is the number of bytes that can be queued waiting for tokens to become available.",
														MarkdownDescription: "Limit is the number of bytes that can be queued waiting for tokens to become available.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(1),
														},
													},

													"minburst": {
														Description:         "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",
														MarkdownDescription: "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),
														},
													},

													"peakrate": {
														Description:         "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",
														MarkdownDescription: "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),
														},
													},

													"rate": {
														Description:         "Rate is the speed knob. Allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",
														MarkdownDescription: "Rate is the speed knob. Allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",

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

											"corrupt": {
												Description:         "Corrupt represents the detail about corrupt action",
												MarkdownDescription: "Corrupt represents the detail about corrupt action",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"correlation": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"corrupt": {
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

											"delay": {
												Description:         "Delay represents the detail about delay action",
												MarkdownDescription: "Delay represents the detail about delay action",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"correlation": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"jitter": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"latency": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"reorder": {
														Description:         "ReorderSpec defines details of packet reorder.",
														MarkdownDescription: "ReorderSpec defines details of packet reorder.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"correlation": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"gap": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"reorder": {
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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"device": {
												Description:         "Device represents the network device to be affected.",
												MarkdownDescription: "Device represents the network device to be affected.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"direction": {
												Description:         "Direction represents the direction, this applies on netem and network partition action",
												MarkdownDescription: "Direction represents the direction, this applies on netem and network partition action",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("to", "from", "both"),
												},
											},

											"duplicate": {
												Description:         "DuplicateSpec represents the detail about loss action",
												MarkdownDescription: "DuplicateSpec represents the detail about loss action",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"correlation": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"duplicate": {
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

											"duration": {
												Description:         "Duration represents the duration of the chaos action",
												MarkdownDescription: "Duration represents the duration of the chaos action",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"external_targets": {
												Description:         "ExternalTargets represents network targets outside k8s",
												MarkdownDescription: "ExternalTargets represents network targets outside k8s",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"loss": {
												Description:         "Loss represents the detail about loss action",
												MarkdownDescription: "Loss represents the detail about loss action",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"correlation": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"loss": {
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

											"mode": {
												Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
												MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
												},
											},

											"selector": {
												Description:         "Selector is used to select pods that are used to inject chaos action.",
												MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression_selectors": {
														Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
														MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "Namespaces is a set of namespace to which objects belong.",
														MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selectors": {
														Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
														MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nodes": {
														Description:         "Nodes is a set of node name and objects must belong to these nodes.",
														MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_phase_selectors": {
														Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
														MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pods": {
														Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
														MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"target": {
												Description:         "Target represents network target, this applies on netem and network partition action",
												MarkdownDescription: "Target represents network target, this applies on netem and network partition action",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"mode": {
														Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
														MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
														},
													},

													"selector": {
														Description:         "Selector is used to select pods that are used to inject chaos action.",
														MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"annotation_selectors": {
																Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
																MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"expression_selectors": {
																Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
																MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"field_selectors": {
																Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
																MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"label_selectors": {
																Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
																MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespaces": {
																Description:         "Namespaces is a set of namespace to which objects belong.",
																MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"node_selectors": {
																Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
																MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"nodes": {
																Description:         "Nodes is a set of node name and objects must belong to these nodes.",
																MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"pod_phase_selectors": {
																Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
																MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"pods": {
																Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
																MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

																Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
														MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

											"target_device": {
												Description:         "TargetDevice represents the network device to be affected in target scope.",
												MarkdownDescription: "TargetDevice represents the network device to be affected in target scope.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
												MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

									"physicalmachine_chaos": {
										Description:         "PhysicalMachineChaosSpec defines the desired state of PhysicalMachineChaos",
										MarkdownDescription: "PhysicalMachineChaosSpec defines the desired state of PhysicalMachineChaos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "the subAction, generate automatically",
												MarkdownDescription: "the subAction, generate automatically",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("stress-cpu", "stress-mem", "disk-read-payload", "disk-write-payload", "disk-fill", "network-corrupt", "network-duplicate", "network-loss", "network-delay", "network-partition", "network-dns", "network-bandwidth", "network-flood", "network-down", "process", "jvm-exception", "jvm-gc", "jvm-latency", "jvm-return", "jvm-stress", "jvm-rule-data", "jvm-mysql", "clock", "redis-expiration", "redis-penetration", "redis-cacheLimit", "redis-restart", "redis-stop", "kafka-fill", "kafka-flood", "kafka-io", "file-create", "file-modify", "file-delete", "file-rename", "file-append", "file-replace", "vm", "user_defined"),
												},
											},

											"address": {
												Description:         "DEPRECATED: Use Selector instead. Only one of Address and Selector could be specified.",
												MarkdownDescription: "DEPRECATED: Use Selector instead. Only one of Address and Selector could be specified.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"clock": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"clock_ids_slice": {
														Description:         "the identifier of the particular clock on which to act. More clock description in linux kernel can be found in man page of clock_getres, clock_gettime, clock_settime. Muti clock ids should be split with ','",
														MarkdownDescription: "the identifier of the particular clock on which to act. More clock description in linux kernel can be found in man page of clock_getres, clock_gettime, clock_settime. Muti clock ids should be split with ','",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pid": {
														Description:         "the pid of target program.",
														MarkdownDescription: "the pid of target program.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"time_offset": {
														Description:         "specifies the length of time offset.",
														MarkdownDescription: "specifies the length of time offset.",

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

											"disk_fill": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fill_by_fallocate": {
														Description:         "fill disk by fallocate",
														MarkdownDescription: "fill disk by fallocate",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
														MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"size": {
														Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
														MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",

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

											"disk_read_payload": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"path": {
														Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
														MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"payload_process_num": {
														Description:         "specifies the number of process work on writing, default 1, only 1-255 is valid value",
														MarkdownDescription: "specifies the number of process work on writing, default 1, only 1-255 is valid value",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"size": {
														Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
														MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",

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

											"disk_write_payload": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"path": {
														Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
														MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"payload_process_num": {
														Description:         "specifies the number of process work on writing, default 1, only 1-255 is valid value",
														MarkdownDescription: "specifies the number of process work on writing, default 1, only 1-255 is valid value",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"size": {
														Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
														MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",

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

											"duration": {
												Description:         "Duration represents the duration of the chaos action",
												MarkdownDescription: "Duration represents the duration of the chaos action",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"file_append": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"count": {
														Description:         "Count is the number of times to append the data.",
														MarkdownDescription: "Count is the number of times to append the data.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"data": {
														Description:         "Data is the data for append.",
														MarkdownDescription: "Data is the data for append.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"file_name": {
														Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
														MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

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

											"file_create": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"dir_name": {
														Description:         "DirName is the directory name to create or delete.",
														MarkdownDescription: "DirName is the directory name to create or delete.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"file_name": {
														Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
														MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

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

											"file_delete": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"dir_name": {
														Description:         "DirName is the directory name to create or delete.",
														MarkdownDescription: "DirName is the directory name to create or delete.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"file_name": {
														Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
														MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

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

											"file_modify": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"file_name": {
														Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
														MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"privilege": {
														Description:         "Privilege is the file privilege to be set.",
														MarkdownDescription: "Privilege is the file privilege to be set.",

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

											"file_rename": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"dest_file": {
														Description:         "DestFile is the name to be renamed.",
														MarkdownDescription: "DestFile is the name to be renamed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source_file": {
														Description:         "SourceFile is the name need to be renamed.",
														MarkdownDescription: "SourceFile is the name need to be renamed.",

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

											"file_replace": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"dest_string": {
														Description:         "DestStr is the destination string of the file.",
														MarkdownDescription: "DestStr is the destination string of the file.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"file_name": {
														Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
														MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"line": {
														Description:         "Line is the line number of the file to be replaced.",
														MarkdownDescription: "Line is the line number of the file to be replaced.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"origin_string": {
														Description:         "OriginStr is the origin string of the file.",
														MarkdownDescription: "OriginStr is the origin string of the file.",

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

											"http_abort": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"code": {
														Description:         "Code is a rule to select target by http status code in response",
														MarkdownDescription: "Code is a rule to select target by http status code in response",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"method": {
														Description:         "HTTP method",
														MarkdownDescription: "HTTP method",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "Match path of Uri with wildcard matches",
														MarkdownDescription: "Match path of Uri with wildcard matches",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "The TCP port that the target service listens on",
														MarkdownDescription: "The TCP port that the target service listens on",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_ports": {
														Description:         "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
														MarkdownDescription: "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"target": {
														Description:         "HTTP target: Request or Response",
														MarkdownDescription: "HTTP target: Request or Response",

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

											"http_config": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"file_path": {
														Description:         "The config file path",
														MarkdownDescription: "The config file path",

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

											"http_delay": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"code": {
														Description:         "Code is a rule to select target by http status code in response",
														MarkdownDescription: "Code is a rule to select target by http status code in response",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"delay": {
														Description:         "Delay represents the delay of the target request/response",
														MarkdownDescription: "Delay represents the delay of the target request/response",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"method": {
														Description:         "HTTP method",
														MarkdownDescription: "HTTP method",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "Match path of Uri with wildcard matches",
														MarkdownDescription: "Match path of Uri with wildcard matches",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "The TCP port that the target service listens on",
														MarkdownDescription: "The TCP port that the target service listens on",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_ports": {
														Description:         "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
														MarkdownDescription: "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"target": {
														Description:         "HTTP target: Request or Response",
														MarkdownDescription: "HTTP target: Request or Response",

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

											"http_request": {
												Description:         "used for HTTP request, now only support GET",
												MarkdownDescription: "used for HTTP request, now only support GET",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"count": {
														Description:         "The number of requests to send",
														MarkdownDescription: "The number of requests to send",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"enable_conn_pool": {
														Description:         "Enable connection pool",
														MarkdownDescription: "Enable connection pool",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"url": {
														Description:         "Request to send'",
														MarkdownDescription: "Request to send'",

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

											"jvm_exception": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"class": {
														Description:         "Java class",
														MarkdownDescription: "Java class",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exception": {
														Description:         "the exception which needs to throw for action 'exception'",
														MarkdownDescription: "the exception which needs to throw for action 'exception'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"method": {
														Description:         "the method in Java class",
														MarkdownDescription: "the method in Java class",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pid": {
														Description:         "the pid of Java process which needs to attach",
														MarkdownDescription: "the pid of Java process which needs to attach",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "the port of agent server, default 9277",
														MarkdownDescription: "the port of agent server, default 9277",

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

											"jvm_gc": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"pid": {
														Description:         "the pid of Java process which needs to attach",
														MarkdownDescription: "the pid of Java process which needs to attach",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "the port of agent server, default 9277",
														MarkdownDescription: "the port of agent server, default 9277",

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

											"jvm_latency": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"class": {
														Description:         "Java class",
														MarkdownDescription: "Java class",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"latency": {
														Description:         "the latency duration for action 'latency', unit ms",
														MarkdownDescription: "the latency duration for action 'latency', unit ms",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"method": {
														Description:         "the method in Java class",
														MarkdownDescription: "the method in Java class",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pid": {
														Description:         "the pid of Java process which needs to attach",
														MarkdownDescription: "the pid of Java process which needs to attach",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "the port of agent server, default 9277",
														MarkdownDescription: "the port of agent server, default 9277",

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

											"jvm_mysql": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"database": {
														Description:         "the match database default value is '', means match all database",
														MarkdownDescription: "the match database default value is '', means match all database",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exception": {
														Description:         "The exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
														MarkdownDescription: "The exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"latency": {
														Description:         "The latency duration for action 'latency' or the latency duration in action 'mysql'",
														MarkdownDescription: "The latency duration for action 'latency' or the latency duration in action 'mysql'",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"mysql_connector_version": {
														Description:         "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
														MarkdownDescription: "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pid": {
														Description:         "the pid of Java process which needs to attach",
														MarkdownDescription: "the pid of Java process which needs to attach",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "the port of agent server, default 9277",
														MarkdownDescription: "the port of agent server, default 9277",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sql_type": {
														Description:         "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
														MarkdownDescription: "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"table": {
														Description:         "the match table default value is '', means match all table",
														MarkdownDescription: "the match table default value is '', means match all table",

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

											"jvm_return": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"class": {
														Description:         "Java class",
														MarkdownDescription: "Java class",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"method": {
														Description:         "the method in Java class",
														MarkdownDescription: "the method in Java class",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pid": {
														Description:         "the pid of Java process which needs to attach",
														MarkdownDescription: "the pid of Java process which needs to attach",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "the port of agent server, default 9277",
														MarkdownDescription: "the port of agent server, default 9277",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "the return value for action 'return'",
														MarkdownDescription: "the return value for action 'return'",

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

											"jvm_rule_data": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"pid": {
														Description:         "the pid of Java process which needs to attach",
														MarkdownDescription: "the pid of Java process which needs to attach",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "the port of agent server, default 9277",
														MarkdownDescription: "the port of agent server, default 9277",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"rule_data": {
														Description:         "RuleData used to save the rule file's data, will use it when recover",
														MarkdownDescription: "RuleData used to save the rule file's data, will use it when recover",

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

											"jvm_stress": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cpu_count": {
														Description:         "the CPU core number need to use, only set it when action is stress",
														MarkdownDescription: "the CPU core number need to use, only set it when action is stress",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"mem_type": {
														Description:         "the memory type need to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
														MarkdownDescription: "the memory type need to locate, only set it when action is stress, the value can be 'stack' or 'heap'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pid": {
														Description:         "the pid of Java process which needs to attach",
														MarkdownDescription: "the pid of Java process which needs to attach",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "the port of agent server, default 9277",
														MarkdownDescription: "the port of agent server, default 9277",

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

											"kafka_fill": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"host": {
														Description:         "The host of kafka server",
														MarkdownDescription: "The host of kafka server",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_bytes": {
														Description:         "The max bytes to fill",
														MarkdownDescription: "The max bytes to fill",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"message_size": {
														Description:         "The size of each message",
														MarkdownDescription: "The size of each message",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"password": {
														Description:         "The password of kafka client",
														MarkdownDescription: "The password of kafka client",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "The port of kafka server",
														MarkdownDescription: "The port of kafka server",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"reload_command": {
														Description:         "The command to reload kafka config",
														MarkdownDescription: "The command to reload kafka config",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topic": {
														Description:         "The topic to attack",
														MarkdownDescription: "The topic to attack",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"username": {
														Description:         "The username of kafka client",
														MarkdownDescription: "The username of kafka client",

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

											"kafka_flood": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"host": {
														Description:         "The host of kafka server",
														MarkdownDescription: "The host of kafka server",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"message_size": {
														Description:         "The size of each message",
														MarkdownDescription: "The size of each message",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"password": {
														Description:         "The password of kafka client",
														MarkdownDescription: "The password of kafka client",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "The port of kafka server",
														MarkdownDescription: "The port of kafka server",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"threads": {
														Description:         "The number of worker threads",
														MarkdownDescription: "The number of worker threads",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topic": {
														Description:         "The topic to attack",
														MarkdownDescription: "The topic to attack",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"username": {
														Description:         "The username of kafka client",
														MarkdownDescription: "The username of kafka client",

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

											"kafka_io": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"config_file": {
														Description:         "The path of server config",
														MarkdownDescription: "The path of server config",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"non_readable": {
														Description:         "Make kafka cluster non-readable",
														MarkdownDescription: "Make kafka cluster non-readable",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"non_writable": {
														Description:         "Make kafka cluster non-writable",
														MarkdownDescription: "Make kafka cluster non-writable",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topic": {
														Description:         "The topic to attack",
														MarkdownDescription: "The topic to attack",

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

											"mode": {
												Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
												MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
												},
											},

											"network_bandwidth": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"buffer": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(1),
														},
													},

													"device": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"hostname": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ip_address": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"limit": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(1),
														},
													},

													"minburst": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"peakrate": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"rate": {
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

											"network_corrupt": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"correlation": {
														Description:         "correlation is percentage (10 is 10%)",
														MarkdownDescription: "correlation is percentage (10 is 10%)",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"device": {
														Description:         "the network interface to impact",
														MarkdownDescription: "the network interface to impact",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"egress_port": {
														Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
														MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"hostname": {
														Description:         "only impact traffic to these hostnames",
														MarkdownDescription: "only impact traffic to these hostnames",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ip_address": {
														Description:         "only impact egress traffic to these IP addresses",
														MarkdownDescription: "only impact egress traffic to these IP addresses",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ip_protocol": {
														Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
														MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"percent": {
														Description:         "percentage of packets to corrupt (10 is 10%)",
														MarkdownDescription: "percentage of packets to corrupt (10 is 10%)",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source_port": {
														Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
														MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

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

											"network_delay": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"accept_tcp_flags": {
														Description:         "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
														MarkdownDescription: "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"correlation": {
														Description:         "correlation is percentage (10 is 10%)",
														MarkdownDescription: "correlation is percentage (10 is 10%)",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"device": {
														Description:         "the network interface to impact",
														MarkdownDescription: "the network interface to impact",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"egress_port": {
														Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
														MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"hostname": {
														Description:         "only impact traffic to these hostnames",
														MarkdownDescription: "only impact traffic to these hostnames",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ip_address": {
														Description:         "only impact egress traffic to these IP addresses",
														MarkdownDescription: "only impact egress traffic to these IP addresses",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ip_protocol": {
														Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
														MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"jitter": {
														Description:         "jitter time, time units: ns, us (or µs), ms, s, m, h.",
														MarkdownDescription: "jitter time, time units: ns, us (or µs), ms, s, m, h.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"latency": {
														Description:         "delay egress time, time units: ns, us (or µs), ms, s, m, h.",
														MarkdownDescription: "delay egress time, time units: ns, us (or µs), ms, s, m, h.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source_port": {
														Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
														MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

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

											"network_dns": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"dns_domain_name": {
														Description:         "map this host to specified IP",
														MarkdownDescription: "map this host to specified IP",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"dns_ip": {
														Description:         "map specified host to this IP address",
														MarkdownDescription: "map specified host to this IP address",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"dns_server": {
														Description:         "update the DNS server in /etc/resolv.conf with this value",
														MarkdownDescription: "update the DNS server in /etc/resolv.conf with this value",

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

											"network_down": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"device": {
														Description:         "The network interface to impact",
														MarkdownDescription: "The network interface to impact",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"duration": {
														Description:         "NIC down time, time units: ns, us (or µs), ms, s, m, h.",
														MarkdownDescription: "NIC down time, time units: ns, us (or µs), ms, s, m, h.",

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

											"network_duplicate": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"correlation": {
														Description:         "correlation is percentage (10 is 10%)",
														MarkdownDescription: "correlation is percentage (10 is 10%)",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"device": {
														Description:         "the network interface to impact",
														MarkdownDescription: "the network interface to impact",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"egress_port": {
														Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
														MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"hostname": {
														Description:         "only impact traffic to these hostnames",
														MarkdownDescription: "only impact traffic to these hostnames",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ip_address": {
														Description:         "only impact egress traffic to these IP addresses",
														MarkdownDescription: "only impact egress traffic to these IP addresses",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ip_protocol": {
														Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
														MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"percent": {
														Description:         "percentage of packets to duplicate (10 is 10%)",
														MarkdownDescription: "percentage of packets to duplicate (10 is 10%)",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source_port": {
														Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
														MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

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

											"network_flood": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"duration": {
														Description:         "The number of seconds to run the iperf test",
														MarkdownDescription: "The number of seconds to run the iperf test",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"ip_address": {
														Description:         "Generate traffic to this IP address",
														MarkdownDescription: "Generate traffic to this IP address",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"parallel": {
														Description:         "The number of iperf parallel client threads to run",
														MarkdownDescription: "The number of iperf parallel client threads to run",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "Generate traffic to this port on the IP address",
														MarkdownDescription: "Generate traffic to this port on the IP address",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"rate": {
														Description:         "The speed of network traffic, allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second",
														MarkdownDescription: "The speed of network traffic, allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second",

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

											"network_loss": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"correlation": {
														Description:         "correlation is percentage (10 is 10%)",
														MarkdownDescription: "correlation is percentage (10 is 10%)",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"device": {
														Description:         "the network interface to impact",
														MarkdownDescription: "the network interface to impact",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"egress_port": {
														Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
														MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"hostname": {
														Description:         "only impact traffic to these hostnames",
														MarkdownDescription: "only impact traffic to these hostnames",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ip_address": {
														Description:         "only impact egress traffic to these IP addresses",
														MarkdownDescription: "only impact egress traffic to these IP addresses",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ip_protocol": {
														Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
														MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"percent": {
														Description:         "percentage of packets to loss (10 is 10%)",
														MarkdownDescription: "percentage of packets to loss (10 is 10%)",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"source_port": {
														Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
														MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

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

											"network_partition": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"accept_tcp_flags": {
														Description:         "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
														MarkdownDescription: "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"device": {
														Description:         "the network interface to impact",
														MarkdownDescription: "the network interface to impact",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"direction": {
														Description:         "specifies the partition direction, values can be 'from', 'to'. 'from' means packets coming from the 'IPAddress' or 'Hostname' and going to your server, 'to' means packets originating from your server and going to the 'IPAddress' or 'Hostname'.",
														MarkdownDescription: "specifies the partition direction, values can be 'from', 'to'. 'from' means packets coming from the 'IPAddress' or 'Hostname' and going to your server, 'to' means packets originating from your server and going to the 'IPAddress' or 'Hostname'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"hostname": {
														Description:         "only impact traffic to these hostnames",
														MarkdownDescription: "only impact traffic to these hostnames",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ip_address": {
														Description:         "only impact egress traffic to these IP addresses",
														MarkdownDescription: "only impact egress traffic to these IP addresses",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ip_protocol": {
														Description:         "only impact egress traffic to these IP addresses",
														MarkdownDescription: "only impact egress traffic to these IP addresses",

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

											"process": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"process": {
														Description:         "the process name or the process ID",
														MarkdownDescription: "the process name or the process ID",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"recover_cmd": {
														Description:         "the command to be run when recovering experiment",
														MarkdownDescription: "the command to be run when recovering experiment",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"signal": {
														Description:         "the signal number to send",
														MarkdownDescription: "the signal number to send",

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

											"redis_cache_limit": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"addr": {
														Description:         "The adress of Redis server",
														MarkdownDescription: "The adress of Redis server",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"cache_size": {
														Description:         "The size of 'maxmemory'",
														MarkdownDescription: "The size of 'maxmemory'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"password": {
														Description:         "The password of Redis server",
														MarkdownDescription: "The password of Redis server",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"percent": {
														Description:         "Specifies maxmemory as a percentage of the original value",
														MarkdownDescription: "Specifies maxmemory as a percentage of the original value",

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

											"redis_expiration": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"addr": {
														Description:         "The adress of Redis server",
														MarkdownDescription: "The adress of Redis server",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expiration": {
														Description:         "The expiration of the keys",
														MarkdownDescription: "The expiration of the keys",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key": {
														Description:         "The keys to be expired",
														MarkdownDescription: "The keys to be expired",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"option": {
														Description:         "Additional options for 'expiration'",
														MarkdownDescription: "Additional options for 'expiration'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"password": {
														Description:         "The password of Redis server",
														MarkdownDescription: "The password of Redis server",

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

											"redis_penetration": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"addr": {
														Description:         "The adress of Redis server",
														MarkdownDescription: "The adress of Redis server",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"password": {
														Description:         "The password of Redis server",
														MarkdownDescription: "The password of Redis server",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"request_num": {
														Description:         "The number of requests to be sent",
														MarkdownDescription: "The number of requests to be sent",

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

											"redis_restart": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"addr": {
														Description:         "The adress of Redis server",
														MarkdownDescription: "The adress of Redis server",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"conf": {
														Description:         "The path of Sentinel conf",
														MarkdownDescription: "The path of Sentinel conf",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"flush_config": {
														Description:         "The control flag determines whether to flush config",
														MarkdownDescription: "The control flag determines whether to flush config",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"password": {
														Description:         "The password of Redis server",
														MarkdownDescription: "The password of Redis server",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"redis_path": {
														Description:         "The path of 'redis-server' command-line tool",
														MarkdownDescription: "The path of 'redis-server' command-line tool",

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

											"redis_stop": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"addr": {
														Description:         "The adress of Redis server",
														MarkdownDescription: "The adress of Redis server",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"conf": {
														Description:         "The path of Sentinel conf",
														MarkdownDescription: "The path of Sentinel conf",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"flush_config": {
														Description:         "The control flag determines whether to flush config",
														MarkdownDescription: "The control flag determines whether to flush config",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"password": {
														Description:         "The password of Redis server",
														MarkdownDescription: "The password of Redis server",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"redis_path": {
														Description:         "The path of 'redis-server' command-line tool",
														MarkdownDescription: "The path of 'redis-server' command-line tool",

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

											"selector": {
												Description:         "Selector is used to select physical machines that are used to inject chaos action.",
												MarkdownDescription: "Selector is used to select physical machines that are used to inject chaos action.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression_selectors": {
														Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
														MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "Namespaces is a set of namespace to which objects belong.",
														MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"physical_machines": {
														Description:         "PhysicalMachines is a map of string keys and a set values that used to select physical machines. The key defines the namespace which physical machine belong, and each value is a set of physical machine names.",
														MarkdownDescription: "PhysicalMachines is a map of string keys and a set values that used to select physical machines. The key defines the namespace which physical machine belong, and each value is a set of physical machine names.",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"stress_cpu": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"load": {
														Description:         "specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
														MarkdownDescription: "specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"options": {
														Description:         "extend stress-ng options",
														MarkdownDescription: "extend stress-ng options",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"workers": {
														Description:         "specifies N workers to apply the stressor.",
														MarkdownDescription: "specifies N workers to apply the stressor.",

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

											"stress_mem": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"options": {
														Description:         "extend stress-ng options",
														MarkdownDescription: "extend stress-ng options",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"size": {
														Description:         "specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB..",
														MarkdownDescription: "specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB..",

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

											"uid": {
												Description:         "the experiment ID",
												MarkdownDescription: "the experiment ID",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"user_defined": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"attack_cmd": {
														Description:         "The command to be executed when attack",
														MarkdownDescription: "The command to be executed when attack",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"recover_cmd": {
														Description:         "The command to be executed when recover",
														MarkdownDescription: "The command to be executed when recover",

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

											"value": {
												Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of physical machines to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of physical machines the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
												MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of physical machines to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of physical machines the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"vm": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"vm_name": {
														Description:         "The name of the VM to be injected",
														MarkdownDescription: "The name of the VM to be injected",

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

									"pod_chaos": {
										Description:         "PodChaosSpec defines the attributes that a user creates on a chaos experiment about pods.",
										MarkdownDescription: "PodChaosSpec defines the attributes that a user creates on a chaos experiment about pods.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "Action defines the specific pod chaos action. Supported action: pod-kill / pod-failure / container-kill Default action: pod-kill",
												MarkdownDescription: "Action defines the specific pod chaos action. Supported action: pod-kill / pod-failure / container-kill Default action: pod-kill",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("pod-kill", "pod-failure", "container-kill"),
												},
											},

											"container_names": {
												Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
												MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"duration": {
												Description:         "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
												MarkdownDescription: "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"grace_period": {
												Description:         "GracePeriod is used in pod-kill action. It represents the duration in seconds before the pod should be deleted. Value must be non-negative integer. The default value is zero that indicates delete immediately.",
												MarkdownDescription: "GracePeriod is used in pod-kill action. It represents the duration in seconds before the pod should be deleted. Value must be non-negative integer. The default value is zero that indicates delete immediately.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"mode": {
												Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
												MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
												},
											},

											"selector": {
												Description:         "Selector is used to select pods that are used to inject chaos action.",
												MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression_selectors": {
														Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
														MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "Namespaces is a set of namespace to which objects belong.",
														MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selectors": {
														Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
														MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nodes": {
														Description:         "Nodes is a set of node name and objects must belong to these nodes.",
														MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_phase_selectors": {
														Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
														MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pods": {
														Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
														MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
												MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

									"schedule": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"starting_deadline_seconds": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"stress_chaos": {
										Description:         "StressChaosSpec defines the desired state of StressChaos",
										MarkdownDescription: "StressChaosSpec defines the desired state of StressChaos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"container_names": {
												Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
												MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"duration": {
												Description:         "Duration represents the duration of the chaos action",
												MarkdownDescription: "Duration represents the duration of the chaos action",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mode": {
												Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
												MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
												},
											},

											"selector": {
												Description:         "Selector is used to select pods that are used to inject chaos action.",
												MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression_selectors": {
														Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
														MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "Namespaces is a set of namespace to which objects belong.",
														MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selectors": {
														Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
														MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nodes": {
														Description:         "Nodes is a set of node name and objects must belong to these nodes.",
														MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_phase_selectors": {
														Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
														MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pods": {
														Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
														MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"stressng_stressors": {
												Description:         "StressngStressors defines plenty of stressors just like 'Stressors' except that it's an experimental feature and more powerful. You can define stressors in 'stress-ng' (see also 'man stress-ng') dialect, however not all of the supported stressors are well tested. It maybe retired in later releases. You should always use 'Stressors' to define the stressors and use this only when you want more stressors unsupported by 'Stressors'. When both 'StressngStressors' and 'Stressors' are defined, 'StressngStressors' wins.",
												MarkdownDescription: "StressngStressors defines plenty of stressors just like 'Stressors' except that it's an experimental feature and more powerful. You can define stressors in 'stress-ng' (see also 'man stress-ng') dialect, however not all of the supported stressors are well tested. It maybe retired in later releases. You should always use 'Stressors' to define the stressors and use this only when you want more stressors unsupported by 'Stressors'. When both 'StressngStressors' and 'Stressors' are defined, 'StressngStressors' wins.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"stressors": {
												Description:         "Stressors defines plenty of stressors supported to stress system components out. You can use one or more of them to make up various kinds of stresses. At least one of the stressors should be specified.",
												MarkdownDescription: "Stressors defines plenty of stressors supported to stress system components out. You can use one or more of them to make up various kinds of stresses. At least one of the stressors should be specified.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cpu": {
														Description:         "CPUStressor stresses CPU out",
														MarkdownDescription: "CPUStressor stresses CPU out",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"load": {
																Description:         "Load specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
																MarkdownDescription: "Load specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	int64validator.AtLeast(0),

																	int64validator.AtMost(100),
																},
															},

															"options": {
																Description:         "extend stress-ng options",
																MarkdownDescription: "extend stress-ng options",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"workers": {
																Description:         "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
																MarkdownDescription: "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	int64validator.AtMost(8192),
																},
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"memory": {
														Description:         "MemoryStressor stresses virtual memory out",
														MarkdownDescription: "MemoryStressor stresses virtual memory out",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"oom_score_adj": {
																Description:         "OOMScoreAdj sets the oom_score_adj of the stress process. See 'man 5 proc' to know more about this option.",
																MarkdownDescription: "OOMScoreAdj sets the oom_score_adj of the stress process. See 'man 5 proc' to know more about this option.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	int64validator.AtLeast(-1000),

																	int64validator.AtMost(1000),
																},
															},

															"options": {
																Description:         "extend stress-ng options",
																MarkdownDescription: "extend stress-ng options",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"size": {
																Description:         "Size specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB.",
																MarkdownDescription: "Size specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"workers": {
																Description:         "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
																MarkdownDescription: "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	int64validator.AtMost(8192),
																},
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

											"value": {
												Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
												MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

									"time_chaos": {
										Description:         "TimeChaosSpec defines the desired state of TimeChaos",
										MarkdownDescription: "TimeChaosSpec defines the desired state of TimeChaos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"clock_ids": {
												Description:         "ClockIds defines all affected clock id All available options are ['CLOCK_REALTIME','CLOCK_MONOTONIC','CLOCK_PROCESS_CPUTIME_ID','CLOCK_THREAD_CPUTIME_ID', 'CLOCK_MONOTONIC_RAW','CLOCK_REALTIME_COARSE','CLOCK_MONOTONIC_COARSE','CLOCK_BOOTTIME','CLOCK_REALTIME_ALARM', 'CLOCK_BOOTTIME_ALARM'] Default value is ['CLOCK_REALTIME']",
												MarkdownDescription: "ClockIds defines all affected clock id All available options are ['CLOCK_REALTIME','CLOCK_MONOTONIC','CLOCK_PROCESS_CPUTIME_ID','CLOCK_THREAD_CPUTIME_ID', 'CLOCK_MONOTONIC_RAW','CLOCK_REALTIME_COARSE','CLOCK_MONOTONIC_COARSE','CLOCK_BOOTTIME','CLOCK_REALTIME_ALARM', 'CLOCK_BOOTTIME_ALARM'] Default value is ['CLOCK_REALTIME']",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"container_names": {
												Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
												MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"duration": {
												Description:         "Duration represents the duration of the chaos action",
												MarkdownDescription: "Duration represents the duration of the chaos action",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mode": {
												Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
												MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
												},
											},

											"selector": {
												Description:         "Selector is used to select pods that are used to inject chaos action.",
												MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotation_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression_selectors": {
														Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
														MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selectors": {
														Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
														MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespaces": {
														Description:         "Namespaces is a set of namespace to which objects belong.",
														MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selectors": {
														Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
														MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"nodes": {
														Description:         "Nodes is a set of node name and objects must belong to these nodes.",
														MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_phase_selectors": {
														Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
														MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pods": {
														Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
														MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

														Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"time_offset": {
												Description:         "TimeOffset defines the delta time of injected program. It's a possibly signed sequence of decimal numbers, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
												MarkdownDescription: "TimeOffset defines the delta time of injected program. It's a possibly signed sequence of decimal numbers, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
												MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

									"type": {
										Description:         "TODO: use a custom type, as 'TemplateType' contains other possible values",
										MarkdownDescription: "TODO: use a custom type, as 'TemplateType' contains other possible values",

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

							"status_check": {
								Description:         "StatusCheck describe the behavior of StatusCheck. Only used when Type is TypeStatusCheck.",
								MarkdownDescription: "StatusCheck describe the behavior of StatusCheck. Only used when Type is TypeStatusCheck.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"duration": {
										Description:         "Duration defines the duration of the whole status check if the number of failed execution does not exceed the failure threshold. Duration is available to both 'Synchronous' and 'Continuous' mode. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
										MarkdownDescription: "Duration defines the duration of the whole status check if the number of failed execution does not exceed the failure threshold. Duration is available to both 'Synchronous' and 'Continuous' mode. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"failure_threshold": {
										Description:         "FailureThreshold defines the minimum consecutive failure for the status check to be considered failed.",
										MarkdownDescription: "FailureThreshold defines the minimum consecutive failure for the status check to be considered failed.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"http": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"body": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"criteria": {
												Description:         "Criteria defines how to determine the result of the status check.",
												MarkdownDescription: "Criteria defines how to determine the result of the status check.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"status_code": {
														Description:         "StatusCode defines the expected http status code for the request. A statusCode string could be a single code (e.g. 200), or an inclusive range (e.g. 200-400, both '200' and '400' are included).",
														MarkdownDescription: "StatusCode defines the expected http status code for the request. A statusCode string could be a single code (e.g. 200), or an inclusive range (e.g. 200-400, both '200' and '400' are included).",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"headers": {
												Description:         "A Header represents the key-value pairs in an HTTP header.  The keys should be in canonical form, as returned by CanonicalHeaderKey.",
												MarkdownDescription: "A Header represents the key-value pairs in an HTTP header.  The keys should be in canonical form, as returned by CanonicalHeaderKey.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

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

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("GET", "POST"),
												},
											},

											"url": {
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

									"interval_seconds": {
										Description:         "IntervalSeconds defines how often (in seconds) to perform an execution of status check.",
										MarkdownDescription: "IntervalSeconds defines how often (in seconds) to perform an execution of status check.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"mode": {
										Description:         "Mode defines the execution mode of the status check. Support type: Synchronous / Continuous",
										MarkdownDescription: "Mode defines the execution mode of the status check. Support type: Synchronous / Continuous",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Synchronous", "Continuous"),
										},
									},

									"records_history_limit": {
										Description:         "RecordsHistoryLimit defines the number of record to retain.",
										MarkdownDescription: "RecordsHistoryLimit defines the number of record to retain.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),

											int64validator.AtMost(1000),
										},
									},

									"success_threshold": {
										Description:         "SuccessThreshold defines the minimum consecutive successes for the status check to be considered successful. SuccessThreshold only works for 'Synchronous' mode.",
										MarkdownDescription: "SuccessThreshold defines the minimum consecutive successes for the status check to be considered successful. SuccessThreshold only works for 'Synchronous' mode.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": {
										Description:         "TimeoutSeconds defines the number of seconds after which an execution of status check times out.",
										MarkdownDescription: "TimeoutSeconds defines the number of seconds after which an execution of status check times out.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"type": {
										Description:         "Type defines the specific status check type. Support type: HTTP",
										MarkdownDescription: "Type defines the specific status check type. Support type: HTTP",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("HTTP"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"stress_chaos": {
								Description:         "StressChaosSpec defines the desired state of StressChaos",
								MarkdownDescription: "StressChaosSpec defines the desired state of StressChaos",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"container_names": {
										Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
										MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"duration": {
										Description:         "Duration represents the duration of the chaos action",
										MarkdownDescription: "Duration represents the duration of the chaos action",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mode": {
										Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
										MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
										},
									},

									"selector": {
										Description:         "Selector is used to select pods that are used to inject chaos action.",
										MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expression_selectors": {
												Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
												MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"label_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "Namespaces is a set of namespace to which objects belong.",
												MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selectors": {
												Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
												MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nodes": {
												Description:         "Nodes is a set of node name and objects must belong to these nodes.",
												MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_phase_selectors": {
												Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
												MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pods": {
												Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
												MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"stressng_stressors": {
										Description:         "StressngStressors defines plenty of stressors just like 'Stressors' except that it's an experimental feature and more powerful. You can define stressors in 'stress-ng' (see also 'man stress-ng') dialect, however not all of the supported stressors are well tested. It maybe retired in later releases. You should always use 'Stressors' to define the stressors and use this only when you want more stressors unsupported by 'Stressors'. When both 'StressngStressors' and 'Stressors' are defined, 'StressngStressors' wins.",
										MarkdownDescription: "StressngStressors defines plenty of stressors just like 'Stressors' except that it's an experimental feature and more powerful. You can define stressors in 'stress-ng' (see also 'man stress-ng') dialect, however not all of the supported stressors are well tested. It maybe retired in later releases. You should always use 'Stressors' to define the stressors and use this only when you want more stressors unsupported by 'Stressors'. When both 'StressngStressors' and 'Stressors' are defined, 'StressngStressors' wins.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"stressors": {
										Description:         "Stressors defines plenty of stressors supported to stress system components out. You can use one or more of them to make up various kinds of stresses. At least one of the stressors should be specified.",
										MarkdownDescription: "Stressors defines plenty of stressors supported to stress system components out. You can use one or more of them to make up various kinds of stresses. At least one of the stressors should be specified.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cpu": {
												Description:         "CPUStressor stresses CPU out",
												MarkdownDescription: "CPUStressor stresses CPU out",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"load": {
														Description:         "Load specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
														MarkdownDescription: "Load specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),

															int64validator.AtMost(100),
														},
													},

													"options": {
														Description:         "extend stress-ng options",
														MarkdownDescription: "extend stress-ng options",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"workers": {
														Description:         "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
														MarkdownDescription: "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtMost(8192),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory": {
												Description:         "MemoryStressor stresses virtual memory out",
												MarkdownDescription: "MemoryStressor stresses virtual memory out",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"oom_score_adj": {
														Description:         "OOMScoreAdj sets the oom_score_adj of the stress process. See 'man 5 proc' to know more about this option.",
														MarkdownDescription: "OOMScoreAdj sets the oom_score_adj of the stress process. See 'man 5 proc' to know more about this option.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(-1000),

															int64validator.AtMost(1000),
														},
													},

													"options": {
														Description:         "extend stress-ng options",
														MarkdownDescription: "extend stress-ng options",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"size": {
														Description:         "Size specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB.",
														MarkdownDescription: "Size specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"workers": {
														Description:         "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
														MarkdownDescription: "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtMost(8192),
														},
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

									"value": {
										Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
										MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

							"task": {
								Description:         "Task describes the behavior of the custom task. Only used when Type is TypeTask.",
								MarkdownDescription: "Task describes the behavior of the custom task. Only used when Type is TypeTask.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"container": {
										Description:         "Container is the main container image to run in the pod",
										MarkdownDescription: "Container is the main container image to run in the pod",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"args": {
												Description:         "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"command": {
												Description:         "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"env": {
												Description:         "List of environment variables to set in the container. Cannot be updated.",
												MarkdownDescription: "List of environment variables to set in the container. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
														MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value_from": {
														Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
														MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map_key_ref": {
																Description:         "Selects a key of a ConfigMap.",
																MarkdownDescription: "Selects a key of a ConfigMap.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "The key to select.",
																		MarkdownDescription: "The key to select.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"name": {
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "Specify whether the ConfigMap or its key must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

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

															"field_ref": {
																Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"api_version": {
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_path": {
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",

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

															"resource_field_ref": {
																Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"container_name": {
																		Description:         "Container name: required for volumes, optional for env vars",
																		MarkdownDescription: "Container name: required for volumes, optional for env vars",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"divisor": {
																		Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																		MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource": {
																		Description:         "Required: resource to select",
																		MarkdownDescription: "Required: resource to select",

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

															"secret_key_ref": {
																Description:         "Selects a key of a secret in the pod's namespace",
																MarkdownDescription: "Selects a key of a secret in the pod's namespace",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"name": {
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "Specify whether the Secret or its key must be defined",
																		MarkdownDescription: "Specify whether the Secret or its key must be defined",

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

											"env_from": {
												Description:         "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
												MarkdownDescription: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"config_map_ref": {
														Description:         "The ConfigMap to select from",
														MarkdownDescription: "The ConfigMap to select from",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"optional": {
																Description:         "Specify whether the ConfigMap must be defined",
																MarkdownDescription: "Specify whether the ConfigMap must be defined",

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

													"prefix": {
														Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
														MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "The Secret to select from",
														MarkdownDescription: "The Secret to select from",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"optional": {
																Description:         "Specify whether the Secret must be defined",
																MarkdownDescription: "Specify whether the Secret must be defined",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image": {
												Description:         "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
												MarkdownDescription: "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image_pull_policy": {
												Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
												MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"lifecycle": {
												Description:         "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
												MarkdownDescription: "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"post_start": {
														Description:         "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
														MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "Exec specifies the action to take.",
																MarkdownDescription: "Exec specifies the action to take.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																		MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_get": {
																Description:         "HTTPGet specifies the http request to perform.",
																MarkdownDescription: "HTTPGet specifies the http request to perform.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																		MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																		MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "The header field name",
																				MarkdownDescription: "The header field name",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",

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

																	"path": {
																		Description:         "Path to access on the HTTP server.",
																		MarkdownDescription: "Path to access on the HTTP server.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																		MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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

															"tcp_socket": {
																Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																		MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
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

													"pre_stop": {
														Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
														MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "Exec specifies the action to take.",
																MarkdownDescription: "Exec specifies the action to take.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																		MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_get": {
																Description:         "HTTPGet specifies the http request to perform.",
																MarkdownDescription: "HTTPGet specifies the http request to perform.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																		MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																		MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "The header field name",
																				MarkdownDescription: "The header field name",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",

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

																	"path": {
																		Description:         "Path to access on the HTTP server.",
																		MarkdownDescription: "Path to access on the HTTP server.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																		MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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

															"tcp_socket": {
																Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																		MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
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

											"liveness_probe": {
												Description:         "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "Exec specifies the action to take.",
														MarkdownDescription: "Exec specifies the action to take.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"failure_threshold": {
														Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc": {
														Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
														MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service": {
																Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_headers": {
																Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

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

															"path": {
																Description:         "Path to access on the HTTP server.",
																MarkdownDescription: "Path to access on the HTTP server.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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

													"initial_delay_seconds": {
														Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"period_seconds": {
														Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
														MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"success_threshold": {
														Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"termination_grace_period_seconds": {
														Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
														MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout_seconds": {
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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

											"name": {
												Description:         "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
												MarkdownDescription: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"ports": {
												Description:         "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",
												MarkdownDescription: "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"container_port": {
														Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
														MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"host_ip": {
														Description:         "What host IP to bind the external port to.",
														MarkdownDescription: "What host IP to bind the external port to.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"host_port": {
														Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
														MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
														MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": {
														Description:         "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
														MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",

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

											"readiness_probe": {
												Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "Exec specifies the action to take.",
														MarkdownDescription: "Exec specifies the action to take.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"failure_threshold": {
														Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc": {
														Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
														MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service": {
																Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_headers": {
																Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

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

															"path": {
																Description:         "Path to access on the HTTP server.",
																MarkdownDescription: "Path to access on the HTTP server.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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

													"initial_delay_seconds": {
														Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"period_seconds": {
														Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
														MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"success_threshold": {
														Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"termination_grace_period_seconds": {
														Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
														MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout_seconds": {
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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

											"resources": {
												Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"limits": {
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"requests": {
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

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

											"security_context": {
												Description:         "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
												MarkdownDescription: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"allow_privilege_escalation": {
														Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"capabilities": {
														Description:         "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"add": {
																Description:         "Added capabilities",
																MarkdownDescription: "Added capabilities",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"drop": {
																Description:         "Removed capabilities",
																MarkdownDescription: "Removed capabilities",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"privileged": {
														Description:         "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"proc_mount": {
														Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only_root_filesystem": {
														Description:         "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_group": {
														Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_non_root": {
														Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user": {
														Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"se_linux_options": {
														Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"level": {
																Description:         "Level is SELinux level label that applies to the container.",
																MarkdownDescription: "Level is SELinux level label that applies to the container.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "Role is a SELinux role label that applies to the container.",
																MarkdownDescription: "Role is a SELinux role label that applies to the container.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "Type is a SELinux type label that applies to the container.",
																MarkdownDescription: "Type is a SELinux type label that applies to the container.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"user": {
																Description:         "User is a SELinux user label that applies to the container.",
																MarkdownDescription: "User is a SELinux user label that applies to the container.",

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

													"seccomp_profile": {
														Description:         "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"localhost_profile": {
																Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",
																MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
																MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",

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

													"windows_options": {
														Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
														MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"gmsa_credential_spec": {
																Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"gmsa_credential_spec_name": {
																Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"host_process": {
																Description:         "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",
																MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_user_name": {
																Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

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

											"startup_probe": {
												Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "Exec specifies the action to take.",
														MarkdownDescription: "Exec specifies the action to take.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"failure_threshold": {
														Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc": {
														Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
														MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service": {
																Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_headers": {
																Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

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

															"path": {
																Description:         "Path to access on the HTTP server.",
																MarkdownDescription: "Path to access on the HTTP server.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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

													"initial_delay_seconds": {
														Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"period_seconds": {
														Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
														MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"success_threshold": {
														Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"termination_grace_period_seconds": {
														Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
														MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout_seconds": {
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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

											"stdin": {
												Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
												MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"stdin_once": {
												Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
												MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_message_path": {
												Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
												MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_message_policy": {
												Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
												MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tty": {
												Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
												MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_devices": {
												Description:         "volumeDevices is the list of block devices to be used by the container.",
												MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"device_path": {
														Description:         "devicePath is the path inside of the container that the device will be mapped to.",
														MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "name must match the name of a persistentVolumeClaim in the pod",
														MarkdownDescription: "name must match the name of a persistentVolumeClaim in the pod",

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

											"volume_mounts": {
												Description:         "Pod volumes to mount into the container's filesystem. Cannot be updated.",
												MarkdownDescription: "Pod volumes to mount into the container's filesystem. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"mount_path": {
														Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
														MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"mount_propagation": {
														Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
														MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "This must match the Name of a Volume.",
														MarkdownDescription: "This must match the Name of a Volume.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
														MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path": {
														Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
														MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path_expr": {
														Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
														MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",

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

											"working_dir": {
												Description:         "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
												MarkdownDescription: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",

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

									"volumes": {
										Description:         "Volumes is a list of volumes that can be mounted by containers in a template.",
										MarkdownDescription: "Volumes is a list of volumes that can be mounted by containers in a template.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"aws_elastic_block_store": {
												Description:         "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
												MarkdownDescription: "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"partition": {
														Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
														MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
														MarkdownDescription: "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_id": {
														Description:         "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
														MarkdownDescription: "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

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

											"azure_disk": {
												Description:         "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
												MarkdownDescription: "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"caching_mode": {
														Description:         "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
														MarkdownDescription: "cachingMode is the Host Caching mode: None, Read Only, Read Write.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"disk_name": {
														Description:         "diskName is the Name of the data disk in the blob storage",
														MarkdownDescription: "diskName is the Name of the data disk in the blob storage",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"disk_uri": {
														Description:         "diskURI is the URI of data disk in the blob storage",
														MarkdownDescription: "diskURI is the URI of data disk in the blob storage",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"fs_type": {
														Description:         "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
														MarkdownDescription: "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

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

											"azure_file": {
												Description:         "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
												MarkdownDescription: "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"read_only": {
														Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_name": {
														Description:         "secretName is the  name of secret that contains Azure Storage Account Name and Key",
														MarkdownDescription: "secretName is the  name of secret that contains Azure Storage Account Name and Key",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"share_name": {
														Description:         "shareName is the azure share Name",
														MarkdownDescription: "shareName is the azure share Name",

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

											"cephfs": {
												Description:         "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
												MarkdownDescription: "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"monitors": {
														Description:         "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"path": {
														Description:         "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
														MarkdownDescription: "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_file": {
														Description:         "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"user": {
														Description:         "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

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

											"cinder": {
												Description:         "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
												MarkdownDescription: "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
														MarkdownDescription: "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"volume_id": {
														Description:         "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

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

											"config_map": {
												Description:         "configMap represents a configMap that should populate this volume",
												MarkdownDescription: "configMap represents a configMap that should populate this volume",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"items": {
														Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
														MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the key to project.",
																MarkdownDescription: "key is the key to project.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mode": {
																Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "optional specify whether the ConfigMap or its keys must be defined",
														MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",

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

											"csi": {
												Description:         "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
												MarkdownDescription: "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"driver": {
														Description:         "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
														MarkdownDescription: "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"fs_type": {
														Description:         "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
														MarkdownDescription: "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_publish_secret_ref": {
														Description:         "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
														MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"read_only": {
														Description:         "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
														MarkdownDescription: "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_attributes": {
														Description:         "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
														MarkdownDescription: "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",

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

											"downward_api": {
												Description:         "downwardAPI represents downward API about the pod that should populate this volume",
												MarkdownDescription: "downwardAPI represents downward API about the pod that should populate this volume",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"items": {
														Description:         "Items is a list of downward API volume file",
														MarkdownDescription: "Items is a list of downward API volume file",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"field_ref": {
																Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"api_version": {
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_path": {
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",

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

															"mode": {
																Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"resource_field_ref": {
																Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"container_name": {
																		Description:         "Container name: required for volumes, optional for env vars",
																		MarkdownDescription: "Container name: required for volumes, optional for env vars",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"divisor": {
																		Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																		MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource": {
																		Description:         "Required: resource to select",
																		MarkdownDescription: "Required: resource to select",

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

											"empty_dir": {
												Description:         "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
												MarkdownDescription: "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"medium": {
														Description:         "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
														MarkdownDescription: "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"size_limit": {
														Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
														MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",

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

											"ephemeral": {
												Description:         "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
												MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"volume_claim_template": {
														Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
														MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"spec": {
																Description:         "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
																MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"access_modes": {
																		Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																		MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"data_source": {
																		Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
																		MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_group": {
																				Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																				MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "Kind is the type of resource being referenced",
																				MarkdownDescription: "Kind is the type of resource being referenced",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name is the name of resource being referenced",
																				MarkdownDescription: "Name is the name of resource being referenced",

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

																	"data_source_ref": {
																		Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
																		MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_group": {
																				Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																				MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "Kind is the type of resource being referenced",
																				MarkdownDescription: "Kind is the type of resource being referenced",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name is the name of resource being referenced",
																				MarkdownDescription: "Name is the name of resource being referenced",

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

																	"resources": {
																		Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																		MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"limits": {
																				Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																				MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"requests": {
																				Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																				MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

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

																	"selector": {
																		Description:         "selector is a label query over volumes to consider for binding.",
																		MarkdownDescription: "selector is a label query over volumes to consider for binding.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																				MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "key is the label key that the selector applies to.",
																						MarkdownDescription: "key is the label key that the selector applies to.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"operator": {
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"match_labels": {
																				Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

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

																	"storage_class_name": {
																		Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																		MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"volume_mode": {
																		Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																		MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"volume_name": {
																		Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
																		MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",

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

											"fc": {
												Description:         "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
												MarkdownDescription: "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"lun": {
														Description:         "lun is Optional: FC target lun number",
														MarkdownDescription: "lun is Optional: FC target lun number",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"target_ww_ns": {
														Description:         "targetWWNs is Optional: FC target worldwide names (WWNs)",
														MarkdownDescription: "targetWWNs is Optional: FC target worldwide names (WWNs)",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"wwids": {
														Description:         "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
														MarkdownDescription: "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"flex_volume": {
												Description:         "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
												MarkdownDescription: "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"driver": {
														Description:         "driver is the name of the driver to use for this volume.",
														MarkdownDescription: "driver is the name of the driver to use for this volume.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"fs_type": {
														Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
														MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"options": {
														Description:         "options is Optional: this field holds extra command options if any.",
														MarkdownDescription: "options is Optional: this field holds extra command options if any.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
														MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

											"flocker": {
												Description:         "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
												MarkdownDescription: "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"dataset_name": {
														Description:         "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
														MarkdownDescription: "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"dataset_uuid": {
														Description:         "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
														MarkdownDescription: "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",

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

											"gce_persistent_disk": {
												Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
												MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"partition": {
														Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pd_name": {
														Description:         "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

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

											"git_repo": {
												Description:         "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
												MarkdownDescription: "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"directory": {
														Description:         "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
														MarkdownDescription: "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"repository": {
														Description:         "repository is the URL",
														MarkdownDescription: "repository is the URL",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"revision": {
														Description:         "revision is the commit hash for the specified revision.",
														MarkdownDescription: "revision is the commit hash for the specified revision.",

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

											"glusterfs": {
												Description:         "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
												MarkdownDescription: "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"endpoints": {
														Description:         "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"path": {
														Description:         "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														MarkdownDescription: "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														MarkdownDescription: "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

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

											"host_path": {
												Description:         "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
												MarkdownDescription: "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"path": {
														Description:         "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
														MarkdownDescription: "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"type": {
														Description:         "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
														MarkdownDescription: "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

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

											"iscsi": {
												Description:         "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
												MarkdownDescription: "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"chap_auth_discovery": {
														Description:         "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
														MarkdownDescription: "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"chap_auth_session": {
														Description:         "chapAuthSession defines whether support iSCSI Session CHAP authentication",
														MarkdownDescription: "chapAuthSession defines whether support iSCSI Session CHAP authentication",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"fs_type": {
														Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"initiator_name": {
														Description:         "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
														MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"iqn": {
														Description:         "iqn is the target iSCSI Qualified Name.",
														MarkdownDescription: "iqn is the target iSCSI Qualified Name.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"iscsi_interface": {
														Description:         "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
														MarkdownDescription: "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"lun": {
														Description:         "lun represents iSCSI Target Lun number.",
														MarkdownDescription: "lun represents iSCSI Target Lun number.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"portals": {
														Description:         "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
														MarkdownDescription: "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
														MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
														MarkdownDescription: "secretRef is the CHAP Secret for iSCSI target and initiator authentication",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"target_portal": {
														Description:         "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
														MarkdownDescription: "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

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

											"name": {
												Description:         "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"nfs": {
												Description:         "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
												MarkdownDescription: "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"path": {
														Description:         "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"server": {
														Description:         "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

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

											"persistent_volume_claim": {
												Description:         "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
												MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"claim_name": {
														Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
														MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
														MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",

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

											"photon_persistent_disk": {
												Description:         "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
												MarkdownDescription: "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pd_id": {
														Description:         "pdID is the ID that identifies Photon Controller persistent disk",
														MarkdownDescription: "pdID is the ID that identifies Photon Controller persistent disk",

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

											"portworx_volume": {
												Description:         "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
												MarkdownDescription: "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_id": {
														Description:         "volumeID uniquely identifies a Portworx volume",
														MarkdownDescription: "volumeID uniquely identifies a Portworx volume",

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

											"projected": {
												Description:         "projected items for all in one resources secrets, configmaps, and downward API",
												MarkdownDescription: "projected items for all in one resources secrets, configmaps, and downward API",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sources": {
														Description:         "sources is the list of volume projections",
														MarkdownDescription: "sources is the list of volume projections",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "configMap information about the configMap data to project",
																MarkdownDescription: "configMap information about the configMap data to project",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"items": {
																		Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "key is the key to project.",
																				MarkdownDescription: "key is the key to project.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"mode": {
																				Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																				MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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

																	"name": {
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "optional specify whether the ConfigMap or its keys must be defined",
																		MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",

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

															"downward_api": {
																Description:         "downwardAPI information about the downwardAPI data to project",
																MarkdownDescription: "downwardAPI information about the downwardAPI data to project",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"items": {
																		Description:         "Items is a list of DownwardAPIVolume file",
																		MarkdownDescription: "Items is a list of DownwardAPIVolume file",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"field_ref": {
																				Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																				MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"api_version": {
																						Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"field_path": {
																						Description:         "Path of the field to select in the specified API version.",
																						MarkdownDescription: "Path of the field to select in the specified API version.",

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

																			"mode": {
																				Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																				MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"resource_field_ref": {
																				Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"container_name": {
																						Description:         "Container name: required for volumes, optional for env vars",
																						MarkdownDescription: "Container name: required for volumes, optional for env vars",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"divisor": {
																						Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																						MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"resource": {
																						Description:         "Required: resource to select",
																						MarkdownDescription: "Required: resource to select",

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

															"secret": {
																Description:         "secret information about the secret data to project",
																MarkdownDescription: "secret information about the secret data to project",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"items": {
																		Description:         "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "key is the key to project.",
																				MarkdownDescription: "key is the key to project.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"mode": {
																				Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																				MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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

																	"name": {
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "optional field specify whether the Secret or its key must be defined",
																		MarkdownDescription: "optional field specify whether the Secret or its key must be defined",

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

															"service_account_token": {
																Description:         "serviceAccountToken is information about the serviceAccountToken data to project",
																MarkdownDescription: "serviceAccountToken is information about the serviceAccountToken data to project",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"audience": {
																		Description:         "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																		MarkdownDescription: "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"expiration_seconds": {
																		Description:         "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																		MarkdownDescription: "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "path is the path relative to the mount point of the file to project the token into.",
																		MarkdownDescription: "path is the path relative to the mount point of the file to project the token into.",

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

											"quobyte": {
												Description:         "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
												MarkdownDescription: "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"group": {
														Description:         "group to map volume access to Default is no group",
														MarkdownDescription: "group to map volume access to Default is no group",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
														MarkdownDescription: "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"registry": {
														Description:         "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
														MarkdownDescription: "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"tenant": {
														Description:         "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
														MarkdownDescription: "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"user": {
														Description:         "user to map volume access to Defaults to serivceaccount user",
														MarkdownDescription: "user to map volume access to Defaults to serivceaccount user",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume": {
														Description:         "volume is a string that references an already created Quobyte volume by name.",
														MarkdownDescription: "volume is a string that references an already created Quobyte volume by name.",

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

											"rbd": {
												Description:         "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
												MarkdownDescription: "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image": {
														Description:         "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"keyring": {
														Description:         "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"monitors": {
														Description:         "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"pool": {
														Description:         "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"user": {
														Description:         "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

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

											"scale_io": {
												Description:         "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
												MarkdownDescription: "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
														MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"gateway": {
														Description:         "gateway is the host address of the ScaleIO API Gateway.",
														MarkdownDescription: "gateway is the host address of the ScaleIO API Gateway.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"protection_domain": {
														Description:         "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
														MarkdownDescription: "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
														MarkdownDescription: "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"ssl_enabled": {
														Description:         "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
														MarkdownDescription: "sslEnabled Flag enable/disable SSL communication with Gateway, default false",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"storage_mode": {
														Description:         "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
														MarkdownDescription: "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"storage_pool": {
														Description:         "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
														MarkdownDescription: "storagePool is the ScaleIO Storage Pool associated with the protection domain.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"system": {
														Description:         "system is the name of the storage system as configured in ScaleIO.",
														MarkdownDescription: "system is the name of the storage system as configured in ScaleIO.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"volume_name": {
														Description:         "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
														MarkdownDescription: "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",

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

											"secret": {
												Description:         "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
												MarkdownDescription: "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"items": {
														Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
														MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the key to project.",
																MarkdownDescription: "key is the key to project.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mode": {
																Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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

													"optional": {
														Description:         "optional field specify whether the Secret or its keys must be defined",
														MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_name": {
														Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
														MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",

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

											"storageos": {
												Description:         "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
												MarkdownDescription: "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
														MarkdownDescription: "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"volume_name": {
														Description:         "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
														MarkdownDescription: "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_namespace": {
														Description:         "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
														MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",

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

											"vsphere_volume": {
												Description:         "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
												MarkdownDescription: "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"storage_policy_id": {
														Description:         "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
														MarkdownDescription: "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"storage_policy_name": {
														Description:         "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
														MarkdownDescription: "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_path": {
														Description:         "volumePath is the path that identifies vSphere volume vmdk",
														MarkdownDescription: "volumePath is the path that identifies vSphere volume vmdk",

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

							"template_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"time_chaos": {
								Description:         "TimeChaosSpec defines the desired state of TimeChaos",
								MarkdownDescription: "TimeChaosSpec defines the desired state of TimeChaos",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"clock_ids": {
										Description:         "ClockIds defines all affected clock id All available options are ['CLOCK_REALTIME','CLOCK_MONOTONIC','CLOCK_PROCESS_CPUTIME_ID','CLOCK_THREAD_CPUTIME_ID', 'CLOCK_MONOTONIC_RAW','CLOCK_REALTIME_COARSE','CLOCK_MONOTONIC_COARSE','CLOCK_BOOTTIME','CLOCK_REALTIME_ALARM', 'CLOCK_BOOTTIME_ALARM'] Default value is ['CLOCK_REALTIME']",
										MarkdownDescription: "ClockIds defines all affected clock id All available options are ['CLOCK_REALTIME','CLOCK_MONOTONIC','CLOCK_PROCESS_CPUTIME_ID','CLOCK_THREAD_CPUTIME_ID', 'CLOCK_MONOTONIC_RAW','CLOCK_REALTIME_COARSE','CLOCK_MONOTONIC_COARSE','CLOCK_BOOTTIME','CLOCK_REALTIME_ALARM', 'CLOCK_BOOTTIME_ALARM'] Default value is ['CLOCK_REALTIME']",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"container_names": {
										Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
										MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"duration": {
										Description:         "Duration represents the duration of the chaos action",
										MarkdownDescription: "Duration represents the duration of the chaos action",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mode": {
										Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
										MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
										},
									},

									"selector": {
										Description:         "Selector is used to select pods that are used to inject chaos action.",
										MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expression_selectors": {
												Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
												MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"label_selectors": {
												Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
												MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespaces": {
												Description:         "Namespaces is a set of namespace to which objects belong.",
												MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selectors": {
												Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
												MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nodes": {
												Description:         "Nodes is a set of node name and objects must belong to these nodes.",
												MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_phase_selectors": {
												Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
												MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pods": {
												Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
												MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

												Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"time_offset": {
										Description:         "TimeOffset defines the delta time of injected program. It's a possibly signed sequence of decimal numbers, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
										MarkdownDescription: "TimeOffset defines the delta time of injected program. It's a possibly signed sequence of decimal numbers, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
										MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

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

						Required: true,
						Optional: false,
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

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_chaos_mesh_org_workflow_v1alpha1")

	var state ChaosMeshOrgWorkflowV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgWorkflowV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("Workflow")

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

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_workflow_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_chaos_mesh_org_workflow_v1alpha1")

	var state ChaosMeshOrgWorkflowV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgWorkflowV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("Workflow")

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

func (r *ChaosMeshOrgWorkflowV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_chaos_mesh_org_workflow_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
