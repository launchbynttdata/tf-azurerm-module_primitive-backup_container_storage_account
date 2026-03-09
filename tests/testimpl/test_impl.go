package testimpl

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableBackupContainerStorageAccount(t *testing.T, ctx types.TestContext) {

	t.Run("validateBackupContainerStorageAccountExists", func(t *testing.T) {

		backupContainerID := terraform.Output(
			t,
			ctx.TerratestTerraformOptions(),
			"backup_container_storage_account_id",
		)

		recoveryVaultName := terraform.Output(
			t,
			ctx.TerratestTerraformOptions(),
			"recovery_vault_name",
		)

		storageAccountID := terraform.Output(
			t,
			ctx.TerratestTerraformOptions(),
			"storage_account_id",
		)

		assert.NotEmpty(t, backupContainerID)
		assert.NotEmpty(t, recoveryVaultName)
		assert.NotEmpty(t, storageAccountID)
	})
}
