import Home from './Form';
import Sub from './Sub';
import Sum from './Sum';

const routes = [
  { path: '/', component: <Home />, exact: true },
  { path: '/sum', component: <Sum /> },
  { path: '/sub', component: <Sub /> },
];

export default routes;