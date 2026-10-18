package khr_synchronization2

import (
	"fmt"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	khr_synchronization2_loader "github.com/vkngwrapper/extensions/v3/khr_synchronization2/loader"
)

type VulkanExtensionDriver struct {
	driver khr_synchronization2_loader.Loader
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// ext_host_query_reset loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver) ExtensionDriver {
	device := coreDriver.Device()

	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		driver: khr_synchronization2_loader.CreateLoaderFromCore(coreDriver.Loader()),
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_synchronization2_loader.Loader) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
	}
}

func (e *VulkanExtensionDriver) CmdPipelineBarrier2(commandBuffer core1_0.CommandBuffer, dependencyInfo DependencyInfo) error {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)
	infoPtr, err := common.AllocOptions(arena, dependencyInfo)
	if err != nil {
		return err
	}

	e.driver.VkCmdPipelineBarrier2KHR(
		commandBuffer.Handle(),
		(*khr_synchronization2_loader.VkDependencyInfo)(infoPtr),
	)

	return nil
}

func (e *VulkanExtensionDriver) CmdResetEvent2(commandBuffer core1_0.CommandBuffer, event core1_0.Event, stageMask PipelineStageFlags2) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !event.Initialized() {
		panic("event cannot be uninitialized")
	}

	e.driver.VkCmdResetEvent2KHR(
		commandBuffer.Handle(),
		event.Handle(),
		khr_synchronization2_loader.VkPipelineStageFlags2(stageMask),
	)
}

func (e *VulkanExtensionDriver) CmdSetEvent2(commandBuffer core1_0.CommandBuffer, event core1_0.Event, dependencyInfo DependencyInfo) error {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !event.Initialized() {
		panic("event cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)
	infoPtr, err := common.AllocOptions(arena, dependencyInfo)
	if err != nil {
		return err
	}

	e.driver.VkCmdSetEvent2KHR(
		commandBuffer.Handle(),
		event.Handle(),
		(*khr_synchronization2_loader.VkDependencyInfo)(infoPtr),
	)

	return nil
}

func (e *VulkanExtensionDriver) CmdWaitEvents2(commandBuffer core1_0.CommandBuffer, events []core1_0.Event, dependencyInfos []DependencyInfo) error {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if len(events) != len(dependencyInfos) {
		panic("the same number of events and dependencyInfos must be provided")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	eventCount := len(events)
	var eventPtr *loader.VkEvent
	var dependencyInfoPtr *khr_synchronization2_loader.VkDependencyInfo

	if eventCount > 0 {
		eventPtr = (*loader.VkEvent)(arena.Malloc(eventCount * int(unsafe.Sizeof([1]loader.VkEvent{}))))
		eventSlice := ([]loader.VkEvent)(unsafe.Slice(eventPtr, eventCount))

		for i := 0; i < eventCount; i++ {
			if events[i].Handle() == 0 {
				return fmt.Errorf("element %d of the events slice was uninitialized", i)
			}
			eventSlice[i] = loader.VkEvent(unsafe.Pointer(events[i].Handle()))
		}

		var err error
		dependencyInfoPtr, err = common.AllocOptionSlice[khr_synchronization2_loader.VkDependencyInfo](arena, dependencyInfos)
		if err != nil {
			return err
		}
	}

	e.driver.VkCmdWaitEvents2KHR(
		commandBuffer.Handle(),
		loader.Uint32(eventCount),
		eventPtr,
		dependencyInfoPtr,
	)
	return nil
}

func (e *VulkanExtensionDriver) CmdWriteTimestamp2(commandBuffer core1_0.CommandBuffer, stage PipelineStageFlags2, queryPool core1_0.QueryPool, query int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	if !queryPool.Initialized() {
		panic("queryPool cannot be uninitialized")
	}

	e.driver.VkCmdWriteTimestamp2KHR(
		commandBuffer.Handle(),
		khr_synchronization2_loader.VkPipelineStageFlags2(stage),
		queryPool.Handle(),
		loader.Uint32(query),
	)
}

func (e *VulkanExtensionDriver) QueueSubmit2(queue core1_0.Queue, submits []SubmitInfo2, fence core1_0.Fence) (common.VkResult, error) {
	if !queue.Initialized() {
		panic("queue cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	var submitInfoPtr *khr_synchronization2_loader.VkSubmitInfo2
	var err error
	submitCount := len(submits)
	var fenceHandle loader.VkFence = loader.VkFence(0)

	if submitCount > 0 {
		submitInfoPtr, err = common.AllocOptionSlice[khr_synchronization2_loader.VkSubmitInfo2](arena, submits)
		if err != nil {
			return core1_0.VKErrorUnknown, err
		}
	}

	if fence.Initialized() {
		fenceHandle = fence.Handle()
	}

	return e.driver.VkQueueSubmit2(
		queue.Handle(),
		loader.Uint32(submitCount),
		submitInfoPtr,
		fenceHandle,
	)
}
