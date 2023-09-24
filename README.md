Golang Simple PegasusHK APi Client
Request PegasusHk Items Data by tiny codes
 
Example

```go
package main

import (
	"fmt"
	GoPegasus "pegasusClient/Pegasus"
)

func main() {
	ItemList := GoPegasus.GetItems(GoPegasus.CPU) //Categories built-in  ([]int{})
	for _, items := range ItemList {
		fmt.Println(items)
	}
}
```

By the above lines , you can get the ItemList from PegasusHK 
