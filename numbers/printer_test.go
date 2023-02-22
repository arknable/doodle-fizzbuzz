package numbers

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrinter(t *testing.T) {
	testCases := []struct {
		name      string
		from      int
		to        int
		stringer  Stringer
		expError  error
		expOutput string
	}{
		{
			name:     "From is negative",
			from:     -1,
			expError: ErrNegativeRange,
		},
		{
			name:     "From is positive, but To is negative",
			from:     0,
			to:       -1,
			expError: ErrNegativeRange,
		},
		{
			name:      "0 to 10",
			from:      0,
			to:        10,
			expOutput: "0\n1\n2\n3\n4\n5\n6\n7\n8\n9\n",
		},
		{
			name:      "5 to 22",
			from:      5,
			to:        22,
			expOutput: "5\n6\n7\n8\n9\n10\n11\n12\n13\n14\n15\n16\n17\n18\n19\n20\n21\n",
		},
		{
			name: "Multiple of 3 print Fizz",
			from: 0,
			to:   10,
			stringer: func(n int) string {
				if n > 0 && n%3 == 0 {
					return "Fizz"
				}
				return strconv.Itoa(n)
			},
			expOutput: "0\n1\n2\nFizz\n4\n5\nFizz\n7\n8\nFizz\n",
		},
		{
			name: "Multiple of 5 print Buzz",
			from: 5,
			to:   21,
			stringer: func(n int) string {
				if (n > 0) && (n%5 == 0) {
					return "Buzz"
				}
				return strconv.Itoa(n)
			},
			expOutput: "Buzz\n6\n7\n8\n9\nBuzz\n11\n12\n13\n14\nBuzz\n16\n17\n18\n19\nBuzz\n",
		},
		{
			name: "Multiple of 3 and 5 print FizzBuzz",
			from: 5,
			to:   20,
			stringer: func(n int) string {
				if (n > 0) && (n%5 == 0) && (n%3 == 0) {
					return "FizzBuzz"
				}
				return strconv.Itoa(n)
			},
			expOutput: "5\n6\n7\n8\n9\n10\n11\n12\n13\n14\nFizzBuzz\n16\n17\n18\n19\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := NewPrinter(tc.from, tc.to)
			if tc.expError != nil {
				require.Nil(t, p)
				require.ErrorIs(t, err, tc.expError)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, p)

			if tc.stringer != nil {
				p.WithStringer(tc.stringer)
			}

			buff := &bytes.Buffer{}
			p.out = buff

			p.Print()

			require.NoError(t, err)
			require.Equal(t, tc.expOutput, buff.String())
		})
	}
}
