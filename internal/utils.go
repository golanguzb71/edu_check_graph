package utils

import "edu_test_graph/graph/model"

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
