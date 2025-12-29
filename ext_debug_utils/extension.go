package ext_debug_utils

import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	ext_driver "github.com/vkngwrapper/extensions/v3/ext_debug_utils/loader"
)

//go:generate mockgen -source extension.go -destination ./mocks/extension.go -package mock_debugutils

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	loader     ext_driver.Loader
	coreLoader loader.Loader
	instance   core1_0.Instance
}

// ExtensionDriver contains all the commands for the ext_debug_utils extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html
type ExtensionDriver interface {
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
	CreateDebugUtilsMessenger(allocator *loader.AllocationCallbacks, o DebugUtilsMessengerCreateInfo) (DebugUtilsMessenger, common.VkResult, error)

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
	SubmitDebugUtilsMessage(severity DebugUtilsMessageSeverityFlags, types DebugUtilsMessageTypeFlags, data DebugUtilsMessengerCallbackData) error

	DestroyDebugUtilsMessenger(messenger DebugUtilsMessenger, callbacks *loader.AllocationCallbacks)
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from an Instance with
// ext_debug_utils loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.CoreInstanceDriver) ExtensionDriver {
	driver := ext_driver.CreateLoaderFromCore(coreDriver.Loader())
	instance := coreDriver.Instance()

	if !instance.IsInstanceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateDriverFromLoader(driver, coreDriver.Loader(), instance)
}

// CreateDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateDriverFromLoader(loader ext_driver.Loader, coreLoader loader.Loader, instance core1_0.Instance) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		loader:     loader,
		coreLoader: coreLoader,
		instance:   instance,
	}
}

func (l *VulkanExtensionDriver) CreateDebugUtilsMessenger(allocation *loader.AllocationCallbacks, o DebugUtilsMessengerCreateInfo) (DebugUtilsMessenger, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return DebugUtilsMessenger{}, core1_0.VKErrorUnknown, err
	}

	var messenger ext_driver.VkDebugUtilsMessengerEXT
	res, err := l.loader.VkCreateDebugUtilsMessengerEXT(l.instance.Handle(), (*ext_driver.VkDebugUtilsMessengerCreateInfoEXT)(createInfo), allocation.Handle(), &messenger)

	if err != nil {
		return DebugUtilsMessenger{}, res, err
	}

	newMessenger := DebugUtilsMessenger{
		handle:     messenger,
		instance:   l.instance.Handle(),
		apiVersion: l.instance.APIVersion(),
	}

	return newMessenger, res, nil
}

func (l *VulkanExtensionDriver) CmdBeginDebugUtilsLabel(commandBuffer core1_0.CommandBuffer, label DebugUtilsLabel) error {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	labelPtr, err := common.AllocOptions(arena, label)
	if err != nil {
		return err
	}

	l.loader.VKCmdBeginDebugUtilsLabelEXT(commandBuffer.Handle(), (*ext_driver.VkDebugUtilsLabelEXT)(labelPtr))

	return nil
}

func (l *VulkanExtensionDriver) CmdEndDebugUtilsLabel(buffer core1_0.CommandBuffer) {
	if !buffer.Initialized() {
		panic("buffer cannot be uninitialized")
	}
	l.loader.VkCmdEndDebugUtilsLabelEXT(buffer.Handle())
}

func (l *VulkanExtensionDriver) CmdInsertDebugUtilsLabel(buffer core1_0.CommandBuffer, label DebugUtilsLabel) error {
	if !buffer.Initialized() {
		panic("buffer cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	labelPtr, err := common.AllocOptions(arena, label)
	if err != nil {
		return err
	}

	l.loader.VkCmdInsertDebugUtilsLabelEXT(buffer.Handle(), (*ext_driver.VkDebugUtilsLabelEXT)(labelPtr))

	return nil
}

func (l *VulkanExtensionDriver) QueueBeginDebugUtilsLabel(queue core1_0.Queue, label DebugUtilsLabel) error {
	if !queue.Initialized() {
		panic("queue cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	labelPtr, err := common.AllocOptions(arena, label)
	if err != nil {
		return err
	}

	l.loader.VkQueueBeginDebugUtilsLabelEXT(queue.Handle(), (*ext_driver.VkDebugUtilsLabelEXT)(labelPtr))

	return nil
}

func (l *VulkanExtensionDriver) QueueEndDebugUtilsLabel(queue core1_0.Queue) {
	if !queue.Initialized() {
		panic("queue cannot be uninitialized")
	}
	l.loader.VkQueueEndDebugUtilsLabelEXT(queue.Handle())
}

func (l *VulkanExtensionDriver) QueueInsertDebugUtilsLabel(queue core1_0.Queue, label DebugUtilsLabel) error {
	if !queue.Initialized() {
		panic("queue cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	labelPtr, err := common.AllocOptions(arena, label)
	if err != nil {
		return err
	}

	l.loader.VkQueueInsertDebugUtilsLabelEXT(queue.Handle(), (*ext_driver.VkDebugUtilsLabelEXT)(labelPtr))

	return nil
}

func (l *VulkanExtensionDriver) SetDebugUtilsObjectName(device core1_0.Device, name DebugUtilsObjectNameInfo) (common.VkResult, error) {
	if !device.Initialized() {
		panic("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	namePtr, err := common.AllocOptions(arena, name)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return l.loader.VkSetDebugUtilsObjectNameEXT(device.Handle(), (*ext_driver.VkDebugUtilsObjectNameInfoEXT)(namePtr))
}

func (l *VulkanExtensionDriver) SetDebugUtilsObjectTag(device core1_0.Device, tag DebugUtilsObjectTagInfo) (common.VkResult, error) {
	if !device.Initialized() {
		panic("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	tagPtr, err := common.AllocOptions(arena, tag)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return l.loader.VkSetDebugUtilsObjectTagEXT(device.Handle(), (*ext_driver.VkDebugUtilsObjectTagInfoEXT)(tagPtr))
}

func (l *VulkanExtensionDriver) SubmitDebugUtilsMessage(severity DebugUtilsMessageSeverityFlags, types DebugUtilsMessageTypeFlags, data DebugUtilsMessengerCallbackData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	callbackPtr, err := common.AllocOptions(arena, &data)
	if err != nil {
		return err
	}

	l.loader.VkSubmitDebugUtilsMessageEXT(l.instance.Handle(),
		ext_driver.VkDebugUtilsMessageSeverityFlagBitsEXT(severity),
		ext_driver.VkDebugUtilsMessageTypeFlagsEXT(types),
		(*ext_driver.VkDebugUtilsMessengerCallbackDataEXT)(callbackPtr))

	return nil
}

func (l *VulkanExtensionDriver) DestroyDebugUtilsMessenger(messenger DebugUtilsMessenger, callbacks *loader.AllocationCallbacks) {
	l.loader.VkDestroyDebugUtilsMessengerEXT(messenger.instance, messenger.handle, callbacks.Handle())
}
