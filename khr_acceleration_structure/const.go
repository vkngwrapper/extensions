package khr_acceleration_structure

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

const (
	// ExtensionName is "VK_KHR_acceleration_structure"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_KHR_acceleration_structure.html
	ExtensionName string = C.VK_KHR_ACCELERATION_STRUCTURE_EXTENSION_NAME
)

// AccelerationStructureType specifies the type of an acceleration structure
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureTypeKHR.html
type AccelerationStructureType int32

var accelerationStructureTypeMapping = make(map[AccelerationStructureType]string)

func (e AccelerationStructureType) Register(str string) {
	accelerationStructureTypeMapping[e] = str
}

func (e AccelerationStructureType) String() string {
	return accelerationStructureTypeMapping[e]
}

const (
	// AccelerationStructureTypeTopLevel is a top-level acceleration structure containing
	// instance data referring to bottom-level acceleration structures.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureTypeKHR.html
	AccelerationStructureTypeTopLevel AccelerationStructureType = C.VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR
	// AccelerationStructureTypeBottomLevel is a bottom-level acceleration structure
	// containing the AABBs or geometry to be intersected.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureTypeKHR.html
	AccelerationStructureTypeBottomLevel AccelerationStructureType = C.VK_ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_KHR
	// AccelerationStructureTypeGeneric is an acceleration structure whose type is determined
	// at build time used for special circumstances. In these cases, the acceleration
	// structure type is not known at creation time, but must be specified at build time as
	// either top or bottom.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureTypeKHR.html
	AccelerationStructureTypeGeneric AccelerationStructureType = C.VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR
)

// BuildAccelerationStructureMode specifies the type of build operation to perform
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkBuildAccelerationStructureModeKHR.html
type BuildAccelerationStructureMode int32

var buildAccelerationStructureModeMapping = make(map[BuildAccelerationStructureMode]string)

func (e BuildAccelerationStructureMode) Register(str string) {
	buildAccelerationStructureModeMapping[e] = str
}

func (e BuildAccelerationStructureMode) String() string {
	return buildAccelerationStructureModeMapping[e]
}

const (
	// BuildAccelerationStructureModeBuild specifies that the destination acceleration structure will be built
	// using the specified geometries.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkBuildAccelerationStructureModeKHR.html
	BuildAccelerationStructureModeBuild BuildAccelerationStructureMode = C.VK_BUILD_ACCELERATION_STRUCTURE_MODE_BUILD_KHR
	// BuildAccelerationStructureModeUpdate specifies that the destination acceleration structure will be built
	// using data in a source acceleration structure, updated by the specified geometries.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkBuildAccelerationStructureModeKHR.html
	BuildAccelerationStructureModeUpdate BuildAccelerationStructureMode = C.VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR
)

// BuildAccelerationStructureFlags specifies additional parameters for acceleration structure
// builds.
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkBuildAccelerationStructureFlagBitsKHR.html
type BuildAccelerationStructureFlags int32

var buildAccelerationStructureFlagsMapping = common.NewFlagStringMapping[BuildAccelerationStructureFlags]()

func (f BuildAccelerationStructureFlags) Register(str string) {
	buildAccelerationStructureFlagsMapping.Register(f, str)
}

func (f BuildAccelerationStructureFlags) String() string {
	return buildAccelerationStructureFlagsMapping.FlagsToString(f)
}

const (
	// BuildAccelerationStructureAllowUpdate specifies that the specified acceleration
	// structure can be updated with a mode of VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR
	// in VkAccelerationStructureBuildGeometryInfoKHR or an update of VK_TRUE in
	// vkCmdBuildAccelerationStructureNV . For sphere and LSS primitives, only positions and
	// radii may be updated, the provided index buffers and flags must remain unchanged from
	// the initial build.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkBuildAccelerationStructureFlagBitsKHR.html
	BuildAccelerationStructureAllowUpdate BuildAccelerationStructureFlags = C.VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR
	// BuildAccelerationStructureAllowCompaction specifies that the specified acceleration structure can act as
	// the source for a copy acceleration structure command with mode of
	// VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR to produce a compacted acceleration structure.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkBuildAccelerationStructureFlagBitsKHR.html
	BuildAccelerationStructureAllowCompaction BuildAccelerationStructureFlags = C.VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR
	// BuildAccelerationStructurePreferFastTrace specifies that the given acceleration structure build should
	// prioritize trace performance over build time.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkBuildAccelerationStructureFlagBitsKHR.html
	BuildAccelerationStructurePreferFastTrace BuildAccelerationStructureFlags = C.VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_KHR
	// BuildAccelerationStructurePreferFastBuild specifies that the given acceleration structure build should
	// prioritize build time over trace performance.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkBuildAccelerationStructureFlagBitsKHR.html
	BuildAccelerationStructurePreferFastBuild BuildAccelerationStructureFlags = C.VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_KHR
	// BuildAccelerationStructureLowMemory specifies that this acceleration structure should minimize the size
	// of the scratch memory and the final result acceleration structure, potentially at the expense of build
	// time or trace performance.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkBuildAccelerationStructureFlagBitsKHR.html
	BuildAccelerationStructureLowMemory BuildAccelerationStructureFlags = C.VK_BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_KHR
)

// GeometryType specifies which type of geometry is provided
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkGeometryTypeKHR.html
type GeometryType int32

var geometryTypeMapping = make(map[GeometryType]string)

func (e GeometryType) Register(str string) {
	geometryTypeMapping[e] = str
}

func (e GeometryType) String() string {
	return geometryTypeMapping[e]
}

const (
	// GeometryTypeTriangles specifies a geometry type consisting of triangles.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkGeometryTypeKHR.html
	GeometryTypeTriangles GeometryType = C.VK_GEOMETRY_TYPE_TRIANGLES_KHR
	// GeometryTypeAABBs specifies a geometry type consisting of axis-aligned bounding boxes.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkGeometryTypeKHR.html
	GeometryTypeAABBs GeometryType = C.VK_GEOMETRY_TYPE_AABBS_KHR
	// GeometryTypeInstances specifies a geometry type consisting of acceleration structure instances.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkGeometryTypeKHR.html
	GeometryTypeInstances GeometryType = C.VK_GEOMETRY_TYPE_INSTANCES_KHR
)

// GeometryFlags specifies additional parameters for a geometry
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkGeometryFlagBitsKHR.html
type GeometryFlags int32

var geometryFlagMapping = common.NewFlagStringMapping[GeometryFlags]()

func (f GeometryFlags) Register(str string) {
	geometryFlagMapping.Register(f, str)
}

func (f GeometryFlags) String() string {
	return geometryFlagMapping.FlagsToString(f)
}

const (
	// GeometryOpaque specifies that this geometry does not invoke the any-hit shaders even if present
	// in a hit group.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkGeometryFlagBitsKHR.html
	GeometryOpaque GeometryFlags = C.VK_GEOMETRY_OPAQUE_BIT_KHR
	// GeometryNoDuplicateAnyHitInvocation specifies that the implementation must only call the any-hit shader
	// a single time for each primitive in this geometry. If this bit is absent an implementation may invoke
	// the any-hit shader more than once for this geometry.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkGeometryFlagBitsKHR.html
	GeometryNoDuplicateAnyHitInvocation GeometryFlags = C.VK_GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_KHR
)

// AccelerationStructureCreateFlags specifies additional creation parameters for acceleration structure
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureCreateFlagBitsKHR.html
type AccelerationStructureCreateFlags int32

var accelerationStructureCreateFlags = common.NewFlagStringMapping[AccelerationStructureCreateFlags]()

func (f AccelerationStructureCreateFlags) Register(str string) {
	accelerationStructureCreateFlags.Register(f, str)
}

func (f AccelerationStructureCreateFlags) String() string {
	return accelerationStructureCreateFlags.FlagsToString(f)
}

const (
	// AccelerationStructureCreateDeviceAddressCaptureReplay specifies that the acceleration structure’s
	// address can be saved and reused on a subsequent run.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureCreateFlagBitsKHR.html
	AccelerationStructureCreateDeviceAddressCaptureReplay AccelerationStructureCreateFlags = C.VK_ACCELERATION_STRUCTURE_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR
)

// CopyAccelerationStructureMode specifies additional operations to perform during the copy
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkCopyAccelerationStructureModeKHR.html
type CopyAccelerationStructureMode int32

var copyAccelerationStructureModeMapping = make(map[CopyAccelerationStructureMode]string)

func (e CopyAccelerationStructureMode) Register(str string) {
	copyAccelerationStructureModeMapping[e] = str
}

func (e CopyAccelerationStructureMode) String() string {
	return copyAccelerationStructureModeMapping[e]
}

const (
	// CopyAccelerationStructureModeClone creates a direct copy of the acceleration structure specified in
	// src into the one specified by dst. The dst acceleration structure must have been created with the same
	// parameters as src. If src contains references to other acceleration structures, dst will reference the
	// same acceleration structures.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkCopyAccelerationStructureModeKHR.html
	CopyAccelerationStructureModeClone CopyAccelerationStructureMode = C.VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR
	// CopyAccelerationStructureModeCompact creates a more compact version of an acceleration structure src
	// into dst. The acceleration structure dst must have been created with a size at least as large as that
	// returned by vkCmdWriteAccelerationStructuresPropertiesNV , vkCmdWriteAccelerationStructuresPropertiesKHR,
	// or vkWriteAccelerationStructuresPropertiesKHR after the build of the acceleration structure specified by
	// src. If src contains references to other acceleration structures, dst will reference the same acceleration
	// structures.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkCopyAccelerationStructureModeKHR.html
	CopyAccelerationStructureModeCompact CopyAccelerationStructureMode = C.VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR
	// CopyAccelerationStructureModeSerialize serializes the acceleration structure to a semi-opaque format
	// which can be reloaded on a compatible implementation.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkCopyAccelerationStructureModeKHR.html
	CopyAccelerationStructureModeSerialize CopyAccelerationStructureMode = C.VK_COPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHR
	// CopyAccelerationStructureModeDeserialize deserializes the semi-opaque serialization format in the buffer
	// to the acceleration structure.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkCopyAccelerationStructureModeKHR.html
	CopyAccelerationStructureModeDeserialize CopyAccelerationStructureMode = C.VK_COPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR
)

// AccelerationStructureBuildType represents the different build types for acceleration structures
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureBuildTypeKHR.html
type AccelerationStructureBuildType int32

var accelerationStructureBuildTypeMapping = make(map[AccelerationStructureBuildType]string)

func (e AccelerationStructureBuildType) Register(str string) {
	accelerationStructureBuildTypeMapping[e] = str
}

func (e AccelerationStructureBuildType) String() string {
	return accelerationStructureBuildTypeMapping[e]
}

const (
	// AccelerationStructureBuildTypeHost requests the memory requirement for operations performed by the host.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureBuildTypeKHR.html
	AccelerationStructureBuildTypeHost AccelerationStructureBuildType = C.VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_KHR
	// AccelerationStructureBuildTypeDevice requests the memory requirement for operations performed by the
	// device.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureBuildTypeKHR.html
	AccelerationStructureBuildTypeDevice AccelerationStructureBuildType = C.VK_ACCELERATION_STRUCTURE_BUILD_TYPE_DEVICE_KHR
	// AccelerationStructureBuildTypeHostOrDevice requests the memory requirement for operations performed by
	// either the host, or the device.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureBuildTypeKHR.html
	AccelerationStructureBuildTypeHostOrDevice AccelerationStructureBuildType = C.VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR
)

// AccelerationStructureCompatibility represents whether an acceleration structure is compatible with a device
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureCompatibilityKHR.html
type AccelerationStructureCompatibility int32

var accelerationStructureCompatibilityMapping = make(map[AccelerationStructureCompatibility]string)

func (e AccelerationStructureCompatibility) Register(str string) {
	accelerationStructureCompatibilityMapping[e] = str
}

func (e AccelerationStructureCompatibility) String() string {
	return accelerationStructureCompatibilityMapping[e]
}

const (
	// AccelerationStructureCompatibilityCompatible indicates the pVersionData version acceleration structure
	// is compatible with device.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureCompatibilityKHR.html
	AccelerationStructureCompatibilityCompatible AccelerationStructureCompatibility = C.VK_ACCELERATION_STRUCTURE_COMPATIBILITY_COMPATIBLE_KHR
	// AccelerationStructureCompatibilityIncompatible indicates the pVersionData version acceleration
	// structure is not compatible with device.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccelerationStructureCompatibilityKHR.html
	AccelerationStructureCompatibilityIncompatible AccelerationStructureCompatibility = C.VK_ACCELERATION_STRUCTURE_COMPATIBILITY_INCOMPATIBLE_KHR
)

func init() {
	AccelerationStructureTypeTopLevel.Register("TopLevel")
	AccelerationStructureTypeBottomLevel.Register("BottomLevel")
	AccelerationStructureTypeGeneric.Register("Generic")

	BuildAccelerationStructureAllowUpdate.Register("AllowUpdate")
	BuildAccelerationStructureAllowCompaction.Register("AllowCompaction")
	BuildAccelerationStructurePreferFastTrace.Register("PreferFastTrace")
	BuildAccelerationStructurePreferFastBuild.Register("PreferFastBuild")
	BuildAccelerationStructureLowMemory.Register("LowMemory")

	BuildAccelerationStructureModeBuild.Register("Build")
	BuildAccelerationStructureModeUpdate.Register("Update")

	GeometryTypeTriangles.Register("Triangles")
	GeometryTypeAABBs.Register("AxisAlignedBoundingBoxes")
	GeometryTypeInstances.Register("Instances")

	GeometryOpaque.Register("Opaque")
	GeometryNoDuplicateAnyHitInvocation.Register("NoDuplicateAnyHitInvocation")

	AccelerationStructureCreateDeviceAddressCaptureReplay.Register("DeviceAddressCaptureReplay")

	CopyAccelerationStructureModeClone.Register("Clone")
	CopyAccelerationStructureModeCompact.Register("Compact")
	CopyAccelerationStructureModeSerialize.Register("Serialize")
	CopyAccelerationStructureModeDeserialize.Register("Deserialize")

	AccelerationStructureBuildTypeHost.Register("Host")
	AccelerationStructureBuildTypeDevice.Register("Device")
	AccelerationStructureBuildTypeHostOrDevice.Register("HostOrDevice")

	AccelerationStructureCompatibilityCompatible.Register("Compatible")
	AccelerationStructureCompatibilityIncompatible.Register("Incompatible")
}
