//go:build !confonly
// +build !confonly

package reverse

import (
	"crypto/rand"
	"io"

	"github.com/clearcodecn/v2ray-core/common/dice"
)

func (c *Control) FillInRandom() {
	randomLength := dice.Roll(64)
	c.Random = make([]byte, randomLength)
	io.ReadFull(rand.Reader, c.Random)
}
