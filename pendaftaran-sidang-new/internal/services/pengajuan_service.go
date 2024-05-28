package services

import "pendaftaran-sidang-new/internal/model/web"

type PengajuanService interface {
	GetAllPengajuan(filter map[string]interface{}) ([]web.PengajuanResponse, error)
	GetPengajuanById(pengajuanId int) (*web.PengajuanDetailResponse, error)
	GetPengajuanByUserId(userId int) (*web.PengajuanDetailResponse, error)
	Create(request *web.PengajuanCreateRequest) (*web.PengajuanResponse, error)
	Update(request *web.PengajuanUpdateRequest) (*web.PengajuanResponse, error)
	AdminStatus(request *web.StatusAdminUpdate) (*web.PengajuanResponse, error)
	ChangeStatus(request *web.ChangeStatusRequest) (*web.PengajuanResponse, error)
	Delete(pengajuanId int) error
}
