# soracom-go

[Soracom](https://dev.soracom.io/jp/) client for Go

## Install

```
> go install github.com/futoase/soracom-go
```

## Usage

- CLI

```
export SORACOM_EMAIL=$your_email
export SORACOM_PASSWORD=$your_password
> soracom-go auth
```

- If you want using implement on your code.

```go
package main

import(
    "log"
    soracomAuth "github.com/futoase/soracom-go/libs/auth"
)

func main() {
    r := soracomAuth.Request{}
    r.Email = "your-email@example.com"
    r.Password = "your-password"
    _, raw, err := r.Auth()
    if err != nil {
      log.Fatal(err)
      return
    }
    println(string(raw))
    return
}
```

# Getting help with command

```
> soracom-go -h
NAME:
   soracom - soracom cli tool for go

USAGE:
   soracom [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   auth                 authorize
   password-reset       password reset
   operator             operator
   help, h              Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h           show help
   --version, -v        print the version

> soracom-go operator -h

NAME:
   soracom operator - operator

USAGE:
   soracom operator command [command options] [arguments...]

COMMANDS:
   token                Update token of operator
   password             Update password
   support_token        get of token for support
   invitation           invitation new user
   verify               verify of invitation user
   info                 get information of operator
   help, h              Shows a list of commands or help for one command

OPTIONS:
   --help, -h   show help

> soracom-go operator token -h

NAME:
   operator token - Update token of operator

USAGE:
   command operator token [command options] [arguments...]

OPTIONS:
   --api-key, -a "api-key"                      api key for SORACOM [$SORACOM_API_KEY]
   --token, -t "token"                          token for SORACOM [$SORACOM_TOKEN]
   --token-timeout-seconds, --timeout "86400"   Timeout seconds of token (default 86400)
   --operator-id, -i "operator-id"              Set operator id
```

# Supported api list

- auth
- operators

# Not implemented api list

- subscriber
- stats
- group
- event handler

# Author

(c) 2015 Keiji Matsuzaki
