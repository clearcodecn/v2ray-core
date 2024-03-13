package json

//go:generate errorgen

import (
	"io"

	"github.com/clearcodecn/v2ray"
	"github.com/clearcodecn/v2ray/common"
	"github.com/clearcodecn/v2ray/common/cmdarg"
	"github.com/clearcodecn/v2ray/infra/conf/serial"
	"github.com/clearcodecn/v2ray/main/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "JSON",
		Extension: []string{"json"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				r, err := confloader.LoadExtConfig(v)
				if err != nil {
					return nil, newError("failed to execute v2ctl to convert config file.").Base(err).AtWarning()
				}
				return core.LoadConfig("protobuf", "", r)
			case io.Reader:
				return serial.LoadJSONConfig(v)
			default:
				return nil, newError("unknow type")
			}
		},
	}))
}
