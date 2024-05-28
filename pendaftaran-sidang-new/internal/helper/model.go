package helper

import (
	"pendaftaran-sidang-new/internal/model/entity"
	"pendaftaran-sidang-new/internal/model/web"
)

func ToPeriodDetailResponse(period *entity.Period) web.PeriodDetailResponse {
	var pengajuans []web.PengajuanResponse
	for _, pengajuan := range period.Pengajuans {
		pengajuans = append(pengajuans, ToPengajuanResponse(&pengajuan))
	}

	return web.PeriodDetailResponse{
		Id:          period.ID,
		Name:        period.Name,
		StartDate:   period.StartDate,
		EndDate:     period.EndDate,
		Description: period.Description,
		Pengajuans:  pengajuans,
		CreatedAt:   period.CreatedAt,
		UpdatedAt:   period.UpdatedAt,
	}

}

func ToPeriodResponse(period *entity.Period) web.PeriodResponse {
	return web.PeriodResponse{
		Id:          period.ID,
		Name:        period.Name,
		StartDate:   period.StartDate,
		EndDate:     period.EndDate,
		Description: period.Description,
		CreatedAt:   period.CreatedAt,
		UpdatedAt:   period.UpdatedAt,
	}

}

func ToPengajuanResponse(pengajuan *entity.Pengajuan) web.PengajuanResponse {
	return web.PengajuanResponse{
		Id:             pengajuan.ID,
		UserId:         pengajuan.UserId,
		Nim:            pengajuan.Nim,
		Pembimbing1Id:  pengajuan.Pembimbing1Id,
		Pembimbing2Id:  pengajuan.Pembimbing2Id,
		Judul:          pengajuan.Judul,
		Eprt:           pengajuan.Eprt,
		DocTa:          pengajuan.DocTa,
		Makalah:        pengajuan.Makalah,
		Tak:            pengajuan.Tak,
		Status:         pengajuan.Status,
		SksLulus:       pengajuan.SksLulus,
		SksBelumLulus:  pengajuan.SksBelumLulus,
		IsEnglish:      pengajuan.IsEnglish,
		PeriodID:       pengajuan.PeriodID,
		SkPenguji:      pengajuan.SkPenguji,
		FormBimbingan1: pengajuan.FormBimbingan1,
		FormBimbingan2: pengajuan.FormBimbingan2,
		Kk:             pengajuan.Kk,
		CreatedAt:      pengajuan.CreatedAt,
		Updated_at:     pengajuan.UpdatedAt,
	}
}

func ToPengajuanDetailResponse(pengajuan *entity.Pengajuan) web.PengajuanDetailResponse {
	var statusLogs []web.StatusLogResponse
	for _, status := range pengajuan.StatusLogs {
		statusLogs = append(statusLogs, ToStatusLogResponse(&status))
	}

	var documentLogResponse web.DocumentLogResponse
	if len(pengajuan.DocumentLogs) > 0 {
		documentLogResponse = ToDocumentLogResponse(&pengajuan.DocumentLogs[0])
		return web.PengajuanDetailResponse{
			Id:             pengajuan.ID,
			UserId:         pengajuan.UserId,
			Nim:            pengajuan.Nim,
			Pembimbing1Id:  pengajuan.Pembimbing1Id,
			Pembimbing2Id:  pengajuan.Pembimbing2Id,
			Judul:          pengajuan.Judul,
			Eprt:           pengajuan.Eprt,
			DocTa:          pengajuan.DocTa,
			Makalah:        pengajuan.Makalah,
			Tak:            pengajuan.Tak,
			Status:         pengajuan.Status,
			SksLulus:       pengajuan.SksLulus,
			SksBelumLulus:  pengajuan.SksBelumLulus,
			IsEnglish:      pengajuan.IsEnglish,
			PeriodID:       pengajuan.PeriodID,
			SkPenguji:      pengajuan.SkPenguji,
			FormBimbingan1: pengajuan.FormBimbingan1,
			FormBimbingan2: pengajuan.FormBimbingan2,
			Kk:             pengajuan.Kk,
			Slide:          &documentLogResponse,
			StatusLogs:     statusLogs,
			CreatedAt:      pengajuan.CreatedAt,
			Updated_at:     pengajuan.UpdatedAt,
		}
	}

	return web.PengajuanDetailResponse{
		Id:             pengajuan.ID,
		UserId:         pengajuan.UserId,
		Nim:            pengajuan.Nim,
		Pembimbing1Id:  pengajuan.Pembimbing1Id,
		Pembimbing2Id:  pengajuan.Pembimbing2Id,
		Judul:          pengajuan.Judul,
		Eprt:           pengajuan.Eprt,
		DocTa:          pengajuan.DocTa,
		Makalah:        pengajuan.Makalah,
		Tak:            pengajuan.Tak,
		Status:         pengajuan.Status,
		SksLulus:       pengajuan.SksLulus,
		SksBelumLulus:  pengajuan.SksBelumLulus,
		IsEnglish:      pengajuan.IsEnglish,
		PeriodID:       pengajuan.PeriodID,
		SkPenguji:      pengajuan.SkPenguji,
		FormBimbingan1: pengajuan.FormBimbingan1,
		FormBimbingan2: pengajuan.FormBimbingan2,
		Kk:             pengajuan.Kk,
		StatusLogs:     statusLogs,
		CreatedAt:      pengajuan.CreatedAt,
		Updated_at:     pengajuan.UpdatedAt,
	}

}

func ToStatusLogResponse(statusLog *entity.StatusLog) web.StatusLogResponse {
	return web.StatusLogResponse{
		Id:           statusLog.ID,
		Feedback:     statusLog.Feedback,
		CreatedBy:    statusLog.CreatedBy,
		WorkFlowType: statusLog.WorkFlowType,
		Name:         statusLog.Name,
		PengajuanId:  statusLog.PengajuanID,
		CreatedAt:    statusLog.CreatedAt,
		UpdatedAt:    statusLog.UpdatedAt,
	}
}

func ToDocumentLogResponse(document *entity.DocumentLog) web.DocumentLogResponse {
	return web.DocumentLogResponse{
		Id:          document.ID,
		PengajuanId: document.PengajuanID,
		FileName:    document.FileName,
		Type:        document.Type,
		FileUrl:     document.FileUrl,
		CreatedBy:   document.CreatedBy,
		CreatedAt:   document.CreatedAt,
		UpdatedAt:   document.UpdatedAt,
	}
}

func ToDocumentLogDetailResponse(document *entity.DocumentLog) web.DocumentLogDetailResponse {
	return web.DocumentLogDetailResponse{
		Id:        document.ID,
		FileName:  document.FileName,
		Type:      document.Type,
		FileUrl:   document.FileUrl,
		CreatedBy: document.CreatedBy,
		Pengajuan: ToPengajuanResponse(&document.Pengajuan),
		CreatedAt: document.CreatedAt,
		UpdatedAt: document.UpdatedAt,
	}
}

func ToTeamResponse(team *entity.Team) web.TeamResponse {
	return web.TeamResponse{
		Id:        team.ID,
		Name:      team.Name,
		CreatedAt: team.CreatedAt,
		UpdatedAt: team.UpdatedAt,
	}
}

func ToTeamDetailResponse(team *entity.Team, teamMembers *web.MemberTeamResponse) web.TeamResponseDetail {
	members := make([]web.MemberData, 0)
	for _, memberDataApi := range teamMembers.Data {
		if memberDataApi.Pengajuan != nil {
			member := web.MemberData{
				UserId:      memberDataApi.UserId,
				TeamId:      memberDataApi.TeamId,
				PeminatanId: memberDataApi.PeminatanId,
				Nim:         memberDataApi.Nim,
				Username:    memberDataApi.User.Username,
				Name:        memberDataApi.User.Nama,
				Pengajuan:   memberDataApi.Pengajuan,
			}

			members = append(members, member)
		}

	}

	return web.TeamResponseDetail{
		Id:        team.ID,
		Name:      team.Name,
		Members:   members,
		CreatedAt: team.CreatedAt,
		UpdatedAt: team.UpdatedAt,
	}
}

func ToNotificationResponse(notification *entity.Notification) web.NotificationResponse {
	return web.NotificationResponse{
		Id:        notification.ID,
		UserId:    notification.UserId,
		Title:     notification.Title,
		Message:   notification.Message,
		Url:       notification.Url,
		ReadAt:    notification.ReadAt,
		CreatedAt: notification.CreatedAt,
		UpdatedAt: notification.UpdatedAt,
	}
}
