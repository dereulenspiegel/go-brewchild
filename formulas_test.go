package brewchild

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSGToPlato(t *testing.T) {
	for _, tc := range []struct {
		sg            float64
		expectedPlato float64
	}{
		{
			sg:            1.048375709,
			expectedPlato: 12.0,
		},
		{
			sg:            1.056,
			expectedPlato: 13.8,
		},
	} {
		plato := SGToPlato(tc.sg)
		assert.InDelta(t, tc.expectedPlato, plato, 0.1)
	}
}
