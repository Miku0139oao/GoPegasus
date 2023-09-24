## GoPegasus
Golang Simple PegasusHK APi Client

## Examples
```go
package main

import (
	"fmt"
	GoPegasus "github.com/Miku0139oao/GoPegasus/Pegasus"
)

func main() {
	ItemList := GoPegasus.GetItems(GoPegasus.CPU) //
	for _, items := range ItemList {
		fmt.Println(items)
	}
}

```

