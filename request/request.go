package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/agungcandra/gojira/entity"
	"github.com/pkg/errors"
)

const TimeoutLimit = 10
const ContentTypeJSON = "application/json"

type Request struct {
	host   string
	client *http.Client
}

func NewRequest(host string) *Request {
	transport := http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	client := http.Client{
		Timeout:   time.Second * TimeoutLimit,
		Transport: &transport,
	}

	return &Request{
		host:   host,
		client: &client,
	}
}

func (r *Request) GetAvailableTransition(issue string, credential *entity.Credential) ([]entity.Transition, error) {
	url := fmt.Sprintf("%s/rest/api/2/issue/%s/transitions", r.host, issue)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(credential.Username, credential.Password)

	response, err := r.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "request")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		io.Copy(io.Discard, response.Body)
	}

	var result entity.TransitionResponse
	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.Wrap(err, "decode")
	}

	return result.Transitions, nil
}

func (r *Request) UpdateJiraState(issue string, transition *entity.Transition, credential *entity.Credential) error {
	url := fmt.Sprintf("%s/rest/api/2/issue/%s/transitions", r.host, issue)
	payload := NewTransitionRequest(transition.ID)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.SetBasicAuth(credential.Username, credential.Password)
	req.Header.Set("Content-Type", ContentTypeJSON)

	response, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	io.Copy(io.Discard, response.Body)

	if response.StatusCode == http.StatusNoContent || response.StatusCode == http.StatusOK {
		return nil
	}

	return errors.New(fmt.Sprintf("invalid status code %v", response.StatusCode))
}

func (r *Request) GetIssue(issue string, credential *entity.Credential) (*entity.Issue, error) {
	url := fmt.Sprintf("%s/rest/api/2/issue/%s", r.host, issue)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("fields", "status")

	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth(credential.Username, credential.Password)

	response, err := r.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "request")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		io.Copy(io.Discard, response.Body)
	}

	var result entity.Issue
	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.Wrap(err, "decode")
	}

	return &result, nil
}
