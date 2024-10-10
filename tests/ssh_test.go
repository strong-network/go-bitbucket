package tests

import (
	"os"
	"testing"

	"github.com/ktrysmt/go-bitbucket"
)

// func getClient(t *testing.T) *bitbucket.Client {
// 	user := os.Getenv("BITBUCKET_TEST_USERNAME")
// 	pass := os.Getenv("BITBUCKET_TEST_PASSWORD")

// 	if user == "" {
// 		t.Error("BITBUCKET_TEST_USERNAME is empty.")
// 	}
// 	if pass == "" {
// 		t.Error("BITBUCKET_TEST_PASSWORD is empty.")
// 	}

// 	return bitbucket.NewBasicAuth(user, pass)
// }

func TestUserSSHKey(t *testing.T) {
	user := os.Getenv("BITBUCKET_TEST_USERNAME")
	// pass := os.Getenv("BITBUCKET_TEST_PASSWORD")
	// owner := os.Getenv("BITBUCKET_TEST_OWNER")
	// repo := os.Getenv("BITBUCKET_TEST_REPOSLUG")

	if user == "" {
		t.Error("BITBUCKET_TEST_USERNAME is empty.")
	}
	c := getClient(t)
	var sshKeyResourceUuid string

	label := "go-user-test"
	key := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAK/b1cHHDr/TEV1JGQl+WjCwStKG6Bhrv0rFpEsYlyTBm1fzN0VOJJYn4ZOPCPJwqse6fGbXntEs+BbXiptR+++HycVgl65TMR0b5ul5AgwrVdZdT7qjCOCgaSV74/9xlHDK8oqgGnfA7ZoBBU+qpVyaloSjBdJfLtPY/xqj4yHnXKYzrtn/uFc4Kp9Tb7PUg9Io3qohSTGJGVHnsVblq/rToJG7L5xIo0OxK0SJSQ5vuId93ZuFZrCNMXj8JDHZeSEtjJzpRCBEXHxpOPhAcbm4MzULgkFHhAVgp4JbkrT99/wpvZ7r9AdkTg7HGqL3rlaDrEcWfL7Lu6TnhBdq5"
	t.Run("create", func(t *testing.T) {
		keyOptions := &bitbucket.SSHKeyOptions{
			Label: label,
			Key:   key,
			Owner: user,
		}
		sshUserKey, err := c.Users.SSHKeys.Create(keyOptions)
		if err != nil {
			t.Error(err)
		}
		if sshUserKey == nil {
			t.Error("The User SSH Key could not be created.")
		}
		sshKeyResourceUuid = sshUserKey.Uuid
	})
	t.Run("get", func(t *testing.T) {
		keyOptions := &bitbucket.SSHKeyOptions{
			Owner: user,
			Uuid:  sshKeyResourceUuid,
		}
		sshKey, err := c.Users.SSHKeys.Get(keyOptions)
		if err != nil {
			t.Error(err)
		}
		if sshKey == nil {
			t.Error("The Deploy Key could not be retrieved.")
		}

		if sshKey.Uuid != sshKeyResourceUuid {
			t.Error("The SSH Key `id` attribute does not match the expected value.")
		}
		if sshKey.Label != label {
			t.Error("The SSH Key `label` attribute does not match the expected value.")
		}
		if sshKey.Key != key {
			t.Error("The SSH Key `key` attribute does not match the expected value.")
		}
	})

	t.Run("delete", func(t *testing.T) {
		keyOptions := &bitbucket.SSHKeyOptions{
			Owner: user,
			Uuid:  sshKeyResourceUuid,
		}
		_, err := c.Users.SSHKeys.Delete(keyOptions)
		if err != nil {
			t.Error(err)
		}
	})
}

// func getUsername(t *testing.T) string {
// 	ev := os.Getenv("BITBUCKET_TEST_USERNAME")
// 	if ev != "" {
// 		return ev
// 	}

// 	return "example-username"
// }

// func getPassword() string {
// 	ev := os.Getenv("BITBUCKET_TEST_PASSWORD")
// 	if ev != "" {
// 		return ev
// 	}

// 	return "password"
// }
