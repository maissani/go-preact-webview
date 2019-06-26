package main

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/zserge/webview"
)

// Counter is a simple example of automatic Go-to-JS data binding
type Counter struct {
	Value int `json:"value"`
}

type Color struct {
	Value string `json:"value"`
}

// Add increases the value of a counter by n
func (c *Counter) Add(n int) {
	c.Value = c.Value + int(n)
}

// Reset sets the value of a counter back to zero
func (c *Counter) Reset() {
	c.Value = 0
}

func (c *Counter) Substract(n int) {
	c.Value = c.Value - int(n)
}

func (o *Color) New() {
	o.Value = string("#FF0000")
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (o *Color) Changec() {
	RandomColor, _ := randomHex(3)
	o.Value = string("#" + RandomColor)
}

func main() {
	w := webview.New(webview.Settings{
		Title:     "CalyScope Launcher ",
		Width:     1000,
		Height:    1000,
		Resizable: true,
	})
	defer w.Exit()
	w.Dispatch(func() {
		// Inject controller
		w.Bind("counter", &Counter{})
		w.Bind("color", &Color{})
		// Inject CSS
		w.InjectCSS(string(MustAsset("js/styles.css")))

		// Inject web UI framework and app UI code
		loadUIFramework(w)
	})
	w.Run()
}
