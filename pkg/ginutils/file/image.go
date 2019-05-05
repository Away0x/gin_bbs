package file

import (
	"github.com/disintegration/imaging"
)

// ReduceImageSize resize 图片
func ReduceImageSize(imgPath string, maxWidth int) error {
	if maxWidth <= 0 {
		return nil
	}

	img, err := imaging.Open(imgPath)
	if err != nil {
		return err
	}

	imgWidth := img.Bounds().Dx()
	if imgWidth <= maxWidth {
		return nil
	}
	// resize
	newImg := imaging.Resize(img, maxWidth, 0, imaging.Lanczos) // 等比 resize
	if err = imaging.Save(newImg, imgPath); err != nil {
		return err
	}

	return nil
}
