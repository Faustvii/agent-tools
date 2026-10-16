package main

import "testing"

func TestHideReviewLink(t *testing.T) {
	for _, tc := range []struct{ name, in, want string }{
		{
			name: "footer with link",
			in:   "<sub>konflate 0.6.4 · rendered `2808f05` · [full diff →](https://konflate.example.com/#/pr/1631)</sub>",
			want: "<sub>konflate 0.6.4 · rendered `2808f05`</sub>",
		},
		{
			name: "footer without link",
			in:   "<sub>konflate 0.6.4 · rendered `2808f05`</sub>",
			want: "<sub>konflate 0.6.4 · rendered `2808f05`</sub>",
		},
		{
			name: "other links stay",
			in:   "see [the chart](https://github.com/example/chart)",
			want: "see [the chart](https://github.com/example/chart)",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := hideReviewLink(tc.in); got != tc.want {
				t.Errorf("hideReviewLink(%q)\n got %q\nwant %q", tc.in, got, tc.want)
			}
		})
	}
}
