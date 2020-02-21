# Semver

[![pipeline status](https://gitlab.com/usvc/modules/go/semver/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/semver/-/commits/master)


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

### Sorting semantic versions

```go
stringVersions := []string{"1.2.1", "3.1.0", "2.0.0", "1.0.0", "2.0.0-alpha"}
semverVersions := semver.Semvers{}
for i := 0; i < len(stringVersions); i++ {
  semverVersion, _ := semver.Parse(stringVersions[i])
  semverVersions = append(semverVersions, *semverVersion)
}
sort.Sort(semverVersions)
for i := 0; i < len(semverVersions); i++ {
  fmt.Println(semverVersions[i].String())
}
// 1.0.0
// 1.2.1
// 2.0.0-alpha
// 2.0.0
// 3.1.0
```

## Development Runbook

### Getting Started

1. Clone this repository
2. Run `make deps` to pull in external dependencies
3. Write some awesome stuff
4. Run `make test` to ensure unit tests are passing
5. Push

### Continuous Integration (CI) Pipeline

To set up the CI pipeline in Gitlab:

1. Run `make .ssh`
2. Copy the contents of the file generated at `./.ssh/id_rsa.base64` into an environment variable named **`DEPLOY_KEY`** in **Settings > CI/CD > Variables**
3. Navigate to the **Deploy Keys** section of the **Settings > Repository > Deploy Keys** and paste in the contents of the file generated at `./.ssh/id_rsa.pub` with the **Write access allowed** checkbox enabled

- **`DEPLOY_KEY`**: generate this by running `make .ssh` and copying the contents of the file generated at `./.ssh/id_rsa.base64`

## Licensing

Code in this package is licensed under the [MIT license (click to see full text))](./LICENSE)