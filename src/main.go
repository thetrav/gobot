package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
)

type App struct {
	*app.Application
	scene  *core.Node
	camera *camera.Camera
}

func (a *App) update(rend *renderer.Renderer, deltaTime time.Duration) {

	a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
	rend.Render(a.scene, a.camera)
}

func (a *App) createScene() {
	a.scene = core.NewNode()
	gui.Manager().Set(a.scene)

	a.camera = camera.New(1)
	a.camera.SetPosition(0, 0, 3)
	a.scene.Add(a.camera)
	camera.NewOrbitControl(a.camera)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		a.camera.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	var err error

	err = addModel(a.scene, "../assets/kenney/space_kit2/alien.obj")
	if err != nil {
		panic(err)
	}

	err = addModel(a.scene, "../assets/levels/level1a.obj")
	if err != nil {
		panic(err)
	}

	

	// Create and add lights to the scene
	a.scene.Add(light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.5))
	sunlight := light.NewDirectional(&math32.Color{1.0, 1.0, 1.0}, 0.5)
	sunlight.SetPosition(0, 1, 0)
	sunlight.SetDirection(-0.1,-1, -0.3)
	a.scene.Add(sunlight)

	// pointLight := light.NewDirection(&math32.Color{1, 1, 1}, 5.0)
	// pointLight.SetPosition(1, 1, 2)
	// a.scene.Add(pointLight)

	// Create and add an axis helper to the scene
	a.scene.Add(helper.NewAxes(0.5))

	// Set background color to gray
	a.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)
}

func (a *App) createGui() {
	// Create and add a button to the scene
	btn := gui.NewButton("Make Red")
	btn.SetPosition(100, 40)
	btn.SetSize(40, 40)
	btn.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		// mat.SetColor(math32.NewColor("DarkRed"))
		fmt.Printf("click!")
	})
	a.scene.Add(btn)
}

func main() {
	a := new(App)
	a.Application = app.App()
	a.createScene()
	a.createGui()

	// Run the application
	a.Run(a.update)
}

func addModel(node *core.Node, path string) error {
	model, err := openModel(path)
	if err != nil {
		return err
	}
	node.Add(model)
	return nil
}

// openModel try to open the specified model and add it to the scene
func openModel(fpath string) (*core.Node, error) {

	dir, file := filepath.Split(fpath)
	ext := filepath.Ext(file)
	if ext == ".obj" {
		// Checks for material file in the same dir
		matfile := file[:len(file)-len(ext)]
		matpath := filepath.Join(dir, matfile)
		_, err := os.Stat(matpath)
		if err != nil {
			matpath = ""
		}
		dec, err := obj.Decode(fpath, matpath)
		if err != nil {
			return nil, err
		}

		return dec.NewGroup()

	}
	return nil, fmt.Errorf("Unrecognized model file extension:[%s]", ext)
}
