package student

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"mind-set/internal/controller"
	"mind-set/internal/model"
	"mind-set/internal/utils"
	"mind-set/internal/utils/errors"
	"mind-set/internal/utils/gen_token"
	"mind-set/internal/validator"
	"mind-set/internal/validator/form"
	"os"
	"path/filepath"

	//"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type StudentLoginController struct {
	controller.Api
}

func NewLoginController() *StudentLoginController {
	return &StudentLoginController{}
}

func (api *StudentLoginController) Login(c *gin.Context) {
	StudentLogin := form.StudentLoginForm()
	if err := validator.CheckPostParams(c, &StudentLogin); err != nil {
		return
	}
	user := model.NewStudent()
	result := api.DB().Where("phone = ?", StudentLogin.Phone).First(user)
	if result.Error != nil {
		api.Fail(c, errors.FAILURE, "account Does Not Exist "+result.Error.Error())
		return
	}
	hashPassWord := utils.GetMD5Hash(StudentLogin.PassWord)
	if user.PassWord != hashPassWord {
		api.Fail(c, errors.FAILURE, "user PassWord error")
		return
	}

	claims := api.newStudentClaims(user)
	accessToken, err := gen_token.Generate(claims)
	if err != nil {
		api.Fail(c, errors.FAILURE, "Get token eror")
		return
	}
	var data map[string]interface{}
	inrec, _ := json.Marshal(user)
	json.Unmarshal(inrec, &data)

	data["access_token"] = accessToken
	api.Success(c, data)
}

// register
func (api *StudentLoginController) Register(c *gin.Context) {
	// Initialize the registration form
	StudentRegister := form.StudentRegisterForm()
	// Validate the form input
	if err := validator.CheckPostParams(c, &StudentRegister); err != nil {
		// Return a specific error if validation fails
		api.Fail(c, errors.FAILURE, "Invalid input data")
		return
	}

	// Create a new student model
	user := model.NewStudent()

	// Check if the phone number already exists in the database
	res := api.DB().Where("phone = ?", StudentRegister.Phone).First(user)
	if res.RowsAffected > 0 {
		// If the phone already exists, return a specific message
		api.Fail(c, errors.FAILURE, "Phone number already registered")
		return
	}

	// Check if the email already exists
	resEmail := api.DB().Where("email = ?", StudentRegister.Email).First(user)
	if resEmail.RowsAffected > 0 {
		// If the email already exists, return a specific message
		api.Fail(c, errors.FAILURE, "Email already registered")
		return
	}
	//
	// Hash the password using MD5 (Consider switching to a more secure hashing algorithm like bcrypt)
	hashPassWord := utils.GetMD5Hash(StudentRegister.PassWord)

	// Assign the student details
	user.Name = StudentRegister.Name
	user.Email = StudentRegister.Email
	user.Phone = StudentRegister.Phone
	user.Token = utils.RandString(32)
	user.Gender = StudentRegister.Gender

	/*// Randomly select an avatar image
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rng.Intn(7) + 1
	randomNumberStr := strconv.Itoa(randomNumber)
	user.Avatar = "images/head" + randomNumberStr + ".png"*/

	user.PassWord = string(hashPassWord)

	// Attempt to create the user in the database
	result := api.DB().Create(user)
	if result.Error != nil {
		// Print detailed error to the console for debugging
		fmt.Println("Database Error:", result.Error)
		api.Fail(c, errors.FAILURE, "Registration failed due to internal error")
		return
	}

	// If everything is successful, send a success response
	api.Success(c, nil)
}

//update password
/*
func (api *StudentLoginController) UpdatePassword(c *gin.Context) {

	user, _ := c.Get("student")
	userinfo := user.(gen_token.UserInfo)
	println(userinfo.Id)
	passwordForm := form.StudentUpdatePasswordForm()
	if err := validator.CheckPostParams(c, &passwordForm); err != nil {
		return
	}
	StudentModel := model.NewStudent()
	result := api.DB().Where("id = ?", userinfo.Id).First(StudentModel)
	if result.Error != nil {
		api.Fail(c, errors.FAILURE, "account Does Not Exist "+result.Error.Error())
		return
	}

	hashPassWord := utils.GetMD5Hash(passwordForm.PassWord)
	hashRePassWord := utils.GetMD5Hash(passwordForm.RePassword)
	if StudentModel.PassWord != hashPassWord {
		api.Fail(c, errors.FAILURE, "user current PassWord error")
		return
	}
	results := api.DB().Model(StudentModel).Where("id = ?", userinfo.Id).UpdateColumn("PassWord", hashRePassWord)
	if results.Error != nil {
		api.Fail(c, errors.FAILURE, "PassWord Update failed")
		return
	}
	api.Success(c, nil)

}*/
func (api *StudentLoginController) newStudentClaims(user *model.Student) gen_token.StudClaims {
	now := time.Now()
	expiresAt := now.Add(time.Second * utils.TTL)
	return gen_token.NewWebStudClaims(user, expiresAt)
}
func (api *StudentLoginController) SendCode(c *gin.Context) {
	verificationForm := form.StudentVerificationForm()
	if err := validator.CheckPostParams(c, &verificationForm); err != nil {
		return
	}

	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	todayEnd := todayStart.Add(24 * time.Hour)

	var count int64
	verificationModel := model.NewVerification()
	result2 := api.DB().Model(verificationModel).
		Where("ip = ?", reqIP).
		Where("created_at >= ? AND created_at < ?", todayStart, todayEnd).
		Count(&count)
	if result2.Error != nil {
		api.Fail(c, errors.FAILURE, "Verification code sent error!")
		return
	}
	if count >= 5 {
		api.Fail(c, errors.FAILURE, "Verification code sent too many times!")
		return
	}

	// Generate 6-digit code
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	verification := model.NewVerification()
	verification.Phone = verificationForm.Phone
	verification.Code = vcode
	verification.Ip = reqIP

	result1 := api.DB().Create(verification)
	if result1.Error != nil {
		api.Fail(c, errors.FAILURE, "Verification code generation failed!")
		return
	}

	// Send the verification code is it real ? NOo ^-^
	err := utils.SendSMSCode(verificationForm.Phone, vcode)
	if err != nil {
		api.Fail(c, errors.FAILURE, "Failed to send SMS: "+err.Error())
		return
	}

	// Optional: log the code (useful in dev/testing)
	fmt.Printf("Verification code sent to %s: %s\n", verificationForm.Phone, vcode)

	api.Success(c, gin.H{"message": "Verification code sent successfully"})
}

func (api *StudentLoginController) UploadFile(c *gin.Context) {
	// Step 1: Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		api.Fail(c, errors.FAILURE, "get file error")
		return
	}

	// Step 2: Restrict allowed file extensions (only images and PDFs in this example)
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".pdf":  true,
	}

	// Get the file extension and check against allowed extensions
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExts[ext] {
		api.Fail(c, errors.FAILURE, "invalid file type")
		return
	}

	// Step 3: Sanitize the file name and generate a unique file name
	// Generate a random file name to prevent overwriting and security issues
	currentDate := time.Now().Format("20060102")
	randomNumber := rand.Intn(1000000)
	uniqueFileName := fmt.Sprintf("%d_%d%s", time.Now().UnixNano(), randomNumber, ext)

	// Step 4: Create the target directory if it doesn't exist
	dir := filepath.Join("uploads", currentDate)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		api.Fail(c, errors.FAILURE, "failed to create directory")
		return
	}

	// Step 5: Set the file path and save the file
	filePath := filepath.Join(dir, uniqueFileName)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		api.Fail(c, errors.FAILURE, "save file error")
		return
	}

	// Step 6: Return the URL of the uploaded file
	url := "uploads/" + currentDate + "/" + uniqueFileName
	api.Success(c, url)
}

func (api *StudentLoginController) GetProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		api.Fail(c, errors.FAILURE, "User not authenticated")
		return
	}
	userinfo, ok := user.(gen_token.UserInfo)
	if !ok {
		api.Fail(c, errors.FAILURE, "Invalid user data")
		return
	}

	StudentModel := model.NewStudent()
	StudentModelResult := api.DB().Where("id =?", userinfo.Id).First(StudentModel)
	if StudentModelResult.Error != nil {
		api.Fail(c, errors.FAILURE, "User not found")
		return
	}

	api.Success(c, StudentModel)
}

func (api *StudentLoginController) UpdatePassword(c *gin.Context) {

	user, _ := c.Get("user")
	userinfo := user.(gen_token.UserInfo)
	println(userinfo.Id)
	passwordForm := form.StudentUpdatePasswordForm()
	if err := validator.CheckPostParams(c, &passwordForm); err != nil {
		return
	}
	studentModel := model.NewStudent()
	result := api.DB().Where("id = ?", userinfo.Id).First(studentModel)
	if result.Error != nil {
		api.Fail(c, errors.FAILURE, "account Does Not Exist "+result.Error.Error())
		return
	}

	hashPassWord := utils.GetMD5Hash(passwordForm.Password)
	hashRePassWord := utils.GetMD5Hash(passwordForm.RePassword)
	if studentModel.PassWord != hashPassWord {
		api.Fail(c, errors.FAILURE, "user current password error")
		return
	}
	results := api.DB().Model(studentModel).Where("id = ?", userinfo.Id).UpdateColumn("password", hashRePassWord)
	if results.Error != nil {
		api.Fail(c, errors.FAILURE, "password Update failed")
		return
	}
	api.Success(c, nil)
	return
}

func (api *StudentLoginController) ForgetPassword(c *gin.Context) {

	forgetForm := form.StudentForgetForm()
	if err := validator.CheckPostParams(c, &forgetForm); err != nil {
		return
	}
	studentModel := model.NewStudent()
	result := api.DB().Where("phone = ?", forgetForm.Phone).First(studentModel)
	if result.Error != nil {
		api.Fail(c, errors.FAILURE, "account Does Not Exist "+result.Error.Error())
		return
	}
	verificationModel := model.NewVerification()
	result1 := api.DB().Where("phone = ?", forgetForm.Phone).Order("id desc").First(verificationModel)
	if result1.Error != nil {
		api.Fail(c, errors.FAILURE, "Verification Code Error")
		return
	}
	if verificationModel.Code != forgetForm.VerificationCode {
		api.Fail(c, errors.FAILURE, "Verification Code Error")
		return
	}

	hashPassWord := utils.GetMD5Hash(forgetForm.PassWord)

	results := api.DB().Model(studentModel).Where("phone = ?", forgetForm.Phone).UpdateColumn("password", hashPassWord)
	if results.Error != nil {
		api.Fail(c, errors.FAILURE, "password Update failed")
		return
	}
	api.Success(c, nil)
	return
}

func (api *StudentLoginController) UpdateUserFcmToken(c *gin.Context) {
	user, _ := c.Get("user")
	userinfo := user.(gen_token.UserInfo)
	fcmtokenForm := form.FcmtokenForm()
	if err := validator.CheckPostParams(c, &fcmtokenForm); err != nil {
		return
	}
	StudentModel := model.NewStudent()
	StudentModel.Fcmtoken = fcmtokenForm.Fcmtoken
	StudentModel.UpdatedAt = time.Now()
	result := api.DB().Where("id =?", userinfo.Id).Updates(StudentModel)
	if result.Error != nil {
		api.Fail(c, errors.FAILURE, "Update fcmtoken error")
		return
	}
	api.Success(c, nil)
}
