# Discord OAuth 2.0
This is a provider for the package [golang.org/x/oauth2](https://godoc.org/golang.org/x/oauth2) implementing authentication endpoints for [Discord](https://discordapp.com)

## Install
```sh
$ go get github.com/ravener/discord-oauth2
```

## Usage
```go
package main

import (
  "github.com/ravener/discord-oauth2"
  "golang.org/x/oauth2"
)

func main() {
  conf := &oauth2.Config{
    Endpoint: discord.Endpoint,
    Scopes: []string{discord.ScopeIdentify},
    RedirectURL: "http://localhost:3000/auth/callback",
    ClientID: "id",
    ClientSecret: "secret",
  }
  // Use oauth2 package as normal, i.e
  // redirect users to conf.AuthCodeURL("state") for initial auth
  // then inside the callback:
  //  - verify the state param as needed.
  //  - exchange code with conf.Exchange(oauth2.NoContext, code)
  //  - Store in session if necessary, etc.
  // to get like user's info use conf.Client(ctx, token) to get a proper http client
  // for such requests.
}
```
A full authentication flow example server can be found in [example directory](example)

You can join [`#oauth2` in my Discord Server](https://discord.gg/wpE3Nfp) for support and updates.

## License
[MIT](LICENSE)
