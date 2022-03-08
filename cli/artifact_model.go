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

// Schema cli for Artifact

// register flags to command
func registerModelArtifactFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerArtifactAdditionLinks(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactAnnotations(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactDigest(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactExtraAttrs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactIcon(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactLabels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactManifestMediaType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactMediaType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactProjectID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactPullTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactPushTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactReferences(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactRepositoryID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactScanOverview(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactTags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerArtifactAdditionLinks(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: addition_links AdditionLinks map type is not supported by go-swagger cli yet

	return nil
}

func registerArtifactAnnotations(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: annotations Annotations map type is not supported by go-swagger cli yet

	return nil
}

func registerArtifactDigest(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	digestDescription := `The digest of the artifact`

	var digestFlagName string
	if cmdPrefix == "" {
		digestFlagName = "digest"
	} else {
		digestFlagName = fmt.Sprintf("%v.digest", cmdPrefix)
	}

	var digestFlagDefault string

	_ = cmd.PersistentFlags().String(digestFlagName, digestFlagDefault, digestDescription)

	return nil
}

func registerArtifactExtraAttrs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: extra_attrs ExtraAttrs map type is not supported by go-swagger cli yet

	return nil
}

func registerArtifactIcon(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	iconDescription := `The digest of the icon`

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

func registerArtifactID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `The ID of the artifact`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault int64

	_ = cmd.PersistentFlags().Int64(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerArtifactLabels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: labels []*Label array type is not supported by go-swagger cli yet

	return nil
}

func registerArtifactManifestMediaType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	manifestMediaTypeDescription := `The manifest media type of the artifact`

	var manifestMediaTypeFlagName string
	if cmdPrefix == "" {
		manifestMediaTypeFlagName = "manifest_media_type"
	} else {
		manifestMediaTypeFlagName = fmt.Sprintf("%v.manifest_media_type", cmdPrefix)
	}

	var manifestMediaTypeFlagDefault string

	_ = cmd.PersistentFlags().String(manifestMediaTypeFlagName, manifestMediaTypeFlagDefault, manifestMediaTypeDescription)

	return nil
}

func registerArtifactMediaType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	mediaTypeDescription := `The media type of the artifact`

	var mediaTypeFlagName string
	if cmdPrefix == "" {
		mediaTypeFlagName = "media_type"
	} else {
		mediaTypeFlagName = fmt.Sprintf("%v.media_type", cmdPrefix)
	}

	var mediaTypeFlagDefault string

	_ = cmd.PersistentFlags().String(mediaTypeFlagName, mediaTypeFlagDefault, mediaTypeDescription)

	return nil
}

func registerArtifactProjectID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	projectIdDescription := `The ID of the project that the artifact belongs to`

	var projectIdFlagName string
	if cmdPrefix == "" {
		projectIdFlagName = "project_id"
	} else {
		projectIdFlagName = fmt.Sprintf("%v.project_id", cmdPrefix)
	}

	var projectIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(projectIdFlagName, projectIdFlagDefault, projectIdDescription)

	return nil
}

func registerArtifactPullTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pullTimeDescription := `The latest pull time of the artifact`

	var pullTimeFlagName string
	if cmdPrefix == "" {
		pullTimeFlagName = "pull_time"
	} else {
		pullTimeFlagName = fmt.Sprintf("%v.pull_time", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(pullTimeFlagName, "", pullTimeDescription)

	return nil
}

func registerArtifactPushTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pushTimeDescription := `The push time of the artifact`

	var pushTimeFlagName string
	if cmdPrefix == "" {
		pushTimeFlagName = "push_time"
	} else {
		pushTimeFlagName = fmt.Sprintf("%v.push_time", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(pushTimeFlagName, "", pushTimeDescription)

	return nil
}

func registerArtifactReferences(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: references []*Reference array type is not supported by go-swagger cli yet

	return nil
}

func registerArtifactRepositoryID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	repositoryIdDescription := `The ID of the repository that the artifact belongs to`

	var repositoryIdFlagName string
	if cmdPrefix == "" {
		repositoryIdFlagName = "repository_id"
	} else {
		repositoryIdFlagName = fmt.Sprintf("%v.repository_id", cmdPrefix)
	}

	var repositoryIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(repositoryIdFlagName, repositoryIdFlagDefault, repositoryIdDescription)

	return nil
}

func registerArtifactScanOverview(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: scan_overview ScanOverview map type is not supported by go-swagger cli yet

	return nil
}

func registerArtifactSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeDescription := `The size of the artifact`

	var sizeFlagName string
	if cmdPrefix == "" {
		sizeFlagName = "size"
	} else {
		sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
	}

	var sizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sizeFlagName, sizeFlagDefault, sizeDescription)

	return nil
}

func registerArtifactTags(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: tags []*Tag array type is not supported by go-swagger cli yet

	return nil
}

func registerArtifactType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `The type of the artifact, e.g. image, chart, etc`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelArtifactFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, additionLinksAdded := retrieveArtifactAdditionLinksFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || additionLinksAdded

	err, annotationsAdded := retrieveArtifactAnnotationsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || annotationsAdded

	err, digestAdded := retrieveArtifactDigestFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || digestAdded

	err, extraAttrsAdded := retrieveArtifactExtraAttrsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || extraAttrsAdded

	err, iconAdded := retrieveArtifactIconFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || iconAdded

	err, idAdded := retrieveArtifactIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, labelsAdded := retrieveArtifactLabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelsAdded

	err, manifestMediaTypeAdded := retrieveArtifactManifestMediaTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || manifestMediaTypeAdded

	err, mediaTypeAdded := retrieveArtifactMediaTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mediaTypeAdded

	err, projectIdAdded := retrieveArtifactProjectIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || projectIdAdded

	err, pullTimeAdded := retrieveArtifactPullTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pullTimeAdded

	err, pushTimeAdded := retrieveArtifactPushTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pushTimeAdded

	err, referencesAdded := retrieveArtifactReferencesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || referencesAdded

	err, repositoryIdAdded := retrieveArtifactRepositoryIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || repositoryIdAdded

	err, scanOverviewAdded := retrieveArtifactScanOverviewFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scanOverviewAdded

	err, sizeAdded := retrieveArtifactSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeAdded

	err, tagsAdded := retrieveArtifactTagsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tagsAdded

	err, typeAdded := retrieveArtifactTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveArtifactAdditionLinksFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	additionLinksFlagName := fmt.Sprintf("%v.addition_links", cmdPrefix)
	if cmd.Flags().Changed(additionLinksFlagName) {
		// warning: addition_links map type AdditionLinks is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveArtifactAnnotationsFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	annotationsFlagName := fmt.Sprintf("%v.annotations", cmdPrefix)
	if cmd.Flags().Changed(annotationsFlagName) {
		// warning: annotations map type Annotations is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveArtifactDigestFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	digestFlagName := fmt.Sprintf("%v.digest", cmdPrefix)
	if cmd.Flags().Changed(digestFlagName) {

		var digestFlagName string
		if cmdPrefix == "" {
			digestFlagName = "digest"
		} else {
			digestFlagName = fmt.Sprintf("%v.digest", cmdPrefix)
		}

		digestFlagValue, err := cmd.Flags().GetString(digestFlagName)
		if err != nil {
			return err, false
		}
		m.Digest = digestFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveArtifactExtraAttrsFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	extraAttrsFlagName := fmt.Sprintf("%v.extra_attrs", cmdPrefix)
	if cmd.Flags().Changed(extraAttrsFlagName) {
		// warning: extra_attrs map type ExtraAttrs is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveArtifactIconFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Icon = iconFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveArtifactIDFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetInt64(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveArtifactLabelsFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	labelsFlagName := fmt.Sprintf("%v.labels", cmdPrefix)
	if cmd.Flags().Changed(labelsFlagName) {
		// warning: labels array type []*Label is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveArtifactManifestMediaTypeFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	manifestMediaTypeFlagName := fmt.Sprintf("%v.manifest_media_type", cmdPrefix)
	if cmd.Flags().Changed(manifestMediaTypeFlagName) {

		var manifestMediaTypeFlagName string
		if cmdPrefix == "" {
			manifestMediaTypeFlagName = "manifest_media_type"
		} else {
			manifestMediaTypeFlagName = fmt.Sprintf("%v.manifest_media_type", cmdPrefix)
		}

		manifestMediaTypeFlagValue, err := cmd.Flags().GetString(manifestMediaTypeFlagName)
		if err != nil {
			return err, false
		}
		m.ManifestMediaType = manifestMediaTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveArtifactMediaTypeFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	mediaTypeFlagName := fmt.Sprintf("%v.media_type", cmdPrefix)
	if cmd.Flags().Changed(mediaTypeFlagName) {

		var mediaTypeFlagName string
		if cmdPrefix == "" {
			mediaTypeFlagName = "media_type"
		} else {
			mediaTypeFlagName = fmt.Sprintf("%v.media_type", cmdPrefix)
		}

		mediaTypeFlagValue, err := cmd.Flags().GetString(mediaTypeFlagName)
		if err != nil {
			return err, false
		}
		m.MediaType = mediaTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveArtifactProjectIDFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	projectIdFlagName := fmt.Sprintf("%v.project_id", cmdPrefix)
	if cmd.Flags().Changed(projectIdFlagName) {

		var projectIdFlagName string
		if cmdPrefix == "" {
			projectIdFlagName = "project_id"
		} else {
			projectIdFlagName = fmt.Sprintf("%v.project_id", cmdPrefix)
		}

		projectIdFlagValue, err := cmd.Flags().GetInt64(projectIdFlagName)
		if err != nil {
			return err, false
		}
		m.ProjectID = projectIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveArtifactPullTimeFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pullTimeFlagName := fmt.Sprintf("%v.pull_time", cmdPrefix)
	if cmd.Flags().Changed(pullTimeFlagName) {

		var pullTimeFlagName string
		if cmdPrefix == "" {
			pullTimeFlagName = "pull_time"
		} else {
			pullTimeFlagName = fmt.Sprintf("%v.pull_time", cmdPrefix)
		}

		pullTimeFlagValueStr, err := cmd.Flags().GetString(pullTimeFlagName)
		if err != nil {
			return err, false
		}
		var pullTimeFlagValue strfmt.DateTime
		if err := pullTimeFlagValue.UnmarshalText([]byte(pullTimeFlagValueStr)); err != nil {
			return err, false
		}
		m.PullTime = pullTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveArtifactPushTimeFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pushTimeFlagName := fmt.Sprintf("%v.push_time", cmdPrefix)
	if cmd.Flags().Changed(pushTimeFlagName) {

		var pushTimeFlagName string
		if cmdPrefix == "" {
			pushTimeFlagName = "push_time"
		} else {
			pushTimeFlagName = fmt.Sprintf("%v.push_time", cmdPrefix)
		}

		pushTimeFlagValueStr, err := cmd.Flags().GetString(pushTimeFlagName)
		if err != nil {
			return err, false
		}
		var pushTimeFlagValue strfmt.DateTime
		if err := pushTimeFlagValue.UnmarshalText([]byte(pushTimeFlagValueStr)); err != nil {
			return err, false
		}
		m.PushTime = pushTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveArtifactReferencesFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	referencesFlagName := fmt.Sprintf("%v.references", cmdPrefix)
	if cmd.Flags().Changed(referencesFlagName) {
		// warning: references array type []*Reference is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveArtifactRepositoryIDFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	repositoryIdFlagName := fmt.Sprintf("%v.repository_id", cmdPrefix)
	if cmd.Flags().Changed(repositoryIdFlagName) {

		var repositoryIdFlagName string
		if cmdPrefix == "" {
			repositoryIdFlagName = "repository_id"
		} else {
			repositoryIdFlagName = fmt.Sprintf("%v.repository_id", cmdPrefix)
		}

		repositoryIdFlagValue, err := cmd.Flags().GetInt64(repositoryIdFlagName)
		if err != nil {
			return err, false
		}
		m.RepositoryID = repositoryIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveArtifactScanOverviewFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scanOverviewFlagName := fmt.Sprintf("%v.scan_overview", cmdPrefix)
	if cmd.Flags().Changed(scanOverviewFlagName) {
		// warning: scan_overview map type ScanOverview is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveArtifactSizeFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sizeFlagName := fmt.Sprintf("%v.size", cmdPrefix)
	if cmd.Flags().Changed(sizeFlagName) {

		var sizeFlagName string
		if cmdPrefix == "" {
			sizeFlagName = "size"
		} else {
			sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
		}

		sizeFlagValue, err := cmd.Flags().GetInt64(sizeFlagName)
		if err != nil {
			return err, false
		}
		m.Size = sizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveArtifactTagsFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tagsFlagName := fmt.Sprintf("%v.tags", cmdPrefix)
	if cmd.Flags().Changed(tagsFlagName) {
		// warning: tags array type []*Tag is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveArtifactTypeFlags(depth int, m *models.Artifact, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}