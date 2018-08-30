package statemachines

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// DeleteInput provides the data required to delete a state machine.
type DeleteInput struct {
	ID      string
	Version int
}

// UpdateInput provides the data required to update a state machine.
type UpdateInput struct {
	ID string

	// The expected version of the state on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict
	// will be returned.
	Version int

	// The list of update actions to be performed on the state.
	Actions commercetools.UpdateActions
}

// GetByID will return a state matching the provided ID. OAuth2 Scopes:
// view_states:{projectKey} (or, deprecated: view_orders:{projectKey})
func (svc *Service) GetByID(id string) (result *State, err error) {
	err = svc.client.Get(fmt.Sprintf("states/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create will create a new state from a draft, and return the newly created
// state. OAuth2 Scopes: manage_states:{projectKey} (or, deprecated:
// manage_orders:{projectKey})
func (svc *Service) Create(draft *StateDraft) (result *State, err error) {
	err = svc.client.Create("states", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update will update a state matching the provided ID with the defined
// UpdateActions. OAuth2 Scopes: manage_states:{projectKey} (or, deprecated:
// manage_orders:{projectKey})
func (svc *Service) Update(input *UpdateInput) (result *State, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid state id passed")
	}

	endpoint := fmt.Sprintf("states/%s", input.ID)
	err = svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByID will delete a state matching the provided ID. OAuth2 Scopes:
// manage_states:{projectKey} (or, deprecated: manage_orders:{projectKey})
func (svc *Service) DeleteByID(id string, version int) (result *State, err error) {
	endpoint := fmt.Sprintf("states/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
