package student

/*
import (
	"fmt"
	"mind-set/internal/controller"
	"mind-set/internal/model"
	"mind-set/internal/service"
	"mind-set/internal/utils"
	"mind-set/internal/utils/errors"
	"mind-set/internal/utils/gen_token"
	"mind-set/internal/validator"
	"mind-set/internal/validator/form"
	"mind-set/ws"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	controller.Api
}

func NewChatController() *ChatController {
	return &ChatController{}
}

// 聊天消息列表
func (api *ChatController) ChatList(c *gin.Context) {
	user, _ := c.Get("user")
	userinfo := user.(gen_token.UserInfo)
	toTokenForm := form.ToTokenForm()
	if err := validator.CheckPostParams(c, &toTokenForm); err != nil {
		return
	}
	datas := []map[string]interface{}{}
	chatModel := model.NewChat()
	result := api.DB().Model(chatModel).
		Where("token = ? AND to_token = ?", userinfo.Token, toTokenForm.ToToken).
		Or("token = ? AND to_token = ?", toTokenForm.ToToken, userinfo.Token).
		Find(&datas)
	if result.Error != nil {
		api.Fail(c, errors.FAILURE, "message Not Exist ")
		return
	}
	api.Success(c, datas)
}

// 发送消息
func (api *ChatController) SendMessage(c *gin.Context) {
	user, _ := c.Get("user")
	userinfo := user.(gen_token.UserInfo)
	chatForm := form.ChatForm()
	if err := validator.CheckPostParams(c, &chatForm); err != nil {
		return
	}
	toCid := fmt.Sprintf("%d", *chatForm.ToCid)
	fmt.Println("student SendMessage---ToCid--", toCid)
	chatModel := model.NewChat()
	chatModel.Cid = userinfo.Cid
	chatModel.Name = userinfo.FirstName
	chatModel.Token = userinfo.Token
	chatModel.Avatar = userinfo.Avatar
	chatModel.Content = chatForm.Content
	chatModel.ToToken = chatForm.ToToken
	chatModel.ToCid = *chatForm.ToCid
	chatModel.CreatedAt = time.Now()
	chatModel.UpdatedAt = time.Now()
	result := api.DB().Create(chatModel)
	if result.Error != nil {
		api.Fail(c, errors.FAILURE, "add message error ")
		return
	}

	hashStr := chatForm.ToToken + toCid
	fmt.Println("student SendMessage---hashStr--", hashStr)
	hashAccessToken := utils.GetMD5Hash(hashStr)
	fmt.Println("student SendMessage-----", hashAccessToken)
	ws.SendMessage(hashAccessToken, chatModel)
	//发送通知
	title := userinfo.FirstName + " send a message"
	body := chatForm.Content
	service.NewMessageService().SendMessage(chatForm.ToToken, title, body)

	api.Success(c, chatModel)
}

func (api *ChatController) SalePointList(c *gin.Context) {
	datas := []map[string]interface{}{}
	salePointModel := model.NewSalePoint()
	result := api.DB().Model(salePointModel).Select("id", "cid", "token", "avatar", "first_name", "middle_name", "last_name", "phone").Find(&datas)
	if result.Error != nil {
		api.Fail(c, errors.FAILURE, "SalePoint Not Exist "+result.Error.Error())
		return
	}
	api.Success(c, datas)
}

func (api *ChatController) AgentList(c *gin.Context) {
	datas := []map[string]interface{}{}
	agentModel := model.NewAgent()
	result := api.DB().Model(agentModel).Select("id", "cid", "token", "avatar", "first_name", "middle_name", "last_name", "phone").Find(&datas)
	if result.Error != nil {
		api.Fail(c, errors.FAILURE, "Agent Not Exist "+result.Error.Error())
		return
	}
	api.Success(c, datas)
}

func (api *ChatController) UpdateUserFcmToken(c *gin.Context) {
	user, _ := c.Get("user")
	userinfo := user.(gen_token.UserInfo)
	fcmtokenForm := form.FcmtokenForm()
	if err := validator.CheckPostParams(c, &fcmtokenForm); err != nil {
		return
	}
	adminUserModel := model.NewAdminUser()
	adminUserModel.Fcmtoken = fcmtokenForm.Fcmtoken
	adminUserModel.UpdatedAt = time.Now()
	result := api.DB().Where("id =?", userinfo.Id).Updates(adminUserModel)
	if result.Error != nil {
		api.Fail(c, errors.FAILURE, "Update fcmtoken error")
		return
	}
	api.Success(c, nil)
}
*/
