import type { BaseRecord } from "@/pocketbase";
import client from "@/pocketbase";
import { defineStore } from "pinia";
import type { RecordSubscription } from "pocketbase";
import { ref } from "vue";

export interface Notification extends BaseRecord {
  title: string;
  body: string;
  type: Type;
  timeout?: number;
}

const LOCAL_STORAGE_KEY = "notifications__read";
const MAXIMUM_NOTIFICATIONS = 2;
const NOTIFICATION_TIMEOUT = 10000000;
const actionsList = ["update", "create", "delete"] as const;
const types = ["info", "success", "warning", "error"] as const;

export type Action = typeof actionsList[number];
export type Type = typeof types[number];

export const useNotifications = defineStore("notifications", () => {
  const notifications = ref([] as Notification[]);
  let readNotifications: string[] = [];
  const timeouts: Record<string, number> = {};

  client
    .collection("notifications")
    .getFullList<Notification>()
    .then((list) => {
      const local = fetchLocal();
      list.filter((n) => !local.includes(n.id)).forEach(internalAdd);
    });

  const handleUpdate: (data: RecordSubscription<Notification>) => void = ({
    action: _action,
    record: notification,
  }) => {
    const action = _action as Action;

    switch (action) {
      case "update":
        update(notification);
        break;
      case "create":
        internalAdd(notification);
        break;
      case "delete":
        setAsRead(notification.id);
        break;
      default:
        console.warn("Unknown action", action);
        break;
    }
  };

  const unsubscribeFn = client
    .collection("notifications")
    .subscribe<Notification>("*", handleUpdate);

  const unsubscribe = async () => {
    (await unsubscribeFn)();
  };

  const setAsRead = (
    id: string,
    optins?: { addToStorage: boolean } = { addToStorage: true }
  ) => {
    if (timeouts[id]) window.clearTimeout(timeouts[id]);

    if (optins?.addToStorage) {
      readNotifications.push(id);
      window.localStorage.setItem(
        LOCAL_STORAGE_KEY,
        JSON.stringify(readNotifications)
      );
    }
    notifications.value = notifications.value.filter((n) => n.id !== id);
  };

  const fetchLocal = () => {
    const localNotiStorage = window.localStorage.getItem(LOCAL_STORAGE_KEY);

    if (!localNotiStorage) {
      return [];
    } else {
      readNotifications = JSON.parse(localNotiStorage);
      return readNotifications;
    }
  };

  const internalAdd = (notification: Notification) => {
    notifications.value.push(notification);
    timeouts[notification.id] = window.setTimeout(() => {
      console.info(`Notification ${notification.id} timed out`);
      setAsRead(notification.id);
    }, notification.timeout || NOTIFICATION_TIMEOUT);
  };

  const add = (notification: Notification, options?: { replace: boolean }) => {
    if (notifications.value.length >= MAXIMUM_NOTIFICATIONS) {
      console.log(
        `Maximum notifications reached ${
          options?.replace ? "(replacing first)" : "(ignoring)"
        }`
      );
      if (options?.replace) {
        const first = notifications.value.shift();
        if (first) setAsRead(first.id, { addToStorage: false });

        internalAdd(notification);
      }
    } else {
      internalAdd(notification);
    }
  };

  const update = (notification: Notification) => {
    notifications.value = notifications.value.map((n) =>
      n.id === notification.id ? notification : n
    );
  };

  return {
    setAsRead,
    add,
    update,
    unsubscribe,
    notifications,
  };
});
