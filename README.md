# Go Steam TOTP
[![Build Status](https://travis-ci.org/fortis/go-steam-totp.svg?branch=master)](https://travis-ci.org/fortis/go-steam-totp)
[![Coverage Status](https://coveralls.io/repos/github/fortis/go-steam-totp/badge.svg?branch=master)](https://coveralls.io/github/fortis/go-steam-totp?branch=master)
[![license](https://img.shields.io/npm/l/steam-totp.svg)](https://github.com/fortis/go-steam-totp/blob/master/LICENSE)

This package generates Steam-style 5-digit alphanumeric two-factor authentication codes given a shared secret.

## Installation

    $ go get github.com/fortis/go-steam-totp

## Usage
Generate 5-digit code to Log on Steam

```go
package main

import (
        "log"
        "github.com/fortis/go-steam-totp"
)

func main() {
    var sharedsecret = "cnOgv/KdpLoP6Nbh0GMkXkPXALQ="
    code, _ := steam_totp.GenerateAuthCode(sharedsecret, time.Now())
    log.Println(code)
}
```

## Acknowledgments
- To [DoctorMcKay](https://github.com/DoctorMcKay) for the great nodejs module [node-steam-totp](https://github.com/DoctorMcKay/node-steam-totp)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details
