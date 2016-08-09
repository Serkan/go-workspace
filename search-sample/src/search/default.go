package search

type defaultMather struct{}

func init() {
	var mather defaultMather
	Register("default", mather)
}

func (m defaultMather) Search(feed *Feed, term string) ([]*Result, error) {
	return nil, nil
}
