// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/artifact"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationArtifactAddLabelCmd returns a cmd to handle operation addLabel
func makeOperationArtifactAddLabelCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "addLabel",
		Short: `Add label to the specified artiact.`,
		RunE:  runOperationArtifactAddLabel,
	}

	if err := registerOperationArtifactAddLabelParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationArtifactAddLabel uses cmd flags to call endpoint api
func runOperationArtifactAddLabel(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := artifact.NewAddLabelParams()
	if err, _ := retrieveOperationArtifactAddLabelXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationArtifactAddLabelLabelFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationArtifactAddLabelProjectNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationArtifactAddLabelReferenceFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationArtifactAddLabelRepositoryNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationArtifactAddLabelResult(appCli.Artifact.AddLabel(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationArtifactAddLabelParamFlags registers all flags needed to fill params
func registerOperationArtifactAddLabelParamFlags(cmd *cobra.Command) error {
	if err := registerOperationArtifactAddLabelXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationArtifactAddLabelLabelParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationArtifactAddLabelProjectNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationArtifactAddLabelReferenceParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationArtifactAddLabelRepositoryNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationArtifactAddLabelXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationArtifactAddLabelLabelParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var labelFlagName string
	if cmdPrefix == "" {
		labelFlagName = "label"
	} else {
		labelFlagName = fmt.Sprintf("%v.label", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(labelFlagName, "", "Optional json string for [label]. The label that added to the artifact. Only the ID property is needed.")

	// add flags for body
	if err := registerModelLabelFlags(0, "label", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationArtifactAddLabelProjectNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationArtifactAddLabelReferenceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationArtifactAddLabelRepositoryNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationArtifactAddLabelXRequestIDFlag(m *artifact.AddLabelParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationArtifactAddLabelLabelFlag(m *artifact.AddLabelParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("label") {
		// Read label string from cmd and unmarshal
		labelValueStr, err := cmd.Flags().GetString("label")
		if err != nil {
			return err, false
		}

		labelValue := models.Label{}
		if err := json.Unmarshal([]byte(labelValueStr), &labelValue); err != nil {
			return fmt.Errorf("cannot unmarshal label string in models.Label: %v", err), false
		}
		m.Label = &labelValue
	}
	labelValueModel := m.Label
	if swag.IsZero(labelValueModel) {
		labelValueModel = &models.Label{}
	}
	err, added := retrieveModelLabelFlags(0, labelValueModel, "label", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Label = labelValueModel
	}
	if dryRun && debug {

		labelValueDebugBytes, err := json.Marshal(m.Label)
		if err != nil {
			return err, false
		}
		logDebugf("Label dry-run payload: %v", string(labelValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationArtifactAddLabelProjectNameFlag(m *artifact.AddLabelParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationArtifactAddLabelReferenceFlag(m *artifact.AddLabelParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationArtifactAddLabelRepositoryNameFlag(m *artifact.AddLabelParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationArtifactAddLabelResult parses request result and return the string content
func parseOperationArtifactAddLabelResult(resp0 *artifact.AddLabelOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning addLabelOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*artifact.AddLabelBadRequest)
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
		resp2, ok := iResp2.(*artifact.AddLabelUnauthorized)
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
		resp3, ok := iResp3.(*artifact.AddLabelForbidden)
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
		resp4, ok := iResp4.(*artifact.AddLabelNotFound)
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
		resp5, ok := iResp5.(*artifact.AddLabelConflict)
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
		resp6, ok := iResp6.(*artifact.AddLabelInternalServerError)
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

	// warning: non schema response addLabelOK is not supported by go-swagger cli yet.

	return "", nil
}