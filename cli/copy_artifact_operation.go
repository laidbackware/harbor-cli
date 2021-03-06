// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/artifact"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationArtifactCopyArtifactCmd returns a cmd to handle operation copyArtifact
func makeOperationArtifactCopyArtifactCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "CopyArtifact",
		Short: `Copy the artifact specified in the 'from' parameter to the repository.`,
		RunE:  runOperationArtifactCopyArtifact,
	}

	if err := registerOperationArtifactCopyArtifactParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationArtifactCopyArtifact uses cmd flags to call endpoint api
func runOperationArtifactCopyArtifact(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := artifact.NewCopyArtifactParams()
	if err, _ := retrieveOperationArtifactCopyArtifactXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationArtifactCopyArtifactFromFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationArtifactCopyArtifactProjectNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationArtifactCopyArtifactRepositoryNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationArtifactCopyArtifactResult(appCli.Artifact.CopyArtifact(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationArtifactCopyArtifactParamFlags registers all flags needed to fill params
func registerOperationArtifactCopyArtifactParamFlags(cmd *cobra.Command) error {
	if err := registerOperationArtifactCopyArtifactXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationArtifactCopyArtifactFromParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationArtifactCopyArtifactProjectNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationArtifactCopyArtifactRepositoryNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationArtifactCopyArtifactXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationArtifactCopyArtifactFromParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	fromDescription := `Required. The artifact from which the new artifact is copied from, the format should be 'project/repository:tag' or 'project/repository@digest'.`

	var fromFlagName string
	if cmdPrefix == "" {
		fromFlagName = "from"
	} else {
		fromFlagName = fmt.Sprintf("%v.from", cmdPrefix)
	}

	var fromFlagDefault string

	_ = cmd.PersistentFlags().String(fromFlagName, fromFlagDefault, fromDescription)

	return nil
}
func registerOperationArtifactCopyArtifactProjectNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	projectNameDescription := `Required. The name of the project`

	var projectNameFlagName string
	if cmdPrefix == "" {
		projectNameFlagName = "project_name"
	} else {
		projectNameFlagName = fmt.Sprintf("%v.project_name", cmdPrefix)
	}

	var projectNameFlagDefault string

	_ = cmd.PersistentFlags().String(projectNameFlagName, projectNameFlagDefault, projectNameDescription)

	return nil
}
func registerOperationArtifactCopyArtifactRepositoryNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	repositoryNameDescription := `Required. The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -> a%252Fb`

	var repositoryNameFlagName string
	if cmdPrefix == "" {
		repositoryNameFlagName = "repository_name"
	} else {
		repositoryNameFlagName = fmt.Sprintf("%v.repository_name", cmdPrefix)
	}

	var repositoryNameFlagDefault string

	_ = cmd.PersistentFlags().String(repositoryNameFlagName, repositoryNameFlagDefault, repositoryNameDescription)

	return nil
}

func retrieveOperationArtifactCopyArtifactXRequestIDFlag(m *artifact.CopyArtifactParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationArtifactCopyArtifactFromFlag(m *artifact.CopyArtifactParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("from") {

		var fromFlagName string
		if cmdPrefix == "" {
			fromFlagName = "from"
		} else {
			fromFlagName = fmt.Sprintf("%v.from", cmdPrefix)
		}

		fromFlagValue, err := cmd.Flags().GetString(fromFlagName)
		if err != nil {
			return err, false
		}
		m.From = fromFlagValue

	}
	return nil, retAdded
}
func retrieveOperationArtifactCopyArtifactProjectNameFlag(m *artifact.CopyArtifactParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("project_name") {

		var projectNameFlagName string
		if cmdPrefix == "" {
			projectNameFlagName = "project_name"
		} else {
			projectNameFlagName = fmt.Sprintf("%v.project_name", cmdPrefix)
		}

		projectNameFlagValue, err := cmd.Flags().GetString(projectNameFlagName)
		if err != nil {
			return err, false
		}
		m.ProjectName = projectNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationArtifactCopyArtifactRepositoryNameFlag(m *artifact.CopyArtifactParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("repository_name") {

		var repositoryNameFlagName string
		if cmdPrefix == "" {
			repositoryNameFlagName = "repository_name"
		} else {
			repositoryNameFlagName = fmt.Sprintf("%v.repository_name", cmdPrefix)
		}

		repositoryNameFlagValue, err := cmd.Flags().GetString(repositoryNameFlagName)
		if err != nil {
			return err, false
		}
		m.RepositoryName = repositoryNameFlagValue

	}
	return nil, retAdded
}

// parseOperationArtifactCopyArtifactResult parses request result and return the string content
func parseOperationArtifactCopyArtifactResult(resp0 *artifact.CopyArtifactCreated, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning copyArtifactCreated is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*artifact.CopyArtifactBadRequest)
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
		resp2, ok := iResp2.(*artifact.CopyArtifactUnauthorized)
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
		resp3, ok := iResp3.(*artifact.CopyArtifactForbidden)
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
		resp4, ok := iResp4.(*artifact.CopyArtifactNotFound)
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
		resp5, ok := iResp5.(*artifact.CopyArtifactMethodNotAllowed)
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
		resp6, ok := iResp6.(*artifact.CopyArtifactInternalServerError)
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

	// warning: non schema response copyArtifactCreated is not supported by go-swagger cli yet.

	return "", nil
}
