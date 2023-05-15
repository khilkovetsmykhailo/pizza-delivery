package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/marijakljestan/golang-web-app/client/dto"
	"github.com/marijakljestan/golang-web-app/client/store"
	"net/http"
	"os"
	"strconv"
)

var baseUrl = os.Getenv("SERVER_URL")

func RegisterUser() {
	var username string
	fmt.Println("Enter username:")
	fmt.Scan(&username)

	var password string
	fmt.Println("Enter password:")
	fmt.Scan(&password)

	url := baseUrl + "/user/register"
	reqBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		handleErrorResponse(resp)
		return
	}

	var registeredUser dto.User
	err = json.NewDecoder(resp.Body).Decode(&registeredUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully registered with username:", registeredUser.Username)
}

func Login() bool {
	var username string
	fmt.Println("Enter username:")
	fmt.Scan(&username)

	var password string
	fmt.Println("Enter password:")
	fmt.Scan(&password)

	url := baseUrl + "/user/login"
	reqBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		fmt.Println(err)
		return false
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		handleErrorResponse(resp)
		return false
	}

	var respBody dto.LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		fmt.Println(err)
		return false
	}

	store.SetLoggedUser(username, respBody.Token)
	fmt.Println("Your authorization token:", respBody.Token)
	return true
}

func LogOut() {
	store.SetLoggedUser("", "")
}

func ListMenu() {
	url := baseUrl + "/pizza"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var menu dto.MenuResponse
	err = json.NewDecoder(resp.Body).Decode(&menu)
	if err != nil {
		fmt.Println(err)
		return
	}
	formatAndPrintResponse(menu)
}

func CreateOrder() {
	var items []dto.OrderItem
	var pizzaName string
	var quantity int
	var orderItem dto.OrderItem
	fmt.Println("-------- Add items to your order ---------")
Loop:
	for {
		fmt.Println("--- Press any key to continue --")
		fmt.Println("--If you want to finish ordering press 0--")
		var choice string
		fmt.Scan(&choice)
		switch choice {
		case "0":
			fmt.Println("Finishing order...")
			break Loop
		default:
			fmt.Println("Pizza name: ")
			fmt.Scan(&pizzaName)

			fmt.Println("Quantity: ")
			fmt.Scan(&quantity)

			orderItem = dto.OrderItem{
				PizzaName: pizzaName,
				Quantity:  quantity,
			}
			items = append(items, orderItem)
		}
	}

	reqBody, err := json.Marshal(map[string]any{
		"items": items,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	bearerToken := getAuthorizationToken()
	url := baseUrl + "/order"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		handleErrorResponse(resp)
		return
	}

	var createOrderResponse dto.CreateOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&createOrderResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully created order!")
	formatAndPrintResponse(createOrderResponse)
}

func CheckOrderStatus() {
	var orderId string
	fmt.Println("Enter order id:")
	fmt.Scan(&orderId)

	url := fmt.Sprintf(baseUrl+"/order/status/%s", orderId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		handleErrorResponse(resp)
		return
	}

	var orderStatusResponse dto.GetOrderStatusResponse
	err = json.NewDecoder(resp.Body).Decode(&orderStatusResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Your order has status:", orderStatusResponse.OrderStatus)
}

func CancelOrder() {
	var orderId string
	fmt.Println("Enter order id:")
	fmt.Scan(&orderId)

	url := fmt.Sprintf(baseUrl+"/order/cancel/%s", orderId)
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		handleErrorResponse(resp)
		return
	}

	var cancelOrderResponse dto.CancelOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&cancelOrderResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully cancelled order")
	formatAndPrintResponse(cancelOrderResponse)
}

func AddPizza() {
	var name string
	fmt.Println("Enter name:")
	fmt.Scan(&name)

	var description string
	fmt.Println("Enter description:")
	fmt.Scan(&description)

	var price string
	fmt.Println("Enter price:")
	fmt.Scan(&price)

	priceConverted, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Error! Expected numeric input!")
		return
	}

	bearerToken := getAuthorizationToken()
	url := baseUrl + "/pizza"

	var pizza = dto.Pizza{
		Name:        name,
		Description: description,
		Price:       priceConverted,
	}

	fmt.Println("Price", price)
	reqBody, err := json.Marshal(pizza)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		handleErrorResponse(resp)
		return
	}

	var respBody dto.MenuResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	formatAndPrintResponse(respBody)
}

func DeletePizza() {
	var pizzaName string
	fmt.Println("Enter name of pizza you want to delete:")
	fmt.Scan(&pizzaName)

	bearerToken := getAuthorizationToken()
	url := fmt.Sprintf(baseUrl+"/pizza/%s", pizzaName)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Add("Authorization", bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		handleErrorResponse(resp)
		return
	}

	var menu dto.MenuResponse
	err = json.NewDecoder(resp.Body).Decode(&menu)
	if err != nil {
		fmt.Println(err)
		return
	}

	formatAndPrintResponse(menu)
}

func CancelOrderRegardlessStatus() {
	var orderId string
	fmt.Println("Enter order id:")
	fmt.Scan(&orderId)

	bearerToken := getAuthorizationToken()
	url := fmt.Sprintf(baseUrl+"/order/%s", orderId)

	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		handleErrorResponse(resp)
		return
	}

	var cancelOrderResponse dto.CancelOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&cancelOrderResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully cancelled order")
	formatAndPrintResponse(cancelOrderResponse)
}

func getAuthorizationToken() string {
	storedUser := store.GetLoggedUser()
	bearerToken := "Bearer " + (*storedUser).Token
	return bearerToken
}

func formatAndPrintResponse(response any) {
	jsonData, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func handleErrorResponse(resp *http.Response) {
	var responseError dto.ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&responseError)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println(responseError.Error)
	return
}

func main() {
	fmt.Println("-----------------------------------------")
	fmt.Println("-------- Library Management CLI ---------")
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("----------- Choose Operation ------------")
		fmt.Println("0) EXIT")
		fmt.Println("1) Register as a customer")
		fmt.Println("2) Login as a customer")
		fmt.Println("3) Login as an admin")
		var flag int
		flag = handlePrompets()
		if flag == 0 {
			break
		}
	}
}

func handlePrompets() int {
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		RegisterUser()
	case "2":
		if successfulLogin := Login(); successfulLogin {
			handleCustomerChoice()
		}
	case "3":
		if successfulLogin := Login(); successfulLogin {
			handleAdminChoice()
		}
	case "0":
		return 0
	}
	return 1
}

func handleCustomerChoice() {
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("----------- Choose Operation ------------")
		fmt.Println("-----------------------------------------")
		fmt.Println("1) List the menu")
		fmt.Println("2) Create an order")
		fmt.Println("3) Check order status")
		fmt.Println("4) Cancel order")
		fmt.Println("0) Log out")
		var choice string
		fmt.Scan(&choice)
		fmt.Println("-----------------------------------------")
		switch choice {
		case "1":
			ListMenu()
		case "2":
			CreateOrder()
		case "3":
			CheckOrderStatus()
		case "4":
			CancelOrder()
		case "0":
			LogOut()
			return
		default:
			fmt.Println("INPUT ERROR !")
		}
	}
}

func handleAdminChoice() {
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("----------- Choose Operation ------------")
		fmt.Println("-----------------------------------------")
		fmt.Println("1) List the menu")
		fmt.Println("2) Add pizza")
		fmt.Println("3) Delete pizza")
		fmt.Println("4) Cancel order")
		fmt.Println("0) Log out")
		var choice string
		fmt.Scan(&choice)
		fmt.Println("-----------------------------------------")
		switch choice {
		case "1":
			ListMenu()
		case "2":
			AddPizza()
		case "3":
			DeletePizza()
		case "4":
			CancelOrderRegardlessStatus()
		case "0":
			LogOut()
			return
		default:
			fmt.Println("INPUT ERROR !")
		}
	}
}
