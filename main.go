package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/imroc/req"
	"github.com/pkg/errors"
)

const notificationURL = "https://6098f32599011f001713fc6e.mockapi.io"

// Check if aa user request is authenticated
func authFunction(r *http.Request) bool {
	return r.URL.Query().Get("auth") == "ABC"
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !authFunction(r) {
			http.Redirect(w, r, "/notloggedIn", http.StatusUnauthorized)
			return
		}

		// User authenticated successfully
		next(w, r)
	}
}

type requestCtx struct {
	Name  string
	Age   int
	Email string
}

func buildRequestContextAndValidate(r *http.Request) (requestCtx, error) {
	ageStr := r.URL.Query().Get("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return requestCtx{}, errors.Wrap(err, "age is not an integer")
	}
	if age < 18 {
		return requestCtx{}, errors.New("user is below 18 yr")
	}

	return requestCtx{
		Name:  r.URL.Query().Get("name"),
		Age:   age,
		Email: r.URL.Query().Get("email"),
	}, nil
}

func insertRequestToDB(ctx context.Context, reqCtx requestCtx) (int, error) {
  // insert record into the db
	return 1, nil
}

func sendNotification(ctx context.Context, reqID int, reqCtx requestCtx, baseURL string) error {
	res, err := req.Post(baseURL+"/mail", req.BodyJSON(req.Param{
		"name":   reqCtx.Name,
		"email":  reqCtx.Email,
		"req_id": reqID,
	}))

	responseMap := map[string]interface{}{}
	res.ToJSON(&responseMap)

	if len(responseMap["email"].(string)) == 0 {
		return errors.New("invalid response from notification provider")
	}
	return err
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqCtx, err := buildRequestContextAndValidate(r)
	if err != nil {
		fmt.Fprintf(w, "There ware a validation error; %v", err)
		return
	}

	reqID, err := insertRequestToDB(ctx, reqCtx)
	if err != nil {
		fmt.Fprintf(w, "There ware an error inserting request; %v", err)
		return
	}

	err = sendNotification(ctx, reqID, reqCtx, notificationURL)
	if err != nil {
		fmt.Fprintf(w, "There was an error sending notification; %v", err)
		return
	}

	fmt.Fprintf(w, "Success in Handling  request")
}

func main() {
	http.HandleFunc("/submit", authMiddleware(handlerFunc))
	http.HandleFunc("/notloggedIn", notLoggedInHandlerFunc)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func notLoggedInHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not logged in"))
}
