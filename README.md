# Semver

[![release github](https://img.shields.io/github/v/release/usvc/go-semver?sort=semver)](https://github.com/usvc/go-semver)

[![pipeline status](https://gitlab.com/usvc/modules/go/semver/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/semver/-/commits/master)
[![Build Status](https://travis-ci.org/usvc/go-semver.svg?branch=master)](https://travis-ci.org/usvc/go-semver)

A Go package to deal with semantic versions as defined at [https://semver.org](https://semver.org).

- [Semver](#semver)
  - [Usage](#usage)
    - [Importing](#importing)
    - [Parsing a semantic version string](#parsing-a-semantic-version-string)
    - [Retrieving semantic version as a string](#retrieving-semantic-version-as-a-string)
    - [Sorting semantic versions](#sorting-semantic-versions)
  - [Development Runbook](#development-runbook)
    - [Getting Started](#getting-started)
    - [Continuous Integration (CI) Pipeline](#continuous-integration-ci-pipeline)
      - [On Github](#on-github)
        - [Releasing](#releasing)
      - [On Gitlab](#on-gitlab)
        - [Version Bumping](#version-bumping)
        - [DockerHub Publishing](#dockerhub-publishing)
  - [Licensing](#licensing)

## Usage

### Importing

```go
import "github.com/usvc/go-semver"
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

#### On Github

Github is used to deploy binaries/libraries because of it's ease of access by other developers.

##### Releasing

Releasing of the binaries can be done via Travis CI.

1. On Github, navigate to the [tokens settings page](https://github.com/settings/tokens) (by clicking on your profile picture, selecting **Settings**, selecting **Developer settings** on the left navigation menu, then **Personal Access Tokens** again on the left navigation menu)
2. Click on **Generate new token**, give the token an appropriate name and check the checkbox on **`public_repo`** within the **repo** header
3. Copy the generated token
4. Navigate to [travis-ci.org](https://travis-ci.org) and access the cooresponding repository there. Click on the **More options** button on the top right of the repository page and select **Settings**
5. Scroll down to the section on **Environment Variables** and enter in a new **NAME** with `RELEASE_TOKEN` and the **VALUE** field cooresponding to the generated personal access token, and hit **Add**

#### On Gitlab

Gitlab is used to run tests and ensure that builds run correctly.

##### Version Bumping

1. Run `make .ssh`
2. Copy the contents of the file generated at `./.ssh/id_rsa.base64` into an environment variable named **`DEPLOY_KEY`** in **Settings > CI/CD > Variables**
3. Navigate to the **Deploy Keys** section of the **Settings > Repository > Deploy Keys** and paste in the contents of the file generated at `./.ssh/id_rsa.pub` with the **Write access allowed** checkbox enabled

- **`DEPLOY_KEY`**: generate this by running `make .ssh` and copying the contents of the file generated at `./.ssh/id_rsa.base64`

##### DockerHub Publishing

1. Login to [https://hub.docker.com](https://hub.docker.com), or if you're using your own private one, log into yours
2. Navigate to [your security settings at the `/settings/security` endpoint](https://hub.docker.com/settings/security)
3. Click on **Create Access Token**, type in a name for the new token, and click on **Create**
4. Copy the generated token that will be displayed on the screen
5. Enter the following varialbes into the CI/CD Variables page at **Settings > CI/CD > Variables** in your Gitlab repository:

- **`DOCKER_REGISTRY_URL`**: The hostname of the Docker registry (defaults to `docker.io` if not specified)
- **`DOCKER_REGISTRY_USERNAME`**: The username you used to login to the Docker registry
- **`DOCKER_REGISTRY_PASSWORD`**: The generated access token

## Licensing

Code in this package is licensed under the [MIT license (click to see full text))](./LICENSE)