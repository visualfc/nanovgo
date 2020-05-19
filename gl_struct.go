package nanovgo

import (
	"github.com/visualfc/gl"
)

const (
	nsvgShaderFILLGRAD = iota
	nsvgShaderFILLIMG
	nsvgShaderSIMPLE
	nsvgShaderIMG
)

type glnvgCallType int

const (
	glnvgNONE glnvgCallType = iota
	glnvgFILL
	glnvgCONVEXFILL
	glnvgSTROKE
	glnvgTRIANGLES
	glnvgTRIANGLESTRIP
)

type glCall struct {
	callType       glnvgCallType
	image          int
	pathOffset     int
	pathCount      int
	triangleOffset int
	triangleCount  int
	uniformOffset  int
	blendFunc      glnvgBlend
}

type glPath struct {
	fillOffset   int
	fillCount    int
	strokeOffset int
	strokeCount  int
}

type glFragUniforms [44]float32

func (u *glFragUniforms) reset() {
	for i := 0; i < 44; i++ {
		u[i] = 0
	}
}

func (u *glFragUniforms) setScissorMat(mat []float32) {
	copy(u[0:12], mat[0:12])
}

func (u *glFragUniforms) clearScissorMat() {
	for i := 0; i < 12; i++ {
		u[i] = 0
	}
}

func (u *glFragUniforms) setPaintMat(mat []float32) {
	copy(u[12:24], mat[0:12])
}

func (u *glFragUniforms) setInnerColor(color Color) {
	copy(u[24:28], color.List())
}

func (u *glFragUniforms) setOuterColor(color Color) {
	copy(u[28:32], color.List())
}

func (u *glFragUniforms) setScissorExt(a, b float32) {
	u[32] = a
	u[33] = b
}

func (u *glFragUniforms) setScissorScale(a, b float32) {
	u[34] = a
	u[35] = b
}

func (u *glFragUniforms) setExtent(ext [2]float32) {
	copy(u[36:38], ext[:])
}

func (u *glFragUniforms) setRadius(radius float32) {
	u[38] = radius
}

func (u *glFragUniforms) setFeather(feather float32) {
	u[39] = feather
}

func (u *glFragUniforms) setStrokeMult(strokeMult float32) {
	u[40] = strokeMult
}

func (u *glFragUniforms) setStrokeThr(strokeThr float32) {
	u[41] = strokeThr
}

func (u *glFragUniforms) setTexType(texType float32) {
	u[42] = texType
}

func (u *glFragUniforms) setType(typeCode float32) {
	u[43] = typeCode
}

type glTexture struct {
	id            int
	tex           gl.Texture
	width, height int
	texType       nvgTextureType
	flags         ImageFlags
}

type glnvgBlend struct {
	srcRGB   gl.Enum
	dstRGB   gl.Enum
	srcAlpha gl.Enum
	dstAlpha gl.Enum
}

func glnvg_convertBlendFuncFactor(factor BlendFactor) gl.Enum {
	switch factor {
	case Zero:
		return gl.ZERO
	case One:
		return gl.ONE
	case SrcColor:
		return gl.SRC_COLOR
	case OneMinusSrcColor:
		return gl.ONE_MINUS_SRC_COLOR
	case DstColor:
		return gl.DST_COLOR
	case OneMinusDstColor:
		return gl.ONE_MINUS_DST_COLOR
	case SrcAlpha:
		return gl.SRC_ALPHA
	case OneMinusSrcAlpha:
		return gl.ONE_MINUS_SRC_ALPHA
	case DstAlpha:
		return gl.DST_ALPHA
	case OneMinusDstAlpha:
		return gl.ONE_MINUS_DST_ALPHA
	case SrcAlphaSaturate:
		return gl.SRC_ALPHA_SATURATE
	}
	return gl.INVALID_ENUM
}

func glnvg__blendCompositeOperation(op *nvgCompositeOperationState) glnvgBlend {
	var blend glnvgBlend
	blend.srcRGB = glnvg_convertBlendFuncFactor(op.srcRGB)
	blend.dstRGB = glnvg_convertBlendFuncFactor(op.dstRGB)
	blend.srcAlpha = glnvg_convertBlendFuncFactor(op.srcAlpha)
	blend.dstAlpha = glnvg_convertBlendFuncFactor(op.dstAlpha)
	if blend.srcRGB == gl.INVALID_ENUM || blend.dstRGB == gl.INVALID_ENUM || blend.srcAlpha == gl.INVALID_ENUM || blend.dstAlpha == gl.INVALID_ENUM {
		blend.srcRGB = gl.ONE
		blend.dstRGB = gl.ONE_MINUS_SRC_ALPHA
		blend.srcAlpha = gl.ONE
		blend.dstAlpha = gl.ONE_MINUS_SRC_ALPHA
	}
	return blend
}
