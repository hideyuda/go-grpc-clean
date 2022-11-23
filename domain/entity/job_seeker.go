package entity

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type JobSeeker struct {
	ID                      uint      `db:"id" json:"id"`
	UUID                    uuid.UUID `db:"uuid" json:"uuid"`
	AgentID                 uint      `db:"agent_id" json:"agent_id"`
	AgentStaffID            null.Int  `db:"agent_staff_id" json:"agent_staff_id"`
	LineID                  string    `db:"line_id" json:"line_id"`
	AgentName               string    `db:"agent_name" json:"agent_name"`
	StaffName               string    `db:"staff_name" json:"staff_name"`
	UserStatus              null.Int  `db:"user_status" json:"user_status"`
	LastName                string    `db:"last_name" json:"last_name"`
	FirstName               string    `db:"first_name" json:"first_name"`
	LastFurigana            string    `db:"last_furigana" json:"last_furigana"`
	FirstFurigana           string    `db:"first_furigana" json:"first_furigana"`
	Gender                  null.Int  `db:"gender" json:"gender"`
	GenderRemarks           string    `db:"gender_remarks" json:"gender_remarks"`
	Birthday                string    `db:"birthday" json:"birthday"`
	Spouse                  null.Int  `db:"spouse" json:"spouse"`
	SupportObligation       null.Int  `db:"support_obligation" json:"support_obligation"`
	Dependents              null.Int  `db:"dependents" json:"dependents"`
	PhoneNumber             string    `db:"phone_number" json:"phone_number"`
	Email                   string    `db:"email" json:"email"`
	EmergencyPhoneNumber    string    `db:"emergency_phone_number" json:"emergency_phone_number"`
	PostCode                string    `db:"post_code" json:"post_code"`
	Prefecture              null.Int  `db:"prefecture" json:"prefecture"`
	Address                 string    `db:"address" json:"address"`
	StateOfEmployment       null.Int  `db:"state_of_employment" json:"state_of_employment"`
	JobSummary              string    `db:"job_summary" json:"job_summary"`
	ResearchContent         string    `db:"research_content" json:"research_content"`
	PcSkill                 null.Int  `db:"pc_skill" json:"pc_skil"`
	JoinCompanyPeriod       null.Int  `db:"join_company_period" json:"join_company_period"`
	JobChange               null.Int  `db:"job_change" json:"job_change"`
	AnnualIncome            null.Int  `db:"annual_income" json:"annual_income"`
	DesiredAnnualIncome     null.Int  `db:"desired_annual_income" json:"desired_annual_income"`
	DesiredCompanyScale     null.Int  `db:"desired_company_scale" json:"desired_company_scale"`
	Transfer                null.Int  `db:"transfer" json:"transfer"`
	TransferRequirement     string    `db:"transfer_requirement" json:"transfer_requirement"`
	ShortResignation        null.Int  `db:"short_resignation" json:"short_resignation"`
	ShortResignationRemarks string    `db:"short_resignation_remarks" json:"short_resignation_remarks"`
	MedicalHistory          null.Int  `db:"medical_history" json:"medical_history"`
	Nationality             null.Int  `db:"nationality" json:"nationality"`
	Appearance              null.Int  `db:"appearance" json:"appearance"`
	AppearanceDetail        string    `db:"appearance_detail" json:"appearance_detail"`
	Communication           null.Int  `db:"communication" json:"communication"`
	CommunicationDetail     string    `db:"communication_detail" json:"communication_detail"`
	Thinking                null.Int  `db:"thinking" json:"thinking"`
	SecretMemo              string    `db:"secret_memo" json:"secret_memo"`
	JobHuntingState         null.Int  `db:"job_hunting_state" json:"job_hunting_state"`
	ContactManner           null.Int  `db:"contact_manner" json:"contact_manner"`
	PublishToAlliance       null.Int  `db:"publish_to_alliance" json:"publish_to_alliance"`
	RecommendReason         string    `db:"recommend_reason" json:"recommend_reason"`
	Phase                   null.Int  `db:"phase" json:"phase"`
	InterviewDate           string    `db:"interview_date" json:"interview_date"`
	CreatedAt               time.Time `db:"created_at" json:"created_at"`
	UpdatedAt               time.Time `db:"updated_at" json:"updated_at"`
	Name                    string    `db:"name" json:"name"`
	Furigana                string    `db:"furigana" json:"furigana"`

	// 他テーブル
	// 求職者の学歴情報
	// StudentHistories []JobSeekerStudentHistory `db:"student_histories" json:"student_histories"`

	// //求職者の職歴
	// WorkHistories []JobSeekerWorkHistory `db:"work_histories" json:"work_histories"`

	// //求職者の所持資格
	// Licenses []JobSeekerLicense `db:"licenses" json:"licenses"`

	// //求職者の自己PR
	// SelfPromotions []JobSeekerSelfPromotion `db:"self_promotions" json:"self_promotions"`

	// //求職者の資料
	// Documents JobSeekerDocument `db:"documents" json:"documents"`

	// //求職者の希望業界
	// DesiredIndustries []JobSeekerDesiredIndustry `db:"desired_industries" json:"desired_industries"`

	// //求職者の希望職種
	// DesiredOccupations []JobSeekerDesiredOccupation `db:"desired_occupations" json:"desired_occupations"`

	// //求職者の希望勤務地
	// DesiredWorkLocations []JobSeekerDesiredWorkLocation `db:"desired_work_locations" json:"desired_work_locations"`

	// //求職者の開発スキル
	// DevelopmentSkills []JobSeekerDevelopmentSkill `db:"development_skills" json:"development_skills"`

	// //求職者の言語スキル
	// LanguageSkills []JobSeekerLanguageSkill `db:"language_skills" json:"language_skills"`

	// //求職者のPCスキル
	// PcSkills []JobSeekerPcSkill `db:"pc_skills" json:"pc_skills"`

	// //求職者の休日タイプ
	// DesiredHolidayTypes []JobSeekerDesiredHolidayType `db:"desired_holiday_types" json:"desired_holiday_types"`

	// // マッチング用
	// MatchingScore null.Int `db:"-" json:"matching_score"`
}

func NewJobSeeker(
	agent_id uint,
	agentStaffID null.Int,
	lineID string,
	userStatus null.Int,
	lastName string,
	firstName string,
	lastFurigana string,
	firstFurigana string,
	gender null.Int,
	genderRemarks string,
	birthday string,
	spouse null.Int,
	supportObligation null.Int,
	dependents null.Int,
	phoneNumber string,
	email string,
	emergencyPhoneNumber string,
	postCode string,
	prefecture null.Int,
	address string,
	stateOfEmployment null.Int,
	jobSummary string,
	researchContent string,
	pcSkill null.Int,
	joinCompanyPeriod null.Int,
	jobChange null.Int,
	annualIncome null.Int,
	desiredAnnualIncome null.Int,
	desiredCompanyScale null.Int,
	transfer null.Int,
	transferRequirement string,
	shortResignation null.Int,
	shortResignationRemarks string,
	medicalHistory null.Int,
	nationality null.Int,
	appearance null.Int,
	appearanceDetail string,
	communication null.Int,
	communicationDetail string,
	thinking null.Int,
	secretMemo string,
	jobHuntingState null.Int,
	contactManner null.Int,
	publishToAlliance null.Int,
	recommendReason string,
	phase null.Int,
	interviewDate string,
) *JobSeeker {
	return &JobSeeker{
		AgentID:                 agent_id,
		AgentStaffID:            agentStaffID,
		LineID:                  lineID,
		UserStatus:              userStatus,
		LastName:                lastName,
		FirstName:               firstName,
		LastFurigana:            lastFurigana,
		FirstFurigana:           firstFurigana,
		Gender:                  gender,
		GenderRemarks:           genderRemarks,
		Birthday:                birthday,
		Spouse:                  spouse,
		SupportObligation:       supportObligation,
		Dependents:              dependents,
		PhoneNumber:             phoneNumber,
		Email:                   email,
		EmergencyPhoneNumber:    emergencyPhoneNumber,
		PostCode:                postCode,
		Prefecture:              prefecture,
		Address:                 address,
		StateOfEmployment:       stateOfEmployment,
		JobSummary:              jobSummary,
		ResearchContent:         researchContent,
		PcSkill:                 pcSkill,
		JoinCompanyPeriod:       joinCompanyPeriod,
		JobChange:               jobChange,
		AnnualIncome:            annualIncome,
		DesiredAnnualIncome:     desiredAnnualIncome,
		DesiredCompanyScale:     desiredCompanyScale,
		Transfer:                transfer,
		TransferRequirement:     transferRequirement,
		ShortResignation:        shortResignation,
		ShortResignationRemarks: shortResignationRemarks,
		MedicalHistory:          medicalHistory,
		Nationality:             nationality,
		Appearance:              appearance,
		AppearanceDetail:        appearanceDetail,
		Communication:           communication,
		CommunicationDetail:     communicationDetail,
		Thinking:                thinking,
		SecretMemo:              secretMemo,
		JobHuntingState:         jobHuntingState,
		ContactManner:           contactManner,
		PublishToAlliance:       publishToAlliance,
		RecommendReason:         recommendReason,
		Phase:                   phase,
		InterviewDate:           interviewDate,
	}
}

type CreateOrUpdateJobSeekerParam struct {
	AgentID                 uint     `db:"agent_id" json:"agent_id" validate:"required"`
	AgentStaffID            null.Int `db:"agent_staff_id" json:"agent_staff_id"`
	LineID                  string   `db:"line_id" json:"line_id"`
	UserStatus              null.Int `db:"user_status" json:"user_status"`
	LastName                string   `db:"last_name" json:"last_name"`
	FirstName               string   `db:"first_name" json:"first_name"`
	LastFurigana            string   `db:"last_furigana" json:"last_furigana"`
	FirstFurigana           string   `db:"first_furigana" json:"first_furigana"`
	Gender                  null.Int `db:"gender" json:"gender"`
	GenderRemarks           string   `db:"gender_remarks" json:"gender_remarks"`
	Birthday                string   `db:"birthday" json:"birthday"`
	Spouse                  null.Int `db:"spouse" json:"spouse"`
	SupportObligation       null.Int `db:"support_obligation" json:"support_obligation"`
	Dependents              null.Int `db:"dependents" json:"dependents"`
	PhoneNumber             string   `db:"phone_number" json:"phone_number"`
	Email                   string   `db:"email" json:"email"`
	EmergencyPhoneNumber    string   `db:"emergency_phone_number" json:"emergency_phone_number"`
	PostCode                string   `db:"post_code" json:"post_code"`
	Prefecture              null.Int `db:"prefecture" json:"prefecture"`
	Address                 string   `db:"address" json:"address"`
	StateOfEmployment       null.Int `db:"state_of_employment" json:"state_of_employment"`
	JobSummary              string   `db:"job_summary" json:"job_summary"`
	ResearchContent         string   `db:"research_content" json:"research_content"`
	PcSkill                 null.Int `db:"pc_skill" json:"pc_skil"`
	JoinCompanyPeriod       null.Int `db:"join_company_period" json:"join_company_period"`
	JobChange               null.Int `db:"job_change" json:"job_change"`
	AnnualIncome            null.Int `db:"annual_income" json:"annual_income"`
	DesiredAnnualIncome     null.Int `db:"desired_annual_income" json:"desired_annual_income"`
	DesiredCompanyScale     null.Int `db:"desired_company_scale" json:"desired_company_scale"`
	Transfer                null.Int `db:"transfer" json:"transfer"`
	TransferRequirement     string   `db:"transfer_requirement" json:"transfer_requirement"`
	ShortResignation        null.Int `db:"short_resignation" json:"short_resignation"`
	ShortResignationRemarks string   `db:"short_resignation_remarks" json:"short_resignation_remarks"`
	MedicalHistory          null.Int `db:"medical_history" json:"medical_history"`
	Nationality             null.Int `db:"nationality" json:"nationality"`
	Appearance              null.Int `db:"appearance" json:"appearance"`
	AppearanceDetail        string   `db:"appearance_detail" json:"appearance_detail"`
	Communication           null.Int `db:"communication" json:"communication"`
	CommunicationDetail     string   `db:"communication_detail" json:"communication_detail"`
	Thinking                null.Int `db:"thinking" json:"thinking"`
	SecretMemo              string   `db:"secret_memo" json:"secret_memo"`
	JobHuntingState         null.Int `db:"job_hunting_state" json:"job_hunting_state"`
	ContactManner           null.Int `db:"contact_manner" json:"contact_manner"`
	PublishToAlliance       null.Int `db:"publish_to_alliance" json:"publish_to_alliance"`
	RecommendReason         string   `db:"recommend_reason" json:"recommend_reason"`
	Phase                   null.Int `db:"phase" json:"phase"`
	InterviewDate           string   `db:"interview_date" json:"interview_date"`

	// 他テーブル
	// 求職者の学歴情報
	// StudentHistories []JobSeekerStudentHistory `db:"student_histories" json:"student_histories"`

	// //求職者の職歴
	// WorkHistories []JobSeekerWorkHistory `db:"work_histories" json:"work_histories"`

	// //求職者の所持資格
	// Licenses []JobSeekerLicense `db:"licenses" json:"licenses"`

	// //求職者の自己PR
	// SelfPromotions []JobSeekerSelfPromotion `db:"self_promotions" json:"self_promotions"`

	// //求職者の希望業界
	// DesiredIndustries []JobSeekerDesiredIndustry `db:"desired_industries" json:"desired_industries"`

	// //求職者の希望職種
	// DesiredOccupations []JobSeekerDesiredOccupation `db:"desired_occupations" json:"desired_occupations"`

	// //求職者の希望勤務地
	// DesiredWorkLocations []JobSeekerDesiredWorkLocation `db:"desired_work_locations" json:"desired_work_locations"`

	// //求職者の開発スキル
	// DevelopmentSkills []JobSeekerDevelopmentSkill `db:"development_skills" json:"development_skills"`

	// //求職者の言語スキル
	// LanguageSkills []JobSeekerLanguageSkill `db:"language_skills" json:"language_skills"`

	// //求職者のPCスキル
	// PcSkills []JobSeekerPcSkill `db:"pc_skills" json:"pc_skills"`

	// //求職者の休日タイプ
	// DesiredHolidayTypes []JobSeekerDesiredHolidayType `db:"desired_holiday_types" json:"desired_holiday_types"`
}

type DeleteJobSeekerParam struct {
	ID uint `json:"id" validate:"required"`
}

type SearchJobSeeker struct {
	FreeWord                 string
	AgentStaffID             string
	PublishToAlliance        string
	GenderTypes              []null.Int
	AgeUnder                 string
	AgeOver                  string
	DesiredIndustries        []null.Int
	DesiredOccupations       []null.Int
	DesiredWorkLocations     []null.Int
	FinalEducationTypes      []null.Int
	SchoolLevelTypes         []null.Int
	NationalityTypes         []null.Int
	MedicalHistoryTypes      []null.Int
	JobChangeTypes           []null.Int
	ShortResignationTypes    []null.Int
	UnderIncome              string
	OverIncome               string
	DesiredTransferTypes     []null.Int
	DesiredHolidayTypes      []null.Int
	DesiredCompanyScaleTypes []null.Int
	ExperienceIndustries     []null.Int
	ExperienceOccupations    []null.Int
	SocialExperiences        []null.Int
	Licenses                 []null.Int
	Languages                []null.Int
	PCSkill                  string
	AnotherPCSkills          []null.Int
	DevelopmentLanguages     []null.Int
	DevelopmentOS            []null.Int
	AppearanceTypes          []null.Int
	CommunicationTypes       []null.Int
	ThinkingTypes            []null.Int
}

func NewSearchJobSeeker(
	freeword string,
	agentStaffId string,
	publishToAlliance string,
	genderTypes []null.Int,
	ageUnder string,
	ageOver string,
	desiredIndustries []null.Int,
	desiredOccupations []null.Int,
	desiredWorkLocations []null.Int,
	finalEducationTypes []null.Int,
	schoolLevelTypes []null.Int,
	nationalityTypes []null.Int,
	medicalHistoryTypes []null.Int,
	jobChangeTypes []null.Int,
	shortResignationTypes []null.Int,
	underIncome string,
	overIncome string,
	desiredTransferTypes []null.Int,
	desiredHolidayTypes []null.Int,
	desiredCompanyScaleTypes []null.Int,
	experienceIndustries []null.Int,
	experienceOccupations []null.Int,
	socialExperiences []null.Int,
	licenses []null.Int,
	languages []null.Int,
	pcSkill string,
	anotherPCSkills []null.Int,
	developmentLanguages []null.Int,
	developmentOS []null.Int,
	appearanceTypes []null.Int,
	communicationTypes []null.Int,
	thinkingTypes []null.Int,
) *SearchJobSeeker {
	return &SearchJobSeeker{
		FreeWord:                 freeword,
		AgentStaffID:             agentStaffId,
		PublishToAlliance:        publishToAlliance,
		GenderTypes:              genderTypes,
		AgeUnder:                 ageUnder,
		AgeOver:                  ageOver,
		DesiredIndustries:        desiredIndustries,
		DesiredOccupations:       desiredOccupations,
		DesiredWorkLocations:     desiredWorkLocations,
		FinalEducationTypes:      finalEducationTypes,
		SchoolLevelTypes:         schoolLevelTypes,
		NationalityTypes:         nationalityTypes,
		MedicalHistoryTypes:      medicalHistoryTypes,
		JobChangeTypes:           jobChangeTypes,
		ShortResignationTypes:    shortResignationTypes,
		UnderIncome:              underIncome,
		OverIncome:               overIncome,
		DesiredTransferTypes:     desiredTransferTypes,
		DesiredHolidayTypes:      desiredHolidayTypes,
		DesiredCompanyScaleTypes: desiredCompanyScaleTypes,
		ExperienceIndustries:     experienceIndustries,
		ExperienceOccupations:    experienceOccupations,
		SocialExperiences:        socialExperiences,
		Licenses:                 licenses,
		Languages:                languages,
		PCSkill:                  pcSkill,
		AnotherPCSkills:          anotherPCSkills,
		DevelopmentLanguages:     developmentLanguages,
		DevelopmentOS:            developmentOS,
		AppearanceTypes:          appearanceTypes,
		CommunicationTypes:       communicationTypes,
		ThinkingTypes:            thinkingTypes,
	}
}

type LineIDParam struct {
	LineID string `db:"line_id" json:"line_id"`
}
