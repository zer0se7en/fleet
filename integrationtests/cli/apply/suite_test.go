package apply

import (
	"context"
	"testing"

	"github.com/rancher/fleet/modules/cli/apply"
	"github.com/rancher/fleet/modules/cli/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFleet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fleet Suite")
}

// simulates fleet cli execution
func fleetApply(name string, dirs []string, options apply.Options) error {

	return apply.Apply(context.Background(), client.NewGetter("", "", "fleet-local"), name, dirs, options)
}
