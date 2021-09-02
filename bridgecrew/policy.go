package bridgecrew

//import (
//"encoding/json"
////"errors"
//"fmt"
//"net/http"
//"strings"
//)

// GetPolicy - Returns a specifc Policy
//func (c *Client) GetPolicy(PolicyID string) (*Policy, error) {
//	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Policys/%s", c.HostURL, PolicyID), nil)
//	if err != nil {
//		return nil, err
//	}
//
//	body, err := c.doRequest(req)
//	if err != nil {
//		return nil, err
//	}
//
//	Policy := Policy{}
//	err = json.Unmarshal(body, &Policy)
//	if err != nil {
//		return nil, err
//	}
//
//	return &Policy, nil
//}

// CreatePolicy - Create new Policy
//func CreatePolicy(Policies []Policy) (*Policy, error) {
//	rb, err := json.Marshal(Policies)
//	if err != nil {
//		return nil, err
//	}
//    HostURL=""
//	req, err := http.NewRequest("POST", fmt.Sprintf("%s/Policys", HostURL), strings.NewReader(string(rb)))
//	if err != nil {
//		return nil, err
//	}
//
//	body, err := c.doRequest(req)
//	if err != nil {
//		return nil, err
//	}
//
//	Policy := Policy{}
//	err = json.Unmarshal(body, &Policy)
//	if err != nil {
//		return nil, err
//	}
//
//	return &Policy, nil
//}

//// UpdatePolicy - Updates an Policy
//func (c *Client) UpdatePolicy(PolicyID string, PolicyItems []PolicyItem) (*Policy, error) {
//	rb, err := json.Marshal(PolicyItems)
//	if err != nil {
//		return nil, err
//	}
//
//	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/Policys/%s", c.HostURL, PolicyID), strings.NewReader(string(rb)))
//	if err != nil {
//		return nil, err
//	}
//
//	body, err := c.doRequest(req)
//	if err != nil {
//		return nil, err
//	}
//
//	Policy := Policy{}
//	err = json.Unmarshal(body, &Policy)
//	if err != nil {
//		return nil, err
//	}
//
//	return &Policy, nil
//}
//
//// DeletePolicy - Deletes an Policy
//func (c *Client) DeletePolicy(PolicyID string) error {
//	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/Policys/%s", c.HostURL, PolicyID), nil)
//	if err != nil {
//		return err
//	}
//
//	body, err := c.doRequest(req)
//	if err != nil {
//		return err
//	}
//
//	if string(body) != "Deleted Policy" {
//		return errors.New(string(body))
//	}
//
//	return nil
//}
