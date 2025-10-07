# go-updater: Build self-updating Go programs

Package updater provides functionality to implement secure, self-updating Go programs (or other single-file targets)
A program can update itself by replacing its executable file with a new version.

It provides the flexibility to implement different updating user experiences
like auto-updating, or manual user-initiated updates. It also boasts
advanced features like binary patching and code signing verification.

Example of updating from a URL:

```go
import (
    "fmt"
    "net/http"

    "github.com/lucasdecamargo/go-updater"
)

func doUpdate(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    err := updater.Apply(resp.Body, updater.Options{})
    if err != nil {
        // error handling
    }
    return err
}
```

## Features

- Cross platform support (Windows too!)
- Binary patch application
- Checksum verification
- Code signing verification
- Support for updating arbitrary files

## License
Apache
