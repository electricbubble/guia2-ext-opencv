package guia2_ext_opencv

import (
	"bytes"
	"github.com/electricbubble/guia2"
	"image"
	"io/ioutil"
	"os"
)
import cvHelper "github.com/electricbubble/opencv-helper"

// TemplateMatchMode is the type of the template matching operation.
type TemplateMatchMode int

const (
	// TmSqdiff maps to TM_SQDIFF
	TmSqdiff TemplateMatchMode = iota
	// TmSqdiffNormed maps to TM_SQDIFF_NORMED
	TmSqdiffNormed
	// TmCcorr maps to TM_CCORR
	TmCcorr
	// TmCcorrNormed maps to TM_CCORR_NORMED
	TmCcorrNormed
	// TmCcoeff maps to TM_CCOEFF
	TmCcoeff
	// TmCcoeffNormed maps to TM_CCOEFF_NORMED
	TmCcoeffNormed
)

type DebugMode int

const (
	// DmOff no output
	DmOff DebugMode = iota
	// DmEachMatch output matched and mismatched values
	DmEachMatch
	// DmNotMatch output only values that do not match
	DmNotMatch
)

type DriverExt struct {
	d *guia2.Driver
	// scale     float64
	MatchMode TemplateMatchMode
	Threshold float64
}

// Extend 获得扩展后的 Driver，
// 并指定匹配阀值，
// 默认匹配模式为 TmCcoeffNormed，
// 默认关闭 OpenCV 匹配值计算后的输出
func Extend(driver *guia2.Driver, threshold float64, matchMode ...TemplateMatchMode) (dExt *DriverExt, err error) {
	dExt = &DriverExt{d: driver}

	// if dExt.scale, err = dExt.d.DeviceScaleRatio(); err != nil {
	// 	return nil, err
	// }

	if len(matchMode) == 0 {
		matchMode = []TemplateMatchMode{TmCcoeffNormed}
	}
	dExt.MatchMode = matchMode[0]
	cvHelper.Debug(cvHelper.DebugMode(DmOff))
	dExt.Threshold = threshold
	return dExt, nil
}

func (dExt *DriverExt) OnlyOnceThreshold(threshold float64) (newExt *DriverExt) {
	newExt = new(DriverExt)
	newExt.d = dExt.d
	// newExt.scale = dExt.scale
	newExt.MatchMode = dExt.MatchMode
	newExt.Threshold = threshold
	return
}

func (dExt *DriverExt) OnlyOnceMatchMode(matchMode TemplateMatchMode) (newExt *DriverExt) {
	newExt = new(DriverExt)
	newExt.d = dExt.d
	// newExt.scale = dExt.scale
	newExt.MatchMode = matchMode
	newExt.Threshold = dExt.Threshold
	return
}

func (dExt *DriverExt) Debug(dm DebugMode) {
	cvHelper.Debug(cvHelper.DebugMode(dm))
}

func getBufFromDisk(name string) (*bytes.Buffer, error) {
	var f *os.File
	var err error
	if f, err = os.Open(name); err != nil {
		return nil, err
	}
	var all []byte
	if all, err = ioutil.ReadAll(f); err != nil {
		return nil, err
	}
	return bytes.NewBuffer(all), nil
}

func (dExt *DriverExt) FindImageRectInUIKit(search string) (x, y, width, height float64, err error) {
	var bufSource, bufSearch *bytes.Buffer
	if bufSearch, err = getBufFromDisk(search); err != nil {
		return 0, 0, 0, 0, err
	}
	if bufSource, err = dExt.d.Screenshot(); err != nil {
		return 0, 0, 0, 0, err
	}

	var rect image.Rectangle
	if rect, err = cvHelper.FindImageRectFromRaw(bufSource, bufSearch, float32(dExt.Threshold), cvHelper.TemplateMatchMode(dExt.MatchMode)); err != nil {
		return 0, 0, 0, 0, err
	}

	x, y, width, height = float64(rect.Min.X), float64(rect.Min.Y), float64(rect.Dx()), float64(rect.Dy())
	return
}
