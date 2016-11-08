package GoCalendar

var julian_months = [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};

var gregorian_months = [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};

//NOTE: This does not provide accurate results leap years before 8 AD

func JulianToGregorian(day int, month int, year int) (int, int, int) {
	dayDifference := GregorianJulianDifference(day, month, year);
	return AddDays(day, month, year, dayDifference, "g");
}

func GregorianToJulian(day int, month int, year int) (int, int, int) {
	dayDifference := GregorianJulianDifference(day, month, year)
	return AddDays(day, month, year, -dayDifference, "j");
}

func GregorianJulianDifference(_ int, month int, year int) int {
	century := year / 100;
	if (month <= 2) {
		century--;
	}
	a := century / 4;
	b := century % 4;
	dayDifference := 3 * a + b - 2;
	return dayDifference;
}

func monthsInYear(year int, calendar string) [13]int {
	months := gregorian_months;
	switch calendar{
	case "j":
		months = julian_months;
		if (year % 4 == 0) {
			months[2] = 29;
		}
		break;
	default:
		months = gregorian_months;
		if (((year % 4 == 0 ) && (year % 100 != 0)) || (year % 400 == 0)) {
			months[2] = 29;
		}
		break;
	}
	return months;
}

func AddDays(day int, month int, year int, difference int, calendar string) (int, int, int) {
	if (difference > 0) {
		day++;
		if (day > monthsInYear(year, calendar)[month]) {
			day = 1;
			month++;
			if (month > 12) {
				month = 1;
				year++;
			}
		}
		difference--;
	} else if (difference < 0) {
		day--;
		if (day < 1) {
			month--;
			if (month < 1) {
				month = 12;
				year--;
			}
			day = monthsInYear(year, calendar)[month];
		}
		difference++;
	} else {
		return day, month, year;
	}
	return AddDays(day, month, year, difference, calendar);
}
