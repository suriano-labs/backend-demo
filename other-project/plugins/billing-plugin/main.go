package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var pluginName = "billing-plugin"
var HandlerRegisterer = registerer(pluginName)

type registerer string

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(
		context.Context,
		map[string]interface{},
		http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

// func Billing(w http.ResponseWriter, r *http.Request) {

// }

type Bill struct {
	UserID  string `json:"user_id"`
	ApiPath string `json:"api_path"`
}

func (r registerer) registerHandlers(_ context.Context, extra map[string]interface{}, handler http.Handler) (http.Handler, error) {

	// return http.HandlerFunc(Billing), nil

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		// //capturando o Header client_id.
		// jwtToken := strings.TrimPrefix(req.Header.Get("Authorization"), "Bearer ")

		// // Parse the token without validating the signature
		// token, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})
		// if err != nil {
		// 	logger.Debug("Failed to parse token: %v", err)
		// }

		// claims, ok := token.Claims.(jwt.MapClaims)

		// if ok {
		// 	clientId, ok := claims["clientId"].(string)
		// 	if !ok {
		// 		logger.Debug("clientId claim not found or not a string")
		// 		return
		// 	}
		// 	logger.Debug("clientId extracted from Token: ", clientId)
		// 	data := map[string]interface{}{
		// 		"client_id": clientId,
		// 		"path":      req.URL.Path,
		// 	}

		// 	jsonData, err := json.Marshal(data)
		// 	if err != nil {
		// 		logger.Debug("Failed to marshal JSON:", err)
		// 		return
		// 	}

		// 	config, _ := extra["api-monetization"].(map[string]interface{})

		// 	endpoint_checker, _ := config["endpoint_checker"].(string)

		// 	logger.Debug("Endpoint for Checking: %s", endpoint_checker)

		// 	resp, err := http.Post(endpoint_checker, "application/json", bytes.NewBuffer(jsonData))
		// 	if err != nil {
		// 		logger.Debug("Failed to call webhook:", err)
		// 		return
		// 	}

		// 	if resp.StatusCode == 500 {

		// 		w.Header().Add("x-qap-error", "Client ID is blocked, you might had exceeded a free plan")
		// 		w.WriteHeader(http.StatusInternalServerError)

		// 		return

		// 	}

		// 	w.Header().Add("x-qap-client-id", clientId)
		// 	w.Header().Add("x-qap-api-path", req.URL.Path)

		// 	handler.ServeHTTP(w, req)

		// 	endpoint_charge, _ := config["endpoint_charge"].(string)

		// 	//fmt.Println("Endpoint for Charging: %s", endpoint_charge)

		// 	go callWebhook(endpoint_charge, data)

		// } else {
		// 	logger.Debug("Invalid token claims")
		// }

	})

	return handlerFunc, nil
}

func init() {
	logger.Info("💰 Billing: ON")
}

func main() {}

func callWebhook(url string, data map[string]interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Debug(fmt.Sprintf("Failed to marshal JSON:", err))
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Debug(fmt.Sprintf("Failed to call webhook:", err))
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Debug(fmt.Sprintf("Failed to read webhook response:", err))
		return
	}

	logger.Debug(fmt.Sprintf("Endpoint response:", string(body)))
}

var logger Logger = noopLogger{}

func (registerer) RegisterLogger(v interface{}) {
	l, ok := v.(Logger)
	if !ok {
		return
	}
	logger = l
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] 💰 API Monetization Plugin Up and Running ", HandlerRegisterer))
}

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
}

// Empty logger implementation
type noopLogger struct{}

func (n noopLogger) Debug(_ ...interface{})    {}
func (n noopLogger) Info(_ ...interface{})     {}
func (n noopLogger) Warning(_ ...interface{})  {}
func (n noopLogger) Error(_ ...interface{})    {}
func (n noopLogger) Critical(_ ...interface{}) {}
func (n noopLogger) Fatal(_ ...interface{})    {}
