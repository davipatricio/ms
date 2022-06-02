# ms

Use this package to easily convert various time formats to milliseconds.

## Installation

```bash
go get github.com/davipatricio/ms
```

## Examples

```go
package main

import (
    "fmt"

    "github.com/davipatricio/ms"
)

func main() {
    // If a invalid string is passed, 0 is returned

    // Integers
    fmt.Println(ms.ConvertInt("1m")) // Returns: 60000
    fmt.Println(ms.ConvertInt("1 m")) // Returns: 60000
    fmt.Println(ms.ConvertInt("1s")) // Returns: 1000
    fmt.Println(ms.ConvertInt("-5h")) // Returns: -18000000
    fmt.Println(ms.ConvertInt("1.5s")) // Returns: 0 (use ms.ConvertFloat instead)
    fmt.Println(ms.ConvertInt("2w"))
    fmt.Println(ms.ConvertInt("5y"))

    // Floats
    fmt.Println(ms.ConvertFloat("2.5h")) // Returns: 9000000
    fmt.Println(ms.ConvertFloat("1s")) // Returns: 1000
    fmt.Println(ms.ConvertFloat("-60s")) // Returns: 60000
}
```

## Related

[ms](https://github.com/vercel/ms) - Tiny millisecond conversion utility
