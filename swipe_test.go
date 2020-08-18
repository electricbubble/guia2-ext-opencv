package guia2_ext_opencv

import (
	"github.com/electricbubble/guia2"
	"testing"
)

func TestDriverExt_Swipe(t *testing.T) {
	driver, err := guia2.NewDriver(nil, uiaServerURL)
	if err != nil {
		t.Fatal(err)
	}

	driverExt, err := Extend(driver, 0.95)
	if err != nil {
		t.Fatal(err)
	}

	// err = driverExt.Swipe("/Users/hero/Desktop/未命名.png", 200, 500)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	err = driverExt.SwipeDown("/Users/hero/Desktop/未命名.png", 0.5)
	if err != nil {
		t.Fatal(err)
	}

}
