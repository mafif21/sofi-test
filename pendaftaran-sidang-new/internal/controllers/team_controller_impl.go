package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"pendaftaran-sidang-new/internal/exception"
	"pendaftaran-sidang-new/internal/model/web"
	"pendaftaran-sidang-new/internal/services"
	"strconv"
)

type TeamControllerImpl struct {
	TeamService      services.TeamService
	PengajuanService services.PengajuanService
}

func NewTeamContoller(services services.TeamService, pengajuanService services.PengajuanService) TeamController {
	return &TeamControllerImpl{
		TeamService:      services,
		PengajuanService: pengajuanService,
	}
}

func (c TeamControllerImpl) FindAll(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	page := ctx.Query("pages", "1")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "value is invalid"))
	}

	limit := ctx.Query("limit", "10")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "value is invalid"))
	}

	if pageInt < 1 {
		pageInt = 1
	}

	var filter = make(map[string]interface{})

	teams, err := c.TeamService.GetAllTeam(filter, pageInt, limitInt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all teams data",
		Data:   teams,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c TeamControllerImpl) FindTeamById(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLSPR", "RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	teamId, err := ctx.ParamsInt("teamId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "id is not valid"))
	}

	foundTeam, err := c.TeamService.GetTeamById(teamId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success get team by id",
		Data:   foundTeam,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)

}

func (c TeamControllerImpl) FindTeamByUserId(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	studentLoggedIn := ctx.Locals("user_id")
	userId := studentLoggedIn.(int)

	foundTeam, err := c.TeamService.GetTeamByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success get user team",
		Data:   foundTeam,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c TeamControllerImpl) CreateTeam(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	studentLoggedIn := ctx.Locals("user_id")
	userId := studentLoggedIn.(int)

	teamRequest := &web.TeamCreateRequest{}
	teamRequest.UserId = userId

	if err := ctx.BodyParser(&teamRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "failed to parse json"))
	}

	newTeam, err := c.TeamService.Create(teamRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "team " + newTeam.Name + " has been created",
		Data:   newTeam,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c TeamControllerImpl) CreatePersonal(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	studentLoggedIn := ctx.Locals("user_id")
	userId := studentLoggedIn.(int)

	teamRequest := &web.TeamCreateRequest{}
	teamRequest.UserId = userId

	foundPengajuan, err := c.PengajuanService.GetPengajuanByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	teamRequest.Name = foundPengajuan.Nim + " Sidang Individu"

	newTeam, err := c.TeamService.Create(teamRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "team " + newTeam.Name + " has been created",
		Data:   newTeam,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c TeamControllerImpl) AddMember(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS", "RLSPR", "RLADM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	addMemberRequest := &web.MemberRequest{}

	if err := ctx.BodyParser(&addMemberRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "failed to parse json"))
	}

	nim, err := c.TeamService.AddMember(addMemberRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success add new member",
		Data:   nim,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c TeamControllerImpl) LeaveTeam(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS", "RLSPR", "RLADM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	addMemberRequest := &web.MemberRequest{}

	if err := ctx.BodyParser(&addMemberRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "failed to parse json"))
	}

	nim, err := c.TeamService.LeaveTeam(addMemberRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success leave the team",
		Data:   nim,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c TeamControllerImpl) AvailableMember(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	member, err := c.TeamService.GetAvailableMember()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "get available member",
		Data:   member,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c TeamControllerImpl) Update(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS", "RLSPR", "RLADM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	teamId, err := ctx.ParamsInt("teamId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "team id is not valid"))
	}

	teamUpdateRequest := &web.TeamUpdateRequest{}
	teamUpdateRequest.Id = teamId

	if err := ctx.BodyParser(&teamUpdateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "failed to parse json"))
	}

	updatedData, err := c.TeamService.Update(teamUpdateRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "team " + updatedData.Name + " has been updated",
		Data:   updatedData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c TeamControllerImpl) Delete(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLSPR", "RLADM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(exception.NewErrorResponse(fiber.StatusUnauthorized, "unauthorized user"))
	}

	userLoggedIn := ctx.Locals("user_id")
	userId := userLoggedIn.(int)

	teamId, err := ctx.ParamsInt("teamId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "team id is not valid"))
	}

	err = c.TeamService.Delete(teamId, userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success delete the team",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
