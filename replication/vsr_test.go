package replication

import (
	"math"
	"testing"
)

func _test_assert(t *testing.T, condition bool, message string, args ...any) {
	if !condition {
		t.Errorf(message, args...)
	}
}

func TestCommitTable_Truncate(t *testing.T) {
	ct := CommitTable{}
	ct.Init(0)

	ct.IncrementCommitCount(10)
	ct.IncrementCommitCount(3)
	ct.IncrementCommitCount(4)
	ct.IncrementCommitCount(1)
	ct.IncrementCommitCount(2)
	ct.IncrementCommitCount(2)
	ct.IncrementCommitCount(2)
	ct.IncrementCommitCount(5)
	ct.IncrementCommitCount(6)

	_test_assert(t, ct.GetCommitCount(1) == 1, "commit index 1 should be 1")
	_test_assert(t, ct.GetCommitCount(2) == 3, "commit index 2 should be 3")
	_test_assert(t, ct.GetCommitCount(3) == 1, "commit index 3 should be 1")
	_test_assert(t, ct.GetCommitCount(4) == 1, "commit index 4 should be 1")
	_test_assert(t, ct.GetCommitCount(5) == 1, "commit index 5 should be 1")
	_test_assert(t, ct.GetCommitCount(6) == 1, "commit index 6 should be 1")
	_test_assert(t, ct.GetCommitCount(7) == 0, "commit index 7 should be 0")
	_test_assert(t, ct.GetCommitCount(8) == 0, "commit index 8 should be 0")
	_test_assert(t, ct.GetCommitCount(9) == 0, "commit index 9 should be 0")
	_test_assert(t, ct.GetCommitCount(10) == 1, "commit index 10 should be 1")

	// truncate at 3
	ct.Truncate(3)
	// [0, 3) should be deleted

	_test_assert(t, ct.GetCommitCount(1) == math.MaxUint64, "commit index 1 should be MaxUint64")
	_test_assert(t, ct.GetCommitCount(2) == math.MaxUint64, "commit index 2 should be MaxUint64")
	_test_assert(t, ct.GetCommitCount(3) == 1, "commit index 3 should be 1, got %d", ct.GetCommitCount(3))
	_test_assert(t, ct.GetCommitCount(4) == 1, "commit index 4 should be 1, got %d", ct.GetCommitCount(4))
	_test_assert(t, ct.GetCommitCount(5) == 1, "commit index 5 should be 1")
	_test_assert(t, ct.GetCommitCount(6) == 1, "commit index 6 should be 1")
	_test_assert(t, ct.GetCommitCount(7) == 0, "commit index 7 should be 0")
	_test_assert(t, ct.GetCommitCount(8) == 0, "commit index 8 should be 0")
	_test_assert(t, ct.GetCommitCount(9) == 0, "commit index 9 should be 0")
	_test_assert(t, ct.GetCommitCount(10) == 1, "commit index 10 should be 1")

}
