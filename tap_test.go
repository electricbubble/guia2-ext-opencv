package guia2_ext_opencv

import (
	"github.com/electricbubble/guia2"
	"testing"
)

var uiaServerURL = "http://localhost:6790/wd/hub"

func TestDriverExt_Tap(t *testing.T) {
	driver, err := guia2.NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	// screenshot, err := driver.Screenshot()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// err = ioutil.WriteFile("/Users/hero/Desktop/d1.png", screenshot.Bytes(), 0600)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	driverExt, err := Extend(driver, 0.95)
	if err != nil {
		t.Fatal(err)
	}

	_ = driverExt
	err = driverExt.Tap("/Users/hero/Desktop/未命名.png")
	if err != nil {
		t.Fatal(err)
	}
}
