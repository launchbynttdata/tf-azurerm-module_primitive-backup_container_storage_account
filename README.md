# tf-azurerm-module_primitive-backup_container_storage_account
<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | ~> 3.117 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) | 3.117.1 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [azurerm_backup_container_storage_account.backup_container_storage_account](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/backup_container_storage_account) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_resource_group_name"></a> [resource\_group\_name](#input\_resource\_group\_name) | Name of the resource group where the Recovery Services Vault exists. | `string` | n/a | yes |
| <a name="input_recovery_vault_name"></a> [recovery\_vault\_name](#input\_recovery\_vault\_name) | Name of the Recovery Services Vault where the storage account will be registered. | `string` | n/a | yes |
| <a name="input_storage_account_id"></a> [storage\_account\_id](#input\_storage\_account\_id) | ID of the Storage Account to register with the Recovery Services Vault. | `string` | n/a | yes |
| <a name="input_create_timeout"></a> [create\_timeout](#input\_create\_timeout) | Timeout for creating the Backup Storage Account Container. | `string` | `"30m"` | no |
| <a name="input_read_timeout"></a> [read\_timeout](#input\_read\_timeout) | Timeout for reading the Backup Storage Account Container. | `string` | `"5m"` | no |
| <a name="input_delete_timeout"></a> [delete\_timeout](#input\_delete\_timeout) | Timeout for deleting the Backup Storage Account Container. | `string` | `"30m"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_backup_container_storage_account_id"></a> [backup\_container\_storage\_account\_id](#output\_backup\_container\_storage\_account\_id) | The ID of the Backup Storage Account Container. |
<!-- END_TF_DOCS -->
