import DavinciMain from "../layout/dashboard/DavinciMain.vue";
// Vistas Comunes
//import NotFound from "../pages/Page404.vue";
//import Default from "../pages/Default.vue";
import Login from "../pages/Login.vue";

// Paginas Oficiales
import Dashboard from "../pages/Dashboard.vue";
import Design from "../pages/Design/Design.vue";
import Portfolio from "../pages/Porfolio/Portfolio.vue";
import Inventions from "../pages/Invention/Inventions.vue";

const routes = [
  {
    path: "/",
    component: DavinciMain,
    redirect: "/davinci",
    children: [
      {
        path: "dashboard",
        name: "Pizarra",
        component: Dashboard
      },
      {
        path: "/inventions",
        name: "Inventos",
        component: Inventions
      },
      {
        path: "/portfolio",
        name: "Portafolio",
        component: Portfolio
      },
      {
        path: "/design",
        name: "Diseño",
        component: Design
      }
    ]
  },
  { path: "*", component: Login },
  //{ path: "/default", component: Default },
  { path: "/login", component: Login }
];

/**
 * Asynchronously load view (Webpack Lazy loading compatible)
 * The specified component must be inside the Views folder
 * @param  {string} name  the filename (basename) of the view to load.
function view(name) {
   var res= require('../components/Dashboard/Views/' + name + '.vue');
   return res;
};**/

export default routes;
