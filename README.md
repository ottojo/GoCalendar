# GoCalendar
## GoCalendar provides utility to convert dates between the julian and the gregorian calendar.


##Usage
```go
import "github.com/ottojo/GoCalendar"

day, month, year := GoCalendar.GregorianToJulian(15, 07, 1337);

println("Date in julian calendar: " +
  strconv.Itoa(day) + "." +
  strconv.Itoa(month) + "." +
  strconv.Itoa(year));
  
day, month, year = GoCalendar.AddDays(day, month, year, 5, "j");    //"j" - Julian calendar, "g" - Gregorian calendar
println("5 days later: : " +
  strconv.Itoa(day) + "." +
  strconv.Itoa(month) + "." +
  strconv.Itoa(year));
```
