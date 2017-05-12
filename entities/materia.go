package entities

type Materia struct {
	Titulo          string `json bson:"titulo"`
	Subtitulo       string `json bson: "subtitulo"`
	Cover           string `json bson:"cover"`
	DataCriacao     string `json bson:"data_criacao"`
	DataModificacao string `json bson:"data_modificacao"`
	DataPublicacao  string `json bson:"data_publicacao"`
	Slug            string `json bson:"slug"`
	Id              string `json bson:"id"`
	Corpo           string `json bson:"corpo"`
}
