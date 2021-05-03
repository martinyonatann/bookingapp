package logging

import (
	"fmt"

	"github.com/martinyonathann/bookingapp/helper"
	log "github.com/sirupsen/logrus"
)

func SuccessLogging(uriReq string, DataResponse interface{}) {
	log.SetFormatter(&log.JSONFormatter{})
	fmt.Print(helper.DateTime("2006-01-02 15:04:05") + " " + uriReq + " ")
	log.WithFields(log.Fields{"response": DataResponse}).Info("Process Successfuly")
}

func ErrorLogging(responseCode int, uriReq string, DataResponse interface{}) {
	log.SetFormatter(&log.JSONFormatter{})
	standardFields := log.Fields{
		"code":   responseCode,
		"detail": "Failed",
	}
	fmt.Print(helper.DateTime("2006-01-02 15:04:05") + " " + uriReq + " ")
	log.WithFields(log.Fields{"errorResult": DataResponse}).WithFields(standardFields).Error("Process Failed")
}
