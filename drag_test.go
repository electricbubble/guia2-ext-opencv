package guia2_ext_opencv

import (
	"github.com/electricbubble/guia2"
	"testing"
)

func TestDriverExt_Drag(t *testing.T) {
	driver, err := guia2.NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	driverExt, err := Extend(driver, 0.95)
	if err != nil {
		t.Fatal(err)
	}

	err = driverExt.Drag("/Users/hero/Desktop/未命名.png", 200, 500)
	if err != nil {
		t.Fatal(err)
	}

	driverExt.Debug(DmNotMatch)

	err = driverExt.OnlyOnceThreshold(0.79).DragOffsetFloat("/Users/hero/Desktop/未命名.png", 1258.3, 238.8, 0.5, 0.85)
	if err != nil {
		t.Fatal(err)
	}
}
