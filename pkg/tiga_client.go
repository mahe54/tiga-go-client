package tigaclient

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// New returns a Tiga Client
// Expects that the following os.environment variables to be exported before use:
// TIGA_HOST=https://api.tiga-sandbox.teliacompany.net
// For SIDM Login (example values provided)
// SIDM_HOST=https://staging.securityservice.teliacompany.com
// SIDM_SECRET=0b6ad3d80d060fa0c6317673b38cad59d8a249611444ecd63c2919fda9ed358dd89f68b6e5feaa2a522b58a578fc0fa925e5a28c101244cbaa82f90f4540e999
// SIDM_SERVICEID=2bee5a4c-6070-4b37-b13f-689e27a4d2a8
func New(caller CallerInterface) (*Client, error) {

	tigaHost := os.Getenv("TIGA_HOST")
	if tigaHost == "" {
		return nil, errors.New("TIGA_HOST environment variable not set")
	}
	c := Client{
		tigaURL: tigaHost,
	}
	c.Caller = caller

	err := c.loginSIDM()
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Client) loginSIDM() error {
	sidmHost := os.Getenv("SIDM_HOST")
	if sidmHost == "" {
		return errors.New("SIDM_HOST environment variable not set")
	}
	sidmServiceID := os.Getenv("SIDM_SERVICEID")
	if sidmServiceID == "" {
		return errors.New("SIDM_SERVICEID environment variable not set")
	}
	sidmSecret := os.Getenv("SIDM_SECRET")
	if sidmSecret == "" {
		return errors.New("SIDM_SECRET environment variable not set")
	}

	encoded := base64.URLEncoding.EncodeToString([]byte(sidmServiceID + ":" + sidmSecret))

	u, _ := url.Parse(sidmHost + "/oauth/token")
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("resource", "https://api.tiga-sandbox.teliacompany.net/v1/")
	data.Set("token_type", "jwt")
	body := strings.NewReader(data.Encode())

	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Basic "+encoded)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.Caller.DoCall(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	var restoken bytes.Buffer
	for scanner.Scan() {
		restoken.Write(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	jwtToken := jwtToken{}
	err = json.Unmarshal(restoken.Bytes(), &jwtToken)
	if err != nil {
		return err
	}

	c.token = &jwtToken
	return nil
}

func (c *Client) GetRole(hid, name string) (*Role, error) {

	baseUrl := "/v1/userRoles?systemId="
	roleName := "&roleName=" + name
	url := c.tigaURL + baseUrl + hid + roleName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.Caller.DoCall(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, &TigaError{
			StatusCode: res.StatusCode,
			Message:    string(body),
		}
	}

	resRoles, errs := io.ReadAll(res.Body)
	if errs != nil {
		return nil, err
	}

	roles := []Role{}
	err = json.Unmarshal(resRoles, &roles)
	if err != nil {
		return nil, err
	}

	return &roles[0], nil
}

func (c *Client) CreateRole(r *Role) (*Role, error) {
	path := "/v1/userRoles"
	u, _ := url.Parse(c.tigaURL + path)
	q := u.Query()
	q.Set("minApprovalLevel", "manager")

	if r.Template != "" {
		q.Set("namingTemplate", strings.TrimLeft(strings.TrimRight(r.Template, `"`), `"`))
	}
	u.RawQuery = q.Encode()

	roleJson, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", u.String(), bytes.NewBuffer(roleJson))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.Caller.DoCall(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		scanner := bufio.NewScanner(res.Body)
		var body bytes.Buffer
		for scanner.Scan() {
			body.Write(scanner.Bytes())
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		return nil, &TigaError{
			StatusCode: res.StatusCode,
			Message:    body.String(),
		}
	}

	scanner := bufio.NewScanner(res.Body)
	var resBody bytes.Buffer
	for scanner.Scan() {
		resBody.Write(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	role := Role{}
	err = json.Unmarshal(resBody.Bytes(), &role)
	if err != nil {
		return nil, err
	}

	return &role, nil
}
