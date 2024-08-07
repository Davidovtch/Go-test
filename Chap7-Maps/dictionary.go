package main

type Dic map[string]string

const (
	ErrNotFound      = DictErr("could not find the word you're looking for")
	ErrWordExists    = DictErr("cannot add word because it already exists")
	ErrWordDontExist = DictErr("cannot update word because is doesn't exist")
)

type DictErr string

func (de DictErr) Error() string {
	return string(de)
}

func (d Dic) Search(key string) (string, error) {
	/*map lookups can return two values the first is the value paired
	with the key,and the second if the key was found successfully*/
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dic) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dic) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDontExist
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dic) Delete(key string) {
	delete(d, key)
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
