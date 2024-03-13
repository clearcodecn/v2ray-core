package outbound_test

import (
	"testing"

	. "github.com/clearcodecn/v2ray/app/proxyman/outbound"
	"github.com/clearcodecn/v2ray/features/outbound"
)

func TestInterfaces(t *testing.T) {
	_ = (outbound.Handler)(new(Handler))
	_ = (outbound.Manager)(new(Manager))
}
