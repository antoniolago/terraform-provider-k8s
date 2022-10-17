---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_k8gb_absa_oss_gslb_v1beta1 Resource - terraform-provider-k8s"
subcategory: "k8gb.absa.oss/v1beta1"
description: |-
  Gslb is the Schema for the gslbs API
---

# k8s_k8gb_absa_oss_gslb_v1beta1 (Resource)

Gslb is the Schema for the gslbs API

## Example Usage

```terraform
resource "k8s_k8gb_absa_oss_gslb_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_k8gb_absa_oss_gslb_v1beta1" "example" {
  metadata = {
    name      = "test-gslb-failover"
    namespace = "test-gslb"
  }
  spec = {
    ingress = {
      rules = [
        {
          host = "failover.test.k8gb.io"
          http = {
            paths = [
              {
                path      = "/"
                path_type = "Prefix"
                backend = {
                  service = {
                    name = "frontend-podinfo"
                    port = {
                      name = "http"
                    }
                  }
                }
              }
            ]
          }
        }
      ]
    }
    strategy = {
      primary_geo_tag = "eu-west-1"
      type            = "failover"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) GslbSpec defines the desired state of Gslb (see [below for nested schema](#nestedatt--spec))

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

- `ingress` (Attributes) Gslb-enabled Ingress Spec (see [below for nested schema](#nestedatt--spec--ingress))
- `strategy` (Attributes) Gslb Strategy spec (see [below for nested schema](#nestedatt--spec--strategy))

<a id="nestedatt--spec--ingress"></a>
### Nested Schema for `spec.ingress`

Optional:

- `backend` (Attributes) A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default. (see [below for nested schema](#nestedatt--spec--ingress--backend))
- `ingress_class_name` (String) IngressClassName is the name of the IngressClass cluster resource. The associated IngressClass defines which controller will implement the resource. This replaces the deprecated 'kubernetes.io/ingress.class' annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this field. For more information, refer to the IngressClass documentation.
- `rules` (Attributes List) A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend. (see [below for nested schema](#nestedatt--spec--ingress--rules))
- `tls` (Attributes List) TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI. (see [below for nested schema](#nestedatt--spec--ingress--tls))

<a id="nestedatt--spec--ingress--backend"></a>
### Nested Schema for `spec.ingress.backend`

Optional:

- `resource` (Attributes) Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with 'Service'. (see [below for nested schema](#nestedatt--spec--ingress--backend--resource))
- `service` (Attributes) Service references a Service as a Backend. This is a mutually exclusive setting with 'Resource'. (see [below for nested schema](#nestedatt--spec--ingress--backend--service))

<a id="nestedatt--spec--ingress--backend--resource"></a>
### Nested Schema for `spec.ingress.backend.service`

Required:

- `kind` (String) Kind is the type of resource being referenced
- `name` (String) Name is the name of resource being referenced

Optional:

- `api_group` (String) APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.


<a id="nestedatt--spec--ingress--backend--service"></a>
### Nested Schema for `spec.ingress.backend.service`

Required:

- `name` (String) Name is the referenced service. The service must exist in the same namespace as the Ingress object.

Optional:

- `port` (Attributes) Port of the referenced service. A port name or port number is required for a IngressServiceBackend. (see [below for nested schema](#nestedatt--spec--ingress--backend--service--port))

<a id="nestedatt--spec--ingress--backend--service--port"></a>
### Nested Schema for `spec.ingress.backend.service.port`

Optional:

- `name` (String) Name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.
- `number` (Number) Number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.




<a id="nestedatt--spec--ingress--rules"></a>
### Nested Schema for `spec.ingress.rules`

Required:

- `http` (Attributes) HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'. (see [below for nested schema](#nestedatt--spec--ingress--rules--http))

Optional:

- `host` (String) Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the 'host' part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the IP in the Spec of the parent Ingress. 2. The ':' delimiter is not respected because ports are not allowed. Currently the port of an Ingress is implicitly :80 for http and :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.  Host can be 'precise' which is a domain name without the terminating dot of a network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain name prefixed with a single wildcard label (e.g. '*.foo.com'). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*'). Requests will be matched against the Host field in the following way: 1. If Host is precise, the request matches this rule if the http host header is equal to Host. 2. If Host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.

<a id="nestedatt--spec--ingress--rules--http"></a>
### Nested Schema for `spec.ingress.rules.host`

Required:

- `paths` (Attributes List) A collection of paths that map requests to backends. (see [below for nested schema](#nestedatt--spec--ingress--rules--host--paths))

<a id="nestedatt--spec--ingress--rules--host--paths"></a>
### Nested Schema for `spec.ingress.rules.host.paths`

Required:

- `backend` (Attributes) Backend defines the referenced service endpoint to which the traffic will be forwarded to. (see [below for nested schema](#nestedatt--spec--ingress--rules--host--paths--backend))
- `path_type` (String) PathType determines the interpretation of the Path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is done on a path element by element basis. A path element refers is the list of labels in the path split by the '/' separator. A request is a match for path p if every p is an element-wise prefix of p of the request path. Note that if the last element of the path is a substring of the last element in request path, it is not a match (e.g. /foo/bar matches /foo/bar/baz, but does not match /foo/barbaz). * ImplementationSpecific: Interpretation of the Path matching is up to the IngressClass. Implementations can treat this as a separate PathType or treat it identically to Prefix or Exact path types. Implementations are required to support all path types.

Optional:

- `path` (String) Path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value 'Exact' or 'Prefix'.

<a id="nestedatt--spec--ingress--rules--host--paths--backend"></a>
### Nested Schema for `spec.ingress.rules.host.paths.path`

Optional:

- `resource` (Attributes) Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with 'Service'. (see [below for nested schema](#nestedatt--spec--ingress--rules--host--paths--path--resource))
- `service` (Attributes) Service references a Service as a Backend. This is a mutually exclusive setting with 'Resource'. (see [below for nested schema](#nestedatt--spec--ingress--rules--host--paths--path--service))

<a id="nestedatt--spec--ingress--rules--host--paths--path--resource"></a>
### Nested Schema for `spec.ingress.rules.host.paths.path.service`

Required:

- `kind` (String) Kind is the type of resource being referenced
- `name` (String) Name is the name of resource being referenced

Optional:

- `api_group` (String) APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.


<a id="nestedatt--spec--ingress--rules--host--paths--path--service"></a>
### Nested Schema for `spec.ingress.rules.host.paths.path.service`

Required:

- `name` (String) Name is the referenced service. The service must exist in the same namespace as the Ingress object.

Optional:

- `port` (Attributes) Port of the referenced service. A port name or port number is required for a IngressServiceBackend. (see [below for nested schema](#nestedatt--spec--ingress--rules--host--paths--path--service--port))

<a id="nestedatt--spec--ingress--rules--host--paths--path--service--port"></a>
### Nested Schema for `spec.ingress.rules.host.paths.path.service.port`

Optional:

- `name` (String) Name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.
- `number` (Number) Number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.







<a id="nestedatt--spec--ingress--tls"></a>
### Nested Schema for `spec.ingress.tls`

Optional:

- `hosts` (List of String) Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
- `secret_name` (String) SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.



<a id="nestedatt--spec--strategy"></a>
### Nested Schema for `spec.strategy`

Required:

- `type` (String) Load balancing strategy type:(roundRobin|failover)

Optional:

- `dns_ttl_seconds` (Number) Defines DNS record TTL in seconds
- `primary_geo_tag` (String) Primary Geo Tag. Valid for failover strategy only
- `split_brain_threshold_seconds` (Number) Split brain TXT record expiration in seconds
- `weight` (Map of String) Weight is defined by map region:weight

