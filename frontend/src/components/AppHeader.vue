<script setup lang="ts">
import type { NavigationItem } from "@/router";
import { RouterLink } from "vue-router";
import PersonIcon from "./icons/MaterialSymbolsPerson.vue";
import { useDark, useToggle } from "@vueuse/core";
import { useAuth } from "@/stores/auth";

const isDark = useDark();
const toggleTheme = useToggle(isDark);
const { isAuthenticated, user } = useAuth();

const navigationItems: NavigationItem[] = [
  { label: "Home", destination: "/" },
  { label: "Create", destination: "/create" },
  { label: "Calendar", destination: "/calendar" },
];
``;
const onPersonClick = () => {
  console.log("Person clicked");
  toggleTheme(); // temporary
};
</script>

<template>
  <nav
    class="fixed z-10 flex h-16 w-full items-center justify-between border-b border-gray-300 bg-gray-200 px-32 dark:border-slate-800 dark:bg-slate-900"
  >
    <router-link to="/">
      <img
        v-if="isDark"
        class="h-auto w-24"
        src="../assets/Schwarz_Logo.svg"
        alt="Logo"
      />
      <img
        v-else
        class="h-auto w-24"
        src="../assets/Schwarz_Logo_Dark.svg"
        alt="Logo"
      />
    </router-link>

    <div class="flex items-center justify-evenly">
      <ul
        class="mr-[14px] flex gap-4 text-base font-semibold text-gray-500 dark:text-white"
        v-for="{ destination, label } in navigationItems"
        v-bind:key="destination"
      >
        <li
          class="hover:cursor-pointer hover:text-brand-500 dark:hover:text-brand-400"
        >
          <router-link
            :to="destination"
            active-class="text-brand-300 hover:text-brand-400"
          >
            {{ label }}
          </router-link>
        </li>
      </ul>
      <div
        class="group flex w-auto cursor-pointer items-center justify-center rounded-full border-2 border-gray-400 hover:border-brand-600 hover:bg-gray-700/10 dark:border-white dark:hover:border-brand-400"
        @click="onPersonClick()"
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
    </div>
  </nav>
</template>
