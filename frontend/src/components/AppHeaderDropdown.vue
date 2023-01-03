<script lang="ts" setup>
import { useAuth } from "@/stores/auth";
import { useToggle } from "@vueuse/shared";
import { UserIcon, ChevronRightIcon } from "@heroicons/vue/24/solid";
import { useDark } from "@vueuse/core";

const [active, toggleActive] = useToggle(false);
const { isAuthenticated, user } = useAuth();

const isDark = useDark();
const toggleTheme = useToggle(isDark);

const sections: {
  label: string;
  items: {
    label: string;
    action?: () => void;
    redirect?: string | { destination: string; external: boolean };
  }[];
}[] = [
  {
    label: "Account",
    items: [
      {
        label: "Your Profile",
      },
      {
        label: "Toggle Theme",
        action: () => toggleTheme(),
      },
      {
        label: "Sign out",
      },
    ],
  },
  {
    label: "Support",
    items: [
      {
        label: "Discord Server",
        redirect: {
          destination: "https://discord.gg/VQhkKMq8dF",
          external: true,
        },
      },
      {
        label: "GitHub Repository",
        redirect: {
          destination: "https://github.com/xrexy/events-fork",
          external: true,
        },
      },
      {
        label: "Report a Bug",
        redirect: {
          destination: "https://github.com/xrexy/events-fork/issues",
          external: true,
        },
      },
    ],
  },
];
</script>

<template>
  <div
    @click="() => toggleActive()"
    class="flex items-center rounded-full border border-white p-2 hover:cursor-pointer"
  >
    <UserIcon class="h-5 w-5" />
  </div>

  <div
    v-show="active"
    class="fixed right-0 top-[63px] m-0 flex h-[calc(100vh-63px)] w-screen flex-col justify-evenly border-t border-gray-300 bg-gray-200 px-32 pt-4 pb-32 text-center dark:border-slate-800 dark:bg-slate-900"
  >
    <div
      class="flex flex-col items-center gap-12"
      v-for="{ items, label } in sections"
      :key="label"
    >
      <p
        class="text-4xl font-black uppercase tracking-[0.15em] text-gray-500 dark:text-white"
      >
        {{ label }}
      </p>
      <div>
        <div
          v-for="{ label, action, redirect } in items"
          :key="label"
          @click="() => !!action && action()"
        >
          <div
            class="flex w-full flex-row items-center justify-start gap-1 rounded-sm py-1 px-2 text-lg tracking-wider text-gray-400 transition-all hover:cursor-pointer hover:bg-slate-300 hover:text-gray-500 dark:text-gray-600 hover:dark:bg-slate-800 hover:dark:text-gray-500"
          >
            <ChevronRightIcon class="h-5 w-5" />
            <p v-if="!redirect">{{ label }}</p>
            <a
              v-if="
                redirect && typeof redirect === 'object' && redirect.external
              "
              :href="redirect.destination"
            >
              {{ label }}
            </a>
            <router-link
              v-if="
                redirect &&
                (typeof redirect === 'string' ||
                  (typeof redirect === 'object' && !redirect.external))
              "
              :to="
                typeof redirect === 'string' ? redirect : redirect.destination
              "
              >{{ label }}</router-link
            >
            <!-- if you're reading this... I'm sorry you had to read this -->
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<!--

      <div
        class="group flex w-auto cursor-pointer items-center justify-center rounded-full border-2 border-gray-400 hover:border-brand-600 hover:bg-gray-700/10 dark:border-white dark:hover:border-brand-400"
        @click="onPersonClick() "
      >
        <PersonIcon
          :scale="{ width: '2em', height: '2em' }"
          :classes="{
            path: 'fill-gray-400 dark:fill-white group-hover:fill-brand-500 dark:group-hover:fill-brand-400',
            svg: 'p-1',
          }"
        />
        <p
          class="select-none pr-2 font-semibold text-gray-400 group-hover:text-brand-500 dark:text-white dark:group-hover:text-brand-400"
          v-if="isAuthenticated && user"
        >
          {{ user.username }}
        </p>
      </div>

-->
