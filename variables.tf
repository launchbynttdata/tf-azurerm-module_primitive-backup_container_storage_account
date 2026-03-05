variable "resource_group_name" {
  description = "Name of the resource group where the Recovery Services Vault exists."
  type        = string
}

variable "recovery_vault_name" {
  description = "Name of the Recovery Services Vault where the storage account will be registered."
  type        = string
}

variable "storage_account_id" {
  description = "ID of the Storage Account to register with the Recovery Services Vault."
  type        = string
}

variable "create_timeout" {
  description = "Timeout for creating the Backup Storage Account Container."
  type        = string
  default     = "30m"
}

variable "read_timeout" {
  description = "Timeout for reading the Backup Storage Account Container."
  type        = string
  default     = "5m"
}

variable "delete_timeout" {
  description = "Timeout for deleting the Backup Storage Account Container."
  type        = string
  default     = "30m"
}
