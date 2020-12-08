terraform {
  required_version = ">= 0.13"
  required_providers {
    null = "~> 2.1"
  }
}

provider "azurerm" {
  version                    = "~> 2.16"
  skip_provider_registration = "true"
  features {}
}