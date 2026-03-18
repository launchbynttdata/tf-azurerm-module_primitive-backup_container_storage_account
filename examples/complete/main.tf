// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

module "resource_names" {
  source  = "terraform.registry.launch.nttdata.com/module_library/resource_name/launch"
  version = "~> 2.0"

  for_each = var.resource_names_map

  region                  = join("", split("-", var.location))
  class_env               = var.class_env
  cloud_resource_type     = each.value.name
  instance_env            = var.instance_env
  instance_resource       = var.instance_resource
  maximum_length          = each.value.max_length
  logical_product_family  = var.logical_product_family
  logical_product_service = var.logical_product_service
}

module "resource_group" {
  source  = "terraform.registry.launch.nttdata.com/module_primitive/resource_group/azurerm"
  version = "~> 1.0"

  name     = module.resource_names["resource_group"].standard
  location = var.location

  tags = merge(var.tags, { resource_name = module.resource_names["resource_group"].standard })
}

module "recovery_services_vault" {
  source              = "terraform.registry.launch.nttdata.com/module_primitive/recovery_services_vault/azurerm"
  version             = "~> 1.0"
  depends_on          = [module.resource_group]
  name                = module.resource_names["recovery_vault"].standard
  location            = var.location
  resource_group_name = module.resource_group.name

  sku = "Standard"

  tags = {
    resource_name = module.resource_names["recovery_vault"].standard
  }
}

module "storage_account" {
  source     = "terraform.registry.launch.nttdata.com/module_primitive/storage_account/azurerm"
  version    = "~> 1.0"
  depends_on = [module.resource_group]
  storage_account_name = lower(
    "${substr(replace(module.resource_names["storage_account"].standard, "-", ""), 0, 14)}${module.resource_names["storage_account"].random_integer}"
  )
  resource_group_name = module.resource_group.name
  location            = var.location

  account_tier             = "Standard"
  account_replication_type = "LRS"
  account_kind             = "StorageV2"

  enable_https_traffic_only     = true
  public_network_access_enabled = true

  storage_shares = {
    fileshare1 = {
      name  = "testshare"
      quota = 50
    }
  }

  tags = merge(var.tags, { resource_name = module.resource_names["storage_account"].standard })
}

module "backup_container_storage_account" {
  source = "../../"

  resource_group_name = module.resource_group.name
  recovery_vault_name = module.recovery_services_vault.vault_name
  storage_account_id  = module.storage_account.id
}
