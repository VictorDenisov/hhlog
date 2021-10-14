package main

type StringArray []string

func (i *StringArray) String() string {
	return "my string representation"
}

func (i *StringArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}
