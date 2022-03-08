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

// makeOperationProjectUpdateProjectCmd returns a cmd to handle operation updateProject
func makeOperationProjectUpdateProjectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "updateProject",
		Short: `This endpoint is aimed to update the properties of a project.`,
		RunE:  runOperationProjectUpdateProject,
	}

	if err := registerOperationProjectUpdateProjectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationProjectUpdateProject uses cmd flags to call endpoint api
func runOperationProjectUpdateProject(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := project.NewUpdateProjectParams()
	if err, _ := retrieveOperationProjectUpdateProjectXIsResourceNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectUpdateProjectXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectUpdateProjectProjectFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectUpdateProjectProjectNameOrIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationProjectUpdateProjectResult(appCli.Project.UpdateProject(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationProjectUpdateProjectParamFlags registers all flags needed to fill params
func registerOperationProjectUpdateProjectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationProjectUpdateProjectXIsResourceNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectUpdateProjectXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectUpdateProjectProjectParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectUpdateProjectProjectNameOrIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationProjectUpdateProjectXIsResourceNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationProjectUpdateProjectXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationProjectUpdateProjectProjectParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var projectFlagName string
	if cmdPrefix == "" {
		projectFlagName = "project"
	} else {
		projectFlagName = fmt.Sprintf("%v.project", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(projectFlagName, "", "Optional json string for [project]. Updates of project.")

	// add flags for body
	if err := registerModelProjectReqFlags(0, "projectReq", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationProjectUpdateProjectProjectNameOrIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationProjectUpdateProjectXIsResourceNameFlag(m *project.UpdateProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationProjectUpdateProjectXRequestIDFlag(m *project.UpdateProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationProjectUpdateProjectProjectFlag(m *project.UpdateProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("project") {
		// Read project string from cmd and unmarshal
		projectValueStr, err := cmd.Flags().GetString("project")
		if err != nil {
			return err, false
		}

		projectValue := models.ProjectReq{}
		if err := json.Unmarshal([]byte(projectValueStr), &projectValue); err != nil {
			return fmt.Errorf("cannot unmarshal project string in models.ProjectReq: %v", err), false
		}
		m.Project = &projectValue
	}
	projectValueModel := m.Project
	if swag.IsZero(projectValueModel) {
		projectValueModel = &models.ProjectReq{}
	}
	err, added := retrieveModelProjectReqFlags(0, projectValueModel, "projectReq", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Project = projectValueModel
	}
	if dryRun && debug {

		projectValueDebugBytes, err := json.Marshal(m.Project)
		if err != nil {
			return err, false
		}
		logDebugf("Project dry-run payload: %v", string(projectValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationProjectUpdateProjectProjectNameOrIDFlag(m *project.UpdateProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationProjectUpdateProjectResult parses request result and return the string content
func parseOperationProjectUpdateProjectResult(resp0 *project.UpdateProjectOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning updateProjectOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*project.UpdateProjectBadRequest)
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
		resp2, ok := iResp2.(*project.UpdateProjectUnauthorized)
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
		resp3, ok := iResp3.(*project.UpdateProjectForbidden)
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
		resp4, ok := iResp4.(*project.UpdateProjectNotFound)
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
		resp5, ok := iResp5.(*project.UpdateProjectInternalServerError)
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

	// warning: non schema response updateProjectOK is not supported by go-swagger cli yet.

	return "", nil
}