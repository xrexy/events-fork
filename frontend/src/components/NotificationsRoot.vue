<script setup lang="ts">
import { useNotifications } from "@/stores/notifications";
import type { Type } from "@/stores/notifications";
const notiStore = useNotifications();

type Icon = {
  paths: string[];
  color?: Partial<{
    background: string;
    icon: string;
  }>;
};

const icons: { [key in Type]: Icon } = {
  error: {
    paths: [
      "M15.362 5.214A8.252 8.252 0 0112 21 8.25 8.25 0 016.038 7.048 8.287 8.287 0 009 9.6a8.983 8.983 0 013.361-6.867 8.21 8.21 0 003 2.48z",
      "M12 18a3.75 3.75 0 00.495-7.467 5.99 5.99 0 00-1.925 3.546 5.974 5.974 0 01-2.133-1A3.75 3.75 0 0012 18z",
    ],
    color: {
      background: "bg-red-500",
      icon: "text-red-200",
    },
  },
  info: {
    paths: [
      "M11.25 11.25l.041-.02a.75.75 0 011.063.852l-.708 2.836a.75.75 0 001.063.853l.041-.021M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-9-3.75h.008v.008H12V8.25z",
    ],
    color: {
      background: "bg-teal-500",
      icon: "text-teal-100",
    },
  },
  success: {
    paths: ["M4.5 12.75l6 6 9-13.5"],
    color: {
      background: "bg-emerald-500",
      icon: "text-emerald-200",
    },
  },
  warning: {
    paths: [
      "M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126zM12 15.75h.007v.008H12v-.008z",
    ],
    color: {
      background: "bg-amber-500",
      icon: "text-amber-100",
    },
  },
};
</script>

<template>
  <div class="fixed bottom-0 right-0 w-screen px-4 sm:mr-4 sm:w-fit sm:p-0">
    <div
      @click="() => notiStore.setAsRead(notification.id)"
      v-for="notification in notiStore.notifications"
      :key="notification.id"
      class="mb-4 flex h-16 w-full flex-row items-center justify-between rounded-lg border border-gray-300 bg-gray-100 py-12 px-4 shadow-lg hover:cursor-pointer sm:w-96"
    >
      <div class="flex w-[85%] max-w-max flex-row items-center">
        <div class="flex w-full flex-col">
          <p
            class="flex-nowrap overflow-x-hidden overflow-ellipsis whitespace-nowrap text-lg font-semibold"
          >
            {{ notification.title }}
          </p>
          <p class="text-sm text-gray-400" v-html="notification.body"></p>
        </div>
      </div>
      <div class="flex flex-row items-center">
        <button
          :class="
            'flex h-10 w-10 flex-row items-center justify-center rounded-full ' +
            (icons[notification.type]?.color?.background || 'bg-gray-200')
          "
          type="button"
          @click="notiStore.setAsRead(notification.id)"
        >
          <svg
            :class="
              'h-6 w-6 ' +
              (icons[notification.type]?.color?.icon || 'text-gray-600')
            "
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="1.65"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              v-for="path in icons[notification.type].paths"
              :key="path"
              :d="path"
            ></path>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>
