package helper

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"pendaftaran-sidang-new/internal/exception"
	"pendaftaran-sidang-new/internal/model/web"
	"strconv"
)

var apiUrl string = "https://sofi.my.id"

func GetDetailStudent(studentId int) (*web.GetDetailStudentResponse, error) {
	client := resty.New()
	result := &web.GetDetailStudentResponse{}

	res, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		Get(apiUrl + "/api/student/" + strconv.Itoa(studentId))

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 || res.IsError() {
		return nil, errors.New("data not found")
	}

	return result, nil
}

func UpdatePeminatanByUserId(peminatanId int, userId int) (*web.GetDetailStudentResponse, error) {
	client := resty.New()
	result := &web.GetDetailStudentResponse{}

	res, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"peminatan_id": peminatanId,
		}).Patch(apiUrl + "/api/sidang/update/" + strconv.Itoa(userId))

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 || res.IsError() {
		return nil, errors.New("failed to update student data")
	}

	return result, nil
}

//	func GetLectureKk(userId int) (*web.DetailLectureResponseApi, error) {
//		client := resty.New()
//		result := &web.DetailLectureResponseApi{}
//
//		res, err := client.R().SetResult(&result).
//			SetHeader("Content-Type", "application/json").
//			Get(apiUrl + "/api/lecturer/" + strconv.Itoa(userId))
//
//		if err != nil {
//			return nil, err
//		}
//
//		fmt.Println(res)
//
//		if res.StatusCode() != 200 || res.IsError() {
//			return nil, errors.New("data lecture not found")
//		}
//
//		return result, nil
//	}
func GetAllTeamMember(teamId int) (*web.MemberTeamResponse, error) {
	client := resty.New()
	result := &web.MemberTeamResponse{}
	res, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		Get(apiUrl + "/api/student/team/" + strconv.Itoa(teamId))

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 || res.IsError() {
		return nil, errors.New("data not found")
	}

	if len(result.Data) < 0 {
		return nil, errors.New("member in team is not found")
	}

	return result, nil
}

//
//func UpdateUserTeamID(userID int, teamID int) (*web.GetDetailStudentResponse, error) {
//	client := resty.New()
//	result := &web.GetDetailStudentResponse{}
//
//	res, err := client.R().SetResult(&result).
//		SetHeader("Content-Type", "application/json").
//		SetBody(map[string]interface{}{
//			"team_id": teamID,
//		}).
//		Patch(apiUrl + "/api/student/team/update/" + strconv.Itoa(userID))
//
//	if err != nil {
//		return nil, err
//	}
//
//	if res.StatusCode() != 200 || res.IsError() {
//		return nil, errors.New("student not found")
//	}
//
//	return result, nil
//}
//
//func ResetTeam(teamID int) (*web.ResetTeamResponse, error) {
//	client := resty.New()
//	result := &web.ResetTeamResponse{}
//	res, err := client.R().SetResult(result).
//		SetHeader("Content-Type", "application/json").
//		SetBody(map[string]interface{}{
//			"team_id": teamID,
//		}).
//		Patch(apiUrl + "/api/student/team/reset")
//
//	if err != nil {
//		return nil, err
//	}
//
//	if res.StatusCode() != 200 || res.IsError() {
//		return nil, errors.New("data not valid")
//	}
//
//	return result, nil
//}

func UpdateUserTeamID(userID int, teamID int) (*web.GetDetailStudentResponse, error) {
	client := resty.New()
	result := &web.GetDetailStudentResponse{}

	res, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"team_id": teamID,
		}).
		Patch(apiUrl + "/api/student/team/update/" + strconv.Itoa(userID))

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 || res.IsError() {
		errResponse := &exception.ErrorResponse{}
		if err = json.Unmarshal(res.Body(), errResponse); err != nil {
			return nil, err
		}
		return nil, errors.New(errResponse.Message)
	}

	return result, nil
}

func GetAdmin() (*web.GetUser, error) {
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	result := &web.GetUser{}

	res, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		Get(apiUrl + "/api/users?role=rladm")

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 || res.IsError() {
		errResponse := &exception.ErrorResponse{}
		if err = json.Unmarshal(res.Body(), errResponse); err != nil {
			return nil, err
		}
		return nil, errors.New(errResponse.Message)
	}

	return result, nil
}
