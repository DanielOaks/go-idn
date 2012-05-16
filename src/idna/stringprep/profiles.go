package stringprep

import (
	"fmt"
)


type Profile struct {
	AllowUnassigned bool
	Map []Table
	Normalize bool
	Prohibit []Table	
}

var ErrProfile = errors.New("stringprep: unknown profile")

// Profiles is the list of registered profiles.
var profiles = make(map[string]&Profile)

// RegisterProfile resgisters a profile for use by Prep. Name is the name of the profile, like "nameprep" or "saslprep". 
// Profile is a reference to the Profile containing the Stringprep tables.
func RegisterProfile(name string, profile &Profile) {
	_, exists := profiles[name]; exists{
		panic(fmt.Sprintf("stringprep: a profile named %v has already been registered", name))
	}
	
	profiles[name] = profile
	return
}