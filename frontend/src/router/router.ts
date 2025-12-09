import {
  createWebHashHistory,
  createRouter,
  type RouteRecordRaw,
} from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: () => import("../views/Home.vue"),
  },
  {
    path: "/pdf-to-image",
    name: "PdfToImage",
    component: () => import("../views/PdfToImage.vue"),
  },
  {
    path: "/image-to-pdf",
    name: "ImageToPdf",
    component: () => import("../views/ImageToPdf.vue"),
  },
  {
    path: "/image-optimizer",
    name: "ImageOptimizer",
    component: () => import("../views/ImageOptimizer.vue"),
  },
  {
    path: "/pdf-optimizer",
    name: "PdfOptimizer",
    component: () => import("../views/PdfOptimizer.vue"),
  },
  {
    path: "/pdf-merger",
    name: "PdfMerger",
    component: () => import("../views/PdfMerger.vue"),
  },
  {
    path: "/extract-pages-pdf",
    name: "ExtractPagesPdf",
    component: () => import("../views/ExtractPagesPdf.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes,
  scrollBehavior(_, __, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    } else {
      return { top: 0 };
    }
  },
});

export default router;
