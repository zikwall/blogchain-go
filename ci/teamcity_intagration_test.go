// +build integration teamcity

package ci

import (
	"testing"
)

func TestTeamcityDummy(t *testing.T) {
	t.Run("it should be always return true", func(t *testing.T) {
		return true
	})
}
