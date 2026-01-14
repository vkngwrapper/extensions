package khr_acceleration_structure

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

type GeometryTrianglesData struct {
	VertexFormat core1_0.Format
	VertexData   DeviceOrHostAddressConst
	VertexStride int
	MaxVertex    int

	IndexType     core1_0.IndexType
	IndexData     DeviceOrHostAddressConst
	TransformData DeviceOrHostAddressConst

	common.NextOptions
}

func (d GeometryTrianglesData) IsGeometryType() {}

func (d GeometryTrianglesData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAccelerationStructureGeometryTrianglesDataKHR{})))
	}

	info := (*C.VkAccelerationStructureGeometryTrianglesDataKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_TRIANGLES_DATA_KHR
	info.pNext = next

	info.vertexFormat = C.VkFormat(d.VertexFormat)
	info.vertexStride = C.VkDeviceSize(d.VertexStride)
	info.maxVertex = C.uint32_t(d.MaxVertex)
	if d.VertexData != nil {
		d.VertexData.PopulateAddressUnion(unsafe.Pointer(&info.vertexData))
	}

	info.indexType = C.VkIndexType(d.IndexType)
	if d.IndexData != nil {
		d.IndexData.PopulateAddressUnion(unsafe.Pointer(&info.indexData))
	}
	if d.TransformData != nil {
		d.TransformData.PopulateAddressUnion(unsafe.Pointer(&info.transformData))
	}

	return preallocatedPointer, nil
}
