package mock_surface

import (
	"math/rand"
	"unsafe"

	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/extensions/v3/khr_surface"
	khr_surface_loader "github.com/vkngwrapper/extensions/v3/khr_surface/loader"
)

func NewDummySurface(instance core1_0.Instance) khr_surface.Surface {
	return khr_surface.InternalSurface(instance.Handle(), NewFakeSurfaceHandle(), instance.APIVersion())
}

func fakePointer() unsafe.Pointer {
	return unsafe.Pointer(uintptr(rand.Int()))
}

func NewFakeSurfaceHandle() khr_surface_loader.VkSurfaceKHR {
	return khr_surface_loader.VkSurfaceKHR(fakePointer())
}
