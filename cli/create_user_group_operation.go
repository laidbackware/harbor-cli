// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/usergroup"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUsergroupCreateUserGroupCmd returns a cmd to handle operation createUserGroup
func makeOperationUsergroupCreateUserGroupCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "createUserGroup",
		Short: `Create user group information`,
		RunE:  runOperationUsergroupCreateUserGroup,
	}

	if err := registerOperationUsergroupCreateUserGroupParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUsergroupCreateUserGroup uses cmd flags to call endpoint api
func runOperationUsergroupCreateUserGroup(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := usergroup.NewCreateUserGroupParams()
	if err, _ := retrieveOperationUsergroupCreateUserGroupXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUsergroupCreateUserGroupUsergroupFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUsergroupCreateUserGroupResult(appCli.Usergroup.CreateUserGroup(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUsergroupCreateUserGroupParamFlags registers all flags needed to fill params
func registerOperationUsergroupCreateUserGroupParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUsergroupCreateUserGroupXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUsergroupCreateUserGroupUsergroupParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUsergroupCreateUserGroupXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	xRequestIdDescription := `An unique ID for the request`

	var xRequestIdFlagName string
	if cmdPrefix == "" {
		xRequestIdFlagName = "X-Request-Id"
	} else {
		xRequestIdFlagName = fmt.Sprintf("%v.X-Request-Id", cmdPrefix)
	}

	var xRequestIdFlagDefault string

	_ = cmd.PersistentFlags().String(xRequestIdFlagName, xRequestIdFlagDefault, xRequestIdDescription)

	return nil
}
func registerOperationUsergroupCreateUserGroupUsergroupParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var usergroupFlagName string
	if cmdPrefix == "" {
		usergroupFlagName = "usergroup"
	} else {
		usergroupFlagName = fmt.Sprintf("%v.usergroup", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(usergroupFlagName, "", "Optional json string for [usergroup]. ")

	// add flags for body
	if err := registerModelUserGroupFlags(0, "userGroup", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationUsergroupCreateUserGroupXRequestIDFlag(m *usergroup.CreateUserGroupParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("X-Request-Id") {

		var xRequestIdFlagName string
		if cmdPrefix == "" {
			xRequestIdFlagName = "X-Request-Id"
		} else {
			xRequestIdFlagName = fmt.Sprintf("%v.X-Request-Id", cmdPrefix)
		}

		xRequestIdFlagValue, err := cmd.Flags().GetString(xRequestIdFlagName)
		if err != nil {
			return err, false
		}
		m.XRequestID = &xRequestIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationUsergroupCreateUserGroupUsergroupFlag(m *usergroup.CreateUserGroupParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("usergroup") {
		// Read usergroup string from cmd and unmarshal
		usergroupValueStr, err := cmd.Flags().GetString("usergroup")
		if err != nil {
			return err, false
		}

		usergroupValue := models.UserGroup{}
		if err := json.Unmarshal([]byte(usergroupValueStr), &usergroupValue); err != nil {
			return fmt.Errorf("cannot unmarshal usergroup string in models.UserGroup: %v", err), false
		}
		m.Usergroup = &usergroupValue
	}
	usergroupValueModel := m.Usergroup
	if swag.IsZero(usergroupValueModel) {
		usergroupValueModel = &models.UserGroup{}
	}
	err, added := retrieveModelUserGroupFlags(0, usergroupValueModel, "userGroup", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Usergroup = usergroupValueModel
	}
	if dryRun && debug {

		usergroupValueDebugBytes, err := json.Marshal(m.Usergroup)
		if err != nil {
			return err, false
		}
		logDebugf("Usergroup dry-run payload: %v", string(usergroupValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationUsergroupCreateUserGroupResult parses request result and return the string content
func parseOperationUsergroupCreateUserGroupResult(resp0 *usergroup.CreateUserGroupCreated, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning createUserGroupCreated is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*usergroup.CreateUserGroupBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*usergroup.CreateUserGroupUnauthorized)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*usergroup.CreateUserGroupForbidden)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*usergroup.CreateUserGroupConflict)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp5 interface{} = respErr
		resp5, ok := iResp5.(*usergroup.CreateUserGroupInternalServerError)
		if ok {
			if !swag.IsZero(resp5) && !swag.IsZero(resp5.Payload) {
				msgStr, err := json.Marshal(resp5.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response createUserGroupCreated is not supported by go-swagger cli yet.

	return "", nil
}