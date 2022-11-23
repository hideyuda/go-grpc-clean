package responses

import (
	"github.com/hidenari-yuda/umerun-resume/domain/entity"
)

type JobSeeker struct {
	JobSeeker *entity.JobSeeker `json:"job_seeker"`
}

func NewJobSeeker(jobSeeker *entity.JobSeeker) JobSeeker {
	return JobSeeker{
		JobSeeker: jobSeeker,
	}
}

type JobSeekerList struct {
	JobSeekerList []*entity.JobSeeker `json:"job_seeker_list"`
}

func NewJobSeekerList(jobSeekers []*entity.JobSeeker) JobSeekerList {
	return JobSeekerList{
		JobSeekerList: jobSeekers,
	}
}
