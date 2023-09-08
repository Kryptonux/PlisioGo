# Plisio-Go

`plisio` is a Go package for wrapping and creating invoices on Plisio.

## Installation

To use `plisio`, you can install it with `go get`:

```bash
go get github.com/Kryptonux/plisio
```

## Example
Here's an example of how to use mongoquery to construct and execute MongoDB queries:
```go
func main() {
	apiKey := ""
	email := ""
	callbackURL := ""

	client := plisio.NewPlisioClient(apiKey, email, callbackURL)

	sourceCurrency := "USD"
	sourceAmount := 15.00
	currency := "BTC"
	orderName := "xkey_payment_" + plisio.GenerateTransactionID(12)
	orderNumber := "xkey_payment_" + plisio.GenerateTransactionID(12)

	invoiceJSON, err := client.CreateInvoice(sourceCurrency, sourceAmount, orderNumber, currency, orderName)
	if err != nil {
		log.Fatal(err)
	}

  // Handle Data
}
```
