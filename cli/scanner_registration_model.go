// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for ScannerRegistration

// register flags to command
func registerModelScannerRegistrationFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerScannerRegistrationAccessCredential(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationAdapter(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationAuth(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationCreateTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationDisabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationHealth(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationIsDefault(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationSkipCertVerify(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationUpdateTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationUseInternalAddr(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationUUID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationVendor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerScannerRegistrationVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerScannerRegistrationAccessCredential(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	accessCredentialDescription := `An optional value of the HTTP Authorization header sent with each request to the Scanner Adapter API.
`

	var accessCredentialFlagName string
	if cmdPrefix == "" {
		accessCredentialFlagName = "access_credential"
	} else {
		accessCredentialFlagName = fmt.Sprintf("%v.access_credential", cmdPrefix)
	}

	var accessCredentialFlagDefault string

	_ = cmd.PersistentFlags().String(accessCredentialFlagName, accessCredentialFlagDefault, accessCredentialDescription)

	return nil
}

func registerScannerRegistrationAdapter(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	adapterDescription := `Optional property to describe the name of the scanner registration`

	var adapterFlagName string
	if cmdPrefix == "" {
		adapterFlagName = "adapter"
	} else {
		adapterFlagName = fmt.Sprintf("%v.adapter", cmdPrefix)
	}

	var adapterFlagDefault string

	_ = cmd.PersistentFlags().String(adapterFlagName, adapterFlagDefault, adapterDescription)

	return nil
}

func registerScannerRegistrationAuth(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	authDescription := `Specify what authentication approach is adopted for the HTTP communications.
Supported types Basic', 'Bearer' and api key header 'X-ScannerAdapter-API-Key'
`

	var authFlagName string
	if cmdPrefix == "" {
		authFlagName = "auth"
	} else {
		authFlagName = fmt.Sprintf("%v.auth", cmdPrefix)
	}

	var authFlagDefault string

	_ = cmd.PersistentFlags().String(authFlagName, authFlagDefault, authDescription)

	return nil
}

func registerScannerRegistrationCreateTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createTimeDescription := `The creation time of this registration`

	var createTimeFlagName string
	if cmdPrefix == "" {
		createTimeFlagName = "create_time"
	} else {
		createTimeFlagName = fmt.Sprintf("%v.create_time", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(createTimeFlagName, "", createTimeDescription)

	return nil
}

func registerScannerRegistrationDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `An optional description of this registration.`

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

func registerScannerRegistrationDisabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	disabledDescription := `Indicate whether the registration is enabled or not`

	var disabledFlagName string
	if cmdPrefix == "" {
		disabledFlagName = "disabled"
	} else {
		disabledFlagName = fmt.Sprintf("%v.disabled", cmdPrefix)
	}

	var disabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(disabledFlagName, disabledFlagDefault, disabledDescription)

	return nil
}

func registerScannerRegistrationHealth(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	healthDescription := `Indicate the healthy of the registration`

	var healthFlagName string
	if cmdPrefix == "" {
		healthFlagName = "health"
	} else {
		healthFlagName = fmt.Sprintf("%v.health", cmdPrefix)
	}

	var healthFlagDefault string

	_ = cmd.PersistentFlags().String(healthFlagName, healthFlagDefault, healthDescription)

	return nil
}

func registerScannerRegistrationIsDefault(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isDefaultDescription := `Indicate if the registration is set as the system default one`

	var isDefaultFlagName string
	if cmdPrefix == "" {
		isDefaultFlagName = "is_default"
	} else {
		isDefaultFlagName = fmt.Sprintf("%v.is_default", cmdPrefix)
	}

	var isDefaultFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isDefaultFlagName, isDefaultFlagDefault, isDefaultDescription)

	return nil
}

func registerScannerRegistrationName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `The name of this registration.`

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

func registerScannerRegistrationSkipCertVerify(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	skipCertVerifyDescription := `Indicate if skip the certificate verification when sending HTTP requests`

	var skipCertVerifyFlagName string
	if cmdPrefix == "" {
		skipCertVerifyFlagName = "skip_certVerify"
	} else {
		skipCertVerifyFlagName = fmt.Sprintf("%v.skip_certVerify", cmdPrefix)
	}

	var skipCertVerifyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(skipCertVerifyFlagName, skipCertVerifyFlagDefault, skipCertVerifyDescription)

	return nil
}

func registerScannerRegistrationUpdateTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	updateTimeDescription := `The update time of this registration`

	var updateTimeFlagName string
	if cmdPrefix == "" {
		updateTimeFlagName = "update_time"
	} else {
		updateTimeFlagName = fmt.Sprintf("%v.update_time", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(updateTimeFlagName, "", updateTimeDescription)

	return nil
}

func registerScannerRegistrationURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	urlDescription := `A base URL of the scanner adapter`

	var urlFlagName string
	if cmdPrefix == "" {
		urlFlagName = "url"
	} else {
		urlFlagName = fmt.Sprintf("%v.url", cmdPrefix)
	}

	var urlFlagDefault string

	_ = cmd.PersistentFlags().String(urlFlagName, urlFlagDefault, urlDescription)

	return nil
}

func registerScannerRegistrationUseInternalAddr(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	useInternalAddrDescription := `Indicate whether use internal registry addr for the scanner to pull content or not`

	var useInternalAddrFlagName string
	if cmdPrefix == "" {
		useInternalAddrFlagName = "use_internal_addr"
	} else {
		useInternalAddrFlagName = fmt.Sprintf("%v.use_internal_addr", cmdPrefix)
	}

	var useInternalAddrFlagDefault bool

	_ = cmd.PersistentFlags().Bool(useInternalAddrFlagName, useInternalAddrFlagDefault, useInternalAddrDescription)

	return nil
}

func registerScannerRegistrationUUID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	uuidDescription := `The unique identifier of this registration.`

	var uuidFlagName string
	if cmdPrefix == "" {
		uuidFlagName = "uuid"
	} else {
		uuidFlagName = fmt.Sprintf("%v.uuid", cmdPrefix)
	}

	var uuidFlagDefault string

	_ = cmd.PersistentFlags().String(uuidFlagName, uuidFlagDefault, uuidDescription)

	return nil
}

func registerScannerRegistrationVendor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vendorDescription := `Optional property to describe the vendor of the scanner registration`

	var vendorFlagName string
	if cmdPrefix == "" {
		vendorFlagName = "vendor"
	} else {
		vendorFlagName = fmt.Sprintf("%v.vendor", cmdPrefix)
	}

	var vendorFlagDefault string

	_ = cmd.PersistentFlags().String(vendorFlagName, vendorFlagDefault, vendorDescription)

	return nil
}

func registerScannerRegistrationVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	versionDescription := `Optional property to describe the version of the scanner registration`

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
func retrieveModelScannerRegistrationFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, accessCredentialAdded := retrieveScannerRegistrationAccessCredentialFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || accessCredentialAdded

	err, adapterAdded := retrieveScannerRegistrationAdapterFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || adapterAdded

	err, authAdded := retrieveScannerRegistrationAuthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authAdded

	err, createTimeAdded := retrieveScannerRegistrationCreateTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createTimeAdded

	err, descriptionAdded := retrieveScannerRegistrationDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, disabledAdded := retrieveScannerRegistrationDisabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || disabledAdded

	err, healthAdded := retrieveScannerRegistrationHealthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || healthAdded

	err, isDefaultAdded := retrieveScannerRegistrationIsDefaultFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isDefaultAdded

	err, nameAdded := retrieveScannerRegistrationNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, skipCertVerifyAdded := retrieveScannerRegistrationSkipCertVerifyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || skipCertVerifyAdded

	err, updateTimeAdded := retrieveScannerRegistrationUpdateTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updateTimeAdded

	err, urlAdded := retrieveScannerRegistrationURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || urlAdded

	err, useInternalAddrAdded := retrieveScannerRegistrationUseInternalAddrFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || useInternalAddrAdded

	err, uuidAdded := retrieveScannerRegistrationUUIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || uuidAdded

	err, vendorAdded := retrieveScannerRegistrationVendorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vendorAdded

	err, versionAdded := retrieveScannerRegistrationVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	return nil, retAdded
}

func retrieveScannerRegistrationAccessCredentialFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	accessCredentialFlagName := fmt.Sprintf("%v.access_credential", cmdPrefix)
	if cmd.Flags().Changed(accessCredentialFlagName) {

		var accessCredentialFlagName string
		if cmdPrefix == "" {
			accessCredentialFlagName = "access_credential"
		} else {
			accessCredentialFlagName = fmt.Sprintf("%v.access_credential", cmdPrefix)
		}

		accessCredentialFlagValue, err := cmd.Flags().GetString(accessCredentialFlagName)
		if err != nil {
			return err, false
		}
		m.AccessCredential = accessCredentialFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationAdapterFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	adapterFlagName := fmt.Sprintf("%v.adapter", cmdPrefix)
	if cmd.Flags().Changed(adapterFlagName) {

		var adapterFlagName string
		if cmdPrefix == "" {
			adapterFlagName = "adapter"
		} else {
			adapterFlagName = fmt.Sprintf("%v.adapter", cmdPrefix)
		}

		adapterFlagValue, err := cmd.Flags().GetString(adapterFlagName)
		if err != nil {
			return err, false
		}
		m.Adapter = adapterFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationAuthFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	authFlagName := fmt.Sprintf("%v.auth", cmdPrefix)
	if cmd.Flags().Changed(authFlagName) {

		var authFlagName string
		if cmdPrefix == "" {
			authFlagName = "auth"
		} else {
			authFlagName = fmt.Sprintf("%v.auth", cmdPrefix)
		}

		authFlagValue, err := cmd.Flags().GetString(authFlagName)
		if err != nil {
			return err, false
		}
		m.Auth = authFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationCreateTimeFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	createTimeFlagName := fmt.Sprintf("%v.create_time", cmdPrefix)
	if cmd.Flags().Changed(createTimeFlagName) {

		var createTimeFlagName string
		if cmdPrefix == "" {
			createTimeFlagName = "create_time"
		} else {
			createTimeFlagName = fmt.Sprintf("%v.create_time", cmdPrefix)
		}

		createTimeFlagValueStr, err := cmd.Flags().GetString(createTimeFlagName)
		if err != nil {
			return err, false
		}
		var createTimeFlagValue strfmt.DateTime
		if err := createTimeFlagValue.UnmarshalText([]byte(createTimeFlagValueStr)); err != nil {
			return err, false
		}
		m.CreateTime = createTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationDescriptionFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveScannerRegistrationDisabledFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	disabledFlagName := fmt.Sprintf("%v.disabled", cmdPrefix)
	if cmd.Flags().Changed(disabledFlagName) {

		var disabledFlagName string
		if cmdPrefix == "" {
			disabledFlagName = "disabled"
		} else {
			disabledFlagName = fmt.Sprintf("%v.disabled", cmdPrefix)
		}

		disabledFlagValue, err := cmd.Flags().GetBool(disabledFlagName)
		if err != nil {
			return err, false
		}
		m.Disabled = &disabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationHealthFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	healthFlagName := fmt.Sprintf("%v.health", cmdPrefix)
	if cmd.Flags().Changed(healthFlagName) {

		var healthFlagName string
		if cmdPrefix == "" {
			healthFlagName = "health"
		} else {
			healthFlagName = fmt.Sprintf("%v.health", cmdPrefix)
		}

		healthFlagValue, err := cmd.Flags().GetString(healthFlagName)
		if err != nil {
			return err, false
		}
		m.Health = healthFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationIsDefaultFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isDefaultFlagName := fmt.Sprintf("%v.is_default", cmdPrefix)
	if cmd.Flags().Changed(isDefaultFlagName) {

		var isDefaultFlagName string
		if cmdPrefix == "" {
			isDefaultFlagName = "is_default"
		} else {
			isDefaultFlagName = fmt.Sprintf("%v.is_default", cmdPrefix)
		}

		isDefaultFlagValue, err := cmd.Flags().GetBool(isDefaultFlagName)
		if err != nil {
			return err, false
		}
		m.IsDefault = &isDefaultFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationNameFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationSkipCertVerifyFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	skipCertVerifyFlagName := fmt.Sprintf("%v.skip_certVerify", cmdPrefix)
	if cmd.Flags().Changed(skipCertVerifyFlagName) {

		var skipCertVerifyFlagName string
		if cmdPrefix == "" {
			skipCertVerifyFlagName = "skip_certVerify"
		} else {
			skipCertVerifyFlagName = fmt.Sprintf("%v.skip_certVerify", cmdPrefix)
		}

		skipCertVerifyFlagValue, err := cmd.Flags().GetBool(skipCertVerifyFlagName)
		if err != nil {
			return err, false
		}
		m.SkipCertVerify = &skipCertVerifyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationUpdateTimeFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	updateTimeFlagName := fmt.Sprintf("%v.update_time", cmdPrefix)
	if cmd.Flags().Changed(updateTimeFlagName) {

		var updateTimeFlagName string
		if cmdPrefix == "" {
			updateTimeFlagName = "update_time"
		} else {
			updateTimeFlagName = fmt.Sprintf("%v.update_time", cmdPrefix)
		}

		updateTimeFlagValueStr, err := cmd.Flags().GetString(updateTimeFlagName)
		if err != nil {
			return err, false
		}
		var updateTimeFlagValue strfmt.DateTime
		if err := updateTimeFlagValue.UnmarshalText([]byte(updateTimeFlagValueStr)); err != nil {
			return err, false
		}
		m.UpdateTime = updateTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationURLFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	urlFlagName := fmt.Sprintf("%v.url", cmdPrefix)
	if cmd.Flags().Changed(urlFlagName) {

		var urlFlagName string
		if cmdPrefix == "" {
			urlFlagName = "url"
		} else {
			urlFlagName = fmt.Sprintf("%v.url", cmdPrefix)
		}

		urlFlagValue, err := cmd.Flags().GetString(urlFlagName)
		if err != nil {
			return err, false
		}
		m.URL = urlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationUseInternalAddrFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	useInternalAddrFlagName := fmt.Sprintf("%v.use_internal_addr", cmdPrefix)
	if cmd.Flags().Changed(useInternalAddrFlagName) {

		var useInternalAddrFlagName string
		if cmdPrefix == "" {
			useInternalAddrFlagName = "use_internal_addr"
		} else {
			useInternalAddrFlagName = fmt.Sprintf("%v.use_internal_addr", cmdPrefix)
		}

		useInternalAddrFlagValue, err := cmd.Flags().GetBool(useInternalAddrFlagName)
		if err != nil {
			return err, false
		}
		m.UseInternalAddr = &useInternalAddrFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationUUIDFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	uuidFlagName := fmt.Sprintf("%v.uuid", cmdPrefix)
	if cmd.Flags().Changed(uuidFlagName) {

		var uuidFlagName string
		if cmdPrefix == "" {
			uuidFlagName = "uuid"
		} else {
			uuidFlagName = fmt.Sprintf("%v.uuid", cmdPrefix)
		}

		uuidFlagValue, err := cmd.Flags().GetString(uuidFlagName)
		if err != nil {
			return err, false
		}
		m.UUID = uuidFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationVendorFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vendorFlagName := fmt.Sprintf("%v.vendor", cmdPrefix)
	if cmd.Flags().Changed(vendorFlagName) {

		var vendorFlagName string
		if cmdPrefix == "" {
			vendorFlagName = "vendor"
		} else {
			vendorFlagName = fmt.Sprintf("%v.vendor", cmdPrefix)
		}

		vendorFlagValue, err := cmd.Flags().GetString(vendorFlagName)
		if err != nil {
			return err, false
		}
		m.Vendor = vendorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveScannerRegistrationVersionFlags(depth int, m *models.ScannerRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Version = versionFlagValue

		retAdded = true
	}

	return nil, retAdded
}
