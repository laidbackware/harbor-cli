// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for ChartMetadata

// register flags to command
func registerModelChartMetadataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerChartMetadataAPIVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerChartMetadataAppVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerChartMetadataDeprecated(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerChartMetadataDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerChartMetadataEngine(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerChartMetadataHome(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerChartMetadataIcon(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerChartMetadataKeywords(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerChartMetadataName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerChartMetadataSources(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerChartMetadataVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerChartMetadataAPIVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	apiVersionDescription := `Required. The API version of this chart`

	var apiVersionFlagName string
	if cmdPrefix == "" {
		apiVersionFlagName = "apiVersion"
	} else {
		apiVersionFlagName = fmt.Sprintf("%v.apiVersion", cmdPrefix)
	}

	var apiVersionFlagDefault string

	_ = cmd.PersistentFlags().String(apiVersionFlagName, apiVersionFlagDefault, apiVersionDescription)

	return nil
}

func registerChartMetadataAppVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	appVersionDescription := `Required. The version of the application enclosed in the chart`

	var appVersionFlagName string
	if cmdPrefix == "" {
		appVersionFlagName = "appVersion"
	} else {
		appVersionFlagName = fmt.Sprintf("%v.appVersion", cmdPrefix)
	}

	var appVersionFlagDefault string

	_ = cmd.PersistentFlags().String(appVersionFlagName, appVersionFlagDefault, appVersionDescription)

	return nil
}

func registerChartMetadataDeprecated(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deprecatedDescription := `Whether or not this chart is deprecated`

	var deprecatedFlagName string
	if cmdPrefix == "" {
		deprecatedFlagName = "deprecated"
	} else {
		deprecatedFlagName = fmt.Sprintf("%v.deprecated", cmdPrefix)
	}

	var deprecatedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(deprecatedFlagName, deprecatedFlagDefault, deprecatedDescription)

	return nil
}

func registerChartMetadataDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `A one-sentence description of chart`

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerChartMetadataEngine(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	engineDescription := `Required. The name of template engine`

	var engineFlagName string
	if cmdPrefix == "" {
		engineFlagName = "engine"
	} else {
		engineFlagName = fmt.Sprintf("%v.engine", cmdPrefix)
	}

	var engineFlagDefault string

	_ = cmd.PersistentFlags().String(engineFlagName, engineFlagDefault, engineDescription)

	return nil
}

func registerChartMetadataHome(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	homeDescription := `The URL to the relevant project page`

	var homeFlagName string
	if cmdPrefix == "" {
		homeFlagName = "home"
	} else {
		homeFlagName = fmt.Sprintf("%v.home", cmdPrefix)
	}

	var homeFlagDefault string

	_ = cmd.PersistentFlags().String(homeFlagName, homeFlagDefault, homeDescription)

	return nil
}

func registerChartMetadataIcon(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	iconDescription := `Required. The URL to an icon file`

	var iconFlagName string
	if cmdPrefix == "" {
		iconFlagName = "icon"
	} else {
		iconFlagName = fmt.Sprintf("%v.icon", cmdPrefix)
	}

	var iconFlagDefault string

	_ = cmd.PersistentFlags().String(iconFlagName, iconFlagDefault, iconDescription)

	return nil
}

func registerChartMetadataKeywords(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: keywords []string array type is not supported by go-swagger cli yet

	return nil
}

func registerChartMetadataName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. The name of the chart`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerChartMetadataSources(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: sources []string array type is not supported by go-swagger cli yet

	return nil
}

func registerChartMetadataVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	versionDescription := `Required. A SemVer 2 version of chart`

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "version"
	} else {
		versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
	}

	var versionFlagDefault string

	_ = cmd.PersistentFlags().String(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelChartMetadataFlags(depth int, m *models.ChartMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, apiVersionAdded := retrieveChartMetadataAPIVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || apiVersionAdded

	err, appVersionAdded := retrieveChartMetadataAppVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || appVersionAdded

	err, deprecatedAdded := retrieveChartMetadataDeprecatedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deprecatedAdded

	err, descriptionAdded := retrieveChartMetadataDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, engineAdded := retrieveChartMetadataEngineFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || engineAdded

	err, homeAdded := retrieveChartMetadataHomeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || homeAdded

	err, iconAdded := retrieveChartMetadataIconFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || iconAdded

	err, keywordsAdded := retrieveChartMetadataKeywordsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || keywordsAdded

	err, nameAdded := retrieveChartMetadataNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, sourcesAdded := retrieveChartMetadataSourcesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sourcesAdded

	err, versionAdded := retrieveChartMetadataVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	return nil, retAdded
}

func retrieveChartMetadataAPIVersionFlags(depth int, m *models.ChartMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	apiVersionFlagName := fmt.Sprintf("%v.apiVersion", cmdPrefix)
	if cmd.Flags().Changed(apiVersionFlagName) {

		var apiVersionFlagName string
		if cmdPrefix == "" {
			apiVersionFlagName = "apiVersion"
		} else {
			apiVersionFlagName = fmt.Sprintf("%v.apiVersion", cmdPrefix)
		}

		apiVersionFlagValue, err := cmd.Flags().GetString(apiVersionFlagName)
		if err != nil {
			return err, false
		}
		m.APIVersion = &apiVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveChartMetadataAppVersionFlags(depth int, m *models.ChartMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	appVersionFlagName := fmt.Sprintf("%v.appVersion", cmdPrefix)
	if cmd.Flags().Changed(appVersionFlagName) {

		var appVersionFlagName string
		if cmdPrefix == "" {
			appVersionFlagName = "appVersion"
		} else {
			appVersionFlagName = fmt.Sprintf("%v.appVersion", cmdPrefix)
		}

		appVersionFlagValue, err := cmd.Flags().GetString(appVersionFlagName)
		if err != nil {
			return err, false
		}
		m.AppVersion = &appVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveChartMetadataDeprecatedFlags(depth int, m *models.ChartMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deprecatedFlagName := fmt.Sprintf("%v.deprecated", cmdPrefix)
	if cmd.Flags().Changed(deprecatedFlagName) {

		var deprecatedFlagName string
		if cmdPrefix == "" {
			deprecatedFlagName = "deprecated"
		} else {
			deprecatedFlagName = fmt.Sprintf("%v.deprecated", cmdPrefix)
		}

		deprecatedFlagValue, err := cmd.Flags().GetBool(deprecatedFlagName)
		if err != nil {
			return err, false
		}
		m.Deprecated = deprecatedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveChartMetadataDescriptionFlags(depth int, m *models.ChartMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveChartMetadataEngineFlags(depth int, m *models.ChartMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	engineFlagName := fmt.Sprintf("%v.engine", cmdPrefix)
	if cmd.Flags().Changed(engineFlagName) {

		var engineFlagName string
		if cmdPrefix == "" {
			engineFlagName = "engine"
		} else {
			engineFlagName = fmt.Sprintf("%v.engine", cmdPrefix)
		}

		engineFlagValue, err := cmd.Flags().GetString(engineFlagName)
		if err != nil {
			return err, false
		}
		m.Engine = &engineFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveChartMetadataHomeFlags(depth int, m *models.ChartMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	homeFlagName := fmt.Sprintf("%v.home", cmdPrefix)
	if cmd.Flags().Changed(homeFlagName) {

		var homeFlagName string
		if cmdPrefix == "" {
			homeFlagName = "home"
		} else {
			homeFlagName = fmt.Sprintf("%v.home", cmdPrefix)
		}

		homeFlagValue, err := cmd.Flags().GetString(homeFlagName)
		if err != nil {
			return err, false
		}
		m.Home = homeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveChartMetadataIconFlags(depth int, m *models.ChartMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	iconFlagName := fmt.Sprintf("%v.icon", cmdPrefix)
	if cmd.Flags().Changed(iconFlagName) {

		var iconFlagName string
		if cmdPrefix == "" {
			iconFlagName = "icon"
		} else {
			iconFlagName = fmt.Sprintf("%v.icon", cmdPrefix)
		}

		iconFlagValue, err := cmd.Flags().GetString(iconFlagName)
		if err != nil {
			return err, false
		}
		m.Icon = &iconFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveChartMetadataKeywordsFlags(depth int, m *models.ChartMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	keywordsFlagName := fmt.Sprintf("%v.keywords", cmdPrefix)
	if cmd.Flags().Changed(keywordsFlagName) {
		// warning: keywords array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveChartMetadataNameFlags(depth int, m *models.ChartMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveChartMetadataSourcesFlags(depth int, m *models.ChartMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sourcesFlagName := fmt.Sprintf("%v.sources", cmdPrefix)
	if cmd.Flags().Changed(sourcesFlagName) {
		// warning: sources array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveChartMetadataVersionFlags(depth int, m *models.ChartMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionFlagName := fmt.Sprintf("%v.version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "version"
		} else {
			versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetString(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = &versionFlagValue

		retAdded = true
	}

	return nil, retAdded
}