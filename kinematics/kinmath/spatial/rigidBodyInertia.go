package spatial

import (
	"github.com/go-gl/mathgl/mgl64"
)

type RigidBodyInertia struct {
	Inertia mgl64.Vec3
	Cog     mgl64.Vec3
	Mass    mgl64.Vec3
}
