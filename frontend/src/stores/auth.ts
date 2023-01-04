import type { BaseRecord } from "@/pocketbase";
import client from "@/pocketbase";
import { defineStore } from "pinia";
import { computed, ref } from "vue";

export interface User extends BaseRecord, UserSettingsDto {
  email: string;
  username: string;
  avatar: string;
  verified: boolean;
}

export interface CreateUserDto {
  username: string;
  email: string;
  password: string;
  passwordConfirm: string;
}

export interface UpdateUserDto {
  username: string;
  email: string;
  password?: string;
  passwordConfirm?: string;
  oldPassword?: string;
  /** Set to null to remove current avatar. */
  avatar?: File | null;
}

export interface LoginPayload {
  identity: string;
  password: string;
}

export interface UserSettingsDto {
  theme: "light" | "dark" | "";
  locale: string;
}

export const useAuth = defineStore("auth", () => {
  const isAuthenticated = computed(() => !!user.value);
  const user = ref(client.authStore.model as User | null);

  client.authStore.onChange((_, model) => {
    user.value = model as typeof user.value;
  });

  if (user.value)
    client
      .collection("users")
      .authRefresh()
      .catch(() => logout()); // logout if refresh fails

  const login = async ({ identity, password }: LoginPayload) => {
    await client.collection("users").authWithPassword(identity, password);
  };

  const logout = () => {
    client.authStore.clear();
  };

  const requestEmailVerification = async () => {
    if (!user.value) return;
    return await client
      .collection("users")
      .requestVerification(user.value?.email);
  };

  const createUser = async (dto: CreateUserDto) => {
    const record = await client
      .collection("users")
      .create<User>({ ...dto, emailVisibility: true });
    if (record.email) {
      await client.collection("users").requestVerification(record.email);
    }
  };

  const updateUser = async (dto: UpdateUserDto) => {
    if (!user.value?.id) {
      throw new Error("Unable to update user because you are not logged in");
    }

    const formData = new FormData();
    Object.entries(dto).forEach(([key, value]) => {
      if (key === "email") return;
      formData.set(key, value);
    });

    const record = await client
      .collection("users")
      .update<User>(user.value.id, formData);
    if (dto.email && record.email !== dto.email) {
      await client.collection("users").requestEmailChange(dto.email);
    }
  };

  const updateSettings = async (dto: Partial<UserSettingsDto>) => {
    if (!user.value?.id) {
      throw new Error("Unable to update user because you are not logged in");
    }
    await client.collection("users").update<User>(user.value.id, dto);
  };

  const deleteUser = async () => {
    if (!user.value) return;
    client.collection("users").delete(user.value.id);
  };

  return {
    isAuthenticated,
    login,
    logout,
    user,
    requestEmailVerification,
    createUser,
    updateUser,
    updateSettings,
    deleteUser,
  };
});
