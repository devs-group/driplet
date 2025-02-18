import { ref, reactive, Ref } from 'vue';

type ToastType = 'success' | 'error' | 'warning' | 'info';

interface Toast {
  id: string;
  message: string;
  type: ToastType;
  title: string;
  duration: number;
  showProgress: boolean;
  progress: number;
}

interface ToastOptions {
  message?: string;
  type?: ToastType;
  title?: string;
  duration?: number;
  showProgress?: boolean;
}

type ToastMethod = (message: string, options?: Partial<Omit<ToastOptions, 'message' | 'type'>>) => string;

interface ToastMethods {
  success: ToastMethod;
  error: ToastMethod;
  warning: ToastMethod;
  info: ToastMethod;
}

interface UseToastReturn {
  toasts: Ref<Toast[]>;
  toast: ToastMethods;
  removeToast: (id: string) => void;
  setMaxToasts: (max: number) => void;
}

interface TimeoutRecord {
  [key: string]: number;
}

const toasts = ref<Toast[]>([]);
const toastTimeouts = reactive<TimeoutRecord>({});
const progressIntervals = reactive<TimeoutRecord>({});
const maxToasts = ref<number>(5);

export function useToast(): UseToastReturn {
  // Create toast
  function addToast({
    message = '',
    type = 'info',
    title = '',
    duration = 5000,
    showProgress = true
  }: ToastOptions): string {
    // Create a new toast
    const id = Date.now().toString();
    const newToast: Toast = {
      id,
      message,
      type,
      title,
      duration,
      showProgress,
      progress: 100
    };

    // Remove oldest toast if at max
    if (toasts.value.length >= maxToasts.value) {
      const oldestToast = toasts.value[0];
      removeToast(oldestToast.id);
    }

    // Add new toast
    toasts.value.push(newToast);

    // Set timeout to remove toast
    if (duration !== Infinity) {
      // Timeout to remove toast
      toastTimeouts[id] = setTimeout(() => {
        removeToast(id);
      }, duration);

      // Progress interval for animation
      if (showProgress) {
        const startTime = Date.now();
        const endTime = startTime + duration;

        progressIntervals[id] = setInterval(() => {
          const now = Date.now();
          const remaining = endTime - now;
          const percent = (remaining / duration) * 100;

          // Find the toast and update its progress
          const toastIndex = toasts.value.findIndex(t => t.id === id);
          if (toastIndex !== -1) {
            toasts.value[toastIndex].progress = Math.max(0, percent);
          }

          // Clear interval when progress reaches 0
          if (percent <= 0) {
            clearInterval(progressIntervals[id]);
          }
        }, 10);
      }
    }

    return id;
  }

  // Remove toast
  function removeToast(id: string): void {
    const index = toasts.value.findIndex(t => t.id === id);
    if (index !== -1) {
      // Clear timeout and interval
      clearTimeout(toastTimeouts[id]);
      clearInterval(progressIntervals[id]);
      delete toastTimeouts[id];
      delete progressIntervals[id];

      // Remove toast
      toasts.value.splice(index, 1);
    }
  }

  // Configure max toasts
  function setMaxToasts(max: number): void {
    maxToasts.value = max;
  }

  // Toast methods exposed
  const toast: ToastMethods = {
    success: (message, options = {}) => addToast({ message, type: 'success', ...options }),
    error: (message, options = {}) => addToast({ message, type: 'error', ...options }),
    warning: (message, options = {}) => addToast({ message, type: 'warning', ...options }),
    info: (message, options = {}) => addToast({ message, type: 'info', ...options })
  };

  return {
    toasts,
    toast,
    removeToast,
    setMaxToasts
  };
}

export const globalToast = useToast();
