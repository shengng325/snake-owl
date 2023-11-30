package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"snake/dto"
)

func createErrorResponse(w http.ResponseWriter, errCode int, err error) {
	errResp := dto.ErrResp{
		Msg: err.Error(),
	}
	log.Printf("err: %v", err.Error())
	if errCode <= 0 {
		errCode = http.StatusInternalServerError
	}
	createHttpResponse(w, errCode, errResp)
}

func createHttpResponse(w http.ResponseWriter, statusCode int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		// Handle the error in case the response failed to encode
		log.Printf("Error on sending JSON response: %v", err)
	}
}

func getIntQueryParams(queryMap url.Values, paramName string) (int, error) {
	strValue := queryMap.Get(paramName)
	if strValue == "" {
		return 0, fmt.Errorf("query parameter '%v' not found", paramName)
	}
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}
