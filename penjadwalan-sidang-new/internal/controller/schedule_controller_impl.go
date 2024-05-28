package controller

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"penjadwalan-sidang-new/internal/exception"
	"penjadwalan-sidang-new/internal/helper"
	"penjadwalan-sidang-new/internal/model/web"
	"penjadwalan-sidang-new/internal/services"
	"strconv"
	"strings"
)

type ScheduleControllerImpl struct {
	ScheduleService services.ScheduleService
}

func NewScheduleController(scheduleService services.ScheduleService) ScheduleController {
	return &ScheduleControllerImpl{ScheduleService: scheduleService}
}

func (c ScheduleControllerImpl) FindAll(ctx *fiber.Ctx) error {
	page := ctx.Query("pages", "1")
	pageInt, err := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	var pengajuanId []int
	pengajuanDatasKk := helper.GetPengajuanKk(ctx.Get("Authorization"))
	for _, pengajuan := range pengajuanDatasKk.Data {
		pengajuanId = append(pengajuanId, pengajuan.Id)
	}

	filter := map[string]interface{}{
		"pengajuan_id": pengajuanId,
	}

	allDatas, err := c.ScheduleService.GetAll(filter, pageInt, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all data schedule",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) FindAllAdmin(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	page := ctx.Query("pages", "1")
	pageInt, err := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	var filter map[string]interface{}
	path := ctx.Path()
	if path == "/api/schedule/get/admin" {
		filter = map[string]interface{}{
			"keputusan": []string{"lulus", "lulus bersyarat"},
		}
	}

	if path == "/api/schedule/get/admin-before" {
		filter = map[string]interface{}{
			"status": []string{"belum dilaksanakan", "sedang dilaksanakan"},
		}
	}

	allDatas, err := c.ScheduleService.GetAll(filter, pageInt, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all data schedule",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) FindAllSuperadmin(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	page := ctx.Query("pages", "1")
	pageInt, err := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	filter := map[string]interface{}{
		"keputusan": []interface{}{"lulus", "lulus bersyarat", ""},
	}

	allDatas, err := c.ScheduleService.GetAll(filter, pageInt, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all data schedule",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) FindAllPenguji(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPGJ"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	userLoggedIn := ctx.Locals("user_id")
	userId := userLoggedIn.(int)

	page := ctx.Query("pages", "1")
	pageInt, err := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	filter := map[string]interface{}{
		"penguji": userId,
	}

	allDatas, err := c.ScheduleService.GetAll(filter, pageInt, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all data schedule",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) FindAllPembimbing(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPBB"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	page := ctx.Query("pages", "1")
	pageInt, err := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	userId := 14
	var pengajuanId []int

	pengajuanDatas := helper.GetPengajuanByPembimbingId(userId, ctx.Get("Authorization"))
	for _, pengajuan := range pengajuanDatas.Data {
		pengajuanId = append(pengajuanId, pengajuan.Id)
	}

	filter := map[string]interface{}{
		"pengajuan_id": pengajuanId,
	}

	allDatas, err := c.ScheduleService.GetAll(filter, pageInt, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all data schedule",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) FindAllPic(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	page := ctx.Query("pages", "1")
	pageInt, err := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	var pengajuanId []int
	pengajuanDatasKk := helper.GetPengajuanKk(ctx.Get("Authorization"))
	for _, pengajuan := range pengajuanDatasKk.Data {
		pengajuanId = append(pengajuanId, pengajuan.Id)
	}

	filter := map[string]interface{}{
		"pengajuan_id": pengajuanId,
		"keputusan":    []interface{}{"lulus bersyarat", "", nil},
	}

	allDatas, err := c.ScheduleService.GetAll(filter, pageInt, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all data schedule",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) FindScheduleList(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLSPR", "RLADM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	page := ctx.Query("pages", "1")
	pageInt, err := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	pengajuanId, err := ctx.ParamsInt("pengajuanId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "pengajuan id is not valid",
		})
	}

	filter := map[string]interface{}{
		"pengajuan_id": pengajuanId,
	}

	allDatas, err := c.ScheduleService.GetAll(filter, pageInt, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all data schedule",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) FindById(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPIC", "RLPGJ", "RLPBB", "RLSPR", "RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	scheduleId, err := ctx.ParamsInt("scheduleId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "schedule id is not valid",
		})
	}

	allDatas, err := c.ScheduleService.GetById(scheduleId, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "schedule has found",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) FindPengajuanSchedules(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPIC", "RLPGJ", "RLPBB", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	status := ctx.Query("status")
	order := ctx.Query("order", "desc")

	pengajuanId, err := ctx.ParamsInt("pengajuanId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "pengajuan id is not valid",
		})
	}

	filter := map[string]interface{}{
		"pengajuan_id": pengajuanId,
	}

	if status != "" {
		filter["status"] = status
	}

	allDatas, err := c.ScheduleService.GetPengajuanSchedules(pengajuanId, filter, order)
	if err != nil || len(allDatas) < 1 {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "schedule has found",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) FindAllMahasiswa(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	page := ctx.Query("pages", "1")
	pageInt, err := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	userPengajuan := helper.GetPengajuanUserLoggedIn(ctx.Get("Authorization"))
	if userPengajuan.Code != 200 {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "user dont have pengajuan data",
		})
	}

	getUserSchedule, err := c.ScheduleService.GetByPengajuan(&userPengajuan.Data)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	//detailTeam := helper.GetTeamUserLoggedIn(ctx.Get("Authorization"))

	//var validMember []map[string]interface{}
	//
	//for _, member := range detailTeam.Data.Members {
	//	_, err := c.ScheduleService.GetByPengajuan(&member.Pengajuan)
	//	if err == nil {
	//		memberDetails := make(map[string]interface{})
	//
	//		memberDetails["name"] = member.Name
	//		memberDetails["nim"] = member.Nim
	//		validMember = append(validMember, memberDetails)
	//	}
	//}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all data schedule",
		Data:   getUserSchedule,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) Create(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}
	newScheduleRequest := web.ScheduleCreateRequest{}

	if err := ctx.BodyParser(&newScheduleRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	newScheduleRequest.Room = strings.ToLower(newScheduleRequest.Room)

	newSchedule, err := c.ScheduleService.Create(&newScheduleRequest, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "new schedule has been created",
		Data:   newSchedule,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (c ScheduleControllerImpl) Update(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	updatedScheduleRequest := web.ScheduleUpdateRequest{}
	scheduleId, err := ctx.ParamsInt("scheduleId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "schedule id is not valid",
		})
	}

	updatedScheduleRequest.Id = scheduleId

	if err := ctx.BodyParser(&updatedScheduleRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	updatedScheduleRequest.Room = strings.ToLower(updatedScheduleRequest.Room)

	newSchedule, err := c.ScheduleService.Update(&updatedScheduleRequest, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "schedule has been updated",
		Data:   newSchedule,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) Delete(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	scheduleId, err := ctx.ParamsInt("scheduleId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "schedule id is not valid",
		})
	}

	err = c.ScheduleService.Delete(scheduleId, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "schedule has been delete",
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) AddFlag(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	code := ctx.Query("code")
	scheduleId, err := ctx.ParamsInt("scheduleId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "schedule id is not valid",
		})
	}

	updatedData, err := c.ScheduleService.AddFlag(code, scheduleId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "schedule has been update",
		Data:   updatedData,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (c ScheduleControllerImpl) ChangeStatus(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	scheduleId, err := ctx.ParamsInt("scheduleId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "schedule id is not valid",
		})
	}

	updateStatusRequest := &web.ScheduleUpdateStatusRequest{}
	if err := ctx.BodyParser(&updateStatusRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	updateStatusRequest.Id = scheduleId

	updatedData, err := c.ScheduleService.ChangeStatus(updateStatusRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "schedule has been update",
		Data:   updatedData,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
