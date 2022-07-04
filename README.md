Shopify reviews crawler
======
[![Go Report Card](https://goreportcard.com/badge/github.com/sudoLife/shopify)](https://goreportcard.com/report/github.com/sudoLife/shopify)  

<a href="https://www.buymeacoffee.com/sudoLife" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>


***Shopify reviews crawler*** fetches all reviews of an app from [Shopify app store](https://apps.shopify.com/)

### Installation

``` shell
$ go get -u github.com/sudoLife/shopify
```

### Example Code

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
### Usage

After above code has been saved as something like *get-reviews.go*, run the following command to scrape the data in JSON:
``` shell
$ go run get-reviews.go
```

### Third party libraries
* [Colly](https://github.com/gocolly/colly/)
