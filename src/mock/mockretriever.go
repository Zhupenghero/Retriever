package mock


type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {

}

func (r *Retriever)Get(url string)string  {
	return r.Contents
}

func (r *Retriever) Post(url string,form map[string]string)string{
	r.Contents=form["contents"]				//map[键]=值
	return "ok"
}
