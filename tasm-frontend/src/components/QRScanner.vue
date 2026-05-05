<template>
  <div class="qr-scanner-container relative overflow-hidden rounded-2xl bg-black aspect-square max-w-sm mx-auto shadow-2xl">
    <div id="qr-reader" class="w-full h-full"></div>
    
    <!-- Scanning Overlay -->
    <div class="absolute inset-0 pointer-events-none border-2 border-primary/50 m-12 rounded-xl animate-pulse">
      <div class="absolute top-0 left-0 w-8 h-8 border-t-4 border-l-4 border-primary"></div>
      <div class="absolute top-0 right-0 w-8 h-8 border-t-4 border-r-4 border-primary"></div>
      <div class="absolute bottom-0 left-0 w-8 h-8 border-b-4 border-l-4 border-primary"></div>
      <div class="absolute bottom-0 right-0 w-8 h-8 border-b-4 border-r-4 border-primary"></div>
      
      <!-- Scanning Line -->
      <div class="absolute top-0 left-0 right-0 h-1 bg-primary/80 shadow-[0_0_15px_rgba(113,44,0,0.8)] animate-scan"></div>
    </div>

    <div v-if="error" class="absolute inset-0 bg-black/80 flex items-center justify-center p-6 text-center">
      <div class="text-white">
        <span class="material-symbols-outlined text-error text-4xl mb-2">error</span>
        <p class="font-body text-sm">{{ error }}</p>
        <button @click="retry" class="mt-4 px-4 py-2 bg-primary text-white rounded-lg text-sm">Retry</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue';
import { Html5Qrcode } from 'html5-qrcode';

const props = defineProps<{
  fps?: number;
  qrbox?: number;
}>();

const emit = defineEmits<{
  (e: 'result', decodedText: string): void;
  (e: 'error', error: string): void;
}>();

const error = ref<string | null>(null);
let html5QrCode: Html5Qrcode | null = null;

const startScanner = async () => {
  try {
    error.value = null;
    html5QrCode = new Html5Qrcode("qr-reader");
    
    const config = { 
      fps: props.fps || 10, 
      qrbox: props.qrbox || 250,
      aspectRatio: 1.0
    };

    await html5QrCode.start(
      { facingMode: "environment" },
      config,
      (decodedText) => {
        emit('result', decodedText);
        stopScanner();
      },
      () => {
        // Soft error (frame scan failed), don't show to user
      }
    );
  } catch (err: any) {
    const msg = err.message || "Failed to start camera";
    error.value = msg;
    emit('error', msg);
  }
};

const stopScanner = async () => {
  if (html5QrCode && html5QrCode.isScanning) {
    try {
      await html5QrCode.stop();
      html5QrCode.clear();
    } catch (err) {
      console.error("Failed to stop scanner", err);
    }
  }
};

const retry = () => {
  stopScanner().then(startScanner);
};

onMounted(() => {
  startScanner();
});

onBeforeUnmount(() => {
  stopScanner();
});

defineExpose({
  stopScanner,
  startScanner
});
</script>

<style scoped>
@keyframes scan {
  0% { top: 0%; opacity: 0; }
  10% { opacity: 1; }
  90% { opacity: 1; }
  100% { top: 100%; opacity: 0; }
}

.animate-scan {
  animation: scan 2s linear infinite;
}

:deep(#qr-reader__scan_region) {
  display: flex;
  justify-content: center;
  align-items: center;
}

:deep(video) {
  width: 100% !important;
  height: 100% !important;
  object-fit: cover !important;
}
</style>
