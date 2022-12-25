import { createRouter, createWebHistory } from "vue-router";
import routes from "~pages";

export type NavigationItem = {
  destination: string;
  label: string;
};

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
