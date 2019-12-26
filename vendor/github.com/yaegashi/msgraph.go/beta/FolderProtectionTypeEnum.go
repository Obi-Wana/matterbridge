// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// FolderProtectionType undocumented
type FolderProtectionType int

const (
	// FolderProtectionTypeVUserDefined undocumented
	FolderProtectionTypeVUserDefined FolderProtectionType = 0
	// FolderProtectionTypeVEnable undocumented
	FolderProtectionTypeVEnable FolderProtectionType = 1
	// FolderProtectionTypeVAuditMode undocumented
	FolderProtectionTypeVAuditMode FolderProtectionType = 2
	// FolderProtectionTypeVBlockDiskModification undocumented
	FolderProtectionTypeVBlockDiskModification FolderProtectionType = 3
	// FolderProtectionTypeVAuditDiskModification undocumented
	FolderProtectionTypeVAuditDiskModification FolderProtectionType = 4
)

// FolderProtectionTypePUserDefined returns a pointer to FolderProtectionTypeVUserDefined
func FolderProtectionTypePUserDefined() *FolderProtectionType {
	v := FolderProtectionTypeVUserDefined
	return &v
}

// FolderProtectionTypePEnable returns a pointer to FolderProtectionTypeVEnable
func FolderProtectionTypePEnable() *FolderProtectionType {
	v := FolderProtectionTypeVEnable
	return &v
}

// FolderProtectionTypePAuditMode returns a pointer to FolderProtectionTypeVAuditMode
func FolderProtectionTypePAuditMode() *FolderProtectionType {
	v := FolderProtectionTypeVAuditMode
	return &v
}

// FolderProtectionTypePBlockDiskModification returns a pointer to FolderProtectionTypeVBlockDiskModification
func FolderProtectionTypePBlockDiskModification() *FolderProtectionType {
	v := FolderProtectionTypeVBlockDiskModification
	return &v
}

// FolderProtectionTypePAuditDiskModification returns a pointer to FolderProtectionTypeVAuditDiskModification
func FolderProtectionTypePAuditDiskModification() *FolderProtectionType {
	v := FolderProtectionTypeVAuditDiskModification
	return &v
}