package ext_debug_utils

import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	ext_driver "github.com/vkngwrapper/extensions/v2/ext_debug_utils/driver"
)

//go:generate mockgen -source extension.go -destination ./mocks/extension.go -package mock_debugutils

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver ext_driver.Driver
}

// Extension contains all the commands for the ext_debug_utils extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html
type Extension interface {
	// CreateDebugUtilsMessenger creates a debug messenger object
	//
	// instance - the instance the messenger will be used with
	//
	// allocator - controls host memory allocation
	//
	// o - contains the callback, as well as defining conditions under which this messenger will
	// trigger the callback
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugUtilsMessengerEXT.html
	CreateDebugUtilsMessenger(instance core1_0.Instance, allocator *driver.AllocationCallbacks, o DebugUtilsMessengerCreateInfo) (DebugUtilsMessenger, common.VkResult, error)

	// CmdBeginDebugUtilsLabel opens a CommandBuffer debug label region
	//
	// commandBuffer - the CommandBuffer into which the command is recorded
	//
	// label - specifies parameters of the label region to open
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginDebugUtilsLabelEXT.html
	CmdBeginDebugUtilsLabel(commandBuffer core1_0.CommandBuffer, label DebugUtilsLabel) error
	// CmdEndDebugUtilsLabel closes a CommandBuffer label region
	//
	// commandBuffer - the CommandBuffer into which the command is recorded
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndDebugUtilsLabelEXT.html
	CmdEndDebugUtilsLabel(commandBuffer core1_0.CommandBuffer)
	// CmdInsertDebugUtilsLabel inserts a label into a CommandBuffer
	//
	// commandBuffer - the CommandBuffer into which the command is recorded
	//
	// label - specifies parameters of the label to insert
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdInsertDebugUtilsLabelEXT.html
	CmdInsertDebugUtilsLabel(commandBuffer core1_0.CommandBuffer, label DebugUtilsLabel) error

	// QueueBeginDebugUtilsLabel opens a Queue debug label region
	//
	// queue - The Queue in which to start a debug label region
	//
	// label - Specifies parameters of the label region to open
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueBeginDebugUtilsLabelEXT.html
	QueueBeginDebugUtilsLabel(queue core1_0.Queue, label DebugUtilsLabel) error
	// QueueEndDebugUtilsLabel closes a Queue debug label region
	//
	// queue - The Queue in which a debug label region should be closed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueEndDebugUtilsLabelEXT.html
	QueueEndDebugUtilsLabel(queue core1_0.Queue)
	// QueueInsertDebugUtilsLabel inserts a label into a Queue
	//
	// queue - The Queue into which a debug label will be inserted
	//
	// label - Specifies parameters of the label to insert
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueInsertDebugUtilsLabelEXT.html
	QueueInsertDebugUtilsLabel(queue core1_0.Queue, label DebugUtilsLabel) error

	// SetDebugUtilsObjectName gives a user-friendly name to an object
	//
	// device - The Device that created the object
	//
	// name - Specifies parameters of the name to set on the object
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetDebugUtilsObjectNameEXT.html
	SetDebugUtilsObjectName(device core1_0.Device, name DebugUtilsObjectNameInfo) (common.VkResult, error)
	// SetDebugUtilsObjectTag attaches arbitrary data to an object
	//
	// device - The Device that created the object
	//
	// tag - Specifies parameters of the tag to attach to the object
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetDebugUtilsObjectTagEXT.html
	SetDebugUtilsObjectTag(device core1_0.Device, tag DebugUtilsObjectTagInfo) (common.VkResult, error)

	// SubmitDebugUtilsMessage injects a message into a debug stream
	//
	// instance - The debug stream's Instance
	//
	// severity - Specifies the severity of this event/message
	//
	// types - Specifies which type of event(s) to identify with this message
	//
	// data - All the callback-related data
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSubmitDebugUtilsMessageEXT.html
	SubmitDebugUtilsMessage(instance core1_0.Instance, severity DebugUtilsMessageSeverityFlags, types DebugUtilsMessageTypeFlags, data DebugUtilsMessengerCallbackData) error
}

// CreateExtensionFromInstance produces an Extension object from an Instance with
// ext_debug_utils loaded
func CreateExtensionFromInstance(instance core1_0.Instance) *VulkanExtension {
	driver := ext_driver.CreateDriverFromCore(instance.Driver())

	if !instance.IsInstanceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateExtensionFromDriver(driver)
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver ext_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (l *VulkanExtension) CreateDebugUtilsMessenger(instance core1_0.Instance, allocation *driver.AllocationCallbacks, o DebugUtilsMessengerCreateInfo) (DebugUtilsMessenger, common.VkResult, error) {
	if instance == nil {
		panic("instance cannot be nil")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var messenger ext_driver.VkDebugUtilsMessengerEXT
	res, err := l.driver.VkCreateDebugUtilsMessengerEXT(instance.Handle(), (*ext_driver.VkDebugUtilsMessengerCreateInfoEXT)(createInfo), allocation.Handle(), &messenger)

	if err != nil {
		return nil, res, err
	}

	coreDriver := instance.Driver()
	newMessenger := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(messenger), driver.Core1_0, func() any {
		return &VulkanDebugUtilsMessenger{
			coreDriver: coreDriver,
			handle:     messenger,
			instance:   instance.Handle(),
			driver:     l.driver,
		}
	}).(*VulkanDebugUtilsMessenger)

	return newMessenger, res, nil
}

func (l *VulkanExtension) CmdBeginDebugUtilsLabel(commandBuffer core1_0.CommandBuffer, label DebugUtilsLabel) error {
	if commandBuffer == nil {
		panic("commandBuffer cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	labelPtr, err := common.AllocOptions(arena, label)
	if err != nil {
		return err
	}

	l.driver.VKCmdBeginDebugUtilsLabelEXT(commandBuffer.Handle(), (*ext_driver.VkDebugUtilsLabelEXT)(labelPtr))

	return nil
}

func (l *VulkanExtension) CmdEndDebugUtilsLabel(buffer core1_0.CommandBuffer) {
	if buffer == nil {
		panic("buffer cannot be nil")
	}
	l.driver.VkCmdEndDebugUtilsLabelEXT(buffer.Handle())
}

func (l *VulkanExtension) CmdInsertDebugUtilsLabel(buffer core1_0.CommandBuffer, label DebugUtilsLabel) error {
	if buffer == nil {
		panic("buffer cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	labelPtr, err := common.AllocOptions(arena, label)
	if err != nil {
		return err
	}

	l.driver.VkCmdInsertDebugUtilsLabelEXT(buffer.Handle(), (*ext_driver.VkDebugUtilsLabelEXT)(labelPtr))

	return nil
}

func (l *VulkanExtension) QueueBeginDebugUtilsLabel(queue core1_0.Queue, label DebugUtilsLabel) error {
	if queue == nil {
		panic("queue cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	labelPtr, err := common.AllocOptions(arena, label)
	if err != nil {
		return err
	}

	l.driver.VkQueueBeginDebugUtilsLabelEXT(queue.Handle(), (*ext_driver.VkDebugUtilsLabelEXT)(labelPtr))

	return nil
}

func (l *VulkanExtension) QueueEndDebugUtilsLabel(queue core1_0.Queue) {
	if queue == nil {
		panic("queue cannot be nil")
	}
	l.driver.VkQueueEndDebugUtilsLabelEXT(queue.Handle())
}

func (l *VulkanExtension) QueueInsertDebugUtilsLabel(queue core1_0.Queue, label DebugUtilsLabel) error {
	if queue == nil {
		panic("queue cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	labelPtr, err := common.AllocOptions(arena, label)
	if err != nil {
		return err
	}

	l.driver.VkQueueInsertDebugUtilsLabelEXT(queue.Handle(), (*ext_driver.VkDebugUtilsLabelEXT)(labelPtr))

	return nil
}

func (l *VulkanExtension) SetDebugUtilsObjectName(device core1_0.Device, name DebugUtilsObjectNameInfo) (common.VkResult, error) {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	namePtr, err := common.AllocOptions(arena, name)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return l.driver.VkSetDebugUtilsObjectNameEXT(device.Handle(), (*ext_driver.VkDebugUtilsObjectNameInfoEXT)(namePtr))
}

func (l *VulkanExtension) SetDebugUtilsObjectTag(device core1_0.Device, tag DebugUtilsObjectTagInfo) (common.VkResult, error) {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	tagPtr, err := common.AllocOptions(arena, tag)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return l.driver.VkSetDebugUtilsObjectTagEXT(device.Handle(), (*ext_driver.VkDebugUtilsObjectTagInfoEXT)(tagPtr))
}

func (l *VulkanExtension) SubmitDebugUtilsMessage(instance core1_0.Instance, severity DebugUtilsMessageSeverityFlags, types DebugUtilsMessageTypeFlags, data DebugUtilsMessengerCallbackData) error {
	if instance == nil {
		panic("instance cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	callbackPtr, err := common.AllocOptions(arena, &data)
	if err != nil {
		return err
	}

	l.driver.VkSubmitDebugUtilsMessageEXT(instance.Handle(),
		ext_driver.VkDebugUtilsMessageSeverityFlagBitsEXT(severity),
		ext_driver.VkDebugUtilsMessageTypeFlagsEXT(types),
		(*ext_driver.VkDebugUtilsMessengerCallbackDataEXT)(callbackPtr))

	return nil
}
