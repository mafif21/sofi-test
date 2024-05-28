package helper

import (
	"penjadwalan-sidang-new/internal/model/entity"
	"penjadwalan-sidang-new/internal/model/web"
)

//func ToScheduleDetailResponse(schedule *entity.Schedule, pengajuan *web.PengajuanResponseApi) web.ScheduleDetailResponse {
//
//	return web.ScheduleDetailResponse{
//		ID:               schedule.ID,
//		DateTime:         schedule.DateTime,
//		Room:             schedule.Room,
//		Penguji1Id:       schedule.Penguji1Id,
//		Penguji2Id:       schedule.Penguji2Id,
//		Status:           schedule.Status,
//		Decision:         schedule.Decision,
//		RevisionDuration: schedule.RevisionDuration,
//		FlagAddRevision:  schedule.FlagAddRevision,
//		FlagChangeScores: schedule.FlagChangeScores,
//		Pengajuan: web.PengajuanRelationship{
//			Id:             pengajuan.Data.Id,
//			UserId:         pengajuan.Data.UserId,
//			Nim:            pengajuan.Data.Nim,
//			Pembimbing1Id:  pengajuan.Data.Pembimbing1Id,
//			Pembimbing2Id:  pengajuan.Data.Pembimbing2Id,
//			Judul:          pengajuan.Data.Judul,
//			Eprt:           pengajuan.Data.Eprt,
//			DocTa:          pengajuan.Data.DocTa,
//			Makalah:        pengajuan.Data.Makalah,
//			Tak:            pengajuan.Data.Tak,
//			Status:         pengajuan.Data.Status,
//			SksLulus:       pengajuan.Data.SksLulus,
//			SksBelumLulus:  pengajuan.Data.SksBelumLulus,
//			IsEnglish:      pengajuan.Data.IsEnglish,
//			PeriodID:       pengajuan.Data.PeriodId,
//			SkPenguji:      pengajuan.Data.SkPenguji,
//			FormBimbingan1: pengajuan.Data.FormBimbingan1,
//			FormBimbingan2: pengajuan.Data.FormBimbingan2,
//			Kk:             pengajuan.Data.Kk,
//			Slide: web.SlideRelationship{
//				Id:          pengajuan.Data.Slide.Id,
//				PengajuanId: pengajuan.Data.Slide.PengajuanId,
//				FileName:    pengajuan.Data.Slide.FileName,
//				Type:        pengajuan.Data.Slide.Type,
//				FileUrl:     pengajuan.Data.Slide.FileUrl,
//				CreatedBy:   pengajuan.Data.Slide.CreatedBy,
//				CreatedAt:   pengajuan.Data.Slide.CreatedAt,
//				UpdatedAt:   pengajuan.Data.Slide.UpdatedAt,
//			},
//			CreatedAt:  pengajuan.Data.CreatedAt,
//			Updated_at: pengajuan.Data.UpdatedAt,
//		},
//		CreatedAt: schedule.CreatedAt,
//		UpdatedAt: schedule.UpdatedAt,
//	}
//}

func ToScheduleDetailResponse(schedule *entity.Schedule, pengajuan *web.Pengajuan) web.ScheduleDetailResponse {

	return web.ScheduleDetailResponse{
		ID:               schedule.ID,
		DateTime:         schedule.DateTime,
		Room:             schedule.Room,
		Penguji1Id:       schedule.Penguji1Id,
		Penguji2Id:       schedule.Penguji2Id,
		Status:           schedule.Status,
		Decision:         schedule.Decision,
		RevisionDuration: schedule.RevisionDuration,
		FlagAddRevision:  schedule.FlagAddRevision,
		FlagChangeScores: schedule.FlagChangeScores,
		Pengajuan: web.PengajuanRelationship{
			Id:             pengajuan.Id,
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
			PeriodID:       pengajuan.PeriodId,
			SkPenguji:      pengajuan.SkPenguji,
			FormBimbingan1: pengajuan.FormBimbingan1,
			FormBimbingan2: pengajuan.FormBimbingan2,
			Kk:             pengajuan.Kk,
			Slide: web.SlideRelationship{
				Id:          pengajuan.Slide.Id,
				PengajuanId: pengajuan.Slide.PengajuanId,
				FileName:    pengajuan.Slide.FileName,
				Type:        pengajuan.Slide.Type,
				FileUrl:     pengajuan.Slide.FileUrl,
				CreatedBy:   pengajuan.Slide.CreatedBy,
				CreatedAt:   pengajuan.Slide.CreatedAt,
				UpdatedAt:   pengajuan.Slide.UpdatedAt,
			},
			CreatedAt:  pengajuan.CreatedAt,
			Updated_at: pengajuan.UpdatedAt,
		},
		CreatedAt: schedule.CreatedAt,
		UpdatedAt: schedule.UpdatedAt,
	}
}

func ToScheduleResponse(schedule *entity.Schedule) web.ScheduleResponse {
	return web.ScheduleResponse{
		ID:               schedule.ID,
		DateTime:         schedule.DateTime,
		Room:             schedule.Room,
		Penguji1Id:       schedule.Penguji1Id,
		Penguji2Id:       schedule.Penguji2Id,
		Status:           schedule.Status,
		Decision:         schedule.Decision,
		RevisionDuration: schedule.RevisionDuration,
		FlagAddRevision:  schedule.FlagAddRevision,
		FlagChangeScores: schedule.FlagChangeScores,
		CreatedAt:        schedule.CreatedAt,
		UpdatedAt:        schedule.UpdatedAt,
	}
}
