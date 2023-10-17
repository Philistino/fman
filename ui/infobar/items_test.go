package infobar

import "testing"

func TestView(t *testing.T) {
	t.Run("should return the correct view", func(t *testing.T) {
		m := itemTracker{
			nItems:    10,
			nSelected: 2,
		}
		m.Init()
		expected := " 2 selected │ 10 items "
		actual := m.View()
		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})
}
