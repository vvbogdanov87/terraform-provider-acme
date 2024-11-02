terraform {
  required_providers {
    acme = {
      source = "registry.terraform.io/vvbogdanov87/acme"
    }
  }
}

provider "acme" {
  namespace = "default"
}

resource "acme_bucket" "bucket" {
  name = "acme-test-bucket-123"

  spec = {
      region = "us-west-2"
  }
}

output "arn" {
  value = acme_bucket.bucket.status.arn
}
