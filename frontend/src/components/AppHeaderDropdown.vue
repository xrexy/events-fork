<script lang="ts" setup>
import { useAuth } from "@/stores/auth";
import { ChevronRightIcon, UserIcon } from "@heroicons/vue/24/solid";
import { onClickOutside, useDark } from "@vueuse/core";
import { useToggle } from "@vueuse/shared";
import { ref } from "vue";

const [active, toggleActive] = useToggle(false);
const { isAuthenticated, user, logout } = useAuth();

const isDark = useDark();
const toggleTheme = useToggle(isDark);

const dropdown = ref<HTMLDivElement>();
const root = ref<HTMLDivElement>();

onClickOutside(
  dropdown,
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  (_) => {
    toggleActive(false);
  },
  { ignore: [root] }
);

const sections: {
  label: string;
  items: {
    label: string;
    action?: () => void;
    redirect?: string | { destination: string; external: boolean };
  }[];
}[] = [
  {
    label: `Account${" " + (user?.username || "")}`,
    items: [
      isAuthenticated
        ? {
            label: "Your Profile",
          }
        : {
            label: "Sign in",
            redirect: "/login",
          },
      {
        label: "Toggle Theme",
        action: () => toggleTheme(),
      },
      isAuthenticated
        ? {
            label: "Sign Out",
            action: () => logout(),
          }
        : {
            label: "Close Menu",
            // action: () => toggleActive(false), // This is the default action
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
    ref="root"
    @click="() => toggleActive()"
    class="flex items-center rounded-full border border-gray-400 p-2 hover:cursor-pointer dark:border-white"
  >
    <UserIcon ref="root__icon" class="h-5 w-5" />
  </div>

  <div
    ref="dropdown"
    v-show="active"
    class="fixed right-0 top-[63px] flex h-[calc(100vh-63px)] w-screen flex-col justify-evenly border-t border-gray-300 bg-gray-200 px-32 pt-4 pb-32 text-center dark:border-slate-800 dark:bg-slate-900 sm:absolute sm:h-fit sm:w-fit sm:gap-4 sm:border sm:px-4 sm:py-4"
  >
    <div
      class="flex h-full flex-col items-center gap-12 sm:w-full sm:gap-2"
      v-for="{ items, label } in sections"
      :key="label"
    >
      <p
        class="w-full text-4xl font-black uppercase tracking-[0.15em] text-gray-600 dark:text-white sm:text-xl"
      >
        {{ label }}
      </p>
      <div @click="toggleActive(false)" class="w-full">
        <div
          v-for="{ label, action, redirect } in items"
          :key="label"
          @click="() => !!action && action()"
        >
          <div
            class="flex w-full flex-row items-center justify-start gap-1 rounded-sm py-1 px-2 text-lg tracking-wider text-gray-400 transition-all hover:cursor-pointer hover:bg-slate-300 hover:text-gray-500 dark:text-gray-600 hover:dark:bg-slate-800 hover:dark:text-gray-500 sm:text-sm"
          >
            <ChevronRightIcon class="h-5 w-5" />
            <p v-if="!redirect">{{ label }}</p>
            <a
              target="_blank"
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
