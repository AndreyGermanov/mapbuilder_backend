package mapbuilder

const SCALE_10M = "10m"
const SCALE_50M = "50m"
const SCALE_110M = "110m"

type ShapeParams struct {
	left   float64
	top    float64
	bottom float64
	right  float64
	width  float64
	height float64
}

type TransformParams struct {
	offsetX     float64
	offsetY     float64
	scaleX      float64
	scaleY      float64
	scaleWidth  float64
	scaleHeight float64
}

type Polygon [][]float64

func PrepareTransformParams(transform TransformParams, shape ShapeParams) TransformParams {
	if transform.offsetX == 0 {
		transform.offsetX = 0.0 + shape.left
	}
	if transform.offsetY == 0 {
		transform.offsetY = 0.0 + shape.top
	}
	if transform.scaleWidth != 0.0 && transform.scaleHeight == 0.0 {
		transform.scaleHeight = transform.scaleWidth * (shape.height / shape.width)
	}
	if transform.scaleHeight != 0.0 && transform.scaleWidth == 0.0 {
		transform.scaleWidth = transform.scaleHeight * (shape.width / shape.height)
	}

	if transform.scaleWidth != 0 {
		transform.scaleX = transform.scaleWidth / shape.width
	}
	if transform.scaleHeight != 0 {
		transform.scaleY = transform.scaleHeight / shape.height
	}
	if transform.scaleX == 0 {
		transform.scaleX = 1
	}
	if transform.scaleY == 0 {
		transform.scaleY = 1
	}
	return transform
}
