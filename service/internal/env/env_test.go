package env

import (
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// #nosec
const testSecret = "tha5hei2Hee5le8n"

var allEnvVars = []string{
	"ARGO_CLOUDOPS_ADMIN_SECRET",
	"VAULT_ROLE",
	"VAULT_SECRET",
	"VAULT_ADDR",
	"ARGO_ADDR",
	"ARGO_CLOUDOPS_WORKFLOW_EXECUTION_NAMESPACE",
	"ARGO_CLOUDOPS_CONFIG",
	"ARGO_CLOUDOPS_GIT_AUTH_METHOD",
	"SSH_PEM_FILE",
	"ARGO_CLOUDOPS_GIT_HTTPS_USER",
	"ARGO_CLOUDOPS_GIT_HTTPS_PASS",
	"ARGO_CLOUDOPS_LOG_LEVEL",
	"ARGO_CLOUDOPS_PORT",
}

func setup() {
	for _, envVar := range allEnvVars {
		os.Unsetenv(envVar)
	}
	instance = Vars{}
	once = sync.Once{}
}

func TestGetEnv(t *testing.T) {
	// Given
	setup()
	os.Setenv("ARGO_CLOUDOPS_ADMIN_SECRET", testSecret)
	os.Setenv("VAULT_ROLE", "vaultRole")
	os.Setenv("VAULT_SECRET", testSecret)
	os.Setenv("VAULT_ADDR", "1.2.3.4")
	os.Setenv("ARGO_ADDR", "2.3.4.5")
	os.Setenv("ARGO_CLOUDOPS_WORKFLOW_EXECUTION_NAMESPACE", "argo-ns")
	os.Setenv("ARGO_CLOUDOPS_CONFIG", "/app/test/config/path")
	os.Setenv("SSH_PEM_FILE", "/app/test/ssh.pem")
	os.Setenv("ARGO_CLOUDOPS_GIT_AUTH_METHOD", "https")
	os.Setenv("ARGO_CLOUDOPS_GIT_HTTPS_USER", "testuser")
	os.Setenv("ARGO_CLOUDOPS_GIT_HTTPS_PASS", "testpass")
	os.Setenv("ARGO_CLOUDOPS_DB_HOST", "localhost")
	os.Setenv("ARGO_CLOUDOPS_DB_NAME", "argocloudops")
	os.Setenv("ARGO_CLOUDOPS_DB_USER", "argoco")
	os.Setenv("ARGO_CLOUDOPS_DB_PASSWORD", "1234")
	os.Setenv("ARGO_CLOUDOPS_LOG_LEVEL", "DEBUG")
	os.Setenv("ARGO_CLOUDOPS_PORT", "1234")

	// When
	var env, _ = GetEnv()

	// Then
	assert.Equal(t, env.AdminSecret, testSecret)
	assert.Equal(t, env.VaultRole, "vaultRole")
	assert.Equal(t, env.VaultSecret, testSecret)
	assert.Equal(t, env.VaultAddress, "1.2.3.4")
	assert.Equal(t, env.ArgoNamespace, "argo-ns")
	assert.Equal(t, env.ConfigFilePath, "/app/test/config/path")
	assert.Equal(t, env.SSHPEMFile, "/app/test/ssh.pem")
	assert.Equal(t, env.GitAuthMethod, "https")
	assert.Equal(t, env.GitHTTPSUser, "testuser")
	assert.Equal(t, env.GitHTTPSPass, "testpass")
	assert.Equal(t, env.LogLevel, "DEBUG")
	assert.Equal(t, env.Port, 1234)
	assert.Equal(t, env.DBHost, "localhost")
	assert.Equal(t, env.DBName, "argocloudops")
	assert.Equal(t, env.DBUser, "argoco")
	assert.Equal(t, env.DBPassword, "1234")
}

func TestDefaults(t *testing.T) {
	// Given
	setup()
	os.Setenv("ARGO_CLOUDOPS_ADMIN_SECRET", testSecret)
	os.Setenv("VAULT_ROLE", "vaultRole")
	os.Setenv("VAULT_SECRET", testSecret)
	os.Setenv("VAULT_ADDR", "1.2.3.4")
	os.Setenv("ARGO_ADDR", "2.3.4.5")
	os.Setenv("ARGO_CLOUDOPS_GIT_AUTH_METHOD", "https")

	// When
	var env, _ = GetEnv()

	// Then
	assert.Equal(t, env.ArgoNamespace, "argo")
	assert.Equal(t, env.ConfigFilePath, "argo-cloudops.yaml")
	assert.Equal(t, env.Port, 8443)
}

func TestValidations(t *testing.T) {
	// Given
	setup()
	os.Setenv("ARGO_CLOUDOPS_ADMIN_SECRET", "PW1234")
	os.Setenv("VAULT_ROLE", "vaultRole")
	os.Setenv("VAULT_SECRET", testSecret)
	os.Setenv("VAULT_ADDR", "1.2.3.4")
	os.Setenv("ARGO_ADDR", "2.3.4.5")
	os.Setenv("ARGO_CLOUDOPS_GIT_AUTH_METHOD", "https")

	// When
	_, err := GetEnv()

	// Then
	assert.Error(t, err)
}

func TestRequiredVars(t *testing.T) {
	// Given
	setup()
	os.Setenv("VAULT_ROLE", "vaultRole")
	os.Setenv("VAULT_SECRET", testSecret)
	os.Setenv("VAULT_ADDR", "1.2.3.4")
	os.Setenv("ARGO_ADDR", "2.3.4.5")
	os.Setenv("ARGO_CLOUDOPS_NAMESPACE", "argo-ns")
	os.Setenv("ARGO_CLOUDOPS_CONFIG", "/app/test/config/path")
	os.Setenv("ARGO_CLOUDOPS_GIT_AUTH_METHOD", "https")
	os.Setenv("ARGO_CLOUDOPS_LOG_LEVEL", "DEBUG")
	os.Setenv("ARGO_CLOUDOPS_PORT", "1234")

	// When
	_, err := GetEnv()

	// Then
	assert.Error(t, err)
}
