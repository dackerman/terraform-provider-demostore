# Example terraform file to extract provider schema

terraform {
  required_providers {
    stlstore = {
      source  = "dackerman/stlstore"
      version = "~> 1.0.0"
    }
  }
}

provider "stlstore" {
}