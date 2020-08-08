package service

import (
	"build-version/model/data"
	"build-version/model/request"
	"build-version/service/repository"
	"github.com/google/uuid"
	"time"
)

func CreateOrganisation(request request.CreateOrg) (*data.OrganisationData, error){
	if db, err := connectDb(); err != nil {
		return nil, err
	} else {
		record := data.OrganisationData{
			Id:        uuid.New().String(),
			Name:      request.Name,
			CreatedTs: time.Time{},
			UpdatedTs: time.Time{},
			PlanType:  request.PlanType,
		}
		if err = repository.CreateOrganisation(db, &record); err != nil {
			return nil, err
		} else {
			return &record, nil
		}
	}
}

func FindOrganisation(orgId string) (*data.OrganisationData, error)  {
	if db, err := connectDb(); err != nil {
		return nil, err
	} else {
		defer db.Close()
		if record, err := repository.GetOrganisationById(db, orgId); err != nil {
			return nil, err
		} else {
			return record, nil
		}
	}
}

func CreateProject(request *request.CreateProject) (*data.ProjectData, error) {
	if db, err := connectDb(); err != nil {
		return nil, err
	} else {
		record := &data.ProjectData{
			Id: uuid.New().String(),
			Name: request.Name,
			Organisation: request.Organisation,
			AccessToken: uuid.New().String(),
		}

		if err = repository.CreateProject(db, record); err != nil {
			return nil, err
		} else {
			return record, nil
		}
	}
}

func FindProject(projectId string) (*data.ProjectData, error) {
	if db, err := connectDb(); err != nil {
		return nil, err
	} else {
		if record, err := repository.GetProjectById(db, projectId); err != nil {
			return nil, err
		} else {
			return record, nil
		}
	}
}

func RegenerateProjectToken(projectId string) (*data.ProjectData, error) {
	if db, err := connectDb(); err != nil {
		return nil, err
	} else {
		if record, err := repository.GetProjectById(db, projectId); err != nil {
			return nil, err
		} else {
			record.AccessToken = uuid.New().String()
			if err = repository.UpdateProject(db, record); err != nil {
				return nil, err
			} else {
				return record, nil
			}
		}
	}
}