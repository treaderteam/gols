package util

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"gitlab.com/alexnikita/translator/trapi"

	"golang.org/x/net/context"

	"gitlab.com/alexnikita/gols/config"

	accounter "gitlab.com/alexnikita/accounter/api"
	bookkeeper "gitlab.com/alexnikita/bookkeeper/api"
	logserver "gitlab.com/alexnikita/logserver/lapi"
	userrer "gitlab.com/alexnikita/userrer/api"
	"google.golang.org/grpc"
)

var (
	// LogserverConnection is permamnent tcp connection to
	// logserver
	LogserverConnection logserver.LogserverClient
)

func init() {
	logserverAddress := config.Request("LOGSERVER_ADDRESS", false)
	if logserverAddress == "" {
		logserverAddress = ":55558"
	}
	c, _, err := ConnectToLogserver(logserverAddress)
	if err != nil {
		log.Panic(err)
	}
	LogserverConnection = c
}

// CORSPOSTHandler CORS POST
func CORSPOSTHandler(rw http.ResponseWriter) int {
	rw.Header().Add("Access-Control-Allow-Method", "POST")
	rw.Header().Add("Access-Control-Allow-Headers", "content-type, user-id")
	return 204
}

// ConnectToBookkeeper func
func ConnectToBookkeeper(address string) (c bookkeeper.BookkeeperClient, conn *grpc.ClientConn, err error) {
	log.Printf("dialing %s\n", address)
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	defer func(e *error) {
		if *e != nil {
			log.Println(*e)
		}
	}(&err)
	if err != nil {
		return
	}

	c = bookkeeper.NewBookkeeperClient(conn)
	return
}

// ConnectToUserrer func
func ConnectToUserrer(address string) (c userrer.UserrerClient, conn *grpc.ClientConn, err error) {
	log.Printf("dialing %s\n", address)
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	defer func(e *error) {
		if *e != nil {
			log.Println(*e)
		}
	}(&err)
	if err != nil {
		return
	}

	c = userrer.NewUserrerClient(conn)
	return
}

// ConnectToLogserver func
func ConnectToLogserver(address string) (c logserver.LogserverClient, conn *grpc.ClientConn, err error) {
	log.Printf("dialing %s\n", address)
	conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithMaxMsgSize(10000000))
	defer func(e *error) {
		if *e != nil {
			log.Println(*e)
		}
	}(&err)
	if err != nil {
		return
	}

	c = logserver.NewLogserverClient(conn)
	return
}

// ConnectToAccounter func
func ConnectToAccounter(address string) (c accounter.AccounterClient, conn *grpc.ClientConn, err error) {
	log.Printf("dialing %s\n", address)
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	defer func(e *error) {
		if *e != nil {
			log.Println(*e)
		}
	}(&err)
	if err != nil {
		return
	}

	c = accounter.NewAccounterClient(conn)
	return
}

// ConnectToTranslator func
func ConnectToTranslator(address string) (c trapi.TranslatorClient, conn *grpc.ClientConn, err error) {
	log.Printf("dialing %s\n", address)
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	defer func(e *error) {
		if *e != nil {
			log.Println(*e)
		}
	}(&err)
	if err != nil {
		return
	}

	c = trapi.NewTranslatorClient(conn)
	return
}

// ReqLogger logger, should be deferred
func ReqLogger(rw *http.ResponseWriter, responseStatus *int, URL *url.URL, Method string, start *time.Time) {
	if *responseStatus != 200 && *responseStatus != 308 {
		(*rw).WriteHeader(*responseStatus)
	}
	log.Printf("%s %s %d %s\n", Method, URL, *responseStatus, time.Since(*start))
}

// ErrLogger should be deferred
func ErrLogger(e *error) {
	if err := recover(); err != nil {
		log.Println(err)
		LogserverConnection.SendError(
			context.Background(),
			&logserver.ErrorMessage{From: "MS-API", Message: fmt.Sprintln(err)})
	}
	if (*e) != nil {
		log.Println(*e)
		LogserverConnection.SendError(
			context.Background(),
			&logserver.ErrorMessage{From: "MS-API", Message: (*e).Error()})
	}
}
