import type { BaseRecord } from "@/pocketbase";
import client from "@/pocketbase";
import { defineStore } from "pinia";
import type { RecordSubscription } from "pocketbase";
import { ref } from "vue";

interface Notification extends BaseRecord {
  title: string;
  body: string;
  type: "info" | "success" | "warning" | "error";
}

export interface CreateNotificationDto {
  title: string;
  body: string;
}

const LOCAL_STORAGE_KEY = "notifications__read";
const actionsList = ["update", "create", "delete"] as const;

type Action = typeof actionsList[number];

export const useNotifications = defineStore("notifications", () => {
  const notifications = ref([] as Notification[]);
  let readNotifications: string[] = [];

  client
    .collection("notifications")
    .getFullList<Notification>()
    .then((list) => {
      const local = fetchLocal();
      notifications.value = list.filter((n) => !local.includes(n.id));
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
        add(notification);
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

  const setAsRead = (id: string) => {
    readNotifications.push(id);
    window.localStorage.setItem(
      LOCAL_STORAGE_KEY,
      JSON.stringify(readNotifications)
    );
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

  const add = (notification: Notification) => {
    notifications.value.push(notification);
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
