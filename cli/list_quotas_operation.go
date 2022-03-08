// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/quota"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationQuotaListQuotasCmd returns a cmd to handle operation listQuotas
func makeOperationQuotaListQuotasCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "listQuotas",
		Short: `List quotas`,
		RunE:  runOperationQuotaListQuotas,
	}

	if err := registerOperationQuotaListQuotasParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationQuotaListQuotas uses cmd flags to call endpoint api
func runOperationQuotaListQuotas(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := quota.NewListQuotasParams()
	if err, _ := retrieveOperationQuotaListQuotasXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationQuotaListQuotasPageFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationQuotaListQuotasPageSizeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationQuotaListQuotasReferenceFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationQuotaListQuotasReferenceIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationQuotaListQuotasSortFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationQuotaListQuotasResult(appCli.Quota.ListQuotas(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationQuotaListQuotasParamFlags registers all flags needed to fill params
func registerOperationQuotaListQuotasParamFlags(cmd *cobra.Command) error {
	if err := registerOperationQuotaListQuotasXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationQuotaListQuotasPageParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationQuotaListQuotasPageSizeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationQuotaListQuotasReferenceParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationQuotaListQuotasReferenceIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationQuotaListQuotasSortParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationQuotaListQuotasXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationQuotaListQuotasPageParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pageDescription := `The page number`

	var pageFlagName string
	if cmdPrefix == "" {
		pageFlagName = "page"
	} else {
		pageFlagName = fmt.Sprintf("%v.page", cmdPrefix)
	}

	var pageFlagDefault int64 = 1

	_ = cmd.PersistentFlags().Int64(pageFlagName, pageFlagDefault, pageDescription)

	return nil
}
func registerOperationQuotaListQuotasPageSizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pageSizeDescription := `The size of per page`

	var pageSizeFlagName string
	if cmdPrefix == "" {
		pageSizeFlagName = "page_size"
	} else {
		pageSizeFlagName = fmt.Sprintf("%v.page_size", cmdPrefix)
	}

	var pageSizeFlagDefault int64 = 10

	_ = cmd.PersistentFlags().Int64(pageSizeFlagName, pageSizeFlagDefault, pageSizeDescription)

	return nil
}
func registerOperationQuotaListQuotasReferenceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	referenceDescription := `The reference type of quota.`

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
func registerOperationQuotaListQuotasReferenceIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	referenceIdDescription := `The reference id of quota.`

	var referenceIdFlagName string
	if cmdPrefix == "" {
		referenceIdFlagName = "reference_id"
	} else {
		referenceIdFlagName = fmt.Sprintf("%v.reference_id", cmdPrefix)
	}

	var referenceIdFlagDefault string

	_ = cmd.PersistentFlags().String(referenceIdFlagName, referenceIdFlagDefault, referenceIdDescription)

	return nil
}
func registerOperationQuotaListQuotasSortParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sortDescription := `Sort method, valid values include:
'hard.resource_name', '-hard.resource_name', 'used.resource_name', '-used.resource_name'.
Here '-' stands for descending order, resource_name should be the real resource name of the quota.
`

	var sortFlagName string
	if cmdPrefix == "" {
		sortFlagName = "sort"
	} else {
		sortFlagName = fmt.Sprintf("%v.sort", cmdPrefix)
	}

	var sortFlagDefault string

	_ = cmd.PersistentFlags().String(sortFlagName, sortFlagDefault, sortDescription)

	return nil
}

func retrieveOperationQuotaListQuotasXRequestIDFlag(m *quota.ListQuotasParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationQuotaListQuotasPageFlag(m *quota.ListQuotasParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("page") {

		var pageFlagName string
		if cmdPrefix == "" {
			pageFlagName = "page"
		} else {
			pageFlagName = fmt.Sprintf("%v.page", cmdPrefix)
		}

		pageFlagValue, err := cmd.Flags().GetInt64(pageFlagName)
		if err != nil {
			return err, false
		}
		m.Page = &pageFlagValue

	}
	return nil, retAdded
}
func retrieveOperationQuotaListQuotasPageSizeFlag(m *quota.ListQuotasParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("page_size") {

		var pageSizeFlagName string
		if cmdPrefix == "" {
			pageSizeFlagName = "page_size"
		} else {
			pageSizeFlagName = fmt.Sprintf("%v.page_size", cmdPrefix)
		}

		pageSizeFlagValue, err := cmd.Flags().GetInt64(pageSizeFlagName)
		if err != nil {
			return err, false
		}
		m.PageSize = &pageSizeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationQuotaListQuotasReferenceFlag(m *quota.ListQuotasParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Reference = &referenceFlagValue

	}
	return nil, retAdded
}
func retrieveOperationQuotaListQuotasReferenceIDFlag(m *quota.ListQuotasParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("reference_id") {

		var referenceIdFlagName string
		if cmdPrefix == "" {
			referenceIdFlagName = "reference_id"
		} else {
			referenceIdFlagName = fmt.Sprintf("%v.reference_id", cmdPrefix)
		}

		referenceIdFlagValue, err := cmd.Flags().GetString(referenceIdFlagName)
		if err != nil {
			return err, false
		}
		m.ReferenceID = &referenceIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationQuotaListQuotasSortFlag(m *quota.ListQuotasParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("sort") {

		var sortFlagName string
		if cmdPrefix == "" {
			sortFlagName = "sort"
		} else {
			sortFlagName = fmt.Sprintf("%v.sort", cmdPrefix)
		}

		sortFlagValue, err := cmd.Flags().GetString(sortFlagName)
		if err != nil {
			return err, false
		}
		m.Sort = &sortFlagValue

	}
	return nil, retAdded
}

// parseOperationQuotaListQuotasResult parses request result and return the string content
func parseOperationQuotaListQuotasResult(resp0 *quota.ListQuotasOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*quota.ListQuotasOK)
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
		resp1, ok := iResp1.(*quota.ListQuotasUnauthorized)
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
		resp2, ok := iResp2.(*quota.ListQuotasForbidden)
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
		resp3, ok := iResp3.(*quota.ListQuotasInternalServerError)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
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
