# go-lingrad

Linear Gradient drawing in Go programming language

## Usage

The small utility here has been written to test the gradient code.
It draws a gradient and save it to a png file. Its usage is quite simple:

```
-width int
    set image width (default 50)
-height int
    set image height (default 512)
-landscape
    set landscape mode
-inverse
    inverse order
-o string
    set file output (default "out.png")
```

## Example

![example](https://raw.githubusercontent.com/xigh/go-lingrad/master/out.png)

## How to program it ?

```golang
gr := New()
gr.Append(color.RGBA{R: 0, G: 0, B: 0, A: 255})
gr.Append(color.RGBA{R: 56, G: 15, B: 109, A: 255})
gr.Append(color.RGBA{R: 182, G: 54, B: 121, A: 255})
gr.Append(color.RGBA{R: 253, G: 154, B: 105, A: 255})
gr.Append(color.RGBA{R: 252, G: 246, B: 184, A: 255})

c1 := gr.ColorAt(0)
c2 := gr.ColorAt(1./3)
c3 := gr.ColorAt(2./3)
c4 := gr.ColorAt(1)
```

## License

BSD 2-Clause License
