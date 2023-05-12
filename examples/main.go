// This file tests the public facing functions only
package main

import (
	"fmt"
	"time"

	tc "github.com/telia-company/tiga-go-client/pkg"
)

func main() {

	//Example data
	HID := "Hid100000006"
	roleName := "AWS_91234998890125_Administrator_Role"

	mockCalls := false
	//Depending on if you are mocking or doing a normal call you could choose to go either way
	//The NewClient method acts as a factory for the client where you choose to pass in the CallerInterface implementation of your choice
	//See example in cmdbmock.go -file
	var ric tc.CallerInterface
	if mockCalls {
		//Needs to match what is exported as an env variable, will be removed/parsed in mock
		ric = &MockCaller{BaseURL: "https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/"}
	} else {
		ric = &tc.Caller{}
	}

	client, err := tc.New(ric, false)
	if err != nil {
		fmt.Print(err.Error())
	}

	// Get Role from Tiga (Lookup)
	role, err := client.GetRole(HID, roleName)
	if err != nil {
		fmt.Printf("\nError: %+v \n\n", err.Error())
	} else {
		fmt.Printf("\nRecieved role: %+v \n\n", role)
	}

	role, err = client.GetRole(HID, roleName+"name_that_doesnt_exist")
	if err != nil {
		fmt.Printf("\nError: %+v \n\n", err.Error())
	} else {
		fmt.Printf("\nRecieved role: %+v \n\n", role)
	}

	//(fail to) Create role in Tiga
	newRole := &tc.Role{
		Name:               "AWS_91234000001225_Administrator_Role",
		Template:           "Amazon Web Services Cloud (AWS)",
		ValidFrom:          "2023-05-03T01:01:01Z",
		ValidTo:            "2024-04-25T01:01:01Z",
		PreventSelfService: false,
		Description:        "Gives access to admin parts of AWS",
		SystemInstance:     "/v1/systems/HID100000006/instances/HID100000006.TEST",
		ProvisioningType:   "activeDirectory",
		Owners:             []string{"zkv293", "mdr449", "nju840"},
		ApprovalSettings: tc.ApprovalSettings{
			SkipSystemOwnerApproval:    true,
			SkipManagerApproval:        true,
			SkipRoleOwnerApproval:      true,
			NamedApprovers:             []string{"sbh881", "zkv293"},
			SecurityClearanceApprovers: []string{"nju840", "zkv293"},
		},
		UserRequirements: tc.UserRequirements{
			DigitalCommittment: true,
			TermsAndConditions: "/v1/termsAndConditions/Terms+and+Conditions+Jfrog",
			Countries:          []string{"SE", "NO", "DK"},
			BusinessContexts:   []string{"/v1/businessContexts/Any"},
		},
	}

	createdRole, err := client.CreateRole(newRole)
	if err != nil {
		fmt.Printf("\nError: %+v \n\n", err.Error())
	} else {
		fmt.Printf("\nCreated role: %+v \n\n", createdRole)
	}

	//Create role with unixtime in name so it allways differs and therefore gets created :)
	currentTime := time.Now().Unix()
	newRole = &tc.Role{
		Name:               "AWS_91234998891225_Administrator_Role_" + fmt.Sprint(currentTime),
		Template:           "Amazon Web Services Cloud (AWS)",
		ValidFrom:          "2023-05-03T01:01:01Z",
		ValidTo:            "2024-04-25T01:01:01Z",
		PreventSelfService: false,
		Description:        "Gives access to admin parts of AWS",
		SystemInstance:     "/v1/systems/HID100000006/instances/HID100000006.TEST",
		ProvisioningType:   "activeDirectory",
		Owners:             []string{"zkv293", "mdr449", "nju840"},
		ApprovalSettings: tc.ApprovalSettings{
			SkipSystemOwnerApproval:    true,
			SkipManagerApproval:        true,
			SkipRoleOwnerApproval:      true,
			NamedApprovers:             []string{"sbh881", "zkv293"},
			SecurityClearanceApprovers: []string{"nju840", "zkv293"},
		},
		UserRequirements: tc.UserRequirements{
			DigitalCommittment: true,
			TermsAndConditions: "/v1/termsAndConditions/Terms+and+Conditions+Jfrog",
			Countries:          []string{"SE", "NO", "DK"},
			BusinessContexts:   []string{"/v1/businessContexts/Any"},
		},
	}

	createdRole, err = client.CreateRole(newRole)
	if err != nil {
		fmt.Printf("\nError: %+v \n\n", err.Error())
	} else {
		fmt.Printf("\nCreated role: %+v \n\n", createdRole)
	}
}
