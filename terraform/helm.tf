resource "helm_release" "releases" {
  for_each = {
    for release in var.helm_releases : release.name => release
  }

  name       = each.value.name
  namespace  = each.value.namespace
  chart      = each.value.chart
  repository = each.value.repository
  version    = each.value.version

  create_namespace = true

  values = [
    file("${path.module}/${each.value.values}"),
  ]

  depends_on = [
    minikube_cluster.cluster,
    kubernetes_config_map.dashboards
  ]
}
