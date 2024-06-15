package web

type StudentData struct {
	Nim    int    `json:"nim"`
	UserId int    `json:"user_id"`
	TeamId int    `json:"team_id"`
	Kk     string `json:"kk"`
	User   struct {
		Name string `json:"nama"`
	}
}

type GetDetailStudentResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   StudentData `json:"data"`
}

type DetailLectureResponseApi struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id     int    `json:"id"`
		Nip    string `json:"nip"`
		Code   string `json:"code"`
		Jfa    string `json:"jfa"`
		Kk     string `json:"kk"`
		UserId int    `json:"user_id"`
	} `json:"data"`
}

type GetUser struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id       int    `json:"id"`
		Username string `json:"nip"`
	} `json:"data"`
}

type MemberTeamResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []struct {
		Nim         int                `json:"nim"`
		UserId      int                `json:"user_id"`
		TeamId      int                `json:"team_id"`
		PeminatanId int                `json:"peminatan_id"`
		Pengajuan   *PengajuanResponse `json:"pengajuan"`
		User        struct {
			Username string `json:"username"`
			Nama     string `json:"nama"`
		} `json:"user"`
	} `json:"data"`
}

type ResetTeamResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []int  `json:"data"`
}
