output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_listener_v3alpha1.minimal.yaml
  }
}
