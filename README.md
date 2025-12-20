# vkngwrapper/extensions/v3

[![Go Reference](https://pkg.go.dev/badge/github.com/vkngwrapper/extensions/v3.svg)](https://pkg.go.dev/github.com/vkngwrapper/extensions/v3)

`go get github.com/vkngwrapper/core/v2`

Vkngwrapper (proununced "Viking Wrapper") is a handwritten cgo wrapper for the Vulkan graphics and compute API.
The goal is to produce fast, easy-to-use, low-go-allocation, and idiomatic Go code to communicate with your graphics
card and enable games and other graphical applications.

To learn more about how to use vkngwrapper, check out [the core library](https://github.com/vkngwrapper/core). For the
 future roadmap, see [the org page](https://github.com/vkngwrapper).

## Using Extensions

Once you've gotten the hang of the core library, you may wish to make use of one of Vullkan's many extensions. This
 library provides first-class support for a small-but-growing list of Vulkan extensions.  It's easy enough to use an extension:
 just like core 1.1 and core 1.2 in the core library, most of the contents of an extension are constants an structures, which
 can be accessed via the package name, such as [khr_imageless_framebuffer.FramebufferAttachmentImageInfo](https://pkg.go.dev/github.com/vkngwrapper/extensions/v3/khr_imageless_framebuffer#FramebufferAttachmentImageInfo)
 or [khr_sampler_mirror_clamp_to_edge.SamplerAddressModeMirrorClampToEdge](https://pkg.go.dev/github.com/vkngwrapper/extensions/v3/khr_sampler_mirror_clamp_to_edge#pkg-constants).

For extensions that add new commands, you can create an extension object using `CreateExtensionFromInstance` for 
 instance extensions, or `CreateExtensionFromDevice` for device extensions.  Commands can be called on the new Extension
 object. These methods will return `nil` if the `Instance` or `Device` was not created with the extension in question
 active.  For your convenience, `ExtensionName` constants, such as `khr_external_memory.ExtensionName` are provided for
 all extensions, which you can pass to `Instance`/`Device` creation.

Example:

```go
	debugLoader := ext_debug_utils.CreateExtensionFromInstance(instance)
	debugMessenger, _, err := debugLoader.CreateDebugUtilsMessenger(instance, nil, ext_debug_utils.DebugUtilsMessengerCreateInfo{
      MessageSeverity: ext_debug_utils.SeverityError | ext_debug_utils.SeverityWarning,
      MessageType:     ext_debug_utils.TypeGeneral | ext_debug_utils.TypeValidation | ext_debug_utils.TypePerformance,
      UserCallback:    logValidationFunc,
    })
	if err != nil {
		return err
	}
```

### Supporting Promoted Extensions

From time to time you may wish to use a particular piece of functionality added by a since-promoted extension if
 either the promoted core API version or the extension are present. But given that the two APIs are often different, how
 can that be easily done?

* **For data structures such as Options or OutData**, simply use the types from the core package.
* **For extension functions**, there should be a shim package included with the extension if the extension has
  callable Vulkan functions. This shim package includes a `Shim` interface that can be received and called from your
  code, as well as a `NewShim` function that allows you to pass in an extension and receive a Shim. The relevant
  core Vulkan object already satisfies the Shim interface.

For instance, `khr_get_memory_requirements2_shim.Shim` adds an interface met by `core1_1.Device`, and also adds a 
 `NewShim()` method that accepts a `khr_get_memory_requirements2.Extension` and `core1_0.Device` and provides
 an object that satisfies `Shim`. By using a `Shim` object to interact with these methods, it is easy to
 interact with them in a flexible way.

## Supported Extensions

* Three major extensions used by most applications developers: vk_ext_debug_utils, vk_khr_surface, and vk_khr_swapchain
* vk_khr_portability_subset and vk_khr_portability_enumeration, which can be used to add Mac/iOS support to your Vulkan applications
* 23 extensions which were promoted to core 1.1
* 24 extensions which were promoted to core 1.2
* Four extensions used by [arsenal/vam](https://github.com/vkngwrapper/arsenal/tree/main/vam) for memory management:
   amd_device_coherent_memory, ext_memory_budget, ext_memory_priority, and khr_maintenance4.

It's also easy to add support for new extensions! Feel free to contribute your favorite Vulkan extension!
