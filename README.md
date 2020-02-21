# Semver

A package to deal with semantic versions as defined at [https://semver.org](https://semver.org).

## Usage

### Importing

```go
import "github.com/usvc/semver"
```

### Parsing a semantic version string

```go
semver.Parse("1.0.0-alpha.1+202022022637")
```

### Retrieving semantic version as a string

```go
// setting up fictitious semantic version
version := Semver{
  Prefix: "v",
  VersionCore: semver.VersionCore{
    Major: 1,
    Minor: 2,
    Patch: 3,
  },
  PreRelease: semver.PreRelease{
    "alpha", "1",
  },
  BuildMetadata: semver.BuildMetadata{
    "202022022637",
  },
}

// the magic follows
fmt.Println(version.String())
// v1.2.3-alpha.1+202022022637
```

## Licensing

Code in this package is licensed under the [MIT license (click to see full text))](./LICENSE)