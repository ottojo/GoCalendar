package GoCalendar

import "testing"

func TestJulianToGregorian(t *testing.T) {
	cases := []struct {
		inDay, inMonth, inYear, outDay, outMonth, outYear int
	}{
		{29, 12, 1620, 8, 1, 1621},
		{19, 2, 1700, 1, 3, 1700},
		{29, 12, 1899, 10, 1, 1900},
	}
	for _, c := range cases {
		day, month, year := JulianToGregorian(c.inDay, c.inMonth, c.inYear);
		if (day != c.outDay || month != c.outMonth || year != c.outYear) {
			t.Errorf("JulianToGregorian(%v.%v.%v) == %v.%v.%v, want %v.%v.%v", c.inDay, c.inMonth, c.inYear,
				day, month, year,
				c.outDay, c.outMonth, c.outYear);
		}
	}
}

func TestGregorianToJulian(t *testing.T) {
	cases := []struct {
		inDay, inMonth, inYear, outDay, outMonth, outYear int
	}{
		{8, 1, 1621, 29, 12, 1620},
		{1, 3, 1700, 19, 2, 1700},
		{10, 1, 1900, 29, 12, 1899},
	}
	for _, c := range cases {
		day, month, year := GregorianToJulian(c.inDay, c.inMonth, c.inYear);
		if (day != c.outDay || month != c.outMonth || year != c.outYear) {
			t.Errorf("GregorianToJulian(%v.%v.%v) == %v.%v.%v, want %v.%v.%v", c.inDay, c.inMonth, c.inYear,
				day, month, year,
				c.outDay, c.outMonth, c.outYear);
		}
	}
}

func TestAddDays(t *testing.T) {
	cases := []struct {
		inDay, inMonth, inYear, inDifference, outDay, outMonth, outYear int
	}{
		{27, 2, 2016, 5, 3, 3, 2016},
		{3, 2, 2017, -35, 30, 12, 2016},
	}
	for _, c := range cases {
		day, month, year := AddDays(c.inDay, c.inMonth, c.inYear, c.inDifference, "g");
		if (day != c.outDay || month != c.outMonth || year != c.outYear) {
			t.Errorf("AddDays(%v.%v.%v + %v) == %v.%v.%v, want %v.%v.%v", c.inDay, c.inMonth, c.inYear, c.inDifference,
				day, month, year,
				c.outDay, c.outMonth, c.outYear);
		}
	}
}