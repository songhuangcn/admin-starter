import { RouteLocationNormalized, RouteRecordRaw } from 'vue-router';
import { useUserStore } from '@/store';

export default function usePermission() {
  const userStore = useUserStore();

  return {
    accessRouter(route: RouteLocationNormalized | RouteRecordRaw) {
      return (
        !route.meta?.requiresAuth ||
        !route.meta?.roles ||
        userStore.is_admin ||
        route.meta?.roles?.some((permission) =>
          userStore.permission_names?.includes(permission)
        )
      );
    },
  };
}
