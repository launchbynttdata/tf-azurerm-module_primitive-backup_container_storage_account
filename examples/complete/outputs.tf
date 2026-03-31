output "backup_container_storage_account_id" {
  value = module.backup_container_storage_account.backup_container_storage_account_id
}

output "storage_account_id" {
  value = module.storage_account.id
}

output "recovery_vault_name" {
  value = module.recovery_services_vault.vault_name
}
