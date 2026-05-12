package common

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNowUnix(t *testing.T) {
	before := time.Now().Unix()
	got := NowUnix()
	after := time.Now().Unix()

	assert.GreaterOrEqual(t, got, before)
	assert.LessOrEqual(t, got, after)
}
