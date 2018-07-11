package models

type HalLinkable interface {
	Clean() interface{}
}

func (m *ResponseListResources) Clean() interface{} {
	return m.Embedded.Clean()
}

func (m *ResourceTfModule) Clean() interface{} {
	m.Links = nil
	return m
}

func (s *ResourceTfStack) Clean() interface{} {
	s.Links = nil
	return s
}

func (resource *Resource) Clean() interface{} {
	return &Resource{
		Name: resource.Name,
	}
}

func (m *ResourceListTfModule) Clean() interface{} {
	rscs := make([]interface{}, len(m.Resources))
	for i, rsc := range m.Resources {
		rscs[i] = rsc.Clean()
	}
	return map[string]interface{}{
		"modules": rscs,
	}
}

func (list *ResourceList) Clean() interface{} {
	rscs := make([]interface{}, len(list.Resources))
	for i, rsc := range list.Resources {
		rscs[i] = rsc.Clean()
	}
	return rscs
}

func (m *ResponseListTfModules) Clean() interface{} {
	return m.Embedded.Clean()
}

func (se *ServerError) Clean() interface{} {
	return se
}
