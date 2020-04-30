package hashmultisets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMultiSet(t *testing.T) {
	set := New()
	set.Add("hello")

	assert.Equal(t, set.Contains("hello"), true)
	assert.Equal(t, set.GetCount("hello"), 1)
	assert.Equal(t, set.Contains("unknown"), false)
	assert.Equal(t, set.GetCount("unknown"), 0)
	assert.Equal(t, set.ContainsAll("hello", "unknown"), false)
	assert.Equal(t, set.ContainsAny("hello", "unknown"), true)
	assert.Equal(t, set.Size(), 1)
	assert.Equal(t, set.GetValues()[0], "hello")
}

func TestHashMultiSetMerge(t *testing.T) {
	set1 := New()
	set1.Add(true, "true")

	set2 := New(true, 5)

	set3 := New()
	set3.Merge(set1, set2)

	assert.Equal(t, set3.Size(), 3)
	assert.Equal(t, set3.GetCount(true), 2)
	assert.Equal(t, set3.GetCount("true"), 1)
	assert.Equal(t, set3.GetCount(5), 1)
	assert.Equal(t, set3.ContainsAll(true, "true", 5), true)
}

func TestHashMultiSetTop(t *testing.T) {
	set := New()
	set.IncrementBy("key1", 10)
	set.IncrementBy("key2", 15)

	pairs := set.GetTopValues()
	assert.Equal(t, len(pairs), 2)

	pair1 := pairs[0]
	assert.Equal(t, pair1.Key, "key2")
	assert.Equal(t, pair1.Count, 15)

	pair2 := pairs[1]
	assert.Equal(t, pair2.Key, "key1")
	assert.Equal(t, pair2.Count, 10)
}
