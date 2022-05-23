package sturct

import (
	"testing"
)

func errorf(shape Shape, want interface{}, t *testing.T) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("Area One", func(t *testing.T) {
		shape := Rectangle{
			width:  9.5,
			height: 10.5,
		}
		want := 99.75
		errorf(shape, want, t)
	})
	t.Run("Area two", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		errorf(circle, want, t)
	})

	t.Run("表驱动测试示例", func(t *testing.T) {
		areaTests := []struct {
			name    string
			shape   Shape
			hasArea float64
		}{
			{"Rectangle", Rectangle{12, 6}, 72.0},
			{"Circle", Circle{10}, 314.1592653589793},
		}
		for _, v := range areaTests {
			got := v.shape.Area()
			if got != v.hasArea {
				t.Errorf("%#v got %.2f want %.2f", v.shape, got, v.hasArea)
			}
		}
	})
}
func BenchmarkArea(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 伪造数据
		areaTests := []struct {
			name    string
			shape   Shape
			hasArea float64
		}{
			{"Rectangle", Rectangle{12, 6}, 72.0},
			{"Circle", Circle{10}, 314.1592653589793},
			{"Rectangle", Rectangle{10000, 3000}, 3e+07},
			{"Circle", Circle{10000000}, 3.141592653589793e+14},
			{"Rectangle", Rectangle{78912345, 45678978}, 3.60463527118341e+15},
		}
		for _, v := range areaTests {
			got := v.shape.Area()
			if v.hasArea != got {
				b.Errorf("%#v got %.2f want %.2f", v.shape, got, v.hasArea)
			}
		}
	}
}
