package helper

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"penjadwalan-sidang-new/internal/model/web"
	"strconv"
)

const PengajuanURI = "http://localhost:3000"

func GetPengajuanById(pengajuanId int, token string) *web.PengajuanResponseApi {
	client := resty.New()
	result := &web.PengajuanResponseApi{}

	_, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		Get(PengajuanURI + "/api/pengajuan/get/" + strconv.Itoa(pengajuanId))

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func GetPengajuanByPembimbingId(pembimbingId int, token string) *web.PengajuanDatasResponseApi {
	client := resty.New()
	result := &web.PengajuanDatasResponseApi{}

	_, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		Get(PengajuanURI + "/api/pengajuan/pembimbing/get")

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func GetPengajuanKk(token string) *web.PengajuanDatasResponseApi {
	client := resty.New()
	result := &web.PengajuanDatasResponseApi{}

	_, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		Get(PengajuanURI + "/api/pengajuan/kk/get")

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func GetTeamUserLoggedIn(token string) *web.GetMemberUserLoggedInApi {
	client := resty.New()
	result := &web.GetMemberUserLoggedInApi{}

	_, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		Get(PengajuanURI + "/api/team/user/get")

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func GetPengajuanUserLoggedIn(token string) *web.PengajuanResponseApi {
	client := resty.New()
	result := &web.PengajuanResponseApi{}

	_, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		Get(PengajuanURI + "/api/pengajuan/user/get")

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func ChangePengajuanStatus(feedback string, status string, workflowType string, pengajuanId int, token string) *web.PengajuanResponseApi {
	client := resty.New()
	result := &web.PengajuanResponseApi{}

	_, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		SetBody(map[string]interface{}{
			"feedback":      feedback,
			"status":        status,
			"workflow_type": workflowType,
		}).
		Patch(PengajuanURI + "/api/pengajuan/change-status/" + strconv.Itoa(pengajuanId))

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func CreateNotification(memberId []int, title string, message string, url string, token string) {
	client := resty.New()

	_, _ = client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		SetBody(map[string]interface{}{
			"user_id": memberId,
			"title":   title,
			"message": message,
			"url":     url,
		}).
		Post(PengajuanURI + "/api/notification/create")

}
