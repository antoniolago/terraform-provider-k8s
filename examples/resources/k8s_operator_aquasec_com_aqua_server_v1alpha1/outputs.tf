output "resources" {
  value = {
    "minimal" = k8s_operator_aquasec_com_aqua_server_v1alpha1.minimal.yaml
    "example" = k8s_operator_aquasec_com_aqua_server_v1alpha1.example.yaml
  }
}
