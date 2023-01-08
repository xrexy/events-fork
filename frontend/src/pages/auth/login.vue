<script setup lang="ts">
import { useAuth } from "@/stores/auth";
import type { LoginPayload, YupSchemaTypeMap } from "@/typings";
import { computed, ref } from "vue";
import { useRouter } from "vue-router";
import { string, ValidationError } from "yup";

/** --- TYPES --- */
interface FormValidationError extends ValidationError {
  key: keyof LoginPayload;
  messages: string[];
}

type FormErrorMessages = { [key in keyof LoginPayload]?: string[] } & {
  login?: string[];
};

/** --- VARIABLES/HOOKS --- */
const { login, isAuthenticated, $subscribe } = useAuth();
const { push: redirect } = useRouter();

/** --- STATE --- */
const mockUserOnMount = computed(() => import.meta.env.DEV && !isAuthenticated);

const state = ref<LoginPayload>({
  identity: mockUserOnMount.value ? "mockuser" : "",
  password: mockUserOnMount.value ? "mockpassword" : "",
});

const errors = ref<FormErrorMessages>({
  identity: [],
  password: [],
  login: [],
});

const timeouts = ref<Record<keyof FormErrorMessages, number | undefined>>({
  identity: undefined,
  password: undefined,
  login: undefined,
});

const validationRules: YupSchemaTypeMap<LoginPayload> = {
  identity: string().min(2).required(),
  password: string().min(6).required(),
};

/** --- HELPERS --- */
const setErrors = (messages: FormErrorMessages) => {
  console.info("setting errors", messages);

  // clear all timeouts
  Object.keys(timeouts.value).forEach((key) => {
    clearTimeout(timeouts.value[key as keyof FormErrorMessages]);
  });

  // set errors
  errors.value = { ...messages, login: [] };

  // set timeouts to clear errors after 5 seconds
  Object.keys(messages).forEach((key) => {
    setError(
      key as keyof LoginPayload,
      messages[key as keyof FormErrorMessages]?.[0]
    );
  });
};

const setError = (key: keyof FormErrorMessages, message?: string) => {
  if (!message) {
    console.info("tried to set error for", key, "but message was empty");
    return;
  }

  console.info(`setting error for ${key} to "${message}"`);

  // clear timeout for the key if it exists
  if (key in timeouts.value) {
    clearTimeout(timeouts.value[key]);
  }

  // set error for the key
  errors.value[key] = [message];

  // set timeout to clear the error after 5 seconds
  timeouts.value[key] = window.setTimeout(() => {
    console.info("clearing error for", key);
    errors.value[key] = [];
  }, 5000);
};

const clearErrors = () => {
  // remove all timeouts
  Object.keys(timeouts.value).forEach((key) => {
    clearTimeout(timeouts.value[key as keyof FormErrorMessages]);
  });

  errors.value = {
    identity: [],
    password: [],
    login: [],
  };
};

const handleSubmit = (
  success: () => void,
  error: (messages: FormErrorMessages) => void
) => {
  // validate all fields
  Promise.all(
    Object.keys(validationRules).map((key) =>
      validationRules[key as keyof LoginPayload]
        .validate((state as any).value[key], { abortEarly: false })
        .catch((err: ValidationError) => {
          return { ...err, key } as FormValidationError;
        })
    )
  ).then((str) => {
    // filter out any errors that are not actually errors
    const errors = str.filter(
      (err) => err && typeof err !== "string"
    ) as FormValidationError[];

    if (errors.length > 0) {
      // reduce error messages to a single object containing just a list of messages
      error(
        errors.reduce((acc, err) => {
          acc[err.key] = err.errors;
          return acc;
        }, {} as FormErrorMessages)
      );
    } else {
      success();
    }
  });
};

/** --- METHODS --- */
const onSubmit = () => {
  handleSubmit(
    () => {
      $subscribe((mutation, state) => {
        if (
          mutation.type !== "direct" &&
          mutation.storeId == "auth" &&
          state.user
        ) {
          return;
        }
        clearErrors();
        redirect("/");
      }, {}); // automatically unsubscribed when component is unmounted
      login(state.value).catch((err) => {
        console.info(err.message);
        setError("login", "Invalid credentials");
      });
    },
    (messages) => {
      setErrors(messages);
    }
  );
};
</script>

<!-- please make an actual design uwu -->
<template>
  <form
    @submit.prevent="onSubmit"
    class="flex w-full flex-col items-center justify-center gap-4 pt-8"
  >
    <div class="flex w-64 flex-col gap-2">
      <label class="text-sm font-semibold text-gray-500" for="identity">
        Email/Username
      </label>
      <input
        class="w-full rounded-sm border border-gray-300 bg-transparent"
        v-model="state.identity"
        name="identity"
      />

      <!-- Will get cancer when I add tailwind classes, will change(if not lazy) <3 -->
      <div v-for="message in errors['identity']" :key="message">
        <p>{{ message }}</p>
      </div>

      <label class="text-sm font-semibold text-gray-500" for="password">
        Password
      </label>
      <input
        class="w-full rounded-sm border border-gray-300 bg-transparent"
        v-model="state.password"
        type="password"
        name="password"
      />
    </div>

    <!-- Will get cancer when I add tailwind classes, will change(if not lazy) <3 -->
    <div v-for="message in errors['password']" :key="message">
      <p>{{ message }}</p>
    </div>

    <button
      class="rounded-md bg-brand-400 px-3 py-1 font-semibold text-white disabled:cursor-not-allowed disabled:bg-brand-200"
    >
      Login
    </button>

    <!-- Will get cancer when I add tailwind classes, will change(if not lazy) <3 -->
    <div v-for="message in errors['login']" :key="message">
      <p>{{ message }}</p>
    </div>
  </form>
</template>
