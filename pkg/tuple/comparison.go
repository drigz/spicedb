package tuple

import (
	core "github.com/authzed/spicedb/pkg/proto/core/v1"
	"google.golang.org/protobuf/proto"
)

// ONREqual checks if two ObjectAndRelation instances are equal.
func ONREqual(lhs, rhs ObjectAndRelation) bool {
	return lhs == rhs
}

// ONREqualOrWildcard checks if an ObjectAndRelation matches another ObjectAndRelation or is a wildcard.
func ONREqualOrWildcard(onr, target ObjectAndRelation) bool {
	return ONREqual(onr, target) || (onr.ObjectID == PublicWildcard && onr.ObjectType == target.ObjectType)
}

// Equal returns true if the two relationships are exactly the same.
func Equal(lhs, rhs Relationship) bool {
	return ONREqual(lhs.Resource, rhs.Resource) && ONREqual(lhs.Subject, rhs.Subject) && caveatEqual(lhs.OptionalCaveat, rhs.OptionalCaveat)
}

func caveatEqual(lhs, rhs *core.ContextualizedCaveat) bool {
	if lhs == nil && rhs == nil {
		return true
	}

	if lhs == nil || rhs == nil {
		return false
	}

	return lhs.CaveatName == rhs.CaveatName && proto.Equal(lhs.Context, rhs.Context)
}