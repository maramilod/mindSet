package middleware

import (
	"mind-set/internal/model"
	"mind-set/internal/utils"
	e "mind-set/internal/utils/errors"
	"mind-set/internal/utils/gen_token"
	"mind-set/internal/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// MIND SET
func AuthStudentAccessHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		accessToken, err := gen_token.GetAccessToken(authorization)
		if err != nil {
			response.Resp().SetHttpCode(e.NotLogin).FailCode(c, e.NotLogin, err.Error())
			return
		}

		studentClaims := new(gen_token.StudClaims)
		err = gen_token.Parse(accessToken, studentClaims, jwt.WithSubject(utils.SubjectStudent))
		if err != nil || studentClaims == nil {
			response.Resp().SetHttpCode(e.NotLogin).FailCode(c, e.NotLogin, err.Error())
			return
		}

		user := model.NewStudent().GetStudentById(studentClaims.UserInfo.Id)

		c.Set("user", gen_token.UserInfo{
			Id:          user.ID,
			Name:        user.Name,
			Gender:      user.Gender,
			Phone:       user.Phone,
			Email:       user.Email,
			Fcmtoken:    user.Fcmtoken,
			Token:       user.Token,
			AccessToken: utils.GetMD5Hash(accessToken),
		})
		// Set student_id in the context (for cart)
		studentID := studentClaims.UserInfo.Id
		if studentID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Student ID not found in token"})
			c.Abort()
			return
		}

		// Log the student ID
		//log.Printf("Student ID found in token: %d", studentID)

		// Store the student_id in the context
		c.Set("student_id", studentID)

		c.Next()
	}
}
