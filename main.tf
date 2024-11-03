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
    tags = {
      "env" = "dev"
    }
  }
}

output "arn" {
  value = acme_bucket.bucket.status.arn
}
