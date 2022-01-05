package models

// El formato bson es necesario para guardar en la BD
type Relacion struct {
	UsuarioID         string `bson:"usuarioId" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuarioRelacionID" json:"usuarioRelacionID"`
}
