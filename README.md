# go-updater: Build self-updating Go programs

Package updater provides functionality to implement secure, self-updating Go programs (or other single-file targets)
A program can update itself by replacing its executable file with a new version.

It provides the flexibility to implement different updating user experiences
like auto-updating, or manual user-initiated updates. It also boasts
advanced features like binary patching and code signing verification.

## Examples

### Basic Update from URL

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

### Update from Zip File

```go
import (
    "github.com/lucasdecamargo/go-updater"
)

func updateFromZip(zipPath string) error {
    opts := updater.Options{
        TargetPath: "/path/to/installation/directory", // Optional: defaults to executable directory
    }
    
    return updater.ApplyZip("file.zip", opts)
}
```

### Update with Verification

```go
import (
    "crypto/sha256"
    "github.com/lucasdecamargo/go-updater"
)

func secureUpdate(updateData []byte, checksum []byte, publicKey crypto.PublicKey, signature []byte) error {
    opts := updater.Options{
        Checksum:  checksum,
        PublicKey: publicKey,
        Signature: signature,
        Hash:      sha256.New(),
    }
    
    return updater.Apply(bytes.NewReader(updateData), opts)
}
```

### Advanced Usage: Binary Patching

```go
import (
    "github.com/lucasdecamargo/go-updater"
    "github.com/lucasdecamargo/go-updater/internal/binarydist"
)

func patchUpdate(patchData []byte) error {
    opts := updater.Options{
        Patcher: &binarydist.Patcher{}, // Use binary patching
    }
    
    return updater.Apply(bytes.NewReader(patchData), opts)
}
```

## Features

- Cross platform support (Windows too!)
- Binary patch application
- Checksum verification
- Code signing verification
- Support for updating arbitrary files
- **NEW**: Zip file update support with `ApplyZip` function
- **NEW**: Smart target file existence handling (creates files if they don't exist)
- **NEW**: Automatic directory creation for target files

## New Features Documentation

### Zip File Updates

The `ApplyZip` function allows you to update your application from a zip file containing multiple files. This is particularly useful for applications that consist of multiple executables or need to update several files at once.

**Key Features:**
- Automatically extracts and applies all files from the zip archive
- Ignores common system files (`.DS_Store`, `__MACOSX`)
- Skips directories automatically
- Uses the same verification and security features as the standard `Apply` function

**Behavior:**
- If `TargetPath` is not specified, files are extracted to the directory containing the current executable
- Each file in the zip is processed individually using the same `Apply` logic
- Directory structure within the zip is preserved

### Smart Target File Handling

The `Apply` function now intelligently handles target file existence:

**For Non-Existent Files:**
- Creates the target file directly with the new content
- Automatically creates parent directories if they don't exist
- Sets appropriate file permissions (default: 0755)
- Ensures data is synced to disk

**For Existing Files:**
- Performs the standard atomic update procedure
- Creates temporary files (`.target.new`, `.target.old`)
- Handles rollback on failure
- Maintains file permissions and attributes

**Patch Application:**
- Patches can only be applied to existing files
- Returns an error if attempting to patch a non-existent file
- This ensures data integrity and prevents corruption

### Enhanced Error Handling

The library provides robust error handling and rollback mechanisms:

**Rollback Support:**
- Automatic rollback on failed updates
- `RollbackError()` function to check rollback status
- Graceful handling of filesystem inconsistencies
- Cross-platform compatibility (Windows file hiding for old executables)

**Error Types:**
- Clear error messages for different failure scenarios
- Distinction between update failures and rollback failures
- Detailed error context for debugging

## License
Apache
