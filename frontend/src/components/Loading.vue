<template>
    <Teleport to="body">
        <Transition name="fade">
            <div
                v-if="visible"
                class="fixed z-50 inset-0 flex flex-col items-center justify-center bg-background-dark/70 backdrop-blur-sm"
            >
                <div class="flex flex-col items-center gap-4">
                    <!-- Spinner -->
                    <div class="relative">
                        <div
                            class="w-16 h-16 rounded-full border-4 border-primary/20 border-t-primary animate-spin"
                        ></div>
                        <div
                            class="absolute inset-0 w-16 h-16 rounded-full border-4 border-transparent border-b-primary/40 animate-spin"
                            style="animation-direction: reverse; animation-duration: 1.5s"
                        ></div>
                    </div>

                    <!-- Text -->
                    <div class="flex flex-col items-center gap-1">
                        <p class="text-text-light font-semibold text-lg">{{ message }}</p>
                        <p class="text-subtle-text-dark text-sm">Please wait...</p>
                    </div>

                    <!-- Dots animation -->
                    <div class="flex gap-1.5">
                        <span
                            class="w-2 h-2 rounded-full bg-primary animate-bounce"
                            style="animation-delay: 0ms"
                        ></span>
                        <span
                            class="w-2 h-2 rounded-full bg-primary animate-bounce"
                            style="animation-delay: 150ms"
                        ></span>
                        <span
                            class="w-2 h-2 rounded-full bg-primary animate-bounce"
                            style="animation-delay: 300ms"
                        ></span>
                    </div>
                </div>
            </div>
        </Transition>
    </Teleport>
</template>

<script lang="ts" setup>
withDefaults(
    defineProps<{
        visible?: boolean;
        message?: string;
    }>(),
    {
        visible: true,
        message: "Processing",
    },
);
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}
</style>
