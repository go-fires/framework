package hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManager(t *testing.T) {
	m := NewManager(&Config{
		Driver: "bcrypt",
	})

	// default driver
	assert.Equal(t, "bcrypt", m.getDefaultDriver())

	// default driver
	assert.True(t, NewBcryptHasher().Check("123456", m.Driver().MustMake("123456")))
	assert.False(t, NewMd5Hasher().Check("123456", m.Driver().MustMake("123456")))
	assert.True(t, NewBcryptHasher().Check("123456", m.MustMake("123456")))
	assert.False(t, NewMd5Hasher().Check("123456", m.MustMake("123456")))
	assert.True(t, m.Check("123456", m.MustMake("123456")))

	// bcrypt driver
	assert.True(t, NewBcryptHasher().Check("123456", m.Driver("bcrypt").MustMake("123456")))
	assert.False(t, NewMd5Hasher().Check("123456", m.Driver("bcrypt").MustMake("123456")))

	// md5 driver
	assert.True(t, NewMd5Hasher().Check("123456", m.Driver("md5").MustMake("123456")))
	assert.False(t, NewBcryptHasher().Check("123456", m.Driver("md5").MustMake("123456")))

	// sha1 driver
	assert.True(t, NewSha1Hasher().Check("123456", m.Driver("sha1").MustMake("123456")))
	assert.False(t, NewSha1Hasher().Check("123456", m.Driver("sha1").MustMake("1234567")))

	// unknown driver
	assert.Panicsf(t, func() { m.Driver("unknown") }, "unknown driver: unknown")
}

func BenchmarkManager(b *testing.B) {
	m := NewManager(&Config{
		Driver: "bcrypt",
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			go m.Driver("md5").MustMake("123456")
		}
	})
}
