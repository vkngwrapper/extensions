package khr_acceleration_structure_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"github.com/vkngwrapper/extensions/v3/khr_acceleration_structure"
	khr_acceleration_structure_loader "github.com/vkngwrapper/extensions/v3/khr_acceleration_structure/loader"
	mock_acceleration_structure "github.com/vkngwrapper/extensions/v3/khr_acceleration_structure/mocks"
	khr_deferred_host_operations_loader "github.com/vkngwrapper/extensions/v3/khr_deferred_host_operations/loader"
	mock_deferred_host_operations "github.com/vkngwrapper/extensions/v3/khr_deferred_host_operations/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtensionDriver_BuildAccelerationStructure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	deferredOperation := mock_deferred_host_operations.NewDummyDeferredOperation(device)
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	srcStructure := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	dstStructure1 := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	dstStructure2 := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkBuildAccelerationStructuresKHR(
		device.Handle(),
		deferredOperation.Handle(),
		loader.Uint32(2),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		deferredOperation khr_deferred_host_operations_loader.VkDeferredOperationKHR,
		infoCount loader.Uint32,
		pInfos *khr_acceleration_structure_loader.VkAccelerationStructureBuildGeometryInfoKHR,
		ppBuildRangeInfos **khr_acceleration_structure_loader.VkAccelerationStructureBuildRangeInfoKHR,
	) (common.VkResult, error) {
		infoSlice := unsafe.Slice(pInfos, 2)

		info := reflect.ValueOf(&infoSlice[0]).Elem()
		require.Equal(t, uint64(1000150000), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_GEOMETRY_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), info.FieldByName("_type").Uint())
		require.Equal(t, uint64(2), info.FieldByName("flags").Uint())
		require.Equal(t, uint64(1), info.FieldByName("mode").Uint())
		require.Equal(t, srcStructure.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("srcAccelerationStructure").UnsafePointer()))
		require.Equal(t, dstStructure1.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("dstAccelerationStructure").UnsafePointer()))
		require.Equal(t, uint64(2), info.FieldByName("geometryCount").Uint())
		require.True(t, info.FieldByName("ppGeometries").IsNil())
		scratchData := info.FieldByName("scratchData").UnsafeAddr()
		require.Equal(t, uint64(13), *(*uint64)(unsafe.Pointer(scratchData)))

		geometriesPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryKHR)(info.FieldByName("pGeometries").UnsafePointer())
		geometriesSlice := unsafe.Slice(geometriesPtr, 2)

		geometry := reflect.ValueOf(&geometriesSlice[0]).Elem()
		require.Equal(t, uint64(1000150006), geometry.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR
		require.True(t, geometry.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), geometry.FieldByName("geometryType").Uint())
		require.Equal(t, uint64(1), geometry.FieldByName("flags").Uint())

		geometryDataUnsafe := geometry.FieldByName("geometry").UnsafeAddr()
		geometryInstanceDataPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryInstancesDataKHR)(unsafe.Pointer(geometryDataUnsafe))
		geometryData := reflect.ValueOf(geometryInstanceDataPtr).Elem()
		require.Equal(t, uint64(1000150004), geometryData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_INSTANCES_DATA_KHR
		require.True(t, geometryData.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), geometryData.FieldByName("arrayOfPointers").Uint())
		data := geometryData.FieldByName("data").UnsafeAddr()
		require.Equal(t, uint64(17), *(*uint64)(unsafe.Pointer(data)))

		geometry = reflect.ValueOf(&geometriesSlice[1]).Elem()
		require.Equal(t, uint64(1000150006), geometry.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR
		require.True(t, geometry.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), geometry.FieldByName("geometryType").Uint())
		require.Equal(t, uint64(2), geometry.FieldByName("flags").Uint())

		geometryDataUnsafe = geometry.FieldByName("geometry").UnsafeAddr()
		geometryAABBDataPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryAabbsDataKHR)(unsafe.Pointer(geometryDataUnsafe))
		geometryData = reflect.ValueOf(geometryAABBDataPtr).Elem()
		require.Equal(t, uint64(1000150003), geometryData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_AABBS_DATA_KHR
		require.True(t, geometryData.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(24), geometryData.FieldByName("stride").Uint())
		data = geometryData.FieldByName("data").UnsafeAddr()
		require.Equal(t, uint64(19), *(*uint64)(unsafe.Pointer(data)))

		info = reflect.ValueOf(&infoSlice[1]).Elem()
		require.Equal(t, uint64(1000150000), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_GEOMETRY_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), info.FieldByName("_type").Uint())
		require.Equal(t, uint64(16), info.FieldByName("flags").Uint())
		require.Equal(t, uint64(0), info.FieldByName("mode").Uint())
		require.Equal(t, khr_acceleration_structure_loader.VkAccelerationStructureKHR(0), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("srcAccelerationStructure").UnsafePointer()))
		require.Equal(t, dstStructure2.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("dstAccelerationStructure").UnsafePointer()))
		require.Equal(t, uint64(1), info.FieldByName("geometryCount").Uint())
		require.True(t, info.FieldByName("ppGeometries").IsNil())
		scratchData = info.FieldByName("scratchData").UnsafeAddr()
		require.Equal(t, uint64(47), *(*uint64)(unsafe.Pointer(scratchData)))

		geometriesPtr = (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryKHR)(info.FieldByName("pGeometries").UnsafePointer())
		geometriesSlice = unsafe.Slice(geometriesPtr, 1)

		geometry = reflect.ValueOf(&geometriesSlice[0]).Elem()
		require.Equal(t, uint64(1000150006), geometry.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR
		require.True(t, geometry.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), geometry.FieldByName("geometryType").Uint())
		require.Equal(t, uint64(0), geometry.FieldByName("flags").Uint())

		geometryDataUnsafe = geometry.FieldByName("geometry").UnsafeAddr()
		geometryTriangleDataPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryTrianglesDataKHR)(unsafe.Pointer(geometryDataUnsafe))
		geometryData = reflect.ValueOf(geometryTriangleDataPtr).Elem()
		require.Equal(t, uint64(1000150005), geometryData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_TRIANGLES_DATA_KHR
		require.True(t, geometryData.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(62), geometryData.FieldByName("vertexFormat").Uint())
		require.Equal(t, uint64(31), geometryData.FieldByName("vertexStride").Uint())
		require.Equal(t, uint64(37), geometryData.FieldByName("maxVertex").Uint())
		require.Equal(t, uint64(1), geometryData.FieldByName("indexType").Uint())
		data = geometryData.FieldByName("vertexData").UnsafeAddr()
		require.Equal(t, uint64(29), *(*uint64)(unsafe.Pointer(data)))
		data = geometryData.FieldByName("indexData").UnsafeAddr()
		require.Equal(t, uint64(41), *(*uint64)(unsafe.Pointer(data)))
		data = geometryData.FieldByName("transformData").UnsafeAddr()
		require.Equal(t, uint64(43), *(*uint64)(unsafe.Pointer(data)))

		pBuildRangeInfosSlice := unsafe.Slice(ppBuildRangeInfos, 2)
		buildRangeInfosSlice := unsafe.Slice(pBuildRangeInfosSlice[0], 2)
		rangeInfo := reflect.ValueOf(&buildRangeInfosSlice[0]).Elem()
		require.Equal(t, uint64(1), rangeInfo.FieldByName("primitiveCount").Uint())
		require.Equal(t, uint64(3), rangeInfo.FieldByName("primitiveOffset").Uint())
		require.Equal(t, uint64(5), rangeInfo.FieldByName("firstVertex").Uint())
		require.Equal(t, uint64(7), rangeInfo.FieldByName("transformOffset").Uint())

		rangeInfo = reflect.ValueOf(&buildRangeInfosSlice[1]).Elem()
		require.Equal(t, uint64(11), rangeInfo.FieldByName("primitiveCount").Uint())
		require.Equal(t, uint64(13), rangeInfo.FieldByName("primitiveOffset").Uint())
		require.Equal(t, uint64(17), rangeInfo.FieldByName("firstVertex").Uint())
		require.Equal(t, uint64(19), rangeInfo.FieldByName("transformOffset").Uint())

		buildRangeInfosSlice = unsafe.Slice(pBuildRangeInfosSlice[1], 1)
		rangeInfo = reflect.ValueOf(&buildRangeInfosSlice[0]).Elem()
		require.Equal(t, uint64(23), rangeInfo.FieldByName("primitiveCount").Uint())
		require.Equal(t, uint64(29), rangeInfo.FieldByName("primitiveOffset").Uint())
		require.Equal(t, uint64(31), rangeInfo.FieldByName("firstVertex").Uint())
		require.Equal(t, uint64(37), rangeInfo.FieldByName("transformOffset").Uint())

		return core1_0.VKSuccess, nil
	})

	res, err := driver.BuildAccelerationStructure(deferredOperation, []khr_acceleration_structure.AccelerationStructureBuildGeometryInfo{
		{
			Type:                     khr_acceleration_structure.AccelerationStructureTypeBottomLevel,
			Flags:                    khr_acceleration_structure.BuildAccelerationStructureAllowCompaction,
			Mode:                     khr_acceleration_structure.BuildAccelerationStructureModeUpdate,
			SrcAccelerationStructure: srcStructure,
			DstAccelerationStructure: dstStructure1,
			Geometries: []khr_acceleration_structure.AccelerationStructureGeometry{
				{
					Type:  khr_acceleration_structure.GeometryTypeInstances,
					Flags: khr_acceleration_structure.GeometryOpaque,
					Geometry: khr_acceleration_structure.GeometryInstancesData{
						ArrayOfPointers: true,
						Data: khr_acceleration_structure.HostAddressConst{
							HostAddress: unsafe.Pointer(uintptr(17)),
						},
					},
				},
				{
					Type:  khr_acceleration_structure.GeometryTypeAABBs,
					Flags: khr_acceleration_structure.GeometryNoDuplicateAnyHitInvocation,
					Geometry: khr_acceleration_structure.GeometryAABBData{
						Data: khr_acceleration_structure.DeviceAddressConst{
							DeviceAddress: 19,
						},
						Stride: 24,
					},
				},
			},
			ScratchData: khr_acceleration_structure.DeviceAddressConst{
				DeviceAddress: 13,
			},
		},
		{
			Type:                     khr_acceleration_structure.AccelerationStructureTypeTopLevel,
			Flags:                    khr_acceleration_structure.BuildAccelerationStructureLowMemory,
			Mode:                     khr_acceleration_structure.BuildAccelerationStructureModeBuild,
			DstAccelerationStructure: dstStructure2,
			Geometries: []khr_acceleration_structure.AccelerationStructureGeometry{
				{
					Type:  khr_acceleration_structure.GeometryTypeTriangles,
					Flags: 0,
					Geometry: khr_acceleration_structure.GeometryTrianglesData{
						VertexFormat: core1_0.FormatA2R10G10B10UnsignedIntPacked,
						VertexData: khr_acceleration_structure.HostAddressConst{
							HostAddress: unsafe.Pointer(uintptr(29)),
						},
						VertexStride: 31,
						MaxVertex:    37,
						IndexType:    core1_0.IndexTypeUInt32,
						IndexData: khr_acceleration_structure.DeviceAddressConst{
							DeviceAddress: 41,
						},
						TransformData: khr_acceleration_structure.HostAddressConst{
							HostAddress: unsafe.Pointer(uintptr(43)),
						},
					},
				},
			},
			ScratchData: khr_acceleration_structure.HostAddressConst{
				HostAddress: unsafe.Pointer(uintptr(47)),
			},
		},
	}, [][]khr_acceleration_structure.AccelerationStructureBuildRangeInfo{
		{
			{
				PrimitiveCount:  1,
				PrimitiveOffset: 3,
				FirstVertex:     5,
				TransformOffset: 7,
			},
			{
				PrimitiveCount:  11,
				PrimitiveOffset: 13,
				FirstVertex:     17,
				TransformOffset: 19,
			},
		},
		{
			{
				PrimitiveCount:  23,
				PrimitiveOffset: 29,
				FirstVertex:     31,
				TransformOffset: 37,
			},
		},
	})

	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
}

func TestVulkanExtensionDriver_CmdBuildAccelerationStructures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	srcStructure := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	dstStructure1 := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	dstStructure2 := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkCmdBuildAccelerationStructuresKHR(
		commandBuffer.Handle(),
		loader.Uint32(2),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		commandBuffer loader.VkCommandBuffer,
		infoCount loader.Uint32,
		pInfos *khr_acceleration_structure_loader.VkAccelerationStructureBuildGeometryInfoKHR,
		ppBuildRangeInfos **khr_acceleration_structure_loader.VkAccelerationStructureBuildRangeInfoKHR,
	) (common.VkResult, error) {
		infoSlice := unsafe.Slice(pInfos, 2)

		info := reflect.ValueOf(&infoSlice[0]).Elem()
		require.Equal(t, uint64(1000150000), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_GEOMETRY_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), info.FieldByName("_type").Uint())
		require.Equal(t, uint64(2), info.FieldByName("flags").Uint())
		require.Equal(t, uint64(1), info.FieldByName("mode").Uint())
		require.Equal(t, srcStructure.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("srcAccelerationStructure").UnsafePointer()))
		require.Equal(t, dstStructure1.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("dstAccelerationStructure").UnsafePointer()))
		require.Equal(t, uint64(2), info.FieldByName("geometryCount").Uint())
		require.True(t, info.FieldByName("ppGeometries").IsNil())
		scratchData := info.FieldByName("scratchData").UnsafeAddr()
		require.Equal(t, uint64(13), *(*uint64)(unsafe.Pointer(scratchData)))

		geometriesPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryKHR)(info.FieldByName("pGeometries").UnsafePointer())
		geometriesSlice := unsafe.Slice(geometriesPtr, 2)

		geometry := reflect.ValueOf(&geometriesSlice[0]).Elem()
		require.Equal(t, uint64(1000150006), geometry.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR
		require.True(t, geometry.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), geometry.FieldByName("geometryType").Uint())
		require.Equal(t, uint64(1), geometry.FieldByName("flags").Uint())

		geometryDataUnsafe := geometry.FieldByName("geometry").UnsafeAddr()
		geometryInstanceDataPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryInstancesDataKHR)(unsafe.Pointer(geometryDataUnsafe))
		geometryData := reflect.ValueOf(geometryInstanceDataPtr).Elem()
		require.Equal(t, uint64(1000150004), geometryData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_INSTANCES_DATA_KHR
		require.True(t, geometryData.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), geometryData.FieldByName("arrayOfPointers").Uint())
		data := geometryData.FieldByName("data").UnsafeAddr()
		require.Equal(t, uint64(17), *(*uint64)(unsafe.Pointer(data)))

		geometry = reflect.ValueOf(&geometriesSlice[1]).Elem()
		require.Equal(t, uint64(1000150006), geometry.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR
		require.True(t, geometry.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), geometry.FieldByName("geometryType").Uint())
		require.Equal(t, uint64(2), geometry.FieldByName("flags").Uint())

		geometryDataUnsafe = geometry.FieldByName("geometry").UnsafeAddr()
		geometryAABBDataPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryAabbsDataKHR)(unsafe.Pointer(geometryDataUnsafe))
		geometryData = reflect.ValueOf(geometryAABBDataPtr).Elem()
		require.Equal(t, uint64(1000150003), geometryData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_AABBS_DATA_KHR
		require.True(t, geometryData.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(24), geometryData.FieldByName("stride").Uint())
		data = geometryData.FieldByName("data").UnsafeAddr()
		require.Equal(t, uint64(19), *(*uint64)(unsafe.Pointer(data)))

		info = reflect.ValueOf(&infoSlice[1]).Elem()
		require.Equal(t, uint64(1000150000), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_GEOMETRY_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), info.FieldByName("_type").Uint())
		require.Equal(t, uint64(16), info.FieldByName("flags").Uint())
		require.Equal(t, uint64(0), info.FieldByName("mode").Uint())
		require.Equal(t, khr_acceleration_structure_loader.VkAccelerationStructureKHR(0), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("srcAccelerationStructure").UnsafePointer()))
		require.Equal(t, dstStructure2.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("dstAccelerationStructure").UnsafePointer()))
		require.Equal(t, uint64(1), info.FieldByName("geometryCount").Uint())
		require.True(t, info.FieldByName("ppGeometries").IsNil())
		scratchData = info.FieldByName("scratchData").UnsafeAddr()
		require.Equal(t, uint64(47), *(*uint64)(unsafe.Pointer(scratchData)))

		geometriesPtr = (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryKHR)(info.FieldByName("pGeometries").UnsafePointer())
		geometriesSlice = unsafe.Slice(geometriesPtr, 1)

		geometry = reflect.ValueOf(&geometriesSlice[0]).Elem()
		require.Equal(t, uint64(1000150006), geometry.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR
		require.True(t, geometry.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), geometry.FieldByName("geometryType").Uint())
		require.Equal(t, uint64(0), geometry.FieldByName("flags").Uint())

		geometryDataUnsafe = geometry.FieldByName("geometry").UnsafeAddr()
		geometryTriangleDataPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryTrianglesDataKHR)(unsafe.Pointer(geometryDataUnsafe))
		geometryData = reflect.ValueOf(geometryTriangleDataPtr).Elem()
		require.Equal(t, uint64(1000150005), geometryData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_TRIANGLES_DATA_KHR
		require.True(t, geometryData.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(62), geometryData.FieldByName("vertexFormat").Uint())
		require.Equal(t, uint64(31), geometryData.FieldByName("vertexStride").Uint())
		require.Equal(t, uint64(37), geometryData.FieldByName("maxVertex").Uint())
		require.Equal(t, uint64(1), geometryData.FieldByName("indexType").Uint())
		data = geometryData.FieldByName("vertexData").UnsafeAddr()
		require.Equal(t, uint64(29), *(*uint64)(unsafe.Pointer(data)))
		data = geometryData.FieldByName("indexData").UnsafeAddr()
		require.Equal(t, uint64(41), *(*uint64)(unsafe.Pointer(data)))
		data = geometryData.FieldByName("transformData").UnsafeAddr()
		require.Equal(t, uint64(43), *(*uint64)(unsafe.Pointer(data)))

		pBuildRangeInfosSlice := unsafe.Slice(ppBuildRangeInfos, 2)
		buildRangeInfosSlice := unsafe.Slice(pBuildRangeInfosSlice[0], 2)
		rangeInfo := reflect.ValueOf(&buildRangeInfosSlice[0]).Elem()
		require.Equal(t, uint64(1), rangeInfo.FieldByName("primitiveCount").Uint())
		require.Equal(t, uint64(3), rangeInfo.FieldByName("primitiveOffset").Uint())
		require.Equal(t, uint64(5), rangeInfo.FieldByName("firstVertex").Uint())
		require.Equal(t, uint64(7), rangeInfo.FieldByName("transformOffset").Uint())

		rangeInfo = reflect.ValueOf(&buildRangeInfosSlice[1]).Elem()
		require.Equal(t, uint64(11), rangeInfo.FieldByName("primitiveCount").Uint())
		require.Equal(t, uint64(13), rangeInfo.FieldByName("primitiveOffset").Uint())
		require.Equal(t, uint64(17), rangeInfo.FieldByName("firstVertex").Uint())
		require.Equal(t, uint64(19), rangeInfo.FieldByName("transformOffset").Uint())

		buildRangeInfosSlice = unsafe.Slice(pBuildRangeInfosSlice[1], 1)
		rangeInfo = reflect.ValueOf(&buildRangeInfosSlice[0]).Elem()
		require.Equal(t, uint64(23), rangeInfo.FieldByName("primitiveCount").Uint())
		require.Equal(t, uint64(29), rangeInfo.FieldByName("primitiveOffset").Uint())
		require.Equal(t, uint64(31), rangeInfo.FieldByName("firstVertex").Uint())
		require.Equal(t, uint64(37), rangeInfo.FieldByName("transformOffset").Uint())

		return core1_0.VKSuccess, nil
	})

	driver.CmdBuildAccelerationStructures(commandBuffer, []khr_acceleration_structure.AccelerationStructureBuildGeometryInfo{
		{
			Type:                     khr_acceleration_structure.AccelerationStructureTypeBottomLevel,
			Flags:                    khr_acceleration_structure.BuildAccelerationStructureAllowCompaction,
			Mode:                     khr_acceleration_structure.BuildAccelerationStructureModeUpdate,
			SrcAccelerationStructure: srcStructure,
			DstAccelerationStructure: dstStructure1,
			Geometries: []khr_acceleration_structure.AccelerationStructureGeometry{
				{
					Type:  khr_acceleration_structure.GeometryTypeInstances,
					Flags: khr_acceleration_structure.GeometryOpaque,
					Geometry: khr_acceleration_structure.GeometryInstancesData{
						ArrayOfPointers: true,
						Data: khr_acceleration_structure.HostAddressConst{
							HostAddress: unsafe.Pointer(uintptr(17)),
						},
					},
				},
				{
					Type:  khr_acceleration_structure.GeometryTypeAABBs,
					Flags: khr_acceleration_structure.GeometryNoDuplicateAnyHitInvocation,
					Geometry: khr_acceleration_structure.GeometryAABBData{
						Data: khr_acceleration_structure.DeviceAddressConst{
							DeviceAddress: 19,
						},
						Stride: 24,
					},
				},
			},
			ScratchData: khr_acceleration_structure.DeviceAddressConst{
				DeviceAddress: 13,
			},
		},
		{
			Type:                     khr_acceleration_structure.AccelerationStructureTypeTopLevel,
			Flags:                    khr_acceleration_structure.BuildAccelerationStructureLowMemory,
			Mode:                     khr_acceleration_structure.BuildAccelerationStructureModeBuild,
			DstAccelerationStructure: dstStructure2,
			Geometries: []khr_acceleration_structure.AccelerationStructureGeometry{
				{
					Type:  khr_acceleration_structure.GeometryTypeTriangles,
					Flags: 0,
					Geometry: khr_acceleration_structure.GeometryTrianglesData{
						VertexFormat: core1_0.FormatA2R10G10B10UnsignedIntPacked,
						VertexData: khr_acceleration_structure.HostAddressConst{
							HostAddress: unsafe.Pointer(uintptr(29)),
						},
						VertexStride: 31,
						MaxVertex:    37,
						IndexType:    core1_0.IndexTypeUInt32,
						IndexData: khr_acceleration_structure.DeviceAddressConst{
							DeviceAddress: 41,
						},
						TransformData: khr_acceleration_structure.HostAddressConst{
							HostAddress: unsafe.Pointer(uintptr(43)),
						},
					},
				},
			},
			ScratchData: khr_acceleration_structure.HostAddressConst{
				HostAddress: unsafe.Pointer(uintptr(47)),
			},
		},
	}, [][]khr_acceleration_structure.AccelerationStructureBuildRangeInfo{
		{
			{
				PrimitiveCount:  1,
				PrimitiveOffset: 3,
				FirstVertex:     5,
				TransformOffset: 7,
			},
			{
				PrimitiveCount:  11,
				PrimitiveOffset: 13,
				FirstVertex:     17,
				TransformOffset: 19,
			},
		},
		{
			{
				PrimitiveCount:  23,
				PrimitiveOffset: 29,
				FirstVertex:     31,
				TransformOffset: 37,
			},
		},
	})
}

func TestVulkanExtensionDriver_CmdBuildAccelerationStructuresIndirect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	srcStructure := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	dstStructure1 := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	dstStructure2 := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkCmdBuildAccelerationStructuresIndirectKHR(
		commandBuffer.Handle(),
		loader.Uint32(2),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		commandBuffer loader.VkCommandBuffer,
		infoCount loader.Uint32,
		pInfos *khr_acceleration_structure_loader.VkAccelerationStructureBuildGeometryInfoKHR,
		pIndirectDeviceAddresses *loader.VkDeviceAddress,
		pIndirectStrides *loader.Uint32,
		ppMaxPrimitiveCounts **loader.Uint32,
	) (common.VkResult, error) {
		infoSlice := unsafe.Slice(pInfos, 2)

		info := reflect.ValueOf(&infoSlice[0]).Elem()
		require.Equal(t, uint64(1000150000), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_GEOMETRY_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), info.FieldByName("_type").Uint())
		require.Equal(t, uint64(2), info.FieldByName("flags").Uint())
		require.Equal(t, uint64(1), info.FieldByName("mode").Uint())
		require.Equal(t, srcStructure.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("srcAccelerationStructure").UnsafePointer()))
		require.Equal(t, dstStructure1.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("dstAccelerationStructure").UnsafePointer()))
		require.Equal(t, uint64(2), info.FieldByName("geometryCount").Uint())
		require.True(t, info.FieldByName("ppGeometries").IsNil())
		scratchData := info.FieldByName("scratchData").UnsafeAddr()
		require.Equal(t, uint64(13), *(*uint64)(unsafe.Pointer(scratchData)))

		geometriesPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryKHR)(info.FieldByName("pGeometries").UnsafePointer())
		geometriesSlice := unsafe.Slice(geometriesPtr, 2)

		geometry := reflect.ValueOf(&geometriesSlice[0]).Elem()
		require.Equal(t, uint64(1000150006), geometry.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR
		require.True(t, geometry.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), geometry.FieldByName("geometryType").Uint())
		require.Equal(t, uint64(1), geometry.FieldByName("flags").Uint())

		geometryDataUnsafe := geometry.FieldByName("geometry").UnsafeAddr()
		geometryInstanceDataPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryInstancesDataKHR)(unsafe.Pointer(geometryDataUnsafe))
		geometryData := reflect.ValueOf(geometryInstanceDataPtr).Elem()
		require.Equal(t, uint64(1000150004), geometryData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_INSTANCES_DATA_KHR
		require.True(t, geometryData.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), geometryData.FieldByName("arrayOfPointers").Uint())
		data := geometryData.FieldByName("data").UnsafeAddr()
		require.Equal(t, uint64(17), *(*uint64)(unsafe.Pointer(data)))

		geometry = reflect.ValueOf(&geometriesSlice[1]).Elem()
		require.Equal(t, uint64(1000150006), geometry.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR
		require.True(t, geometry.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), geometry.FieldByName("geometryType").Uint())
		require.Equal(t, uint64(2), geometry.FieldByName("flags").Uint())

		geometryDataUnsafe = geometry.FieldByName("geometry").UnsafeAddr()
		geometryAABBDataPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryAabbsDataKHR)(unsafe.Pointer(geometryDataUnsafe))
		geometryData = reflect.ValueOf(geometryAABBDataPtr).Elem()
		require.Equal(t, uint64(1000150003), geometryData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_AABBS_DATA_KHR
		require.True(t, geometryData.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(24), geometryData.FieldByName("stride").Uint())
		data = geometryData.FieldByName("data").UnsafeAddr()
		require.Equal(t, uint64(19), *(*uint64)(unsafe.Pointer(data)))

		info = reflect.ValueOf(&infoSlice[1]).Elem()
		require.Equal(t, uint64(1000150000), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_GEOMETRY_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), info.FieldByName("_type").Uint())
		require.Equal(t, uint64(16), info.FieldByName("flags").Uint())
		require.Equal(t, uint64(0), info.FieldByName("mode").Uint())
		require.Equal(t, khr_acceleration_structure_loader.VkAccelerationStructureKHR(0), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("srcAccelerationStructure").UnsafePointer()))
		require.Equal(t, dstStructure2.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("dstAccelerationStructure").UnsafePointer()))
		require.Equal(t, uint64(1), info.FieldByName("geometryCount").Uint())
		require.True(t, info.FieldByName("ppGeometries").IsNil())
		scratchData = info.FieldByName("scratchData").UnsafeAddr()
		require.Equal(t, uint64(47), *(*uint64)(unsafe.Pointer(scratchData)))

		geometriesPtr = (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryKHR)(info.FieldByName("pGeometries").UnsafePointer())
		geometriesSlice = unsafe.Slice(geometriesPtr, 1)

		geometry = reflect.ValueOf(&geometriesSlice[0]).Elem()
		require.Equal(t, uint64(1000150006), geometry.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR
		require.True(t, geometry.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), geometry.FieldByName("geometryType").Uint())
		require.Equal(t, uint64(0), geometry.FieldByName("flags").Uint())

		geometryDataUnsafe = geometry.FieldByName("geometry").UnsafeAddr()
		geometryTriangleDataPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryTrianglesDataKHR)(unsafe.Pointer(geometryDataUnsafe))
		geometryData = reflect.ValueOf(geometryTriangleDataPtr).Elem()
		require.Equal(t, uint64(1000150005), geometryData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_TRIANGLES_DATA_KHR
		require.True(t, geometryData.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(62), geometryData.FieldByName("vertexFormat").Uint())
		require.Equal(t, uint64(31), geometryData.FieldByName("vertexStride").Uint())
		require.Equal(t, uint64(37), geometryData.FieldByName("maxVertex").Uint())
		require.Equal(t, uint64(1), geometryData.FieldByName("indexType").Uint())
		data = geometryData.FieldByName("vertexData").UnsafeAddr()
		require.Equal(t, uint64(29), *(*uint64)(unsafe.Pointer(data)))
		data = geometryData.FieldByName("indexData").UnsafeAddr()
		require.Equal(t, uint64(41), *(*uint64)(unsafe.Pointer(data)))
		data = geometryData.FieldByName("transformData").UnsafeAddr()
		require.Equal(t, uint64(43), *(*uint64)(unsafe.Pointer(data)))

		addressSlice := unsafe.Slice(pIndirectDeviceAddresses, 2)
		require.Equal(t, loader.VkDeviceAddress(1), addressSlice[0])
		require.Equal(t, loader.VkDeviceAddress(3), addressSlice[1])

		strideSlice := unsafe.Slice(pIndirectStrides, 2)
		require.Equal(t, loader.Uint32(5), strideSlice[0])
		require.Equal(t, loader.Uint32(7), strideSlice[1])

		pMaxPrimitiveCounts := unsafe.Slice(ppMaxPrimitiveCounts, 2)
		maxPrimitivesSlice := unsafe.Slice(pMaxPrimitiveCounts[0], 2)
		require.Equal(t, loader.Uint32(11), maxPrimitivesSlice[0])
		require.Equal(t, loader.Uint32(13), maxPrimitivesSlice[1])
		maxPrimitivesSlice = unsafe.Slice(pMaxPrimitiveCounts[1], 1)
		require.Equal(t, loader.Uint32(17), maxPrimitivesSlice[0])

		return core1_0.VKSuccess, nil
	})

	driver.CmdBuildAccelerationStructuresIndirect(
		commandBuffer,
		[]khr_acceleration_structure.AccelerationStructureBuildGeometryInfo{
			{
				Type:                     khr_acceleration_structure.AccelerationStructureTypeBottomLevel,
				Flags:                    khr_acceleration_structure.BuildAccelerationStructureAllowCompaction,
				Mode:                     khr_acceleration_structure.BuildAccelerationStructureModeUpdate,
				SrcAccelerationStructure: srcStructure,
				DstAccelerationStructure: dstStructure1,
				Geometries: []khr_acceleration_structure.AccelerationStructureGeometry{
					{
						Type:  khr_acceleration_structure.GeometryTypeInstances,
						Flags: khr_acceleration_structure.GeometryOpaque,
						Geometry: khr_acceleration_structure.GeometryInstancesData{
							ArrayOfPointers: true,
							Data: khr_acceleration_structure.HostAddressConst{
								HostAddress: unsafe.Pointer(uintptr(17)),
							},
						},
					},
					{
						Type:  khr_acceleration_structure.GeometryTypeAABBs,
						Flags: khr_acceleration_structure.GeometryNoDuplicateAnyHitInvocation,
						Geometry: khr_acceleration_structure.GeometryAABBData{
							Data: khr_acceleration_structure.DeviceAddressConst{
								DeviceAddress: 19,
							},
							Stride: 24,
						},
					},
				},
				ScratchData: khr_acceleration_structure.DeviceAddressConst{
					DeviceAddress: 13,
				},
			},
			{
				Type:                     khr_acceleration_structure.AccelerationStructureTypeTopLevel,
				Flags:                    khr_acceleration_structure.BuildAccelerationStructureLowMemory,
				Mode:                     khr_acceleration_structure.BuildAccelerationStructureModeBuild,
				DstAccelerationStructure: dstStructure2,
				Geometries: []khr_acceleration_structure.AccelerationStructureGeometry{
					{
						Type:  khr_acceleration_structure.GeometryTypeTriangles,
						Flags: 0,
						Geometry: khr_acceleration_structure.GeometryTrianglesData{
							VertexFormat: core1_0.FormatA2R10G10B10UnsignedIntPacked,
							VertexData: khr_acceleration_structure.HostAddressConst{
								HostAddress: unsafe.Pointer(uintptr(29)),
							},
							VertexStride: 31,
							MaxVertex:    37,
							IndexType:    core1_0.IndexTypeUInt32,
							IndexData: khr_acceleration_structure.DeviceAddressConst{
								DeviceAddress: 41,
							},
							TransformData: khr_acceleration_structure.HostAddressConst{
								HostAddress: unsafe.Pointer(uintptr(43)),
							},
						},
					},
				},
				ScratchData: khr_acceleration_structure.HostAddressConst{
					HostAddress: unsafe.Pointer(uintptr(47)),
				},
			},
		},
		[]uint64{1, 3},
		[]int{5, 7},
		[][]int{{11, 13}, {17}},
	)
}

func TestVulkanExtensionDriver_CmdCopyAccelerationStructure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	srcStruct := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	dstStruct := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkCmdCopyAccelerationStructureKHR(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		commandBuffer loader.VkCommandBuffer,
		infoPtr *khr_acceleration_structure_loader.VkCopyAccelerationStructureInfoKHR,
	) {
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, uint64(1000150010), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_ACCELERATION_STRUCTURE_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, srcStruct.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("src").UnsafePointer()))
		require.Equal(t, dstStruct.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("dst").UnsafePointer()))
		require.Equal(t, uint64(3), info.FieldByName("mode").Uint())
	})

	driver.CmdCopyAccelerationStructure(commandBuffer, khr_acceleration_structure.CopyAccelerationStructureInfo{
		Src:  srcStruct,
		Dst:  dstStruct,
		Mode: khr_acceleration_structure.CopyAccelerationStructureModeDeserialize,
	})
}

func TestVulkanExtensionDriver_CmdCopyAccelerationStructureToMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	srcStruct := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkCmdCopyAccelerationStructureToMemoryKHR(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		commandBuffer loader.VkCommandBuffer,
		infoPtr *khr_acceleration_structure_loader.VkCopyAccelerationStructureToMemoryInfoKHR,
	) {
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, uint64(1000150011), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_ACCELERATION_STRUCTURE_TO_MEMORY_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, srcStruct.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("src").UnsafePointer()))
		require.Equal(t, uint64(0), info.FieldByName("mode").Uint())

		data := info.FieldByName("dst").UnsafeAddr()
		require.Equal(t, uint64(7), *(*uint64)(unsafe.Pointer(data)))
	})

	driver.CmdCopyAccelerationStructureToMemory(commandBuffer, khr_acceleration_structure.CopyAccelerationStructureToMemoryInfo{
		Src: srcStruct,
		Dst: khr_acceleration_structure.DeviceAddressConst{
			DeviceAddress: uint64(7),
		},
		Mode: khr_acceleration_structure.CopyAccelerationStructureModeClone,
	})
}

func TestVulkanExtensionDriver_CmdCopyMemoryToAccelerationStructure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	dstStruct := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkCmdCopyMemoryToAccelerationStructureKHR(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		commandBuffer loader.VkCommandBuffer,
		infoPtr *khr_acceleration_structure_loader.VkCopyMemoryToAccelerationStructureInfoKHR,
	) {
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, uint64(1000150012), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_MEMORY_TO_ACCELERATION_STRUCTURE_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, dstStruct.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("dst").UnsafePointer()))
		require.Equal(t, uint64(1), info.FieldByName("mode").Uint())

		data := info.FieldByName("src").UnsafeAddr()
		require.Equal(t, uint64(11), *(*uint64)(unsafe.Pointer(data)))
	})

	driver.CmdCopyMemoryToAccelerationStructure(commandBuffer, khr_acceleration_structure.CopyMemoryToAccelerationStructureInfo{
		Src: khr_acceleration_structure.HostAddressConst{
			HostAddress: unsafe.Pointer(uintptr(11)),
		},
		Dst:  dstStruct,
		Mode: khr_acceleration_structure.CopyAccelerationStructureModeCompact,
	})
}

func TestVulkanExtensionDriver_CmdWriteAccelerationStructuresProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	accel1 := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	accel2 := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	accel3 := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	queryPool := mocks.NewDummyQueryPool(device)

	mockLoader.EXPECT().VkCmdWriteAccelerationStructuresPropertiesKHR(
		commandBuffer.Handle(),
		loader.Uint32(3),
		gomock.Not(gomock.Nil()),
		loader.VkQueryType(0),
		queryPool.Handle(),
		loader.Uint32(5),
	).DoAndReturn(func(
		commandBuffer loader.VkCommandBuffer,
		accelerationStructureCount loader.Uint32,
		pAccelerationStructures *khr_acceleration_structure_loader.VkAccelerationStructureKHR,
		queryType loader.VkQueryType,
		queryPool loader.VkQueryPool,
		firstQuery loader.Uint32,
	) {
		accelStructureSlice := unsafe.Slice(pAccelerationStructures, 3)
		require.Equal(t, []khr_acceleration_structure_loader.VkAccelerationStructureKHR{
			accel1.Handle(), accel2.Handle(), accel3.Handle(),
		}, accelStructureSlice)
	})

	driver.CmdWriteAccelerationStructuresProperties(
		commandBuffer,
		[]khr_acceleration_structure.AccelerationStructure{accel1, accel2, accel3},
		core1_0.QueryTypeOcclusion,
		queryPool,
		5,
	)
}

func TestVulkanExtensionDriver_CopyAccelerationStructure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	deferredOperation := mock_deferred_host_operations.NewDummyDeferredOperation(device)
	srcAccel := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	dstAccel := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkCopyAccelerationStructureKHR(
		device.Handle(),
		deferredOperation.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		deferredOperation khr_deferred_host_operations_loader.VkDeferredOperationKHR,
		pInfo *khr_acceleration_structure_loader.VkCopyAccelerationStructureInfoKHR,
	) (common.VkResult, error) {
		info := reflect.ValueOf(pInfo).Elem()
		require.Equal(t, uint64(1000150010), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_ACCELERATION_STRUCTURE_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, srcAccel.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("src").UnsafePointer()))
		require.Equal(t, dstAccel.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("dst").UnsafePointer()))
		require.Equal(t, uint64(2), info.FieldByName("mode").Uint())

		return core1_0.VKSuccess, nil
	})

	res, err := driver.CopyAccelerationStructure(
		deferredOperation,
		khr_acceleration_structure.CopyAccelerationStructureInfo{
			Src:  srcAccel,
			Dst:  dstAccel,
			Mode: khr_acceleration_structure.CopyAccelerationStructureModeSerialize,
		})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
}

func TestVulkanExtensionDriver_CopyAccelerationStructureToMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	deferredOperation := mock_deferred_host_operations.NewDummyDeferredOperation(device)
	srcStruct := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkCopyAccelerationStructureToMemoryKHR(
		device.Handle(),
		deferredOperation.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		deferredOperation khr_deferred_host_operations_loader.VkDeferredOperationKHR,
		infoPtr *khr_acceleration_structure_loader.VkCopyAccelerationStructureToMemoryInfoKHR,
	) (common.VkResult, error) {
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, uint64(1000150011), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_ACCELERATION_STRUCTURE_TO_MEMORY_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, srcStruct.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("src").UnsafePointer()))
		require.Equal(t, uint64(0), info.FieldByName("mode").Uint())

		data := info.FieldByName("dst").UnsafeAddr()
		require.Equal(t, uint64(7), *(*uint64)(unsafe.Pointer(data)))

		return core1_0.VKSuccess, nil
	})

	res, err := driver.CopyAccelerationStructureToMemory(deferredOperation, khr_acceleration_structure.CopyAccelerationStructureToMemoryInfo{
		Src: srcStruct,
		Dst: khr_acceleration_structure.DeviceAddressConst{
			DeviceAddress: uint64(7),
		},
		Mode: khr_acceleration_structure.CopyAccelerationStructureModeClone,
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
}

func TestVulkanExtensionDriver_CopyMemoryToAccelerationStructure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	deferredOperation := mock_deferred_host_operations.NewDummyDeferredOperation(device)
	dstStruct := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkCopyMemoryToAccelerationStructureKHR(
		device.Handle(),
		deferredOperation.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		deferredOperation khr_deferred_host_operations_loader.VkDeferredOperationKHR,
		infoPtr *khr_acceleration_structure_loader.VkCopyMemoryToAccelerationStructureInfoKHR,
	) (common.VkResult, error) {
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, uint64(1000150012), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_MEMORY_TO_ACCELERATION_STRUCTURE_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, dstStruct.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("dst").UnsafePointer()))
		require.Equal(t, uint64(1), info.FieldByName("mode").Uint())

		data := info.FieldByName("src").UnsafeAddr()
		require.Equal(t, uint64(11), *(*uint64)(unsafe.Pointer(data)))

		return core1_0.VKSuccess, nil
	})

	res, err := driver.CopyMemoryToAccelerationStructure(deferredOperation, khr_acceleration_structure.CopyMemoryToAccelerationStructureInfo{
		Src: khr_acceleration_structure.HostAddressConst{
			HostAddress: unsafe.Pointer(uintptr(11)),
		},
		Dst:  dstStruct,
		Mode: khr_acceleration_structure.CopyAccelerationStructureModeCompact,
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
}

func TestVulkanExtensionDriver_CreateAccelerationStructure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	expectedAccelStruct := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	buffer := mocks.NewDummyBuffer(device)

	mockLoader.EXPECT().VkCreateAccelerationStructureKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		pCreateInfo *khr_acceleration_structure_loader.VkAccelerationStructureCreateInfoKHR,
		pAllocator *loader.VkAllocationCallbacks,
		pAccelerationStructure *khr_acceleration_structure_loader.VkAccelerationStructureKHR,
	) (common.VkResult, error) {
		info := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(1000150017), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), info.FieldByName("createFlags").Uint())
		require.Equal(t, buffer.Handle(), loader.VkBuffer(info.FieldByName("buffer").UnsafePointer()))
		require.Equal(t, uint64(1), info.FieldByName("offset").Uint())
		require.Equal(t, uint64(3), info.FieldByName("size").Uint())
		require.Equal(t, uint64(1), info.FieldByName("_type").Uint())
		require.Equal(t, uint64(5), info.FieldByName("deviceAddress").Uint())

		*pAccelerationStructure = khr_acceleration_structure_loader.VkAccelerationStructureKHR(expectedAccelStruct.Handle())

		return core1_0.VKSuccess, nil
	})

	accelStructure, res, err := driver.CreateAccelerationStructure(khr_acceleration_structure.AccelerationStructureCreateInfo{
		CreateFlags:   khr_acceleration_structure.AccelerationStructureCreateDeviceAddressCaptureReplay,
		Buffer:        buffer,
		Offset:        1,
		Size:          3,
		Type:          khr_acceleration_structure.AccelerationStructureTypeBottomLevel,
		DeviceAddress: 5,
	}, nil)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.Equal(t, expectedAccelStruct.Handle(), accelStructure.Handle())
}

func TestVulkanExtensionDriver_DestroyAccelerationStructure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	accelStructure := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkDestroyAccelerationStructureKHR(
		device.Handle(),
		accelStructure.Handle(),
		gomock.Nil(),
	)

	driver.DestroyAccelerationStructure(accelStructure, nil)
}

func TestVulkanExtensionDriver_GetAccelerationStructureBuildSizes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	srcStructure := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	dstStructure := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkGetAccelerationStructureBuildSizesKHR(
		device.Handle(),
		khr_acceleration_structure_loader.VkAccelerationStructureBuildTypeKHR(2),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		buildType khr_acceleration_structure_loader.VkAccelerationStructureBuildTypeKHR,
		pBuildInfo *khr_acceleration_structure_loader.VkAccelerationStructureBuildGeometryInfoKHR,
		pMaxPrimitiveCounts *loader.Uint32,
		pSizeInfo *khr_acceleration_structure_loader.VkAccelerationStructureBuildSizesInfoKHR,
	) {
		info := reflect.ValueOf(pBuildInfo).Elem()
		require.Equal(t, uint64(1000150000), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_GEOMETRY_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), info.FieldByName("_type").Uint())
		require.Equal(t, uint64(2), info.FieldByName("flags").Uint())
		require.Equal(t, uint64(1), info.FieldByName("mode").Uint())
		require.Equal(t, srcStructure.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("srcAccelerationStructure").UnsafePointer()))
		require.Equal(t, dstStructure.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("dstAccelerationStructure").UnsafePointer()))
		require.Equal(t, uint64(2), info.FieldByName("geometryCount").Uint())
		require.True(t, info.FieldByName("ppGeometries").IsNil())
		scratchData := info.FieldByName("scratchData").UnsafeAddr()
		require.Equal(t, uint64(13), *(*uint64)(unsafe.Pointer(scratchData)))

		geometriesPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryKHR)(info.FieldByName("pGeometries").UnsafePointer())
		geometriesSlice := unsafe.Slice(geometriesPtr, 2)

		geometry := reflect.ValueOf(&geometriesSlice[0]).Elem()
		require.Equal(t, uint64(1000150006), geometry.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR
		require.True(t, geometry.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), geometry.FieldByName("geometryType").Uint())
		require.Equal(t, uint64(1), geometry.FieldByName("flags").Uint())

		geometryDataUnsafe := geometry.FieldByName("geometry").UnsafeAddr()
		geometryInstanceDataPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryInstancesDataKHR)(unsafe.Pointer(geometryDataUnsafe))
		geometryData := reflect.ValueOf(geometryInstanceDataPtr).Elem()
		require.Equal(t, uint64(1000150004), geometryData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_INSTANCES_DATA_KHR
		require.True(t, geometryData.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), geometryData.FieldByName("arrayOfPointers").Uint())
		data := geometryData.FieldByName("data").UnsafeAddr()
		require.Equal(t, uint64(17), *(*uint64)(unsafe.Pointer(data)))

		geometry = reflect.ValueOf(&geometriesSlice[1]).Elem()
		require.Equal(t, uint64(1000150006), geometry.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR
		require.True(t, geometry.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), geometry.FieldByName("geometryType").Uint())
		require.Equal(t, uint64(2), geometry.FieldByName("flags").Uint())

		geometryDataUnsafe = geometry.FieldByName("geometry").UnsafeAddr()
		geometryAABBDataPtr := (*khr_acceleration_structure_loader.VkAccelerationStructureGeometryAabbsDataKHR)(unsafe.Pointer(geometryDataUnsafe))
		geometryData = reflect.ValueOf(geometryAABBDataPtr).Elem()
		require.Equal(t, uint64(1000150003), geometryData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_AABBS_DATA_KHR
		require.True(t, geometryData.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(24), geometryData.FieldByName("stride").Uint())
		data = geometryData.FieldByName("data").UnsafeAddr()
		require.Equal(t, uint64(19), *(*uint64)(unsafe.Pointer(data)))

		maxPrimitiveSlice := unsafe.Slice(pMaxPrimitiveCounts, 2)
		require.Equal(t, []loader.Uint32{3, 5}, maxPrimitiveSlice)

		size := reflect.ValueOf(pSizeInfo).Elem()
		require.Equal(t, uint64(1000150020), size.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_SIZES_INFO_KHR
		require.True(t, size.FieldByName("pNext").IsNil())
		*(*uint64)(unsafe.Pointer(size.FieldByName("accelerationStructureSize").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(size.FieldByName("updateScratchSize").UnsafeAddr())) = uint64(3)
		*(*uint64)(unsafe.Pointer(size.FieldByName("buildScratchSize").UnsafeAddr())) = uint64(5)
	})

	var buildSizes khr_acceleration_structure.AccelerationStructureBuildSizesInfo
	driver.GetAccelerationStructureBuildSizes(
		khr_acceleration_structure.AccelerationStructureBuildTypeHostOrDevice,
		khr_acceleration_structure.AccelerationStructureBuildGeometryInfo{
			Type:                     khr_acceleration_structure.AccelerationStructureTypeBottomLevel,
			Flags:                    khr_acceleration_structure.BuildAccelerationStructureAllowCompaction,
			Mode:                     khr_acceleration_structure.BuildAccelerationStructureModeUpdate,
			SrcAccelerationStructure: srcStructure,
			DstAccelerationStructure: dstStructure,
			Geometries: []khr_acceleration_structure.AccelerationStructureGeometry{
				{
					Type:  khr_acceleration_structure.GeometryTypeInstances,
					Flags: khr_acceleration_structure.GeometryOpaque,
					Geometry: khr_acceleration_structure.GeometryInstancesData{
						ArrayOfPointers: true,
						Data: khr_acceleration_structure.HostAddressConst{
							HostAddress: unsafe.Pointer(uintptr(17)),
						},
					},
				},
				{
					Type:  khr_acceleration_structure.GeometryTypeAABBs,
					Flags: khr_acceleration_structure.GeometryNoDuplicateAnyHitInvocation,
					Geometry: khr_acceleration_structure.GeometryAABBData{
						Data: khr_acceleration_structure.DeviceAddressConst{
							DeviceAddress: 19,
						},
						Stride: 24,
					},
				},
			},
			ScratchData: khr_acceleration_structure.DeviceAddressConst{
				DeviceAddress: 13,
			},
		}, []int{3, 5}, &buildSizes)
	require.Equal(t, khr_acceleration_structure.AccelerationStructureBuildSizesInfo{
		AccelerationStructureSize: 1,
		UpdateScratchSize:         3,
		BuildScratchSize:          5,
	}, buildSizes)
}

func TestVulkanExtensionDriver_GetAccelerationStructureDeviceAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	accel := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkGetAccelerationStructureDeviceAddressKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		pInfo *khr_acceleration_structure_loader.VkAccelerationStructureDeviceAddressInfoKHR,
	) loader.VkDeviceAddress {
		info := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000150002), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_DEVICE_ADDRESS_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, accel.Handle(), khr_acceleration_structure_loader.VkAccelerationStructureKHR(info.FieldByName("accelerationStructure").UnsafePointer()))

		return loader.VkDeviceAddress(31)
	})

	deviceAddress := driver.GetAccelerationStructureDeviceAddress(khr_acceleration_structure.AccelerationStructureDeviceAddressInfo{
		AccelerationStructure: accel,
	})
	require.Equal(t, uint64(31), deviceAddress)
}

func TestVulkanExtensionDriver_GetDeviceAccelerationStructureCompatibility(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	mockLoader.EXPECT().VkGetDeviceAccelerationStructureCompatibilityKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		pVersionInfo *khr_acceleration_structure_loader.VkAccelerationStructureVersionInfoKHR,
		pCompatibility *khr_acceleration_structure_loader.VkAccelerationStructureCompatibilityKHR,
	) {
		info := reflect.ValueOf(pVersionInfo).Elem()
		require.Equal(t, uint64(1000150009), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_VERSION_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		data := *(*uint64)(unsafe.Pointer(info.FieldByName("pVersionData").UnsafeAddr()))
		require.Equal(t, uint64(23), data)

		*pCompatibility = khr_acceleration_structure_loader.VkAccelerationStructureCompatibilityKHR(1)
	})

	compat := driver.GetDeviceAccelerationStructureCompatibility(khr_acceleration_structure.AccelerationStructureVersionInfo{
		VersionData: unsafe.Pointer(uintptr(23)),
	})
	require.Equal(t, khr_acceleration_structure.AccelerationStructureCompatibilityIncompatible, compat)
}

func TestVulkanExtensionDriver_WriteAccelerationStructuresProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_acceleration_structure.NewMockLoader(ctrl)
	driver := khr_acceleration_structure.CreateExtensionDriverFromLoader(mockLoader, device)

	accel1 := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	accel2 := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	mockLoader.EXPECT().VkWriteAccelerationStructuresPropertiesKHR(
		device.Handle(),
		loader.Uint32(2),
		gomock.Not(gomock.Nil()),
		loader.VkQueryType(1),
		loader.Size(3),
		gomock.Not(gomock.Nil()),
		loader.Size(7),
	).DoAndReturn(func(
		device loader.VkDevice,
		accelerationStructureCount loader.Uint32,
		pAccelerationStructures *khr_acceleration_structure_loader.VkAccelerationStructureKHR,
		queryType loader.VkQueryType,
		dataSize loader.Size,
		pData unsafe.Pointer,
		stride loader.Size,
	) (common.VkResult, error) {
		accelStructSlice := unsafe.Slice(pAccelerationStructures, 2)
		require.Equal(t,
			[]khr_acceleration_structure_loader.VkAccelerationStructureKHR{accel1.Handle(), accel2.Handle()},
			accelStructSlice,
		)

		require.Equal(t, uint64(5), (uint64)(uintptr(pData)))

		return core1_0.VKSuccess, nil
	})

	res, err := driver.WriteAccelerationStructuresProperties(
		[]khr_acceleration_structure.AccelerationStructure{accel1, accel2},
		core1_0.QueryTypePipelineStatistics,
		3,
		unsafe.Pointer(uintptr(5)),
		7,
	)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
}

func TestPhysicalDeviceAccelerationStructureFeatures_PopulateCPointer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, mockLoader)

	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	expectedDevice := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	mockLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {

		info := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
		require.False(t, info.FieldByName("pNext").IsNil())

		pInnerFeatures := (*khr_acceleration_structure_loader.VkPhysicalDeviceAccelerationStructureFeaturesKHR)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()

		require.Equal(t, uint64(1000150013), features.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ACCELERATION_STRUCTURE_FEATURES_KHR
		require.True(t, features.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), features.FieldByName("accelerationStructure").Uint())
		require.Equal(t, uint64(0), features.FieldByName("accelerationStructureCaptureReplay").Uint())
		require.Equal(t, uint64(1), features.FieldByName("accelerationStructureIndirectBuild").Uint())
		require.Equal(t, uint64(0), features.FieldByName("accelerationStructureHostCommands").Uint())
		require.Equal(t, uint64(1), features.FieldByName("descriptorBindingAccelerationStructureUpdateAfterBind").Uint())

		*pDevice = expectedDevice.Handle()

		return core1_0.VKSuccess, nil
	})

	device, res, err := driver.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueueFamilyIndex: 0,
				QueuePriorities:  []float32{1},
			},
		},
		NextOptions: common.NextOptions{khr_acceleration_structure.PhysicalDeviceAccelerationStructureFeatures{
			AccelerationStructure:                                 true,
			AccelerationStructureCaptureReplay:                    false,
			AccelerationStructureIndirectBuild:                    true,
			AccelerationStructureHostCommands:                     false,
			DescriptorBindingAccelerationStructureUpdateAfterBind: true,
		}},
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NotNil(t, device)
	require.Equal(t, expectedDevice.Handle(), device.Device().Handle())
}

func TestPhysicalDeviceAccelerationStructureFeatures_PopulateOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, mockLoader)

	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	mockLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		info := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
		require.False(t, info.FieldByName("pNext").IsNil())

		pInnerFeatures := (*khr_acceleration_structure_loader.VkPhysicalDeviceAccelerationStructureFeaturesKHR)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()
		require.Equal(t, uint64(1000150013), features.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ACCELERATION_STRUCTURE_FEATURES_KHR
		require.True(t, features.FieldByName("pNext").IsNil())
		*(*uint64)(unsafe.Pointer(features.FieldByName("accelerationStructure").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(features.FieldByName("accelerationStructureCaptureReplay").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(features.FieldByName("accelerationStructureIndirectBuild").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(features.FieldByName("accelerationStructureHostCommands").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(features.FieldByName("descriptorBindingAccelerationStructureUpdateAfterBind").UnsafeAddr())) = uint64(1)
	})

	var baseFeatures core1_1.PhysicalDeviceFeatures2
	var features khr_acceleration_structure.PhysicalDeviceAccelerationStructureFeatures

	baseFeatures.NextOutData.Next = &features

	err := driver.GetPhysicalDeviceFeatures2(physicalDevice, &baseFeatures)
	require.NoError(t, err)

	require.Equal(t, khr_acceleration_structure.PhysicalDeviceAccelerationStructureFeatures{
		AccelerationStructure:                                 true,
		AccelerationStructureCaptureReplay:                    false,
		AccelerationStructureIndirectBuild:                    true,
		AccelerationStructureHostCommands:                     false,
		DescriptorBindingAccelerationStructureUpdateAfterBind: true,
	}, features)
}

func TestPhysicalDeviceAccelerationStructureProperties_PopulateOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, mockLoader)

	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	mockLoader.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pBaseProperties *loader.VkPhysicalDeviceProperties2) {
		baseProperties := reflect.ValueOf(pBaseProperties).Elem()

		require.Equal(t, uint64(1000059001), baseProperties.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
		require.False(t, baseProperties.FieldByName("pNext").IsNil())

		pProperties := (*khr_acceleration_structure_loader.VkPhysicalDeviceAccelerationStructurePropertiesKHR)(baseProperties.FieldByName("pNext").UnsafePointer())
		properties := reflect.ValueOf(pProperties).Elem()

		require.Equal(t, uint64(1000150014), properties.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ACCELERATION_STRUCTURE_PROPERTIES_KHR
		require.True(t, properties.FieldByName("pNext").IsNil())

		*(*uint64)(unsafe.Pointer(properties.FieldByName("maxGeometryCount").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("maxInstanceCount").UnsafeAddr())) = uint64(3)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("maxPrimitiveCount").UnsafeAddr())) = uint64(5)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("maxPerStageDescriptorAccelerationStructures").UnsafeAddr())) = uint64(7)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("maxPerStageDescriptorUpdateAfterBindAccelerationStructures").UnsafeAddr())) = uint64(11)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("maxDescriptorSetAccelerationStructures").UnsafeAddr())) = uint64(13)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("maxDescriptorSetUpdateAfterBindAccelerationStructures").UnsafeAddr())) = uint64(17)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("minAccelerationStructureScratchOffsetAlignment").UnsafeAddr())) = uint64(19)
	})

	var baseProperties core1_1.PhysicalDeviceProperties2
	var properties khr_acceleration_structure.PhysicalDeviceAccelerationStructureProperties

	baseProperties.Next = &properties

	err := driver.GetPhysicalDeviceProperties2(physicalDevice, &baseProperties)
	require.NoError(t, err)

	require.Equal(t, khr_acceleration_structure.PhysicalDeviceAccelerationStructureProperties{
		MaxGeometryCount:  1,
		MaxInstanceCount:  3,
		MaxPrimitiveCount: 5,
		MaxPerStageDescriptorAccelerationStructures:                7,
		MaxPerStageDescriptorUpdateAfterBindAccelerationStructures: 11,
		MaxDescriptorSetAccelerationStructures:                     13,
		MaxDescriptorSetUpdateAfterBindAccelerationStructures:      17,
		MinAccelerationStructureScratchOffsetAlignment:             19,
	}, properties)
}

func TestWriteDescriptorSetAccelerationStructure_PopulateCPointer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	driver := mocks1_1.InternalDeviceDriver(device, mockLoader)

	accel1 := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	accel2 := mock_acceleration_structure.NewDummyAccelerationStructure(device)
	accel3 := mock_acceleration_structure.NewDummyAccelerationStructure(device)

	sampler := mocks.NewDummySampler(device)
	imageView := mocks.NewDummyImageView(device)
	descriptorPool := mocks.NewDummyDescriptorPool(device)
	descriptorSet := mocks.NewDummyDescriptorSet(descriptorPool, device)

	mockLoader.EXPECT().VkUpdateDescriptorSets(
		device.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
		loader.Uint32(0),
		gomock.Nil(),
	).DoAndReturn(func(device loader.VkDevice, descriptorWriteCount loader.Uint32, pDescriptorWrites *loader.VkWriteDescriptorSet, descriptorCopyCount loader.Uint32, pDescriptorCopies *loader.VkCopyDescriptorSet) error {
		descriptorWriteSlice := unsafe.Slice(pDescriptorWrites, 1)
		descriptorWrite := reflect.ValueOf(&descriptorWriteSlice[0]).Elem()

		require.Equal(t, uint64(35), descriptorWrite.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
		require.False(t, descriptorWrite.FieldByName("pNext").IsNil())

		pInfo := (*khr_acceleration_structure_loader.VkWriteDescriptorSetAccelerationStructureKHR)(descriptorWrite.FieldByName("pNext").UnsafePointer())
		info := reflect.ValueOf(pInfo).Elem()
		require.Equal(t, uint64(1000150007), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), info.FieldByName("accelerationStructureCount").Uint())

		pStructures := (*khr_acceleration_structure_loader.VkAccelerationStructureKHR)(unsafe.Pointer(info.FieldByName("pAccelerationStructures").Elem().UnsafeAddr()))
		structs := unsafe.Slice(pStructures, 3)
		require.Equal(t, []khr_acceleration_structure_loader.VkAccelerationStructureKHR{
			accel1.Handle(), accel2.Handle(), accel3.Handle(),
		}, structs)

		return nil
	})

	err := driver.UpdateDescriptorSets([]core1_0.WriteDescriptorSet{
		{
			DstSet: descriptorSet,
			ImageInfo: []core1_0.DescriptorImageInfo{
				{
					Sampler:     sampler,
					ImageView:   imageView,
					ImageLayout: core1_0.ImageLayoutGeneral,
				},
			},
			NextOptions: common.NextOptions{khr_acceleration_structure.WriteDescriptorSetAccelerationStructure{
				AccelerationStructures: []khr_acceleration_structure.AccelerationStructure{
					accel1,
					accel2,
					accel3,
				},
			}},
		},
	}, nil)
	require.NoError(t, err)
}
