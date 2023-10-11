
terraform {
  backend "gcs" {
    bucket = var.state_bucket
    prefix = "terraform/state/${var.gh_repo_id}"
  }
}
