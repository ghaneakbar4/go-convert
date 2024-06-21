package convert

import "strconv"

//تبدیل عدد به استرینگ
func ToString(number int) string {
	return strconv.Itoa(number)
}

//تبدیل رشته به عدد
func ToInt(str string) int {
   res,_:=	strconv.Atoi(str);
   return res
}
