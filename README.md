# soracom-go

Soracom client for Go

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
