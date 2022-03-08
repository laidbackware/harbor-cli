// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/user"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUserUpdateUserPasswordCmd returns a cmd to handle operation updateUserPassword
func makeOperationUserUpdateUserPasswordCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "updateUserPassword",
		Short: `This endpoint is for user to update password. Users with the admin role can change any user's password. Regular users can change only their own password.
`,
		RunE: runOperationUserUpdateUserPassword,
	}

	if err := registerOperationUserUpdateUserPasswordParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUserUpdateUserPassword uses cmd flags to call endpoint api
func runOperationUserUpdateUserPassword(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := user.NewUpdateUserPasswordParams()
	if err, _ := retrieveOperationUserUpdateUserPasswordXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserUpdateUserPasswordPasswordFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserUpdateUserPasswordUserIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUserUpdateUserPasswordResult(appCli.User.UpdateUserPassword(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUserUpdateUserPasswordParamFlags registers all flags needed to fill params
func registerOperationUserUpdateUserPasswordParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUserUpdateUserPasswordXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserUpdateUserPasswordPasswordParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserUpdateUserPasswordUserIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUserUpdateUserPasswordXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationUserUpdateUserPasswordPasswordParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var passwordFlagName string
	if cmdPrefix == "" {
		passwordFlagName = "password"
	} else {
		passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(passwordFlagName, "", "Optional json string for [password]. Password to be updated, the attribute old_password is optional when the API is called by the system administrator.")

	// add flags for body
	if err := registerModelPasswordReqFlags(0, "passwordReq", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationUserUpdateUserPasswordUserIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	userIdDescription := `Required. `

	var userIdFlagName string
	if cmdPrefix == "" {
		userIdFlagName = "user_id"
	} else {
		userIdFlagName = fmt.Sprintf("%v.user_id", cmdPrefix)
	}

	var userIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(userIdFlagName, userIdFlagDefault, userIdDescription)

	return nil
}

func retrieveOperationUserUpdateUserPasswordXRequestIDFlag(m *user.UpdateUserPasswordParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationUserUpdateUserPasswordPasswordFlag(m *user.UpdateUserPasswordParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("password") {
		// Read password string from cmd and unmarshal
		passwordValueStr, err := cmd.Flags().GetString("password")
		if err != nil {
			return err, false
		}

		passwordValue := models.PasswordReq{}
		if err := json.Unmarshal([]byte(passwordValueStr), &passwordValue); err != nil {
			return fmt.Errorf("cannot unmarshal password string in models.PasswordReq: %v", err), false
		}
		m.Password = &passwordValue
	}
	passwordValueModel := m.Password
	if swag.IsZero(passwordValueModel) {
		passwordValueModel = &models.PasswordReq{}
	}
	err, added := retrieveModelPasswordReqFlags(0, passwordValueModel, "passwordReq", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Password = passwordValueModel
	}
	if dryRun && debug {

		passwordValueDebugBytes, err := json.Marshal(m.Password)
		if err != nil {
			return err, false
		}
		logDebugf("Password dry-run payload: %v", string(passwordValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationUserUpdateUserPasswordUserIDFlag(m *user.UpdateUserPasswordParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("user_id") {

		var userIdFlagName string
		if cmdPrefix == "" {
			userIdFlagName = "user_id"
		} else {
			userIdFlagName = fmt.Sprintf("%v.user_id", cmdPrefix)
		}

		userIdFlagValue, err := cmd.Flags().GetInt64(userIdFlagName)
		if err != nil {
			return err, false
		}
		m.UserID = userIdFlagValue

	}
	return nil, retAdded
}

// parseOperationUserUpdateUserPasswordResult parses request result and return the string content
func parseOperationUserUpdateUserPasswordResult(resp0 *user.UpdateUserPasswordOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning updateUserPasswordOK is not supported

		// Non schema case: warning updateUserPasswordBadRequest is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*user.UpdateUserPasswordUnauthorized)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning updateUserPasswordForbidden is not supported

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*user.UpdateUserPasswordInternalServerError)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response updateUserPasswordOK is not supported by go-swagger cli yet.

	return "", nil
}
