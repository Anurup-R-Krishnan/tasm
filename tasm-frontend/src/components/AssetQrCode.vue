<template>
  <div
    class="p-6 bg-surface-subtle rounded-2xl flex flex-col items-center gap-4 border border-border-default"
  >
    <div
      class="w-32 h-32 bg-white rounded-xl flex items-center justify-center border border-border-default p-2 overflow-hidden"
    >
      <img v-if="qrDataUrl" :src="qrDataUrl" :alt="label" class="w-full h-full object-contain" />
      <span
        v-else
        class="text-[10px] text-text-secondary font-bold uppercase tracking-widest text-center px-2"
      >
        {{ loading ? 'Generating QR' : 'QR unavailable' }}
      </span>
    </div>
    <div class="text-center space-y-1">
      <p class="text-xs font-bold text-text-primary uppercase tracking-widest">{{ label }}</p>
      <p class="text-[10px] font-mono text-text-secondary break-all max-w-[220px]">{{ value }}</p>
    </div>
    <button
      type="button"
      class="text-[10px] font-bold text-primary uppercase tracking-widest hover:underline disabled:opacity-50"
      :disabled="!qrDataUrl"
      @click="downloadQr"
    >
      Download Asset Tag
    </button>
    <p v-if="error" class="text-[10px] font-bold text-status-critical text-center">
      {{ error }}
    </p>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import QRCode from 'qrcode';

const props = defineProps<{
  value: string;
  label?: string;
}>();

const qrDataUrl = ref('');
const loading = ref(false);
const error = ref('');

const label = computed(() => props.label || 'Asset QR Code');

const generateQr = async () => {
  const text = props.value.trim();
  if (!text) {
    qrDataUrl.value = '';
    error.value = 'No tag ID available';
    return;
  }

  loading.value = true;
  error.value = '';

  try {
    qrDataUrl.value = await QRCode.toDataURL(text, {
      errorCorrectionLevel: 'M',
      margin: 1,
      width: 512,
      color: {
        dark: '#0f172a',
        light: '#ffffff',
      },
    });
  } catch (err) {
    console.error('Failed to generate QR code:', err);
    qrDataUrl.value = '';
    error.value = 'Failed to generate QR code';
  } finally {
    loading.value = false;
  }
};

const downloadQr = () => {
  if (!qrDataUrl.value) return;

  const anchor = document.createElement('a');
  anchor.href = qrDataUrl.value;
  anchor.download = `${props.value || 'asset-tag'}.png`;
  anchor.click();
};

onMounted(generateQr);
watch(() => props.value, generateQr);
</script>
