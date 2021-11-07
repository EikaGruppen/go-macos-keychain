# Go macOS keychain

An intutive golang API for the macOS keychain.

Create a better user experience by storing tokens etc. in the macOS keychain, protected by the user's system password.

The only dependency is [keybase/go-keychain](https://github.com/keybase/go-keychain), which is used to communicate to the macOS keychain API directly, in stead for through the `security`-binary, which makes the keys available to all processes that can run that binary.

When fetching a key, the user will be asked if the app should be granted permission to retrieve the key: ["Allow", "Deny", "Always allow]

If the user chooses "Always allow", the apps will be added in the access-list for that key. NOTE: A new verion of the app will ultimatly make the user allow it again, as the checksum of the app has changed.
More details [here](https://developer.apple.com/documentation/security/keychain_services/access_control_lists)

## Usage

```go
keys := keychain.NewKeychainClient("my-app")

name, value := "id_for_pass", "the pass"

err := keys.Update(name, value)
if err != nil {
	if err == keychain.KeyNotFoundError {
		// do something
	} else {
	//handle err
	}
}

fromKeychain, err := keys.Get(name)
//handle err

err = keys.Delete(name)
//handle err
```

## Non-goals

- Creating custom keychains
  - As Apple states:
> ...unnecessary for anything other than an app trying to replicate the keychain access utility. [source](https://developer.apple.com/documentation/security/keychain_services/keychains)