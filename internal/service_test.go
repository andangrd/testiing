package internal_test

import (
	"testing"

	"github.com/andangrd/testiing/internal"
	"github.com/stretchr/testify/assert"
)

var (
	svc *internal.Service = internal.ProvideService(3)
)

func AddCounterTest(t *testing.T) {
	expectedResult := 7
	t.Logf("HelloCount : %d", svc.HelloCount)
	result := svc.AddCounter(4)
	t.Logf("HelloCount : %d", result)

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, result, expectedResult)
	})
}
