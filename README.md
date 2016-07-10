# go-ufc

Golang interface for UFC API v3.

## Installation

    go get github.com/sungwoncho/go-ufc

## Usage

`ufc.New()` returns a new API client.

API client has the following methods

* Fighters()
* Fighter(id int)
* Events()
* Event(id int)

*example*

```go
import "github.com/sungwoncho/go-ufc"

func main() {
  client := ufc.New()

  res, err := client.Fighters()

  if err != nil {
    // do something
  }
}
```

## License

MIT
