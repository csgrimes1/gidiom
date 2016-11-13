package iterate_test

import (
	"github.com/csgrimes1/gidiom/iterate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToInt(t *testing.T) {
	assert.Equal(t, 2, iterate.CreateAny(2.0).To().Int())
	assert.Equal(t, 2, iterate.CreateAny(2).To().Int())
	assert.Equal(t, 2, iterate.CreateAny("2").To().Int())
	assert.Equal(t, int8(-8), iterate.CreateAny("-8").To().Int8())
	assert.Equal(t, int16(-16), iterate.CreateAny("-16").To().Int16())
	assert.Equal(t, int32(-32), iterate.CreateAny("-32").To().Int32())
}

func TestToUint(t *testing.T) {
	assert.Equal(t, uint(2), iterate.CreateAny(2.0).To().Uint())
	assert.Equal(t, uint(2), iterate.CreateAny(2).To().Uint())
	assert.Equal(t, uint(2), iterate.CreateAny("2").To().Uint())
	assert.Equal(t, uint8(8), iterate.CreateAny("8").To().Uint8())
	assert.Equal(t, uint16(16), iterate.CreateAny("16").To().Uint16())
	assert.Equal(t, uint32(32), iterate.CreateAny("32").To().Uint32())
}

func TestToFloat(t *testing.T) {
	assert.Equal(t, float64(2.1), iterate.CreateAny(2.1).To().Float64())
	assert.Equal(t, float32(-32), iterate.CreateAny(-32).To().Float32())
	assert.Equal(t, float64(-64), iterate.CreateAny("-64").To().Float64())
}

