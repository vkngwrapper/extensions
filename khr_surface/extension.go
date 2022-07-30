package khr_surface

import (
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	khr_surface_driver "github.com/vkngwrapper/extensions/khr_surface/driver"
)

// Extension contains all commands for the khr_surface extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html
type Extension interface {
	// CreateSurfaceFromHandle is used to create a Surface object from native platform or window data.
	// Generally, this method should be called from an integration for a windowing system that produces
	// Surface handles, and not by end users
	//
	// surfaceHandle - A surface handle produced by a windowing system
	CreateSurfaceFromHandle(surfaceHandle khr_surface_driver.VkSurfaceKHR) (Surface, error)
}

// CreateExtensionFromInstance produces an Extension object from an Insstance with
// khr_surface loaded
func CreateExtensionFromInstance(instance core1_0.Instance) *VulkanExtension {
	if !instance.IsInstanceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver:   khr_surface_driver.CreateDriverFromCore(instance.Driver()),
		instance: instance,
	}
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_surface_driver.Driver, instance core1_0.Instance) *VulkanExtension {
	return &VulkanExtension{
		driver:   driver,
		instance: instance,
	}
}

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver   khr_surface_driver.Driver
	instance core1_0.Instance
}

func (e *VulkanExtension) CreateSurfaceFromHandle(surfaceHandle khr_surface_driver.VkSurfaceKHR) (Surface, error) {
	coreDriver := e.instance.Driver()
	instanceHandle := e.instance.Handle()
	apiVersion := e.instance.APIVersion()

	surface := e.instance.Driver().ObjectStore().GetOrCreate(driver.VulkanHandle(surfaceHandle), driver.Core1_0, func() any {
		return &VulkanSurface{
			handle:            surfaceHandle,
			coreDriver:        coreDriver,
			instance:          instanceHandle,
			driver:            e.driver,
			minimumAPIVersion: apiVersion,
		}
	}).(*VulkanSurface)
	return surface, nil
}
