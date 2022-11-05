package utils

import (
	"encoding/json"
	"go-authentication/dtos"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

func FileWithLineNum() string {
	// the second caller usually from gorm internal, so set i start from 2
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.HasSuffix(file, "_test.go")) {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}

	return ""
}

func WriteResponse(resWriter http.ResponseWriter, response dtos.Response) error {
	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(response.StatusCode)
	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}
	resWriter.Write(marshaledResponse)
	return nil
}
