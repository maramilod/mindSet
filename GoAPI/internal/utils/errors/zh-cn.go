package errors

var zhCNText = map[int]string{
	SUCCESS:            "تم",
	FAILURE:            "فشل",
	NotFound:           "غير متوفر",
	ServerError:        "خطاء في السيرفر",
	TooManyRequests:    "تداخل الطلبات",
	InvalidParameter:   "ادخال غير متوفر",
	UserDoesNotExist:   "المستخدم غير متوفر",
	AuthorizationError: "خطاء في التوثيق",
	NotLogin:           "لم يتم تسجيل الدخول",
}
