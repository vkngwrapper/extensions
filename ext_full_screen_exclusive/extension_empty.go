//go:build !windows

package ext_full_screen_exclusive

func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver, instance core1_0.Instance) ExtensionDriver {
	return nil
}
