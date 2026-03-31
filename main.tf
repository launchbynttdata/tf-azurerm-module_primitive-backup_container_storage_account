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

resource "azurerm_backup_container_storage_account" "backup_container_storage_account" {
  resource_group_name = var.resource_group_name
  recovery_vault_name = var.recovery_vault_name
  storage_account_id  = var.storage_account_id

  timeouts {
    create = var.create_timeout
    read   = var.read_timeout
    delete = var.delete_timeout
  }
}
