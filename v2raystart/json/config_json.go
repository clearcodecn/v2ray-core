package json

//go:generate go run github.com/v2fly/v2ray-core/v4/common/errors/errorgen

import (
	"io"
	"os"

	core "github.com/clearcodecn/v2ray-core"
	"github.com/clearcodecn/v2ray-core/common"
	"github.com/clearcodecn/v2ray-core/common/cmdarg"
	"github.com/clearcodecn/v2ray-core/v2raystart/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "JSON",
		Extension: []string{"json"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				r, err := confloader.LoadExtConfig(v, os.Stdin)
				if err != nil {
					return nil, newError("failed to execute v2ctl to convert config file.").Base(err).AtWarning()
				}
				return core.LoadConfig("protobuf", "", r)
			case io.Reader:
				r, err := confloader.LoadExtConfig([]string{"stdin:"}, os.Stdin)
				if err != nil {
					return nil, newError("failed to execute v2ctl to convert config file.").Base(err).AtWarning()
				}
				return core.LoadConfig("protobuf", "", r)
			default:
				return nil, newError("unknown type")
			}
		},
	}))
}
