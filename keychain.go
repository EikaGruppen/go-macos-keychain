package keychain

import (
	"fmt"

	kc "github.com/keybase/go-keychain"
)

type keychainClient struct {
	name string
}

// Returns a KeychainClient for specified app name
// That struct can be used to fetch, update and delete keys
func NewKeychainClient(name string) *keychainClient {
	return &keychainClient{name}
}

type UserAbortedPromptError struct{}

func (e *UserAbortedPromptError) Error() string {
	return "The user aborted the prompt for password"
}

type KeyNotFoundError struct {
	name string
}

func (e *KeyNotFoundError) Error() string {
	return fmt.Sprintf("Key with id '%s' not found", e.name)
}

// Update a key
// Will create one if it does not exist
func (k *keychainClient) Update(name string, password string) error {
	item := k.buildItem(name)
	item.SetData([]byte(password))

	err := kc.AddItem(item)
	if err == kc.Error(-25299) {
	  existing := k.buildItem(name)
		return kc.UpdateItem(existing, item)
	} else {
		return err
	}
}

// Get a key with the specified name
// It will return KeyNotFoundError when not found
// Also, if user cancels the prompt, UserAbortedPromptError will be returned
func (k *keychainClient) Get(name string) (string, error) {
	query := k.buildItem(name)
	query.SetMatchLimit(kc.MatchLimitOne)
	query.SetReturnData(true)
	results, err := kc.QueryItem(query)
	if err != nil {
		if err == kc.Error(-128) {
			return "", &UserAbortedPromptError{}
		}
		return "", err
	} else if len(results) != 1 {
		// kc.ErrorItemNotFound
		return "", &KeyNotFoundError{name}
	}
	password := string(results[0].Data)

	return password, err
}

// Delete a key with the specified name
func (k *keychainClient) Delete(name string) error {
	item := k.buildItem(name)
	return kc.DeleteItem(item)
}

func (k *keychainClient) buildItem(keyName string) kc.Item {
	item := kc.NewItem()
	item.SetSecClass(kc.SecClassGenericPassword)
	item.SetService(k.name)
	item.SetLabel(keyName)
	return item
}
