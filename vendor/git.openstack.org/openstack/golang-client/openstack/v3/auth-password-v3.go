// auth-password - Username/Password Authentication
// Copyright 2015 Dean Troyer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v3

import (
	"encoding/json"
	"errors"
	// "strings"
)

type Domain struct {
	Id string `json:"id"`
	// Name string `json:"name"`
}

type User struct {
	Username string `json:"name"`
	Password string `json:"password"`
	Domain   `json:"domain"`
}

type PasswordCredentials struct {
	User `json:"user"`
}

type IdentityEntry struct {
	Methods             []string `json:"methods"`
	PasswordCredentials `json:"password"`
}

type Project struct {
	Id string `json:"id"`
}

type ScopeEntry struct {
	Project `json:"project"`
}

type OSAuth struct {
	IdentityEntry `json:"identity"`
	ScopeEntry    `json:"scope"`
}

type UserPassV3 struct {
	AuthUrl string `json:"-"`
	OSAuth  `json:"auth"`
}

func NewUserPassV3(ao AuthOpts) (*UserPassV3, error) {
	// Validate incoming values
	if ao.AuthUrl == "" {
		err := errors.New("AuthUrl required")
		return nil, err
	}
	if ao.Methods[0] != "password" {
		err := errors.New("Password method required")
		return nil, err
	}
	if ao.Username == "" {
		err := errors.New("Username required")
		return nil, err
	}
	if ao.Password == "" {
		err := errors.New("Password required")
		return nil, err
	}

	upv3 := &UserPassV3{
		AuthUrl: ao.AuthUrl,
		OSAuth: OSAuth{
			IdentityEntry: IdentityEntry{
				Methods: ao.Methods,
				PasswordCredentials: PasswordCredentials{
					User: User{
						Domain: Domain{
							Id: "default",
						},
						Username: ao.Username,
						Password: ao.Password,
					},
				},
			},
			ScopeEntry: ScopeEntry{
				Project: Project{
					Id: ao.ProjectId,
				},
			},
		},
	}
	return upv3, nil
}

// Produce JSON output
func (s *UserPassV3) JSON() []byte {
	reqAuth, err := json.Marshal(s)
	if err != nil {
		// Return an empty structure
		reqAuth = []byte{'{', '}'}
	}
	return reqAuth
}
