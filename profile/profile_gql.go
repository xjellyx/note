package ctrl

import "git.yichui.net/tudy/go-rest"

// UID
func (n *NodeBusiness) UID() (ret string, err error) {
	if n.Data != nil {
		ret = n.Data.UID
	}
	return
}

// Category
func (n *NodeBusiness) Category() (ret string, err error) {
	if n.Data != nil {
		ret = n.Data.Category
	}
	return
}

// Name
func (n *NodeBusiness) Name() (ret string, err error) {
	if n.Data != nil {
		ret = n.Data.Name
	}
	return
}

// Icon
func (n *NodeBusiness) Icon() (ret string, err error) {
	if n.Data != nil {
		ret = n.Data.Icon
	}
	return
}

// Address
func (n *NodeBusiness) Address() (ret string, err error) {
	if n.Data != nil {
		ret = n.Data.Address
	}
	return
}

// Contact
func (n *NodeBusiness) Contact() (ret string, err error) {
	if n.Data != nil {
		ret = n.Data.Contact
	}
	return
}

// Phone
func (n *NodeBusiness) Phone() (ret string, err error) {
	if n.Data != nil {
		ret = n.Data.Phone
	}
	return
}

// CreateTime
func (n *NodeBusiness) CreateTime() (ret *string) {
	if n.Data != nil && n.Data.CreateTime.Unix() > 0 {
		_t := rest.PubTimeToStr(n.Data.CreateTime)
		ret = &_t
	}
	return
}

// UpdateTime
func (n *NodeBusiness) UpdateTime() (ret *string) {
	if n.Data != nil && n.Data.UpdateTime.Unix() > 0 {
		_t := rest.PubTimeToStr(n.Data.UpdateTime)
		ret = &_t
	}
	return
}

const standSchema = `
	type Business {
		#
		uid: String!
		#
		category: String!
		#
		name: String!
		#
		icon: String!
		#
		address: String!
		#
		contact: String!
		#
		phone: String!
		# 
		createTime: String
		# 
		updateTime: String
	}
`
