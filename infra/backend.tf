
terraform {
  backend "gcs" {
    bucket = "staging.francl-in.appspot.com"
    prefix = "terraform/state/fssou-api"
  }
}
