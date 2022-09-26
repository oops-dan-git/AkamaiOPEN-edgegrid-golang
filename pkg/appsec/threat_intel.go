package appsec

import (
	"context"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// The ThreatIntel interface supports retrieving and modifying the operational settings for adaptive intelligence.
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#threatintel
	ThreatIntel interface {
		// GetThreatIntel retrieves the current threat intel settings.
		GetThreatIntel(ctx context.Context, params GetThreatIntelRequest) (*GetThreatIntelResponse, error)

		// UpdateThreatIntel modifies the current threat intel settings.
		UpdateThreatIntel(ctx context.Context, params UpdateThreatIntelRequest) (*UpdateThreatIntelResponse, error)
	}

	// GetThreatIntelRequest is used to retrieve the threat intel settings.
	GetThreatIntelRequest struct {
		ConfigID int    `json:"-"`
		Version  int    `json:"-"`
		PolicyID string `json:"-"`
	}

	// GetThreatIntelResponse is returned from a call to GetThreatIntel.
	GetThreatIntelResponse struct {
		ThreatIntel string `json:"threatIntel,omitempty"`
	}

	// UpdateThreatIntelRequest is used to update the threat intel settings.
	UpdateThreatIntelRequest struct {
		ConfigID    int    `json:"-"`
		Version     int    `json:"-"`
		PolicyID    string `json:"-"`
		ThreatIntel string `json:"threatIntel"`
	}

	// UpdateThreatIntelResponse is returned from a call to UpdateThreatIntel.
	UpdateThreatIntelResponse struct {
		ThreatIntel string `json:"threatIntel,omitempty"`
	}
)

// Validate validates a GetAttackGroupConditionExceptionRequest.
func (v GetThreatIntelRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"PolicyID": validation.Validate(v.PolicyID, validation.Required),
	}.Filter()
}

// Validate validates an UpdateAttackGroupConditionExceptionRequest.
func (v UpdateThreatIntelRequest) Validate() error {
	return validation.Errors{
		"ConfigID": validation.Validate(v.ConfigID, validation.Required),
		"Version":  validation.Validate(v.Version, validation.Required),
		"PolicyID": validation.Validate(v.PolicyID, validation.Required),
	}.Filter()
}

func (p *appsec) GetThreatIntel(ctx context.Context, params GetThreatIntelRequest) (*GetThreatIntelResponse, error) {
	logger := p.Log(ctx)
	logger.Debug("GetThreatIntel")

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	var result GetThreatIntelResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/security-policies/%s/threat-intel",
		params.ConfigID,
		params.Version,
		params.PolicyID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetThreatIntel request: %w", err)
	}
	logger.Debugf("BEFORE GetThreatIntel %v", result)
	resp, err := p.Exec(req, &result)
	if err != nil {
		return nil, fmt.Errorf("GetThreatIntel  request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}
	logger.Debugf("GetThreatIntel %v", result)
	return &result, nil

}

func (p *appsec) UpdateThreatIntel(ctx context.Context, params UpdateThreatIntelRequest) (*UpdateThreatIntelResponse, error) {
	logger := p.Log(ctx)
	logger.Debug("UpdateThreatIntel")

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/security-policies/%s/threat-intel",
		params.ConfigID,
		params.Version,
		params.PolicyID,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create UpdateThreatIntel request: %w", err)
	}

	var result UpdateThreatIntelResponse
	resp, err := p.Exec(req, &result, params)
	if err != nil {
		return nil, fmt.Errorf("UpdateThreatIntel request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &result, nil
}
