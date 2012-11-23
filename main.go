package main

import (
	"bytes"
	"errors"
	"fmt"
  gl "github.com/chsc/gogl/gl21"
  "github.com/jteeuwen/glfw"
  "image"
	"image/png"
	"io"
	"os"
  "time"
)

const (
	Title  = "star"
	Width  = 640/3
	Height = 480/3
)

var (
  exit = false
  last_time time.Time
)

func main() {
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.Terminate()

	//glfw.OpenWindowHint(glfw.WindowNoResize, 1)
	glfw.OpenWindowHint(glfw.WindowNoResize, gl.TRUE)
	//glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 2)

	if err := glfw.OpenWindow(Width, Height, 0, 0, 0, 0, 32, 32, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(Title)

	if err := gl.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "gl: %s\n", err)
	}

	if err := initScene(); err != nil {
		fmt.Fprintf(os.Stderr, "init: %s\n", err)
		return
	}
	defer destroyScene()

  last_time = time.Now()
	for glfw.WindowParam(glfw.Opened) == 1 && !exit {
    updateScene()
		drawScene()
		glfw.SwapBuffers()
    since := time.Since(last_time).Seconds()
    if (since > 0.02) {
      fmt.Println("frame under 50fps:", since)
    }
    last_time = time.Now()
  }
}

func createTexture(r io.Reader) (textureId gl.Uint, err error) {
	img, err := png.Decode(r)
	if err != nil {
		return 0, err
	}

	rgbaImg, ok := img.(*image.NRGBA)
	if !ok {
		return 0, errors.New("texture must be an NRGBA image")
	}

	gl.GenTextures(1, &textureId)
	gl.BindTexture(gl.TEXTURE_2D, textureId)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	// flip image: first pixel is lower left corner
	imgWidth, imgHeight := img.Bounds().Dx(), img.Bounds().Dy()
	data := make([]byte, imgWidth * imgHeight * 4)
	lineLen := imgWidth * 4
	dest := len(data)-lineLen
	for src := 0; src < len(rgbaImg.Pix); src+=rgbaImg.Stride {
		copy(data[dest:dest+lineLen], rgbaImg.Pix[src:src+rgbaImg.Stride])
		dest-=lineLen
	}
	gl.TexImage2D(gl.TEXTURE_2D, 0, 4, gl.Sizei(imgWidth), gl.Sizei(imgHeight), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(&data[0]))

	return textureId, nil
}

func createTextureFromBytes(data []byte) (gl.Uint, error) {
	r := bytes.NewBuffer(data)
	return createTexture(r)
}

