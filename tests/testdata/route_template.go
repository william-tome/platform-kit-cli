package testdata

const ROUTE_TEMPLATE = `import * as route_services from 'route_services_path'
	
/** @type {import('vue-router').RouteRecordRaw} */
export const route_method_name = {
	path: '/route_name',
	name: '',
	children: [
		{
			path: '',
			name: 'route_name',
			component: () => import('component_path'),
			props: {
				service_method_name: route_services.service_method_name
			},
			meta: {}
		},
	]
}
`
