// handler.go

package handler

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"inventory-management/model/dto"
	"inventory-management/src/usecase"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
    iCustomerUsecase usecase.CustomerUsecaseInterface
}

func NewCustomerHandler(iCustomerUsecase usecase.CustomerUsecaseInterface) CustomerHandlerInterface {
    return &CustomerHandler{iCustomerUsecase}
}

func (h *CustomerHandler) RegisterCustomer(c *gin.Context) {
		var request dto.RegisterCustomerRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Println("register customer bad request ", err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "bad request", "message": err})
			return
		}

		if !validateCustomer(&request) {
			log.Println("register customer bad request")
			c.JSON(http.StatusBadRequest, gin.H{"error": "request doesn’t pass validation"})
			return
		}

    userID, err := h.iCustomerUsecase.RegisterCustomer(request)
    if err != nil {
				log.Println("register customer internal server error:", err)
				c.JSON(500, gin.H{"status": "internal server error", "message": err})
        return
    }

    log.Println("register customer success")
		// Mocking the response
		response := gin.H{
			"message": "success",
			"data": gin.H{
				"userId":        userID,
				"phoneNumber": request.PhoneNumber,
				"name": request.Name,
			},
		}

		c.JSON(http.StatusCreated, response)
}

func (h *CustomerHandler) SearchCustomers(c *gin.Context) {
    // Parse query parameters
    // phoneNumber := r.URL.Query().Get("phoneNumber")
    // name := r.URL.Query().Get("name")

    // customers, err := h.iCustomerUsecase.SearchCustomers(r.Context(), phoneNumber, name)
    // if err != nil {
    //     // Handle error from usecase
    //     http.Error(w, err.Error(), http.StatusInternalServerError)
    //     return
    // }

    // resp := dto.SearchCustomersResponse{
    //     Message: "success",
    //     Data:    customers,
    // }

    // w.WriteHeader(http.StatusOK)
    // json.NewEncoder(w).Encode(resp)
}

func validateCustomer(customer *dto.RegisterCustomerRequest) bool {
	// Validate phoneNumber format
	if !isValidPhoneNumber(customer.PhoneNumber) {
		return false
	}

	// Validate name length
	if len(customer.Name) < 5 || len(customer.Name) > 50 {
		return false
	}

	return true
}

func isValidPhoneNumber(phoneNumber string) bool {

	if(len(phoneNumber) < 10 || len(phoneNumber) > 16) {return false}

	countryCodes := []string{"93","355","213","1-684","376","244","1-264","672","1-268","54","374","297","61","43","994","1-242","973","880","1-246","375","32","501","229","1-441","975","591","387","267","55","246","1-284","673","359","226","95","257","855","237","1","238","1-345","236","235","56","86","61","61","57","269","242","243","682","506","385","53","599","357","420","45","253","1-767","1-849","670","593","20","503","240","291","372","251","500","298","679","358","33","689","241","220","995","49","233","350","30","299","1-473","1-671","502","44-1481","224","245","592","509","504","852","36","354","91","62","98","964","353","44-1624","972","39","225","1-876","81","44-1534","962","7","254","686","383","965","996","856","371","961","266","231","218","423","370","352","853","389","261","265","60","960","223","356","692","222","230","262","52","691","373","377","976","382","1-664","212","258","264","674","977","31","599","687","64","505","227","234","683","1-670","850","47","968","92","680","970","507","675","595","51","63","64","48","351","1-939","974","262","40","7","250","590","685","378","239","966","221","381","248","232","65","1-721","421","386","677","252","27","82","211","34","94","290","1-869","1-758","590","508","1-784","249","597","47","268","46","41","963","886","992","255","66","228","690","676","1-868","216","90","993","1-649","688","971","256","44","380","598","1","998","678","379","58","84","1-340","681","212","967","260","263"}
	pattern := "^\\+(" + strings.Join(countryCodes, "|") + ")\\d+$"

	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phoneNumber)
}

