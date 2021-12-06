# Go macOS keychain

An intutive golang API for the macOS keychain.

Create a better user experience by storing tokens etc. in the macOS keychain, protected by the user's system password.

The only dependency is [keybase/go-keychain](https://github.com/keybase/go-keychain), which is used to communicate to the macOS keychain API directly, in stead for through the `security`-binary, which makes the keys available to all processes that can run that binary.

When fetching a key, the user will be asked if the app should be granted permission to retrieve the key: ["Allow", "Deny", "Always allow]

If the user chooses "Always allow", the app will be added in the access-list for that key. NOTE: A new version of the app will ultimatly prompt the user to allow it again, as the checksum of the app has changed.
More details [here](https://developer.apple.com/documentation/security/keychain_services/access_control_lists)

## Get

```bash
go get github.com/EikaGruppen/go-macos-keychain@v0.1.0
```

## Usage

```go
keys := keychain.NewKeychainClient("my-app")

name, value := "id_for_pass", "the pass"

err := keys.Update(name, value)
if err != nil {
	//handle err
}

fromKeychain, err := keys.Get(name)
if err != nil {
	if _, ok := err.(*keychain.KeyNotFoundError); ok {
		//handle
	} else if _, ok := err.(*keychain.UserAbortedPromptError); ok {
		//handle
	} else {
		//handle
	}
}

err = keys.Delete(name)
//handle err
```

## Prompt

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃                                                                ┃
┃     _my app_ wants to access _my item_ in OS keychain          ┃
┃   =====================================================        ┃
┃                                                                ┃
┃                                                                ┃
┃        | Always allow |      | Deny |       | Allow |          ┃
┃                                                                ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

The user will be prompted with something like this. If the user chooses `Always allow`, they will not be prompted until the checksum of the app changes (new build). 

## Non-goals

- Creating custom keychains
  - As Apple states:
> ...unnecessary for anything other than an app trying to replicate the keychain access utility. [source](https://developer.apple.com/documentation/security/keychain_services/keychains)
