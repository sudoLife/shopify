Shopify reviews crawler
======
***Shopify reviews crawler*** fetches all reviews of an app from [Shopify app store](https://apps.shopify.com/)

### Installation

``` shell
$ go get -u github.com/sudoLife/shopify
```

### Example

``` go
import (
	"encoding/json"
	"os"
	"github.com/sudoLife/shopify"
)

func main() {
	reviews := shopify.Parse("https://apps.shopify.com/YourApp/reviews")
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	
	enc.Encode(reviews)
}
```

### Third party libraries
* [Colly](https://github.com/gocolly/colly/)
