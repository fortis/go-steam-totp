# Go Steam TOTP
[![Build Status](https://travis-ci.org/fortis/go-steam-totp.svg?branch=master)](https://travis-ci.org/fortis/go-steam-totp)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This package generates Steam-style 5-digit alphanumeric two-factor authentication codes given a shared secret.

## Usage
##### Generate 5-digit code to Log on Steam

```go
	code, _ := steam_totp.GenerateAuthCode("cnOgv/KdpLoP6Nbh0GMkXkPXALQ=", time.Now())
```

## Acknowledgments
- To [DoctorMcKay](https://github.com/DoctorMcKay) for the great nodejs module [node-steam-totp](https://github.com/DoctorMcKay/node-steam-totp)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details
