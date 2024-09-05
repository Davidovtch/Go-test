package refl

import (
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Austin"},
			},
			[]string{"London", "Austin"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Austin"},
			},
			[]string{"London", "Austin"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)

			})
			if !slices.Equal(got, test.ExpectedCalls) {
				t.Errorf("got %q want %q", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		m := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(m, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("channels", func(t *testing.T) {
		var c = make(chan Profile)

		go func() {
			c <- Profile{33, "Berlin"}
			c <- Profile{34, "Paris"}
			close(c)
		}()

		var got []string
		want := []string{"Berlin", "Paris"}

		walk(c, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v expected:%v", got, want)
		}
	})

	t.Run("function", func(t *testing.T) {
		fn := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Paris"}
		}
		var got []string
		want := []string{"Berlin", "Paris"}

		walk(fn, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v expected:%v", got, want)
		}

	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
