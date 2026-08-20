package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask13(t *testing.T) {
	now := time.Now()
	s := NewService(NewRegistry(), func() time.Time { return now })
	due, err := s.CheckCorrectiveDeadline(context.Background(), "critical", now)
	require.NoError(t, err)
	require.Equal(t, now.UTC().Add(24*time.Hour), due)
}
