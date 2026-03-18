package testimpl

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableBackupContainerStorageAccount(t *testing.T, ctx types.TestContext) {
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		t.Fatal("ARM_SUBSCRIPTION_ID environment variable is not set")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Unable to get Azure credentials: %v", err)
	}

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

		client, err := armrecoveryservicesbackup.NewBackupProtectionContainersClient(
			subscriptionID,
			cred,
			nil,
		)
		if err != nil {
			t.Fatalf("Failed to create Backup Protection Containers client: %v", err)
		}

		pager := client.NewListPager(
			recoveryVaultName,
			resourceGroupNameFromResourceID(backupContainerID),
			nil,
		)

		container := findBackupContainerByID(t, pager, backupContainerID)

		assert.Equal(t, backupContainerID, *container.ID, "Backup container resource ID mismatch")
		assert.Contains(
			t,
			*container.Name,
			lastSegment(storageAccountID),
			"Backup container resource name should include the storage account name",
		)
	})
}

func findBackupContainerByID(
	t *testing.T,
	pager interface {
		More() bool
		NextPage(context.Context) (armrecoveryservicesbackup.BackupProtectionContainersClientListResponse, error)
	},
	backupContainerID string,
) *armrecoveryservicesbackup.ProtectionContainerResource {
	t.Helper()

	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			t.Fatalf("Failed to list backup protection containers: %v", err)
		}

		for _, container := range page.Value {
			if container.ID != nil && strings.EqualFold(*container.ID, backupContainerID) {
				return container
			}
		}
	}

	t.Fatalf("Backup container %q was not found in Azure", backupContainerID)
	return nil
}

func resourceGroupNameFromResourceID(resourceID string) string {
	segments := strings.Split(strings.Trim(resourceID, "/"), "/")
	for i := 0; i < len(segments)-1; i++ {
		if strings.EqualFold(segments[i], "resourceGroups") {
			return segments[i+1]
		}
	}

	return ""
}

func lastSegment(resourceID string) string {
	segments := strings.Split(strings.Trim(resourceID, "/"), "/")
	return segments[len(segments)-1]
}
