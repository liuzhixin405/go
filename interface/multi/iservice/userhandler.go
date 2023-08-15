package iservice

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type UserDataStore interface {
	GetUserScore(ctx context.Context, id string) (int, error)
	DeleteUser(ctx context.Context, id string) error
}

func GetUserScoreHandler(userDataStore UserDataStore) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		id := req.Header.Get("x-user-id")

		score, err := userDataStore.GetUserScore(req.Context(), id)
		if err != nil {
			fmt.Println("userDataStore.GetUserScore: ", err)
			res.WriteHeader(500)
			return
		}
		res.Write([]byte(fmt.Sprintf("%d", score)))
	}
}

func DeleteUserHandler(userDataStore UserDataStore) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Println("io.readall(req.body): ", err)
			res.WriteHeader(500)
			return
		}
		id := string(body)
		err = userDataStore.DeleteUser(req.Context(), id)
		if err != nil {
			fmt.Println("userDataSource.GetUserScore: ", err)
			res.WriteHeader(500)
			return
		}
	}
}
