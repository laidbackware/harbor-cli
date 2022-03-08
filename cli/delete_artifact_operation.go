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

// makeOperationArtifactDeleteArtifactCmd returns a cmd to handle operation deleteArtifact
func makeOperationArtifactDeleteArtifactCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteArtifact",
		Short: `Delete the artifact specified by the reference under the project and repository. The reference can be digest or tag`,
		RunE:  runOperationArtifactDeleteArtifact,
	}

	if err := registerOperationArtifactDeleteArtifactParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationArtifactDeleteArtifact uses cmd flags to call endpoint api
func runOperationArtifactDeleteArtifact(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := artifact.NewDeleteArtifactParams()
	if err, _ := retrieveOperationArtifactDeleteArtifactXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationArtifactDeleteArtifactProjectNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationArtifactDeleteArtifactReferenceFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationArtifactDeleteArtifactRepositoryNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationArtifactDeleteArtifactResult(appCli.Artifact.DeleteArtifact(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationArtifactDeleteArtifactParamFlags registers all flags needed to fill params
func registerOperationArtifactDeleteArtifactParamFlags(cmd *cobra.Command) error {
	if err := registerOperationArtifactDeleteArtifactXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationArtifactDeleteArtifactProjectNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationArtifactDeleteArtifactReferenceParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationArtifactDeleteArtifactRepositoryNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationArtifactDeleteArtifactXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationArtifactDeleteArtifactProjectNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationArtifactDeleteArtifactReferenceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	referenceDescription := `Required. The reference of the artifact, can be digest or tag`

	var referenceFlagName string
	if cmdPrefix == "" {
		referenceFlagName = "reference"
	} else {
		referenceFlagName = fmt.Sprintf("%v.reference", cmdPrefix)
	}

	var referenceFlagDefault string

	_ = cmd.PersistentFlags().String(referenceFlagName, referenceFlagDefault, referenceDescription)

	return nil
}
func registerOperationArtifactDeleteArtifactRepositoryNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationArtifactDeleteArtifactXRequestIDFlag(m *artifact.DeleteArtifactParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationArtifactDeleteArtifactProjectNameFlag(m *artifact.DeleteArtifactParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationArtifactDeleteArtifactReferenceFlag(m *artifact.DeleteArtifactParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("reference") {

		var referenceFlagName string
		if cmdPrefix == "" {
			referenceFlagName = "reference"
		} else {
			referenceFlagName = fmt.Sprintf("%v.reference", cmdPrefix)
		}

		referenceFlagValue, err := cmd.Flags().GetString(referenceFlagName)
		if err != nil {
			return err, false
		}
		m.Reference = referenceFlagValue

	}
	return nil, retAdded
}
func retrieveOperationArtifactDeleteArtifactRepositoryNameFlag(m *artifact.DeleteArtifactParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationArtifactDeleteArtifactResult parses request result and return the string content
func parseOperationArtifactDeleteArtifactResult(resp0 *artifact.DeleteArtifactOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteArtifactOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*artifact.DeleteArtifactUnauthorized)
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
		resp2, ok := iResp2.(*artifact.DeleteArtifactForbidden)
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
		resp3, ok := iResp3.(*artifact.DeleteArtifactNotFound)
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
		resp4, ok := iResp4.(*artifact.DeleteArtifactInternalServerError)
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

	// warning: non schema response deleteArtifactOK is not supported by go-swagger cli yet.

	return "", nil
}