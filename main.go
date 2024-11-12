package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	etag "github.com/pablor21/echo-etag/v4"

	"go-store/templates"
	"go-store/types"
)

const taxRate = 0.07 // Hardcoded tax rate

// Product prices
var productPrices = map[string]float64{
	"drive shaft": 250.0,
	"piston rod": 60.0,
	"exhaust header": 200.0,
}


func main() {
	e := echo.New()
	e.Use(etag.Etag())

	// Serve static assets (CSS, images, etc.)
	e.Static("assets", "./assets")

	// Define routes
	e.GET("/store", func(ctx echo.Context) error {
		return Render(ctx, http.StatusOK, templates.Base(templates.Store()))
	})
	e.POST("/purchase", func(ctx echo.Context) error {
		// Get form values from POST request
		firstName := ctx.FormValue("first-name")
		lastName := ctx.FormValue("last-name")
		email := ctx.FormValue("email")
		product := ctx.FormValue("product")
		quantityStr := ctx.FormValue("quantity")
		donation := ctx.FormValue("donation")
	
		// Convert quantity to integer
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			return fmt.Errorf("invalid quantity: %v", err)
		}
	
		// Get product price
		price, ok := productPrices[product]
		if !ok {
			return fmt.Errorf("invalid product selected")
		}
	
		// Calculate subtotal
		subtotal := float64(quantity) * price
	
		// Calculate total with tax
		totalWithTax := subtotal * (1 + taxRate)
	
		// Calculate final total (round up if donation is selected)
		finalTotal := totalWithTax
		if donation == "yes" {
			finalTotal = math.Ceil(finalTotal)
		}
	
		// Format data for rendering
		data := types.ConfirmationData{
			FirstName:    firstName,
			LastName:     lastName,
			Email:        email,
			Product:      product,
			Price:        price,  // Ensure price is float64
			Quantity:     quantity,
			Subtotal:     subtotal,
			TotalWithTax: totalWithTax,
			FinalTotal:   finalTotal,
			TaxRate:      taxRate,
		}
	
		// Render confirmation page using Templ
		err = Render(ctx, http.StatusOK, templates.Confirmation(data))
		if err != nil {
			ctx.Logger().Error(err)
			return err
		}
		return nil
		})

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}


// Custom Render method for Templ templates
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
