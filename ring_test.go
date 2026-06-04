package ring_test

import (
	"strconv"
	"testing"

	"github.com/mga135/ring"
	"github.com/stretchr/testify/require"
)

func TestRing(t *testing.T) {
	t.Parallel()

	type run struct {
		toAdd   int64
		headExp int64
		tailExp int64
		dataExp []int64
	}

	type testCase struct {
		capacity int
		errExp   error
		runs     []run
	}

	testCases := []testCase{
		{
			capacity: 0,
			errExp:   ring.ErrInvalidCapacity,
			runs:     nil,
		},
		{
			capacity: 1,
			errExp:   nil,
			runs: []run{
				{1, 1, 1, []int64{1}},
				{2, 2, 2, []int64{2}},
			},
		},
		{
			capacity: 3,
			errExp:   nil,
			runs: []run{
				{1, 1, 1, []int64{1}},
				{2, 1, 2, []int64{1, 2}},
				{3, 1, 3, []int64{1, 2, 3}},
				{4, 2, 4, []int64{2, 3, 4}},
				{5, 3, 5, []int64{3, 4, 5}},
				{6, 4, 6, []int64{4, 5, 6}},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			rg, err := ring.New[int64](tc.capacity)
			require.ErrorIs(t, err, tc.errExp)

			if err != nil {
				require.Nil(t, rg)

				return
			}

			require.NotNil(t, rg)
			require.Equal(t, tc.capacity, rg.Cap())

			for j, r := range tc.runs {
				rg.Add(r.toAdd)

				require.Equal(t, r.headExp, rg.Head())
				require.Equal(t, r.tailExp, rg.Tail())
				require.Equal(t, r.dataExp, rg.Data())
				require.Equal(t, min(tc.capacity, j+1), rg.Size())
				require.Equal(t, tc.capacity, rg.Cap())
			}
		})
	}
}
