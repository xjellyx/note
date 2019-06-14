package main

type Serve struct {
}

func NewServe(s *Serve) (d *Serve) {
	if s != nil {
		d = s
	} else {
		d = new(Serve)
	}
	return
}

type Data struct {
	Name string
}

func (s *Serve) GetData() (ret *Data, err error) {

	return
}
