package starportconf

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	confyml := `
accounts:
  - name: me
    coins: ["1000token", "100000000stake"]
  - name: you
    coins: ["5000token"]
validator:
  name: user1
  staked: "100000000stake"
`

	conf, err := Parse(strings.NewReader(confyml))
	require.NoError(t, err)
	require.Equal(t, Config{
		Accounts: []Account{
			{
				Name:  "me",
				Coins: []string{"1000token", "100000000stake"},
			},
			{
				Name:  "you",
				Coins: []string{"5000token"},
			},
		},
		Validator: Validator{
			Name:   "user1",
			Staked: "100000000stake",
		},
	}, conf)
}

func TestParseInvalid(t *testing.T) {
	confyml := `
accounts:
  - name: me
    coins: ["1000token", "100000000stake"]
  - name: you
    coins: ["5000token"]
`

	_, err := Parse(strings.NewReader(confyml))
	require.Equal(t, &ValidationError{"validator is required"}, err)
}
