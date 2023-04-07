package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ojam-test/c3/s1/app/constant"
	"ojam-test/c3/s1/app/model"
)

func Post(userId, randomPostId, lastId int) int {
	var firstTitle, lastTitle, firstBody, optionalBody, lastBody string
	if userId > 90 {
		firstTitle = "amor spero"
		lastBody = "amor est facillimum"
	} else if userId > 80 {
		firstTitle = "defectum"
		lastBody = "calamus igniti falcatus"
	} else if userId > 70 {
		firstTitle = "dominus gustum"
		lastBody = "romance in aperto universum amoris"
	} else if userId > 60 {
		firstTitle = "vadum libidine"
		lastBody = "spe comeditur acerbitate vitae"
	} else if userId > 50 {
		firstTitle = "sugar tata"
		lastBody = "ursi pulmonis sanguinem movet hominem"
	} else if userId > 40 {
		firstTitle = "paenitet"
		lastBody = "cor regis in saltu vasculum dolet"
	} else if userId > 30 {
		firstTitle = "dens eruit"
		lastBody = "vexas in tenebris"
	} else if userId > 20 {
		firstTitle = "promisit"
		lastBody = "ardenti tempore flammas aquae sculpit fabulas"
	} else if userId > 10 {
		firstTitle = "et consolidationis"
		lastBody = "sacram caudam mirantis viduae"
	} else {
		firstTitle = "puppy amor"
		lastBody = "amor fabula in teenage ludum in lecto libidinis"
	} 

	if randomPostId > 20 {
		lastTitle = "masculum passionem"
		optionalBody = "virilitas virilis est in semine producendo"
	} else if randomPostId > 15 {
		lastTitle = "confractus felis"
		optionalBody = "incidit in canalem penis consilio"
	} else if randomPostId > 10 {
		lastTitle = "in tace"
		optionalBody = "intellige regem qui in aetemitate vastavit"
	} else if randomPostId > 5 {
		lastTitle = "lapis te ipsum"
		optionalBody = "amantes se non cognoscunt"
	} else {
		lastTitle = "et mori"
		optionalBody = "memorans vitam, quae numquam durat"
	}

	if (userId > 85 && randomPostId < 3) {
		firstBody = "multum libidinis hospitio principis"
	} else if (userId > 70 && randomPostId < 6) {
		firstBody = "dator victoria ad dominum"
	} else if (userId > 55 && randomPostId < 9) {
		firstBody = "nisi gay vidui sunt aduncum"
	} else if (userId > 40 && randomPostId < 12)  {
		firstBody = "consumpta est in amore pueri"
	} else if (userId > 25 && randomPostId < 15) {
		firstBody = "faciens in furore amoris turbulenti"
	} else if (userId > 10 && randomPostId < 18) {
		firstBody = "vile bellum viduarum et vidui"
	} else if (userId > 85 && randomPostId > 18) {
		firstBody = "de fine fabulae poetae in genitalibus"
	} else if (userId > 70 && randomPostId > 15) {
		firstBody = "sterilitas hominis semen vitae"
	} else if (userId > 55 && randomPostId > 12) {
		firstBody = "spes amoris suilla tempore"
	} else {
		firstBody = "homo enim qui vult ut simul"
	}

	postTitle := fmt.Sprintf("%s %s", firstTitle, lastTitle)
	postBody := fmt.Sprintf("%s %s %s", firstBody, optionalBody, lastBody)

	dataRequest := model.PostStatus{
		UserId: userId,
		Title: postTitle,
		Body: postBody,
	}

	bodyRequest, err := json.Marshal(dataRequest)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest(http.MethodPost, constant.POST_PATH, bytes.NewBuffer(bodyRequest))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	bodyResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// For Print Actual Response Post, the ID always 101
	// fmt.Println(string(bodyResponse))

	dataResponse := model.PostStatus{}
	json.Unmarshal([]byte(bodyResponse), &dataResponse)

	specificDataResponse := model.PostStatus{
		UserId: dataResponse.UserId,
		Id: lastId,
		Title: dataResponse.Title,
		Body: dataResponse.Body,
	}

	body, err := json.MarshalIndent(specificDataResponse, "", "   ")
	if err != nil {
		panic(err)
	}

	// For Print Custom Response that handle the ID
	fmt.Println(string(body))

	return lastId
}