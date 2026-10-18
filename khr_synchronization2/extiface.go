package khr_synchronization2

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

//go:generate go run go.uber.org/mock/mockgen -source extiface.go -destination ./mocks/extension.go -package mock_synchronization2

// ExtensionDriver contains all commands for the khr_synchronization2
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VK_KHR_synchronization2.html
type ExtensionDriver interface {
	// CmdPipelineBarrier2 inserts a memory dependency
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// dependencyInfo - defines the scopes of this operation
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdPipelineBarrier2.html
	CmdPipelineBarrier2(commandBuffer core1_0.CommandBuffer, dependencyInfo DependencyInfo) error

	// CmdResetEvent2 resets an event object to non-signaled state
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// event - the event that will be unsignaled.
	//
	// stageMask - a PipelineStageFlags2 mask of pipeline stages used to determine the first synchronization scope
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdResetEvent2.html
	CmdResetEvent2(commandBuffer core1_0.CommandBuffer, event core1_0.Event, stageMask PipelineStageFlags2)

	// CmdSetEvent2 sets an event object to signaled state
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// event - the event that will be signaled
	//
	// dependencyInfo - defines the first scopes of this operation
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetEvent2.html
	CmdSetEvent2(commandBuffer core1_0.CommandBuffer, event core1_0.Event, dependencyInfo DependencyInfo) error

	// CmdWaitEvents2 waits for one or more events
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// events - a slice of events to wait on
	//
	// dependencyInfos - a slice of DependencyInfo structures, defining the second synchronization scope
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdWaitEvents2.html
	CmdWaitEvents2(commandBuffer core1_0.CommandBuffer, events []core1_0.Event, dependencyInfos []DependencyInfo) error

	// CmdWriteTimestamp2 writes a device timestamp into a query object
	//
	// commandBuffer - the command buffer into which the command will be recorded
	//
	// stage - specifies a stage of the pipeline
	//
	// queryPool - the query pool that will manage the timestamp
	//
	// query - the query within the query pool that will contain the timestamp
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdWriteTimestamp2.html
	CmdWriteTimestamp2(commandBuffer core1_0.CommandBuffer, stage PipelineStageFlags2, queryPool core1_0.QueryPool, query int)

	// QueueSubmit2 submits command buffers to a queue
	//
	// queue - the queue that the command buffers will be submitted to
	//
	// submits - a slice of SubmitInfo2 structures, each specifying a command buffer submission
	//           batch. Command buffers and semaphores specified in this array may be accessed
	//           at any point until the queue operations they define complete execution on the
	//           device.
	//
	// fence -  is an optional handle to a fence to be signaled once all submitted command
	//          buffers have completed execution. If fence is not nil, it defines a fence signal
	//          operation. If it is not nil, fence may be accessed at any point until this command
	//          completes on the device.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkQueueSubmit2.html
	QueueSubmit2(queue core1_0.Queue, submits []SubmitInfo2, fence core1_0.Fence) (common.VkResult, error)
}
