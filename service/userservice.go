package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetUserList
// @Summary 用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"code": 0, //0成功	-1失败
		"data": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/creatUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	//因为前端是表单的格式因此不能使用query
	//user.Name = c.Query("name")
	//password := c.Query("password")
	//repassword := c.Query("repassword")

	user.Name = c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("repassword")

	salt := fmt.Sprintf("%06d", rand.Int31())

	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"code":    -1, //0成功	-1失败
			"message": "用户名已被占用！",
		})
		return
	}

	if password != repassword {
		c.JSON(-1, gin.H{
			"code":    -1, //0成功	-1失败
			"message": "两次密码不一致",
		})
		return
	} else if len(password) < 6 {
		c.JSON(-1, gin.H{
			"code":    -1, //0成功	-1失败
			"message": "密码长度需要大于等于6位",
		})
		return
	}

	//user.PassWord = password
	//加密
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreatUser(user)

	c.JSON(http.StatusOK, gin.H{
		"code":    0, //0成功	-1失败
		"message": "新增用户成功",
	})

}

// FindUserByNameAndPwd
// @Summary 登录
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}

	//因为前端是表单的格式因此不能使用query
	//name := c.Query("name")
	//password := c.Query("password")
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	fmt.Println("....")
	fmt.Println(name, password)
	fmt.Println("....")
	//先判断用户存不存在
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //0成功	-1失败
			"message": "该用户不存在",
		})
		return
	}

	//解密
	//添加功能，添加校验次数，超过多久之后提醒冻结
	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //0成功	-1失败
			"message": "密码不正确",
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)

	c.JSON(http.StatusOK, gin.H{
		"code":    0, //0成功	-1失败
		"message": "登录成功",
		"data":    data,
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0, //0成功	-1失败
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Summary 修改个人信息
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //0成功	-1失败
			"message": "修改参数不匹配成功",
		})
		return
	}

	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"code":    -1, //0成功	-1失败
			"message": "用户名已被占用！",
		})
		return
	}

	phone := models.FindUserByPhone(user.Phone)
	if phone.Phone != "" {
		c.JSON(-1, gin.H{
			"code":    -1, //0成功	-1失败
			"message": "电话号码已被绑定！",
		})
		return
	}

	email := models.FindUserByEmail(user.Email)
	if email.Email != "" {
		c.JSON(-1, gin.H{
			"code":    -1, //0成功	-1失败
			"message": "邮箱已被绑定！",
		})
		return
	}

	if len(user.PassWord) < 6 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //0成功	-1失败
			"message": "密码长度需要大于等于6",
		})
		return
	}

	salt := fmt.Sprintf("%06d", rand.Int31())
	user.PassWord = utils.MakePassword(user.PassWord, salt)
	models.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0, //0成功	-1失败
		"message": "修改用户成功",
	})
}

// 防止跨域站点的伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(ws)
	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println(err)
		}

		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}

func SearchFriends(c *gin.Context) {
	id, _ := strconv.Atoi(c.Request.FormValue("userId"))
	users := models.SearchFriend(uint(id))
	//c.JSON(200, gin.H{
	//	"code":    0,
	//	"message": "success",
	//	"data":    users,
	//})

	utils.RespOKList(c.Writer, users, len(users))
}

func AddFriend(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
	targetId, _ := strconv.Atoi(c.Request.FormValue("targetId"))
	code, msg := models.AddFriend(uint(userId), uint(targetId))
	// c.JSON(200, gin.H{
	// "code": 0, // 0成功 -1失败
	// "message": "查询好友列表成功！",
	// "data": users,
	// })
	if code == 0 {
		utils.RespOK(c.Writer, code, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}

func CreateCommunity(c *gin.Context) {
	ownerId, _ := strconv.Atoi(c.Request.FormValue("ownerId"))
	name := c.Request.FormValue("name")
	community := models.Community{}
	community.OwnerId = uint(ownerId)
	community.Name = name
	code, msg := models.CreateCommunity(community)

	// c.JSON(200, gin.H{
	// "code": 0, // 0成功 -1失败
	// "message": "查询好友列表成功！",
	// "data": users,
	// })
	if code == 0 {
		utils.RespOK(c.Writer, code, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}

// 加载群列表
func LoadCommunity(c *gin.Context) {
	ownerId, _ := strconv.Atoi(c.Request.FormValue("ownerId"))

	data, msg := models.LoadCommunity(uint(ownerId))
	if len(data) != 0 {
		utils.RespList(c.Writer, 0, data, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}

//func JoinGroup(c *gin.Context) {
//	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
//	targetId, _ := strconv.Atoi(c.Request.FormValue("targetId"))
//	code, msg := models.JoinGroup(uint(userId), uint(targetId))
//	// c.JSON(200, gin.H{
//	// "code": 0, // 0成功 -1失败
//	// "message": "查询好友列表成功！",
//	// "data": users,
//	// })
//	if code == 0 {
//		utils.RespOK(c.Writer, code, msg)
//	} else {
//		utils.RespFail(c.Writer, msg)
//	}
//}
