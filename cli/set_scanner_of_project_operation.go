// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/project"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationProjectSetScannerOfProjectCmd returns a cmd to handle operation setScannerOfProject
func makeOperationProjectSetScannerOfProjectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "setScannerOfProject",
		Short: `Set one of the system configured scanner registration as the indepndent scanner of the specified project.`,
		RunE:  runOperationProjectSetScannerOfProject,
	}

	if err := registerOperationProjectSetScannerOfProjectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationProjectSetScannerOfProject uses cmd flags to call endpoint api
func runOperationProjectSetScannerOfProject(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := project.NewSetScannerOfProjectParams()
	if err, _ := retrieveOperationProjectSetScannerOfProjectXIsResourceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectSetScannerOfProjectXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectSetScannerOfProjectPayloadFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectSetScannerOfProjectProjectNameOrIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationProjectSetScannerOfProjectResult(appCli.Project.SetScannerOfProject(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationProjectSetScannerOfProjectParamFlags registers all flags needed to fill params
func registerOperationProjectSetScannerOfProjectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationProjectSetScannerOfProjectXIsResourceNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectSetScannerOfProjectXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectSetScannerOfProjectPayloadParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectSetScannerOfProjectProjectNameOrIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationProjectSetScannerOfProjectXIsResourceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationProjectSetScannerOfProjectXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationProjectSetScannerOfProjectPayloadParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var payloadFlagName string
	if cmdPrefix == "" {
		payloadFlagName = "payload"
	} else {
		payloadFlagName = fmt.Sprintf("%v.payload", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(payloadFlagName, "", "Optional json string for [payload]. ")

	// add flags for body
	if err := registerModelProjectScannerFlags(0, "projectScanner", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationProjectSetScannerOfProjectProjectNameOrIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationProjectSetScannerOfProjectXIsResourceNameFlag(m *project.SetScannerOfProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationProjectSetScannerOfProjectXRequestIDFlag(m *project.SetScannerOfProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationProjectSetScannerOfProjectPayloadFlag(m *project.SetScannerOfProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("payload") {
		// Read payload string from cmd and unmarshal
		payloadValueStr, err := cmd.Flags().GetString("payload")
		if err != nil {
			return err, false
		}

		payloadValue := models.ProjectScanner{}
		if err := json.Unmarshal([]byte(payloadValueStr), &payloadValue); err != nil {
			return fmt.Errorf("cannot unmarshal payload string in models.ProjectScanner: %v", err), false
		}
		m.Payload = &payloadValue
	}
	payloadValueModel := m.Payload
	if swag.IsZero(payloadValueModel) {
		payloadValueModel = &models.ProjectScanner{}
	}
	err, added := retrieveModelProjectScannerFlags(0, payloadValueModel, "projectScanner", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Payload = payloadValueModel
	}
	if dryRun && debug {

		payloadValueDebugBytes, err := json.Marshal(m.Payload)
		if err != nil {
			return err, false
		}
		logDebugf("Payload dry-run payload: %v", string(payloadValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationProjectSetScannerOfProjectProjectNameOrIDFlag(m *project.SetScannerOfProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationProjectSetScannerOfProjectResult parses request result and return the string content
func parseOperationProjectSetScannerOfProjectResult(resp0 *project.SetScannerOfProjectOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning setScannerOfProjectOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*project.SetScannerOfProjectBadRequest)
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
		resp2, ok := iResp2.(*project.SetScannerOfProjectUnauthorized)
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
		resp3, ok := iResp3.(*project.SetScannerOfProjectForbidden)
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
		resp4, ok := iResp4.(*project.SetScannerOfProjectNotFound)
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
		resp5, ok := iResp5.(*project.SetScannerOfProjectInternalServerError)
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

	// warning: non schema response setScannerOfProjectOK is not supported by go-swagger cli yet.

	return "", nil
}
