package types

type ConfirmationData struct {
    FirstName    string
    LastName     string
    Email        string
    Product      string
    Price        float64  // Ensure price is float64
    Quantity     int
    Subtotal     float64
    TotalWithTax float64
    FinalTotal   float64
    TaxRate      float64
}