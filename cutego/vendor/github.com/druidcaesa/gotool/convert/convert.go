package convert

//Conversion of Gregorian Calendar to Lunar Calendar Tool Set
//Thanks to the open source project provided by zhuqiyang, most of this tool is compiled from zhuqiyang
import (
	"errors"
	"fmt"
	"math"
	"time"
)

type Convert struct {
}

var lunarInfomation = [200]int{
	0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, //1900-1909
	0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, //1910-1919
	0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, //1920-1929
	0x06566, 0x0d4a0, 0x0ea50, 0x16a95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, //1930-1939
	0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, //1940-1949
	0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5b0, 0x14573, 0x052b0, 0x0a9a8, 0x0e950, 0x06aa0, //1950-1959
	0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, //1960-1969
	0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b6a0, 0x195a6, //1970-1979
	0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, //1980-1989
	0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x05ac0, 0x0ab60, 0x096d5, 0x092e0, //1990-1999
	0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, //2000-2009
	0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, //2010-2019
	0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, //2020-2029
	0x05aa0, 0x076a3, 0x096d0, 0x04afb, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, //2030-2039
	0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, //2040-2049
	0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06aa0, 0x1a6c4, 0x0aae0, //2050-2059
	0x092e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, //2060-2069
	0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, //2070-2079
	0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, //2080-2089
	0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252, //2090-2099
}

var yearSum = map[int]int{
	1900: 384, 1901: 354, 1902: 355, 1903: 383, 1904: 354, 1905: 355, 1906: 384, 1907: 354, 1908: 355, 1909: 384,
	1910: 354, 1911: 384, 1912: 354, 1913: 354, 1914: 384, 1915: 354, 1916: 355, 1917: 384, 1918: 355, 1919: 384,
	1920: 354, 1921: 354, 1922: 384, 1923: 354, 1924: 354, 1925: 385, 1926: 354, 1927: 355, 1928: 384, 1929: 354,
	1930: 383, 1931: 354, 1932: 355, 1933: 384, 1934: 355, 1935: 354, 1936: 384, 1937: 354, 1938: 384, 1939: 354,
	1940: 354, 1941: 384, 1942: 355, 1943: 354, 1944: 385, 1945: 354, 1946: 354, 1947: 384, 1948: 354, 1949: 384,
	1950: 354, 1951: 355, 1952: 384, 1953: 354, 1954: 355, 1955: 384, 1956: 354, 1957: 383, 1958: 355, 1959: 354,
	1960: 384, 1961: 355, 1962: 354, 1963: 384, 1964: 355, 1965: 353, 1966: 384, 1967: 355, 1968: 384, 1969: 354,
	1970: 355, 1971: 384, 1972: 354, 1973: 354, 1974: 384, 1975: 354, 1976: 384, 1977: 354, 1978: 355, 1979: 384,
	1980: 355, 1981: 354, 1982: 384, 1983: 354, 1984: 384, 1985: 354, 1986: 354, 1987: 384, 1988: 355, 1989: 355,
	1990: 384, 1991: 354, 1992: 354, 1993: 383, 1994: 355, 1995: 384, 1996: 354, 1997: 355, 1998: 384, 1999: 354,
	2000: 354, 2001: 384, 2002: 354, 2003: 355, 2004: 384, 2005: 354, 2006: 385, 2007: 354, 2008: 354, 2009: 384,
	2010: 354, 2011: 354, 2012: 384, 2013: 355, 2014: 384, 2015: 354, 2016: 355, 2017: 384, 2018: 354, 2019: 354,
	2020: 384, 2021: 354, 2022: 355, 2023: 384, 2024: 354, 2025: 384, 2026: 354, 2027: 354, 2028: 384, 2029: 355,
	2030: 354, 2031: 384, 2032: 355, 2033: 384, 2034: 354, 2035: 354, 2036: 384, 2037: 354, 2038: 354, 2039: 384,
	2040: 355, 2041: 355, 2042: 384, 2043: 354, 2044: 384, 2045: 354, 2046: 354, 2047: 384, 2048: 354, 2049: 355,
	2050: 384, 2051: 355, 2052: 384, 2053: 354, 2054: 354, 2055: 383, 2056: 355, 2057: 354, 2058: 384, 2059: 355,
	2060: 354, 2061: 384, 2062: 354, 2063: 384, 2064: 354, 2065: 355, 2066: 384, 2067: 354, 2068: 355, 2069: 384,
	2070: 354, 2071: 384, 2072: 354, 2073: 354, 2074: 384, 2075: 355, 2076: 354, 2077: 384, 2078: 355, 2079: 354,
	2080: 384, 2081: 354, 2082: 384, 2083: 354, 2084: 355, 2085: 384, 2086: 354, 2087: 355, 2088: 383, 2089: 354,
	2090: 384, 2091: 354, 2092: 355, 2093: 384, 2094: 355, 2095: 354, 2096: 384, 2097: 354, 2098: 354, 2099: 384,
}

// GregorianToLunarCalendar 公历转农历
// GregorianToLunarCalendar Gregorian to Lunar Calendar
func (c Convert) GregorianToLunarCalendar(solarYear, solarMonth, solarDay int) [3]int {
	var date string = fmt.Sprintf("%04d-%02d-%02d", solarYear, solarMonth, solarDay)

	// 1900-1-31 时间戳
	var timestamp int64 = -2206425600
	timenumber, err := c.getTimestamp(date)
	if err != nil {
		fmt.Println("getTimestamp err = ", err)
	}
	days := int(math.Ceil(float64(timenumber-timestamp)/86400) + 1)

	var sum int = 0 // 农历天数之和
	var count int = len(yearSum) + 1900
	var lunarYear int

	for lunarYear = 1900; lunarYear < count; lunarYear++ {
		sum += yearSum[lunarYear]
		if sum >= days {
			break
		}
	}
	olddays := yearSum[lunarYear] - (sum - days)
	hex := lunarInfomation[lunarYear-1900]
	sumMonth := 0
	leapMonth := hex & 0xf
	isleap := false

	var i int
	var month int = 1
	for i = 0x08000; i >= 0x00010; i >>= 1 {
		if hex&i > 0 {
			sumMonth += 30
		} else {
			sumMonth += 29
		}
		if sumMonth >= olddays {
			break
		}
		if leapMonth == month {
			if hex&0xf0000 > 0 {
				sumMonth += 30
			} else {
				sumMonth += 29
			}
			isleap = true
			if sumMonth >= olddays {
				break
			}
		}
		month++
	}
	var currentMonthDays int
	if leapMonth == month && isleap {
		if hex&0xf0000 > 0 {
			currentMonthDays = 30
		} else {
			currentMonthDays = 29
		}
	} else {
		if hex&(0x08000>>(uint(month)-1)) > 0 {
			currentMonthDays = 30
		} else {
			currentMonthDays = 29
		}
		leapMonth = 0
	}
	day := currentMonthDays - (sumMonth - olddays)
	return [3]int{lunarYear, month, day}
}

// LunarToGregorian 农历转公历
func (c Convert) LunarToGregorian(year, month, day int, leap bool) [3]int {
	var sum int = 0
	for i := 1900; i < year; i++ {
		sum += yearSum[i]
	}
	// 加上最后一年
	hex := lunarInfomation[year-1900]
	leapMonth := hex & 0xf
	lunarMonth := 1
	for i := 0x08000; i >= 0x00010; i >>= 1 {
		if month == lunarMonth {
			break
		}
		if hex&i > 0 {
			sum += 30
		} else {
			sum += 29
		}
		if leapMonth == lunarMonth {
			if hex&0xf0000 > 0 {
				sum += 30
			} else {
				sum += 29
			}
		}
		lunarMonth++
	}
	if month == leapMonth && leap {
		if hex&(0x08000>>(uint(lunarMonth)-1)) > 0 {
			sum += 30
		} else {
			sum += 29
		}
	}
	sum = sum + day // 这个和是农历天数的总和

	count := len(yearSum) + 1900
	solarSum := -30
	var solarYear int
	for solarYear = 1900; solarYear < count; solarYear++ {
		solarSum += c.GetSolarYearDays(solarYear)
		if solarSum >= sum {
			break
		}
	}

	sumMonth := c.GetSolarYearDays(solarYear) - (solarSum - sum)
	tempSum := 0
	solarMonth := 1
	for i := 0; i < 12; i++ {
		tempSum += c.GetSolarMonthDays(solarYear, solarMonth)
		if tempSum >= sumMonth {
			break
		}
		solarMonth++
	}
	lastMonth := c.GetSolarMonthDays(solarYear, solarMonth)
	solarDay := lastMonth - (tempSum - sumMonth)
	return [3]int{solarYear, solarMonth, solarDay}
}

// 获取时间戳
func (Convert) getTimestamp(timeStr string) (int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return -1, errors.New(fmt.Sprintf("time.LoadLocation %s", err))
	}
	Time, err := time.ParseInLocation("2006-01-02", timeStr, loc)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("time.ParseInLocation %s", err))
	}
	return Time.Unix(), nil
}

// GetLunarYearDays 获取农历某年全年天数
//Get the number of days in a year in the lunar calendar
func (Convert) GetLunarYearDays(year int) int {
	hex := lunarInfomation[year-1900]
	sum := 0
	for i := 0x08000; i >= 0x00010; i >>= 1 {
		if hex&i > 0 {
			sum += 30
		} else {
			sum += 29
		}
	}
	if hex&0xf > 0 {
		if hex&0xf0000 > 0 {
			sum += 30
		} else {
			sum += 29
		}
	}
	return sum
}

// 获取个农历日期到1900年的所有天数
// Get all the days from a lunar date to 1900
func getLunarTotalDays(year, month, day int, leap bool) int {
	var sum int = 0
	for i := 1900; i < year; i++ {
		sum += yearSum[i]
	}
	// 加上最后一年
	hex := lunarInfomation[year-1900]
	leapMonth := hex & 0xf
	lunarMonth := 1
	for i := 0x08000; i >= 0x00010; i >>= 1 {
		if month == lunarMonth {
			break
		}
		if hex&i > 0 {
			sum += 30
		} else {
			sum += 29
		}
		if leapMonth == lunarMonth {
			if hex&0xf0000 > 0 {
				sum += 30
			} else {
				sum += 29
			}
		}
		lunarMonth++
	}
	if month == leapMonth && leap {
		if hex&(0x08000>>(uint(lunarMonth)-1)) > 0 {
			sum += 30
		} else {
			sum += 29
		}
	}
	return sum + day // 这个和是农历天数的总和
}

// GetDayGanZhiIndex Get the index value of the flow of heaven and earth, starting from 0
func (Convert) GetDayGanZhiIndex(year, month, day int, leap bool) int {
	days := 43041 // 甲子日天数和
	differ := getLunarTotalDays(year, month, day, leap) - days
	index := differ % 60
	if index < 0 {
		index += 60
	}
	return index
}

// GetLunarMonthDays 获取农历某月天数
//Get the number of days in a lunar month
func (Convert) GetLunarMonthDays(year, month int, leapMonth bool) int {
	hex := lunarInfomation[year-1900]
	if leapMonth {
		if hex&0xf > 0 {
			if hex&0xf0000 > 0 {
				return 30
			} else {
				return 29
			}
		} else {
			return 0
		}
	} else {
		if hex&(0x08000>>(uint(month)-1)) > 0 {
			return 30
		} else {
			return 29
		}
	}
}

// GetLunarLeapMonth 获取农历某年闰月的月份,为零的时候没有闰月
//Get the month of the leap month of a certain year in the lunar calendar, when it is zero, there is no leap month
func (Convert) GetLunarLeapMonth(year int) int {
	return lunarInfomation[year-1900] & 0xf
}

// GetLunarLeapMonthDays 某年闰月天数
//Leap month days in a year
func (Convert) GetLunarLeapMonthDays(year int) int {
	hex := lunarInfomation[year-1900]
	if hex&0xf > 0 {
		if hex&0xf0000 > 0 {
			return 30
		} else {
			return 29
		}
	}
	return -1
}

// GetSolarYearDays 获取公历全年天数
//Get the number of days in the Gregorian calendar year
func (Convert) GetSolarYearDays(year int) int {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return 366
	} else {
		return 365
	}
}

// GetSolarMonthDays 获取公历某月天数
//Get the number of days in the Gregorian calendar month
func (Convert) GetSolarMonthDays(year, month int) int {
	var days int
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		days = 29
	} else {
		days = 28
	}
	monthDays := map[int]int{1: 31, 2: days, 3: 31, 4: 30, 5: 31, 6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 31}
	return monthDays[month]
}

// IsLeapYear 判断某年是否是闰年
//Determine whether a year is a leap year
func (Convert) IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	} else {
		return false
	}
}

// GetLeapMonth 获取某年闰月月份
//Get the leap month month of a certain year
func (Convert) GetLeapMonth(year int) int {
	return lunarInfomation[year-1900] & 0xf
}
