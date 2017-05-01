package gocraft

type Mesher interface {
	Generate(chunk *Chunk, vao *Vao)
}

type StupidMesher struct {
}

func (sm *StupidMesher) Generate(chunk *Chunk, vao *Vao) {

}
