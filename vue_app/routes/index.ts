import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
const routes: Array<RouteRecordRaw> = [
	{
		// add routes here
		// path: "/",
		// name: "Home",
		// component: HomeView,
	}
];
		
const router = createRouter({
	history: createWebHashHistory(),
	routes
});
		
export default router;