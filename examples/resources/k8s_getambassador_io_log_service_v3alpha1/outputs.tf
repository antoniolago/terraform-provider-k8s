output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_log_service_v3alpha1.minimal.yaml
  }
}
