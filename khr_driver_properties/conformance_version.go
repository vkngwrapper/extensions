package khr_driver_properties

// ConformanceVersion contains the comformance test suite version the implementation is
// compliant with
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConformanceVersionKHR.html
type ConformanceVersion struct {
	// Major is the major version number of the conformance test suite
	Major uint8
	// Minor is the minor version number of the conformance test suite
	Minor uint8
	// Subminor is the subminor version number of the conformance test suite
	Subminor uint8
	// Patch is the patch version number of the conformance test suite
	Patch uint8
}

// IsAtLeast returns true if the other ConformanceVersion is at least as high as this one
func (v ConformanceVersion) IsAtLeast(other ConformanceVersion) bool {
	if v.Major > other.Major {
		return true
	} else if v.Major < other.Major {
		return false
	}

	if v.Minor > other.Minor {
		return true
	} else if v.Minor < other.Minor {
		return false
	}

	if v.Subminor > other.Subminor {
		return true
	} else if v.Subminor < other.Subminor {
		return false
	}

	return v.Patch >= other.Patch
}
