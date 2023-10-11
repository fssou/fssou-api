provider "google" {
  project = var.project_name
  region  = "us-central1"
}

resource "google_cloud_run_v2_service" "main" {
  name     = "api-francl-in"
  location = "us-central1"
  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}

resource "google_cloud_run_domain_mapping" "main" {
  depends_on = [ google_cloud_run_v2_service.main ]
  name     = "api.francl.in"
  location = "us-central1"
  metadata {
    namespace = var.project_id
  }
  spec {
    route_name = google_cloud_run_service.main.name
  }
}
