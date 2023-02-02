package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(compress.New())
	rg := app.Group("/api/v1/diagrams")
	ruler, _ := textmeasure.NewRuler()
	defaultLayout := func(ctx context.Context, g *d2graph.Graph) error {
		return d2dagrelayout.Layout(ctx, g, nil)
	}
	rg.Post("", func(c *fiber.Ctx) error {
		req := map[string]string{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		diagram, _, _ := d2lib.Compile(context.Background(), req["body"], &d2lib.CompileOptions{
			Layout:  defaultLayout,
			Ruler:   ruler,
			ThemeID: d2themescatalog.GrapeSoda.ID,
		})
		svg, _ := d2svg.Render(diagram, &d2svg.RenderOpts{
			Pad: d2svg.DEFAULT_PADDING,
		})
		c.Set("Content-Type", "image/svg+xml")
		return c.SendString(string(svg))
	})
	app.Listen(":3000")

	// diagram, _, _ := d2lib.Compile(context.Background(), "x -> y", &d2lib.CompileOptions{
	// 	Layout:  defaultLayout,
	// 	Ruler:   ruler,
	// 	ThemeID: d2themescatalog.GrapeSoda.ID,
	// })
	// a
	// out, _ := d2svg.Render(diagram, &d2svg.RenderOpts{
	// 	Pad: d2svg.DEFAULT_PADDING,
	// })
	// _ = ioutil.WriteFile(filepath.Join("out.svg"), out, 0600)
}
