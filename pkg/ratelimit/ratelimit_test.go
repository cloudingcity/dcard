package ratelimit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLimiter(t *testing.T) {
	const key = "my-key"

	t.Run("default configuration", func(t *testing.T) {
		l := New()
		assert.Equal(t, defaultMax, l.Max)
		assert.Equal(t, defaultTimeout, l.timeout)
	})

	t.Run("hits", func(t *testing.T) {
		l := New(Config{Max: 3, Timeout: 1})

		for i := 1; i <= 3; i++ {
			hit, err := l.Hit(key)
			assert.NoError(t, err)
			assert.Equal(t, i, hit.Count)
		}
		hit, err := l.Hit(key)
		assert.Error(t, err)
		assert.Equal(t, 4, hit.Count)
		assert.Equal(t, 0, hit.Remaining)
		assert.Equal(t, 1, hit.ResetTime)

		time.Sleep(1100 * time.Millisecond)

		hit, err = l.Hit(key)
		assert.NoError(t, err)
		assert.Equal(t, 1, hit.Count)
		assert.Equal(t, 2, hit.Remaining)
		assert.Equal(t, 1, hit.ResetTime)
	})
}
