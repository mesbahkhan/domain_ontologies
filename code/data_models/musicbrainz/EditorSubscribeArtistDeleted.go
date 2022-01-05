package dto

type EditorSubscribeArtistDeleted struct {
	Editor    int            `db:"editor"`
	GID       sql.NullString `db:"gid"`
	DeletedBy int            `db:"deleted_by"`
}
