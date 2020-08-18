package guia2_ext_opencv

func (dExt *DriverExt) Tap(pathname string) error {
	return dExt.TapOffset(pathname, 0.5, 0.5)
}

func (dExt *DriverExt) TapOffset(pathname string, xOffset, yOffset float64) (err error) {
	var x, y, width, height float64
	if x, y, width, height, err = dExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	return dExt.d.TapFloat(x+width*xOffset, y+height*yOffset)
}
