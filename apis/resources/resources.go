package resources

type Resource struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func NewResource() *Resource {
	return &Resource{}
}

// todo
func (u *Resource) CreatedSuccess()
func (u *Resource) DeletedSuccess()
func (u *Resource) UpdatedSuccess()
func (u *Resource) ListAndCount()
