import Vue from "vue";
import VueRouter from "vue-router";
import Login from "../pages/SubDomain/Login.vue";
import Main from "../pages/SubDomain/Main.vue";
import Epic from "../pages/SubDomain/Epic/Epic.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Login,
  },
  { path: "*", component: Login },
  //{ path: "/default", component: Default },
  { path: "/login", component: Login },
  { path: "/main", component: Main },
  { path: "/epic", component: Epic },
];

Vue.filter('readMore', function (text, length, suffix) {
  return (text.length> length)?text.substring(0, length) + suffix : text;
});
// configure router
const router = new VueRouter({
  routes, // short for routes: routes
  linkActiveClass: "active"
});

export default router;

