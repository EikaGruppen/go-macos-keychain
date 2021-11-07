package keychain_test

import (
	"testing"

	"github.com/EikaGruppen/go-macos-keychain"
)

func TestKeyItems(t *testing.T) {
	keys := keychain.NewKeychainClient("my-app")

	name, value := "id_for_pass", "the pass"

	err := keys.Update(name, value)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { 
		err = keys.Delete(name)
		if err != nil {
			t.Fatal(err)
		}
	}()

	fromKeychain, err := keys.Get(name)
	if err != nil {
		t.Fatal(err)
	}
	if fromKeychain != value {
		t.Errorf("Expected '%s' from keychain, got '%s'", fromKeychain, value)
	}

	newValue := "the new value"

	err = keys.Update(name, newValue)
	if err != nil {
		t.Fatal(err)
	}

	updatedfromKeychain, err := keys.Get(name)
	if err != nil {
		t.Fatal(err)
	}
	if updatedfromKeychain != newValue {
		t.Errorf("Expected '%s' from keychain, got '%s'", updatedfromKeychain, value)
	}

}
