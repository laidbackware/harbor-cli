// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/robotv1"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationRobotv1UpdateRobotV1Cmd returns a cmd to handle operation updateRobotV1
func makeOperationRobotv1UpdateRobotV1Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "UpdateRobotV1",
		Short: `Used to disable/enable a specified robot account.`,
		RunE:  runOperationRobotv1UpdateRobotV1,
	}

	if err := registerOperationRobotv1UpdateRobotV1ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationRobotv1UpdateRobotV1 uses cmd flags to call endpoint api
func runOperationRobotv1UpdateRobotV1(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := robotv1.NewUpdateRobotV1Params()
	if err, _ := retrieveOperationRobotv1UpdateRobotV1XIsResourceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRobotv1UpdateRobotV1XRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRobotv1UpdateRobotV1ProjectNameOrIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRobotv1UpdateRobotV1RobotFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRobotv1UpdateRobotV1RobotIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationRobotv1UpdateRobotV1Result(appCli.Robotv1.UpdateRobotV1(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationRobotv1UpdateRobotV1ParamFlags registers all flags needed to fill params
func registerOperationRobotv1UpdateRobotV1ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationRobotv1UpdateRobotV1XIsResourceNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRobotv1UpdateRobotV1XRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRobotv1UpdateRobotV1ProjectNameOrIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRobotv1UpdateRobotV1RobotParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRobotv1UpdateRobotV1RobotIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationRobotv1UpdateRobotV1XIsResourceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	xIsResourceNameDescription := `The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.`

	var xIsResourceNameFlagName string
	if cmdPrefix == "" {
		xIsResourceNameFlagName = "X-Is-Resource-Name"
	} else {
		xIsResourceNameFlagName = fmt.Sprintf("%v.X-Is-Resource-Name", cmdPrefix)
	}

	var xIsResourceNameFlagDefault bool

	_ = cmd.PersistentFlags().Bool(xIsResourceNameFlagName, xIsResourceNameFlagDefault, xIsResourceNameDescription)

	return nil
}
func registerOperationRobotv1UpdateRobotV1XRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationRobotv1UpdateRobotV1ProjectNameOrIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	projectNameOrIdDescription := `Required. The name or id of the project`

	var projectNameOrIdFlagName string
	if cmdPrefix == "" {
		projectNameOrIdFlagName = "project_name_or_id"
	} else {
		projectNameOrIdFlagName = fmt.Sprintf("%v.project_name_or_id", cmdPrefix)
	}

	var projectNameOrIdFlagDefault string

	_ = cmd.PersistentFlags().String(projectNameOrIdFlagName, projectNameOrIdFlagDefault, projectNameOrIdDescription)

	return nil
}
func registerOperationRobotv1UpdateRobotV1RobotParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var robotFlagName string
	if cmdPrefix == "" {
		robotFlagName = "robot"
	} else {
		robotFlagName = fmt.Sprintf("%v.robot", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(robotFlagName, "", "Optional json string for [robot]. The JSON object of a robot account.")

	// add flags for body
	if err := registerModelRobotFlags(0, "robot", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationRobotv1UpdateRobotV1RobotIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	robotIdDescription := `Required. Robot ID`

	var robotIdFlagName string
	if cmdPrefix == "" {
		robotIdFlagName = "robot_id"
	} else {
		robotIdFlagName = fmt.Sprintf("%v.robot_id", cmdPrefix)
	}

	var robotIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(robotIdFlagName, robotIdFlagDefault, robotIdDescription)

	return nil
}

func retrieveOperationRobotv1UpdateRobotV1XIsResourceNameFlag(m *robotv1.UpdateRobotV1Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("X-Is-Resource-Name") {

		var xIsResourceNameFlagName string
		if cmdPrefix == "" {
			xIsResourceNameFlagName = "X-Is-Resource-Name"
		} else {
			xIsResourceNameFlagName = fmt.Sprintf("%v.X-Is-Resource-Name", cmdPrefix)
		}

		xIsResourceNameFlagValue, err := cmd.Flags().GetBool(xIsResourceNameFlagName)
		if err != nil {
			return err, false
		}
		m.XIsResourceName = &xIsResourceNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationRobotv1UpdateRobotV1XRequestIDFlag(m *robotv1.UpdateRobotV1Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationRobotv1UpdateRobotV1ProjectNameOrIDFlag(m *robotv1.UpdateRobotV1Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("project_name_or_id") {

		var projectNameOrIdFlagName string
		if cmdPrefix == "" {
			projectNameOrIdFlagName = "project_name_or_id"
		} else {
			projectNameOrIdFlagName = fmt.Sprintf("%v.project_name_or_id", cmdPrefix)
		}

		projectNameOrIdFlagValue, err := cmd.Flags().GetString(projectNameOrIdFlagName)
		if err != nil {
			return err, false
		}
		m.ProjectNameOrID = projectNameOrIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationRobotv1UpdateRobotV1RobotFlag(m *robotv1.UpdateRobotV1Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("robot") {
		// Read robot string from cmd and unmarshal
		robotValueStr, err := cmd.Flags().GetString("robot")
		if err != nil {
			return err, false
		}

		robotValue := models.Robot{}
		if err := json.Unmarshal([]byte(robotValueStr), &robotValue); err != nil {
			return fmt.Errorf("cannot unmarshal robot string in models.Robot: %v", err), false
		}
		m.Robot = &robotValue
	}
	robotValueModel := m.Robot
	if swag.IsZero(robotValueModel) {
		robotValueModel = &models.Robot{}
	}
	err, added := retrieveModelRobotFlags(0, robotValueModel, "robot", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Robot = robotValueModel
	}
	if dryRun && debug {

		robotValueDebugBytes, err := json.Marshal(m.Robot)
		if err != nil {
			return err, false
		}
		logDebugf("Robot dry-run payload: %v", string(robotValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationRobotv1UpdateRobotV1RobotIDFlag(m *robotv1.UpdateRobotV1Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("robot_id") {

		var robotIdFlagName string
		if cmdPrefix == "" {
			robotIdFlagName = "robot_id"
		} else {
			robotIdFlagName = fmt.Sprintf("%v.robot_id", cmdPrefix)
		}

		robotIdFlagValue, err := cmd.Flags().GetInt64(robotIdFlagName)
		if err != nil {
			return err, false
		}
		m.RobotID = robotIdFlagValue

	}
	return nil, retAdded
}

// parseOperationRobotv1UpdateRobotV1Result parses request result and return the string content
func parseOperationRobotv1UpdateRobotV1Result(resp0 *robotv1.UpdateRobotV1OK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning updateRobotV1OK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*robotv1.UpdateRobotV1BadRequest)
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
		resp2, ok := iResp2.(*robotv1.UpdateRobotV1Unauthorized)
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
		resp3, ok := iResp3.(*robotv1.UpdateRobotV1Forbidden)
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
		resp4, ok := iResp4.(*robotv1.UpdateRobotV1NotFound)
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
		resp5, ok := iResp5.(*robotv1.UpdateRobotV1Conflict)
		if ok {
			if !swag.IsZero(resp5) && !swag.IsZero(resp5.Payload) {
				msgStr, err := json.Marshal(resp5.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp6 interface{} = respErr
		resp6, ok := iResp6.(*robotv1.UpdateRobotV1InternalServerError)
		if ok {
			if !swag.IsZero(resp6) && !swag.IsZero(resp6.Payload) {
				msgStr, err := json.Marshal(resp6.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response updateRobotV1OK is not supported by go-swagger cli yet.

	return "", nil
}
