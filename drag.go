package guia2_ext_opencv

func (dExt *DriverExt) Drag(pathname string, toX, toY int, steps ...int) (err error) {
	return dExt.DragFloat(pathname, float64(toX), float64(toY), steps...)
}

func (dExt *DriverExt) DragFloat(pathname string, toX, toY float64, steps ...int) (err error) {
	return dExt.DragOffsetFloat(pathname, toX, toY, 0.5, 0.5, steps...)
}

func (dExt *DriverExt) DragOffset(pathname string, toX, toY int, xOffset, yOffset float64, steps ...int) (err error) {
	return dExt.DragOffsetFloat(pathname, float64(toX), float64(toY), xOffset, yOffset, steps...)
}

func (dExt *DriverExt) DragOffsetFloat(pathname string, toX, toY, xOffset, yOffset float64, steps ...int) (err error) {
	if len(steps) == 0 {
		steps = []int{12}
	}

	var x, y, width, height float64
	if x, y, width, height, err = dExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	fromX := x + width*xOffset
	fromY := y + height*yOffset

	return dExt.d.DragFloat(fromX, fromY, toX, toY, steps[0])
}
