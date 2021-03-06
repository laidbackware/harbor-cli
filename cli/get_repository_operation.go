// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/repository"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationRepositoryGetRepositoryCmd returns a cmd to handle operation getRepository
func makeOperationRepositoryGetRepositoryCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getRepository",
		Short: `Get the repository specified by name`,
		RunE:  runOperationRepositoryGetRepository,
	}

	if err := registerOperationRepositoryGetRepositoryParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationRepositoryGetRepository uses cmd flags to call endpoint api
func runOperationRepositoryGetRepository(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := repository.NewGetRepositoryParams()
	if err, _ := retrieveOperationRepositoryGetRepositoryXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRepositoryGetRepositoryProjectNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationRepositoryGetRepositoryRepositoryNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationRepositoryGetRepositoryResult(appCli.Repository.GetRepository(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationRepositoryGetRepositoryParamFlags registers all flags needed to fill params
func registerOperationRepositoryGetRepositoryParamFlags(cmd *cobra.Command) error {
	if err := registerOperationRepositoryGetRepositoryXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRepositoryGetRepositoryProjectNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationRepositoryGetRepositoryRepositoryNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationRepositoryGetRepositoryXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationRepositoryGetRepositoryProjectNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationRepositoryGetRepositoryRepositoryNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationRepositoryGetRepositoryXRequestIDFlag(m *repository.GetRepositoryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationRepositoryGetRepositoryProjectNameFlag(m *repository.GetRepositoryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationRepositoryGetRepositoryRepositoryNameFlag(m *repository.GetRepositoryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationRepositoryGetRepositoryResult parses request result and return the string content
func parseOperationRepositoryGetRepositoryResult(resp0 *repository.GetRepositoryOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*repository.GetRepositoryOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*repository.GetRepositoryBadRequest)
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
		resp2, ok := iResp2.(*repository.GetRepositoryUnauthorized)
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
		resp3, ok := iResp3.(*repository.GetRepositoryForbidden)
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
		resp4, ok := iResp4.(*repository.GetRepositoryNotFound)
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
		resp5, ok := iResp5.(*repository.GetRepositoryInternalServerError)
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

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
