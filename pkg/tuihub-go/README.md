# TuiHub Go

TuiHub Client & Plugin helper for Go developers.

## Installation

```bash
go get github.com/tuihub/librarian/pkg/tuihub-go
```

## Getting Started

**Client**

```go
package main

import (
	"context"
	"fmt"
	"github.com/tuihub/librarian/pkg/tuihub-go"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	"os"
)

func main() {
	ctx := context.Background()
	c, err := tuihub.LoginByPassword(ctx, "username", "password")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	information, err := c.GetServerInformation(ctx, &pb.GetServerInformationRequest{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(information)
}
```

**Plugin**

```go
package main

import (
	"context"
	"fmt"
	"github.com/tuihub/librarian/pkg/tuihub-go"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	"os"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// version is the version of the compiled software.
	version string
)

// impl tuihub.Handler
type Handler struct{}

func main() {
	ctx := context.Background()
	plugin, err := tuihub.NewPorter(
		ctx,
		tuihub.PorterConfig{
			Name:           "plugin-name",
			Version:        version,
			GlobalName:     "YOUR_PROJECT_URL",
			FeatureSummary: &porter.PorterFeatureSummary{},
			Server:         tuihub.ServerConfig{},
		},
		Handler{},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = plugin.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

```