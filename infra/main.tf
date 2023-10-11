provider "google" {
  project = var.project_name
  region  = "us-central1"
}

resource "google_cloud_run_service" "main" {
  name     = "api-francl-in"
  location = "us-central1"
  metadata {
    namespace = var.project_name
  }
  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}

resource "google_cloud_run_domain_mapping" "main" {
  name     = "api.francl.in"
  location = "us-central1"
  metadata {
    namespace = var.project_name
  }
  spec {
    route_name = google_cloud_run_service.main.name
  }
}
