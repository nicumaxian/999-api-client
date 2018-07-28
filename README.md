# 999.md GO API Client

API Client was based on the [Official documentation](https://999.md/api/documentation) offered.

## Initializing client

Import library
```lang
import "github.com/maxiannicu/999-api-client/client"

...

client := client.BuildClient(key)
```

## Usage

### Fetch categories

```
categories, e := categories.GetCategories(client)
```

### Fetch subcategories

```
subCategories, e := categories.GetSubCategories(client,270)
```

### Fetch offers

```
offers, err := categories.GetOfferTypes(client, "270","1404")
```

