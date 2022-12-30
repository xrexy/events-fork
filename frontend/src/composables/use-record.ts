import type { BaseRecord } from "@/pocketbase";
import client from "@/pocketbase";
import { useAuthStore } from "@/stores/auth";
import { ref } from "vue";

export const useRecord = <T extends BaseRecord, CreateDto extends object>(
  collection: string
) => {
  const record = ref<T>();
  const isFetching = ref(false);
  const authStore = useAuthStore();

  const fetchAll = () => {
    isFetching.value = true;
    try {
      return client.collection(collection).getFullList<T>(200, {
        sort: "-created",
        perPage: 4,
      });
    } finally {
      isFetching.value = false;
    }
  };

  const fetch = async (id: string) => {
    isFetching.value = true;
    try {
      record.value = await client.collection(collection).getOne<T>(id);
    } finally {
      isFetching.value = false;
    }
  };

  const create = async (dto: CreateDto) => {
    if (!authStore.user) throw new Error("unauthenticated");
    record.value = await client
      .collection(collection)
      .create<T>({ ...dto, userId: authStore.user.id });
  };

  const update = async (dto: Partial<CreateDto>) => {
    if (!record.value) return;
    if (!record.value.id) return;

    record.value = await client
      .collection(collection)
      .update<T>(record.value.id, dto);
  };

  const remove = async () => {
    if (!record.value) return;
    return client.collection(collection).delete(record.value.id);
  };

  return {
    record,
    fetchAll,
    fetch,
    create,
    update,
    remove,
    isFetching,
  };
};
