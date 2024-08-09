package utils

import (
	"context"
	"edu_test_graph/graph/model"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func AbsResponseChecking(err error, msg string) (*model.Response, error) {
	if err != nil {
		return &model.Response{
			StatusCode: 409,
			Message:    err.Error(),
		}, nil
	}
	return &model.Response{
		StatusCode: 200,
		Message:    msg,
	}, nil
}

type Response struct {
	UserID int `json:"user_id"`
	Code   int `json:"code"`
}

func SearchByValue(rdb *redis.Client, targetCode int) (Response, error) {
	ctx := context.TODO()
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		return Response{}, err
	}

	for _, key := range keys {
		val, err := rdb.Get(ctx, key).Result()
		if err != nil {
			return Response{}, err
		}

		var res Response
		err = json.Unmarshal([]byte(val), &res)
		if err != nil {
			return Response{}, err
		}

		if res.Code == targetCode {
			return res, nil
		}
	}

	return Response{}, fmt.Errorf("no match found")
}
