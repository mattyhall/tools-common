package value

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersionOlder(t *testing.T) {
	type testCase struct {
		name     string
		first    Version
		second   Version
		expected bool
	}

	tests := []testCase{
		{
			name:   "SameVersion",
			first:  Version5_0_0,
			second: Version5_0_0,
		},
		{
			name:     "5.0.0ToUnknown",
			first:    Version5_0_0,
			second:   VersionUnknown,
			expected: true,
		},
		{
			name:   "UnknownTo5.0.0",
			first:  VersionUnknown,
			second: Version5_0_0,
		},
		{
			name:   "UnknownTo7.0.0",
			first:  VersionUnknown,
			second: Version7_0_0,
		},
		{
			name:   "UnknownToUnknown",
			first:  VersionUnknown,
			second: VersionUnknown,
		},
		{
			name:     "5.0.0ToEmpty",
			first:    Version5_0_0,
			expected: true,
		},
		{
			name:   "EmptyTo5.0.0",
			second: Version5_0_0,
		},
		{
			name:     "10.2.4To42.1.5",
			first:    Version("10.2.4"),
			second:   Version("42.1.5"),
			expected: true,
		},
		{
			name:   "Columnar1.0.0To7.2.0",
			first:  VersionColumnar1_0_0,
			second: Version7_2_0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.first.Older(tc.second))
		})
	}
}

func TestVersionNewer(t *testing.T) {
	type testCase struct {
		name     string
		first    Version
		second   Version
		expected bool
	}

	tests := []testCase{
		{
			name:   "SameVersion",
			first:  Version5_0_0,
			second: Version5_0_0,
		},
		{
			name:   "5.0.0ToUnknown",
			first:  Version5_0_0,
			second: VersionUnknown,
		},
		{
			name:     "UnknownTo5.0.0",
			first:    VersionUnknown,
			second:   Version5_0_0,
			expected: true,
		},
		{
			name:   "UnknownTo8.0.0",
			first:  VersionUnknown,
			second: Version8_0_0,
		},
		{
			name:   "UnknownToUnknown",
			first:  VersionUnknown,
			second: VersionUnknown,
		},
		{
			name:  "5.0.0ToEmpty",
			first: Version5_0_0,
		},
		{
			name:     "EmptyTo5.0.0",
			second:   Version5_0_0,
			expected: true,
		},
		{
			name:   "10.2.4To42.1.5",
			first:  Version("10.2.4"),
			second: Version("42.1.5"),
		},
		{
			name:     "Columnar1.0.0To7.2.0",
			first:    VersionColumnar1_0_0,
			second:   Version7_2_0,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.first.Newer(tc.second))
		})
	}
}

func TestVersionAtLeast(t *testing.T) {
	type test struct {
		name     string
		first    Version
		second   Version
		expected bool
	}

	tests := []*test{
		{
			name:     "SameVersion",
			first:    Version5_0_0,
			second:   Version5_0_0,
			expected: true,
		},
		{
			name:   "5.0.0ToUnknown",
			first:  Version5_0_0,
			second: VersionUnknown,
		},
		{
			name:     "UnknownTo5.0.0",
			first:    VersionUnknown,
			second:   Version5_0_0,
			expected: true,
		},
		{
			name:   "5.0.0To6.5.1",
			first:  Version5_0_0,
			second: Version("6.5.1"),
		},
		{
			name:   "5.0.0To5.5.0",
			first:  Version5_0_0,
			second: Version("5.5.0"),
		},
		{
			name:     "6.5.1To5.0",
			first:    Version("6.5.0"),
			second:   Version5_0_0,
			expected: true,
		},
		{
			name:     "5.5.0To5.0.0",
			first:    Version("5.5.0"),
			second:   Version5_0_0,
			expected: true,
		},
		{
			name:     "UnknownTo6.6.0",
			first:    VersionUnknown,
			second:   Version6_6_0,
			expected: true,
		},
		{
			name:     "UnknownTo7.0.0",
			first:    VersionUnknown,
			second:   Version7_0_0,
			expected: true,
		},
		{
			name:     "EmptyTo6.6.0",
			second:   Version6_6_0,
			expected: true,
		},
		{
			name:  "6.6.0ToEmpty",
			first: Version6_0_0,
		},
		{
			name:   "10.2.4To42.1.5",
			first:  Version("10.2.4"),
			second: Version("42.1.5"),
		},
		{
			name:     "7.6.0ToColumnar1.0.0",
			first:    Version7_6_0,
			second:   VersionColumnar1_0_0,
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, test.first.AtLeast(test.second))
		})
	}
}

func TestVersionEqual(t *testing.T) {
	type test struct {
		name     string
		first    Version
		second   Version
		expected bool
	}

	tests := []*test{
		{
			name:     "Equal",
			first:    Version5_0_0,
			second:   Version5_0_0,
			expected: true,
		},
		{
			name:   "NotEqual",
			first:  Version5_0_0,
			second: VersionLatest,
		},
		{
			name:     "EqualJustAString",
			first:    Version("test"),
			second:   Version("test"),
			expected: true,
		},
		{
			name:   "NotEqualJustAString",
			first:  Version("test"),
			second: Version("testing"),
		},
		{
			name:   "10.2.4To42.1.5",
			first:  Version("10.2.4"),
			second: Version("42.1.5"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, test.first.Equal(test.second))
		})
	}
}
