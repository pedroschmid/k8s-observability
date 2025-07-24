resource "kubernetes_config_map" "dashboards" {
  for_each = fileset("${path.module}/helm/grafana/dashboards", "*.json")

  metadata {
    name      = replace(each.key, ".json", "")
    namespace = "monitoring"
    labels = {
      grafana_dashboard = "1"
    }
  }

  data = {
    "${each.key}" = file("${path.module}/helm/grafana/dashboards/${each.key}")
  }
}
