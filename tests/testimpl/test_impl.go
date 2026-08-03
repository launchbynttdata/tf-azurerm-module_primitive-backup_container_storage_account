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
	validateBackupContainerStorageAccount(t, ctx)
}

func TestComposableReadonlyBackupContainerStorageAccount(t *testing.T, ctx types.TestContext) {
	validateBackupContainerStorageAccount(t, ctx)
}

func validateBackupContainerStorageAccount(t *testing.T, ctx types.TestContext) {
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

		client, err := armrecoveryservicesbackup.NewProtectionContainersClient(
			subscriptionID,
			cred,
			nil,
		)
		if err != nil {
			t.Fatalf("Failed to create Protection Containers client: %v", err)
		}

		fabricName, containerName := backupContainerPathParts(t, backupContainerID)

		container, err := client.Get(
			context.Background(),
			recoveryVaultName,
			resourceGroupNameFromResourceID(backupContainerID),
			fabricName,
			containerName,
			nil,
		)
		if err != nil {
			t.Fatalf("Failed to get backup protection container %q: %v", backupContainerID, err)
		}

		assert.Equal(t, backupContainerID, *container.ID, "Backup container resource ID mismatch")
		assert.Contains(
			t,
			*container.Name,
			lastSegment(storageAccountID),
			"Backup container resource name should include the storage account name",
		)
	})
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

func backupContainerPathParts(t *testing.T, resourceID string) (string, string) {
	t.Helper()

	segments := strings.Split(strings.Trim(resourceID, "/"), "/")
	for i := 0; i < len(segments)-1; i++ {
		if strings.EqualFold(segments[i], "backupFabrics") && i+2 < len(segments) && strings.EqualFold(segments[i+2], "protectionContainers") && i+3 < len(segments) {
			return segments[i+1], segments[i+3]
		}
	}

	t.Fatalf("Failed to parse backup fabric and protection container name from %q", resourceID)
	return "", ""
}

func lastSegment(resourceID string) string {
	segments := strings.Split(strings.Trim(resourceID, "/"), "/")
	return segments[len(segments)-1]
}
