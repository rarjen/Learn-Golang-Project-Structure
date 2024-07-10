package request

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_LIMIT = 10
)

type Pagination struct {
	Limit int
	Page  int
}

func (pag Pagination) Offset() int {
	return (pag.Page - 1) * pag.Limit
}

func parsePagination(ctx *gin.Context) Pagination {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	return Pagination{
		Limit: limit,
		Page:  page,
	}
}

type CreatePipelineRequest struct {
	Name                string  `json:"name" validate:"required"`
	PlaceOfBirth        string  `json:"place_of_birth" validate:"required"`
	DateOfBirth         string  `json:"date_of_birth" validate:"required"`
	Gender              string  `json:"gender" validate:"required"`
	EconomicSector      string  `json:"economic_sector" validate:"required"`
	EconomicSubsector   string  `json:"economic_subsector" validate:"required"`
	BusinessPeriod      int     `json:"business_period" validate:"required"`
	BusinessPlaceStatus string  `json:"business_place_status" validate:"required"`
	ProductPlan         int     `json:"product_plan" validate:"required"`
	LoanPlan            float64 `json:"loan_plan" validate:"required"`
	PhoneNumber         string  `json:"phone_number" validate:"required"`
}

type CreateUserRequest struct {
	IDEmployee   string    `json:"id_employee" validate:"required"`
	Username     string    `json:"username" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	IsActive     int       `json:"is_active" validate:"required"`
	CreatedBy    string    `json:"created_by" validate:"required"`
	CreatedTime  time.Time `json:"created_time"`
	ModifiedBy   string    `json:"modified_by"`
	ModifiedTime time.Time `json:"modified_time"`
}

type CreateProductRequest struct {
	ProductName        string    `json:"product_name" validate:"required"`
	ProductCode        string    `json:"product_code" validate:"required"`
	InterestRate       float64   `json:"interest_rate" validate:"required"`
	InterestRateAnnual float64   `json:"interest_rate_annual" validate:"required"`
	LimitLoanLower     float64   `json:"limit_loan_lower" validate:"required"`
	LimitLoanUpper     float64   `json:"limit_loan_upper" validate:"required"`
	TimePeriodLower    int       `json:"time_period_lower" validate:"required"`
	TimePeriodUpper    int       `json:"time_period_upper" validate:"required"`
	IsActive           int       `json:"is_active" validate:"required"`
	CreatedBy          string    `json:"created_by" validate:"required"`
	CreatedTime        time.Time `json:"created_time"`
	ModifiedBy         string    `json:"modified_by"`
	ModifiedTime       time.Time `json:"modified_time"`
}

type CreateProgramRequest struct {
	ProgramName string    `json:"program_name" validate:"required"`
	IsActive    int       `json:"is_active" validate:"required"`
	CreatedBy   string    `json:"created_by" validate:"required"`
	CreatedTime time.Time `json:"created_time"`
}

type UpdateProgramRequest struct {
	ProgramName string `json:"program_name" validate:"required"`
	IsActive    int    `json:"is_active" validate:"required"`
	ModifiedBy  string `json:"modified_by" validate:"required"`
}

type GetIdUserRequest struct {
	IDEmployee string `uri:"id_employee" binding:"required"`
}

type IdProgramRequest struct {
	IDProgram int `uri:"id_program" binding:"required"`
}

type IdProductParam struct {
	IDProduct int `uri:"id_product" binding:"required"`
}

type UpdateUserRequest struct {
	IDEmployee   string    `json:"id_employee"`
	Username     string    `json:"username"`
	Name         string    `json:"name"`
	IsActive     int       `json:"is_active"`
	CreatedBy    string    `json:"created_by"`
	CreatedTime  time.Time `json:"created_time"`
	ModifiedBy   string    `json:"modified_by"`
	ModifiedTime time.Time `json:"modified_time"`
}

type CreateCommunityPipelineRequest struct {
	PipelineIds       []int64 `json:"id_pipelines" validate:"required"`
	ChiefId           int64   `json:"chief_id" validate:"required"`
	CommunityName     string  `json:"community_name" validate:"required"`
	EconomicSector    string  `json:"economic_sector" validate:"required"`
	EconomicSubsector string  `json:"economic_subsector" validate:"required"`
}

type CreateMemberPipelineRequest struct {
	PipelineIds []int64 `json:"id_pipelines" validate:"required"`
	CommunityId int64   `json:"id_community" validate:"required"`
}

type FetchPipelineRequest struct {
	CommunityId        int64
	ExcludeCommunityId int64
	Keyword            string
	Pagination

	// Not mean to be used directly
	ExcludePipelineIds []int64
}

// Change chief community
type ChangeCommunityChiefRequest struct {
	CommunityId int64 `json:"id_community" validate:"required"`
	NewChiefId  int64 `json:"new_chief_id" validate:"required"`
}
