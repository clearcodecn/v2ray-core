package reverse_test

import (
	"testing"

	"github.com/clearcodecn/v2ray-core/app/reverse"
	"github.com/clearcodecn/v2ray-core/common"
)

func TestStaticPickerEmpty(t *testing.T) {
	picker, err := reverse.NewStaticMuxPicker()
	common.Must(err)
	worker, err := picker.PickAvailable()
	if err == nil {
		t.Error("expected error, but nil")
	}
	if worker != nil {
		t.Error("expected nil worker, but not nil")
	}
}
