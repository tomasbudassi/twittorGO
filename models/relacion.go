package models

type Relacion struct {
	UsuarioID         string `bson:"usuarioId" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuarioRelacionID" json:"usuarioRelacionID"`
}
