# HLTB Crashdummy GO

Go implementation of the "Crashdummy HowLongToBeat API": https://codeberg.org/Crashdummy/HowLongToBeatApi

This client works with the public instance at https://hltbapi.codepotatoes.de.

#### Install:

```bash
go get -u github.com/ShadowDash2000/hltb-crashdummy-go
```

#### How to use:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ShadowDash2000/hltb-crashdummy-go"
)

func main() {
    client := hltb.New(
        hltb.WithTimeout(30),
        hltb.WithRetryCount(3),
        // hltb.WithBaseUrl("https://hltbapi.codepotatoes.de"),
    )

    game, err := client.GetByHltbId(context.Background(), 10)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(game.Title, game.MainStory)
}
```
