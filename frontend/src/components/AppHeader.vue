<script setup lang="ts">
import Dropdown from "@/components/AppHeaderDropdown.vue";
import type { NavigationItem } from "@/router";
import { useDark } from "@vueuse/core";
import { RouterLink } from "vue-router";

const isDark = useDark();

const navigationItems: NavigationItem[] = [
  { label: "Home", destination: "/" },
  { label: "Create", destination: "/create" },
  { label: "Calendar", destination: "/calendar" },
];
</script>

<template>
  <nav
    class="fixed z-10 flex h-16 w-full items-center justify-between border-b border-gray-300 bg-gray-200 px-20 dark:border-slate-800 dark:bg-slate-900"
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

    <div class="flex h-full items-center justify-evenly">
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
      <Dropdown />
    </div>
  </nav>
</template>
