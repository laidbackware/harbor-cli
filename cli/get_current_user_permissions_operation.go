// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/user"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUserGetCurrentUserPermissionsCmd returns a cmd to handle operation getCurrentUserPermissions
func makeOperationUserGetCurrentUserPermissionsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getCurrentUserPermissions",
		Short: ``,
		RunE:  runOperationUserGetCurrentUserPermissions,
	}

	if err := registerOperationUserGetCurrentUserPermissionsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUserGetCurrentUserPermissions uses cmd flags to call endpoint api
func runOperationUserGetCurrentUserPermissions(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := user.NewGetCurrentUserPermissionsParams()
	if err, _ := retrieveOperationUserGetCurrentUserPermissionsXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserGetCurrentUserPermissionsRelativeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserGetCurrentUserPermissionsScopeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUserGetCurrentUserPermissionsResult(appCli.User.GetCurrentUserPermissions(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUserGetCurrentUserPermissionsParamFlags registers all flags needed to fill params
func registerOperationUserGetCurrentUserPermissionsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUserGetCurrentUserPermissionsXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserGetCurrentUserPermissionsRelativeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserGetCurrentUserPermissionsScopeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUserGetCurrentUserPermissionsXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationUserGetCurrentUserPermissionsRelativeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	relativeDescription := `If true, the resources in the response are relative to the scope,
eg for resource '/project/1/repository' if relative is 'true' then the resource in response will be 'repository'.
`

	var relativeFlagName string
	if cmdPrefix == "" {
		relativeFlagName = "relative"
	} else {
		relativeFlagName = fmt.Sprintf("%v.relative", cmdPrefix)
	}

	var relativeFlagDefault bool

	_ = cmd.PersistentFlags().Bool(relativeFlagName, relativeFlagDefault, relativeDescription)

	return nil
}
func registerOperationUserGetCurrentUserPermissionsScopeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	scopeDescription := `The scope for the permission`

	var scopeFlagName string
	if cmdPrefix == "" {
		scopeFlagName = "scope"
	} else {
		scopeFlagName = fmt.Sprintf("%v.scope", cmdPrefix)
	}

	var scopeFlagDefault string

	_ = cmd.PersistentFlags().String(scopeFlagName, scopeFlagDefault, scopeDescription)

	return nil
}

func retrieveOperationUserGetCurrentUserPermissionsXRequestIDFlag(m *user.GetCurrentUserPermissionsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationUserGetCurrentUserPermissionsRelativeFlag(m *user.GetCurrentUserPermissionsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("relative") {

		var relativeFlagName string
		if cmdPrefix == "" {
			relativeFlagName = "relative"
		} else {
			relativeFlagName = fmt.Sprintf("%v.relative", cmdPrefix)
		}

		relativeFlagValue, err := cmd.Flags().GetBool(relativeFlagName)
		if err != nil {
			return err, false
		}
		m.Relative = &relativeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationUserGetCurrentUserPermissionsScopeFlag(m *user.GetCurrentUserPermissionsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("scope") {

		var scopeFlagName string
		if cmdPrefix == "" {
			scopeFlagName = "scope"
		} else {
			scopeFlagName = fmt.Sprintf("%v.scope", cmdPrefix)
		}

		scopeFlagValue, err := cmd.Flags().GetString(scopeFlagName)
		if err != nil {
			return err, false
		}
		m.Scope = &scopeFlagValue

	}
	return nil, retAdded
}

// parseOperationUserGetCurrentUserPermissionsResult parses request result and return the string content
func parseOperationUserGetCurrentUserPermissionsResult(resp0 *user.GetCurrentUserPermissionsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*user.GetCurrentUserPermissionsOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning getCurrentUserPermissionsUnauthorized is not supported

		// Non schema case: warning getCurrentUserPermissionsInternalServerError is not supported

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
