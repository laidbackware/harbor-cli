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

// makeOperationUsergroupUpdateUserGroupCmd returns a cmd to handle operation updateUserGroup
func makeOperationUsergroupUpdateUserGroupCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "updateUserGroup",
		Short: `Update user group information`,
		RunE:  runOperationUsergroupUpdateUserGroup,
	}

	if err := registerOperationUsergroupUpdateUserGroupParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUsergroupUpdateUserGroup uses cmd flags to call endpoint api
func runOperationUsergroupUpdateUserGroup(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := usergroup.NewUpdateUserGroupParams()
	if err, _ := retrieveOperationUsergroupUpdateUserGroupXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUsergroupUpdateUserGroupGroupIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUsergroupUpdateUserGroupUsergroupFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUsergroupUpdateUserGroupResult(appCli.Usergroup.UpdateUserGroup(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUsergroupUpdateUserGroupParamFlags registers all flags needed to fill params
func registerOperationUsergroupUpdateUserGroupParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUsergroupUpdateUserGroupXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUsergroupUpdateUserGroupGroupIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUsergroupUpdateUserGroupUsergroupParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUsergroupUpdateUserGroupXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationUsergroupUpdateUserGroupGroupIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	groupIdDescription := `Required. Group ID`

	var groupIdFlagName string
	if cmdPrefix == "" {
		groupIdFlagName = "group_id"
	} else {
		groupIdFlagName = fmt.Sprintf("%v.group_id", cmdPrefix)
	}

	var groupIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(groupIdFlagName, groupIdFlagDefault, groupIdDescription)

	return nil
}
func registerOperationUsergroupUpdateUserGroupUsergroupParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationUsergroupUpdateUserGroupXRequestIDFlag(m *usergroup.UpdateUserGroupParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationUsergroupUpdateUserGroupGroupIDFlag(m *usergroup.UpdateUserGroupParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("group_id") {

		var groupIdFlagName string
		if cmdPrefix == "" {
			groupIdFlagName = "group_id"
		} else {
			groupIdFlagName = fmt.Sprintf("%v.group_id", cmdPrefix)
		}

		groupIdFlagValue, err := cmd.Flags().GetInt64(groupIdFlagName)
		if err != nil {
			return err, false
		}
		m.GroupID = groupIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationUsergroupUpdateUserGroupUsergroupFlag(m *usergroup.UpdateUserGroupParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationUsergroupUpdateUserGroupResult parses request result and return the string content
func parseOperationUsergroupUpdateUserGroupResult(resp0 *usergroup.UpdateUserGroupOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning updateUserGroupOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*usergroup.UpdateUserGroupBadRequest)
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
		resp2, ok := iResp2.(*usergroup.UpdateUserGroupUnauthorized)
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
		resp3, ok := iResp3.(*usergroup.UpdateUserGroupForbidden)
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
		resp4, ok := iResp4.(*usergroup.UpdateUserGroupNotFound)
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
		resp5, ok := iResp5.(*usergroup.UpdateUserGroupInternalServerError)
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

	// warning: non schema response updateUserGroupOK is not supported by go-swagger cli yet.

	return "", nil
}
