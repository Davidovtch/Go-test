package smi

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, s Shape, want float64) {
		t.Helper()
		got := s.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0

		checkArea(t, rectangle, want)

	})
	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})
}

/*table tests are good for when you need to test implementations
of a interface or if the data being
passed needs lots of requirements*/

func TestArea2(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		/*the fields were named only to increase readability,
		their names need to be the same as in the struct*/
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
