package routers

import (
	"mind-set/internal/controller/student"
	"mind-set/internal/middleware"
	"mind-set/ws"

	"github.com/gin-gonic/gin"
)

func SetWebStudentRoute(r *gin.Engine) {
	r.GET("/web_student", middleware.AuthStudentAccessHandler(), ws.NewStudentServer)

	v1 := r.Group("webstudent/")
	{

		loginC := student.NewLoginController() // for login and related actions

		// Public routes
		v1.POST("login", loginC.Login)
		v1.POST("register", loginC.Register)
		v1.POST("upload_file", loginC.UploadFile)
		v1.POST("send_code", loginC.SendCode)
		v1.POST("forget_password", loginC.ForgetPassword)

		// Protected routes
		v1.POST("get_profile", middleware.AuthStudentAccessHandler(), loginC.GetProfile)
		v1.POST("update_password", middleware.AuthStudentAccessHandler(), loginC.UpdatePassword)
		v1.POST("bind_fcmtoken", middleware.AuthStudentAccessHandler(), loginC.UpdateUserFcmToken)

	}
}
