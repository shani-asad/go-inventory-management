package handler

import (
	"errors"
	"inventory-management/model/dto"
	"inventory-management/src/usecase"
	"log"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	iAuthUsecase usecase.AuthUsecaseInterface
}

func NewAuthHandler(iAuthUsecase usecase.AuthUsecaseInterface) AuthHandlerInterface {
	return &AuthHandler{iAuthUsecase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request dto.RequestCreateStaff
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Register bad request (ShouldBindJSON) >> ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	// Validate request payload
	err = ValidateRegisterRequest(request.PhoneNumber, request.Name, request.Password)
	if err != nil {
		log.Println("Register bad request (ValidateRegisterRequest) >> ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	// Check if phoneNumber already exists
	exists, _ := h.iAuthUsecase.GetStaffByPhoneNumber(request.PhoneNumber)
	if exists {
		log.Println("Register bad request (GetStaffByPhoneNumber) >> ", err)
		c.JSON(409, gin.H{"status": "bad request", "message": "phoneNumber already exists"})
		return
	}

	token, err := h.iAuthUsecase.Register(request)
	if err != nil {
		log.Println("Register bad request (Register) >> ", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	log.Println("Register successful")
	c.JSON(201, gin.H{
		"message": "Staff registered successfully",
		"data": gin.H{
			"phoneNumber":       request.PhoneNumber,
			"name":        request.Name,
			"accessToken": token,
		},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request dto.RequestAuth
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Login bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	err = ValidateLoginRequest(request.PhoneNumber, request.Password)
	if err != nil {
		log.Println("Login bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	token, userData, err := h.iAuthUsecase.Login(request)
	if err != nil {
		log.Println("Login bad request ", err)
		if err.Error() == "staff not found" {
			c.JSON(404, gin.H{"status": "bad request", "message": "staff not found"})
			return
		}
		if err.Error() == "wrong password" {
			c.JSON(400, gin.H{"status": "bad request", "message": "wrong password"})
			return
		} else {
			c.JSON(500, gin.H{"status": "internal server error", "message": err})
			return
		}
	}

	log.Println("Login successful")
	c.JSON(200, gin.H{
		"message": "Staff logged successfully",
		"data": gin.H{
			"userId": userData.Id,
			"phoneNumber": userData.PhoneNumber,
			"name": userData.Name,
			"accessToken": token,
		},
	})
}

// ValidateRegisterRequest validates the register user request payload
func ValidateRegisterRequest(phoneNumber, name, password string) error {
	// Validate phoneNumber format
	if !isValidPhoneNumber(phoneNumber) {
		return errors.New("phoneNumber must be in valid phoneNumber format")
	}

	// Validate name length
	if len(name) < 5 || len(name) > 50 {
		return errors.New("name length must be between 5 and 50 characters")
	}

	// Validate password length
	if len(password) < 5 || len(password) > 15 {
		return errors.New("password length must be between 5 and 15 characters")
	}

	return nil
}

func ValidateLoginRequest(phoneNumber, password string) error {
	// Validate phoneNumber format
	if !isValidPhoneNumber(phoneNumber) {
		return errors.New("phoneNumber must be in valid phoneNumber format")
	}

	// Validate password length
	if len(password) < 5 || len(password) > 15 {
		return errors.New("password length must be between 5 and 15 characters")
	}

	return nil
}

// Helper function to validate phoneNumber format
func isValidPhoneNumber(phoneNumber string) bool {

	if(len(phoneNumber) < 10 || len(phoneNumber) > 16) {return false}

	countryCodes := []string{"93","355","213","1-684","376","244","1-264","672","1-268","54","374","297","61","43","994","1-242","973","880","1-246","375","32","501","229","1-441","975","591","387","267","55","246","1-284","673","359","226","95","257","855","237","1","238","1-345","236","235","56","86","61","61","57","269","242","243","682","506","385","53","599","357","420","45","253","1-767","1-849","670","593","20","503","240","291","372","251","500","298","679","358","33","689","241","220","995","49","233","350","30","299","1-473","1-671","502","44-1481","224","245","592","509","504","852","36","354","91","62","98","964","353","44-1624","972","39","225","1-876","81","44-1534","962","7","254","686","383","965","996","856","371","961","266","231","218","423","370","352","853","389","261","265","60","960","223","356","692","222","230","262","52","691","373","377","976","382","1-664","212","258","264","674","977","31","599","687","64","505","227","234","683","1-670","850","47","968","92","680","970","507","675","595","51","63","64","48","351","1-939","974","262","40","7","250","590","685","378","239","966","221","381","248","232","65","1-721","421","386","677","252","27","82","211","34","94","290","1-869","1-758","590","508","1-784","249","597","47","268","46","41","963","886","992","255","66","228","690","676","1-868","216","90","993","1-649","688","971","256","44","380","598","1","998","678","379","58","84","1-340","681","212","967","260","263"}
	pattern := "^\\+(" + strings.Join(countryCodes, "|") + ")\\d+$"

	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phoneNumber)
}
