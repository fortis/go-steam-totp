# Go Steam TOTP
This package generates Steam-style 5-digit alphanumeric two-factor authentication codes given a shared secret.

## Usage
##### Generate 5-digit code to Log on Steam

```go
	code, _ := steam_totp.GenerateAuthCode("cnOgv/KdpLoP6Nbh0GMkXkPXALQ=", time.Now()))
```
