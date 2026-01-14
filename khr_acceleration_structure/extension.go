package khr_acceleration_structure

import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	khr_acceleration_structure_loader "github.com/vkngwrapper/extensions/v3/khr_acceleration_structure/loader"
	"github.com/vkngwrapper/extensions/v3/khr_deferred_host_operations"
)

type VulkanExtensionDriver struct {
	loader khr_acceleration_structure_loader.Loader
	device core1_0.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_acceleration_structure loaded
func CreateExtensionDriverFromCoreDriver(loader core1_0.DeviceDriver) ExtensionDriver {
	device := loader.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		loader: khr_acceleration_structure_loader.CreateLoaderFromCore(loader.Loader()),
		device: device,
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(loader khr_acceleration_structure_loader.Loader, device core1_0.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		loader: loader,
		device: device,
	}
}

func (e *VulkanExtensionDriver) BuildAccelerationStructure(deferredOperation khr_deferred_host_operations.DeferredOperation,
	infos []AccelerationStructureBuildGeometryInfo,
	rangeInfos [][]AccelerationStructureBuildRangeInfo,
) (common.VkResult, error) {
	if !deferredOperation.Initialized() {
		return core1_0.VKErrorUnknown, errors.New("deferredOperation cannot be uninitialized")
	}
	if len(infos) != len(rangeInfos) {
		return core1_0.VKErrorUnknown, fmt.Errorf("the length of infos is %d, but the length of rangeInfos is %d - they must be equal", len(infos), len(rangeInfos))
	}

	for index := range infos {
		if len(infos[index].Geometries) != len(rangeInfos[index]) {
			return core1_0.VKErrorUnknown, fmt.Errorf("infos[%d].Geometries has a length of %d, but the length of rangeInfos[%d] is %d - they must be equal", index, len(infos[index].Geometries), index, len(rangeInfos[index]))
		}
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoCount := len(infos)

	infoPtr, err := common.AllocOptionSlice[khr_acceleration_structure_loader.VkAccelerationStructureBuildGeometryInfoKHR](arena, infos)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	rangeInfoAlloc := arena.Malloc(int(unsafe.Sizeof([1]*khr_acceleration_structure_loader.VkAccelerationStructureBuildRangeInfoKHR{})) * infoCount)
	rangeInfoPtr := (**khr_acceleration_structure_loader.VkAccelerationStructureBuildRangeInfoKHR)(rangeInfoAlloc)
	rangeInfoSlice := unsafe.Slice(rangeInfoPtr, infoCount)
	for i := 0; i < infoCount; i++ {
		rangeInfoSlice[i], err = common.AllocSlice[khr_acceleration_structure_loader.VkAccelerationStructureBuildRangeInfoKHR](arena, rangeInfos[i])
		if err != nil {
			return core1_0.VKErrorUnknown, err
		}
	}

	return e.loader.VkBuildAccelerationStructuresKHR(
		e.device.Handle(),
		deferredOperation.Handle(),
		loader.Uint32(infoCount),
		infoPtr,
		rangeInfoPtr,
	)
}

func (e *VulkanExtensionDriver) CmdBuildAccelerationStructuresIndirect(commandBuffer core1_0.CommandBuffer,
	infos []AccelerationStructureBuildGeometryInfo,
	indirectDeviceAddresses []uint64,
	strides []int,
	maxPrimitiveCounts [][]int,
) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	infoCount := len(infos)
	if len(indirectDeviceAddresses) != infoCount {
		panic(fmt.Sprintf("the length of indirectDeviceAddresses is %d, but the length of infos is %d - they must be equal", len(indirectDeviceAddresses), infoCount))
	}
	if len(strides) != infoCount {
		panic(fmt.Sprintf("the length of strides is %d, but the length of the infos is %d - they must be equal", len(strides), infoCount))
	}
	if len(maxPrimitiveCounts) != infoCount {
		panic(fmt.Sprintf("the length of maxPrimitiveCounts is %d, but the length of infos is %d - they must be equal", len(maxPrimitiveCounts), infoCount))
	}
	for infoIndex, info := range infos {
		if len(info.Geometries) != len(maxPrimitiveCounts[infoIndex]) {
			panic(fmt.Sprintf("the length of infos[%d].Geometries is %d, but the length of maxPrimitiveCounts[%d] is %d", infoIndex, len(info.Geometries), infoIndex, len(maxPrimitiveCounts[infoIndex])))
		}
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptionSlice[khr_acceleration_structure_loader.VkAccelerationStructureBuildGeometryInfoKHR](arena, infos)
	if err != nil {
		panic(err.Error())
	}

	deviceAddressAlloc := arena.Malloc(int(unsafe.Sizeof(loader.VkDeviceAddress(0))) * infoCount)
	deviceAddressPtr := (*loader.VkDeviceAddress)(deviceAddressAlloc)
	deviceAddressSlice := unsafe.Slice(deviceAddressPtr, infoCount)

	stridesAlloc := arena.Malloc(int(unsafe.Sizeof(loader.Uint32(0))) * infoCount)
	stridesPtr := (*loader.Uint32)(stridesAlloc)
	stridesSlice := unsafe.Slice(stridesPtr, infoCount)

	primitiveCountAlloc := arena.Malloc(int(unsafe.Sizeof([1]*loader.Uint32{})) * infoCount)
	primitiveCountPtr := (**loader.Uint32)(primitiveCountAlloc)
	primitiveCountSlice := unsafe.Slice(primitiveCountPtr, infoCount)

	for i := 0; i < infoCount; i++ {
		deviceAddressSlice[i] = loader.VkDeviceAddress(indirectDeviceAddresses[i])
		stridesSlice[i] = loader.Uint32(strides[i])

		primitivesAlloc := arena.Malloc(int(unsafe.Sizeof(loader.Uint32(0))) * len(maxPrimitiveCounts[i]))
		primitivesPtr := (*loader.Uint32)(primitivesAlloc)
		primitivesSlice := unsafe.Slice(primitivesPtr, len(maxPrimitiveCounts[i]))

		for primIndex, prim := range maxPrimitiveCounts[i] {
			primitivesSlice[primIndex] = loader.Uint32(prim)
		}
		primitiveCountSlice[i] = primitivesPtr
	}

	e.loader.VkCmdBuildAccelerationStructuresIndirectKHR(
		commandBuffer.Handle(),
		loader.Uint32(infoCount),
		infoPtr,
		deviceAddressPtr,
		stridesPtr,
		primitiveCountPtr,
	)
}

func (e *VulkanExtensionDriver) CmdBuildAccelerationStructures(commandBuffer core1_0.CommandBuffer,
	infos []AccelerationStructureBuildGeometryInfo,
	rangeInfos [][]AccelerationStructureBuildRangeInfo,
) {
	if !commandBuffer.Initialized() {
		panic("deferredOperation cannot be uninitialized")
	}
	if len(infos) != len(rangeInfos) {
		panic(fmt.Sprintf("the length of infos is %d, but the length of rangeInfos is %d - they must be equal", len(infos), len(rangeInfos)))
	}

	for index := range infos {
		if len(infos[index].Geometries) != len(rangeInfos[index]) {
			panic(fmt.Sprintf("infos[%d].Geometries has a length of %d, but the length of rangeInfos[%d] is %d - they must be equal", index, len(infos[index].Geometries), index, len(rangeInfos[index])))
		}
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoCount := len(infos)

	infoPtr, err := common.AllocOptionSlice[khr_acceleration_structure_loader.VkAccelerationStructureBuildGeometryInfoKHR](arena, infos)
	if err != nil {
		panic(err)
	}

	rangeInfoAlloc := arena.Malloc(int(unsafe.Sizeof([1]*khr_acceleration_structure_loader.VkAccelerationStructureBuildRangeInfoKHR{})) * infoCount)
	rangeInfoPtr := (**khr_acceleration_structure_loader.VkAccelerationStructureBuildRangeInfoKHR)(rangeInfoAlloc)
	rangeInfoSlice := unsafe.Slice(rangeInfoPtr, infoCount)
	for i := 0; i < infoCount; i++ {
		rangeInfoSlice[i], err = common.AllocSlice[khr_acceleration_structure_loader.VkAccelerationStructureBuildRangeInfoKHR](arena, rangeInfos[i])
		if err != nil {
			panic(err)
		}
	}

	e.loader.VkCmdBuildAccelerationStructuresKHR(
		commandBuffer.Handle(),
		loader.Uint32(infoCount),
		infoPtr,
		rangeInfoPtr,
	)
}

func (e *VulkanExtensionDriver) CmdCopyAccelerationStructure(commandBuffer core1_0.CommandBuffer,
	info CopyAccelerationStructureInfo,
) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, info)
	if err != nil {
		panic(err)
	}

	e.loader.VkCmdCopyAccelerationStructureKHR(commandBuffer.Handle(), (*khr_acceleration_structure_loader.VkCopyAccelerationStructureInfoKHR)(infoPtr))
}

func (e *VulkanExtensionDriver) CmdCopyAccelerationStructureToMemory(commandBuffer core1_0.CommandBuffer,
	info CopyAccelerationStructureToMemoryInfo,
) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, info)
	if err != nil {
		panic(err)
	}

	e.loader.VkCmdCopyAccelerationStructureToMemoryKHR(commandBuffer.Handle(), (*khr_acceleration_structure_loader.VkCopyAccelerationStructureToMemoryInfoKHR)(infoPtr))
}

func (e *VulkanExtensionDriver) CmdCopyMemoryToAccelerationStructure(commandBuffer core1_0.CommandBuffer,
	info CopyMemoryToAccelerationStructureInfo,
) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, info)
	if err != nil {
		panic(err)
	}

	e.loader.VkCmdCopyMemoryToAccelerationStructureKHR(commandBuffer.Handle(), (*khr_acceleration_structure_loader.VkCopyMemoryToAccelerationStructureInfoKHR)(infoPtr))
}

func (e *VulkanExtensionDriver) CmdWriteAccelerationStructuresProperties(commandBuffer core1_0.CommandBuffer,
	accelerationStructures []AccelerationStructure,
	queryType core1_0.QueryType,
	queryPool core1_0.QueryPool,
	firstQuery int,
) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	structuresAlloc := arena.Malloc(int(unsafe.Sizeof([1]khr_acceleration_structure_loader.VkAccelerationStructureKHR{})) * len(accelerationStructures))
	structuresPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureKHR)(structuresAlloc)
	structuresSlice := unsafe.Slice(structuresPtr, len(accelerationStructures))

	for index, structure := range accelerationStructures {
		structuresSlice[index] = structure.Handle()
	}

	e.loader.VkCmdWriteAccelerationStructuresPropertiesKHR(
		commandBuffer.Handle(),
		loader.Uint32(len(accelerationStructures)),
		structuresPtr,
		loader.VkQueryType(queryType),
		queryPool.Handle(),
		loader.Uint32(firstQuery),
	)
}

func (e *VulkanExtensionDriver) CopyAccelerationStructure(deferredOperation khr_deferred_host_operations.DeferredOperation,
	info CopyAccelerationStructureInfo,
) (common.VkResult, error) {
	if !deferredOperation.Initialized() {
		panic("deferredOperation cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, info)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.loader.VkCopyAccelerationStructureKHR(e.device.Handle(),
		deferredOperation.Handle(),
		(*khr_acceleration_structure_loader.VkCopyAccelerationStructureInfoKHR)(infoPtr),
	)
}

func (e *VulkanExtensionDriver) CopyAccelerationStructureToMemory(deferredOperation khr_deferred_host_operations.DeferredOperation,
	info CopyAccelerationStructureToMemoryInfo,
) (common.VkResult, error) {
	if !deferredOperation.Initialized() {
		panic("deferredOperation cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, info)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.loader.VkCopyAccelerationStructureToMemoryKHR(e.device.Handle(),
		deferredOperation.Handle(),
		(*khr_acceleration_structure_loader.VkCopyAccelerationStructureToMemoryInfoKHR)(infoPtr),
	)
}

func (e *VulkanExtensionDriver) CopyMemoryToAccelerationStructure(deferredOperation khr_deferred_host_operations.DeferredOperation,
	info CopyMemoryToAccelerationStructureInfo,
) (common.VkResult, error) {
	if !deferredOperation.Initialized() {
		panic("deferredOperation cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, info)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.loader.VkCopyMemoryToAccelerationStructureKHR(e.device.Handle(),
		deferredOperation.Handle(),
		(*khr_acceleration_structure_loader.VkCopyMemoryToAccelerationStructureInfoKHR)(infoPtr),
	)
}

func (e *VulkanExtensionDriver) CreateAccelerationStructure(createInfo AccelerationStructureCreateInfo,
	allocator *loader.AllocationCallbacks,
) (AccelerationStructure, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, createInfo)
	if err != nil {
		return AccelerationStructure{}, core1_0.VKErrorUnknown, err
	}

	structureAlloc := arena.Malloc(int(unsafe.Sizeof([1]*khr_acceleration_structure_loader.VkAccelerationStructureKHR{})))
	structurePtr := (*khr_acceleration_structure_loader.VkAccelerationStructureKHR)(structureAlloc)

	res, err := e.loader.VkCreateAccelerationStructureKHR(
		e.device.Handle(),
		(*khr_acceleration_structure_loader.VkAccelerationStructureCreateInfoKHR)(infoPtr),
		allocator.Handle(),
		structurePtr,
	)
	if err != nil {
		return AccelerationStructure{}, res, err
	}

	return AccelerationStructure{
		handle:     *structurePtr,
		device:     e.device.Handle(),
		apiVersion: e.device.APIVersion(),
	}, res, nil
}

func (e *VulkanExtensionDriver) DestroyAccelerationStructure(accelerationStructure AccelerationStructure,
	pAllocator *loader.AllocationCallbacks,
) {
	if !accelerationStructure.Initialized() {
		panic("accelerationStructure cannot be uninitialized")
	}

	e.loader.VkDestroyAccelerationStructureKHR(e.device.Handle(), accelerationStructure.Handle(), pAllocator.Handle())
}

func (e *VulkanExtensionDriver) GetAccelerationStructureBuildSizes(buildType AccelerationStructureBuildType,
	buildInfo AccelerationStructureBuildGeometryInfo,
	maxPrimitiveCounts []int,
	out *AccelerationStructureBuildSizesInfo) {
	if len(buildInfo.Geometries) != len(maxPrimitiveCounts) {
		panic(fmt.Sprintf("buildInfo.Geometries has length %d and maxPrimitiveCounts has length %d -- they must be equal", len(buildInfo.Geometries), len(maxPrimitiveCounts)))
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	buildInfoPtr, err := common.AllocOptions(arena, buildInfo)
	if err != nil {
		panic(err)
	}

	primitiveCountsAlloc := arena.Malloc(int(unsafe.Sizeof(loader.Uint32(0))) * len(maxPrimitiveCounts))
	primitiveCountsPtr := (*loader.Uint32)(primitiveCountsAlloc)
	primitiveCountsSlice := unsafe.Slice(primitiveCountsPtr, len(maxPrimitiveCounts))

	for index, val := range maxPrimitiveCounts {
		primitiveCountsSlice[index] = loader.Uint32(val)
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		panic(err)
	}

	e.loader.VkGetAccelerationStructureBuildSizesKHR(
		e.device.Handle(),
		khr_acceleration_structure_loader.VkAccelerationStructureBuildTypeKHR(buildType),
		(*khr_acceleration_structure_loader.VkAccelerationStructureBuildGeometryInfoKHR)(buildInfoPtr),
		primitiveCountsPtr,
		(*khr_acceleration_structure_loader.VkAccelerationStructureBuildSizesInfoKHR)(outDataPtr),
	)

	err = common.PopulateOutData(out, outDataPtr)
	if err != nil {
		panic(err)
	}
}

func (e *VulkanExtensionDriver) GetAccelerationStructureDeviceAddress(o AccelerationStructureDeviceAddressInfo) uint64 {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		panic(err)
	}

	return uint64(e.loader.VkGetAccelerationStructureDeviceAddressKHR(
		e.device.Handle(),
		(*khr_acceleration_structure_loader.VkAccelerationStructureDeviceAddressInfoKHR)(infoPtr),
	))
}

func (e *VulkanExtensionDriver) GetDeviceAccelerationStructureCompatibility(o AccelerationStructureVersionInfo) AccelerationStructureCompatibility {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		panic(err)
	}

	compatibilityAlloc := arena.Malloc(int(unsafe.Sizeof([1]*khr_acceleration_structure_loader.VkAccelerationStructureCompatibilityKHR{})))
	compatibilityPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureCompatibilityKHR)(compatibilityAlloc)

	e.loader.VkGetDeviceAccelerationStructureCompatibilityKHR(
		e.device.Handle(),
		(*khr_acceleration_structure_loader.VkAccelerationStructureVersionInfoKHR)(infoPtr),
		compatibilityPtr,
	)

	return AccelerationStructureCompatibility(*compatibilityPtr)
}

func (e *VulkanExtensionDriver) WriteAccelerationStructuresProperties(accelerationStructures []AccelerationStructure,
	queryType core1_0.QueryType,
	dataSize int,
	data unsafe.Pointer,
	stride int,
) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	structuresAlloc := arena.Malloc(int(unsafe.Sizeof([1]khr_acceleration_structure_loader.VkAccelerationStructureKHR{})) * len(accelerationStructures))
	structuresPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureKHR)(structuresAlloc)
	structuresSlice := unsafe.Slice(structuresPtr, len(accelerationStructures))

	for index, structure := range accelerationStructures {
		structuresSlice[index] = structure.Handle()
	}

	return e.loader.VkWriteAccelerationStructuresPropertiesKHR(
		e.device.Handle(),
		loader.Uint32(len(accelerationStructures)),
		structuresPtr,
		loader.VkQueryType(queryType),
		loader.Size(dataSize),
		data,
		loader.Size(stride),
	)
}
