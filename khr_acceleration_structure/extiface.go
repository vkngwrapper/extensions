package khr_acceleration_structure

import (
	"unsafe"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_deferred_host_operations"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_acceleration_structure

type ExtensionDriver interface {
	BuildAccelerationStructure(deferredOperation khr_deferred_host_operations.DeferredOperation,
		infos []AccelerationStructureBuildGeometryInfo,
		rangeInfos [][]AccelerationStructureBuildRangeInfo,
	) (common.VkResult, error)
	CmdBuildAccelerationStructuresIndirect(commandBuffer core1_0.CommandBuffer,
		infos []AccelerationStructureBuildGeometryInfo,
		indirectDeviceAddresses []uint64,
		strides []int,
		maxPrimitiveCounts [][]int,
	)
	CmdBuildAccelerationStructures(commandBuffer core1_0.CommandBuffer,
		infos []AccelerationStructureBuildGeometryInfo,
		rangeInfos [][]AccelerationStructureBuildRangeInfo,
	)
	CmdCopyAccelerationStructure(commandBuffer core1_0.CommandBuffer,
		info CopyAccelerationStructureInfo,
	)
	CmdCopyAccelerationStructureToMemory(commandBuffer core1_0.CommandBuffer,
		info CopyAccelerationStructureToMemoryInfo,
	)
	CmdCopyMemoryToAccelerationStructure(commandBuffer core1_0.CommandBuffer,
		info CopyMemoryToAccelerationStructureInfo,
	)
	CmdWriteAccelerationStructuresProperties(commandBuffer core1_0.CommandBuffer,
		accelerationStructures []AccelerationStructure,
		queryType core1_0.QueryType,
		queryPool core1_0.QueryPool,
		firstQuery int,
	)
	CopyAccelerationStructure(deferredOperation khr_deferred_host_operations.DeferredOperation,
		info CopyAccelerationStructureInfo,
	) (common.VkResult, error)
	CopyAccelerationStructureToMemory(deferredOperation khr_deferred_host_operations.DeferredOperation,
		info CopyAccelerationStructureToMemoryInfo,
	) (common.VkResult, error)
	CopyMemoryToAccelerationStructure(deferredOperation khr_deferred_host_operations.DeferredOperation,
		info CopyMemoryToAccelerationStructureInfo,
	) (common.VkResult, error)
	CreateAccelerationStructure(createInfo AccelerationStructureCreateInfo,
		allocator *loader.AllocationCallbacks,
	) (AccelerationStructure, common.VkResult, error)
	DestroyAccelerationStructure(accelerationStructure AccelerationStructure,
		pAllocator *loader.AllocationCallbacks,
	)
	GetAccelerationStructureBuildSizes(buildType AccelerationStructureBuildType,
		buildInfo AccelerationStructureBuildGeometryInfo,
		maxPrimitiveCounts []int,
		out *AccelerationStructureBuildSizesInfo)
	GetAccelerationStructureDeviceAddress(o AccelerationStructureDeviceAddressInfo) uint64
	GetDeviceAccelerationStructureCompatibility(o AccelerationStructureVersionInfo) AccelerationStructureCompatibility
	WriteAccelerationStructuresProperties(accelerationStructures []AccelerationStructure,
		queryType core1_0.QueryType,
		dataSize int,
		data unsafe.Pointer,
		stride int,
	) (common.VkResult, error)
}
