package ext_private_data

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

// PipelineCacheCreateFlags specifies the behavior of the pipeline cache
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineCacheCreateFlagBits.html
type PrivateDataSlotCreateFlags uint32

var privateDataSlotCreateFlags = common.NewFlagStringMapping[PrivateDataSlotCreateFlags]()

func (f PrivateDataSlotCreateFlags) Register(str string) {
	privateDataSlotCreateFlags.Register(f, str)
}

func (f PrivateDataSlotCreateFlags) String() string {
	return privateDataSlotCreateFlags.FlagsToString(f)
}

const (
	// ObjectTypePrivateDataSlot specifies a PrivateDataSlot handle
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkObjectType.html
	ObjectTypePrivateDataSlot core1_0.ObjectType = C.VK_OBJECT_TYPE_PRIVATE_DATA_SLOT_EXT

	// ExtensionName is "VK_EXT_private_data"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_EXT_private_data.html
	ExtensionName string = C.VK_EXT_PRIVATE_DATA_EXTENSION_NAME
)
