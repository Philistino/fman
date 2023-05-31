package entry

// reimplement with mock
// func TestGetEntries(t *testing.T) {

// 	testCases := []struct {
// 		desc         string
// 		path         string
// 		expectedSize int
// 	}{
// 		{
// 			desc:         "cur dir",
// 			path:         "./",
// 			expectedSize: 4,
// 		},
// 	}

// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			entries, _ := GetEntries(tC.path, true, false)
// 			if len(entries) != tC.expectedSize {
// 				t.Errorf("expecting %d entries, got %d", tC.expectedSize, len(entries))
// 			}
// 		})
// 	}
// }

// func TestSortEntries(t *testing.T) {
// 	tt := []struct {
// 		name      string
// 		inEntries []Entry
// 		want      []Entry
// 	}{
// 		{
// 			name: "first",
// 			inEntries: []Entry{
// 				{Name: "1 file", IsDir: false},
// 				{Name: "2 file", IsDir: false},
// 				{Name: "1 dir", IsDir: true},
// 				{Name: "2 dir", IsDir: true},
// 			},
// 			want: []Entry{
// 				{Name: "1 dir", IsDir: true},
// 				{Name: "2 dir", IsDir: true},
// 				{Name: "1 file", IsDir: false},
// 				{Name: "2 file", IsDir: false},
// 			},
// 		},
// 	}
// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			entries := sortEntries(tc.inEntries)
// 			for _, e := range entries {
// 				log.Println(e)
// 			}
// 			if reflect.DeepEqual(tc.want, entries) {
// 				t.Errorf("expecting %v entries, got %v", entries, len(entries))
// 			}
// 		})
// 		t.Error()
// 	}

// }
