package service

import (
	"ojam-test/c3/s1/app/generator"
)

func PostStatus(lastId int) int {
	maxUserId := 100
	maxRandomPostId := 25
	userId := generator.RandomNumber(maxUserId)
	randomPostId := generator.RandomNumber(maxRandomPostId)
	lastPostId := generator.Post(userId, randomPostId, lastId)
	return lastPostId
}