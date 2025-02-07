package stats_test

import (
	"context"
	"testing"

	. "github.com/clearcodecn/v2ray-core/app/stats"
	"github.com/clearcodecn/v2ray-core/common"
	"github.com/clearcodecn/v2ray-core/features/stats"
)

func TestInternface(t *testing.T) {
	_ = (stats.Manager)(new(Manager))
}

func TestStatsCounter(t *testing.T) {
	raw, err := common.CreateObject(context.Background(), &Config{})
	common.Must(err)

	m := raw.(stats.Manager)
	c, err := m.RegisterCounter("test.counter")
	common.Must(err)

	if v := c.Add(1); v != 1 {
		t.Fatal("unpexcted Add(1) return: ", v, ", wanted ", 1)
	}

	if v := c.Set(0); v != 1 {
		t.Fatal("unexpected Set(0) return: ", v, ", wanted ", 1)
	}

	if v := c.Value(); v != 0 {
		t.Fatal("unexpected Value() return: ", v, ", wanted ", 0)
	}
}
