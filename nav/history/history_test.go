package history

import (
	"reflect"
	"testing"
)

func TestBack(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc         string
		tracker      History[int]
		wantTracker  History[int]
		wantNavState int
	}{
		{
			desc:         "0 entry back",
			tracker:      NewHistory[int](100),
			wantTracker:  NewHistory[int](100),
			wantNavState: 0,
		},
		{
			desc: "1 entry back",
			tracker: History[int]{
				maxStackSize: 100,
				backStack:    []int{1},
				fwdStack:     nil,
			},
			wantTracker: History[int]{
				maxStackSize: 100,
				backStack:    []int{},
				fwdStack:     []int{2},
			},
			wantNavState: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, commit, _ := tc.tracker.Back(2)
			commit()
			if !reflect.DeepEqual(got, tc.wantNavState) {
				t.Errorf("Invalid response for %+q. Got %+q, want %+q", tc.tracker, got, tc.wantNavState)
			}
			if !reflect.DeepEqual(tc.tracker, tc.wantTracker) {
				t.Errorf("Invalid response for tracker state. Got %+q, want %+q", tc.tracker, tc.wantTracker)
			}
		})
	}
}

func TestForward(t *testing.T) {
	t.Parallel()
	tracker := History[string]{
		maxStackSize: -1,
		backStack:    []string{"/1", "/1/2", "/1/2/3"},
		fwdStack:     []string{},
	}
	final := "/1/2/3/4" // start and end state wanted
	_, commit, _ := tracker.Back(final)
	commit()
	_, commit, _ = tracker.Back("/1/2/3")
	commit()
	s, commit, _ := tracker.Back("/1/2")
	commit()
	commit()
	if s != "/1" {
		t.Errorf("Invalid response for Back %+q. Got %+q, want %+q", tracker, s, "/1")
	}

	_, commit, _ = tracker.Foreward("/1")
	commit()
	commit() // call commit multiple times should be a no-op
	commit()
	_, commit, _ = tracker.Foreward("/1/2")
	commit()
	s, commit, _ = tracker.Foreward("/1/2/3")
	commit()
	if s != final {
		t.Errorf("Invalid response for %+q. Got %+q, want %+q", tracker, s, final)
	}

	// fwdStack should be empty so this will cause an error
	_, commit, err := tracker.Foreward("/1/2/3/4")
	commit()
	if err == nil {
		t.Errorf("Invalid response for error. Got %+q, want not nil", err)
	}

	s, commit, _ = tracker.Back(final)
	commit()
	tracker.Go(s)
	if !tracker.ForewardEmpty() {
		t.Errorf("Tracker forward stack should be empty.  Got %+q", tracker.fwdStack)
	}
}

func TestPop(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		desc    string
		in      []int
		out     int
		outRest []int
	}{
		{
			desc:    "regular",
			in:      []int{1, 2, 3},
			out:     3,
			outRest: []int{1, 2},
		},
		{
			desc:    "single item",
			in:      []int{1},
			out:     1,
			outRest: []int{},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			gotVal, gotRest := pop(tc.in)
			if len(tc.in) != len(gotRest)+1 {
				t.Error("In should not have been modified")
			}
			if gotVal != tc.out {
				t.Errorf("Invalid response for %+q. Got %+q, want %+q", tc.in, gotVal, tc.out)
			}
			if !reflect.DeepEqual(gotRest, tc.outRest) {
				t.Errorf("Invalid response for %+q. Got %+q, want %+q", tc.in, gotRest, tc.outRest)
			}
		})
	}
}

func TestAppendMax(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		desc    string
		inSlice []int
		inVal   int
		maxLen  int
		want    []int
	}{
		{
			desc:    "not at limit",
			inSlice: []int{1, 2, 3},
			inVal:   4,
			maxLen:  4,
			want:    []int{1, 2, 3, 4},
		},
		{
			desc:    "at limit",
			inSlice: []int{1, 2, 3},
			inVal:   4,
			maxLen:  3,
			want:    []int{2, 3, 4},
		},
		{
			desc:    "limit 1",
			inSlice: []int{1, 2, 3},
			inVal:   4,
			maxLen:  1,
			want:    []int{4},
		},
		{
			desc:    "limit 2",
			inSlice: []int{1, 2, 3},
			inVal:   4,
			maxLen:  2,
			want:    []int{3, 4},
		},
		{
			desc:    "maxlen less than 1",
			inSlice: []int{1, 2, 3},
			inVal:   4,
			maxLen:  0,
			want:    []int{1, 2, 3, 4},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			got := appendMaxLen(tc.inSlice, tc.inVal, tc.maxLen)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Invalid response for %s. Got %v, want %v", tc.desc, got, tc.want)
			}
		})
	}
}
