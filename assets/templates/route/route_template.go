package route

import (
	"fmt"
	"text/template"
)

type RouteTemplateFields struct {
	RouteName         string
	RouteServices     string
	RouteServicesPath string
	RouteMethodName   string
	ComponentPath     string
	ServiceMethodName string
}

func GenerateRouteTemplate() (*template.Template, error) {
	tmpl, err := template.New("index").Parse(`import * as {{ .RouteServices }} from '{{ .RouteServicesPath }}'
	
/** @type {import('vue-router').RouteRecordRaw} */
export const {{ .RouteMethodName }} = {
	path: '/{{ .RouteName }}',
	name: '',
	children: [
		{
			path: '',
			name: '{{ .RouteName }}',
			component: () => import('{{ .ComponentPath }}'),
			props: {
				{{ .ServiceMethodName }}: {{ .RouteServices }}.{{ .ServiceMethodName }}
			},
			meta: {}
		},
	]
}
`)

	if err != nil {
		return nil, fmt.Errorf("route template: %s", err.Error())
	}

	return tmpl, nil
}
