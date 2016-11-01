package detour

const (
	NAVMESHSET_MAGIC   int32 = 'M'<<24 | 'S'<<16 | 'E'<<8 | 'T' //'MSET';
	NAVMESHSET_VERSION int32 = 1
)

const (
	/// The maximum number of vertices per navigation polygon.
	DT_VERTS_PER_POLYGON int = 6
)
