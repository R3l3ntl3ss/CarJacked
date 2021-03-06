package mongo

import (
	"github.com/R3l3ntl3ss/CarJacked/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AssignCaseToOfficer : Assign case to unassigned officer
func (m Mongo) AssignCaseToOfficer(caseID primitive.ObjectID) (assigned bool, officer models.Officer, err error) {

	officers, err := m.GetUnassignedOfficers()

	// Check for Error And No Assigned Officers
	if err != nil {
		return false, models.Officer{}, err
	} else if (len(officers)) == 0 {
		return false, models.Officer{}, nil
	}

	officer = officers[0]

	err = m.UpdateCaseWithOfficerID(caseID, officer.ID)

	if err != nil {
		return false, models.Officer{}, err
	}

	err = m.MakeOfficerAssigned(officer.ID)

	if err != nil {
		return false, models.Officer{}, err
	}

	return true, officer, nil
}

// AssignOfficerToCase : Assign case to unassigned officer
func (m Mongo) AssignOfficerToCase(officerID string) ( bool,  models.Case,  error) {

	// Convert case ID string to primitive Object
	officerObjID, err := primitive.ObjectIDFromHex(officerID)

	if err != nil {
		return false, models.Case{}, err
	}

	cases, err := m.GetAllUnassignedCases()

	// Check for Error And No Assigned Officers
	if err != nil {
		return false, models.Case{}, nil
	} else if (len(cases)) == 0 {
		return false, models.Case{}, nil
	}

	oneCase := cases[0]

	err = m.UpdateCaseWithOfficerID(oneCase.ID, officerObjID)

	if err != nil {
		return false, models.Case{}, err
	}

	err = m.MakeOfficerAssigned(officerObjID)

	if err != nil {
		return false, models.Case{}, err
	}

	return true, oneCase, nil
}
