// credit to https://github.com/poi5305/go-yuv2webRTC/blob/master/webrtc/webrtc.go
package util

import (
	"image"
	"unsafe"

	"github.com/giongto35/cloud-game/config"
)

// https://stackoverflow.com/questions/9465815/rgb-to-yuv420-algorithm-efficiency

/*
void rgba2yuv(void *destination, void *source, int width, int height, int stride) {
	const int image_size = width * height;
	unsigned char *rgba = source;
  unsigned char *dst_y = destination;
  unsigned char *dst_u = destination + image_size;
  unsigned char *dst_v = destination + image_size + image_size/4;

	// Y plane
	for( int y=0; y<height; ++y ) {
    for( int x=0; x<width; ++x ) {
      const int i = y*(width+stride) + x;
			*dst_y++ = ( ( 66*rgba[4*i] + 129*rgba[4*i+1] + 25*rgba[4*i+2] ) >> 8 ) + 16;
		}
  }
  // U plane
  for( int y=0; y<height; y+=2 ) {
    for( int x=0; x<width; x+=2 ) {
      const int i = y*(width+stride) + x;
			*dst_u++ = ( ( -38*rgba[4*i] + -74*rgba[4*i+1] + 112*rgba[4*i+2] ) >> 8 ) + 128;
		}
  }
  // V plane
  for( int y=0; y<height; y+=2 ) {
    for( int x=0; x<width; x+=2 ) {
      const int i = y*(width+stride) + x;
			*dst_v++ = ( ( 112*rgba[4*i] + -94*rgba[4*i+1] + -18*rgba[4*i+2] ) >> 8 ) + 128;
		}
  }
}
*/
import "C"

// RgbaToYuv convert to yuv from rgba
func RgbaToYuv(rgba *image.RGBA) []byte {
	w := rgba.Rect.Max.X
	h := rgba.Rect.Max.Y
	size := int(float32(w*h) * 1.5)
	stride := rgba.Stride - w*4
	yuv := make([]byte, size, size)
	C.rgba2yuv(unsafe.Pointer(&yuv[0]), unsafe.Pointer(&rgba.Pix[0]), C.int(w), C.int(h), C.int(stride))
	return yuv
}

// RgbaToYuvInplace convert to yuv from rgba inplace to yuv. Avoid reallocation
func RgbaToYuvInplace(rgba *image.RGBA, yuv []byte) {
	stride := rgba.Stride - config.Width*4
	C.rgba2yuv(unsafe.Pointer(&yuv[0]), unsafe.Pointer(&rgba.Pix[0]), C.int(config.Width), C.int(config.Height), C.int(stride))
}
