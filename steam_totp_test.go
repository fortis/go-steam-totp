package steam_totp

import (
	"github.com/stretchr/testify/require"

	"reflect"
	"testing"
	"time"
)

func TestGenerateAuthCode(t *testing.T) {
	k, err := GenerateAuthCode("cnOgv/KdpLoP6Nbh0GMkXkPXALQ=", time.Now())
	require.NoError(t, err, "Failed to generate auth code with valid shared secret")
	require.Equal(t, "string", reflect.TypeOf(k).String(), "Expected string code type")
}
